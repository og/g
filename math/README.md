# math


## Float64ToFixed

`Float64ToFixed(f float64, digit int) float64`

按指定位数截取 float64 小数点。（四舍**六**入五成双）

```go
v := float64(0.123456789123456789123456789)
Float64ToFixed(v, 0) 0
Float64ToFixed(v, 1) 0.1
Float64ToFixed(v, 2) 0.12
Float64ToFixed(v, 3) 0.123
Float64ToFixed(v, 4) 0.1235	// round off 5
Float64ToFixed(v, 5) 0.12346 // round off 6
Float64ToFixed(v, 6) 0.123457 // round off 7
Float64ToFixed(v, 7) 0.1234568 // round off 8
Float64ToFixed(v, 8) 0.12345679 // round off 9
Float64ToFixed(v, 9) 0.123456789
Float64ToFixed(v, 10) 0.1234567891
```


## IntPercent

求2个整数的百分比，返回 0 ~ 无限大

函数内部会避免除 0 的情况，如果  `IntPercent(14,0)` 将返回 0

计算百分比时候会做舍入操作，采取的是四舍**六**入五成双

```go
IntPercent(14,0)// 0
IntPercent(4,10)// 40
IntPercent(333,1000)// 33
```

## Float64Percent

实现类型 `IntPercent`，区别是函数参数接收的是float64

`Float64Percent (part float64, total float64) int`


```go
IntPercent(0.14,0)// 0
IntPercent(0.4,1)// 40
IntPercent(0.333,1)// 33
```