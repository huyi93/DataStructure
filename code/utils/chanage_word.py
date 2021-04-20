word = """
执行用时：
4 ms
, 在所有 Go 提交中击败了
89.77%
的用户
内存消耗：
4.2 MB
, 在所有 Go 提交中击败了
95.34%
的用户
"""

types = """
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 """


def change():
    lines = word.replace("\n", " ").replace("内存消耗", "\n内存消耗")
    return lines[1:]


def remove():
    lines = types.replace("*", "").replace("/", "")
    print(lines)


if __name__ == '__main__':
    print(change())
    remove()
