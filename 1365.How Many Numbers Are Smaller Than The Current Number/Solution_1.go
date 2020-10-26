package smallerNumbersThanCurrent
//Counting sort
func smallerNumbersThanCurrent(nums []int) []int {
    length := len(nums)
    cnt := make([]int, 101)
    result := make([]int, length)
    sum := 0
    for i := range nums {
        cnt[nums[i]]++
    }
    for i := range result {
        for j := 0; j < nums[i]; j++ {
            sum += cnt[j]
        }
        result[i] = sum
        sum = 0
    }
    return result
}