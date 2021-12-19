package hello
/**
1. Writing a good suite of tests is not always easy, but the process of writing tests
often reveals more about a problem than you may at first realize. For example,
with our Average function, what happens if you pass in an empty list
([]float64{})? How could the function be modified to return 0 in this case?
An early return could be added that exits early with 0 if it identifies an empty list.

2. Write a series of tests for the Max function.
**/
func Max( arr []int) int{
  max := arr[0]
  for _, val := range arr{
    if val > max{
      max = val
    }
  }
  return max
}

