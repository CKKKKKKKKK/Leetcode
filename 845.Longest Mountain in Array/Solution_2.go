package longestMountain
func longestMountain(A []int) int {
    length := len(A)
    if length < 3 {
        return 0
    }
    result := 0
    left := make([]int, length)
    right := make([]int, length)
    left[0] = 0
    right[length - 1] = 0
    for i := 1; i < length; i++ {
        if A[i] > A[i - 1] {
            left[i] = left[i - 1] + 1
        }
    }
    for i := length - 2; i >= 0; i-- {
        if A[i] > A[i + 1] {
            right[i] = right[i + 1] + 1
        }
    }
    for i := 0; i < length; i++ {
        if left[i] > 0 && right[i] > 0 && left[i] + right[i] + 1 > result {
            result = left[i] + right[i] + 1
        }
    }
    return result
}