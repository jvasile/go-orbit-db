// replicator is a package managing the replication of data among peers
package replicator

import (
	ipfslog "berty.tech/go-ipfs-log"
	"berty.tech/go-ipfs-log/identityprovider"
	"context"
	"github.com/berty/go-orbit-db/accesscontroller"
	"github.com/berty/go-orbit-db/events"
	"github.com/ipfs/go-cid"
	coreapi "github.com/ipfs/interface-go-ipfs-core"
)

// StoreInterface An interface used to avoid import cycles
type StoreInterface interface {
	OpLog() *ipfslog.Log
	Ipfs() coreapi.CoreAPI
	Identity() *identityprovider.Identity
	AccessController() accesscontroller.Interface
}

// Replicator Replicates stores information among peers
type Replicator interface {
	events.EmitterInterface

	// Stop Stops the replication
	Stop()

	// Load Loads new data to replicate
	Load(ctx context.Context, cids []cid.Cid)

	// GetQueue Returns the list of CID in the queue
	GetQueue() []cid.Cid

	// GetBufferLen Gets the length of the buffer
	GetBufferLen() int
}

// ReplicationInfo Holds information about the current replication state
type ReplicationInfo interface {
	// GetProgress Get the value of progress
	GetProgress() int

	// GetMax Get the value of max
	GetMax() int

	// GetBuffered Get the value of buffered
	GetBuffered() int

	// GetQueued Get the value of queued
	GetQueued() int

	// IncQueued Increments the value of queued
	IncQueued()

	// Reset Resets all values to 0
	Reset()

	// SetProgress Sets the value of progress
	SetProgress(i int)

	// SetMax Sets the value of max
	SetMax(i int)

	// DecreaseQueued Decrements the value of queued of i
	DecreaseQueued(i int)

	// SetBuffered Sets the value of buffered
	SetBuffered(i int)
}