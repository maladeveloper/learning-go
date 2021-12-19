package main

import "fmt"

func main(){
  fmt.Println("Hello World")

  fmt.Println("Individual characters are represented as bytes-")
  fmt.Println("Accessing the first character of 'Hello World' gives:","Hello World"[0])

  var str string
  fmt.Printf("An initialised string that has not been defined defaults to '%s'\n",str)

  fmt.Println("Go can infer types but what happens if we let it infer and then assign a different type value 🤔")
  x := "hello"
  fmt.Printf("Initialised x as '%s'", x)
  fmt.Println("Now lets assign x as an integer - go throws a compiler error !\n\n")
  //x = 5

  fmt.Println("Dividing two integers that would result in floating point i.e 5/9")
  val := 5/9
  fmt.Printf("... equals %d\n", val)
  fmt.Println("This is 2 integers are being used so the value of the expression must be integer.")
  fmt.Println("To prevent this from happen and allow the value to be float we have to have atleast one of the numbers be a float.")
  newVal := 5.0/9
  fmt.Printf("Now the value equals %f\n", newVal)

  fmt.Printf("When float and integer interact, the result is always a float i.e 6.0/3 equals %f\n", 6.0/3)

  fmt.Println("Using make to create new slice specifying the length and capacity give:")
  a := make([]int, 5, 10)
  printArrInfo(a)

  fmt.Println("\nUsing make to create new slice specifying only the length of 5 gives:")
  b := make([]int, 5)
  printArrInfo(b)

  fmt.Println("\nWithout using make to create a new slice with 5 elements give:")
  c := []int{ 1,2, 3, 4, 5, }
  printArrInfo(c)

  fmt.Println("\nWithout using make to create a new slice with no elements give:")
  d := []int{}
  printArrInfo(d)
}

func printArrInfo( a []int){
  fmt.Printf("Length: %d\n", len(a))
  fmt.Printf("Capacity: %d\n", cap(a))
  fmt.Println("The Array:", a)
}
