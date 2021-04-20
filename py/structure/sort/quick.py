# -*- coding:utf-8 -*-
"""
author:
date:
version:
content:
"""
from typing import List


# 快速排序
def quick_sort(ls: List, low: int, high: int) -> None:
    """
    """
    if low < high:
        pivot = partition(ls, low, high)
        quick_sort(ls, low, pivot - 1)
        quick_sort(ls, pivot + 1, high)


def partition(ls: List, low: int, high: int) -> int:
    pivot = ls[low]
    while low < high:
        while low < high and pivot <= ls[high]:
            high -= 1
        ls[low] = ls[high]

        while low < high and pivot >= ls[low]:
            low += 1
        ls[high] = ls[low]

    ls[low] = pivot

    return low


if __name__ == '__main__':
    s = [1, 3, 6, 7, 8, 9, 2, 5, 1]
    quick_sort(s, 0, len(s) - 1)
    print(s)
