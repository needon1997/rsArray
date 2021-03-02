package rsArray_test

import "testing"
import "rsArray"
func TestNew(t *testing.T) {
	_ =rsArray.NewRSArray()
}
func TestRead(t *testing.T) {
	arr :=rsArray.NewRSArray()
	arr.Grow()
	arr.Write(0,1)
	arr.Grow()
	arr.Write(1,1)
	arr.Grow()
	arr.Write(2,1)
	arr.Grow()
	arr.Write(3,1)
	arr.Grow()
	arr.Write(4,1)
	arr.Shrink()
	arr.Shrink()
	arr.Shrink()
	arr.Shrink()
	arr.Shrink()
}