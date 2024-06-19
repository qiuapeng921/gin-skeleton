package helper

import "time"

// Duration 记录时间点
type Duration interface {
	Observe(handle func(time.Time, time.Time, time.Duration))
}

type duration struct {
	start time.Time
	end   time.Time
}

// Observe 查看时间详细信息
func (d *duration) Observe(handle func(time.Time, time.Time, time.Duration)) {
	delta := d.end.Sub(d.start)
	handle(d.start, d.end, delta)
}

// TimeLog 记录执行时间
func TimeLog(f func()) Duration {
	d := &duration{
		start: time.Now(),
	}
	f()
	d.end = time.Now()
	return d
}
