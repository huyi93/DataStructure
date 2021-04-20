from typing import List


class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        temp = {}
        for v in nums:
            if v in temp:
                return True
            temp[v] = 1

        return False

    def containsDuplicate2(self, nums: List[int]) -> bool:
        """
        执行用时： 48 ms , 在所有 Python3 提交中击败了 56.14% 的用户
        内存消耗： 20.2 MB , 在所有 Python3 提交中击败了 44.43% 的用户
        """
        return True if len(set(nums)) != len(nums) else False


if __name__ == '__main__':
    ls = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
    print(Solution().containsDuplicate(ls))
