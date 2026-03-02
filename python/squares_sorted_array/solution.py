class Solution:
    def sortedSquares(self, nums: list[int]) -> list[int]:
        # Problem
        # - Given int array of nums
        # - sorted in ascending
        # - return array of the squares of each number asc

        # Solution
        # - use two pointer
        # - square both and take the biggest
        # - allocate new array for result
        # - take largest and prepend to array
        # - return result
        result = []
        r = len(nums) - 1
        l = 0
        while l <= r:
            l_sq = nums[l] * nums[l]
            r_sq = nums[r] * nums[r]
            if l_sq > r_sq:
                result.insert(0, l_sq)
                l += 1
            else:
                result.insert(0, r_sq)
                r -= 1
        return result

if __name__ == "__main__":
    solution = Solution()
    result = solution.sortedSquares([-4, -1, 0, 3, 10])
    print(result)
