# -*- coding:utf-8 -*-
"""
author:
date:
version:
content:
"""
from typing import List


# 冒泡排序
def bubble_sort(s: List) -> List:
    length = len(s)
    for i in range(length):
        for j in range(i + 1, length):
            if s[i] > s[j]:
                s[i], s[j] = s[j], s[i]
    return s
