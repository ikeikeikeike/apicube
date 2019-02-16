package dlm

import (
	"time"

	"github.com/gomodule/redigo/redis"
	redsync "gopkg.in/redsync.v1"
)

type (
	// DLM is called distributed lock manager.
	DLM struct {
		Pool *redis.Pool
	}
)

// Mutex returns MUTual EXclusion
func (d *DLM) Mutex(name string, expires time.Duration) *redsync.Mutex {
	rs := redsync.New([]redsync.Pool{d.Pool})
	return rs.NewMutex(name, redsync.SetExpiry(expires))
}

// Close dlm connection pooling
func (d *DLM) Close() error {
	return d.Pool.Close()
}
