// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5f\x6b\xdb\x3e\x14\x7d\xb6\x3e\xc5\xf9\x99\xfe\xa8\xdd\xa5\x4a\xdb\xb7\x0d\xf2\x50\xda\x0c\x32\xb6\x76\x90\xc2\x1e\xba\x52\x14\xfb\x3a\x11\x75\x24\xef\x4a\x29\x0b\x42\xdf\x7d\x48\x4e\xc2\xf6\x64\x4b\xe7\xdc\xf3\x47\x37\x84\xe9\x85\xb8\xb3\xc3\x9e\xf5\x7a\xe3\x71\x73\x75\xfd\xf1\x72\x60\x72\x64\x3c\x3e\xab\x86\x56\xd6\xbe\x61\x61\x1a\x89\xdb\xbe\x47\x26\x39\x24\x9c\xdf\xa9\x95\xe2\x69\xa3\x1d\x9c\xdd\x71\x43\x68\x6c\x4b\xd0\x0e\xbd\x6e\xc8\x38\x6a\xb1\x33\x2d\x31\xfc\x86\x70\x3b\xa8\x66\x43\xb8\x91\x57\x47\x14\x9d\xdd\x99\x56\x68\x93\xf1\xaf\x8b\xbb\xf9\xc3\x72\x8e\x4e\xf7\x84\xc3\x1d\x5b\xeb\xd1\x6a\xa6\xc6\x5b\xde\xc3\x76\xf0\x7f\x99\x79\x26\x92\xe2\x62\x1a\xa3\x10\x21\xa0\xa5\x4e\x1b\x42\xb9\x55\xda\x94\x88\x51\x4c\xa7\xb8\x4b\x79\xd6\x64\x88\x95\xa7\x16\xab\x3d\xce\xc9\xf8\xe6\x74\x75\x2e\x71\xff\x88\x87\xc7\x27\xcc\xef\x17\x4f\x52\x0c\xaa\x79\x53\x6b\x42\xd2\x10\x42\x6f\x07\xcb\x1e\x95\x28\x4a\xeb\x4a\x51\x94\xab\xbd\xa7\xf4\x13\x02\x3c\x6d\x87\x5e\x79\x42\x39\xb2\x5c\xb6\xcc\xd0\xc0\xda\xf8\x0e\xe5\xff\xbf\x4a\xc8\xef\x07\xc5\x18\x45\x9d\x63\x9e\xad\x94\x23\x7c\x9a\x21\x7f\x8f\x78\x9a\x7d\x57\x0c\xd7\x6c\x68\xab\x1c\x66\x78\x7e\x21\xe3\xe5\xc2\x78\xe2\x4e\x35\x14\xb2\x34\x2b\xb3\x26\x9c\xbd\x4e\x70\x66\xd4\x36\xcb\xc8\x07\xb5\x25\x97\xf4\x8b\x22\x84\xcb\x83\x7e\x8c\x32\x1d\x4e\x51\x5c\x88\xe5\x61\x26\xc6\x49\xd6\x22\xd3\xe2\x32\x46\x11\x85\xe8\x76\xa6\xc9\x9d\xab\x1a\x41\x14\x29\x48\xaf\x0d\x39\x3c\xbf\x3c\xbf\xa4\xd2\xa2\xe8\x2c\xe3\x75\x72\xc8\x97\x7c\xc7\x28\xc7\xbc\x41\x14\xc5\x6a\x02\x62\x4e\xd8\x37\xc5\x6e\xa3\xfa\x65\x06\xab\x91\x53\x8b\xa2\xd0\x5d\x66\xfc\x37\x83\xd1\x7d\x9e\x29\x3a\xa5\xfb\x8a\x98\x13\x9c\x2a\x8c\xbe\x33\xa8\x61\x20\xd3\x56\xf9\x38\xc1\xaa\x16\x09\xb5\x4e\x2e\x7d\x6b\x77\x5e\xfe\x60\xed\xa9\xca\xfb\x90\x5f\xac\x36\x47\xe2\x18\xb7\x2a\x7f\x9a\xb2\xae\xeb\x53\xb7\xa3\x4b\xb2\xb7\x9c\x4b\x8e\x5a\xc4\x3c\x6a\x2d\x3d\x6b\xb3\x4e\x1c\x39\x4f\x9c\xaa\xfe\x90\x45\x32\x71\xfe\x5b\xfb\xea\x3a\xcb\xfd\xb3\xfa\xb1\xd9\xb8\xf9\xc3\x8b\xc6\x28\xfe\x04\x00\x00\xff\xff\x95\x06\x0f\xa4\x50\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 848, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\xef\x6f\xdc\x36\xef\x7f\x7d\xfe\x2b\xd8\x00\x0d\xec\xe2\xe6\xeb\x86\x61\xf8\x7e\xaf\xcf\x3d\xc0\xb0\xb5\xd8\x3d\x7b\x9a\x16\x6b\xbb\x37\x45\xd1\x39\xb6\x7c\xa7\xc6\x96\x5d\x4b\x4e\x73\xcb\xf2\xbf\x3f\x20\x29\xd9\xb2\xcf\x77\xe9\x8f\xe4\xde\xc4\xa6\x48\x91\xfc\x88\x22\x29\x39\x8b\x05\xfc\x52\xd5\xbb\x46\x6e\xb6\x06\x7e\x78\xfc\xfd\xff\x7f\x57\x37\x42\x0b\x65\xe0\x59\x92\x8a\xf3\xaa\xba\x80\xb5\x4a\x63\xf8\xb9\x28\x80\x98\x34\xe0\x78\x73\x29\xb2\x38\x58\x2c\xe0\xf5\x56\x6a\xd0\x55\xdb\xa4\x02\xd2\x2a\x13\x20\x35\x14\x32\x15\x4a\x8b\x0c\x5a\x95\x89\x06\xcc\x56\xc0\xcf\x75\x92\x6e\x05\xfc\x10\x3f\x76\xa3\x90\x57\xad\xca\x70\x0a\xa9\x88\xe5\xbf\xeb\x5f\x9e\x9e\xbd\x7a\x0a\xb9\x2c\x84\xa3\x35\x55\x65\x20\x93\x8d\x48\x4d\xd5\xec\xa0\xca\xc1\x78\xfa\x4c\x23\x44\x1c\x04\x75\x92\x5e\x24\x1b\x01\x45\x95\x64\x41\x20\xcb\xba\x6a\x0c\x84\xc1\xec\x44\xa8\xb4\xca\xa4\xda\x2c\x3e\xe8\x4a\x9d\x04\xb3\x93\xbc\x34\xf8\xa7\x11\x79\x21\x52\x73\x12\x04\xb3\x93\x8d\x34\xdb\xf6\x3c\x4e\xab\x72\x91\x5b\x87\xa5\x4a\xdb\xf3\xc4\x54\xcd\x42\x28\xe2\xbf\x8d\x67\xa1\xd3\xad\x28\x93\x85\xc8\x36\xe2\x4b\xf8\x73\x29\x8a\xec\x4b\x04\xa4\xca\xc4\xd5\x49\x10\x05\x08\xdb\x2b\xa2\x41\x23\xec\x82\x69\x48\x14\x08\x65\x62\x3b\x60\xb6\x89\x81\x4f\x89\x26\x5c\x44\x06\x79\x53\x95\x90\x40\x5a\x95\x75\x21\x71\x71\xb4\x68\xc0\x62\x17\x07\x66\x57\x0b\x37\xa5\x36\x4d\x9b\x1a\xb8\x0e\x66\x67\x49\x29\x00\x00\x29\x52\x6d\x80\x7e\x7f\x21\x9a\xcb\x13\x95\x94\x62\x5e\x95\xd2\x88\xb2\x36\xbb\x93\xbf\x82\xd9\x2f\x95\xca\xe5\x06\xc8\x06\xf7\x6c\x99\x53\x7a\x1d\xb2\x3f\xcd\x36\x42\x03\xc0\xdb\x77\x8f\xf0\xd1\x9f\x1b\x81\xd4\x43\xee\x67\x88\x95\x26\x6e\x7a\xf4\xb8\x09\xc6\x11\xfb\x1a\x91\x12\x1a\xd9\xe9\xd1\x63\x97\x3c\x34\xe4\xff\xad\xaa\x2e\xac\x31\x2f\x2b\x2d\x8d\xac\x94\xe3\xdf\xe2\xd0\x90\xfb\x65\x55\xc8\x74\x07\x70\x5e\x55\x05\xc0\x00\x96\x9a\x86\x06\xec\x37\xb4\x5c\xdd\xb4\x99\xd0\x69\x23\xcf\x85\x86\x04\xc8\x74\xa8\xdd\x90\x8d\x7a\x5e\x6d\xbb\x26\x9d\x5c\xbf\x2a\x9d\x47\x00\x52\x19\x80\xc5\x02\x18\x13\x72\xcd\xcd\xc2\x73\x17\x52\x9b\x38\x98\x3d\x97\x57\x22\x5b\x2b\x14\x21\xa3\x17\x0b\x58\xab\x4c\xa6\x89\x11\x1a\x64\xee\x09\x60\xc4\x94\xc8\xfd\x9d\x54\x2c\x28\xd5\xda\xce\xcb\xba\x88\x34\xd4\x55\x12\x89\x75\xb1\xbb\x6c\xd0\x7e\x70\x32\xfd\x2b\x62\x93\x05\xf7\x43\x93\x7f\x7e\x80\xde\x12\xa6\x6b\x95\x57\x3d\xdb\x23\xf2\x3a\x7e\xbd\xab\x85\x1d\xb0\x82\xa8\x74\x28\xf8\x3a\xf1\x15\x1c\xd4\x68\x92\x51\xa0\xbf\x92\x7f\x7b\x96\x3e\x92\xca\xfc\xf4\xe3\x84\x9c\x96\x7f\x8f\x14\x3e\x55\x6d\xa9\x3b\xb6\xb7\xef\xc6\x2a\xdd\x6e\x41\xb6\xa1\xe4\x1b\x25\x3f\xb6\x9d\x52\x3f\x4c\x07\x92\x2d\xb1\x0d\x45\xcf\x64\x51\x24\xe7\x85\xb8\x45\x54\x59\xb6\xa1\xf0\x8b\x1a\x43\x35\x29\x6e\x11\xae\x2c\xdb\x50\xf8\x57\x91\x27\x6d\x61\x6e\x33\x3a\x63\xb6\x49\xd9\x3f\x93\x02\xdd\x96\xca\x88\x06\x33\xe9\xf5\xcd\xa4\xec\xfb\x4b\xe4\x1b\x41\x56\x67\x89\x11\xce\x86\xc3\x90\x11\xdb\xfb\x49\x23\xd6\x65\xd9\x9a\x0e\xbb\x83\x53\x48\xc7\x36\x94\xfe\x33\x29\x64\x86\x19\x9f\x96\x9c\x36\xdb\x94\xf4\x65\xc7\x36\x8a\x32\x53\x35\xc9\x46\xfc\x2e\x76\x70\x2c\x3a\x35\xb3\xbd\xbf\x10\xbb\x71\x4e\xb3\x79\x86\x7e\x8f\x86\xaf\x7e\x7e\x63\xfa\x48\xb9\x50\x48\xbe\xbc\xc5\x73\xed\xd8\x46\xd2\x94\xef\x70\x0b\x22\x6f\x99\xd4\x6f\xd9\x7c\x17\xf0\x4e\x9a\xd8\xde\xef\x6d\x4c\x4e\x38\x54\x43\xf6\xf3\x0d\x91\xbf\x22\xdd\x90\xdc\x44\xb6\x19\x9a\xb4\x9f\x5d\x9c\x17\x23\xc6\x23\xd9\x64\xc4\x38\xce\x1e\x7f\x88\x9c\x95\x0f\xf9\x1a\x91\xbf\xdf\xd7\xfe\x87\xc8\xed\xfa\x71\x49\xed\x99\x0f\xe4\x07\xbb\x54\x47\xf2\xc1\x5a\x5d\x8a\x46\x8b\x31\xab\x64\xf2\x58\xfd\xc7\x56\x36\x22\x1b\xf1\x36\x96\x3c\xb1\x6a\x5c\x59\xf6\x97\x8d\xe9\x5f\xb1\x6e\x2c\xd8\x2f\x9c\x97\x09\xbb\xb0\x3c\xe2\xad\x6b\x4a\xfc\x7c\x7b\x7b\x53\x32\xc1\x3d\xd5\x94\x78\x5b\xb4\xdb\x9f\xb7\x6c\x4b\x46\xe9\x4c\x7c\xa2\xf5\x4c\x1b\x41\x05\x3b\x51\x0e\x11\x34\x8a\x61\xa1\x27\xee\x2d\x6a\x53\x35\x71\x90\xb7\x2a\x75\x92\xa1\xc8\xe0\x11\x72\xc4\xbf\x76\x1c\x91\x0d\x92\xeb\x60\xa6\x04\x2c\x57\x70\x8a\xaf\xd7\xc1\x0c\x43\x73\xc9\x18\x88\x2c\x7e\x9d\x6c\xe6\x48\xdb\xd5\x62\xd9\xd1\x30\x9a\x83\x19\xed\x8a\x8e\x88\x2f\x48\x64\xc4\x97\x4c\xe4\x17\x24\xdb\x38\x5a\x12\xd9\xbe\x20\xdd\xc5\xcc\x12\xe9\xee\x85\x07\x72\x3b\x3f\x0d\xe4\x76\xfe\x9b\x60\x26\x73\x68\x44\x8e\x26\xf3\xc8\x13\x7a\x7d\xb0\x02\x25\x0b\x74\x67\xa6\x04\x92\x61\xd5\xb9\xdf\x88\x3c\x22\xd1\x46\x98\xb6\x51\xa0\x44\x8f\x2c\x37\x16\xfb\xd0\x72\x3b\x44\xd8\xf2\xe3\x14\xb8\x24\x1c\xe6\x99\xeb\x23\x7c\x78\x43\xee\x54\xe7\x20\x9a\x06\xdf\xaf\x83\x99\x26\xab\x4f\x89\x7e\x3d\x00\x90\x7e\x79\x8f\x22\x36\x23\xc3\x11\xa4\xcc\x07\xab\xe3\x46\xec\x12\x51\xbb\xb0\xf4\x07\x88\x32\x5c\x13\x37\xd4\x2f\x8c\x2b\xf8\xcb\xde\x06\x57\xdb\x83\x59\x57\xd1\xfb\x51\x47\xc1\x51\x5b\x2e\x97\xfd\xbc\xae\x80\xf2\x6a\x90\x6e\xbf\xb0\x2e\x49\xf7\xa0\xd4\xf6\x9c\x5d\xfd\x5c\x76\x3e\x77\xa5\x32\x98\x79\xdb\x67\x69\x87\x7b\x0a\x8e\xf7\x05\x94\xc6\x0b\xa1\xc2\x3c\x8b\x7b\x6a\x44\x93\xb8\x12\xd4\xe9\xe8\x28\x34\xdc\x95\xa2\x4e\x47\x47\xe9\x82\x4f\xe7\xb4\x18\xb0\xea\x23\xce\xc5\x95\x2c\xe6\x90\x97\x26\x7e\x8a\x4b\x9e\x87\x27\xa5\xd4\x1a\x37\x3a\xe5\x26\x89\x42\x79\xd5\xd8\x78\x7a\xf8\xf1\x64\x8e\x73\xe1\x92\x47\xdd\xdc\xd8\x2c\x2e\x57\x40\x5d\x22\xda\x8f\xdd\x63\xf4\x84\xe9\x0f\x56\xf0\x98\xd4\xe9\x9c\xe8\xb0\x82\x53\x1c\x20\x61\xcc\xa6\xdc\xc8\xdb\xe6\x04\xa8\xcb\x81\x34\x51\x70\x2e\x80\x0e\xc3\x22\x03\x53\x11\xcf\x46\x28\xd1\x24\x14\xcb\x28\xf9\xac\x6a\x40\x5c\x25\x65\x5d\x88\x39\xa8\xca\xe0\xd9\xa4\x55\x29\x75\x00\x85\xbc\x10\x60\x64\x29\xe2\xb3\xea\x53\x4c\x56\xbe\xa7\xa0\x46\x3b\x31\x7d\xc5\xcf\x93\x46\x6f\x93\x22\xec\xd7\x3f\x7a\x42\x0c\x1e\x42\x3a\x8f\x07\x4d\xda\xca\x8b\x16\xe7\xbc\x8d\x72\xca\x2f\x28\xdb\xf7\xe6\x6f\xde\xac\x7f\x85\xd3\xd3\xfd\x08\xa3\xb9\xcd\xae\x46\x5b\xec\xb9\x9e\x04\x5e\xe4\xbe\x35\xc1\x0c\xa7\x37\xbb\x3a\xfe\x5d\xaa\x2c\x8c\x50\xd8\x71\x3f\xc3\xad\xfc\xcf\x3f\x34\x7a\xd6\x96\x6b\xc5\xc3\x8f\x3d\xda\x8b\xd6\x30\xf1\x7b\x47\x44\xca\xe3\x28\x7e\x45\x69\x9c\xc7\x9c\xf1\x1d\x0d\x2d\x3b\x18\x18\xe2\xaa\x16\xa9\xe1\xb8\x08\x11\xea\x30\x82\x87\x3a\xa2\xf0\x68\x5b\x99\x0d\x17\xf1\x64\xbe\x37\x3d\xfa\x74\xe3\xe7\x34\x9d\xcf\x51\x4d\x9f\xd8\xb8\x14\xee\x27\x36\x3e\xb9\x51\x62\xe3\xc7\xa9\xc4\x46\xc2\xa1\xcc\xae\xf0\xc0\x92\x89\xab\x61\xe1\xe0\xa9\xaf\x3b\xdd\xa7\x44\x40\x87\xa9\x80\xda\x7c\x20\xb3\x2b\xea\xbf\x28\x05\x71\xad\x5c\x76\x03\xfc\x3e\x4e\x4e\x38\xd2\xa7\x26\x7f\xc7\xe3\xc8\x60\xbf\xdf\x58\x4f\x6d\xf0\xd9\xbb\x0b\x0e\x73\x0a\x71\xef\x2e\xa4\x3b\x10\xe0\x53\x05\x09\xfc\xe7\xd5\x8b\x33\x14\xa6\x0e\xc3\xee\x90\x4c\xf0\x0e\x21\x16\x9c\xc0\x0a\x57\xe7\x1f\x70\xa9\xf8\x8f\x45\x68\xa0\x34\xd4\x4e\x37\x36\x2e\x56\x53\x04\xe1\x39\xbc\x7d\x77\xbe\x33\x82\x37\x8b\x57\x05\xa8\x08\xb0\x2c\x62\xc6\x97\x25\x4b\x77\xee\xe7\xd7\x30\xf2\x2b\xac\x54\x7c\x0b\x16\x8e\x62\x9c\x45\xa2\x88\x92\x08\x89\xf0\x4e\xb2\xbb\x53\xc7\xb8\xe6\x74\x60\x77\xac\xbc\x31\x1f\xdc\x9e\xba\xac\x53\x0f\x3f\x2e\xe1\xe1\x25\x66\x2a\xae\x4d\x28\x1e\x4d\xaa\xe1\x15\xbd\x7b\x3d\xdc\x78\x75\xba\x92\x5c\x50\x50\x39\x45\x9d\x21\x77\xa1\x0b\xb7\x1f\x26\x37\x4a\x27\x89\xda\x08\xea\xab\x34\x67\x30\x0e\x66\x58\x41\x52\xd7\x42\x65\xa1\x25\xcc\xfb\x2e\xcb\xdb\x25\x61\x14\x59\x98\xec\x7d\x93\xef\x80\xbd\x9e\xba\x4f\x17\x70\xeb\x76\x4e\x58\x1b\xac\x1b\xee\x72\xcc\x73\x64\xed\x8c\xf4\xb7\xfe\xa4\x37\xa3\x45\xa7\x8b\xb3\xfb\x8f\x2d\xbe\x71\xbb\x7b\x3d\x56\x70\x50\xc5\x74\x64\x33\xcb\x1b\x55\x0e\x72\x0b\x27\x08\xcd\xf5\x53\x5e\x0a\x05\xe7\x6d\x9e\x8b\x06\x28\xa5\xd8\xec\xea\x2e\xef\x28\x4d\x8c\x66\x08\xcf\xdb\xdc\xe6\x04\x6c\x0f\x99\x38\x3f\x94\x19\x06\x30\x90\x85\xdd\x74\x38\xd1\x1c\xf4\x71\x20\x44\xd3\xf8\x01\x91\xf7\xe1\xa0\x6d\xf6\x25\x91\x5e\x47\x1e\xdb\xa2\xa3\xc3\xfd\x99\xf7\xa7\x1e\x95\x1f\xbf\xfa\x74\x59\x87\x9e\xb4\xbd\x1f\x34\x95\x45\xc7\x9e\x58\xfc\x74\x69\x01\x0b\x35\x58\x58\x22\x18\xa7\xae\x71\x7e\x25\xd8\xd0\x36\x9a\x7d\xb0\xbf\x06\x19\xef\xc8\xee\xf2\x21\x92\x73\x28\xbd\x2d\xc3\x26\xd3\x69\x02\x8f\xdf\xd4\x92\x4d\xe7\xe0\xf2\xaa\xcb\xbf\xc1\x6c\x66\x0f\x7e\xbe\x35\x36\x31\x96\x57\x51\x0f\xf7\x04\xb2\xc3\xbe\x11\xb5\x77\x71\xab\xbc\xa8\x45\x7b\xc9\xe0\x0f\x83\x35\xcd\xfb\x15\x9d\x61\x2b\x60\xf5\xf7\x67\x94\xe1\x6e\x46\xb6\x09\x53\xbe\xd4\x16\x32\x06\x7b\xbb\xee\xbe\x68\x05\xa7\xee\x99\x67\xa4\x74\x62\x3b\x82\x0f\x73\x22\xd9\xdb\x68\x22\x9a\x86\x6b\xfd\xcc\xbb\x6a\x5e\x82\x9c\xf7\x93\xbb\x60\xf5\xd2\x95\x6d\x1e\x40\xe7\x0e\x90\x43\x45\xe2\xae\x41\x3f\x54\x1c\xbe\xaa\x3a\xd0\xac\xc7\xea\xc3\x3d\x58\x7f\xb0\x2e\x7c\x4b\x61\x20\x05\xfc\xa1\xc4\x77\x83\x8b\xc3\x9d\xc7\x7d\x6f\x3f\xa9\x74\xd6\xf3\x37\x1c\xcf\xf6\xdf\xd8\xa0\x3b\x8c\xc7\xbd\xa6\x7b\x98\xf2\x6c\xa0\x72\xce\xe3\x43\xde\x57\xe4\xbc\x41\x1f\x75\x30\xe9\x1d\xce\x33\x5f\x9c\xf6\xa6\xb3\xc8\xe7\x25\x91\xc3\xcb\xda\xd5\x88\x83\xe9\xc1\x61\x4b\x3c\xb7\xed\xf2\x3d\xcc\x27\xb1\xf3\xdb\x91\x83\xd0\x1d\x0a\xd4\x2f\x04\x6e\x2a\x0c\x3f\x37\x0a\xbb\x20\xe4\xc0\xea\x02\x30\x4f\x0a\xbe\x16\xbb\xf9\x6c\x97\x07\xad\xd1\x41\x9f\xed\x77\x49\xdf\xe9\x61\x4f\xf5\x19\x5e\xeb\xd8\x7e\xf8\x5c\x01\x4f\x67\x79\xa7\xcd\xcc\x81\xef\xbf\x22\xe8\xbb\x8a\xde\x1e\x99\xc3\x83\xee\x46\x00\x4f\xd5\x0f\xf8\x52\x05\x8f\xdb\xa2\x91\xa9\x3d\x3f\x7b\x13\xa3\x05\x6a\x0e\xd5\x05\xb7\x2a\xfe\x65\x42\x1c\xe6\x45\x95\x98\x9f\x7e\x64\x2f\x1e\x54\x17\xbe\xb0\x9f\x5f\x5a\xc5\x07\x6f\x31\x3a\x60\xf3\x41\xbc\xbb\x9b\x59\xf2\xe5\x8c\x7f\x37\xa3\x3f\x49\x93\x6e\xc1\xb0\xf6\xee\x9a\xe2\x09\x6a\x4a\x13\x2d\xc0\xc0\xbf\xfd\x1b\x8b\xb5\x32\xff\x07\xa7\xa7\x60\xe0\x5f\x23\xf2\x4f\x3f\x2e\x31\x93\x8d\xaf\x43\xf8\xc6\x47\x45\xd3\xd3\xbd\x91\xd3\xf3\xbd\x91\x07\x27\x6c\xfb\x19\xa7\x12\x56\x9f\x31\xe0\x53\x93\xd4\xda\xff\xf4\x6c\xe9\x89\xca\xb8\x0f\x72\x84\x52\x98\x6d\x95\xc1\x27\x69\xb6\xd0\x88\xb4\xba\xe4\xe6\x57\x28\xdd\x36\x02\x54\x05\x75\xa2\x64\xaa\x41\x2a\xb0\x9d\xaa\x54\x1b\x9b\xe6\xbc\x0c\x95\x67\xde\x27\x3a\xb0\xc4\x08\xde\xbe\xeb\xbf\x10\xdf\x44\x10\xda\x64\xe4\x91\xc7\x27\xe9\x4c\x60\xfb\x6d\xaf\x4f\x6c\x33\x7b\xc9\x57\x41\x64\x1c\xf6\xb1\x97\x83\xe4\x44\xb7\x52\x83\x90\x78\xf8\xda\x79\xc7\xc6\xdb\xd2\x93\x67\x73\xb8\xa4\x16\x27\x77\x89\x89\xa2\x90\xf2\x3f\x76\x7a\x2e\xba\xb2\xd8\x39\x30\x1f\xa1\xcb\x0d\xc1\x1e\xb8\x4c\xfe\x56\x28\xfd\x33\xb0\x8f\x26\xd3\x1d\x98\x74\xc3\x8f\x58\x72\xa7\xd2\x13\xef\x03\xc9\x81\x7f\x03\x30\x19\x48\x61\x1b\xa4\x49\x1c\x7d\xe1\x7d\x28\x5d\x67\xb2\x07\xa6\x1b\xf8\x56\x38\x87\x27\x72\x1f\x50\x37\xe2\x20\xe5\xbb\x2f\xc4\x54\x76\xff\x64\xd2\xd1\xef\x11\x56\xe7\xe9\x04\xb0\xb2\xeb\xdb\x8e\x41\xdb\x39\x32\x06\x97\x4f\x6a\x7b\xd0\x32\xf9\x5b\x81\x3d\x76\x82\x0b\xb9\xdd\x63\xfc\x9e\xf7\xa7\xb8\x7b\xc1\x8f\xdd\x99\x40\x8f\x8d\x38\x8e\x1d\x7b\xb1\x87\x1c\x17\xfb\x3d\xe4\x98\xfc\xad\xc8\x0d\x7a\x19\x2f\x20\x99\xee\xc2\x11\xdf\x28\x1a\xb9\x09\xe9\x89\xf7\x08\x25\xfb\x37\x01\xe5\xd6\x36\x3f\xc7\xa0\xb4\xe6\x8f\xa1\xb4\xad\xc5\x1e\x96\x96\xfe\xad\x60\x1e\xed\x92\x42\xdb\xce\x20\xf9\xa5\xd7\x28\xdd\x0b\x78\xd6\xa1\x09\xf4\x6a\xd7\x5d\x1d\x83\xcf\x3a\xd2\xe3\x47\x2e\x76\x77\x13\x66\xf0\x15\x24\x1a\xbc\xd1\xb1\xa1\x6a\xc0\xb8\xaf\x20\xab\xfe\x2b\xc8\x4b\xd3\xf0\xa7\x14\x58\x81\x89\x9f\x16\xa2\x0c\x07\x7d\x83\x09\x6e\x82\xff\x05\x00\x00\xff\xff\x76\x23\x4a\xaa\x2c\x2a\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 10796, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
