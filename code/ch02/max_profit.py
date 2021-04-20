"""
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
    随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3


"""
from typing import List


# 执行用时： 52 ms , 在所有 Python3 提交中击败了 42.74% 的用户
# 内存消耗： 15.4 MB , 在所有 Python3 提交中击败了 97.72% 的用户
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        ans = 0
        n = len(prices)
        for i in range(1, n):
            ans += max(0, prices[i] - prices[i - 1])
        return ans


# 执行用时： 40 ms , 在所有 Python3 提交中击败了 88.84% 的用户
# 内存消耗： 15.6 MB , 在所有 Python3 提交中击败了 49.37% 的用户
class MySolution:
    def maxProfit(self, prices: List[int]) -> int:
        max_num = prices[0]
        min_num = prices[0]
        now = 0
        max_profit = 0
        length = len(prices) - 1
        for i, v in enumerate(prices):
            if i < length:
                if v < prices[i + 1]:
                    if min_num > v:
                        min_num = v
                    if prices[i + 1] > max_num:
                        max_num = prices[i + 1]
                    now = max_num - min_num
                else:
                    if now != 0:
                        max_profit += now
                        now = 0
                    max_num = min_num = prices[i + 1]
        if now != 0:
            max_profit += now
        return max_profit


if __name__ == '__main__':
    print(Solution().maxProfit([7, 1, 5, 3, 6, 4]))
