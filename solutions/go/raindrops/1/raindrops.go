package raindrops

import "fmt"

func Convert(number int) string {
	ans := ""
    if number % 3 == 0 {
        ans += "Pling"
    } 
    if number % 5 == 0 {
        ans += "Plang"
    }
    if number % 7 == 0 {
        ans += "Plong"
    }
    if number % 3 != 0 && number % 5 != 0 && number % 7 != 0 {
        return fmt.Sprintf("%d", number)
    }
    return ans
}
