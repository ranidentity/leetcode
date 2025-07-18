1. 如何理解oop
- encapsulations
 - properties, functions
- abstraction
 - hide complex, only show public easy function 
- inheritance
    - parent functions 
- polymorphism
2. singleton 模式design 是什么
3. Thread & process 解释一下， 优缺点，为什么是这样
4. 解释一下死锁是什么 dead lock
5. 解释一下前段 IP 到domain 的过程
6. networking中的特点解释一下
7. 解释你理解的 golang vs java
8. 解释一下 golang： make side 的区别
9. golang 如何define var =/ := 区别
10.  数据结构和算法： 
    1. 给你一个string， 里面是0-9 的数值 但是是string 然后包含运算符 +-：， 写出 +- 法
    2. 问如果不止有+-， * / 怎么处理


Problem solving + coding in Golang – next round
o You'll face algorithmic and real-world coding problems in Go.
o Expect a strong focus on problem-solving under constraints, debugging,
and optimizing code.
Topics to Prepare:
o Arrays, strings, maps, recursion, sorting
o Concurrency patterns in Go
o Time/space complexity analysis

Tip:
Brush up on –
• Basic syntax and working of goroutines
• Usage of mutex
• How and where to close the channels and about waitgroup

1. basics of Goroutine
- channel
    - worker pool
        - fan in/out
        - pipeline
        - worker pool
        - pub sub
        - Request/Response Pattern
        - Timeout Pattern
    - worker pool with context
    - channel merging
    - channel goroutine closure
- mutex
- waitgroup
2. One of the question was increasing memory space in app/website. How do you continuously increase the memory?

3. Related to given 2 log file of timestamp, customer ID and page ID
- Solve to get a list of loyal customer who visit 2 days and 2 unique page id
Advance question if you can solve the first 1: 
- What other approach can it be solve if there is a need to scale for more than 2 days