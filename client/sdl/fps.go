package sdl

import "time"

type FPS struct {
	lasttime time.Time
}

func (f *FPS) Tick() int {
	now := time.Now()
	ns := now.Sub(f.lasttime).Nanoseconds()
	current := int(int64(time.Second) / ns)
	f.lasttime = now
	return current
}
