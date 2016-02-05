// Code generated by go-bindata.
// sources:
// templates/.gitignore
// templates/Dockerfile
// templates/Godeps
// templates/build.sh
// templates/entrypoint/Dockerfile
// templates/main.go
// templates/version.go
// DO NOT EDIT!

package main

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

var _Gitignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8f\x41\x4b\xc4\x30\x10\x85\xef\xf3\x2b\x02\x7b\xd1\x22\x59\x14\x14\xaf\xa2\x78\x54\x61\x8f\x22\x25\x4d\x27\xdd\x2c\x6d\x26\x66\xa6\xd2\xb2\xec\x7f\x77\xba\x55\x2f\x1e\x26\x33\x2f\xf3\xf2\xf1\xb2\x31\x8f\x34\xe4\xd8\x63\x6b\x5e\x9b\x03\x7a\x31\x41\x05\x5f\x99\x9d\x38\x89\xde\xb8\xd4\x9a\xa7\x39\xb9\x41\xe7\x3e\x36\x6c\x2e\x76\x7b\x57\xfe\xdc\x7c\x09\x95\x25\x2d\xa7\xc5\x04\xb0\x31\xcf\xd4\xb7\x58\x18\x6a\x6a\x0e\x50\x0b\xb2\x2c\xb7\x0f\xc5\xef\xa3\xe8\x8b\xb1\xa0\xe1\x8c\x3e\x06\x25\xe2\x24\x98\x38\x52\xe2\x6d\x2e\x18\xe2\x84\xac\x9c\xf7\xdb\xbb\xfb\xaf\xcf\x0f\xf8\xe9\x96\x46\x45\x54\xd6\x77\x74\x6d\x3b\x5a\xa7\x1b\xeb\xa1\xd6\x5e\xb7\x18\xc6\xf4\x2b\x3a\x92\x39\x23\x2f\xae\xb3\xc6\x29\x53\x11\x5b\xc1\x1a\x64\x70\x31\x2d\x3b\x45\xe0\x84\x7a\x9e\xd3\x55\x36\x17\x0a\x70\x3c\xda\xb7\x42\xcb\xa7\x5e\xdc\x80\xa7\x13\x60\x92\x32\x67\x8a\x49\xb6\xff\x76\xdf\x01\x00\x00\xff\xff\x6a\x2d\xcc\x84\x37\x01\x00\x00")

func GitignoreBytes() ([]byte, error) {
	return bindataRead(
		_Gitignore,
		".gitignore",
	)
}

func Gitignore() (*asset, error) {
	bytes, err := GitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".gitignore", size: 311, mode: os.FileMode(420), modTime: time.Unix(1454710734, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x5f\x4b\xc3\x30\x14\xc5\xdf\xf3\x29\x2e\x7b\xd9\x83\x36\xf1\xcf\x9b\xa8\x58\xd7\x3a\x8a\xb6\x1d\x45\x11\x11\x91\x34\x8d\x69\x35\xcd\x2d\x4d\x3a\x85\xb1\xef\x6e\x9b\x0d\x5f\x7c\xda\x43\x20\x39\x39\xf7\x77\x0e\xf7\xae\xc8\x53\x50\xa8\xb9\x51\x24\x0d\x93\xec\x71\x3c\x71\x01\xb9\x70\x58\xea\xe1\x18\x12\x23\x28\x5c\x56\x28\xbe\x64\x7f\x83\x3b\x95\x0a\x6c\xaf\x09\x09\xa3\x08\x6a\xe7\x3a\x7b\xc1\x58\xcf\xbf\xa9\x6a\x5c\x3d\x94\x83\x95\xbd\x40\xe3\xa4\x71\x93\x8f\x75\xe8\x24\x53\x5d\xcb\xd6\xa7\xf4\x9c\x9e\xb1\xb2\x31\xd3\x13\x98\x42\x7f\x27\xc5\x53\x06\xa2\x6e\xb1\x82\xa3\x9f\x3f\x75\x74\x10\xb2\xc8\x57\x2f\xb0\xc4\x4a\x76\xd6\x7f\xd8\x5e\xb0\x5d\x88\x27\xef\xdb\xb0\xcd\x86\xae\x7a\xfc\x94\xc2\x65\xbc\x95\xdb\x2d\x23\xcf\x79\x71\x1f\x25\xc5\x21\x43\xbe\xc6\xd4\xab\x31\xd6\x71\xad\xf7\xe9\xf4\x20\x86\x87\x48\xb3\x86\xc5\x32\x7f\x8f\xb3\xf0\xf6\x21\x8e\xae\x4e\xc6\xf5\x42\x39\x34\xba\x82\x80\x43\xa0\xab\x0f\xcd\x95\x85\x79\x60\xe7\x40\xc7\x98\x34\x82\xd7\x19\xfd\x07\x9b\xbd\x91\xdf\x00\x00\x00\xff\xff\xd9\xca\xa9\x72\x9c\x01\x00\x00")

func dockerfileBytes() ([]byte, error) {
	return bindataRead(
		_dockerfile,
		"Dockerfile",
	)
}

func dockerfile() (*asset, error) {
	bytes, err := dockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Dockerfile", size: 412, mode: os.FileMode(420), modTime: time.Unix(1454710734, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _godeps = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8d\x4d\x0e\x82\x21\x0c\x44\xf7\xdf\x5d\xb4\x67\xe2\xa7\x96\x2a\x30\x09\x2d\x9e\x5f\x4c\x5c\x40\x5c\x4d\xf2\xde\x24\x4f\xd4\xcb\x8c\xf7\x84\x46\x09\x99\x25\x74\x31\x0f\x94\xaa\x5e\x87\x1a\x0c\x23\xc1\xcd\xb8\xbd\x79\xec\xee\x11\x5c\xcb\x7a\x54\x1c\x18\xdd\x94\x44\xfb\x4b\xf0\x9b\x7f\x8b\xb6\x82\x3b\xf6\xe7\xb7\x91\x39\x4e\xb9\x3e\x01\x00\x00\xff\xff\xee\xda\x45\x34\x9a\x00\x00\x00")

func godepsBytes() ([]byte, error) {
	return bindataRead(
		_godeps,
		"Godeps",
	)
}

func godeps() (*asset, error) {
	bytes, err := godepsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Godeps", size: 154, mode: os.FileMode(420), modTime: time.Unix(1454710697, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _buildSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x4b\x6b\xfa\x40\x14\xc5\xf7\xf3\x29\xee\xdf\x7f\x90\x76\x11\x83\x5b\xc1\x85\x50\x29\x2e\x94\x50\xba\x2c\xc8\x38\x19\x75\x6a\xe6\xc1\x64\x2c\x8a\xe6\xbb\x77\x5e\x71\x12\x68\xa1\xbb\xe4\x77\xcf\x39\xf7\x31\xff\xff\x15\x3b\x26\x8a\x1d\x6e\x8e\x08\x2d\xca\x72\xbb\x59\xac\x97\xf3\xdb\x6d\x52\x6a\xf9\x49\x89\xd9\x60\x4e\xdb\x16\xbd\xaf\xcb\xed\xcb\xea\x6d\x5e\x18\xae\x8a\xac\xd3\xa1\xd5\x7a\xf1\xba\x0c\x96\x5a\x12\x5c\xf7\x4a\x68\x77\x66\x75\xf5\xf4\x0c\x37\x04\x50\x49\x72\xa2\x1a\x3c\x82\x3c\x37\xf8\x00\x59\xf2\xce\x1c\x37\x30\x41\x2d\x42\x44\xaa\x6b\x34\x11\x05\x59\xec\x9b\x72\xad\xea\xb7\x0a\x15\x46\x5f\x95\x64\xc2\x14\x2e\x89\x09\x66\x62\x92\xe6\x90\xeb\x7d\xf2\xc0\x87\x85\x30\x1e\x03\x3f\x55\x4c\x43\xde\x8b\x73\x4e\x85\xc9\x09\x1f\xe8\xdf\x66\xaf\xb1\xa1\x8d\xe9\x35\x77\x09\xfa\x2c\x86\x6e\x0b\xac\xd7\xce\xe1\x3b\xdb\xcf\x2f\x59\x9f\x39\x7d\xf4\x9d\x15\xf4\xa2\xa4\x36\x71\xb4\x1f\xae\x13\x78\xd8\xfd\xb1\x72\x74\x85\xa1\x05\x23\xb1\xa9\x7f\x0b\xe0\xb4\x69\xec\x1a\xf3\x6c\x6a\x11\x25\x47\x09\x59\x44\xee\xff\xc2\x0c\x4c\x9d\x91\x63\xd6\x0d\xeb\x4e\xe6\x7a\xdc\xef\xe0\xe3\x60\xe4\xc9\x1e\xb3\x9a\x56\x23\x2b\x08\x47\xe8\x0b\x02\x49\x0a\xb7\xe8\x30\xc2\x91\x54\x77\xef\x3b\xac\x7b\x92\x04\xf1\xf8\x3d\x41\x47\x3a\x4d\xeb\x47\x46\xdf\x01\x00\x00\xff\xff\x60\x90\x6e\x65\xba\x02\x00\x00")

func buildShBytes() ([]byte, error) {
	return bindataRead(
		_buildSh,
		"build.sh",
	)
}

func buildSh() (*asset, error) {
	bytes, err := buildShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "build.sh", size: 698, mode: os.FileMode(493), modTime: time.Unix(1454710734, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _entrypointDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x28\x4e\x2e\x4a\x2c\x49\xce\xe0\xf2\x75\xf4\xf4\x0b\x01\x62\xd7\x20\x05\xff\xe4\x92\xfc\xa4\x9c\x52\x1d\x05\xcf\xbc\x64\x3d\x05\x9b\x94\xfc\xe4\xec\xd4\x22\x87\x7c\x88\xa8\x5e\x72\x7e\xae\x1d\x17\x97\xa3\x8b\x8b\x42\x75\xb5\x5e\x40\x51\x7e\x56\x6a\x72\x89\x5f\x62\x6e\x6a\x6d\x2d\x86\x00\x97\xab\x5f\x48\x50\x64\x80\x3f\xd0\x64\x85\x68\x25\x3d\x7d\x74\x79\xa5\x58\x2e\x40\x00\x00\x00\xff\xff\x2b\xf7\x77\x5e\x84\x00\x00\x00")

func entrypointDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_entrypointDockerfile,
		"entrypoint/Dockerfile",
	)
}

func entrypointDockerfile() (*asset, error) {
	bytes, err := entrypointDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "entrypoint/Dockerfile", size: 132, mode: os.FileMode(420), modTime: time.Unix(1454710734, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\x4d\x6f\xe3\x36\x10\x3d\x4b\xbf\x62\x2a\xa0\x28\xb5\x48\xe4\xee\xd5\xa8\x0f\xc6\x46\x6d\x53\xd4\x4e\x60\xa7\x41\x81\xa2\x28\x68\x79\x24\x73\x57\x22\x55\x92\x72\xb2\x08\xfc\xdf\x3b\x43\x4a\x59\xaf\x37\xbd\x48\x24\x67\x38\x1f\xef\xbd\x61\x2f\xab\x4f\xb2\x41\xe8\xa4\xd2\x69\xaa\xba\xde\x58\x0f\x22\x4d\xb2\xba\xf3\x19\xfd\x5a\xd3\xf0\xcf\xb8\xf8\x9d\x39\xd5\x68\xd9\xf2\xc6\x7d\x76\x95\x6c\xc3\xd2\xab\x0e\xb3\x94\x16\x8d\xf2\x87\x61\x57\x54\xa6\x9b\x55\x66\x8f\x8d\xd4\x8d\xf3\x72\x56\xb5\x2a\xbb\xb4\x5a\xa4\x68\x8d\xb9\x76\xd8\x1d\xd1\xce\xe2\xef\xc2\xab\x96\x5e\x1d\xc8\xb7\x35\x6c\xb9\x41\x38\x37\xfa\x8f\x7c\x7d\x8f\xbb\x81\x2a\xcc\xd3\xf4\x28\x2d\x84\x1d\x2c\xe0\x06\x8b\x1b\x5e\x8a\xec\xe5\xa5\xb8\xb7\xe6\x23\x56\x7e\x2d\x3b\x3c\x9d\xe6\xdc\x68\x46\xee\xf5\xa0\xab\xd0\xb5\xc8\xe1\x25\x4d\x64\xdf\xc3\x7c\x01\x54\x69\xb1\xc6\xa7\x65\xdf\x8b\x3c\x1c\x16\x7c\x8d\x22\x7e\x13\x28\x8b\xe6\x47\xb4\x4e\x19\x4d\x1e\xc7\xb8\x9a\xee\x2d\x2b\x1f\xcf\xed\xa0\xe3\xc9\xcf\xad\x6c\x1c\x1d\xfc\xf5\x37\x67\xe1\x1d\xe5\x4d\x78\xbd\xf5\x56\xe9\x66\x3a\x49\x38\xc1\x1c\x00\x32\x7c\x96\x5d\xdf\xe2\x15\x60\x76\xc5\x86\x52\x1f\x1f\xa5\x9d\xc7\x62\x36\x77\xbf\x95\x1f\x1e\xd6\xcb\x55\x79\x3a\xfd\x53\xfe\xb9\x5c\xdd\xff\x5e\x46\xb7\x3f\x1c\x31\x4a\x01\xb2\x32\xde\x07\x17\xe2\x43\x4d\x09\x82\xc7\x89\x3e\xa7\x58\xd4\x66\xd0\xc2\xb8\x62\x69\x1b\x97\xa7\xa7\x11\x16\x2a\x59\x54\x46\x7b\x7c\xf6\xf0\x8e\x0b\xfc\x10\x37\x01\xa9\xb1\x28\x46\xab\x41\x7f\xd7\x7b\x37\xf9\x12\xaa\x09\xe9\xe3\x01\x6d\xc7\xd6\x4e\x7e\x42\x51\x1d\xa4\x06\x4a\xb0\x0d\xba\xc9\x83\x03\x2d\x8a\xb5\xf1\xaa\xfe\x2c\x46\xf7\x2b\x18\xd5\x54\x6c\x6f\x7f\x79\x28\x37\xab\xb3\x50\x1b\xac\x50\x1d\x71\xcf\x21\x6b\xd9\x3a\x24\x53\x63\x80\x0b\x8d\xd4\x25\x3f\x5d\x8f\xae\xb4\x26\xd9\x12\x4f\x4a\xfb\x56\x8b\x6c\x0c\x06\x76\x0c\x71\x05\x4f\x52\x79\x86\xc2\x1b\xc0\x67\xe5\x49\x08\xc9\x37\x79\x16\xe0\xed\x80\x84\x90\xe0\x2a\x6a\x63\x43\x12\x55\xc3\xa5\x63\x20\xeb\xab\x84\xb7\x3f\xb4\x2d\xec\x10\x76\x34\x55\x45\x08\x9e\x50\xef\x25\x65\x12\x3f\xf2\x8e\x00\x4e\x92\xfd\xdb\xca\x2c\x5a\x63\xfa\x39\x7c\x7f\xcc\x88\xef\x88\x31\x5f\xe1\xd9\x2a\xb6\x2d\x62\x2f\xde\xc3\x3b\x88\x5b\x24\xc4\xf7\x39\x93\x38\x51\x76\x41\xc5\x05\x6d\xa3\x00\xbe\x66\x6f\x74\x1d\xd5\x27\x26\xb5\xf1\x70\x70\xb7\x93\xe7\x82\xc4\x9f\xc1\xab\x54\x0f\x86\xa7\xe3\x57\x6c\xfb\x73\xda\xdf\xbe\x90\x84\xd1\x2d\x36\xb8\x17\x19\xc0\x4a\x39\xc7\x55\x58\xfc\x77\x50\x96\xf0\x63\x3d\xc2\xf5\xf5\x74\x8f\x81\xfe\x3f\x5d\x47\xf0\xbe\xa0\xf9\x3e\x34\x9f\x26\x16\xfd\x60\xf5\x94\xfa\x15\x8d\xd7\x59\x3c\xeb\x7c\x3c\x23\x70\xad\xe5\xf6\xe3\x8b\xc3\xd3\x3e\xce\xb0\x78\x2c\x37\xdb\xdb\xbb\x75\x1e\xdb\x27\xaf\xef\x16\xa0\x55\x1b\x5a\xa1\xad\xb1\x2b\x74\x3c\x5a\x41\x89\x44\xfb\xb6\x67\xde\x6b\x91\x95\x6c\x84\x27\x7a\x9b\xa6\xcc\xa0\x87\x6e\x87\x76\xa4\xf3\x4b\xe0\x84\xde\xd3\xe2\x5e\x6a\x55\x91\x5c\xce\x63\x86\xb2\x8a\x10\x48\xe4\xa1\xb9\xa9\xb7\x31\xe2\x44\x13\x4f\xe9\x7f\x01\x00\x00\xff\xff\xc1\xd5\xdb\x13\xb5\x05\x00\x00")

func mainGoBytes() ([]byte, error) {
	return bindataRead(
		_mainGo,
		"main.go",
	)
}

func mainGo() (*asset, error) {
	bytes, err := mainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go", size: 1461, mode: os.FileMode(420), modTime: time.Unix(1454710953, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _versionGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\xd2\xd7\x57\x08\x73\x0d\x0a\xf6\xf4\xf7\x53\xc8\x2c\x56\x28\xc9\x48\x55\x48\x2e\x2d\x2a\x4a\xcd\x2b\x51\x48\x2c\x28\xc8\xc9\x4c\x4e\x2c\xc9\xcc\xcf\x53\x08\x4b\x2d\x2a\x06\xd2\x5c\x65\x89\x45\x70\xe5\xb6\x0a\x4a\x06\x7a\x40\xa8\xc4\x05\x08\x00\x00\xff\xff\xbb\x3d\x68\x63\x52\x00\x00\x00")

func versionGoBytes() ([]byte, error) {
	return bindataRead(
		_versionGo,
		"version.go",
	)
}

func versionGo() (*asset, error) {
	bytes, err := versionGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "version.go", size: 82, mode: os.FileMode(420), modTime: time.Unix(1454710984, 0)}
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
	".gitignore": Gitignore,
	"Dockerfile": dockerfile,
	"Godeps": godeps,
	"build.sh": buildSh,
	"entrypoint/Dockerfile": entrypointDockerfile,
	"main.go": mainGo,
	"version.go": versionGo,
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
	".gitignore": &bintree{Gitignore, map[string]*bintree{}},
	"Dockerfile": &bintree{dockerfile, map[string]*bintree{}},
	"Godeps": &bintree{godeps, map[string]*bintree{}},
	"build.sh": &bintree{buildSh, map[string]*bintree{}},
	"entrypoint": &bintree{nil, map[string]*bintree{
		"Dockerfile": &bintree{entrypointDockerfile, map[string]*bintree{}},
	}},
	"main.go": &bintree{mainGo, map[string]*bintree{}},
	"version.go": &bintree{versionGo, map[string]*bintree{}},
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

