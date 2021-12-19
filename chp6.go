package main
import (
  "fmt"
  "math/rand"
  "time"
)

func main(){
  sum(4, 5)

  var num int
  var isEven bool

  fmt.Println("Running half num on '5' returns:")
  num, isEven = halfNum(5)
  fmt.Printf("Number is Even: %t\n", isEven)
  fmt.Printf("Half of the number: %d\n", num)


  fmt.Println("Running half num on '16' returns:")
  num, isEven = halfNum(16)
  fmt.Printf("Number is Even: %t\n", isEven)
  fmt.Printf("Half of the number: %d\n", num)

  fmt.Println("\nThe largest number out of this array of numbers - 8, 123, 3, 2 is:")
  fmt.Printf("%d\n",findLargestNumber( 8, 123, 3, 2 ))

  fmt.Println("\nUse closure to initialise an odd number generator")
  oddNumberGen := makeOddGenerator()
  fmt.Printf("Calling odd number generator once: %d\n", oddNumberGen())
  fmt.Printf("Calling odd number generator twice: %d\n", oddNumberGen())
  fmt.Printf("Calling odd number generator third time: %d\n", oddNumberGen())

  fmt.Println("\nThe 20th fibonacci number is:", fib(20))

  fmt.Println("\nTest out the function which causes a panic...")
  testPanic()
  fmt.Println("Success!")

  fmt.Println("\nGet memory address of random integer...")
  intPointer := getRandomIntPointer()
  fmt.Println("Memory Address:", intPointer)
  fmt.Println("Value:", *intPointer)

  fmt.Println("\nAssign an initialised interger to 10 via passing by pointer...")
  x := 123
  fmt.Printf("The current integer values is: %d\n", x)
  assignedToTen(&x)
  fmt.Printf("After the calling 'assignedToTen' the integer has a value: %d\n", x)

  fmt.Println("\nInitialising pointer via 'y := new(int)'")
  y := new(int)
  fmt.Println("Printing y:", y)
  fmt.Println("Initialising pointer via 'var z *int'")
  var z *int
  fmt.Println("Printing z:", z)

  fmt.Println("\nSwapping a and b using pass by pointer function- ")
  a := 10
  b := 32
  fmt.Printf("Before swapping a equals %d and b equals %d\n", a, b)
  swapInt(&a, &b)
  fmt.Printf("After swapping a equals %d and b equals %d\n", a, b)

}

/**
1. sum is a function that takes a slice of numbers and adds them together. What
would its function signature look like in Go?
**/
func sum( num1 int, num2 int) int{
  fmt.Printf("Adding two numbers %d and %d...", num1, num2)
  sum :=  num1 + num2
  fmt.Printf("equals %d\n", sum)
  return sum
}

/**
Write a function that takes an integer and halves it and returns true if it was even
or false if it was odd. For example, half(1) should return (0, false) and
half(2) should return (1, true).
**/
func halfNum( num int ) (int, bool) {
  halvedNum := num/2
  if num % 2 == 0 {
    return halvedNum, true
  }
  return halvedNum, false
}
/**
3. Write a function with one variadic parameter that finds the greatest number in a
list of numbers.
**/
func findLargestNumber( nums...int) int {
  largest := nums[0]
  for _, val := range nums{
    if val > largest{
      largest = val
    }
  }
  return largest
}

/**
4. Using makeEvenGenerator as an example, write a makeOddGenerator function
that generates odd numbers.
**/
func makeOddGenerator() func()int {
  oddNum := 1
  return func()int {
    oddNum += 2
    return oddNum
  }
}

/**
5. The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) =
fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).
**/
func fib(n int) int {
  if n == 0 {
    return 0
  }
  if n == 1 {
    return 1
  }
  return fib(n-1) + fib(n-2)
}

/**
6. What are defer, panic, and recover? How do you recover from a runtime panic?
Anything defered runs right before a function exits, whilst the panic raises a runtime error and recover 
handles a runtime error. 
**/
func testPanic(){
  defer func(){
    fmt.Println("I am actually running")
    recover()
  }()
  panic("We lost :(")
}

/**
7. How do you get the memory address of a variable?
Use the '&' symbol.
**/
func getRandomIntPointer() *int{
  rand.Seed(time.Now().UnixNano())
  randInt := rand.Intn(100)
  return &randInt
}

/**
8. How do you assign a value to a pointer?
The pointer can be deferenced before assignement.
**/
func assignedToTen( x *int){
  *x = 10
}

/**
9. How do you create a new pointer?
Initialising it via pointer initialisation like so: var x *int and then default address with 'new' like so x = new(int)

10. What is the value of x after running this program:
func square(x *float64) {
   *x = *x * *x
  }
func main() {
  x := 1.5
  square(&x)
}
X is values as: 1.5

11. Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y)
should give you x=2 and y=1).
**/
func swapInt(x *int, y *int){
  tmp := *x
  *x = *y
  *y = tmp
}


