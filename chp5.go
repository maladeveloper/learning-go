package main
import "fmt"

func main(){
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }
  findMinInArr(x)
}

/**
1. How do you access the fourth element of an array or slice?
Suffixing the array or slice with '[3]'.

2. What is the length of a slice created using make([]int, 3, 9)?
The length is 3 but it has the capacity to be 9.

3. Given the following array, what would x[2:5] give you?
x := [6]string{"a","b","c","d","e","f"}
It would give = ["c","d","e"]

4. Write a program that finds the smallest number in this list:
x := []int{
 48,96,86,68,
 57,82,63,70,
 37,34,83,27,
 19,97, 9,17,
}
**/
func findMinInArr(arr []int){
  fmt.Println("Given this array:", arr)
  min := arr[0]
  for _, val := range arr {
    if val < min{
      min = val
    }
  }
  fmt.Println("The smallest value is:", min)
}
