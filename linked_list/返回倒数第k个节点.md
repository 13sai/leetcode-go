# 面试题 02.02. 返回倒数第 k 个节点

https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/

实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动
```
示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。
```

---

1. 计算长度
2. 遍历到倒数第k个节点

```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
    var i int
    tmp := head
    for tmp.Next != nil {
        tmp = tmp.Next
        i++
    }
    i++


    tmp = head
    for j:=0; j < (i-k); j++ {
        tmp = tmp.Next
    }

    return tmp.Val
}
```

使用链表翻转显然也是一种解法。


题解看到这么一种解法，很不错。算是双指针吧，fast先移动k个偏移量，再slow和fast一起向后，fast到末尾，slow也刚好到达题目中所需的倒数第k个节点。

```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
    fast := head
    slow := head
    for k > 0 {
        fast = fast.Next
        k--
    }
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    return slow.Val
}
```