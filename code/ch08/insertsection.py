import collections
from typing import List


class Solution:

    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        """
        执行用时： 44 ms , 在所有 Python3 提交中击败了 89.55% 的用户
        内存消耗： 14.7 MB , 在所有 Python3 提交中击败了 95.85% 的用户
        """
        num1 = collections.Counter(nums1)
        num2 = collections.Counter(nums2)
        num = num1 & num2
        return list(num.elements())


if __name__ == '__main__':
    print(Solution().intersect(nums1=[1, 2, 2, 1], nums2=[2, 2]))
