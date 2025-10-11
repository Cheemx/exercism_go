package armstrong

import "math"

func IsNumber(n int) bool {
    num := n
    nums := []int{}
    cnt := 0
    sum := 0
    for n > 0 {
        digi := n % 10
        nums = append(nums, digi)
        cnt++
        n = n / 10
    }
    for _, num := range nums{
        sum += int(math.Pow(float64(num), float64(cnt)))
    }
    return sum == num
}
