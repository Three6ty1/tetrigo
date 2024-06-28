package game

import "fmt"

type LockDelay struct {
	initiated bool
	moves     int32
	moveLimit int32
	startTick uint
	tickLimit uint
}

// 32 ticks = 0.5 seconds
func NewLockDelay(moveLimit int32, tickLimit uint) *LockDelay {
	return &LockDelay{
		initiated: false,
		moves:     0,
		moveLimit: moveLimit,
		startTick: 0,
		tickLimit: tickLimit,
	}
}

// If the lock is already initialised, will do nothing
func (ld *LockDelay) InitiateLockDelay(tick uint) {
	if ld.initiated {
		return
	}
	fmt.Printf("Lock Started %v\n", tick)
	ld.initiated = true
	ld.moves = 0
	ld.startTick = tick
}

func (ld *LockDelay) IncrementLockMovement(tick uint) {
	if ld.initiated {
		ld.moves++
		ld.startTick = tick
	}
}

func (ld *LockDelay) IsLockActive(tick uint) bool {
	if !ld.initiated {
		return false
	}
	if ld.moves >= ld.moveLimit {
		return true
	}

	if tick < ld.startTick && tick < ld.tickLimit {
		return tick >= ld.startTick+ld.tickLimit-^uint(0) // TODO: fix this seems unsafe
	} else {
		return tick >= ld.startTick+ld.tickLimit
	}
}

func (ld *LockDelay) ResetLockDelay() {
	ld.initiated = false
}
