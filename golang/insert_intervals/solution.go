func insert(intervals [][]int, newInterval []int) [][]int {
    result := make([][]int, 0)
    // before overlap
    for _, v := range intervals {
        if v[1] < newInterval[0] {
            result = append(result, v)
        }
    }

    // merge overlap
    for _, v := range intervals {
        if v[0] <= newInterval[1] && v[1] >= newInterval[0] {
            start := min(v[0], newInterval[0])
            end := max(v[1], newInterval[1])
            newInterval = []int{start, end}
        }
    }

    result = append(result, newInterval)
    
    // after overlap
    for _, v := range intervals {
       if v[0] > newInterval[1] {
            result = append(result, v)
        }
    }

    return result
}
