package hello
import "testing"

func TestMax( t *testing.T){
  arr := []int{1, 123, 2, 13, 45}
  max := Max(arr)

  if max != 123 {
    t.Error("Expected 123 but got ", max)
  }
}
