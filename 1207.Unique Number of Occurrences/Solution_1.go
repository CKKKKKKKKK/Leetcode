package uniqueOccurrences

// Double hash table
func uniqueOccurrences(arr []int) bool {
	cnt := map[int]int{}
	times := map[int]bool{}
	for _, v := range arr {
		cnt[v]++
	}
	for _, v := range cnt {
		times[v] = true
	}
	return len(cnt) == len(times)
}
