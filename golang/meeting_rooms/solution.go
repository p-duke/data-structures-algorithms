func canAttendMeetings(intervals [][]int) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    for i := 0; i < len(intervals) - 1; i++ {

        j := i + 1
        end := intervals[i][1]
        start := intervals[j][0]

        if end > start {
            return false
        }
    }
    
    return true
}
