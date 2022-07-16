# 蓝山后端Go语言第一次作业

## 1.[解方程](https://www.luogu.com.cn/problem/P1022)

因为是一元一次方程，所以最终一定可以化成
$$
x=-\frac{b}{k}
$$
 的形式，为了减少变量数方便，将等号右边的多项式直接移到等式左边，即系数均乘 -1，用变量now标记一下。



数字读入:   x = x*10+ c - '0'  c为ASCII码，所以需要减去‘0’

未知数处理： 若遇到小写字母，则判定为未知数，将记录的系数累加至未知数系数上，同时清零系数。

遇到 +0x 和 -0x 的情况，则用r进行系数是否读入的记录，若未曾读过忽略跳过即可

遇到 -x = 0 情况， 会出现-0.000的情况，特判下改成0.000

## 2.[ The Tamworth Two ](https://www.luogu.com.cn/problem/P1518)

这题主要采用模拟运动的方法，不断按照规则在地图内运动，两人的坐标在同一个地方则表示相遇。

地图可以映射到一个二维int数组，空地可以映射为0，障碍物映射为1，C映射为2，Farmer映射为3.

因为牛和farmer的运动方式是一样的，每个人的状态可以由一个struct来描述，如下：

```go
type status struct {
   direction int
   x         int
   y         int
}
```

其中direction是东南西北的映射，采取如下的map映射

```go
const (
   NORTH = iota
   EAST
   SOUTH
   WEST
)
```

这样做的好处之一就是，当遇到障碍物或地图边界的时候，其direction可以利用数学方法，快速变换

```go
character.direction = (character.direction + 1) % 4
```

此外还需要考虑到不能相遇的情况，即两人陷入循环，且在第一次循环的时候就无法碰到，自然下次循环也无法碰到。这种情况下采取了一个特征值计算其两人的轨迹，boolean数组来记录，如果特征值相同，表示之前出现过这个情况，已经开始了循环。

特征值计算方法：

```go
f := John.x + John.y*10 + cow.x*100 + cow.y*1000 + cow.direction*10000 + John.direction*40000
```

## 3.[凯撒密码 ](https://www.luogu.com.cn/problem/P1914)

挺简单的，注意越界别超出'z'就行，超出了要重新从‘a’开始，用取余一句代码即可

```go
97+(b-97+offset)%26
```

97为‘a’的ASCII码，b为当前字母，offset表示凯撒密码的向后位移数量