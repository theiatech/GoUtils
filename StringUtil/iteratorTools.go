package StringUtil

import (
	"fmt"
	"math"
)

/**
* Example: IteratorTools StringUtil.IteratorTools{'a','b','c','d'}
*/
type IteratorTools []int8

type Iterator struct {
	len		int
	data	IteratorTools
	idx		float64
	count	float64
	start	float64
	stop	float64
}

/**
 * @Param length
 * @Param startIdx  0 ~ (stopIdx - 1)
 * @Param stopIdx   startIdx +1 ~ math.Pow(float64(len(IteratorTools)), float64(length))
 **/
func (c IteratorTools) Product(length int,startIdx ,stopIdx float64) *Iterator {
	end := math.Pow(float64(len(c)), float64(length))
	if 0 < stopIdx && end >= stopIdx {
		end = stopIdx
	}

	begin := startIdx
	if 0 > begin || end <= begin {
		begin = 0
	}
	return &Iterator{
		data: c,
		idx: begin,
		count: math.Pow(float64(len(c)), float64(length)),
		start: begin,
		stop: end,
		len: length,
	}
}

/**
 * @Param idx < stopIdx
 **/
func (iter *Iterator)SetIndex(idx int) bool {
	i := math.Floor(math.Abs(float64(idx)))
	if i < iter.stop {
		iter.idx = i
	}
	return false
}

func (iter *Iterator)HasNext() bool {
	return iter.idx < iter.stop
}

func (iter *Iterator)Next() string {
	res := ""
	if iter.idx < iter.stop {
		for i := 0; i < iter.len; i++ {
			idx := int(iter.idx)%int(math.Pow(float64(len(iter.data)),float64(i + 1)))/int(math.Pow(float64(len(iter.data)),float64(i)))
			res = fmt.Sprintf("%c%s",iter.data[idx],res)
		}
		iter.idx += 1
	}
	return res
}

func (iter *Iterator)Count() float64 {
	return iter.stop - iter.start
}
