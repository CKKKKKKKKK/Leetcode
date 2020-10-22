package partitionLabels

func partitionLabels(S string) []int {
    endList := make([]int, 26);
    length := len(S);
    result := make([]int, 0);
    //mark end point for each letter
    for i := 0; i < length; i++ {
        endList[S[i] - 'a'] = i;
    }
    //partition
    for start, end := 0, 0; end != length - 1; {
        end = endList[S[start] - 'a'];
        for i := start; i <= end; i++ {
            if endList[S[i] - 'a'] > end {
                end = endList[S[i] - 'a'];
            }
        }
        //report
        result = append(result, end + 1 - start);
        //update
        start = end + 1;
    }
    return result;
}