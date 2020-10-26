// Fast sort
package smallerNumbersThanCurrent
import (
	"sort"
)
type pair struct{v, pos int}
func smallerNumbersThanCurrent(nums []int) []int {
    length := len(nums)
    data := make([]pair, length)
    result := make([]int, length)
    prev := -1
    for i, v := range nums {
        data[i] = pair{v, i}
    }
    sort.Slice(data, func(i, j int) bool {
        return data[i].v < data[j].v
    })
    for i := 0; i < length; i++ {
        if prev == -1 || data[i].v > data[i - 1].v {
            prev = i
        }
        result[data[i].pos] = prev 
    }
    return result
}