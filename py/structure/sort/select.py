# -*- coding:utf-8 -*-
"""
author:
date:
version:
content:
"""


# 简单选择排序
def simple_select_sort(s: list) -> list:
    for k in range(len(s) - 1):
        min_key = select_min_key(s[k + 1:]) + k + 1
        if s[k] > s[min_key]:
            s[k], s[min_key] = s[min_key], s[k]
    return s


def select_min_key(s: list) -> int:
    min_key, min_value = 0, s[0]
    for k, v in enumerate(s):
        if v < min_value:
            min_key = k
            min_value = v
    return min_key
