package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
    hour int
    min int
}

func normalize(m int) Clock {
    m = ((m % 1440) + 1440) % 1440 
    hour := m / 60
    min := m % 60
    return Clock{
        hour: hour,
        min:  min,
    }
}

func New(h, m int) Clock {
    return normalize(h*60+m)
}

func (c Clock) Add(m int) Clock {
	return normalize(c.hour*60+c.min+m)
}

func (c Clock) Subtract(m int) Clock {
	return normalize(c.hour*60+c.min-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}
