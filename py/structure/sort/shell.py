# -*- coding:utf-8 -*-
"""
author:
date:
version:
content:
"""
from typing import List


def shell_sort(s: List) -> List:
    length = len(s)
    d = int(length / 2)
    while d >= 1:
        for i in range(d, length):
            if s[i] < s[i - d]:
                temp = s[i]
                j = i - d
                while j >= 0 and s[j] > temp:
                    s[j + d] = s[j]
                    j -= d
                s[j + d] = temp
        d = int(d / 2)
    return s
