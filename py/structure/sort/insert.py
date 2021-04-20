# -*- coding:utf-8 -*-
"""
author:
date:
version:
content:
"""
from typing import List


def insert_sort(s: List) -> List:
    for i in range(1, len(s)):
        v = s[i]
        if v < s[i - 1]:
            j = i - 1
            while j >= 0 and s[j] > v:
                s[j + 1] = s[j]
                j -= 1
            s[j + 1] = v
    return s


def insert_sort_desc(s: List) -> List:
    for i in range(1, len(s)):
        v = s[i]
        if v > s[i - 1]:
            j = i - 1
            while j >= 0 and s[j] < v:
                s[j + 1] = s[j]
                j -= 1
            s[j + 1] = v
    return s
