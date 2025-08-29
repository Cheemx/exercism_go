package expenses

import (
    "errors"
    "fmt"
)
// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var res []Record
    for _, rec := range in {
        if predicate(rec) {
            res = append(res, rec)
        }
    }
    return res
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
        return r.Day >= p.From && r.Day <= p.To
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func (r Record) bool {
        return c == r.Category
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    expense := 0.0
	for _, rec := range in {
        predicate := ByDaysPeriod(p)
        if predicate(rec) {
            expense += float64(rec.Amount)
        }
    }
    return expense
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var records []Record
    for _, rec := range in {
        predicate := ByCategory(c)
        if predicate(rec) {
            records = append(records, rec)
        }
    }
    if len(records) < 1 {
        return 0.0, errors.New(fmt.Sprintf("unknown category %s", c)) 
    }
    return TotalByPeriod(records, p), nil
}
