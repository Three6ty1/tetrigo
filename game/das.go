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
		dasDelay:  12, // 60fps ebiten
	}
}

func (d *DelayAutoShift) InitiateDAS(isLeft bool, tick uint) {
	d.isLeft = isLeft
	d.startTick = tick
}

func (d *DelayAutoShift) IsDASActive(isLeft bool, tick uint) bool {
	if isLeft != d.isLeft {
		return false
	}

	// Deal with uint wrapping
	if tick < d.startTick && tick < uint(d.dasDelay) {
		return tick >= d.startTick+uint(d.dasDelay)-^uint(0)
	}

	return tick >= d.startTick+uint(d.dasDelay)
}
