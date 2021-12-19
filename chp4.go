package main
import "fmt"

func main(){
  multiplesOfThree(103)
  fizzBuzzPrinter(100)
}

/**
1. What does the following program print?
i := 10
if i > 10 {
 fmt.Println("Big")
} else {
 fmt.Println("Small")
}
It prints "Small".

2. Write a program that prints out all the numbers between 1 and 100 that are
evenly divisible by 3 (i.e., 3, 6, 9, etc.).
**/
func multiplesOfThree(end int){
  fmt.Printf("The multiples of 3 from 0 to %d are:\n", end )
  for i := 1; i <= end; i++ {
    if i % 3 == 0 {
      if end - i <= 3 {
        fmt.Printf("%d\n", i)
      }else{
        fmt.Printf("%d, ", i)
      }
    }
  }
}

/**
3. Write a program that prints the numbers from 1 to 100, but for multiples of
three, print “Fizz” instead of the number, and for the multiples of five, print
“Buzz.” For numbers that are multiples of both three and five, print “FizzBuzz.”
**/
func fizzBuzzPrinter(end int){
  fmt.Printf("Printing fizzbuz for numbers 1 to %d:\n", end )
  for i := 1; i <= end; i++ {
    res := ""
    if i % 3 == 0 { res += "Fizz" }
    if i % 5 == 0 { res += "Buzz" }
    if len(res) > 0{
      if i == end{
        fmt.Printf("%s\n", res)
      }else{
        fmt.Printf("%s, ", res)
      }
    }else{
      if i == end{
        fmt.Printf("%d\n", i)
      }else{
        fmt.Printf("%d, ", i)
      }
    }
  }
}
