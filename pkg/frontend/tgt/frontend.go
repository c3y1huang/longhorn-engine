package tgt

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/longhorn/go-iscsi-helper/longhorndev"
	"github.com/longhorn/longhorn-engine/pkg/frontend/socket"
	"github.com/longhorn/longhorn-engine/pkg/types"
)

const (
	DevPath = "/dev/longhorn/"

	DefaultTargetID = 1
)

type Tgt struct {
	s *socket.Socket

	isUp         bool
	dev          longhorndev.DeviceService
	frontendName string
}

// New returns new Tgt with new socket and given frontend name
func New(frontendName string) types.Frontend {
	s := socket.New()
	return &Tgt{s, false, nil, frontendName}
}

// FrontendName returns the Tgt.frontendName
func (t *Tgt) FrontendName() string {
	return t.frontendName
}

// Init populates the server object and intialize a new frontend device
func (t *Tgt) Init(name string, size, sectorSize int64) error {
	if err := t.s.Init(name, size, sectorSize); err != nil {
		return err
	}

	ldc := longhorndev.LonghornDeviceCreator{}
	dev, err := ldc.NewDevice(name, size, t.frontendName)
	if err != nil {
		return err
	}
	t.dev = dev
	if err := t.dev.InitDevice(); err != nil {
		return err
	}

	t.isUp = false

	return nil
}

// Startup listen to the server endpont and starts the tgs target and initiator
func (t *Tgt) Startup(rw types.ReaderWriterAt) error {
	if err := t.s.Startup(rw); err != nil {
		return err
	}

	if err := t.dev.Start(); err != nil {
		return err
	}

	t.isUp = true

	return nil
}

// Shutdown closes the tgt initiator, target, updates the server object and
// close the socket server
func (t *Tgt) Shutdown() error {
	if t.dev != nil {
		if err := t.dev.Shutdown(); err != nil {
			return err
		}
	}
	if err := t.s.Shutdown(); err != nil {
		return err
	}
	t.isUp = false

	return nil
}

// State return the server status if is up
func (t *Tgt) State() types.State {
	if t.isUp {
		return types.StateUp
	}
	return types.StateDown
}

// Endpoint returns the device endpoint if server status is up
func (t *Tgt) Endpoint() string {
	if t.isUp {
		return t.dev.GetEndpoint()
	}
	return ""
}

// Upgrade creates new longhorn device and frontend, removes socket, populates
// device server object, start the socket servre, reload iscsi connection
// and starts the device
func (t *Tgt) Upgrade(name string, size, sectorSize int64, rw types.ReaderWriterAt) error {
	ldc := longhorndev.LonghornDeviceCreator{}
	dev, err := ldc.NewDevice(name, size, t.frontendName)
	if err != nil {
		return err
	}
	t.dev = dev

	if err := t.dev.PrepareUpgrade(); err != nil {
		return err
	}

	if err := t.s.Init(name, size, sectorSize); err != nil {
		return err
	}
	if err := t.s.Startup(rw); err != nil {
		return err
	}

	if err := t.dev.FinishUpgrade(); err != nil {
		return err
	}
	t.isUp = true
	logrus.Infof("engine: Finish upgrading for %v", name)

	return nil
}

// Expand the device for the given size if device server is up
func (t *Tgt) Expand(size int64) error {
	if t.isUp {
		return fmt.Errorf("cannot expand the active frontend %v", t.frontendName)
	}
	if t.dev != nil {
		return t.dev.Expand(size)
	}
	return nil
}
