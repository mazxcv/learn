package clock

import "fmt"

// Define the Clock type here.

type Clock struct {
	Hour   int
	Minute int
}

func New(h, m int) Clock {
	for m < 0 {
		m += 60
		h--
	}

	for h < 0 {
		h += 24
	}

	minute := m % 60
	hour := (h + m/60) % 24

	return Clock{
		Hour:   hour,
		Minute: minute,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hour, c.Minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hour, c.Minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
