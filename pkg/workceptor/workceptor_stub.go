//go:build no_workceptor
// +build no_workceptor

package workceptor

// Stub file to satisfy dependencies when Workceptor is not compiled in

import (
	"context"
	"fmt"

	"github.com/ansible/receptor/pkg/controlsvc"
	"github.com/ansible/receptor/pkg/netceptor"
)

// ErrNotImplemented is returned from functions that are stubbed out
var ErrNotImplemented = fmt.Errorf("not implemented")

// Workceptor is the main object that handles unit-of-work management
type Workceptor struct{}

// New constructs a new Workceptor instance
func New(ctx context.Context, nc *netceptor.Netceptor, dataDir string) (*Workceptor, error) {
	return &Workceptor{}, nil
}

// MainInstance is the global instance of Workceptor instantiated by the command-line main() function
var MainInstance *Workceptor

// RegisterWithControlService registers this workceptor instance with a control service instance
func (w *Workceptor) RegisterWithControlService(cs *controlsvc.Server) error {
	return nil
}

// RegisterWorker notifies the Workceptor of a new kind of work that can be done
func (w *Workceptor) RegisterWorker(typeName string, newWorkerFunc NewWorkerFunc, verifySignature bool) error {
	return ErrNotImplemented
}

// AllocateUnit creates a new local work unit and generates an identifier for it
func (w *Workceptor) AllocateUnit(workTypeName, signature string, params string) (WorkUnit, error) {
	return nil, ErrNotImplemented
}

// AllocateRemoteUnit creates a new remote work unit and generates a local identifier for it
func (w *Workceptor) AllocateRemoteUnit(remoteNode string, remoteWorkType string, tlsclient string, ttl string, signWork bool, params string) (WorkUnit, error) {
	return nil, ErrNotImplemented
}

// StartUnit starts a unit of work
func (w *Workceptor) StartUnit(unitID string) error {
	return ErrNotImplemented
}

// ListKnownUnitIDs returns a slice containing the known unit IDs
func (w *Workceptor) ListKnownUnitIDs() []string {
	return []string{}
}

// UnitStatus returns the state of a unit
func (w *Workceptor) UnitStatus(unitID string) (*StatusFileData, error) {
	return nil, ErrNotImplemented
}

// CancelUnit cancels a unit of work, killing any processes
func (w *Workceptor) CancelUnit(unitID string) error {
	return ErrNotImplemented
}

// ReleaseUnit releases (deletes) resources from a unit of work, including stdout.  Release implies Cancel.
func (w *Workceptor) ReleaseUnit(unitID string, force bool) error {
	return ErrNotImplemented
}

// GetResults returns a live stream of the results of a unit
func (w *Workceptor) GetResults(unitID string, startPos int64, doneChan chan struct{}) (chan []byte, error) {
	return nil, ErrNotImplemented
}
