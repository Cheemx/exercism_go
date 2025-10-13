package prime

func Factors(n int64) []int64 {
    ans := []int64{}
    for i := int64(2);i <= n; i++ {
        for n % i == 0 && n > 0 {
            ans = append(ans, i)
            n = n / i
        }
    }
    return ans
}