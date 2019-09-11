# go-map

## Keys

用于解决迭代顺序问题

> https://blog.golang.org/go-maps-in-action
> 当使用范围循环在映射上迭代时，没有指定迭代顺序，也不能保证从一个迭代到下一个迭代的顺序是相同的。

使用 `Keys(sMap).String()` 可以获取稳定的 `[]string`，确保每次遍历获取的 key 的顺序一致。

```go
sMap := map[string]string{
    "name": "nimo",
    "title": "abc",
    "lang": "go",
}

for _, key := range Keys(sMap).String() {
    log.Print(key+":"+sMap[key])
}
// lang:go
// name:nimo
// title:abc
```


使用 `Keys(iMap).Int()` 可以获取稳定的 `[]int`，确保每次遍历获取的 key 的顺序一致。

```go
iMap := map[int]string{
    6: "nimo",
    2: "abc",
    9: "go",
}
for _, key := range Keys(iMap).Int() {
    log.Print(key, ":", iMap[key])
}
// 2:abc
// 6:nimo
// 9:go
```

使用 `Keys(fMap).Float64()` 可以获取稳定的 `[]float64`，确保每次遍历获取的 key 的顺序一致。

```go
fMap := map[float64]string{
    6.1: "nimo",
    2.2: "abc",
    9.3: "go",
}
for _, key := range Keys(fMap).Float64() {
    log.Print(key, ":", fMap[key])
}
// 2.2:abc
// 6.1:nimo
// 9.3:go
```
