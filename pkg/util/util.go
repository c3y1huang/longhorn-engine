package util

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	iutil "github.com/longhorn/go-iscsi-helper/util"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"

	"github.com/longhorn/longhorn-engine/pkg/types"
)

var (
	// MaximumVolumeNameSize is the maximum number of the charactors can a
	// volume contain
	MaximumVolumeNameSize = 64
	validVolumeName       = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]+$`)

	cmdTimeout = time.Minute // one minute by default

	// HostProc directory
	HostProc = "/host/proc"
)

const (
	// BlockSizeLinux is the default Linux block byte
	BlockSizeLinux = 512
)

// ParseAddresses returns the given host:port, host:port+1, host:port+2, port+2
func ParseAddresses(name string) (string, string, string, int, error) {
	host, strPort, err := net.SplitHostPort(name)
	if err != nil {
		return "", "", "", 0, fmt.Errorf("Invalid address %s : couldn't find host and port", name)
	}

	port, _ := strconv.Atoi(strPort)

	return net.JoinHostPort(host, strconv.Itoa(port)),
		net.JoinHostPort(host, strconv.Itoa(port+1)),
		net.JoinHostPort(host, strconv.Itoa(port+2)),
		port + 2, nil
}

// GetGRPCAddress trims the URL prefix, suffix and returns address:port
func GetGRPCAddress(address string) string {
	if strings.HasPrefix(address, "tcp://") {
		address = strings.TrimPrefix(address, "tcp://")
	}

	if strings.HasPrefix(address, "http://") {
		address = strings.TrimPrefix(address, "http://")
	}

	if strings.HasSuffix(address, "/v1") {
		address = strings.TrimSuffix(address, "/v1")
	}

	return address
}

// GetPortFromAddress returns the port for the given address
func GetPortFromAddress(address string) (int, error) {
	if strings.HasSuffix(address, "/v1") {
		address = strings.TrimSuffix(address, "/v1")
	}

	_, strPort, err := net.SplitHostPort(address)
	if err != nil {
		return 0, fmt.Errorf("Invalid address %s, must have a port in it", address)
	}

	port, err := strconv.Atoi(strPort)
	if err != nil {
		return 0, err
	}

	return port, nil
}

// UUID returns random UUID string
func UUID() string {
	return uuid.NewV4().String()
}

// Filter returns a list if item in list passed the given check
func Filter(list []string, check func(string) bool) []string {
	result := make([]string, 0, len(list))
	for _, i := range list {
		if check(i) {
			result = append(result, i)
		}
	}
	return result
}

// Contains returns if given list contains value
func Contains(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

type filteredLoggingHandler struct {
	filteredPaths  map[string]struct{}
	handler        http.Handler
	loggingHandler http.Handler
}

// FilteredLoggingHandler returns filteredLoggingHandler
func FilteredLoggingHandler(filteredPaths map[string]struct{}, writer io.Writer, router http.Handler) http.Handler {
	return filteredLoggingHandler{
		filteredPaths:  filteredPaths,
		handler:        router,
		loggingHandler: handlers.CombinedLoggingHandler(writer, router),
	}
}

// ServeHTTP respond to a HTTP request
func (h filteredLoggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if _, exists := h.filteredPaths[req.URL.Path]; exists {
			h.handler.ServeHTTP(w, req)
			return
		}
	}
	h.loggingHandler.ServeHTTP(w, req)
}

// DuplicateDevice creates device with same major/minor version for the given
// source
func DuplicateDevice(src, dest string) error {
	stat := unix.Stat_t{}
	if err := unix.Stat(src, &stat); err != nil {
		return fmt.Errorf("Cannot duplicate device because cannot find %s: %v", src, err)
	}
	major := int(stat.Rdev / 256)
	minor := int(stat.Rdev % 256)
	if err := mknod(dest, major, minor); err != nil {
		return fmt.Errorf("Cannot duplicate device %s to %s", src, dest)
	}
	if err := os.Chmod(dest, 0660); err != nil {
		return fmt.Errorf("Couldn't change permission of the device %s: %s", dest, err)
	}
	return nil
}

// mknod creates device
func mknod(device string, major, minor int) error {
	var fileMode os.FileMode = 0660
	fileMode |= unix.S_IFBLK
	dev := int((major << 8) | (minor & 0xff) | ((minor & 0xfff00) << 12))

	logrus.Infof("Creating device %s %d:%d", device, major, minor)
	return unix.Mknod(device, uint32(fileMode), dev)
}

// RemoveDevice removes host device
func RemoveDevice(dev string) error {
	if _, err := os.Stat(dev); err == nil {
		if err := remove(dev); err != nil {
			return fmt.Errorf("Failed to removing device %s, %v", dev, err)
		}
	}
	return nil
}

// removeAsync removes host path and sents to channel
func removeAsync(path string, done chan<- error) {
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		logrus.Errorf("Unable to remove: %v", path)
		done <- err
	}
	done <- nil
}

// remove with go routine and monitor over channel
func remove(path string) error {
	done := make(chan error)
	go removeAsync(path, done)
	select {
	case err := <-done:
		return err
	case <-time.After(30 * time.Second):
		return fmt.Errorf("timeout trying to delete %s", path)
	}
}

// ValidVolumeName validates the volume name
func ValidVolumeName(name string) bool {
	if len(name) > MaximumVolumeNameSize {
		return false
	}
	return validVolumeName.MatchString(name)
}

// Volume2ISCSIName replace the given name first "_" with ":"
func Volume2ISCSIName(name string) string {
	return strings.Replace(name, "_", ":", -1)
}

// Now returns current time
func Now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// GetFileActualSize returns the SIZE-byte blocks for the given file
func GetFileActualSize(file string) int64 {
	var st syscall.Stat_t
	if err := syscall.Stat(file, &st); err != nil {
		logrus.Errorf("Fail to get size of file %v", file)
		return -1
	}
	return st.Blocks * BlockSizeLinux
}

// ParseLabels validate the given list of labels and return an object maps
// value to the label name
func ParseLabels(labels []string) (map[string]string, error) {
	result := map[string]string{}
	for _, label := range labels {
		kv := strings.SplitN(label, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid label not in <key>=<value> format %v", label)
		}
		key := kv[0]
		value := kv[1]
		if errList := IsQualifiedName(key); len(errList) > 0 {
			return nil, fmt.Errorf("invalid key %v for label: %v", key, errList[0])
		}
		// We don't need to validate the Label value since we're allowing for any form of data to be stored, similar
		// to Kubernetes Annotations. Of course, we should make sure it isn't empty.
		if value == "" {
			return nil, fmt.Errorf("invalid empty value for label with key %v", key)
		}
		result[key] = value
	}
	return result, nil
}

// UnescapeURL remove the escape charactors for the given URL
func UnescapeURL(url string) string {
	// Deal with escape in url inputed from bash
	result := strings.Replace(url, "\\u0026", "&", 1)
	result = strings.Replace(result, "u0026", "&", 1)
	result = strings.TrimLeft(result, "\"'")
	result = strings.TrimRight(result, "\"'")
	return result
}

// CheckBackupType returns the URL scheme, for example "https"
func CheckBackupType(backupTarget string) (string, error) {
	u, err := url.Parse(backupTarget)
	if err != nil {
		return "", err
	}

	return u.Scheme, nil
}

// GetBackupCredential get s3 credentials from environment variables
func GetBackupCredential(backupURL string) (map[string]string, error) {
	credential := map[string]string{}
	backupType, err := CheckBackupType(backupURL)
	if err != nil {
		return nil, err
	}
	if backupType == "s3" {
		accessKey := os.Getenv(types.AWSAccessKey)
		if accessKey == "" {
			return nil, fmt.Errorf("missing environment variable AWS_ACCESS_KEY_ID for s3 backup")
		}
		secretKey := os.Getenv(types.AWSSecretKey)
		if secretKey == "" {
			return nil, fmt.Errorf("missing environment variable AWS_SECRET_ACCESS_KEY for s3 backup")
		}
		credential[types.AWSAccessKey] = accessKey
		credential[types.AWSSecretKey] = secretKey
		credential[types.AWSEndPoint] = os.Getenv(types.AWSEndPoint)
		credential[types.AWSCert] = os.Getenv(types.AWSCert)
		credential[types.HTTPSProxy] = os.Getenv(types.HTTPSProxy)
		credential[types.HTTPProxy] = os.Getenv(types.HTTPProxy)
		credential[types.NOProxy] = os.Getenv(types.NOProxy)
		credential[types.VirtualHostedStyle] = os.Getenv(types.VirtualHostedStyle)
	}
	return credential, nil
}

// ResolveBackingFilepath checks the backing file and returns the path or the
// given fileOrDirpath
func ResolveBackingFilepath(fileOrDirpath string) (string, error) {
	fileOrDir, err := os.Open(fileOrDirpath)
	if err != nil {
		return "", err
	}
	defer fileOrDir.Close()

	fileOrDirInfo, err := fileOrDir.Stat()
	if err != nil {
		return "", err
	}

	if fileOrDirInfo.IsDir() {
		files, err := fileOrDir.Readdir(-1)
		if err != nil {
			return "", err
		}
		if len(files) != 1 {
			return "", fmt.Errorf("expected exactly one file, found %d files/subdirectories", len(files))
		}
		if files[0].IsDir() {
			return "", fmt.Errorf("expected exactly one file, found a subdirectory")
		}
		return filepath.Join(fileOrDirpath, files[0].Name()), nil
	}

	return fileOrDirpath, nil
}

// GetInitiatorNS returns a path for namespace "/host/proc/PID/ns/"
func GetInitiatorNS() string {
	return iutil.GetHostNamespacePath(HostProc)
}

// GetFunctionName to show as part of log
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
