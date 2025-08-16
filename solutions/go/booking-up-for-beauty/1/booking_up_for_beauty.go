package booking

import(
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        fmt.Println(date, err)
        return time.Now()
    }
    return t  
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        fmt.Println(err)
    }
    if t.Compare(time.Now()) == -1{
		return true
    }
    return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        fmt.Println(err)
    }
    h, _, _ := t.Clock()
    if h >= 12 && h < 18 {
        return true
    }
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        fmt.Println(err)
    }
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %02d:%02d.", 
        t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())  
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    t := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return t
}
