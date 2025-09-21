package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    mep := map[string]int{}
    for k, v := range in {
        for _, str := range v {
            mep[strings.ToLower(str)] = k
        }
    }
    return mep
}
