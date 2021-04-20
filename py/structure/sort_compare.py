# -*- coding:utf-8 -*-
"""
author:
date:
version:
content:
"""
from sort import bubble, insert, shell, quick, select, heap
from tools import generate_list

if __name__ == '__main__':
    source = generate_list(14, True)
    print(bubble.bubble_sort(source.copy()))
    print(insert.insert_sort(source.copy()))
    print(shell.shell_sort(source.copy()))
    s = source.copy()
    quick.quick_sort(s, 0, len(s) - 1)
    print(s)
    print(select.simple_select_sort(source.copy()))
    print(heap.heap_sort(source.copy()))
    print(source)
