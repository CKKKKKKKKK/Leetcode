package longestMountain
func longestMountain(A []int) int {
    var prev int
    asc := false
    dec := false
    count := 0
    max := 0
    length := len(A)
    if length < 3 {
        return 0
    }
    for i := 0; i < length; i++ {
        if i == 0 {
            prev = A[i]
            continue
        }
        if !asc && !dec && A[i] > prev {
            prev = A[i]
            asc = true
            count = count + 2
            continue
        }else if asc && !dec && A[i] > prev {
            prev = A[i]
            count++
            continue
        }else if asc && !dec && A[i] < prev {
            prev = A[i]
            dec = true
            count++
            continue
        }else if asc && dec && A[i] < prev {
            prev = A[i]
            count++
            continue
        }else if asc && dec && A[i] > prev {
            prev = A[i]
            if count > max {
                max = count
            }
            count = 2
            dec = false
            continue
        }else {
            if asc && dec && count > max {
                max = count
            }
            prev = A[i]
            asc = false
            dec = false
            count = 0
            continue
        }
    }
    if asc && dec && count > max {
        max = count
    }
    return max
}