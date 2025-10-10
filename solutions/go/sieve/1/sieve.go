package sieve

func Sieve(limit int) []int {
    if limit < 2 {
        return []int{}
    }
    mep := map[int]bool{}
    for i := 2; i <= limit; i++ {
        if _, ok := mep[i]; !ok {
        	j := i + 1
            for j <= limit {
                if j % i == 0 {
                    mep[j] = true
                } 
                j++
            }
        }
    }
    ans := []int{}
    for i := 2;i <= limit;i++ {
        if _, ok := mep[i]; !ok {
        	ans = append(ans, i)
        }
    }
    return ans
}
