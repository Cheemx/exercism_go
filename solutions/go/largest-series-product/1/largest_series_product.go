package lsproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if span > len(digits) || span < 0 {
        return 0, errors.New("something wrong")
    }
    var maxProd int64 = 0
	for i := 0;i <= len(digits)-span; i++ {
        var prod int64 = 1
        for j := i;j < i + span; j++ {
            if digits[j] < '0' || digits[j] > '9' {
    			return 0, errors.New("digits allowed only")
			}
            prod = prod * int64(digits[j] - '0')
        }
        if prod > maxProd {
            maxProd = prod
        }
    }
    return maxProd, nil
}
