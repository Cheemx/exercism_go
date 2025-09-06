package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 1000000000 seconds to the time and returns it
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
