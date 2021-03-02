package rsArray

import (
	"errors"
	"math"
)

type RSArray struct {
	b indexArr
	k int
	n int
}

func NewRSArray() *RSArray {
	arr := &RSArray{}
	arr.b = indexArr{arr: make([]*[]int, 1), k: 0}
	arr.k = 0
	arr.n = 0
	return arr
}
func (this *RSArray) Read(i int) (int, error) {
	if i > this.n {
		return 0, errors.New("index out of bound")
	}
	hi := int(math.Ceil(math.Sqrt((9+8*float64(i))-3) / 2))
	lo := i - hi*(hi+1)/2
	return (*this.b.arr[hi])[lo], nil
}
func (this *RSArray) Write(i int, val int) error {
	if i > this.n {
		return errors.New("index out of bound")
	}
	hi := int(math.Ceil((math.Sqrt(9+8*float64(i)) - 3) / 2))
	lo := i - hi*(hi+1)/2
	(*this.b.arr[hi])[lo] = val
	return nil
}
func (this *RSArray) Grow() {
	if this.n == this.k*(this.k+1)/2 {
		ak := make([]int, this.k+1)
		this.b.grow()
		this.b.arr[this.k] = &ak
		this.k += 1
	}
	this.n += 1
}

func (this *RSArray) Shrink() {
	if this.n >= 2 {
		if this.n-1 == (this.k-1)*this.k/2 {
			this.b.shrink()
			this.k -= 1
		}
		this.n -= 1
	}
}

type indexArr struct {
	arr []*[]int
	k   int
}

func (this *indexArr) grow() {
	if this.k >= len(this.arr) {
		newArr := make([]*[]int, 2*len(this.arr))
		for i := 0; i < this.k; i++ {
			newArr[i] = this.arr[i]
		}
		this.arr = newArr
	}
	this.k += 1
}

func (this *indexArr) shrink() {
	if this.k > 0 {
		this.k -= 1
		this.arr[this.k] = nil
		if this.k < len(this.arr)/3 {
			newArr := make([]*[]int, len(this.arr)/2)
			for i := 0; i < this.k; i++ {
				newArr[i] = this.arr[i]
			}
			this.arr = newArr
		}
	}
}
