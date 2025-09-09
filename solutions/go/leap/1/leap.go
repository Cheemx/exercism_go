
package leap

// IsLeapYear checks if int year it accepts is leap year or not
func IsLeapYear(year int) bool {
	if year % 100 == 0 && year % 400 != 0{
        return false
    } else if year % 4 == 0 {
         
            return true
        
    }
    return false
}
