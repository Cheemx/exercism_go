package phonenumber

import (
    "strings"
    "unicode"
    "errors"
    "fmt"
)

func Number(phoneNumber string) (string, error) {
    if strings.HasPrefix(phoneNumber, "+") {
        if string(phoneNumber[1]) != "1"{
            return "", errors.New("wrong formatting")
        }
        phoneNumber = phoneNumber[3:]
    } else if strings.HasPrefix(phoneNumber, "1") {
        if string(phoneNumber[1]) == " " {
    		phoneNumber = phoneNumber[2:]
        } else {
            phoneNumber = phoneNumber[1:]
        }
    }
    number := ""
    cnt := 0
    for _, c := range phoneNumber {
        if unicode.IsDigit(c) {
            if cnt == 0 || cnt == 3{
                if c - '0' < 2{
                    return "", errors.New("N should've been greater than 2")
                }
            }
            number += string(c)
            cnt++
        }
    }
    if cnt != 10{
        return "", errors.New("wrong number of digits")
    }
    return number, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return phoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
    phoneNumber, err :=  Number(phoneNumber)
    if err != nil {
        return "", err
    }
	return fmt.Sprintf("(%s) %s-%s", phoneNumber[:3], phoneNumber[3:6], phoneNumber[6:]), nil
}
