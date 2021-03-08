## 剑指 Offer 46. 把数字翻译成字符串

https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
 

示例 1:
```
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
```

提示：

> 0 <= num < 231

-------

有点类似爬楼梯的题目，推导方程：

记数字 num 第 i 位数字为 X(i)

f(n) = f(n-1) + f(n-2) (假如X(i-1)X(i) >= 10 && X(i-1)X(i) <= 25)
f(n) = f(n-1) (其他情况)

```go
func translateNum(num int) int {
    // 小于10的直接返回
    if num < 10 {
        return 1
    }
    
    // 数字转字符串
    str := strconv.Itoa(num)
    // 初始化返回结果为1， f2 没有数字为0； f1 第一位数，为1
    f2, f1, fn := 0, 1, 1
    // 从第二位数遍历
    for i:=1; i < len(str); i++ {
        // 先更新f0, f1
        f2, f1 = f1, fn
        // 得到X(i-1)X(i)
        v := str[i-1:i+1]
        if (v <= "25" && v >= "10") {
            fn += f2
        }
    }
    return fn
}
```
