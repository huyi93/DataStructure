"""
 堆排序
"""


# 正序
def heap_sort(s: list) -> list:
    length = len(s)
    num = length - 1
    while num >= 0:
        build_max_heap(s, length)
        s[0], s[num] = s[num], s[0]
        num -= 1
    return s


def build_max_heap(s: list, length: int):
    i = int(length / 2)
    while i >= 0:
        j = 2 * i
        temp = s[i]
        while j < length:
            if j <= length - 1 and s[j] < s[j + 1]:
                j += 1
            if s[i] >= s[j]:
                break
            else:
                s[i] = s[j]
                i = j
            j *= 2
        s[i] = temp
        i -= 1
