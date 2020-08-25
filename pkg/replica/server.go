package replica

import (
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	// Initial state initial
	Initial    = State("initial")
	// Open state open
	Open       = State("open")
	// Closed state closed
	Closed     = State("closed")
	// Dirty state dirty
	Dirty      = State("dirty")
	// Rebuilding state rebuilding
	Rebuilding = State("rebuilding")
	// Error state error
	Error      = State("error")
)

// State type
type State string

// Server object
type Server struct {
	sync.RWMutex
	r                 *Replica
	dir               string
	defaultSectorSize int64
	backing           *BackingFile
}

// NewServer returns new Server
func NewServer(dir string, backing *BackingFile, sectorSize int64) *Server {
	return &Server{
		dir:               dir,
		backing:           backing,
		defaultSectorSize: sectorSize,
	}
}

// getSectorSize returns the existing backing.SectorSize if Server.backing is
// not nil and SectorSize is greater than 0, or returns the defaultSector size
func (s *Server) getSectorSize() int64 {
	if s.backing != nil && s.backing.SectorSize > 0 {
		return s.backing.SectorSize
	}
	return s.defaultSectorSize
}

// getSize returns the existing backing size if Server.backing is not nil and
// Size is greater than 0, or returns the given size
func (s *Server) getSize(size int64) int64 {
	if s.backing != nil && s.backing.Size > 0 {
		return s.backing.Size
	}
	return size
}

// Create new repliac with no head for the given size
func (s *Server) Create(size int64) error {
	s.Lock()
	defer s.Unlock()

	state, _ := s.Status()
	if state != Initial {
		return nil
	}

	size = s.getSize(size)
	sectorSize := s.getSectorSize()

	logrus.Infof("Creating volume %s, size %d/%d", s.dir, size, sectorSize)
	r, err := New(size, sectorSize, s.dir, s.backing)
	if err != nil {
		return err
	}

	return r.Close()
}

// Open creates new replica with no head for the exiting size
func (s *Server) Open() error {
	s.Lock()
	defer s.Unlock()

	if s.r != nil {
		return fmt.Errorf("Replica is already open")
	}

	_, info := s.Status()
	size := s.getSize(info.Size)
	sectorSize := s.getSectorSize()

	logrus.Infof("Opening volume %s, size %d/%d", s.dir, size, sectorSize)
	r, err := New(size, sectorSize, s.dir, s.backing)
	if err != nil {
		return err
	}
	s.r = r
	return nil
}

// Reload create new replica with no head and close existing replica fd
func (s *Server) Reload() error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Reloading volume")
	newReplica, err := s.r.Reload()
	if err != nil {
		return err
	}

	oldReplica := s.r
	s.r = newReplica
	oldReplica.Close()
	return nil
}

// Status returns replica status
func (s *Server) Status() (State, Info) {
	if s.r == nil {
		info, err := ReadInfo(s.dir)
		if os.IsNotExist(err) {
			return Initial, Info{}
		} else if err != nil {
			return Error, Info{}
		}
		return Closed, info
	}
	info := s.r.Info()
	switch {
	case info.Error != "":
		return Error, info
	case info.Rebuilding:
		return Rebuilding, info
	case info.Dirty:
		return Dirty, info
	default:
		return Open, info
	}
}

// SetRebuilding updates the given boolean to volume.meta and
// Server.Replica.info. This returns error when:
// * it needs rebuild but server state not open or not dirty
// * it does not need rebuild but server is already rebuilding
func (s *Server) SetRebuilding(rebuilding bool) error {
	s.Lock()
	defer s.Unlock()

	state, _ := s.Status()
	// Must be Open/Dirty to set true or must be Rebuilding to set false
	if (rebuilding && state != Open && state != Dirty) ||
		(!rebuilding && state != Rebuilding) {
		return fmt.Errorf("Can not set rebuilding=%v from state %s", rebuilding, state)
	}

	return s.r.SetRebuilding(rebuilding)
}

// Replica returns Server.Replica object
func (s *Server) Replica() *Replica {
	return s.r
}

// Revert to the given disk name
func (s *Server) Revert(name, created string) error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Reverting to snapshot [%s] on volume at %s", name, created)
	r, err := s.r.Revert(name, created)
	if err != nil {
		return err
	}

	s.r = r
	return nil
}

// Snapshot creates the snapshot and meta hardlinked to the head
func (s *Server) Snapshot(name string, userCreated bool, createdTime string, labels map[string]string) error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Replica server starts to snapshot [%s] volume, user created %v, created time %v, labels %v",
		name, userCreated, createdTime, labels)
	return s.r.Snapshot(name, userCreated, createdTime, labels)
}

// Expand creates expand disk, appends block index to location and update
// the size
func (s *Server) Expand(size int64) error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Replica server starts to expand to size %v", size)

	return s.r.Expand(size)
}

// RemoveDiffDisk removes the disk from the chain and its host files for the
// given name
func (s *Server) RemoveDiffDisk(name string, force bool) error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Removing disk %s, force %v", name, force)
	return s.r.RemoveDiffDisk(name, force)
}

// ReplaceDisk links source to target and remove source
func (s *Server) ReplaceDisk(target, source string) error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Replacing disk %v with %v", target, source)
	return s.r.ReplaceDisk(target, source)
}

// MarkDiskAsRemoved update disk.Remove to true for the given name
func (s *Server) MarkDiskAsRemoved(name string) error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Marking disk %v as removed", name)
	return s.r.MarkDiskAsRemoved(name)
}

// PrepareRemoveDisk ensures the given name is not the head disk and returns
// the PrepareRemoveAction
func (s *Server) PrepareRemoveDisk(name string) ([]PrepareRemoveAction, error) {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil, nil
	}

	logrus.Infof("Prepare removing disk: %s", name)
	return s.r.PrepareRemoveDisk(name)
}

// Delete closes all fd except the head and removes all host files
func (s *Server) Delete() error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Deleting volume")
	if err := s.r.Close(); err != nil {
		return err
	}

	err := s.r.Delete()
	s.r = nil
	return err
}

// Close all fd except the head and set Server.Replica to nil
func (s *Server) Close() error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}

	logrus.Infof("Closing volume")
	if err := s.r.Close(); err != nil {
		return err
	}

	s.r = nil
	return nil
}

// WriteAt writes the data at offset and increase revision counter
func (s *Server) WriteAt(buf []byte, offset int64) (int, error) {
	s.RLock()
	defer s.RUnlock()

	if s.r == nil {
		return 0, fmt.Errorf("Volume no longer exist")
	}
	i, err := s.r.WriteAt(buf, offset)
	return i, err
}

// ReadAt returns the length of byte array
func (s *Server) ReadAt(buf []byte, offset int64) (int, error) {
	s.RLock()
	defer s.RUnlock()

	if s.r == nil {
		return 0, fmt.Errorf("Volume no longer exist")
	}
	i, err := s.r.ReadAt(buf, offset)
	return i, err
}

// SetRevisionCounter writes revision counter to fd and update the Server
// revisionCache
func (s *Server) SetRevisionCounter(counter int64) error {
	s.Lock()
	defer s.Unlock()

	if s.r == nil {
		return nil
	}
	return s.r.SetRevisionCounter(counter)
}

// PingResponse returns errro if state not match condition
func (s *Server) PingResponse() error {
	state, info := s.Status()
	if state == Error {
		return fmt.Errorf("ping failure due to %v", info.Error)
	}
	if state != Open && state != Dirty && state != Rebuilding {
		return fmt.Errorf("ping failure: replica state %v", state)
	}
	return nil
}
