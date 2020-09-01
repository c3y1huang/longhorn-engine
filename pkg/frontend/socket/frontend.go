package socket

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/longhorn/longhorn-engine/pkg/dataconn"
	"github.com/longhorn/longhorn-engine/pkg/types"
)

const (
	frontendName = "socket"

	SocketDirectory = "/var/run"
	DevPath         = "/dev/longhorn/"
)

// New return new Socket
func New() *Socket {
	return &Socket{}
}

type Socket struct {
	Volume     string
	Size       int64
	SectorSize int

	isUp         bool
	socketPath   string
	socketServer *dataconn.Server
}

// FrontendName returns "socket"
func (t *Socket) FrontendName() string {
	return frontendName
}

// Init populates the Socket object and shuts down the socket server
func (t *Socket) Init(name string, size, sectorSize int64) error {
	t.Volume = name
	t.Size = size
	t.SectorSize = int(sectorSize)

	return t.Shutdown()
}

// Startup starts the socket server
func (t *Socket) Startup(rw types.ReaderWriterAt) error {
	if err := t.startSocketServer(rw); err != nil {
		return err
	}

	t.isUp = true

	return nil
}

// Shutdown stops socket server if Volume exist and socketServer exist
func (t *Socket) Shutdown() error {
	if t.Volume != "" {
		if t.socketServer != nil {
			logrus.Infof("Shutdown TGT socket server for %v", t.Volume)
			t.socketServer.Stop()
			t.socketServer = nil
		}
	}
	t.isUp = false

	return nil
}

// State returns socket state if socket is up
func (t *Socket) State() types.State {
	if t.isUp {
		return types.StateUp
	}
	return types.StateDown
}

// Endpoint returns the socket path if socket is up
func (t *Socket) Endpoint() string {
	if t.isUp {
		return t.GetSocketPath()
	}
	return ""
}

// GetSocketPath returns longhorn socket path
func (t *Socket) GetSocketPath() string {
	if t.Volume == "" {
		panic("Invalid volume name")
	}
	return filepath.Join(SocketDirectory, "longhorn-"+t.Volume+".sock")
}

// startSocketServer creates socket directory abd remove existing socket path and 
func (t *Socket) startSocketServer(rw types.ReaderWriterAt) error {
	socketPath := t.GetSocketPath()
	if err := os.MkdirAll(filepath.Dir(socketPath), 0700); err != nil {
		return fmt.Errorf("Cannot create directory %v", filepath.Dir(socketPath))
	}
	// Check and remove existing socket
	if st, err := os.Stat(socketPath); err == nil && !st.IsDir() {
		if err := os.Remove(socketPath); err != nil {
			return err
		}
	}

	t.socketPath = socketPath
	go t.startSocketServerListen(rw)
	return nil
}

// startSocketServerListen listens to the socker path and starts infinite loop
// processing server requests
func (t *Socket) startSocketServerListen(rw types.ReaderWriterAt) error {
	ln, err := net.Listen("unix", t.socketPath)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			logrus.Errorln("Fail to accept socket connection")
			continue
		}
		go t.handleServerConnection(conn, rw)
	}
}

// handleServerConnection starts new server and handles the request read and
// write
func (t *Socket) handleServerConnection(c net.Conn, rw types.ReaderWriterAt) {
	defer c.Close()

	server := dataconn.NewServer(c, NewDataProcessorWrapper(rw))
	logrus.Infoln("New data socket connnection established")
	if err := server.Handle(); err != nil && err != io.EOF {
		logrus.Errorln("Fail to handle socket server connection due to ", err)
	} else if err == io.EOF {
		logrus.Warnln("Socket server connection closed")
	}
}

// DataProcessorWrapper object
type DataProcessorWrapper struct {
	rw types.ReaderWriterAt
}

// NewDataProcessorWrapper returns new DataProcessorWrapper
func NewDataProcessorWrapper(rw types.ReaderWriterAt) DataProcessorWrapper {
	return DataProcessorWrapper{
		rw: rw,
	}
}

// ReadAt returns the DataProcessorWrapper to read at the given offset to the
// data byte
func (d DataProcessorWrapper) ReadAt(p []byte, off int64) (n int, err error) {
	return d.rw.ReadAt(p, off)
}

// WriteAt returns the DataProcessorWrapper ti read at the given offset and
// writes to the backend
func (d DataProcessorWrapper) WriteAt(p []byte, off int64) (n int, err error) {
	return d.rw.WriteAt(p, off)
}

// PingResponse returns nil
func (d DataProcessorWrapper) PingResponse() error {
	return nil
}

// Upgrade not supported
func (t *Socket) Upgrade(name string, size, sectorSize int64, rw types.ReaderWriterAt) error {
	return fmt.Errorf("Upgrade is not supported")
}

// Expand not supported
func (t *Socket) Expand(size int64) error {
	return fmt.Errorf("Expand is not supported")
}
