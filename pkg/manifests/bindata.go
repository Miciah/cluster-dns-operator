// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/dns/cluster-role-binding.yaml (223B)
// assets/dns/cluster-role.yaml (397B)
// assets/dns/daemonset.yaml (2.563kB)
// assets/dns/daemonset.yaml~ (2.564kB)
// assets/dns/metrics/cluster-role-binding.yaml (279B)
// assets/dns/metrics/cluster-role.yaml (246B)
// assets/dns/metrics/role-binding.yaml (293B)
// assets/dns/metrics/role.yaml (284B)
// assets/dns/namespace.yaml (369B)
// assets/dns/service-account.yaml (85B)
// assets/dns/service.yaml (520B)
// assets/dns/update-node-resolver.sh (2.215kB)

package manifests

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

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x8e\x83\x40\x0c\x05\xd0\x7e\x4e\xe1\x0b\xc0\x6a\xbb\xd5\x74\x9b\xdc\x80\x48\xe9\xcd\x8c\x09\x0e\x60\xa3\xb1\x87\x22\xa7\x8f\x10\x4a\x45\x3a\x17\xfe\xff\xfd\x89\x25\x47\xb8\xce\xd5\x9c\x4a\xa7\x33\x5d\x58\x32\xcb\x23\xe0\xca\x77\x2a\xc6\x2a\x11\x4a\x8f\xa9\xc5\xea\xa3\x16\x7e\xa1\xb3\x4a\x3b\xfd\x59\xcb\xfa\xb3\xfd\x86\x85\x1c\x33\x3a\xc6\x00\x00\x20\xb8\x50\x04\x5d\x49\x6c\xe4\xc1\x9b\x2c\x16\xac\xf6\x4f\x4a\x6e\x31\x34\x70\x78\x37\x2a\x1b\x27\xfa\x4f\x49\xab\x78\xf8\xc4\xf6\xe7\xe3\xb6\x15\xd3\xa9\xa7\xe8\x4c\x1d\x0d\x3b\x74\x9a\x1d\xbe\xd3\xef\x00\x00\x00\xff\xff\xfa\x62\xe7\x50\xdf\x00\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 223, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xf6, 0x2a, 0x3b, 0x84, 0xd7, 0x3e, 0xc4, 0xe1, 0x70, 0x66, 0x31, 0xda, 0xc4, 0x2f, 0x53, 0x27, 0x29, 0x13, 0xfe, 0x80, 0x36, 0xc5, 0xa1, 0x70, 0xdc, 0x2d, 0xef, 0xcf, 0xe0, 0xc4, 0xeb}}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\x85\x75\xfd\x05\xd1\xa1\xb4\x14\xf4\x14\xf4\x1b\x67\x50\x96\xe4\x76\xad\xdd\x75\x4e\xe2\xeb\x51\x2e\x57\xa0\x8b\xa0\x9b\x19\xd9\xf3\x3c\x9e\x59\xc6\x3e\xbf\x2e\xcd\x03\xf6\xae\x0b\x12\x55\xfe\x80\x39\xab\xf4\xd9\x06\x2a\x1d\xb5\x98\xd4\xf8\x9b\x82\x55\xba\xf9\xc5\x3b\xd6\xa7\xf5\x39\x5d\x10\x34\x52\x50\x9f\x72\x16\xba\xa0\xcf\x5a\x21\x3e\xf1\x67\x9c\x47\xf1\x64\x6d\x81\xf7\xe9\x9c\xa9\xf2\x9b\x69\xab\xbe\x9d\x3c\xe7\xd3\x29\xe5\x6c\x70\x6d\x56\x70\xcf\x20\x63\x55\x96\xf0\x9b\x73\xd8\xca\x05\xbb\xa9\x3a\xee\x62\x63\x78\xa5\x3d\x5f\x61\xc3\xfd\xee\xc2\x1e\x37\x71\xa5\x28\x53\x3a\x02\xb7\x01\x90\xe0\xf2\x7b\xc1\xf1\x0d\xa1\x33\xc4\xb0\x32\xae\x0f\x84\x62\xa0\xc0\x1f\xcd\x8f\x5f\x73\x2c\xf6\x36\x7c\xa1\x04\x95\x02\xf7\xff\x00\x3f\x01\x00\x00\xff\xff\x76\x1b\x55\x2e\x8d\x01\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 397, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x84, 0xae, 0xd1, 0xba, 0xfa, 0x6b, 0xf8, 0x6e, 0x8d, 0x28, 0xc2, 0xa7, 0xaf, 0xc9, 0x3b, 0xc7, 0xcd, 0x80, 0xbe, 0xec, 0x98, 0xb4, 0x61, 0xa0, 0x9, 0xae, 0xa, 0xd8, 0xb2, 0x2e, 0x16, 0xf2}}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4d\x6f\x1b\x37\x10\xbd\xeb\x57\x0c\xe4\xab\x57\x1f\x76\x94\xb8\x04\x7c\x70\x25\xd5\x0e\x10\x3b\x42\xa5\xb6\x87\xa2\x10\x68\xee\xd8\x22\xc4\x25\xb7\x33\xb3\xaa\xf7\xdf\x17\x5c\x7d\xed\x5a\xb2\x12\x27\x17\x61\xc1\xf7\x66\xf8\xf8\x38\x9c\xd1\xd2\xfa\x54\xc1\x48\x63\x16\xfc\x14\xa5\xa5\x73\xfb\x27\x12\xdb\xe0\x15\xe8\x3c\xe7\xee\xaa\xdf\x3a\x03\xaf\x33\x3c\xaf\x7e\x39\xd7\x06\x41\xfb\x14\x9c\x7e\x44\xc7\xa0\x09\x81\x51\x40\x0b\x50\xe1\xc5\x66\xd8\xe2\x1c\x8d\x6a\x01\x08\x66\xb9\xd3\x82\xf1\x1b\x60\xbb\x5a\x7d\x23\xad\xac\xc1\x1b\x63\x42\xe1\xe5\x41\x67\xa8\x20\xf5\xbc\x41\x73\xb2\x81\xac\x94\x43\xa7\x99\xd7\x20\x97\x2c\x98\x25\x3e\xa4\x98\x18\xb2\x62\x8d\x76\x1b\xb6\x09\x5e\xb4\xf5\x48\xbc\xcd\x9e\x54\x4a\xeb\x19\x01\xce\xc0\x66\xfa\x19\xc1\xf2\x6b\xb5\x5b\x46\x85\x4f\x0a\xe7\x26\xc1\x59\x53\x2a\xf8\xfc\xf4\x10\x64\x42\xc8\xe8\x65\xc7\x12\xa4\xcc\x7a\x2d\x36\xf8\x7b\x64\x8e\x21\x1b\xfa\x6f\xda\xb9\x47\x6d\x96\xb3\xf0\x25\x3c\xf3\x57\x3f\x26\x0a\xb4\x8b\x33\x21\xcb\x74\xb4\xfa\x6f\x68\x9b\x40\x98\x7a\x6e\xc3\x3f\x3b\x58\xd3\x33\x57\x58\x62\x82\x7f\x6a\x9f\x43\xbb\x8b\x62\xba\x1b\x66\x77\x18\x08\x9f\xac\xc3\x7a\xc8\x2a\xb8\x22\xc3\xfb\x68\xe0\xee\xe4\xfb\xb3\xc7\x34\xf6\x39\x59\x93\x76\x28\x40\x16\xf9\x13\x2d\x0b\x05\xf5\x1d\x6a\x0c\x42\x9d\x7e\xf5\xae\x54\x20\x54\xec\x43\xf3\x40\xcd\x7d\x76\xbe\x4f\x02\x89\x82\xc1\xe5\xe0\xb2\x96\xe5\xf0\x06\xe2\xbd\x06\x09\x26\x38\x05\x7f\x8c\x26\xef\xcf\x94\x88\xc9\x8f\x66\x9b\x0d\xf7\xd9\xa2\x7a\xeb\x91\x79\x42\xe1\x11\x55\x8d\xbf\x10\xc9\x6f\x51\xea\x4b\x00\xf9\xda\x89\x18\x55\x36\x81\x4a\xca\x55\xff\xaa\xdf\x58\x66\xb3\xc0\x28\xe7\x6e\x36\x9b\xd4\x00\xeb\xad\x58\xed\x46\xe8\x74\x39\x45\x13\x7c\xca\x0a\xfa\xbd\xba\x5a\x24\x1b\xd2\xe3\x18\x17\xc6\x20\xf3\x6c\x41\xc8\x8b\xe0\x52\x05\xf5\x4d\x9f\xb4\x75\x05\x61\x0d\xad\xbb\x13\x4b\x38\x14\x72\x2c\xb1\xb3\x2b\x7c\xb7\x13\x0b\xd4\x4e\x16\xc7\xac\xe8\x5d\xf5\x7e\xd8\x8a\x8f\xbd\x13\x92\x07\x3f\x61\xc5\xa0\x76\xf1\x1c\x0a\x32\xc8\xaa\x51\xcb\xff\x16\xc8\xc2\xcd\xa3\x9a\xbc\x50\x30\xe8\x65\x8d\xc5\x0c\xb3\x40\xa5\x82\x4f\xbd\x7b\xfb\xaa\x8f\x2c\x8b\x47\x4c\xe8\x51\x9b\x24\xa7\xf0\x52\xbe\xa3\xa7\x54\xcf\xba\x56\xe9\x49\xe2\xc2\xb3\x04\x96\x14\x89\x1a\xeb\x8c\xa6\x20\x4c\x9c\x65\x41\x9f\xe8\x34\x25\x64\xbe\x56\xbf\xf4\x07\x1f\x1a\x3c\x71\x9c\x18\x9b\x2f\x90\x12\x2e\xac\x20\x5f\xcf\xbe\x4c\xe7\xe3\xe1\xe8\x6e\x3c\xff\x7d\x7a\x33\xff\xeb\xf3\xec\x6e\x7e\x33\x9e\xce\xfb\x17\x57\xf3\xdb\xe1\xfd\x7c\x7a\x77\x73\x31\xf8\x78\xbe\x67\x8d\x87\xa3\x6f\xf0\x0e\xf2\x0c\x7f\x1d\x7e\x57\x9e\xa3\xbc\x13\xd9\x1a\x27\x2b\x72\x16\x42\x9d\x5d\xc7\xf2\x54\xdd\x6e\xff\xe2\x53\xa7\xd7\xe9\x75\xfa\xd1\x84\xcb\xee\xa1\x0b\x48\x92\xc4\xa6\x78\x5d\x35\x32\x71\xdc\xcd\xc9\xae\xb4\x60\xfc\xee\x18\x92\x83\x90\x0d\x9e\x2c\xb1\x3c\x11\xb9\xc4\xf2\xbb\xbb\x5e\xe3\x7e\xb6\xbd\x2a\x43\x21\x6b\xf8\x87\x4b\xb3\xff\x46\x69\x7e\xd8\x97\xe6\xdb\xed\xff\x75\x83\xaf\x9d\xee\x2d\xa1\xd1\x9b\x6f\x0d\x80\xd4\xf3\x76\xd0\x8d\xf0\x49\x17\x6e\xeb\x6e\x1c\xc9\x53\x74\x68\x24\xd0\x5e\x47\x7c\x32\xe4\x51\x90\x3b\x36\x74\x03\x2b\x70\xd6\x17\x2f\x6b\x70\xc3\xaa\x8d\xd2\x5b\xd2\x06\x27\xaf\x5a\xe4\xc5\xb6\x69\xac\x8f\x7a\x30\xde\x8f\x8f\xb8\xf5\xea\xbd\xce\x55\xed\x9d\xc6\x3f\x11\x27\x9e\x29\x80\x15\xcc\x1a\x57\x91\xc0\x12\x4b\x05\xdb\xc1\x7b\xa4\x55\xbe\x82\x92\x13\x9e\x9e\x01\xa3\x21\x94\x93\x32\x24\x38\xa4\xca\x8d\x9d\x90\x33\x18\x3d\x4c\xc1\x23\xa6\x0c\x12\x22\x1f\x70\x85\x54\xfe\xb7\x40\xc2\x4e\xab\xa1\xb4\x5d\xfd\x37\xa2\xe0\xb0\xd3\x34\x3f\xd3\x2c\x48\xed\x9d\x98\x90\xc7\x6d\x02\x29\x18\xbf\x58\x16\x6e\xfd\x1f\x00\x00\xff\xff\x4a\x26\x6a\xde\x03\x0a\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 2563, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdc, 0x49, 0xae, 0x5a, 0x1f, 0xad, 0xee, 0x14, 0xa5, 0x19, 0x88, 0xc9, 0x2a, 0xa6, 0xc, 0xec, 0x81, 0xa5, 0x35, 0x59, 0x89, 0x34, 0x65, 0xd3, 0x24, 0xbf, 0xf7, 0x35, 0xe, 0x4f, 0x59, 0xe5}}
	return a, nil
}

var _assetsDnsDaemonsetYaml2 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4d\x6f\x1b\x37\x10\xbd\xeb\x57\x0c\xe4\xab\x57\x1f\x76\x94\xb8\x04\x7c\x70\x25\xd5\x0e\x10\x3b\x42\xa5\xb6\x87\xa2\x10\x68\xee\xd8\x22\xc4\x25\xb7\x33\xb3\xaa\xf7\xdf\x17\x5c\x7d\xed\x5a\xb2\x12\x27\x17\x61\xc1\xf7\x66\xf8\xf8\x38\x9c\xd1\xd2\xfa\x54\xc1\x48\x63\x16\xfc\x14\xa5\xa5\x73\xfb\x27\x12\xdb\xe0\x15\xe8\x3c\xe7\xee\xaa\xdf\x3a\x03\xaf\x33\x3c\xaf\x7e\x39\xd7\x06\x41\xfb\x14\x9c\x7e\x44\xc7\xa0\x09\x81\x51\x40\x0b\x50\xe1\xc5\x66\xd8\xe2\x1c\x8d\x6a\x01\x08\x66\xb9\xd3\x82\xf1\x1b\x60\xbb\x5a\x7d\x23\xad\xac\xc1\x1b\x63\x42\xe1\xe5\x41\x67\xa8\x20\xf5\xbc\x41\x73\xb2\x81\xac\x94\x43\xa7\x99\xd7\x20\x97\x2c\x98\x25\x3e\xa4\x98\x18\xb2\x62\x8d\x76\x1b\xb6\x09\x5e\xb4\xf5\x48\xbc\xcd\x9e\x54\x4a\xeb\x19\x01\xce\xc0\x66\xfa\x19\xc1\xf2\x6b\xb5\x5b\x46\x85\x4f\x0a\xe7\x26\xc1\x59\x53\x2a\xf8\xfc\xf4\x10\x64\x42\xc8\xe8\x65\xc7\x12\xa4\xcc\x7a\x2d\x36\xf8\x7b\x64\x8e\x21\x1b\xfa\x6f\xda\xb9\x47\x6d\x96\xb3\xf0\x25\x3c\xf3\x57\x3f\x26\x0a\xb4\x8b\x33\x21\xcb\x74\xb4\xfa\x6f\x68\x9b\x40\x98\x7a\x6e\xc3\x3f\x3b\x58\xd3\x33\x57\x58\x62\x82\x7f\x6a\x9f\x43\xbb\x8b\x62\xba\x1b\x66\x77\x18\x08\x9f\xac\xc3\x7a\xc8\x2a\xb8\x22\xc3\xfb\x68\xe0\xee\xe4\xfb\xb3\xc7\x34\xf6\x39\x59\x93\x76\x28\x40\x16\xf9\x13\x2d\x0b\x05\xf5\x1d\x6a\x0c\x42\x9d\x7e\xf5\xae\x54\x20\x54\xec\x43\xf3\x40\xcd\x7d\x76\xbe\x4f\x02\x89\x82\xc1\xe5\xe0\xb2\x96\xe5\xf0\x06\xe2\xbd\x06\x09\x26\x38\x05\x7f\x8c\x26\xef\xcf\x94\x88\xc9\x8f\x66\x9b\x0d\xf7\xd9\xa2\x7a\xeb\x91\x79\x42\xe1\x11\x55\x8d\xbf\x10\xc9\x6f\x51\xea\x4b\x00\xf9\xda\x89\x18\x55\x36\x81\x4a\xca\x55\xff\xaa\xdf\x58\x66\xb3\xc0\x28\xe7\x6e\x36\x9b\xd4\x00\xeb\xad\x58\xed\x46\xe8\x74\x39\x45\x13\x7c\xca\x0a\xfa\xbd\xba\x5a\x24\x1b\xd2\xe3\x18\x17\xc6\x20\xf3\x6c\x41\xc8\x8b\xe0\x52\x05\xf5\x4d\x9f\xb4\x75\x05\x61\x0d\xad\xbb\x13\x4b\x38\x14\x72\x2c\xb1\xb3\x2b\x7c\xb7\x13\x0b\xd4\x4e\x16\xc7\xac\xe8\x5d\xf5\x7e\xd8\x8a\x8f\xbd\x13\x92\x07\x3f\x61\xc5\xa0\x76\xf1\x1c\x0a\x32\xc8\xaa\x51\xcb\xff\x16\xc8\xc2\xcd\xa3\x9a\xbc\x50\x30\xe8\x65\x8d\xc5\x0c\xb3\x40\xa5\x82\x4f\xbd\x7b\xfb\xaa\x8f\x2c\x8b\x47\x4c\xe8\x51\x9b\x24\xa7\xf0\x52\xbe\xa3\xa7\x54\xcf\xba\x56\xe9\x49\xe2\xc2\xb3\x04\x96\x14\x89\x1a\xeb\x8c\xa6\x20\x4c\x9c\x65\x41\x9f\xe8\x34\x25\x64\xbe\x56\xbf\xf4\x07\x1f\x1a\x3c\x71\x9c\x18\x9b\x2f\x90\x12\x2e\xac\x20\x5f\xcf\xbe\x4c\xe7\xe3\xe1\xe8\x6e\x3c\xff\x7d\x7a\x33\xff\xeb\xf3\xec\x6e\x7e\x33\x9e\xce\xfb\x17\x57\xf3\xdb\xe1\xfd\x7c\x7a\x77\x73\x31\xf8\x78\xbe\x67\x8d\x87\xa3\x6f\xf0\x0e\xf2\x0c\x7f\x1d\x7e\x57\x9e\xa3\xbc\x13\xd9\x1a\x27\x2b\x72\x16\x42\x9d\x5d\xc7\xf2\x54\xdd\x6e\xff\xe2\x53\xa7\xd7\xe9\x75\xfa\xd1\x84\xcb\xee\xa1\x0b\x48\x92\xc4\xa6\x78\x5d\x35\x32\x71\xdc\xcd\xc9\xae\xb4\x60\xfc\xee\x18\x92\x83\x90\x0d\x9e\x2c\xb1\x3c\x11\xb9\xc4\xf2\xbb\xbb\x5e\xe3\x7e\xb6\xbd\x2a\x43\x21\x6b\xf8\x87\x4b\xb3\xff\x46\x69\x7e\xd8\x97\xe6\xdb\xed\xff\x75\x83\xaf\x9d\xee\x2d\xa1\xd1\x9b\x6f\x0d\x80\xd4\xf3\x76\xd0\x8d\xf0\x49\x17\x6e\xeb\x6e\x1c\xc9\x53\x74\x68\x24\xd0\x5e\x47\x7c\x32\xe4\x51\x90\x3b\x36\x74\x03\x2b\x70\xd6\x17\x2f\x6b\x70\xc3\xaa\x8d\xd2\x5b\xd2\x06\x27\xaf\x5a\xe4\xc5\xb6\x69\xac\x8f\x7a\x30\xde\x8f\x8f\xb8\xf5\xea\xbd\xce\x55\xed\x9d\xc6\x3f\x11\x27\x9e\x29\x80\x15\xcc\x1a\x57\x91\xc0\x12\x4b\x05\xdb\xc1\x7b\xa4\x55\xbe\x82\x92\x13\x9e\x9e\x01\xa3\x21\x94\x93\x32\x24\x38\xa4\xca\x8d\x9d\x90\x33\x18\x3d\x4c\xc1\x23\xa6\x0c\x12\x22\x1f\x70\x85\x54\xfe\xb7\x40\xc2\x4e\xab\xa1\xb4\x5d\xfd\x37\xa2\xe0\xb0\xd3\x34\x3f\xd3\x2c\x48\xed\x9d\x98\x90\xc7\x6d\x02\x29\x18\xbf\x58\x16\x6e\xb5\xfe\x0f\x00\x00\xff\xff\x3b\xc4\x00\xa4\x04\x0a\x00\x00")

func assetsDnsDaemonsetYaml2Bytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml2,
		"assets/dns/daemonset.yaml~",
	)
}

func assetsDnsDaemonsetYaml2() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYaml2Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml~", size: 2564, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb2, 0x96, 0x57, 0xd3, 0x88, 0xc6, 0x7b, 0x3f, 0xdb, 0x8a, 0x6f, 0xdf, 0x3c, 0xbc, 0x65, 0x96, 0xdd, 0xae, 0xf8, 0xcc, 0x6b, 0xb8, 0xba, 0xd6, 0x15, 0xc7, 0xde, 0xbf, 0x3e, 0x60, 0xdd, 0x92}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x4a\x04\x41\x0c\x86\xfb\x79\x8a\xbc\xc0\xae\xd8\x1d\xd3\xa9\x85\xfd\x09\xf6\xb9\x99\x9c\x1b\x77\x27\x19\x92\xcc\x16\x3e\xbd\x2c\x8a\x08\xe2\xb5\x81\x7c\xdf\xff\xad\x2c\x35\xc3\xd3\x36\x3c\xc8\xce\xba\xd1\x23\x4b\x65\x79\x4b\xd8\xf9\x95\xcc\x59\x25\x83\x5d\xb0\xcc\x38\x62\x51\xe3\x0f\x0c\x56\x99\xd7\x93\xcf\xac\x77\xfb\x7d\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xaa\xf8\xd4\x54\x38\xd4\x0e\x92\x8f\xcb\x3b\x95\xf0\x9c\x26\xf8\xd2\xbd\x90\xed\x5c\xe8\xa1\x14\x1d\x12\x3f\x7f\xdd\xb4\x51\x2c\x34\x7c\x5a\x4f\xfe\x7d\xf6\x8e\x85\x32\x68\x27\xf1\x85\xaf\xf1\x9b\x6c\xba\xd1\x99\xae\x87\xf9\x4f\xc7\x7f\x6b\x00\xb0\xf3\xb3\xe9\xe8\x37\xba\xd2\x67\x00\x00\x00\xff\xff\x5b\x52\x00\xaa\x17\x01\x00\x00")

func assetsDnsMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleBindingYaml,
		"assets/dns/metrics/cluster-role-binding.yaml",
	)
}

func assetsDnsMetricsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role-binding.yaml", size: 279, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0x95, 0x6f, 0xa4, 0xd5, 0xed, 0x48, 0x27, 0x41, 0x56, 0x5c, 0xea, 0x5c, 0x89, 0xdc, 0xc1, 0x44, 0x91, 0xd4, 0xb, 0x18, 0x85, 0x79, 0x75, 0xaa, 0x6e, 0xb5, 0x98, 0xbe, 0xc6, 0x33, 0x43}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x31\x4b\x34\x41\x0c\x87\xf1\x7e\x3e\x45\xe0\xad\x77\x5f\xec\x64\x5a\x05\x3b\x0b\x05\xfb\xec\xce\xdf\xdb\x70\x3b\xc9\x90\x64\x0e\xf4\xd3\x8b\x70\xb6\x0f\x3f\x78\xfe\xd1\xd3\x39\x23\xe1\xe4\x76\x22\x48\x81\x86\x46\xdb\x17\x0d\xb7\x8e\x3c\x30\x83\xd2\x28\x76\xe7\x01\x7a\x7e\x7d\xa7\x8e\x74\xd9\x83\xa0\x6d\x98\x68\x16\x1e\xf2\x01\x0f\x31\xad\xe4\x1b\xef\x2b\xcf\x3c\xcc\xe5\x9b\x53\x4c\xd7\xeb\x63\xac\x62\xff\x6f\x0f\xe5\x2a\xda\xea\xdf\xf0\xcd\x4e\x94\x8e\xe4\xc6\xc9\xb5\x10\x29\x77\x54\x6a\x1a\x4b\x37\x95\x34\x17\xbd\x14\x9f\x27\xa2\x96\x85\x78\xc8\x8b\xdb\x1c\xf1\x4b\x17\xb2\x01\xe7\x34\x5f\x6d\x40\xe3\x90\xcf\x5c\xc5\x0a\x91\x23\x6c\xfa\x8e\x3b\x6b\x1a\x88\x42\x74\x83\x6f\xf7\x74\x41\x96\x9f\x00\x00\x00\xff\xff\x9f\xa8\x4d\x6c\xf6\x00\x00\x00")

func assetsDnsMetricsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleYaml,
		"assets/dns/metrics/cluster-role.yaml",
	)
}

func assetsDnsMetricsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role.yaml", size: 246, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0xdb, 0xe0, 0x95, 0x65, 0xae, 0x53, 0x96, 0x3a, 0x5f, 0x5e, 0x8b, 0x69, 0xe2, 0x7d, 0x5, 0xbf, 0x1f, 0x3a, 0xf, 0xff, 0xd0, 0x6b, 0x23, 0x4f, 0xfd, 0x11, 0x7f, 0x57, 0xd4, 0x4a, 0x8b}}
	return a, nil
}

var _assetsDnsMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x4e\xc4\x40\x0c\x04\xd0\x7e\xbf\xc2\x3f\x90\x20\xba\xd3\x76\xd0\xd0\x1f\x12\xbd\x6f\xd7\x97\x98\x64\xed\x95\xed\x4d\xc1\xd7\x23\xa4\x48\x54\x20\x5d\x3b\x9a\xd1\x1b\xec\xfc\x41\xe6\xac\x92\xc1\x6e\x58\x66\x1c\xb1\xaa\xf1\x17\x06\xab\xcc\xdb\xc5\x67\xd6\xa7\xe3\x39\x6d\x2c\x35\xc3\x55\x77\x7a\x65\xa9\x2c\x4b\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xba\x69\xa3\x58\x69\xf8\xb4\x5d\xfc\x8c\xbd\x63\xa1\x0c\xda\x49\x7c\xe5\x7b\x4c\x55\x3c\x99\xee\x74\xa5\xfb\xcf\x14\x3b\xbf\x99\x8e\xfe\x8f\x9f\x00\x7e\xf9\xbf\x34\x1f\xb7\x4f\x2a\xe1\x39\x4d\x67\xfb\x9d\xec\xe0\x42\x2f\xa5\xe8\x90\x78\xf0\x65\x53\xe1\x50\x63\x59\x20\x7d\x07\x00\x00\xff\xff\xb9\xd9\xab\x8d\x25\x01\x00\x00")

func assetsDnsMetricsRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleBindingYaml,
		"assets/dns/metrics/role-binding.yaml",
	)
}

func assetsDnsMetricsRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role-binding.yaml", size: 293, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x7d, 0xc7, 0x45, 0x33, 0xc4, 0xd8, 0xf, 0x8d, 0x89, 0x8d, 0x6, 0x47, 0xa7, 0xa, 0x6b, 0x17, 0xf5, 0x5f, 0x5a, 0x2f, 0xd8, 0xf9, 0x6, 0x71, 0xaa, 0x78, 0x8d, 0xb5, 0x7a, 0xf6, 0x99}}
	return a, nil
}

var _assetsDnsMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x4e\xec\x40\x0c\x45\xfb\xf9\x0a\x6b\x5f\x9d\x7d\xa2\x5b\x4d\x8d\x44\x47\x01\x12\xbd\x77\xe6\x42\xac\x24\xe3\x91\xed\x04\xc1\xd7\xa3\xcd\x06\x89\xca\xf7\x1e\x59\x3e\xfe\x47\x2f\x3a\xc3\xa9\x01\x15\x95\xae\x5f\xd4\x4d\x17\xc4\x88\xd5\x29\x94\xbc\x18\x77\xd0\xe3\xf3\x2b\x2d\x08\x93\xe2\x84\x56\xbb\x4a\x8b\xc4\x5d\xde\x60\x2e\xda\x32\xd9\x95\xcb\x99\xd7\x18\xd5\xe4\x9b\x43\xb4\x9d\xa7\x8b\x9f\x45\xff\x6f\x0f\x69\x92\x56\xf3\x2e\x4a\x0b\x82\x2b\x07\xe7\x44\xd4\x78\x41\xfe\xe3\x1b\xa6\x8b\x1f\xd8\x3b\x17\x64\xd2\x8e\xe6\xa3\xbc\xc7\x50\x9b\x27\x5b\x67\x78\x4e\x03\x71\x97\x27\xd3\xb5\xfb\xed\xca\x40\xa7\x53\x22\x32\xb8\xae\x56\x70\x30\x87\x6d\x52\xe0\x7b\xf9\xfd\xf8\xde\xba\xd6\x5b\xd8\x60\xd7\x63\xf9\x03\xb1\xcf\x59\xfc\x1e\x3e\x39\xca\x98\x7e\x02\x00\x00\xff\xff\x29\x39\xda\x05\x1c\x01\x00\x00")

func assetsDnsMetricsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleYaml,
		"assets/dns/metrics/role.yaml",
	)
}

func assetsDnsMetricsRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role.yaml", size: 284, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0xf2, 0x4e, 0x40, 0x91, 0xd8, 0x5e, 0x1c, 0x98, 0xb6, 0x2f, 0x11, 0x2a, 0x15, 0x8f, 0xe4, 0x7c, 0xfe, 0xc6, 0x31, 0xf3, 0xb2, 0xa0, 0x38, 0xb2, 0x3f, 0x15, 0x5a, 0x33, 0x12, 0xd2, 0x88}}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcd\x4e\xc4\x30\x0c\x84\xef\x79\x8a\x51\x38\x2f\x3f\xd7\xbc\x03\x5c\x90\xb8\xbb\x8d\x97\x35\x4d\xed\x2a\x76\xcb\xeb\xa3\xb2\x15\xac\xb4\xc7\x68\x46\xf3\x7d\xf1\x24\x5a\x0b\xde\x68\x66\x5f\x68\xe4\x44\x8b\x7c\x70\x77\x31\x2d\xd8\x5e\xd2\xcc\x41\x95\x82\x4a\x02\x48\xd5\x82\x42\x4c\x7d\x7f\x02\xb6\xb0\xfa\x45\xce\xf1\x28\xf6\xa4\x56\xf9\xe4\xdc\x78\x0c\xeb\x05\x39\x27\x40\x69\xe6\xf2\x5f\x3b\x55\xf5\x04\x34\x1a\xb8\x1d\x13\x0f\x70\x0e\x6c\xd4\x56\x46\x18\x68\x33\xa9\xa8\xbc\xb0\x56\xd1\x4f\x98\x62\x5a\x07\x06\xd5\x59\x7c\x97\x42\x5c\x28\x8e\x82\xef\xf1\xdf\x38\x68\x11\xbf\xd7\xea\xab\x9e\x1a\x6f\xdc\x0a\xf2\x73\x3e\x98\xd4\x9a\x7d\xdf\x78\xcd\xa6\x12\xd6\x77\x62\x18\x9a\xd9\x84\xb3\x75\xbc\x73\xdf\x64\xe4\xd7\x6b\x0a\x1b\xbe\x78\x0c\x87\xec\x16\xe2\xbf\xbf\xbb\x1e\xed\x8e\x3a\xb6\xd5\x83\xfb\xcd\x70\x41\x8e\xbe\x72\x4e\x3f\x01\x00\x00\xff\xff\x82\x6d\x29\x03\x71\x01\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 369, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xab, 0x50, 0x84, 0x61, 0x5f, 0x41, 0xf4, 0x17, 0x3b, 0x6, 0x84, 0xc0, 0x5f, 0x4f, 0xbb, 0xd8, 0x1d, 0xae, 0x26, 0x3e, 0x1f, 0x29, 0x2c, 0x84, 0x6d, 0x5e, 0xc1, 0x87, 0x97, 0x5f, 0xc9}}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x09\xc4\x30\x0c\x05\xd0\xde\x53\x68\x81\x2b\xae\x55\x77\x33\x1c\xa4\x17\xf2\x0f\x11\xc1\xb2\xb1\x14\xcf\x1f\x02\xe9\x1e\xbc\xd3\xbc\x32\xfd\x31\x97\x29\x7e\xaa\xfd\xf2\x2c\x32\x6c\xc3\x0c\xeb\xce\xb4\xbe\xa5\x21\xa5\x4a\x0a\x17\x22\x97\x06\xa6\xea\xf1\x3a\x86\x28\x98\xfa\x80\xc7\x61\x7b\x7e\x9e\xba\x03\x00\x00\xff\xff\x8e\x2c\xf1\x2e\x55\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 85, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x12, 0x50, 0x4d, 0x67, 0x2f, 0x1b, 0x74, 0xa0, 0xa4, 0xbb, 0xa7, 0x59, 0xe9, 0x5a, 0xc6, 0xc1, 0x1a, 0xf8, 0x5f, 0xff, 0x5, 0xdb, 0xc, 0x10, 0x8b, 0xc1, 0x0, 0xcc, 0xf, 0x9f, 0x3a}}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x31\x6f\xe2\x40\x10\x85\x7b\xff\x8a\x27\xe8\x4e\xc0\x09\xdd\x51\x9c\xdb\xa3\x89\x52\x80\x14\x48\x3f\x5e\x4f\xcc\x8a\xf5\x8c\xb5\x33\x06\xf1\xef\x23\x4c\x42\x80\x14\x69\x56\xda\x7d\x9f\x3e\x3d\xbd\xdd\x47\xa9\x4b\xbc\x70\x3e\xc4\xc0\x05\x75\xf1\x95\xb3\x45\x95\x12\x87\x79\x31\x86\x50\xcb\x93\xe1\xb4\x8e\x02\x4f\x12\x55\x9c\x0c\x24\x35\x48\x44\x9d\x3c\xaa\x18\x28\x33\x8c\x1d\xe4\xc8\xbd\x78\x6c\xb9\xb0\x8e\x43\x59\x00\x63\x84\xd4\x9b\x73\x7e\x5a\xe3\x18\x53\x42\xc5\xa0\xde\xb5\x25\x8f\x81\x52\x3a\xa1\x25\xa1\x86\xeb\xd9\x00\x1b\x27\x0e\xae\x19\xd1\x1e\x8d\x40\xa7\xd9\xed\x2c\x9d\x0e\x95\x4a\xd4\x62\x05\x70\x09\x4a\x2c\xfe\x0c\x17\xa7\xdc\xb0\xaf\x87\xa7\x2b\x90\xd5\x35\x68\x2a\xb1\x5d\xae\xef\x05\x53\x0f\xdd\x8f\x92\x2f\xe8\x2a\xda\xfc\xbf\x15\xb5\xec\x39\x86\xdb\x36\xff\xe6\x8b\xbf\xdf\x54\x77\xd8\x83\x6a\x8c\xcd\x6a\xb9\x2a\xb1\x95\xa0\x6d\xcb\xe2\x38\xee\x58\x60\x97\xbf\x81\x6b\xa7\x49\x9b\x13\xde\x98\xbc\xcf\x8c\x86\x9c\xcf\x33\xb1\x50\x95\x3e\xf6\xfb\x84\x9e\xf9\x64\x97\xf5\x31\xc5\x68\xdf\x57\x9c\x85\x9d\x6d\x16\xf5\xf7\x4e\xcd\xcf\xa5\x47\xd7\xfc\xd7\xa8\x78\x0f\x00\x00\xff\xff\x82\x42\x75\xa4\x08\x02\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 520, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x69, 0xc5, 0xf1, 0xe, 0xc, 0x77, 0xe5, 0x78, 0xce, 0xfc, 0xc2, 0x41, 0xf8, 0x21, 0x87, 0x8a, 0xb7, 0x67, 0xdd, 0x48, 0x94, 0x63, 0x79, 0x69, 0x4e, 0x38, 0x53, 0x3c, 0xdb, 0xc7, 0x13}}
	return a, nil
}

var _assetsDnsUpdateNodeResolverSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xc1\x6e\xdb\x38\x10\xbd\xeb\x2b\x5e\x15\xa3\x4e\xda\x2a\xc9\x62\x8b\x1e\xda\xa6\x5b\x6f\xe3\xa0\x46\x9b\xd8\xb0\xdd\xed\x21\x30\x02\x5a\x1a\x59\xdc\xd0\x24\x4b\x52\x72\x8d\xc4\xff\xbe\x20\x25\xd9\x49\xdd\xdd\xc3\xa2\x3e\x48\x10\x3d\x7a\x33\x6f\xf8\xe6\x51\x07\x4f\x4e\xe6\x5c\x9e\xcc\x99\x2d\x22\x4b\x0e\x49\xa9\xa0\xb9\xa6\x9c\x71\x11\x45\xce\x30\x8d\xee\xdf\x6a\x6e\x91\x68\xdc\xe3\x3b\x33\x0b\x8b\x5b\x2e\x04\xee\xef\xe1\x4c\x49\x6f\xb0\x62\xdc\xbd\x01\x7d\xe7\x0e\xa7\x5d\x4c\xfb\xe3\xcb\x28\x1a\x8e\xfa\x57\x93\x8f\x83\x8b\xe9\xcd\x65\x6f\xfc\xa9\x3f\x3e\x8b\x95\x26\x69\x0b\x9e\xbb\x64\x41\x92\x0c\x73\x94\x25\x52\x65\x94\x18\xb2\x4a\x54\x64\xe2\xe8\xe3\x70\x32\x9d\xdc\x5c\x0c\x3e\xf7\xcf\xe2\x13\x72\xe9\x49\xa1\xac\xb3\x71\x34\xed\x5f\x8e\xf6\x96\x8f\xdd\x52\xc7\x51\x34\xb8\x98\x9c\x75\x5f\xa0\x0b\x43\x2c\x43\x62\x90\x30\x58\x32\x15\x4f\xc9\xe2\xed\xdb\xb7\x88\x3b\x77\x93\xfe\xf8\xaf\xc1\x87\xfe\x64\x13\x47\xd1\x01\x2e\xd9\x2d\x81\xc1\xd1\x52\x2b\xc3\xcc\x1a\x39\x17\x84\x15\x77\x05\x5c\x41\x50\x22\x43\x48\x10\xd6\xbb\x16\xcc\x39\xc3\xe7\xa5\x23\x7b\x1c\xa5\x1a\x49\x8e\x24\xd9\xad\x25\x4a\x8a\xb5\x4f\xb2\x2b\x7e\x13\xfb\xe7\x6d\xd1\x3e\xeb\xaa\xf0\x39\xea\x76\x65\x2a\x02\x32\x4a\x05\x33\x84\xa4\x07\x5b\xa5\x37\x5c\xdb\x08\xc8\x95\xf1\x4f\xe0\xd2\x03\xb4\x2c\xae\xdf\xcf\x36\x71\xf3\x1a\x70\x80\x0b\x72\x69\xd1\x72\xc4\x60\x84\xdc\xa8\x25\x52\x51\x5a\x47\x06\x99\xb4\xe0\x39\xb4\x21\x4b\xd2\x1d\xe3\x2b\x61\xe9\x09\x5b\xaa\xc8\x30\x01\x67\x38\xd9\x06\xc9\x29\x64\x0a\xdc\xbd\xc6\x60\x54\xbd\x7c\xe1\xaf\xaf\xc2\xf5\x25\x54\x45\x06\xd3\x0f\x23\x30\x99\x85\xf5\xed\xca\x31\xa6\x05\xc1\xad\x14\x04\xb3\x0e\x4a\x6e\xe1\x3c\x1f\x4f\x21\x23\x2d\xd4\x7a\x49\xd2\xd9\xba\xad\x9f\x4a\xb3\x36\x50\xd2\xf7\x96\x0c\x86\x9a\xe4\xc4\xb1\xf4\x16\x87\xc3\xc9\xe8\xb7\xdf\x8f\x90\xc0\x15\xca\x92\xaf\x46\x2a\xd7\xc0\xd9\x52\x6b\x65\x1c\xbe\x9c\x8f\x20\x14\xcb\xe6\x4c\x30\x99\x92\xb1\xa1\x26\x43\xdf\x4a\x6e\xc8\x6f\x7b\x5a\x70\xb9\xc0\xf9\xd5\x04\xae\x30\xaa\x5c\x14\xa1\xcc\x80\x92\x2e\x33\x7b\x76\xd8\xcd\xf8\x02\x89\x43\x0f\xef\xe3\xce\xdd\x55\xef\xb2\xef\x15\xd1\x1f\x6f\x62\x3c\xb7\x85\xcf\xe1\xdb\x5d\xa5\x9b\xe3\xce\xdd\x87\xcf\x5f\x26\xd3\xfe\xf8\xe6\x7c\x78\xd9\x1b\x5c\x6d\xe2\x6e\xc0\xa9\x7f\x5b\x9c\x5e\xef\x97\x41\xe1\xb9\x4b\xf5\xaf\xac\xeb\xff\xe2\x1d\x05\x40\xbf\x7f\xdc\x0b\xb0\x73\xf7\xc4\x37\xef\xfa\xd9\x6c\x13\xfe\x68\xe4\x07\x70\x6d\xcf\x0e\x3b\x87\x54\x31\xe1\xe1\x42\x10\x9f\x6d\xe2\xa3\xa3\x36\x20\xc7\xf5\x35\xe2\xce\x1f\x31\x12\xfa\x86\x53\x3c\x7d\xea\x03\x0f\xb8\xae\xa5\x8c\x44\x12\x4e\x31\x9b\xbd\xf1\xf3\x26\xb7\x3c\x9a\x39\xb8\x6e\x6a\x8c\x67\x67\x71\xe7\xae\x7d\x69\x1b\x35\x37\xc4\x6e\x9b\xa7\x9c\x37\xa5\x49\x8a\x9a\x5b\xe4\x95\xf3\x45\x67\xcc\x11\x76\x5e\x81\x30\xa4\x3c\xc7\x8a\xb0\x20\x87\x8a\x09\x9e\x3d\x18\x21\x1b\x5e\xfb\xea\x5d\x40\x08\x2f\x41\x94\x7b\x10\xab\x82\xa4\x2f\xd8\x10\xb8\x45\xaa\x0c\xf9\x51\x6b\x31\x54\xe9\xd8\x82\xa0\x0c\x98\xe6\x28\x25\xab\x18\x17\x6c\xce\x05\x77\xeb\x00\x3e\x71\x4c\x10\x48\x86\xf9\x43\xaa\x4a\x91\x79\xcb\xb4\xce\x37\xfb\x41\x1a\x9e\x07\x17\x6a\x71\xb9\x45\x46\x82\x1c\x65\x51\xdb\xd9\x44\x36\xdb\x18\xba\xf5\x6c\x96\x6c\xe2\xc7\xcd\x3c\xc0\x9f\x25\x17\x19\x18\x24\xad\x1e\x78\x59\x6d\x14\x0f\x29\xf9\xe1\x54\xa5\x41\x5a\x5a\xa7\x96\xdb\xea\x72\x2e\x1c\x19\xca\x3c\xab\x80\xb8\x30\xa4\x91\x54\x88\x0f\xd0\xb9\xfb\xd1\xda\x6b\xbb\x7b\x64\x7f\xef\xf6\x0c\xb0\xae\xab\xa7\x35\x85\xd9\xad\x1d\x7f\x97\xd0\x1b\x5f\x63\x77\x5b\x19\xee\x9c\xf0\x49\x4b\xf6\x91\x13\x36\x62\xd5\xb5\x5a\xdb\x90\x5a\x3c\xb3\xcd\x83\x30\x80\xd2\x42\x21\xa8\x69\x53\x87\xb6\xb7\xfd\x41\xc0\xbf\x50\x7c\xb7\xc7\xa9\x86\x6e\xc4\xb7\x93\x9f\xe7\x39\x1d\x9e\x0f\x5f\xff\x44\x86\xcc\xa9\x25\x4f\x99\x10\x6b\xef\xbd\xac\x52\x3c\x03\x93\x6b\x70\x99\x2a\x69\xb9\x75\x24\x1d\xe6\x54\xb0\x8a\x2b\xd3\x60\x8d\x49\x0b\x96\xd2\x4f\xf7\x6d\xa9\x32\x9e\x73\xca\x50\x91\xb1\x5c\x49\xaf\x10\x49\x94\x05\xb5\x78\xf7\xd3\x3f\x14\xbd\xb7\x53\xf7\xf7\xa8\xcf\xb4\xff\x8e\x6b\x79\xb5\x11\x5e\x97\x7e\x4c\x0c\x2d\x55\x45\xd9\x8e\x4d\xd0\x59\x6a\x88\x39\x3a\xa9\x85\x1b\xdc\x7a\x77\x5e\x22\x55\x7a\x8d\xb4\x28\x8d\x17\x6b\x98\x60\x2b\x88\x34\x5e\x9d\xe2\x69\xf8\x94\x88\x80\x52\xfa\x2f\x91\xf6\x60\x0c\x9d\xfd\x27\x00\x00\xff\xff\x30\x5b\x4f\x00\xa7\x08\x00\x00")

func assetsDnsUpdateNodeResolverShBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsUpdateNodeResolverSh,
		"assets/dns/update-node-resolver.sh",
	)
}

func assetsDnsUpdateNodeResolverSh() (*asset, error) {
	bytes, err := assetsDnsUpdateNodeResolverShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/update-node-resolver.sh", size: 2215, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x20, 0xc, 0xf4, 0xe9, 0x43, 0xee, 0x60, 0x7e, 0x65, 0x2d, 0x8c, 0x54, 0x79, 0x4, 0x95, 0xfc, 0xb, 0x71, 0x44, 0x99, 0x5f, 0x89, 0xff, 0x24, 0x10, 0x2c, 0x6, 0x15, 0x2c, 0xe9, 0x80, 0x8c}}
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
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,

	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,

	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,

	"assets/dns/daemonset.yaml~": assetsDnsDaemonsetYaml2,

	"assets/dns/metrics/cluster-role-binding.yaml": assetsDnsMetricsClusterRoleBindingYaml,

	"assets/dns/metrics/cluster-role.yaml": assetsDnsMetricsClusterRoleYaml,

	"assets/dns/metrics/role-binding.yaml": assetsDnsMetricsRoleBindingYaml,

	"assets/dns/metrics/role.yaml": assetsDnsMetricsRoleYaml,

	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,

	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,

	"assets/dns/service.yaml": assetsDnsServiceYaml,

	"assets/dns/update-node-resolver.sh": assetsDnsUpdateNodeResolverSh,
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
	"assets": {nil, map[string]*bintree{
		"dns": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"daemonset.yaml":            {assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"daemonset.yaml~":           {assetsDnsDaemonsetYaml2, map[string]*bintree{}},
			"metrics": {nil, map[string]*bintree{
				"cluster-role-binding.yaml": {assetsDnsMetricsClusterRoleBindingYaml, map[string]*bintree{}},
				"cluster-role.yaml":         {assetsDnsMetricsClusterRoleYaml, map[string]*bintree{}},
				"role-binding.yaml":         {assetsDnsMetricsRoleBindingYaml, map[string]*bintree{}},
				"role.yaml":                 {assetsDnsMetricsRoleYaml, map[string]*bintree{}},
			}},
			"namespace.yaml":          {assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml":    {assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml":            {assetsDnsServiceYaml, map[string]*bintree{}},
			"update-node-resolver.sh": {assetsDnsUpdateNodeResolverSh, map[string]*bintree{}},
		}},
	}},
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
