// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/topfreegames/maestro/api"
)

var bind string
var port int

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts maestro",
	Long:  `starts maestro api`,
	Run: func(cmd *cobra.Command, args []string) {
		ll := logrus.InfoLevel
		switch Verbose {
		case 0:
			ll = logrus.InfoLevel
		case 1:
			ll = logrus.WarnLevel
		case 3:
			ll = logrus.DebugLevel
		}

		var log = logrus.New()
		if json {
			log.Formatter = new(logrus.JSONFormatter)
		}
		log.Level = ll

		cmdL := log.WithFields(logrus.Fields{
			"source":    "startCmd",
			"operation": "Run",
			"bind":      bind,
			"port":      port,
		})

		cmdL.Info("starting maestro")

		app, err := api.NewApp(bind, port, config, log)
		if err != nil {
			cmdL.Fatal(err)
		}

		app.ListenAndServe()

	},
}

func init() {
	startCmd.Flags().StringVarP(&bind, "bind", "b", "0.0.0.0", "bind address")
	startCmd.Flags().IntVarP(&port, "port", "p", 8080, "bind port")
	RootCmd.AddCommand(startCmd)
}
