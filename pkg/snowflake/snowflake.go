// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package snowflake

import (
	"fmt"
	"hash/fnv"
	"os"
	"sync"
	"time"
)

const (
	epoch        = 1735689600000 // 2025-01-01T00:00:00Z in millis
	workerBits   = 10
	sequenceBits = 12
	workerMax    = -1 ^ (-1 << workerBits)   // 1023
	sequenceMax  = -1 ^ (-1 << sequenceBits) // 4095
	timeShift    = workerBits + sequenceBits
	workerShift  = sequenceBits
)

// Node is a snowflake ID generator.
type Node struct {
	mu        sync.Mutex
	lastStamp int64
	workerID  int64
	sequence  int64
}

var global *Node

// Init initializes the global snowflake node with the given worker ID.
// If workerID is out of range [0, 1023], it auto-derives from hostname.
func Init(workerID int64) {
	if workerID < 0 || workerID > workerMax {
		workerID = autoWorkerID()
	}
	global = &Node{workerID: workerID}
}

// Next returns the next globally unique snowflake ID.
func Next() int64 {
	if global == nil {
		Init(autoWorkerID())
	}
	return global.next()
}

func (n *Node) next() int64 {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixMilli()
	if now < n.lastStamp {
		panic(fmt.Sprintf("snowflake: clock moved backwards by %dms", n.lastStamp-now))
	}

	if now == n.lastStamp {
		n.sequence = (n.sequence + 1) & sequenceMax
		if n.sequence == 0 {
			for now <= n.lastStamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		n.sequence = 0
	}
	n.lastStamp = now

	return (now-epoch)<<timeShift | n.workerID<<workerShift | n.sequence
}

func autoWorkerID() int64 {
	h, _ := os.Hostname()
	hash := fnv.New32a()
	hash.Write([]byte(h))
	return int64(hash.Sum32() & workerMax)
}
