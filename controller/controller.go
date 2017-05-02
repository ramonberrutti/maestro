// maestro
// https://github.com/topfreegames/maestro
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package controller

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
	pginterfaces "github.com/topfreegames/extensions/pg/interfaces"
	redisinterfaces "github.com/topfreegames/extensions/redis/interfaces"
	"github.com/topfreegames/maestro/models"
	yaml "gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

// CreateScheduler creates a new scheduler from a yaml configuration
func CreateScheduler(logger logrus.FieldLogger, mr *models.MixedMetricsReporter, db pginterfaces.DB, redisClient redisinterfaces.RedisClient, clientset kubernetes.Interface, configYAML *models.ConfigYAML) error {
	// TODO only return when cluster is ready
	// TODO on case of error rollback everything
	configBytes, err := yaml.Marshal(configYAML)
	if err != nil {
		return err
	}
	yamlString := string(configBytes)
	config := models.NewConfig(configYAML.Name, configYAML.Game, yamlString)
	err = mr.WithSegment(models.SegmentInsert, func() error {
		return config.Create(db)
	})
	if err != nil {
		return err
	}

	namespace := models.NewNamespace(config.Name)
	err = mr.WithSegment(models.SegmentNamespace, func() error {
		return namespace.Create(clientset)
	})

	if err != nil {
		deleteErr := deleteSchedulerHelper(logger, mr, db, clientset, config, namespace)
		if deleteErr != nil {
			return deleteErr
		}
		return err
	}

	// TODO: optimize creation (in parallel?)
	// TODO nao esta removendo do banco
	// TODO pegar timeout da config
	err = ScaleUp(logger, mr, db, redisClient, clientset, config.Name, configYAML.AutoScaling.Min, 300)
	return err
}

// DeleteScheduler deletes a scheduler from a yaml configuration
func DeleteScheduler(logger logrus.FieldLogger, mr *models.MixedMetricsReporter, db pginterfaces.DB, clientset kubernetes.Interface, configName string) error {
	config := models.NewConfig(configName, "", "")
	err := mr.WithSegment(models.SegmentSelect, func() error {
		return config.Load(db)
	})
	if err != nil {
		return err
	}
	namespace := models.NewNamespace(config.Name)
	return deleteSchedulerHelper(logger, mr, db, clientset, config, namespace)
}

// GetSchedulerScalingInfo returns the scheduler scaling policies and room count by status
func GetSchedulerScalingInfo(logger logrus.FieldLogger, mr *models.MixedMetricsReporter, db pginterfaces.DB, client redisinterfaces.RedisClient, configName string) (*models.AutoScaling, *models.RoomsStatusCount, error) {
	config := models.NewConfig(configName, "", "")
	err := mr.WithSegment(models.SegmentSelect, func() error {
		return config.Load(db)
	})
	if err != nil {
		return nil, nil, err
	}
	scalingPolicy := config.GetAutoScalingPolicy()
	var roomCountByStatus *models.RoomsStatusCount
	err = mr.WithSegment(models.SegmentGroupBy, func() error {
		roomCountByStatus, err = models.GetRoomsCountByStatus(client, config.Name)
		return err
	})
	if err != nil {
		return nil, nil, err
	}
	return scalingPolicy, roomCountByStatus, nil
}

// GetSchedulerStateInfo returns the scheduler state information
func GetSchedulerStateInfo(logger logrus.FieldLogger, mr *models.MixedMetricsReporter, client redisinterfaces.RedisClient, configName string) (*models.SchedulerState, error) {
	state := models.NewSchedulerState(configName, "", 0, 0)
	err := mr.WithSegment(models.SegmentHGetAll, func() error {
		return state.Load(client)
	})
	if err != nil {
		return nil, err
	}
	return state, nil
}

// SaveSchedulerStateInfo updates the scheduler state information
func SaveSchedulerStateInfo(logger logrus.FieldLogger, mr *models.MixedMetricsReporter, client redisinterfaces.RedisClient, state *models.SchedulerState) error {
	return mr.WithSegment(models.SegmentHMSet, func() error {
		return state.Save(client)
	})
}

// ScaleUp scales up a scheduler using its config
func ScaleUp(logger logrus.FieldLogger, mr *models.MixedMetricsReporter, db pginterfaces.DB, redisClient redisinterfaces.RedisClient, clientset kubernetes.Interface, configName string, amount, timeoutSec int) error {
	l := logger.WithFields(logrus.Fields{
		"source":    "scaleUp",
		"scheduler": configName,
		"amount":    amount,
	})
	l.Info("scaling scheduler up")
	config := models.NewConfig(configName, "", "")
	err := mr.WithSegment(models.SegmentSelect, func() error {
		return config.Load(db)
	})
	if err != nil {
		return err
	}
	configYAML, _ := models.NewConfigYAML(config.YAML)

	timeout := make(chan bool, 1)
	pods := make([]string, amount)
	go func() {
		time.Sleep(time.Duration(timeoutSec) * time.Second)
		timeout <- true
	}()
	for i := 0; i < amount; i++ {
		podName, err := createServiceAndPod(l, mr, redisClient, clientset, configYAML, config.Name)
		if err != nil {
			//TODO logica quando da erro
			l.WithError(err).Error("scale up error")
		}
		pods[i] = podName
	}
	select {
	case <-timeout:
		return errors.New("timeout scaling up scheduler")
	default:
		exit := true
		for i := 0; i < amount; i++ {
			pod, err := clientset.CoreV1().Pods(config.Name).Get(pods[i], metav1.GetOptions{})
			if err != nil {
				l.WithError(err).Error("scale up pod error")
			}
			for _, containerStatus := range pod.Status.ContainerStatuses {
				if !containerStatus.Ready {
					exit = false
				}
			}
			if exit {
				l.Info("finished scaling up scheduler")
				break
			}
		}
		l.Debug("scaling up scheduler...")
		time.Sleep(time.Duration(1) * time.Second)
	}
	// TODO: set SchedulerState.State back to "in-sync"
	return nil
}

func deleteSchedulerHelper(logger logrus.FieldLogger, mr *models.MixedMetricsReporter, db pginterfaces.DB, clientset kubernetes.Interface, config *models.Config, namespace *models.Namespace) error {
	err := mr.WithSegment(models.SegmentNamespace, func() error {
		return namespace.Delete(clientset) // TODO: we're assuming that deleting a namespace gracefully terminates all its pods
	})
	if err != nil {
		logger.WithError(err).Error("Failed to delete namespace while rolling back cluster creation.")
		return err
	}
	err = mr.WithSegment(models.SegmentDelete, func() error {
		return config.Delete(db)
	})
	if err != nil {
		logger.WithError(err).Error("Failed to delete config while rolling back cluster creation.")
		return err
	}
	return nil
}

func buildNodePortEnvVars(ports []v1.ServicePort) []*models.EnvVar {
	nodePortEnvVars := make([]*models.EnvVar, len(ports))
	for idx, port := range ports {
		nodePortEnvVars[idx] = &models.EnvVar{
			Name:  fmt.Sprintf("MAESTRO_NODE_PORT_%d_%s", port.Port, port.Protocol),
			Value: strconv.FormatInt(int64(port.NodePort), 10),
		}
	}
	return nodePortEnvVars
}

func createServiceAndPod(logger logrus.FieldLogger, mr *models.MixedMetricsReporter, redisClient redisinterfaces.RedisClient, clientset kubernetes.Interface, configYAML *models.ConfigYAML, configName string) (string, error) {
	randID := strings.SplitN(uuid.NewV4().String(), "-", 2)[0]
	name := fmt.Sprintf("%s-%s", configYAML.Name, randID)
	room := models.NewRoom(name, configName)
	err := mr.WithSegment(models.SegmentInsert, func() error {
		return room.Create(redisClient)
	})
	if err != nil {
		return "", err
	}
	service := models.NewService(name, configYAML.Name, configYAML.Ports)
	var kubeService *v1.Service
	err = mr.WithSegment(models.SegmentService, func() error {
		kubeService, err = service.Create(clientset)
		return err
	})
	if err != nil {
		return "", err
	}
	namesEnvVars := []*models.EnvVar{
		&models.EnvVar{
			Name:  "MAESTRO_SCHEDULER_NAME",
			Value: configYAML.Name,
		},
		&models.EnvVar{
			Name:  "MAESTRO_ROOM_ID",
			Value: name,
		},
	}
	env := append(configYAML.Env, namesEnvVars...)
	nodePortEnvVars := buildNodePortEnvVars(kubeService.Spec.Ports)
	env = append(env, nodePortEnvVars...)
	pod := models.NewPod(
		configYAML.Game,
		configYAML.Image,
		name,
		configYAML.Name,
		configYAML.Limits.CPU,
		configYAML.Limits.Memory,
		configYAML.Limits.CPU,    // TODO: requests should be < limits calculate it
		configYAML.Limits.Memory, // TODO: requests should be < limits calculate it
		configYAML.ShutdownTimeout,
		configYAML.Ports,
		configYAML.Cmd,
		env,
	)
	var kubePod *v1.Pod
	err = mr.WithSegment(models.SegmentPod, func() error {
		kubePod, err = pod.Create(clientset)
		return err
	})
	if err != nil {
		return "", err
	}
	nodeName := kubePod.Spec.NodeName
	logger.WithFields(logrus.Fields{
		"node": nodeName,
		"name": name,
	}).Info("Created GRU (service and pod) successfully.")
	return name, nil
}
