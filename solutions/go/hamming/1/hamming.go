package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("a and b should be of same size")
    }
    cnt := 0
    for i, _ := range a {
        if a[i] != b[i] {
            cnt++
        }
    }
    return cnt, nil
}
