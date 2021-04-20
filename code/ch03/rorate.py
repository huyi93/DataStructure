"""
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

进阶：
尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
"""
from typing import List


class Solution:
    """
    执行用时： 76 ms , 在所有 Python3 提交中击败了 19.37% 的用户
    内存消耗： 21 MB , 在所有 Python3 提交中击败了 5.00% 的用户
    """

    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        k %= len(nums)
        self.reverse(nums, 0, len(nums) - 1)
        self.reverse(nums, 0, k - 1)
        self.reverse(nums, k, len(nums) - 1)

    def reverse(self, nums, i, j):
        while i <= j:
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
            j -= 1


class MySolution:
    """
    执行用时： 48 ms , 在所有 Python3 提交中击败了 39.63% 的用户
    内存消耗： 21.4 MB , 在所有 Python3 提交中击败了 5.07% 的用户
    """

    def rotate_1(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        [1,2,3,4,5,6,7], k = 3
        """
        k %= len(nums)
        temp = nums[:len(nums) - k]
        nums[:len(nums) - k] = nums[len(nums) - k:len(nums)]
        nums[len(nums) - k:len(nums)] = temp

    def rotate_2(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        [1,2,3,4,5,6,7], k = 3
        """
        k %= len(nums)
        nums[:] = nums[-k:] + nums[:-k]


if __name__ == '__main__':
    ls = [1, 2, 3, 4, 5, 6, 7]
    MySolution().rotate_2(ls, 3)
    print(ls)
