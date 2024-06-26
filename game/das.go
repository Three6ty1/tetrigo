package game

type DelayAutoShift struct {
	isLeft    bool
	startTick uint
	dasDelay  int32
}

func NewDelayAutoShift() *DelayAutoShift {
	return &DelayAutoShift{
		isLeft:    true,
		startTick: 0,
		dasDelay:  15, // 60fps ebiten
	}
}

func (d *DelayAutoShift) MovementKeyPressed(isLeft bool, tick uint) {
	d.isLeft = isLeft
	d.startTick = tick
}

func (d *DelayAutoShift) IsActivateDAS(isLeft bool, tick uint) bool {
	if isLeft != d.isLeft {
		return false
	}

	// Deal with uint wrapping
	if tick < d.startTick && tick < uint(d.dasDelay) {
		return tick >= d.startTick+uint(d.dasDelay)-^uint(0) // TODO: fix this seems unsafe
	} else {
		return tick >= d.startTick+uint(d.dasDelay)
	}
}

// remember wrapping tick based on limit of uint
