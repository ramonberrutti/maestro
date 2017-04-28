// Code generated by go-bindata.
// sources:
// migrations/0001-CreteConfigTable.sql
// migrations/0002-CreteRoomsTable.sql
// DO NOT EDIT!

package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _migrations0001CreteconfigtableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xcd\x6e\xd3\x40\x14\x85\xf7\x7e\x8a\xa3\xae\x12\x09\xc7\x50\x51\x21\xa5\x08\x11\x9a\x09\x8c\x70\xc6\x90\x8c\x55\x97\x8d\x35\xb5\x6f\xec\x11\xb5\xc7\xcc\x4f\xa3\x3e\x12\xaf\xc1\x93\x21\xb7\x4d\x22\x75\xd3\xe5\x9c\x7b\xce\x9d\xef\x9e\x38\x46\xa7\xc8\x79\x6b\xa2\x38\x46\xeb\xfd\xe0\xe6\x49\xd2\x68\xdf\x86\xdb\x59\x65\xba\xc4\x9b\x61\x67\x89\x1a\xd5\x91\x4b\x4e\xd6\xd1\x9d\xea\x8a\x7a\x47\x35\x42\x5f\x93\x85\x6f\x09\x6b\x2e\x71\xf7\x24\xcf\x0f\x0b\xe7\x49\xb2\xdf\xef\x67\x66\xa0\xde\x99\x60\x2b\x9a\x19\xdb\x24\xcf\x2e\x97\x74\xda\xc7\xcf\x8f\x31\x71\x65\x86\x07\xab\x9b\xd6\xe3\xdf\x5f\x9c\xbf\x7d\xf7\x01\xd2\x0c\x58\x59\x22\x7c\x1d\x19\xf0\xf1\x56\x55\xbf\xa9\xaf\x3f\xfb\x5d\x53\x99\x91\xf1\x53\x14\x5d\x6d\xd8\x42\x32\xb0\x42\x32\xb1\xe5\x99\x00\x5f\x41\x64\x12\xac\xe0\x5b\xb9\xc5\x59\x08\xba\x8e\x8d\x73\xc3\xd9\xe5\xd1\x2c\x17\x5f\x52\x86\xca\xf4\x3b\xdd\x38\x4c\x22\x00\xd0\x35\xf2\x9c\x2f\xf1\x63\xc3\xd7\x8b\xcd\x0d\xbe\xb3\x1b\x2c\xd9\x6a\x91\xa7\x12\xe3\x8e\xb2\xa1\x9e\xac\xf2\x54\xde\xbf\x9f\x4c\xdf\x3c\x66\x7a\xd5\x11\xee\x95\xad\x5a\x65\x27\xe7\x17\x17\xd3\xc7\x9f\x45\x9e\xa6\x4f\xf3\xe6\x95\xf9\x83\xea\xee\x20\x59\x21\x5f\xe8\x95\x25\xe5\xa9\x2e\x95\x87\xd7\x1d\x39\xaf\xba\x01\xd7\x5c\x7e\x83\xe4\x6b\x86\x5f\x99\x60\xc7\xc4\x11\x52\x64\xd7\x07\xae\x30\xd4\xaf\xe7\xf3\x34\x8d\xa6\xa7\x4e\x72\xc1\x7f\xe6\x0c\x5c\x2c\x59\x71\xa8\xa6\x1c\x0f\x2c\x43\xaf\xff\x04\x42\x26\x4e\x8d\x8d\xfa\xf4\x32\xfa\x1f\x00\x00\xff\xff\xfe\xa6\x90\x9c\x43\x02\x00\x00")

func migrations0001CreteconfigtableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0001CreteconfigtableSql,
		"migrations/0001-CreteConfigTable.sql",
	)
}

func migrations0001CreteconfigtableSql() (*asset, error) {
	bytes, err := migrations0001CreteconfigtableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0001-CreteConfigTable.sql", size: 579, mode: os.FileMode(420), modTime: time.Unix(1492543985, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations0002CreteroomstableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xd1\x6a\xdb\x30\x18\x85\xef\xfd\x14\x87\x5e\xc5\x30\xc7\x5b\xa1\x0c\xd2\x31\xe6\xc5\xca\x66\x70\x65\x58\x64\x1a\x76\x13\x54\xe9\x8f\x2d\x56\x5b\x42\x92\x17\xf6\x48\x7b\x8d\x3d\xd9\x48\xe3\x36\xbb\xea\x2e\xf5\x9f\xef\x1c\x1d\x4e\x96\x61\x90\x14\xa2\xb7\x49\x96\xa1\x8f\xd1\x85\x55\x9e\x77\x26\xf6\xd3\xc3\x52\xd9\x21\x8f\xd6\x1d\x3c\x51\x27\x07\x0a\xf9\x05\x3d\xd1\xb5\x51\x34\x06\xd2\x98\x46\x4d\x1e\xb1\x27\xdc\x55\x02\x8f\xe7\xf3\xea\x39\x70\x95\xe7\xc7\xe3\x71\x69\x1d\x8d\xc1\x4e\x5e\xd1\xd2\xfa\x2e\x9f\xa9\x90\x0f\x26\x66\xf3\xe3\xe4\x58\x5b\xf7\xcb\x9b\xae\x8f\xf8\xf3\x1b\xd7\x6f\xdf\xbd\x87\xb0\x0e\x1b\x4f\x84\x2f\xa7\x0e\xf8\xf0\x20\xd5\x0f\x1a\xf5\xa7\x78\xe8\x94\x3d\x75\xfc\x98\x24\xeb\x6f\xac\x10\x0c\x6c\x27\x18\xdf\x56\x0d\x47\xb5\x01\x6f\x04\xd8\xae\xda\x8a\x2d\xae\xa6\xc9\xe8\xcc\x86\xe0\xae\x6e\x5f\x60\x51\x7c\xae\x19\xbc\xb5\x43\xc0\x22\x01\x00\xa3\xf1\x53\x7a\xd5\x4b\xbf\xb8\xbe\xb9\x49\x9f\x12\x78\x5b\xd7\x6f\x9e\x54\x65\xc7\x83\xe9\xf6\x46\xa3\x6d\xab\x12\x9e\x0e\xe4\x69\x54\x14\x66\x25\x2c\x8c\x4e\xd1\x70\x94\xac\x66\x82\x61\x5d\x6c\xd7\x45\xc9\xce\xe6\x10\x65\x9c\xc2\x6b\xf1\x8f\x32\xc4\xbd\x33\x63\xb7\x97\x11\xd1\x0c\x14\xa2\x1c\x1c\xee\x2b\xf1\x15\xa2\xba\x63\xf8\xde\x70\xf6\x6f\x1d\x4f\x32\x92\x7e\x9d\x9e\x7f\x40\xc9\x36\x45\x5b\x0b\xf0\xe6\x7e\x91\x9e\xfd\x93\xd3\xff\xf7\xb7\x75\x9d\xa4\x97\xc9\x2a\x5e\xb2\xdd\x79\xb2\xfd\x65\x8e\x86\x3f\xaf\xf8\x72\x4b\x6f\x93\xbf\x01\x00\x00\xff\xff\xa2\xc3\x3e\x51\x5a\x02\x00\x00")

func migrations0002CreteroomstableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0002CreteroomstableSql,
		"migrations/0002-CreteRoomsTable.sql",
	)
}

func migrations0002CreteroomstableSql() (*asset, error) {
	bytes, err := migrations0002CreteroomstableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0002-CreteRoomsTable.sql", size: 602, mode: os.FileMode(420), modTime: time.Unix(1493386725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"migrations/0001-CreteConfigTable.sql": migrations0001CreteconfigtableSql,
	"migrations/0002-CreteRoomsTable.sql": migrations0002CreteroomstableSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"migrations": &bintree{nil, map[string]*bintree{
		"0001-CreteConfigTable.sql": &bintree{migrations0001CreteconfigtableSql, map[string]*bintree{}},
		"0002-CreteRoomsTable.sql": &bintree{migrations0002CreteroomstableSql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

