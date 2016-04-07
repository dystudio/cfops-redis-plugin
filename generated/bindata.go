// Code generated by go-bindata.
// sources:
// scripts/backupDedicated.sh
// scripts/backupShared.sh
// scripts/restoreDedicated.sh
// scripts/restoreShared.sh
// DO NOT EDIT!

package generated

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

var _scriptsBackupdedicatedSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x52\x31\x4f\xf3\x30\x10\xdd\xfd\x2b\x4e\x91\xd5\x7e\x19\x5c\x7f\x2d\x03\x53\x87\x22\x15\x84\xd4\xad\x88\x05\x50\x74\xb5\x4d\xb0\x9a\xd8\xc6\x76\x13\x89\xaa\xff\x9d\xa4\xa9\x12\xca\x00\x13\x95\x07\xfb\xd9\x7e\xef\xdd\x3b\xdd\xed\xfd\x6a\x39\x4f\x78\x85\x9e\x57\x02\x1d\x0f\xd1\x7a\xc5\xbd\x92\x3a\x24\xe4\xb8\x65\xc2\x96\x25\x1a\x39\xa7\xff\xb0\xde\xc2\xb8\x79\x34\x58\x2a\x76\xba\x86\x9b\xbb\xf5\xe2\x71\xc9\xf7\xce\x6b\x13\x81\x5e\x1d\xc6\x40\x5b\xd5\x4e\x64\x22\xac\x79\x4d\x4f\x4a\xce\xfa\xd8\xcb\xb4\xa0\x67\xcd\x7e\x62\xd5\x5f\xbd\xdf\x77\xda\x2b\x87\x21\xfc\xc2\x2d\x30\xc4\x2c\x60\xa5\x1a\xee\x10\xcf\xa1\xd8\x62\xae\x42\xf7\x95\x6f\xb4\xe9\x4e\x4c\x14\x1a\xd8\x1b\x4c\x67\xd7\x93\xff\xcd\x9a\x02\x73\x40\x87\xaa\x81\x61\x0f\x6b\x09\xab\xc5\xfa\xa1\x4d\x9d\x92\x4d\xfe\x67\x26\x74\x7f\xd6\x7f\xce\x9f\x13\x7e\x48\x89\x51\x75\x76\xa1\x70\x3b\x13\x75\x01\x4f\x40\xcf\x3c\x81\xe5\x4d\xd7\x07\xf8\x42\xa4\x25\x00\x17\xac\x4b\x5a\xa3\x88\x90\xf0\x7d\x6c\x61\x34\x82\x88\x1e\xc4\x07\x1c\x69\xe4\x33\x00\x00\xff\xff\xfd\xa4\xac\xd6\xde\x02\x00\x00")

func scriptsBackupdedicatedShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsBackupdedicatedSh,
		"scripts/backupDedicated.sh",
	)
}

func scriptsBackupdedicatedSh() (*asset, error) {
	bytes, err := scriptsBackupdedicatedShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/backupDedicated.sh", size: 734, mode: os.FileMode(420), modTime: time.Unix(1459951760, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptsBackupsharedSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x53\xcf\x4f\xc2\x30\x18\xbd\xf3\x57\xbc\x98\x05\xd4\xa4\x54\xf0\xe0\x81\x78\xc0\x04\x8d\x09\x37\x8c\x17\x24\xe4\xa3\xed\xa0\x01\xda\xd9\x16\x96\x48\xf8\xdf\x2d\x1b\x6e\x72\x92\x8b\xda\x5d\xbe\xb7\xed\xbd\xf7\xfd\xe8\x97\x5a\x87\xc7\xe7\xe1\x00\xda\x80\x6f\xc9\xf1\xad\xa0\x8c\xfb\x60\x9d\xe2\x22\x65\x4e\x49\xed\xd9\xcc\xd9\xa5\x72\xbc\x04\x92\x02\xf1\xeb\x1e\xa4\x6d\x20\x1e\x9d\x62\x3c\x06\x93\x48\x0a\x99\xc9\xa4\x87\xb0\x50\xa6\xf8\x16\x4f\xc1\x99\x0a\xbb\x5e\x93\x91\xf7\xc9\x25\xe5\x4b\xb4\xa2\x92\xa1\xb5\x62\xc7\xd7\x78\x78\x1a\xf5\x5f\x07\x7c\x97\x39\x6d\x02\x92\xdb\x7d\xab\x54\x2b\x1d\xdb\xc2\x9a\xf4\xea\x54\x30\xb3\x2e\x54\x6a\x07\x50\x91\xbb\x67\x90\xf3\xef\x99\xbc\x6f\xb4\x53\x19\x79\x7f\x9e\xc4\x8a\x7c\x98\x7a\xda\xaa\x28\x51\x37\x2c\x23\xb1\xa4\xb9\xf2\x25\x83\xcf\xb4\x39\x76\x4b\xac\x34\xd8\x02\x9d\xee\x5d\xfb\x26\x3e\x1d\xb0\x0c\x49\x5d\x03\x18\x55\x30\x97\x18\xf6\x47\x2f\x87\x56\x54\x6e\xb3\xf9\xaf\x79\x25\xbb\x93\xd9\x70\xfe\x76\xc1\xf7\x95\xb1\x51\xf9\xf4\x6f\x4b\xdd\x98\xa0\x57\x18\x23\x39\xb1\x06\x9b\xc7\x89\xd4\x70\xf2\xf5\xfb\xf1\xfa\xfd\x4f\xae\xd2\x1a\x55\xc4\xa9\x6e\x14\xb1\x90\x3f\x6e\x0f\x9a\x4d\x04\x72\x10\x1f\xa8\x17\xa9\xf1\x19\x00\x00\xff\xff\x8f\xd2\x82\x93\x80\x03\x00\x00")

func scriptsBackupsharedShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsBackupsharedSh,
		"scripts/backupShared.sh",
	)
}

func scriptsBackupsharedSh() (*asset, error) {
	bytes, err := scriptsBackupsharedShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/backupShared.sh", size: 896, mode: os.FileMode(420), modTime: time.Unix(1459891238, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptsRestorededicatedSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x54\x5f\x6f\xd3\x3e\x14\x7d\xf7\xa7\xb8\x4b\xab\xf5\xd7\x9f\xe4\xb8\xdd\xc3\x10\x82\x3e\x80\x18\x68\x4f\x4c\x05\x89\x07\x84\x22\xd7\xbe\x6d\xac\x3a\xb6\x67\x3b\x0d\x13\xf0\xdd\x71\xba\x6e\x6d\x46\xd7\x22\xe0\x61\xc9\x43\x74\x7d\xff\x1d\xe7\x9e\x7b\x7a\x27\x6c\xa6\x0c\x9b\xf1\x50\x92\xc8\xfd\x02\x63\x21\x95\x9f\xb0\x15\xf7\x6c\x25\xb8\x63\x21\x5a\x8f\xcc\xa3\x54\x81\x11\x21\xe1\xa1\x27\x56\xae\x98\x71\xb1\xac\x1d\x83\xd3\x53\x48\x35\xe0\xeb\x6a\x7e\x28\x6c\x5d\x8b\x46\xa5\x31\x4f\xd1\x04\x45\x69\x21\x9b\x62\x1b\xa7\xcc\x02\xd6\xee\x8c\xac\x3f\x85\xb3\x3e\x4e\xfa\xff\xf1\x66\x09\x03\xd6\x1a\xec\x9b\x4b\x51\x11\xfa\x67\x3f\x06\xd0\xdf\x22\xbe\xad\x9a\x0b\x6b\xe6\xc3\xbb\xdc\x46\xde\xa7\x7a\xbc\xae\x95\x47\xc7\x43\xf8\xad\x0a\x57\x97\x6f\x52\xae\xe0\xf1\xd7\x08\x1a\xd0\xaf\xd0\xe7\x4e\xc9\x21\x59\x2a\xad\x81\x3e\x87\x7e\x4a\x20\xbe\x02\xea\xe7\x07\x33\x48\xaf\x07\x01\x23\x44\x0b\xdc\x39\x34\xd2\x1a\x7d\x03\xc6\x92\x80\x12\xa8\x82\x41\x60\x3b\xe7\x37\xd8\x31\x8d\x65\x8b\xc7\x20\xef\x14\xf0\x68\x78\x85\x54\xd8\xaa\xe2\x46\xc2\xeb\x77\xd3\x8b\x4f\xd3\xcb\x8f\x17\xaf\xde\xbf\x85\x2c\x3b\xec\xde\xb5\xb2\x03\xdd\xf6\x5c\xf5\xff\x9c\xdb\x47\x1c\x5e\xce\x88\x70\x47\x39\xc1\x64\x5d\xb9\x36\xb8\x93\x9e\x13\x51\x56\x56\xc2\xf9\xf9\xa8\x73\x7e\x17\x4c\x48\x53\x26\x2e\xc1\x09\x18\x01\xf4\x1a\xc6\xa0\xad\xe0\xba\xb4\x21\x0d\x6f\xcb\x22\x78\xc9\x24\xae\x98\xa9\xb5\x7e\x01\xd2\x42\xd0\x88\x0e\xc6\xa3\xd6\x30\xb8\xe1\xfe\x95\xb7\x0b\x8f\x21\x7c\x88\x3c\xd6\x21\x31\x60\x0b\xd9\x25\xa0\x7c\x81\x61\x03\xb5\x5d\x9a\xdb\xe1\x0a\xad\x80\x96\x30\x3e\x7b\x96\x8f\xd2\x3b\x06\xea\x3a\x7d\x29\xbf\x37\x1b\x09\xca\xcc\x2d\x7c\x87\xd4\xc5\xc1\x20\xfd\xb0\xc2\x63\xe3\x55\xc4\x42\x99\xc2\x6d\xba\x0f\x86\xe4\x9f\xf7\xdd\x1d\x2c\x71\x4f\xe5\x9a\xb5\x49\x32\x00\x9f\x21\xeb\x77\x21\x65\x30\x99\xa4\xc3\x7d\x43\xc9\xe0\x0b\x91\x96\x40\x7a\x9e\xcc\x35\x5a\x30\x1b\x3e\x91\x35\x9b\xfe\x4c\x3f\xf6\x6a\x40\x5a\xfa\x07\x8a\xf0\xd7\x1a\xd0\x5d\xf2\x23\x72\x71\x54\x03\x8e\x6d\x35\xf9\x19\x00\x00\xff\xff\x30\x59\x7f\xe4\x63\x06\x00\x00")

func scriptsRestorededicatedShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsRestorededicatedSh,
		"scripts/restoreDedicated.sh",
	)
}

func scriptsRestorededicatedSh() (*asset, error) {
	bytes, err := scriptsRestorededicatedShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/restoreDedicated.sh", size: 1635, mode: os.FileMode(420), modTime: time.Unix(1460001364, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptsRestoresharedSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\x51\x4f\xdb\x30\x10\x7e\xef\xaf\x38\x42\x44\x07\x52\x6a\xca\x03\xd3\xc4\xfa\xb0\x09\x98\x90\xa6\x0d\xb1\x49\x7b\x40\xa8\x72\xe3\x4b\x63\xd5\xb1\x8d\xed\xb6\xa0\x6d\xff\x7d\xd7\x04\x5a\xa7\x74\xc0\x24\x26\x61\x1e\xa8\x9b\xef\xbe\xbb\xef\x7a\xf7\x65\x7b\x8b\x8d\xa4\x66\x23\xee\xcb\x4e\xe0\x6e\x8c\x61\x28\xa4\x1b\xb0\x19\x77\x6c\x96\x73\xcb\x7c\x30\x0e\x59\x5e\x64\x0e\x85\xf4\xd9\xc8\x99\x09\x3a\xd6\x5c\x04\x0f\x9c\x75\x72\x01\xeb\xf0\x50\xd9\xe1\x88\xe7\x93\xa9\x65\xb0\xb3\x03\x44\x0c\x37\xb3\xe2\x31\x58\x43\x18\xa4\xc2\x1e\xa1\x3b\x85\x71\x70\x7a\xf6\xf9\x04\xa4\x7e\x3a\xaa\x2e\x63\xef\x08\x84\xe9\x00\x1d\x59\xc0\xe5\x25\x64\x02\xd2\x9a\xe1\xea\xea\x08\x42\x89\xba\x7e\x46\x87\xe4\x61\x4e\x44\xb7\x5f\x78\x85\x83\xf4\x0d\x49\x47\x4d\x1f\x1b\xf8\xee\x3d\x4c\x6a\x1f\xb8\x52\xc7\xd4\x8c\x74\xd5\x98\xb4\x15\xbd\xc4\x52\xc6\x45\xc2\x24\x5d\x45\x25\xb0\x96\x97\x0e\xe6\xa5\x81\xe4\x02\x17\x3a\xa4\x1e\x43\x9b\x2d\x89\x90\xb5\xb0\xa1\x35\x2e\x50\x85\x7c\x3e\x81\x2e\x5b\x5c\xd8\x4f\x4b\x81\x01\xd2\x83\xdf\x5d\x88\x92\x35\x8d\xe8\xe5\x46\x17\xbb\x0f\x59\xe6\x62\x49\xe2\xf0\x7a\x4a\x39\x2d\xf7\xfe\x1f\xb9\xce\xcf\x8e\x89\x25\xe7\xe1\x21\x36\xf3\xe8\x66\xe8\x7a\x56\x8a\x38\x62\x22\x95\x82\xec\x1d\xa4\x14\x1a\x17\x55\x41\xe6\x8a\x47\x59\x22\xf4\xf6\x36\x78\x0c\x10\x0c\x70\x6b\x51\x0b\xa3\xd5\x2d\x68\x13\x21\x3c\x0a\xc8\x24\x74\x3d\x8b\x10\xb7\xd8\xba\x6a\xc3\xc6\x7f\x13\xb9\x91\xca\xd5\x23\x91\xe5\xa6\xaa\xb8\x16\xf0\xf1\xd3\xc5\xc9\x8f\x8b\xb3\xef\x27\x1f\xbe\x9e\x42\x92\x3c\xfe\x38\xbe\x25\xcf\xca\xbb\xa1\x25\x62\xc4\xf6\x22\x44\x6e\x9f\xb9\x06\xed\x91\x5a\xd0\x88\x69\x65\x7b\x4e\x8c\xd6\xe9\x7b\x31\x7d\x59\x19\x01\x87\x87\xfb\xeb\xa0\xfb\xe0\x4e\x04\x9e\x97\xb4\xa5\xb0\x05\x3a\x87\xec\x1a\xfa\xa0\x4c\xce\x55\x69\x3c\x4d\xc6\x6a\x6c\xe1\x3d\x13\x38\x63\x7a\xaa\xd4\x62\x31\xc1\x2b\x44\x0b\xfd\xfd\xc5\x45\x63\xc4\xd6\xec\xd6\xb9\x33\x63\x87\xde\x7f\x0b\x3c\x4c\x3d\x0d\xda\x4a\xad\x25\x8d\x7c\x8c\xbe\x51\x59\xbb\x55\xa3\x37\x57\x12\xb2\x12\xfa\x07\x6f\x7b\xfb\xf4\xd7\x87\xcc\xb6\x2a\xc8\xf8\xf2\x3a\x17\xb4\xcf\x85\x81\x5f\x40\x59\x2c\x74\xb9\x29\x86\x0e\xe7\x4e\x06\x1c\x4a\x3d\xb4\x77\xd9\xbb\xf1\xf8\xbe\x78\x05\xf1\x5c\x44\x79\xec\xeb\x93\x3e\xd5\x64\xc4\x64\x69\x49\xda\x2e\x2e\x81\xc1\x80\xbe\xdc\xf4\x93\x91\xdd\x45\x04\x22\x5e\xd0\x57\x29\x11\xee\x27\xb2\x55\x35\x4d\xe6\xff\xb0\xbc\x8d\x16\x45\x9e\xb4\x66\x58\x2f\x68\x51\x6d\x0f\x7a\xc2\xcd\x9e\xcc\x8b\xca\xe3\x83\xd7\x58\x14\x42\x9d\x43\xaf\xbb\x01\xf0\x46\xfa\xb0\x7c\x8f\x15\xb2\x73\xf7\xaf\xee\xec\x9d\xcb\x3d\xcb\xc7\x3a\x7f\x02\x00\x00\xff\xff\x79\x96\x74\x8b\x98\x08\x00\x00")

func scriptsRestoresharedShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsRestoresharedSh,
		"scripts/restoreShared.sh",
	)
}

func scriptsRestoresharedSh() (*asset, error) {
	bytes, err := scriptsRestoresharedShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/restoreShared.sh", size: 2200, mode: os.FileMode(420), modTime: time.Unix(1460001077, 0)}
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
	"scripts/backupDedicated.sh": scriptsBackupdedicatedSh,
	"scripts/backupShared.sh": scriptsBackupsharedSh,
	"scripts/restoreDedicated.sh": scriptsRestorededicatedSh,
	"scripts/restoreShared.sh": scriptsRestoresharedSh,
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
	"scripts": &bintree{nil, map[string]*bintree{
		"backupDedicated.sh": &bintree{scriptsBackupdedicatedSh, map[string]*bintree{}},
		"backupShared.sh": &bintree{scriptsBackupsharedSh, map[string]*bintree{}},
		"restoreDedicated.sh": &bintree{scriptsRestorededicatedSh, map[string]*bintree{}},
		"restoreShared.sh": &bintree{scriptsRestoresharedSh, map[string]*bintree{}},
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

