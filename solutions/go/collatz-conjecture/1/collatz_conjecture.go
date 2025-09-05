package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n < 1 {
        return 0, errors.New("zero or negative number found")
    }
    cnt := 0
	for n > 1 {
        if n % 2 == 0{
            n = n / 2
        } else {
            n = n * 3 + 1
        }
        cnt++
    }
    return cnt, nil
}
