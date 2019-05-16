# Golang标准库sort
> The GO Programming Language读书笔记，7.6节
> 改进的自定排序实现 
> [tempflag code from the book](https://github.com/adonovan/gopl.io/tree/master/ch7/tempflag)

## Sort Tracks 
### 包装待排序的切片
```go 
type By func(t1, t2 *Track) bool

type TrackSorter struct {
	tracks []*Track
	by     By
}
```
随后使用包装类型（或其指针）去实现`sort.Interface`接口

### 包装的Sort
```go 
func (by By) Sort(tracks []*Track) {
	ts := TrackSorter{
		tracks: tracks,
		by:     by,
	}
	sort.Sort(&ts)
}
```