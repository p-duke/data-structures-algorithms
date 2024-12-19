func minMeetingRooms(intervals [][]int) int {
	rooms := 0
	maxRooms := 0
	start := make([]int, 0)
	end := make([]int, 0)

	for _, v := range intervals {
		start = append(start, v[0])
		end = append(end, v[1])
	}

	sort.Slice(start, func(i int, j int) bool {
		return start[i] < start[j]
	})

	sort.Slice(end, func(i int, j int) bool {
		return end[i] < end[j]
	})

    i, j := 0, 0
	for i < len(start) && j < len(end) {
		if start[i] < end[j] {
			rooms++
            if i < len(start) {
                i++
            }
		} else {
			rooms--
            j++
		}

		if rooms > maxRooms {
			maxRooms = rooms
		}
	}

	return maxRooms
}
