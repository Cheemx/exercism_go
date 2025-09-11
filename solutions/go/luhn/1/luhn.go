package luhn

func Valid(id string) bool {
    if len(id) < 2 {
        return false
    }
    number := ""
    for _, char := range id {
        if char == ' '{
            continue
        }
        if char < '0' || char > '9' {
            return false
        }

        number += string(char)
    }

    if len(number) < 2{
        return false
    }

    
    sum := 0
    isEven := false
	for i := len(number) - 1;i >= 0; i--{
        digit := int(number[i] - '0')

        if isEven {
            digit *= 2
            if digit > 9{
                digit -= 9
            } 
        }

        sum += digit
        isEven = !isEven
    }
    return sum % 10 == 0
}
