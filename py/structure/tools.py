# -*- coding:utf-8 -*-
"""
author:
date:
version:
content:
"""
from random import randint
from typing import List


# 生成随机数列表 0 ~ 999
def generate_list(length: int, is_different: bool = False) -> List:
    simple_list = []
    while len(simple_list) < length:
        num = randint(0, 999)
        if is_different:
            if num in simple_list:
                continue
        simple_list.append(num)
    return simple_list


if __name__ == '__main__':
    pass
