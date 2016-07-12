// Code generated by go-bindata.
// sources:
// templates/.dockerignore
// templates/.gitignore
// templates/Dockerfile
// templates/build.sh
// templates/cross-compile.sh
// templates/entrypoint/Dockerfile
// templates/main.go
// templates/vendor/manifest
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

var _Dockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xcf\x2c\xe1\xaa\xae\xd6\x0b\x28\xca\xcf\x4a\x4d\x2e\xf1\x4b\xcc\x4d\xad\xad\xe5\x4a\xcd\x2b\x29\xaa\x2c\xc8\xcf\xcc\x2b\xd1\xc7\x90\x4b\xcf\xd7\xc5\xa7\x1e\x9b\x74\x52\x69\x66\x4e\x8a\x5e\x71\x06\x57\x72\x51\x7e\x71\xb1\x6e\x72\x7e\x6e\x41\x66\x4e\x2a\x48\x20\xc8\xd5\xd1\xc5\xd7\x55\x2f\x37\x85\x0b\x10\x00\x00\xff\xff\x6b\xa3\x55\x57\x89\x00\x00\x00")

func DockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_Dockerignore,
		".dockerignore",
	)
}

func Dockerignore() (*asset, error) {
	bytes, err := DockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".dockerignore", size: 137, mode: os.FileMode(420), modTime: time.Unix(1468360567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _Gitignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x8f\x41\x4b\xfc\x30\x10\xc5\xef\xf3\x29\x02\x7b\xf9\xff\x8b\x66\x51\x50\xbc\x8a\xe2\x51\x85\x3d\x8a\x94\x34\x9d\x74\xb3\xb4\x99\x98\x99\x4a\xcb\xb2\xdf\xdd\xe9\xae\x7a\x51\x3c\x4c\x66\x5e\xde\xcb\x8f\xc9\xca\xdc\xd1\x90\x63\x8f\xad\x79\x6a\x76\xe8\xc5\x04\x15\x7c\x66\x36\xe2\x24\x7a\xe3\x52\x6b\xee\xe7\xe4\x06\x9d\xfb\xd8\xb0\xf9\xb7\xd9\xba\xf2\x9d\xe6\xff\x50\x59\xd2\x72\x5a\x4c\x00\x2b\xf3\x40\x7d\x8b\x85\xa1\xa6\x66\x07\xb5\x20\xcb\x72\x7b\x5b\xfc\x36\x8a\xbe\x18\x0b\x1a\xce\xe8\x63\x50\x22\x4e\x82\x89\x23\x25\x5e\xe7\x82\x21\x4e\xc8\xca\x79\xb9\xba\xbe\x79\x7f\x7b\x85\xcf\x6e\x69\x54\x44\x65\x7d\x47\x17\xb6\xa3\xd3\x74\x69\x3d\xd4\xda\xeb\x16\xc3\x98\xbe\x44\x47\x32\x67\xe4\x25\x75\xd4\x38\x65\x2a\x62\x2b\x38\x2d\x32\xb8\x98\x16\x4f\x11\x38\xa1\x9e\xc7\xed\x2a\x9b\x0b\x05\xd8\xef\xed\x73\xa1\xe5\x53\x8f\x6e\xc0\xc3\x01\x30\x49\x99\x33\xc5\x24\xeb\x1f\x5e\x47\xe7\x7f\xe5\x7f\xb3\x3f\x02\x00\x00\xff\xff\xbb\x83\x6f\xb2\x6a\x01\x00\x00")

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

	info := bindataFileInfo{name: ".gitignore", size: 362, mode: os.FileMode(420), modTime: time.Unix(1468360567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x48\xcf\xcf\x49\xcc\x4b\xb7\x32\xd4\x33\xe3\xf2\x75\xf4\xf4\x0b\x01\x62\xd7\x20\x05\xff\xe4\x92\xfc\xa4\x9c\x52\x1d\x05\xcf\xbc\x64\x3d\x05\x9b\x94\xfc\xe4\xec\xd4\x22\x87\x7c\x88\xa8\x5e\x72\x7e\xae\x1d\x17\x57\xb8\x7f\x90\xb7\x8b\x67\x90\x82\x7e\x7a\xbe\x7e\x71\x51\xb2\x7e\x7a\x66\x49\x46\x69\x12\x48\x52\x1f\xaa\x50\xbf\xba\x5a\x2f\xa0\x28\x3f\x2b\x35\xb9\xc4\x2f\x31\x37\xb5\xb6\x96\xcb\xd9\x3f\x20\x52\x41\x8f\x24\x2d\x5c\x41\xa1\x7e\x0a\xa9\x79\x65\x0a\xce\xee\xfe\xf1\xae\x7e\x8e\x4e\x3e\xae\x2e\xb6\x06\x40\x77\x2b\x24\x95\x66\xe6\xa4\x28\xe8\xe6\x2b\xa0\xeb\x51\xd0\x4d\x54\xd0\xcd\x49\x49\xcb\x49\x4c\x2f\x56\x50\xd7\x2d\x56\x57\xd0\xe3\xe2\x72\xf6\x75\x51\x88\x56\xd2\xc3\xb0\x40\x29\x96\x0b\x10\x00\x00\xff\xff\x03\x3a\x2f\xc1\x09\x01\x00\x00")

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

	info := bindataFileInfo{name: "Dockerfile", size: 265, mode: os.FileMode(420), modTime: time.Unix(1461000606, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _buildSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\x6d\x4f\xdb\x30\x10\xfe\x9e\x5f\x71\xf3\x22\x5e\x26\xa5\x11\xdb\x84\xa6\xa2\x6e\x03\x15\x21\xa4\x01\x15\x08\x31\x69\x9b\x2a\x93\x38\xad\x47\x62\x57\xb1\x5b\x0a\xa5\xff\x7d\x3e\xdb\x69\xd2\x12\xa6\xf6\x4b\x92\xbb\xe7\xee\xf1\xe3\x7b\xe9\xfb\x77\xf1\x3d\x17\xf1\x3d\x55\xe3\x20\x38\x1e\x0c\x86\x97\xc7\x17\xa7\xbd\xc5\xa2\x33\x28\xe5\x5f\x96\xe8\x4b\x5a\xb0\xe5\x32\x38\xb9\x3d\xff\xd1\x1f\xf6\xcf\xaf\x7b\xe1\xe0\xae\x1f\xdf\x4f\x79\x9e\x06\xe7\x17\xc7\x67\xa7\x2e\x20\x97\x09\xcd\xe3\xb0\x4a\x10\x04\x16\x31\x94\x62\x98\xca\xe4\x81\x95\x7b\xfb\xb0\x08\x00\xdc\x07\x58\x27\x44\x91\xa6\x23\x08\xeb\x2c\x5d\xb4\x6b\xe8\x04\xcb\x46\xbc\xcd\xec\xc3\xed\x3b\x8c\xa4\x54\x3d\x12\x1e\x90\x86\x85\x96\xc9\xd8\xd8\x3e\x92\xc0\x18\x99\x98\xc1\xd9\xd5\xd5\x4d\x2f\xe7\x62\x3a\x37\xde\x8a\x90\x02\x52\x2a\x10\x4c\x1b\x63\xc4\x85\xd2\x34\xcf\xd5\x34\xcb\xf8\x1c\x12\x34\xe5\x69\x96\x23\x62\x37\x7a\xdc\x85\x48\x02\x09\x17\x2b\xe9\xcb\x38\x5c\x54\x02\x97\x51\xe8\x48\xc9\xda\x71\xd5\x7c\xf3\xc8\xeb\x3a\x80\xa4\xb4\x7c\xe4\x82\x00\xa1\x45\x7a\xf8\x99\x60\x6c\x22\x27\x4f\x1e\x9d\x4c\x20\x5c\xf1\xd5\xd7\xf9\xc1\x48\xd2\xe5\xd3\x44\x72\xa1\x63\x0c\xe1\x82\x6b\x1f\x52\x16\x10\x95\x59\x33\x0c\x7e\x1b\xf3\xce\x0e\x14\x0f\x29\x2f\x21\x5a\x4b\x89\xc1\x13\x9a\x3c\xd0\x11\xdb\xae\x24\x39\xd5\x4c\xe9\x06\x3f\x66\x28\xa7\x62\x3d\xda\x18\x4c\xac\x39\x0a\x52\x83\x79\x9d\xc9\x7c\x5a\xb0\x06\x73\x37\x66\xf3\x89\x2c\x75\xec\x21\xaf\xcb\xee\xec\xee\x0e\x2a\xe1\xe0\xa3\x90\x34\xa3\x7a\xa3\x11\x0a\xa6\x94\x11\xd2\x0b\x0f\xb0\xe8\xc9\x58\x42\xe8\x4d\xf8\x3d\xe7\x1a\x0e\xec\xfd\x96\x52\xa9\x61\x22\x8b\x09\xcf\xd9\xd0\x0a\xdd\xdb\xc7\x2c\x99\x2c\x6d\x33\x01\x17\xe0\xea\x02\xae\x63\xcc\x5b\x2a\x1f\xd5\x91\x91\x67\x0f\xe5\x80\x58\x6e\x84\x7e\xfa\x72\x08\xb6\x7a\x2b\xbf\x67\x27\x36\x37\x17\xa3\x2e\x84\x0b\x4c\x6c\xba\x64\xe1\xe2\x96\xc4\x03\x37\xfb\x21\x44\x1c\xb1\x4f\xd7\x4e\x5f\x21\x4e\xd9\x2c\x16\xd3\x3c\xb7\x21\xa9\x14\x2c\xf0\x0f\x23\xc6\x5d\x78\xa5\xc2\x5e\x06\x36\x03\x34\x7e\x2f\x2f\x60\xef\x0a\x88\xf5\x64\xd4\xc8\x4e\x49\xb3\x17\x7d\xd5\x6a\xe0\xa6\xa7\x8e\xc1\xca\xb6\x27\x47\x4f\x8d\xc3\x26\x6e\xc7\x59\x4f\x0d\xf4\xdd\xd7\x02\xac\x3c\x15\xd6\xa8\xc5\x81\xfa\x9f\xd4\x2d\xd4\x36\x46\xf2\xb5\xe0\xa6\xb3\x41\x5b\xb2\x9c\x51\xc5\x86\x9b\xf4\xab\x89\x4a\xb9\xd2\xd5\x98\xad\x40\x95\x41\x53\x83\x49\x9e\x33\x5c\x1e\xf5\xbe\x30\xb0\x8e\xf1\x74\x46\xcf\x64\xcd\x41\x56\xe3\x3a\x7b\x3b\x00\xf9\xe2\xaa\xc9\xc9\x5d\x29\x35\x73\xb6\x37\x02\x50\x44\x41\xb9\x58\x1f\x17\x99\xb2\x6a\x6f\xf2\x0c\x7e\x19\x3a\x34\x11\xe8\xf5\xcc\x5e\xb2\x85\x27\xf0\xe7\x08\xf4\x98\x89\xa0\x6e\xea\xbe\x6b\x89\x13\x94\x48\x7c\x4b\xd6\x3d\xe8\x80\x38\x6b\xe1\x37\x1c\x29\x1e\xb4\x65\x37\x47\x6b\x4b\x7d\x75\xf3\xb3\x99\x77\x75\x93\xdb\x25\xf5\x55\x8a\xde\x48\x7e\xed\xdc\x4d\x82\x57\x75\xdd\x8e\xc8\x2e\x90\xc8\x2f\x90\x2d\xa9\x5a\x96\x4e\x2b\x99\x4b\x70\x8b\x4b\xab\x0b\x1d\xf7\xc7\xda\x51\x63\x7f\xc7\xb1\x39\x69\xdc\xd0\x19\xaf\x1f\xa5\xb9\xe6\xb0\xdc\x10\x7e\x0f\xfe\x05\x00\x00\xff\xff\xdd\xae\xf4\xc6\xd1\x07\x00\x00")

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

	info := bindataFileInfo{name: "build.sh", size: 2001, mode: os.FileMode(493), modTime: time.Unix(1468363009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _crossCompileSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x90\x5f\x4b\xc3\x30\x14\xc5\xdf\xf3\x29\x8e\xb1\x0c\x7d\xe8\xca\x50\x86\x38\x26\x88\x03\xf5\x41\x36\xa6\xe0\x83\x8a\x64\x4d\xba\x5e\x4c\x93\xd1\x76\x7f\x70\xeb\x77\x37\x89\xad\xc3\xa7\x34\xe7\xfe\xce\x49\xef\x39\x3d\x49\x16\x64\x92\x85\xa8\x72\xc6\x5e\x9e\x66\x9f\x93\xc7\xf9\x98\x47\xb3\xd7\x49\x52\x17\x2b\xce\x18\x65\x78\x43\xfc\x0d\x1e\xb5\x53\x8e\x8f\x11\xea\x5c\x19\x06\xa8\x34\xb7\xe0\xc6\xc2\xb1\x90\x54\x72\xaf\xed\xa8\xc6\x80\x65\xc4\xd8\x62\x4d\x5a\x9e\x9d\x63\xef\xe4\xcc\x96\x58\x5a\x5b\x81\x0c\xa4\x28\xb7\xee\xd0\x64\xd6\x3b\xb8\x2f\x69\xb7\xd5\x08\xd2\x3a\xac\x03\x45\x99\xe6\x1e\xbd\xb8\x1a\x42\x14\x72\x78\xf9\x37\xef\x5e\x0d\xe1\x64\x96\xd7\x88\xf6\x3e\xb8\x89\xfd\xe9\x7d\x0d\xef\x40\xb3\xc1\xfd\x74\xfa\xec\x16\xf2\x04\x77\x97\xdb\xf9\xdd\x43\xb8\x7a\x90\xa3\x9f\x84\x98\x7e\x95\xe3\x06\x89\x54\x9b\xc4\xac\xb5\x0e\x76\x69\x8d\x62\xed\xd1\xb8\x1e\x0c\xd5\xed\x2a\x65\x81\xb8\xcc\xd0\x15\x92\xe0\xdd\x1b\x7a\x3d\x14\x5f\xae\x03\xc4\xab\xe3\xc8\x3b\x57\xc2\x50\xda\x5a\xb5\x4d\x85\x46\xa1\xaa\x4a\x2c\xd5\x38\x1a\x74\x1d\x46\xad\x74\xec\xcf\x19\x0b\x41\xa6\xf5\xf9\xd7\xfd\x4f\x1d\x0e\x08\x71\xe0\x41\xc9\x04\x69\x25\xfd\xba\x61\x8d\x7f\xc0\xaf\xd2\x11\x4d\x48\x63\x3f\x01\x00\x00\xff\xff\x7b\xc1\x31\x83\xef\x01\x00\x00")

func crossCompileShBytes() ([]byte, error) {
	return bindataRead(
		_crossCompileSh,
		"cross-compile.sh",
	)
}

func crossCompileSh() (*asset, error) {
	bytes, err := crossCompileShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cross-compile.sh", size: 495, mode: os.FileMode(493), modTime: time.Unix(1468362687, 0)}
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

var _mainGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\x4d\x6f\xe3\x36\x10\x3d\x4b\xbf\x62\x2a\xa0\x28\xb5\x48\xe4\xee\xd5\xa8\x0f\xc6\x46\x6d\x53\xd4\x4e\x60\xa7\x41\x81\xa2\x28\x68\x79\x24\xb3\x2b\x91\x2a\x49\x39\x59\x04\xfe\xef\x9d\x21\xa5\xac\xd7\x9b\xbd\x48\x24\x67\x38\x1f\xef\xbd\x61\x2f\xab\x8f\xb2\x41\xe8\xa4\xd2\x69\xaa\xba\xde\x58\x0f\x22\x4d\xb2\xba\xf3\x19\xfd\x5a\xd3\xf0\xcf\xb8\xf8\x9d\x39\xd5\x68\xd9\xf2\xc6\x7d\x72\x95\x6c\xc3\xd2\xab\x0e\xb3\x94\x16\x8d\xf2\x87\x61\x57\x54\xa6\x9b\x55\x66\x8f\x8d\xd4\x8d\xf3\x72\x56\xb5\x2a\xbb\xb4\x5a\xa4\x68\x8d\xb9\x76\xd8\x1d\xd1\xce\xe2\xef\xc2\xab\x96\x5e\x1d\xc8\xb7\x35\x6c\xb9\x41\x38\x37\x1e\x95\x53\x46\x77\xb8\x57\x92\xe3\xec\x71\x37\x50\xa9\x79\x9a\x1e\xa5\x85\xb0\x83\x05\xdc\x60\x71\xc3\x4b\x91\xbd\xbc\x14\xf7\xd6\xfc\x8b\x95\x5f\xcb\x0e\x4f\xa7\x39\x77\x9c\x91\x7b\x3d\xe8\x2a\xb4\x2f\x72\x78\x49\x13\xd9\xf7\x30\x5f\x00\x95\x5c\xac\xf1\x69\xd9\xf7\x22\x0f\x87\x05\x5f\xa3\x88\x5f\x05\xca\xa2\xf9\x11\x2d\xd7\x43\x1e\xc7\xb8\x9a\xee\x2d\x2b\x1f\xcf\xed\xa0\xe3\xc9\xcf\xad\x6c\x1c\x1d\xfc\xf5\x37\x67\xe1\x1d\xe5\x4d\x78\xbd\xf5\x56\xe9\x66\x3a\x49\x38\xc1\x1c\x00\x32\x7c\x96\x5d\xdf\xe2\x15\x60\x76\xc5\x86\x52\x1f\x1f\xa5\x9d\xc7\x62\x36\x77\xbf\x95\x1f\x1e\xd6\xcb\x55\x79\x3a\xfd\x53\xfe\xb9\x5c\xdd\xff\x5e\x46\xb7\x3f\x1c\x51\x4b\x01\xb2\x32\xde\x07\x17\xe2\x43\x4d\x09\x82\xc7\x89\x3e\xa7\x58\xd4\x66\xd0\xc2\xb8\x62\x69\x1b\x97\xa7\xa7\x11\x16\x2a\x59\x54\x46\x7b\x7c\xf6\xf0\x8e\x0b\xfc\x10\x37\x01\xa9\xb1\x28\x46\xab\x41\x7f\xd7\x7b\x37\xf9\x12\xaa\x09\x09\xe5\x01\x6d\xc7\xd6\x4e\x7e\x44\x51\x1d\xa4\x06\x4a\xb0\x0d\x02\xca\x83\x03\x2d\x8a\xb5\xf1\xaa\xfe\x24\x46\xf7\x2b\x18\x65\x55\x6c\x6f\x7f\x79\x28\x37\xab\xb3\x50\x1b\xac\x50\x1d\x71\xcf\x21\x6b\xd9\x3a\x24\x53\x63\x80\x0b\x8d\xd4\x25\x3f\x5d\x8f\xae\xb4\x26\xfd\x12\x4f\x4a\xfb\x56\x8b\x6c\x0c\x06\x76\x0c\x71\x05\x4f\x52\x79\x86\xc2\x1b\xc0\x67\xe5\x49\x08\xc9\x57\x79\x16\xe0\xed\x80\x84\x90\xe0\x2a\x6a\x63\x43\x12\x55\xc3\xa5\x63\x20\xeb\x8b\x84\xb7\x3f\xb4\x2d\xec\x10\x76\x34\x5e\x45\x08\x9e\x50\xef\x25\x65\x12\x3f\xf2\x8e\x00\x4e\x92\xfd\xdb\xca\x2c\x5a\x63\xfa\x39\x7c\x7f\xcc\x88\xef\x88\x31\x5f\xe1\x21\x2b\xb6\x2d\x62\x2f\xde\xc3\x3b\x88\x5b\x24\xc4\xf7\x39\x93\x38\x51\x76\x41\xc5\x05\x6d\xa3\x00\xbe\x64\x6f\x74\x1d\xd5\x27\x26\xb5\xf1\x70\x70\xb7\x93\xe7\x82\xc4\x9f\xc1\xab\x54\x0f\x86\xa7\xe3\x57\x6c\xfb\x73\xda\xdf\xbe\x90\x84\x19\x2e\x36\xb8\x17\x19\xc0\x4a\x39\xc7\x55\x58\xfc\x6f\x50\x96\xf0\x63\x3d\xc2\xf5\xf5\x74\x8f\x81\xfe\x96\xae\x23\x78\x9f\xd1\x7c\x1f\x9a\x4f\x13\x8b\x7e\xb0\x7a\x4a\xfd\x8a\xc6\xeb\x2c\x9e\x75\x3e\x9e\x11\xb8\xd6\x72\xfb\xf1\xe9\xe1\x69\x1f\x67\x58\x3c\x96\x9b\xed\xed\xdd\x3a\x8f\xed\x93\xd7\x77\x0b\xd0\xaa\x0d\xad\xd0\xd6\xd8\x15\x3a\x1e\xad\xa0\x44\xa2\x7d\xdb\x33\xef\xb5\xc8\x4a\x36\xc2\x13\x3d\x52\x53\x66\xd0\x43\xb7\x43\x3b\xd2\xf9\x39\x70\x42\x0f\x6b\x71\x2f\xb5\xaa\x48\x2e\xe7\x31\x43\x59\x45\x08\x24\xf2\xd0\xdc\xd4\xdb\x18\x71\xa2\x89\xa7\xf4\xff\x00\x00\x00\xff\xff\xa3\x6d\x3c\x8c\xbe\x05\x00\x00")

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

	info := bindataFileInfo{name: "main.go", size: 1470, mode: os.FileMode(420), modTime: time.Unix(1468363442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorManifest = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\xd3\xcf\x8a\xdb\x30\x10\x06\xf0\xb3\xf3\x14\xc5\xe7\xb4\x96\x22\x4b\x1a\xed\xab\x94\x1e\x66\x46\xa3\xc4\x10\xff\xc1\x56\x16\x96\xb2\xef\x5e\x25\x9b\x92\x6e\x59\x82\x77\x21\x27\xdb\x83\x3e\xf8\x7e\x92\xfc\x7b\x53\xd5\xcf\x32\x2f\xdd\x38\xd4\x4f\xdf\xd4\xb6\x7c\x46\x99\x64\x88\x32\x70\x27\x4b\x99\xfd\xdc\x54\x55\x59\x55\x55\x75\xd7\x4f\xe3\x9c\x27\xcc\x87\x32\xae\xf7\x5d\x3e\x9c\xe8\x07\x8f\x7d\xc3\x63\x94\x3d\x0e\xfb\x25\x63\xc3\xc7\xae\xde\x5e\x96\xcf\x32\x8d\x4b\x97\xc7\xf9\xe5\xbc\xfc\x90\xf3\xb4\x3c\x35\xcd\x8a\xd8\x73\x77\xad\x53\x7b\x9d\xac\x8f\x46\xa9\x18\x1d\x7a\x50\xc8\x1a\xac\x63\xa5\x2c\xb7\x64\x35\x70\x8a\x6d\x00\xe1\x6b\x94\x66\x1c\xf8\x52\xae\xc7\x25\xcb\x7c\x1d\x0f\x63\x96\x25\x9f\x2d\x79\x3e\x49\x19\xbd\x6e\xd7\x98\x66\x19\x97\x66\x3f\x7e\x5f\xa4\x2f\x3b\xd4\xbc\x3d\xd6\xdb\xde\xc7\x3f\xc0\x19\x4f\xe4\x20\x01\x70\xdb\x6a\x00\x08\x49\x61\xa0\x24\xd8\xa2\x35\xc6\x44\xe3\x28\x6a\x1b\xed\x7d\xdc\xdf\xea\xef\xfb\x7d\x89\x9c\x30\x77\x87\xd2\xfc\x38\xae\x66\x7e\x14\xb9\x09\x0b\x83\xa3\x4f\x11\x10\x6c\x50\x36\x39\x8f\xda\x5b\x83\x89\xda\x28\x31\x32\x58\x29\x83\x36\x3d\xe0\xf8\x7a\xcc\x79\x38\x6f\xff\xa5\x1b\xd2\x51\xd6\x92\xee\x24\x6f\xb2\xc0\x24\xc9\xb3\xb1\x26\x68\x66\x2c\x36\x6d\x13\x68\xd0\x91\x15\x69\x20\x0e\xde\xb8\x48\xf4\x48\x59\xb7\x94\xb7\x97\x4f\xb3\xfe\x8b\xfd\x73\x5a\x8e\x7c\xb9\x72\xc9\xea\xe4\xbd\x02\x6f\x95\x20\x83\x4a\x68\x80\x82\xdd\x11\x85\x64\x76\xce\x84\x07\x98\xde\x4a\xf4\x12\x3b\x3c\x57\x8c\x42\xa7\xfd\x5a\xd8\xdd\xec\x4d\x97\x52\xf9\xad\x2c\xee\x14\x82\x0b\xa1\xd5\x1a\x5c\xdb\x12\x95\x5b\xe8\x70\xa7\x5d\x44\x1d\x0c\xb3\x36\x9f\xd5\x6d\xaa\x5f\x9b\xd7\xcd\x9f\x00\x00\x00\xff\xff\xed\xbe\xc0\xb1\x47\x05\x00\x00")

func vendorManifestBytes() ([]byte, error) {
	return bindataRead(
		_vendorManifest,
		"vendor/manifest",
	)
}

func vendorManifest() (*asset, error) {
	bytes, err := vendorManifestBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/manifest", size: 1351, mode: os.FileMode(420), modTime: time.Unix(1468362540, 0)}
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
	".dockerignore": Dockerignore,
	".gitignore": Gitignore,
	"Dockerfile": dockerfile,
	"build.sh": buildSh,
	"cross-compile.sh": crossCompileSh,
	"entrypoint/Dockerfile": entrypointDockerfile,
	"main.go": mainGo,
	"vendor/manifest": vendorManifest,
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
	".dockerignore": &bintree{Dockerignore, map[string]*bintree{}},
	".gitignore": &bintree{Gitignore, map[string]*bintree{}},
	"Dockerfile": &bintree{dockerfile, map[string]*bintree{}},
	"build.sh": &bintree{buildSh, map[string]*bintree{}},
	"cross-compile.sh": &bintree{crossCompileSh, map[string]*bintree{}},
	"entrypoint": &bintree{nil, map[string]*bintree{
		"Dockerfile": &bintree{entrypointDockerfile, map[string]*bintree{}},
	}},
	"main.go": &bintree{mainGo, map[string]*bintree{}},
	"vendor": &bintree{nil, map[string]*bintree{
		"manifest": &bintree{vendorManifest, map[string]*bintree{}},
	}},
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

