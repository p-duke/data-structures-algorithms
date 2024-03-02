// O(n) time | O(n) space
func containsDuplicate(nums []int) bool {
    occur := make(map[int]int)
    for _, v := range nums {
        if _, ok := occur[v]; !ok {
            occur[v] = 1
        }

        if occur[v] > 1 {
            return true
        } else {
            occur[v] += 1
        }
    }

    return false
}
