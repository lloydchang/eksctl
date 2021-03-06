// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/schema.json (45.076kB)

package v1alpha5

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x6b\x73\xdb\x38\x92\xdf\xf3\x2b\x50\xca\xd6\xee\xa4\x4a\xb4\x76\xe6\xea\x1e\x95\x9d\x73\x9d\x62\xe7\xe1\x4b\x62\xeb\xac\xcc\x4c\xdd\xda\x53\x15\x88\x6c\x49\x58\x93\x04\x07\x00\x6d\x2b\x3b\xf9\xef\x57\x78\x90\xe2\x03\x7c\x80\xa2\x13\xe7\x6a\xbe\xd9\x14\xd9\x68\xf4\x0b\x8d\xee\x46\xe3\x9f\x4f\x10\x9a\xfc\x89\xc1\x7a\xf2\x1c\x4d\x9e\xce\x02\x58\x93\x98\x08\x42\x63\x3e\x3b\x09\x53\x2e\x80\x9d\xd0\x78\x4d\x36\x93\xa9\x7c\x51\xec\x12\x90\x2f\xd2\xd5\x3f\xc0\x17\xfa\xd9\x9f\xb8\xbf\x85\x08\xcb\xc7\x5b\x21\x92\xe7\xb3\xd9\x3f\x38\x8d\x3d\xfd\xd4\xa3\x6c\x33\x0b\x18\x5e\x0b\xef\xaf\xff\x3e\xd3\xcf\x9e\xea\xef\x0a\x43\x4d\x9e\x23\x89\x07\x42\x93\x6c\xcc\x90\xa6\xc1\x2f\x58\xf8\xdb\xfc\x27\x84\x26\x09\xa3\x09\x30\x41\x80\x17\x9e\x22\x34\xf1\xf5\x47\xef\xe8\x66\x43\xe2\x4d\xe9\xb7\xce\xc9\xe5\x03\x65\x5f\xe7\x9f\x7e\x36\x7f\x7d\x9e\xee\xc7\x87\x35\x30\x06\xc1\x05\x0b\x80\x4d\x9e\xa3\xab\x46\x1c\xcc\x0f\xbf\xe6\xdf\xe2\x20\x50\x23\xe3\x70\x51\x9c\xc5\x1a\x87\x1c\xf2\x97\x02\xe0\x3e\x23\x89\x7c\x4f\x62\xec\xd3\x58\x60\x12\x73\xe4\x2b\x16\xa0\x04\x33\x1c\x81\x00\xc6\x11\x83\x10\x0b\x08\x90\xa0\xa8\x40\xab\x1c\xd0\xbd\x47\x62\x01\x61\x48\xfe\xe1\x6d\x45\x14\x7a\x87\x02\x7e\x52\x20\x44\x9d\x47\x75\xc2\x37\xb2\x0a\x62\xbc\x0a\xe1\xc3\x2e\xa9\xfc\x80\xd0\x84\x08\x88\xaa\x0f\x0b\x22\xc7\x05\x93\x63\x4c\xcb\xbf\x06\xb0\xc6\x69\x28\xe4\x0b\x93\xc2\x2f\x9f\x8b\xaf\xe5\x20\x30\x63\x78\x57\x82\x50\xa5\xb8\x42\x0c\xd1\x35\x0a\xf5\x94\x24\x19\x34\xca\xe8\x3b\x0e\x80\xae\xf6\x93\x46\x01\xf5\xf9\xaf\xdf\xcd\x52\x8e\x37\x30\xf3\xe5\xf3\x3b\xf9\xdc\x33\x92\xe0\x19\x10\xb3\xa7\xe6\x81\xa6\xb5\x07\xf7\x38\x4a\x42\xe0\xcf\x9e\x1d\xa1\x0f\x14\xf9\x5b\x4a\x39\xa0\x35\xa3\xd1\x73\xf4\x11\x27\xe4\xe3\x14\x7d\xc4\x69\x40\x84\xfe\x43\x6c\x21\x16\xc4\xc7\x82\x32\xf9\x40\x72\x8e\xd1\x30\x04\xf6\x1e\xc7\x78\x03\xea\xa1\x54\xab\x20\x0d\x81\x7d\x2c\x4f\xae\x43\x0a\xba\x26\xfb\x23\x46\x5b\x06\xeb\xff\xbc\x9e\x0c\x9e\xe4\xf5\xe4\xb8\x42\xb1\x1f\x67\xf8\xd8\x32\xf3\x1f\x7d\x1a\xc0\x31\x4e\xc8\x8f\x33\xf5\xd7\x34\x7b\x22\x29\x51\x7b\x56\x20\x4a\xe5\xb7\x1a\x7d\x2a\xbf\xe7\xa4\x32\xcf\x87\xaa\x7b\x51\x8e\xc7\xd4\x75\x60\xed\x3a\x69\xc8\x9c\xb1\xcc\x55\xe3\x5d\xc1\x5b\xf5\x5e\xaf\x07\x05\x65\x67\xf0\x5b\x4a\x18\x04\x65\x12\x45\x20\x70\x80\x05\xae\xd3\xa7\xc9\x38\xe0\x84\xfc\x0c\x8c\x6b\x94\xff\x69\xd3\x60\x8b\x11\x28\x99\x80\xd2\x0f\x99\x10\x96\xd0\xd2\xbf\xdc\x70\x5f\x84\x47\x84\xce\x6e\xbf\xc7\x61\xb2\xc5\xff\x5a\x34\x1e\xbf\x3e\xb1\x98\x91\x09\xbe\xc5\x24\xc4\x2b\x12\x12\xb1\xfb\x3b\x8d\xbf\x9e\xfd\xb2\x62\xe7\xdb\x16\x4c\xe4\xb0\xfc\xb5\x9a\xc5\x65\xc5\xf4\xf1\x34\x49\x28\x13\x7d\xac\xdf\x33\x27\x93\xb4\x74\x34\x3b\x65\xfb\x62\xd0\x92\x26\xc6\x4e\xa5\x35\x66\x1b\x2c\x60\xc1\xe8\x9a\x84\xbd\x39\x68\xa7\xe0\xab\x12\xac\x83\x98\xb7\x21\xa2\x1f\xd7\x5e\x13\x61\x87\x40\x70\xe4\xc4\xf7\xb3\xf9\x7b\x3b\xa0\x1b\x12\x07\x0f\xac\x7c\x65\x33\xd2\xa9\x77\x91\xb2\xe3\xc1\x39\x0d\xe0\x35\xa3\x69\x72\x18\xd7\xde\x57\xa0\x8d\xe1\x34\x28\xed\x90\x10\x37\x0a\x3f\xa4\xa4\x36\xd7\x0d\x85\x3f\x89\x37\x5e\x9c\xbf\xf1\x0c\xe1\x38\x40\x57\x66\x66\x68\xff\x43\xfe\x11\xdc\x70\xcf\xfc\xac\xbe\xe3\x63\xe8\x91\x05\x93\xeb\xc9\x71\x15\x71\xa9\x3d\x0a\xbf\xda\xf7\x75\xa4\xae\x27\xc7\xf5\x49\x34\xab\x5f\xbe\x22\xb8\x88\xea\x7b\xd8\x2f\x21\x65\x70\xf1\x38\x22\x31\xaa\x2c\xbc\xa2\x0c\x91\x78\x4d\x59\x84\xe5\x23\x45\xc8\x4c\x15\x90\xf2\x20\x2d\xdc\xb6\x89\x88\x13\xbb\x3b\x47\xed\x29\x0b\x7d\x98\x98\x30\x72\x8b\x05\x18\xee\xf4\x63\xe5\xa2\xfc\x4d\x1b\x01\x71\x18\xd2\xbb\x6c\x63\x92\x4a\x83\x83\x30\x5a\xa7\x61\xb8\xf3\xcc\xc8\xb9\x93\x42\x62\x74\xb7\x25\xfe\x16\xc5\x54\x89\x1f\xda\x62\x8e\x68\x2a\x56\x34\x8d\x03\x24\x09\xc6\x62\x10\x08\xfb\x3e\x70\x3e\x55\x44\xc9\x40\xe8\x67\xd2\xe3\x99\xff\xb2\x44\x1c\xd8\x2d\xf1\x81\x23\xc2\x8d\x03\x1c\xa0\x5b\x82\xd1\xcf\x8b\x13\x04\x71\x90\x50\x12\x0b\xee\xc4\x90\xc7\x3b\x0b\x2b\x4f\x39\xf8\x0c\x04\x7f\x19\xfb\x6c\x97\xcd\xa1\x07\x5b\x97\xb5\xcf\xec\xd0\x05\x16\x69\x4d\x47\x5b\x95\x7e\xa9\x3f\xb1\x82\xbb\x4d\x7c\x27\x58\x3f\x2f\x4e\x86\x7a\xf9\x6a\x35\x9c\x5a\x1d\x54\x9b\x59\xab\x2c\xc8\x15\x9c\x9b\x75\xc8\x6e\xd3\x5a\xd7\xc0\x16\xa7\xa6\xd5\x67\xb5\xbb\x8c\xad\xa2\x50\xe7\x64\xc5\x75\x19\x65\xff\x83\x11\x27\xd2\x5e\x19\x9d\x99\x4a\xa9\x5e\x01\x62\x90\x84\xd8\x87\x00\xdd\x11\xb1\x45\x86\x60\x68\xbe\x38\xeb\xbd\xf3\x71\x06\x6c\xdb\xf3\xbc\xcc\xf5\xa7\x47\x8c\xc3\x70\x77\xae\xb4\xb3\xc9\x99\x5a\x51\x1a\x02\x6e\xd0\x98\x24\x5d\x85\xc4\x77\x05\xe0\x24\xda\x65\x24\x9b\xc6\x1e\x85\xb5\x5b\x1a\x06\x3c\xb7\x77\x38\x21\xca\x54\x01\xcb\xad\x52\x66\xc8\x0a\x4b\x58\x6f\xf6\x0e\x02\x6e\x63\xb1\x74\x89\x7b\x30\x37\xd3\x36\x1a\xbc\xbc\x07\x3f\x95\xe0\x2e\x69\x08\xf3\xcb\xf3\x0e\xb7\xb9\x75\x43\x52\x81\xb6\x00\x16\x11\x2e\xed\x0c\x7f\x21\x97\x01\xcc\x76\x43\xa0\x9b\x15\x61\xee\xfb\x34\x2d\xcb\x2e\x72\xf6\x92\xf6\x54\x5a\x96\xa0\x1e\xb4\xf1\x31\x08\x1e\x40\xc0\x02\x84\x91\x88\x26\xed\xc1\xc5\xd9\xe9\xc9\x03\xe9\x5d\x65\xca\xfd\xa7\xd2\x2d\x35\x15\x78\x0e\xb2\x65\x9b\x7e\x8b\x1c\x8d\x68\x15\x70\x18\xa2\xb3\xf9\x7b\x84\x85\x60\x64\x95\x0a\x1d\x90\xc4\x99\x42\x3b\x9a\x81\x2e\x68\x0d\x7a\x5f\x91\xe8\x1e\x56\x00\xc7\x31\x15\xb8\x9c\xbb\x68\xa7\xc5\xc3\x45\x85\x0a\xb9\x18\x1b\x80\x7f\x7e\xb6\xcb\x39\x16\x02\xfb\xdb\x05\x0d\x89\x5f\xd3\x13\xbb\x09\x38\x8b\x43\x12\xc3\x29\xf5\xd3\x08\xe2\xda\x80\x36\x76\xa0\x44\x81\x47\x81\xf9\x46\xae\xbd\x7a\x5c\xf9\x97\xd8\x12\x9e\x79\xad\xd2\x4a\x2b\xe2\xbb\x78\xd9\xc3\x47\xe9\xa4\xc8\xfc\xf2\xfc\x71\x05\xf8\x42\xbc\x82\xf0\x9b\x15\xb6\x18\x47\x30\x34\xac\xd4\x08\x90\x27\xd8\x1f\x17\x6a\xe2\xbc\x84\xb8\xc1\x1f\xb0\x13\xaa\x99\xa7\xb6\xad\x91\xc0\x9b\x6f\x4b\x44\x9c\x96\x4e\x25\x44\x56\x19\xa8\xeb\xc9\xd4\x6e\xab\xdb\xb4\xbd\xc9\x36\x76\xc8\x47\xeb\xf6\x48\x31\x64\xcc\xe5\x32\x46\x04\x47\xc6\x9a\x19\x63\x86\xb2\xfd\xa7\x8a\x0d\x64\x21\x87\x21\x5e\xb4\x2b\xf4\x5e\xcb\xe9\xb2\x2a\xf3\x8d\x8b\x2a\x73\x74\x03\x9d\x84\x27\x03\x3e\x22\x33\x34\xbb\xa5\x87\x51\xa3\x9a\x23\xe1\xdb\x20\xd9\x88\xac\x62\xa2\x9d\x09\xb8\xaa\xbe\x30\xd8\x14\x62\x34\xdd\xe9\xb8\xd1\x8d\xb6\x41\x60\x4c\x90\xdf\x9c\xc5\x2b\x06\x85\x0e\xcb\x75\x8e\x64\x45\x0d\x53\x2c\x78\x3d\x88\x1d\xcb\x4b\x4f\x36\x10\x03\xc3\x61\x21\x0e\xea\xbe\xf9\xef\x05\xcc\xa6\x40\xe7\xf3\x0f\x7d\x2c\x92\xdc\x3f\xdd\xe1\xfe\xdb\x48\x27\x46\x64\xc0\x47\xb4\x48\xe7\xf3\x0f\xc8\x80\x2d\x9b\x6a\x44\x93\xf2\x02\xd8\xcf\x2e\x75\xc3\xb3\x11\xb7\xbf\xc5\xc7\x6c\x5c\x6b\xe0\xcb\x41\xd6\xc4\xc7\x02\xe6\xa9\xd8\x52\x46\xc4\xee\xd4\x92\x70\xea\xe7\xc9\x1f\xe2\xae\x67\xd1\xa7\xb1\xdd\x47\xff\xe6\x7c\x14\xab\xec\x24\xa9\xf9\x64\xfa\x90\x7a\x5a\x66\xaf\x15\xfd\xd1\x44\x1e\x31\xc0\x81\x47\xe3\x70\x37\x4a\x04\xa1\x07\x38\xab\xc0\xa7\xab\x18\x9c\x62\xc3\xc3\x96\xac\x86\xfc\x25\x88\x3b\xca\x6e\x1e\x6c\x99\xd2\x01\xe1\x47\x8f\xb1\x93\x44\x67\x6c\xa8\x4f\x73\x44\x6b\x9c\xa7\xe9\xe2\x00\x69\xe8\x88\x1b\x49\x71\x33\xc3\x2d\x80\x6c\xe2\xf8\xf3\xe2\xa4\x97\xf1\x4d\x05\x9d\x87\x21\x95\x2a\x7c\xb6\xb8\xfd\xb7\x41\x99\x0a\x9f\x04\x3d\x73\xc0\x1b\x22\xb6\xe9\xea\xc8\xa7\xd1\xef\x77\x80\x6f\x41\x4a\x00\xff\x5d\x57\x62\xfd\x9e\xdc\x6c\x7e\x4f\x05\x09\xf9\xef\x24\x89\x41\x1c\x9d\x2d\xce\xa1\x21\x4a\xe3\x37\x67\x64\x5a\x46\xaf\xe5\x71\xec\x56\xfb\x5e\x30\x7c\x72\x76\x7a\x79\x58\xac\xfc\x90\xa9\x0e\x2f\x3d\x58\x53\x86\xf6\xc2\x8a\xe4\x34\x10\xe6\x9c\xfa\x44\x6f\x7e\xa7\x08\x8e\x36\x47\x48\x50\x94\x72\xd0\x69\x2f\x0e\x09\x66\x52\xb2\xd4\xcb\x12\x40\x26\x6a\x46\xbe\x90\x84\x19\xef\x10\x0e\xbc\x2d\xad\x8b\x6f\x1f\x11\xfe\x82\x68\x59\x79\x4a\x06\xd7\x33\x59\xc1\xc5\xb8\x67\xad\x56\xc1\xd3\x6c\x31\xab\x3a\xcf\xe6\x22\x72\x5f\x26\xdc\xc8\xc1\x4f\xe5\x92\xae\x0b\x63\x46\x75\x62\xb6\x98\xe9\x54\xf6\x72\xf8\x18\x35\x09\x4b\x18\x78\x8a\xfa\x10\x20\x3d\x82\x2e\xa6\x58\xbe\x76\x16\xd6\xbe\xa0\xba\x67\x5a\x73\x0b\xba\xc5\x65\x69\xd3\xb0\x0a\x92\x99\x12\x60\x06\x08\x88\xd8\x02\xcb\x56\x85\x82\xa6\xc8\x99\xd4\x15\x6a\x5f\xe0\x33\x45\x62\x0b\x1c\x14\x90\x1b\xd8\x41\x80\x56\x3b\x34\xff\xbb\xfa\xce\xa7\xf1\x2d\xc4\x04\xe2\x52\x68\xad\x9b\x7a\x5f\x14\xb1\x81\x2b\x3f\x29\x15\x74\xa8\xe5\xab\x51\xec\x2d\xbc\xb4\x2f\x16\x3d\xc4\x7b\xda\xb2\xf0\x56\xcc\x4b\xdb\x62\xd7\x6a\x40\x46\xf4\x5d\x36\x21\x5d\xe1\xd0\x58\x56\xe5\x78\xe0\x30\x44\xfe\x96\x84\x99\x0b\x32\x2b\xdb\x64\x47\x97\xc6\x1d\x7e\xc9\xd3\xa9\x14\xdf\xf6\x0b\x85\xd5\xc8\x33\x5e\xe0\xab\x34\x43\xba\x96\x22\x8c\x0c\x8e\x28\xd1\x48\x1e\x39\xa9\x52\x2f\x18\xdd\xf9\x0c\xe7\x2a\x85\xb6\x79\x9d\xcd\xdf\x23\x46\x43\xf8\x0b\x47\xf3\xcb\xf3\x6c\xc5\x16\x14\xb1\x34\x46\x09\x0d\x38\xa2\xb1\xa0\x19\xce\x6e\xf3\x3d\x08\x76\xb7\x25\x86\x10\x7c\x41\xd9\x98\x15\xdf\x4b\x03\x73\x0c\xd7\x4d\x2f\x37\x8a\xe3\x2c\x0d\x41\x95\xf4\x69\x9c\x91\xf4\x1d\x43\x8a\xd5\xe9\x8c\xec\xf0\xca\x01\x74\x3e\x6c\x24\x97\x65\xee\x31\x1c\xe7\xd2\x85\x95\x99\x12\xf1\x2d\x4d\xc3\x20\x13\xac\x80\xa2\x58\xef\x43\x91\xaa\x04\x53\x99\x63\xa3\x76\x9a\x22\x10\xe4\x34\x39\x42\x67\x6b\x14\xd3\x58\x69\xe2\x2d\x09\x20\x98\x2a\x83\x95\xad\x78\x72\x71\x92\x1f\x66\xf1\xc7\x3b\x12\x86\x68\x05\x72\xac\xc0\x8d\x41\x8f\x04\x65\x2b\xa7\x1f\x71\xb0\xbd\x44\xc3\x9f\xb8\x3e\xcd\x24\xf0\x46\x4d\x71\xfe\xcb\x12\x31\xe0\x34\x65\x3e\xb8\x6d\x5e\x1c\x20\x3d\x48\x8a\xd3\x66\xc0\xad\x76\xad\xdd\x55\x19\x2f\x7c\xaf\xed\x07\x37\x22\x27\x04\x89\x37\x5c\x89\x4c\xc9\x6a\xe4\xa6\xc4\x6e\xa8\xfa\x19\xa9\x81\x83\xb4\xf8\x09\xb9\xc9\xee\xe5\x2f\xe8\xd4\x72\x6f\xa7\xe1\x51\x17\x68\x94\xc8\xfb\x36\x5d\xa9\x9a\x72\xe0\x48\x21\x8d\x72\x31\x2a\xac\xbb\x95\x45\xc1\xcd\x88\x8d\x30\x42\xcf\xa2\x92\x01\x35\x20\x4d\x98\xe6\xe0\xd4\x91\x54\x53\x94\x3f\x22\x25\x86\xc1\x1f\x29\xe3\xd7\x54\x27\x31\xaa\x51\x38\xc0\x77\xe9\x6b\x12\x86\x3a\x2d\x99\x41\x78\x4d\x7a\xd5\xf7\xad\x28\x15\x5c\x30\x9c\xd4\x77\x18\xa8\xd9\x41\xcc\x5e\x6e\x13\xb8\xab\xb3\x98\x0b\x1c\x86\xfa\x48\xc8\xff\xa4\xc4\xbf\xe1\x02\x33\x91\xb9\xf8\xf9\xa1\xa0\x0d\x11\x34\xe1\xb3\xa7\x24\x7f\xdf\xc3\xde\x6f\xf9\xfb\x9e\x79\xdf\x23\xb1\xb7\xa3\x29\xcb\x8e\x46\xba\x1d\x1c\xaa\x9d\x0b\x1a\x38\xea\xf5\xe4\xb8\x63\x5e\xcd\x07\x8a\x24\x07\x70\xd9\x2a\xb7\xd0\xf8\x22\x7b\xbb\x95\xc8\x2f\xf5\x31\xf6\x4b\x48\x68\x1b\x41\xd7\x61\x7a\x3f\x3e\xc1\x24\xd4\xeb\xc9\x71\x01\x87\xe6\xc9\x33\x48\x68\xbf\x89\x4b\x38\xdf\xf2\xa4\x9d\x4c\x16\x2b\x4f\x76\x2f\x23\xd3\x16\x1d\x1d\xc5\x96\x99\xa3\x90\x2a\x1a\x61\xcb\x78\x17\x4f\xcb\xab\x53\x56\x52\xe0\x5f\x13\x71\x91\xc8\x2d\xea\x3e\x51\xa8\x62\x1a\x21\x89\x6f\xe4\xef\x44\xd7\xa4\xca\xf7\x90\x9c\x1a\x27\x82\xb2\xdd\x11\xba\x7a\xad\x08\x89\x5e\xa7\x24\x90\x9a\xaf\xe9\x5a\xd0\xb7\xc2\xb1\xcf\x2e\x26\x7d\x51\xc4\x0b\x12\x51\xc7\xf9\x7a\x72\x5c\x9c\xd7\x5e\x0e\x32\x23\x5c\x29\x24\x2e\xd8\xe3\x26\x77\x69\x2f\x35\x0d\x7e\x4e\x53\xb9\xdc\x0e\x61\xb6\x22\x82\x61\xb6\x43\xff\xbd\xbc\x38\x9f\xfd\xef\xfc\xfd\xbb\xbc\x52\x98\x4f\x11\x4f\xfd\x2d\xc2\x1c\xa9\x68\x9e\xa5\x39\x02\x65\xaa\xa2\x5c\x95\x18\x13\x70\xcd\xdc\x3d\x24\x02\x16\x0f\x29\x23\x70\xed\x94\xf3\xc8\x01\x32\x1c\x91\x57\x38\x22\xa1\x63\x75\x6c\xf9\x88\x72\x02\x3e\x59\xef\xd0\x95\x9f\x72\x41\x23\x34\x7f\x7f\x56\x68\xa7\xa2\x9e\x79\x38\x22\x9e\x39\xd4\x3f\x7b\x36\x45\xd7\x2a\x74\xea\x71\x1e\x5d\x4f\xb2\xff\xe4\x5f\x94\xa1\x6b\x55\x7c\x49\xfc\xeb\x89\x93\x65\xcb\x90\xa8\xf7\x1c\xa8\x23\x70\x3d\x39\x2e\xa0\x2a\xa5\x7a\x8a\xfe\xfc\x5b\x4a\xc5\xdf\x32\xac\xf4\x7f\xc5\xa7\xd9\x13\xca\xcc\x43\x8d\xa5\xfe\xdb\x31\x88\xf5\x68\x1a\x50\xb4\x71\xf5\x1d\x89\x88\xd0\x87\x9c\xb5\x9b\xa6\x08\x4c\x7c\x34\xff\xfb\x9e\xbb\x92\x32\xdc\xc7\xa1\x6a\x1c\xf3\x89\xc6\xe0\xe1\x3b\xcc\xc0\xd3\x74\xd4\x3f\xb8\xad\x50\x7a\xd8\x1a\x17\xfb\x0c\x64\x8e\x3d\xd7\xb0\x6d\x5e\xb3\x03\xe0\x52\x81\x4e\x70\x82\x7d\x22\x1a\x75\x40\x22\xbc\x01\x76\x60\xc3\x86\x5c\x87\x1b\x5b\x36\xa8\x45\x38\xf6\x55\x33\x9a\x51\x33\x75\x8f\x7a\x47\xdb\xb9\x3b\x8c\xf0\xfd\x92\x7c\x6a\xa4\x48\x2b\x77\x22\x12\x0f\xfe\x76\xf4\xaa\x59\x93\x07\x31\x45\x33\x96\x06\x67\xd5\x42\x8d\x36\xfd\x34\xfe\xd9\x55\x96\x5c\x89\x73\xa8\xb9\x76\xde\x26\xbe\xb7\x7f\x3c\x7b\x9a\x72\xc8\x0e\xad\x7b\x26\xb6\xe4\xad\x29\xf3\x94\x88\xe2\x70\xdf\x41\xe0\x99\x8a\x29\xe6\xff\x3a\x29\xb0\xc1\xab\xa6\xc1\x83\x91\xb9\x9e\x1c\xd7\xe7\xa8\x7a\x5b\xb4\x20\x59\xe0\x85\x72\x12\x1b\x02\xde\xbc\x67\x93\x9d\x5c\x77\x97\xcb\x37\xdf\x5c\x44\xb5\xbb\x7c\x99\x86\x69\x04\x7d\xf4\xa4\x4d\x20\x37\x64\x83\x57\x3b\xe1\x18\x97\x6d\xf8\xaa\x80\xf5\x7f\xfc\x75\xac\x18\xec\xde\xd9\x69\x32\xba\xad\x8b\x74\xcb\xc2\x61\xb1\x38\x16\x03\x66\x27\x79\x45\x1e\xeb\x66\xbb\xd5\x82\x54\x45\xb0\xb2\x32\xd5\x1c\xc1\x0a\xf9\x19\x24\x0c\xb8\xf4\x5f\x11\x8e\xd1\xcb\xb7\x4b\xaf\xd6\x06\x06\x7d\xb8\x38\xbd\x40\x3f\xe3\x90\x04\x79\xd6\x3f\x8e\x70\x92\x40\x80\xd6\x04\xb4\x73\x1c\x20\xb1\x65\xf4\x4e\x02\x01\xc6\x68\xff\x62\xcd\x07\x43\xa0\xec\x46\x83\x60\xc4\xe7\x27\x34\x0c\xc1\x17\xe5\x52\xfd\x06\x3f\x7a\xc3\x70\x9c\x86\x98\x49\xf6\xf6\x76\xa7\x8b\x1f\x8d\xb9\x72\x44\x1a\xff\xaf\x5f\x58\xe4\xa4\x7f\x45\x6a\x58\x26\x33\xca\x2e\x5f\x25\x13\x56\x3b\x9d\x61\xd0\x7e\x60\xde\x1f\x42\x75\xe9\x51\xcd\x31\xf6\x0d\x75\x74\xaf\xc7\xad\x10\x09\x7f\x3e\x9b\xc9\xff\x8e\xf0\x1d\x3f\xc2\x11\xfe\x44\xe3\x23\x9f\x46\xb3\xf9\x2f\x4b\xd5\x89\xec\x55\xf6\xcd\x4c\xee\xb6\xb9\x98\xfd\xc4\x81\xa9\x7d\xf0\x0c\xdf\x71\x6f\x2f\x02\x1e\xe6\x9e\x99\x93\x9f\x0b\xd8\x91\x14\xf6\xfe\x7b\xfe\xae\x69\xec\x97\xd3\x2f\x84\xba\xdc\x23\xd5\x29\x57\x8f\x00\x64\x15\xc0\x3d\x42\xb1\x5f\xbe\xc2\x74\x8c\x8a\x41\x27\x89\xb7\xd4\x24\x8d\x22\xe5\x3a\x06\x71\x76\xaa\x0c\xdd\xc9\xd9\xe9\xa5\x63\xf4\xa2\xf8\x65\x99\x7d\x0f\x18\x58\x18\xd5\x02\xfe\x11\xa8\xf8\x1a\x81\x0a\xbe\x69\x5b\x39\x51\xcb\x1a\xd4\xd0\xb4\xaf\x06\xcd\x79\x05\xfa\x23\x8e\xd2\x4b\xcc\xbe\x68\x1c\x65\x45\x85\x08\x81\x51\xff\x06\x7a\xd6\x54\xe7\x76\xe7\x45\xf1\xd3\x56\xe5\x5d\xbe\xc9\xc3\xe1\xc0\x11\xe7\xdb\xac\x1f\x90\xae\x7f\x21\x7c\xe0\x86\xd5\x09\xb0\x75\xfa\x7e\x88\x39\x27\xfe\x3b\x8a\x83\x17\x38\x94\x9b\x09\x76\x8e\xa3\xc7\x29\x8e\x73\x53\xb0\x0f\x48\x25\x5b\x57\x06\x5f\xae\xab\x78\xa5\x18\xe4\xfe\x87\x3b\x29\x9d\x81\x37\x90\x53\xa5\x2f\x4e\xcf\x97\x07\xd8\xfa\xab\x13\x6d\x38\x71\x10\x30\xe0\xbc\x31\x26\x62\xac\x6e\xd6\x8e\x36\x88\xb9\x67\x3e\x79\xa6\x0b\x55\xa4\x10\x9c\x9e\x2f\x51\x48\xe9\x4d\xb9\x5f\xdb\x80\x74\x5b\xff\xd1\xaf\x27\xc7\xe5\x19\xa8\x88\x47\x37\x46\xdd\x27\x2b\x93\xf4\x84\x41\x40\xea\x75\x7e\x0e\xd4\x2d\xe8\xcb\xd5\x87\x7f\x41\x3f\xc5\xa1\x34\x37\x10\x74\x7a\xd6\x2f\x4f\x7e\xa8\xfb\xa4\xab\x94\x71\x81\x57\x21\x78\x09\x30\xe5\x70\xc6\x3e\x78\xd9\xc6\x9c\x7b\x69\x06\xde\x8b\x68\x00\xda\xb3\x9e\xa2\x5b\xb9\x21\x44\xea\x88\x9f\xa4\xc8\x07\x4f\xe2\x8f\xf2\xaf\x9c\x18\x55\x98\x4f\x6f\x4f\x7b\xac\xa9\x5c\x4f\x8e\x8b\x24\xd4\xab\x7c\xd7\xe4\xac\xac\x1d\x23\x9c\x0d\x2b\x7e\x91\x08\x12\x91\x4f\xd0\xe8\x3c\xf7\x09\x4f\xea\xbe\x95\x1c\x5d\xbd\x7c\xb1\x54\x89\xcb\x88\x7c\x52\xfb\x88\x61\x22\x02\x2b\xee\xd1\x0c\xaf\xca\xde\xaa\x0f\x83\x33\x74\x0e\xe3\x6e\x1d\x8b\xeb\xc9\x71\x75\x82\xcd\xcb\xe4\x03\xa4\x0a\xc6\x39\x43\x6c\x01\xbc\x60\xb0\x26\xf7\x0f\x02\x7a\xf4\xf4\x46\xae\x16\xa7\x84\xeb\xb3\xbe\xbd\x1b\x9f\xee\x29\x6d\x85\x61\x1d\xee\x26\x5d\x41\x08\xe2\xa5\x3a\x35\x52\x6d\xa1\xdf\x32\x96\x43\xd7\x2d\xb3\x7a\x91\x4f\x80\x3e\x9a\xe1\x3e\x1a\x27\xa5\xb2\x61\x21\x9f\x48\xbc\xf1\xc4\x16\x3c\xf3\x9e\x63\x7f\xe9\x86\x6d\x48\x1d\x6c\xbe\x20\x49\xa4\xf4\x15\x0c\xe6\x27\x73\x01\x83\xc1\xaf\x59\xfc\xbf\xf9\x2c\xd4\x82\x06\x7c\x01\x4c\xca\xcc\xb0\x64\xd4\xff\x97\x44\x16\xbd\x05\xc6\x48\x00\x2f\xb2\xc2\x9d\x13\x1a\x45\xd8\xb5\xdf\x7c\x49\x0e\x2f\x0c\x48\xf4\x51\x87\x79\x3e\xfe\x85\xa3\xbc\x2e\x28\x91\x1e\xa3\x7e\xdd\x49\xb8\x73\xa0\x5a\x5e\x35\x64\x23\xae\x4d\xf0\xad\x13\x4e\x58\x6d\xae\x8f\xd2\xbb\x07\x55\x62\x0e\x01\x5a\xc1\x9a\x32\xa8\xcc\x30\xb7\x93\xba\x9b\x1e\xd4\x3a\x2b\xf4\xa1\xe9\xc0\x21\x1a\xc8\xfa\x47\x42\xf4\x71\x25\x44\x8b\x27\x1d\x7b\x9e\x77\xdd\xe7\x46\x5f\x37\x35\x25\xff\x23\xcd\xaa\x71\xb7\x75\x39\xf8\x76\xb0\x67\x1b\x10\x8a\x3b\x5f\xb3\xa5\x66\xbf\xd8\x86\x46\x56\x07\x19\x46\x8e\x6c\xf4\x02\x6d\xa5\xa0\xce\xf9\x9a\x7e\xec\xdd\xbb\xab\x16\x18\x67\x17\x8b\xc6\xd8\x48\xab\x27\xa0\x3f\x7f\x1b\xf1\xb7\xb0\x3b\x3b\xed\xdd\x0d\xab\x06\xa1\xc7\xb6\xa3\xe5\xeb\x6f\xa2\xd2\xa0\x86\xb5\xfb\xb6\xa5\x34\xbc\xca\x5f\xa3\x5b\xcc\x08\x8e\xf5\xc9\xfa\xe7\xe8\xe3\xf5\x64\x93\xfc\x70\x3d\xf9\x88\x08\x47\xaf\x4d\xe7\xb3\x45\xca\x12\xca\x01\x2d\x97\xa7\xe8\x3b\x83\xdd\xb3\xa9\x7c\x97\xd0\xef\xcd\xbb\x0b\x46\x6f\x09\x27\x34\x86\x00\x49\x61\x90\x2f\xab\x57\xb8\x9f\xbd\xf2\x61\xcb\x68\xba\xd9\x26\xa9\x40\xf9\x86\x1e\xbd\x39\x35\xaf\x89\xec\xb5\x13\x1a\xaa\xc7\x6e\x87\x61\x6c\x93\xd1\x3e\x96\xce\x33\x6c\x92\x1f\xf4\x1f\xd9\xde\xa0\x7b\x7e\xc5\xcf\x09\xfd\xbe\xf6\xb9\x7d\xca\xc5\xaf\xb8\x5f\xff\xaa\x99\x0a\xa5\x2f\x45\xfd\xcb\x06\xc2\x14\xc4\x65\x93\xfc\x50\xfe\x0d\xe2\x34\xaa\xdf\x92\x54\x7d\x4d\x1a\x4b\xfa\x7d\xf5\x11\xf7\xeb\x8f\xc4\xf7\x45\xab\x58\xb8\x54\xe9\x49\x45\x46\xff\xa8\x82\xa9\x78\x08\x95\xa0\x4b\x79\xf6\xb6\x79\x97\xb7\xfd\xcd\xf1\x90\xa6\x10\x4c\xb3\x0f\xd5\x95\x4e\x6b\x8a\xbd\xd9\x43\xb6\x76\xab\x64\xb7\xcf\x2d\x4b\x4f\xf3\x92\x60\x5f\x6b\x9a\xf7\xc1\x75\xff\xa6\x4f\x62\xa4\xc5\xaf\x68\x4a\x2b\x75\xed\xc6\xfa\x6c\x4f\xed\x49\x86\xf6\xd0\x8e\xf9\x71\x94\x8b\x33\x4a\xa7\x39\x0a\x7d\xe9\xc4\x16\x0b\xd5\x26\x25\x4f\xb7\xa9\xb3\x1a\x75\x57\xbe\xe7\x1d\x1a\x83\xc7\x51\xc3\xd4\x6a\x03\x5e\xd8\xd3\x7b\x1d\x57\xc3\xce\x83\x88\xc4\x27\xd9\x6d\x95\x83\xbc\x9d\xec\xc8\xf0\xe8\xf1\xb5\xbc\xdd\x28\x8e\x77\xe8\xaa\x28\x67\xf9\x31\xe5\x7d\x9c\x7a\x5f\x8d\x32\x2b\xbe\xe9\x51\x5e\xfa\x7f\xf6\xb4\x30\x88\x47\xd7\x5e\x06\xc9\x2d\x20\x57\x42\xad\x1e\xae\x3e\x14\x99\xeb\xc9\xb1\x75\xba\x87\x9c\xf1\xb2\xf2\xdb\xc6\xc6\x11\x75\x49\xc5\x16\x4a\x72\x2e\x37\xbf\x45\x49\x45\x2b\xcc\x21\x40\xfb\xab\x96\xfa\x9f\x51\x3d\x60\x08\xbb\x06\xf5\xbc\x92\xe6\x51\xdf\x5b\xb0\x5f\x09\xd5\x21\x3d\xe7\x0e\x34\x3d\x83\xf8\x83\xba\xdb\x38\xc0\x7e\xb0\x8c\xc9\xb0\x9b\x6b\xdc\xc6\x92\xfb\xcd\x79\x10\xd0\x78\x91\x1d\x22\x73\xce\x27\x95\x3f\x1f\xa8\xf1\x6d\x2d\xf7\x2d\x72\xd2\xc2\xe6\x36\x2e\x39\x10\xb9\x95\x46\x23\x9a\x9d\xa6\x7b\x69\xf6\x75\x74\x6e\x36\xa6\x1b\x5e\xa3\x41\x69\x92\x83\x66\xeb\x12\xae\xce\xe2\x0d\x1b\x7a\x95\x19\x4e\x92\xf7\x50\x8f\xe4\xb9\x7c\xbb\x60\x70\x4b\xe0\x6e\x18\x88\x54\xd0\xa5\x8f\xc3\x81\xae\x84\x0f\x4c\x98\x8b\xc4\x87\x7d\xdf\x78\x2d\x73\xaf\xcf\x61\x35\x8c\xe8\xb0\x1e\xf8\xdd\xbd\x00\x16\xe3\xb0\xa5\x88\xa6\xf5\xfb\x35\x6f\xcc\xfd\xb6\x7e\x47\x22\xbc\x81\x17\x29\x09\x83\x81\x74\xbe\xbf\x6c\xee\xd2\x7e\xe0\x85\x5d\x25\xdc\xec\x92\xd5\x40\xc1\x06\x39\xb2\x28\x47\xb3\xcc\x57\x84\xa1\x42\xeb\x0a\xcb\xa7\x56\xad\xad\x92\xc9\x2e\x9e\x0f\x61\xed\xa4\xa9\x19\x7c\x72\xda\x0e\xa4\xc1\xae\x75\xa4\xf2\x1b\x6a\x93\x8b\xe1\x0b\x8b\xbd\x6f\xb2\x88\xe5\xcf\x1e\x93\xb3\x25\x77\xd9\x8c\x34\x37\xc4\x89\xd3\x68\xd5\x14\xe4\xa5\xf1\x29\xc8\xed\xee\x0b\xcc\xe1\xa0\x5a\xa0\x0c\xd0\x02\x98\x0f\xb1\xc0\x1b\x98\xaf\xe8\x2d\x1c\x0c\x97\x27\x54\x98\xb6\x99\x84\xc6\x4b\xc1\xb0\x80\xcd\xb0\x7b\x11\x13\x2a\x32\x91\x59\x50\x5a\xaf\x2e\x68\xc6\xc7\xcd\x74\x94\x04\xc5\xc6\xa7\x2e\xfa\x3b\x92\xb5\x75\x8e\xdd\xa4\x1c\xd1\x06\xd8\x77\x41\x57\x72\xe0\x7d\xb6\x37\xcf\xb0\xca\xc7\xfb\x72\x37\x87\xa6\x17\x6d\x83\xd5\x52\xa7\x95\x51\xae\x27\xc7\x65\x74\x2c\xc7\x54\x8a\x49\xca\xde\x3b\xb1\xb3\xd3\x47\x99\xe7\xd2\xc8\x01\x2f\x76\xe3\xce\xa2\x8e\xc8\x74\x0f\x31\xe9\xf7\x61\x29\xeb\x41\x03\x34\x6e\x58\xde\x51\x1f\x87\x87\x64\xf6\xcd\x15\x80\xb8\x82\x03\x92\x62\x1f\xe6\x37\x03\x1e\x32\xd5\x81\xb0\x0b\x7c\x15\x2c\x6d\xc8\xa9\x4b\x12\x2c\x55\x2f\xe1\x11\x68\xa0\x1b\xe9\x95\x30\x35\x9d\xad\x71\x44\xe3\x8d\x5a\x6c\xf7\x1d\x98\x11\x89\x07\xd7\x79\x8c\x3f\x60\x33\xb5\x9c\x6c\xf1\x5e\x35\xed\x44\xb6\x4a\xdf\x28\x06\xd1\xa7\xb1\x60\x34\xe4\x35\x5d\x68\x39\x81\xd0\x27\xdc\xd7\x17\x66\x83\x45\x5b\xbe\xe9\xb7\xfb\x0b\xe9\xb0\x9d\x97\xee\xe2\xfc\x16\x06\xad\xd0\xf9\xc7\x43\x53\xc6\x39\x80\x05\x16\x8d\x7b\xaf\x56\x1f\x41\xf5\xf8\x2c\xf5\xf0\x3e\xfb\x7a\xd5\x5b\x43\x85\x5e\x71\xaf\x91\x2c\x56\x6e\x35\x72\xa1\x9b\x38\x23\xef\x21\x94\x11\xd9\x9f\xab\x29\xaf\xf0\x2a\x05\x71\x48\xfc\xc4\x05\x7a\x49\x85\x2e\xea\xfd\xea\x9a\x0f\xaf\xd2\x28\x22\x22\xfb\xe2\x3d\x8e\xc9\x1a\x78\xf3\x19\x8a\x3e\x26\xfd\x44\x81\x34\x77\xc1\xf0\x2d\x7a\x15\xa6\xf7\x28\xca\x20\x67\x0b\xec\x6b\x22\x54\x13\x36\x44\x63\x64\xba\xb4\x39\xd9\xf1\xe1\xa3\x58\xb5\x49\x65\x63\x0f\x28\x83\x90\x03\xe9\x4e\xa2\x82\xa2\x1b\x80\x04\x09\x86\xfd\x1b\x44\xd7\x0a\xb3\xbf\x70\xc4\x77\xb1\x8f\x12\x46\xd5\x9e\xf7\x6f\xda\x06\x12\x8e\xe4\xbe\xef\x16\x87\xe6\x56\x60\x93\xe2\x23\xf1\x06\x79\xde\x86\x08\x4f\x7e\xe5\x09\xbc\x51\x13\xd5\x8f\x62\x2a\x80\x7b\x0c\xd6\x72\x55\x92\xc0\x9d\xe8\xf6\x78\x10\x7d\xa8\x2b\x74\xcb\x62\x62\x7a\xc6\xed\x7b\x9b\xde\x6d\x81\xa9\xf6\xae\x46\x1e\xb4\xe4\xe8\x4e\x0b\x80\xde\x40\x18\xa1\x4c\x1d\xf4\xc5\x2f\x6b\x57\x12\x3f\xc8\x98\x3d\xee\xae\xc4\xc1\x45\xdc\x7c\x90\xb8\x8f\xea\xca\xfd\x18\x4b\x7d\xa1\xf1\x13\xb4\x70\xdd\x58\x44\x03\x7d\xbf\x93\xcf\x40\xd5\x8b\x6d\x01\x05\x90\x84\x74\x87\x6e\x60\x87\x30\xdf\xbf\xeb\x44\xac\x87\x18\xb2\x5f\x51\xa8\xf4\xa2\x24\xe9\x0f\x25\x58\x66\xab\x4b\x6c\x74\xa6\x81\x1d\xca\xc0\x65\xb5\xc9\xaa\xd7\x0c\x9e\x55\xdb\x6c\x34\xb2\x09\xda\x28\xab\xa9\x4b\x57\x48\x49\x9f\xac\xd1\x66\xde\x6e\x5b\xdb\xb0\x42\x9f\xf8\x4c\xad\xca\x2d\x21\xa5\x0d\x92\x36\xaa\x7f\xb6\xf4\xcb\x63\x56\x5a\xc5\x17\xba\x1a\xc8\x18\x93\x5e\xae\x70\x4e\xfe\xec\x3a\x16\x73\xd1\xf1\xa3\xdc\xed\x9b\x5a\x8d\xf2\x6e\x3c\xbb\xc8\x31\xbb\x0f\xdf\xd4\x76\x44\x29\x17\x68\x05\xba\x2d\xa8\x39\x2a\x9a\xdf\x39\xa7\xdd\xa3\xac\x74\x2f\x4c\x41\xdf\xef\x53\x6e\xf3\x31\x45\x85\x23\xe2\xba\x83\x28\xdd\xb8\x15\x5a\x3e\x0a\x84\xad\x96\xcc\x8c\x32\xc6\xb9\x46\x65\x6a\x95\x90\xaf\x11\x46\xeb\x34\x0c\x77\x59\xf5\xfe\xb0\x63\x15\x43\xe1\xb6\xd8\x70\x27\x53\x98\xd1\x66\xda\x4b\x4f\x46\x31\x69\xc5\xdb\x01\xea\x51\xbf\xae\xd9\xbb\xdc\x3d\xd0\x1f\x7a\xc5\xb4\xd4\x6e\x25\x6a\xb2\x29\x34\x15\x49\x2a\x7a\xec\x51\xdb\x84\xeb\x42\x01\x41\x01\x61\xaa\xc7\xfe\x2e\xbf\xda\x23\x61\x54\xea\x02\x04\x59\x13\x6e\x24\x20\x4a\xd4\xf9\x54\xf4\x9d\xbe\x32\x79\x7f\xb7\x10\xf2\x75\xb9\x9b\x5b\xdd\xd1\x83\x8e\x5d\x10\xd2\xa3\xd9\x8f\x85\x2e\xe4\xd2\xba\x7b\x72\x45\x6d\xec\xaa\xad\x8b\x7e\x0f\x20\xaa\xb9\x56\x45\x35\x32\x47\xcb\x62\x27\xf3\x23\x74\x82\x63\x69\x7d\x30\x5a\x31\x1c\xfb\xdb\xa9\xba\x1e\x44\x5d\x4f\xa6\x36\x4d\x5b\x5c\x4a\xe9\xf5\xbe\xec\x69\x9c\xb1\xba\x2f\x43\x52\x9b\xf7\x03\x48\x73\x8e\x23\x90\x28\xfc\x74\xf9\x0e\xb5\xa1\xfe\x4a\xba\xdb\xf7\x38\x4a\x42\x98\x22\x9c\x24\x5e\x00\xb7\x4e\x74\x19\x6f\xa0\x11\x9a\x34\x19\xb2\xd9\xa4\x6c\x6a\x55\xe8\xb1\xfd\xb7\x00\x04\x26\xa1\x69\xd6\xfd\x5b\xad\xc1\x7e\xde\xd7\x1b\xe4\x1b\x55\xbf\x08\x07\x41\x71\xd3\x5e\xe8\xe5\x3d\xc4\x61\x7b\x28\x54\x4a\x66\xf4\xb2\xdc\x1a\xbf\xf9\xae\x06\xa5\x18\x07\xc8\xf3\x87\x2d\xa0\x0d\x11\x46\xc3\x50\x1a\x07\xc0\xcc\xf5\x1c\x19\xde\x95\x95\x80\xc8\x35\x37\xbb\x30\x49\x6b\xa2\xf4\x54\xff\xac\x62\x25\x10\x98\x2b\x4d\x23\xec\xbc\x9e\x8f\x87\x0a\x8e\x92\xbf\x75\xa3\xd3\x69\x2c\x20\xc2\xe4\xd0\xc0\x8d\x82\x61\x66\x51\xbc\x65\x4a\x4a\x81\x31\x63\xfe\x16\xc7\x1b\xc7\x93\x39\x87\x80\xee\x9c\xf7\x3a\x4c\xef\x0f\x5c\x96\x25\x2f\xf7\xeb\x62\x91\x95\x6a\x2f\xde\xc6\xc7\x3b\x26\xb9\x18\x4f\xf7\xd1\x8a\x99\xb3\x18\x3d\xd4\xd0\xdd\xfd\x8a\xb1\xd8\x3e\xca\x2d\xd1\xa5\xdc\x50\x92\x5b\x40\x0a\x43\x75\x0a\xcf\x24\x9e\x2a\x3b\x46\x73\x35\x8f\xfe\x41\xdd\x8f\x90\xed\x3d\x15\x95\x22\x1a\xcb\xf7\xa4\x8c\xad\x49\x1c\xa0\xc2\x1d\x3f\xa5\x08\x29\x4e\x92\x70\x67\x08\x79\x75\xad\x4e\x09\x78\x7c\xc7\x05\x98\xae\x70\x2b\xcc\xe1\x7a\xf2\xab\x13\x67\xbf\xea\x1c\xf4\x69\xa7\xc2\x3c\xca\x7d\xe4\xe4\x7c\xf4\x5f\xbf\xb6\x1e\xdc\x5e\x2e\xdf\xf4\x4b\xcc\xb4\x31\x53\x7e\x9e\x2d\x23\x99\x37\xbe\x5c\xbe\x51\xa1\xab\xfd\x1d\x53\x38\x15\x5b\x88\x05\xf1\xcb\x97\xbd\x77\xd3\xf9\x50\xf0\x9d\x5a\x92\xb2\x43\xcc\xea\x07\xc3\x70\x89\x92\xf4\x8e\x0c\xa6\x35\xfe\x2b\x5e\x9b\xb3\x01\xa5\x85\xb8\x64\x02\x0c\x7f\x37\x44\xfc\xd7\xfe\xa0\xc0\x73\xca\x36\x33\x56\xbb\x74\xa6\x87\xe5\xf9\x3a\x88\x75\x53\x9c\x37\x17\x17\xf6\x5c\xc9\x24\x88\x87\x59\xc8\x06\x42\x1e\xc1\xaf\x95\x82\x38\xad\x79\x53\x35\x7b\x6e\x5b\x1b\xab\xc4\xad\xf9\x0d\xad\xba\xff\xc5\xe3\x9b\xd5\x6b\x6c\xf6\x7d\xd1\xb4\x89\x7c\x98\xd8\x65\xf7\xa8\x25\xaf\x77\x09\x3e\x03\xc1\xcd\xa1\xbb\x5e\x65\x8c\x37\xb0\x9b\x5f\x9e\xf7\xaf\x5f\x34\xef\xf7\x4d\x8a\x3b\x49\x53\x13\x2e\xe3\x07\x7a\xde\xbe\x5f\x22\xc8\xa9\x94\x5d\x72\x3a\x56\xa0\xa7\x09\x7a\x89\x57\x03\x1a\x0e\x17\x98\xd9\x60\x7b\x6a\x25\x3c\x31\x3a\x5b\x64\xed\xfc\x10\x89\x55\x3b\x5e\x14\x53\x81\x4b\x07\x41\x3b\x0b\x73\xda\xc1\x3c\xc9\x58\xfd\xf9\xc9\xe7\x27\xff\x17\x00\x00\xff\xff\xe1\x95\xd6\x61\x14\xb0\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x80, 0x12, 0x6e, 0xf0, 0xa7, 0xf8, 0x88, 0xc0, 0xce, 0xa2, 0x28, 0xff, 0x55, 0x96, 0xbf, 0x94, 0xba, 0x3, 0xad, 0x19, 0xad, 0xff, 0x4f, 0x37, 0xf5, 0x49, 0x55, 0xab, 0xe3, 0xa1, 0xff, 0xe9}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"schema.json": schemaJson,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
