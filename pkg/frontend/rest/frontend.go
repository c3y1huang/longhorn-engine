package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"

	"github.com/longhorn/longhorn-engine/pkg/types"
)

const (
	frontendName = "rest"
)

var (
	log = logrus.WithFields(logrus.Fields{"pkg": "rest-frontend"})
)

// Device object
type Device struct {
	Name       string
	Size       int64
	SectorSize int64

	isUp    bool
	backend types.ReaderWriterAt
}

// New returns new Device
func New() types.Frontend {
	return &Device{}
}

// FrontendName returns "rest"
func (d *Device) FrontendName() string {
	return frontendName
}

// Init populates the Device object
func (d *Device) Init(name string, size, sectorSize int64) error {
	d.Name = name
	d.Size = size
	d.SectorSize = sectorSize
	return nil
}

// Startup starts new HTTP server
func (d *Device) Startup(rw types.ReaderWriterAt) error {
	d.backend = rw
	if err := d.start(); err != nil {
		return err
	}

	d.isUp = true
	return nil
}

// Shutdown updates the Device.isUp to false
func (d *Device) Shutdown() error {
	return d.stop()
}

// start listen and server a new HTTP server at localhost:9414
func (d *Device) start() error {
	listen := "localhost:9414"
	server := NewServer(d)
	router := http.Handler(NewRouter(server))
	router = handlers.LoggingHandler(os.Stdout, router)
	router = handlers.ProxyHeaders(router)

	log.Infof("Rest Frontend listening on %s", listen)

	go func() {
		http.ListenAndServe(listen, router)
	}()
	return nil
}

// stop updates Device.isUp to false
func (d *Device) stop() error {
	d.isUp = false
	return nil
}

// State returns the Device state
func (d *Device) State() types.State {
	if d.isUp {
		return types.StateUp
	}
	return types.StateDown
}

// Endpoint returns HTTP url if Device is up else returns empty string
func (d *Device) Endpoint() string {
	if d.isUp {
		return "http://localhost:9414"
	}
	return ""
}

// Upgrade not supported
func (d *Device) Upgrade(name string, size, sectorSize int64, rw types.ReaderWriterAt) error {
	return fmt.Errorf("Upgrade is not supported")
}

// Expand not supported
func (d *Device) Expand(size int64) error {
	return fmt.Errorf("Expand is not supported")
}
