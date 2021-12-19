package main
import "fmt"

func main(){
  printMult()
  lenOfStr()
}
/**
1. How are integers stored on a computer?
Stored in binary as bytes, the size of which is determined by the machine.

2. We know that (in base 10) the largest one-digit number is 9 and the largest twodigit number is 99. Given that in binary the largest two-digit number is 11 (3),
the largest three-digit number is 111 (7) and the largest four-digit number is 1111
(15), what’s the largest eight-digit number?
255.

3. Although overpowered for the task, you can use Go as a calculator. Write a pro‐
gram that computes 32,132 × 42,452 and prints it to the terminal (use the * oper‐
ator for multiplication).
**/
func printMult(){
  fmt.Println("32,132 x 42,452 =", 32132 * 42452)
}

/**
4. What is a string? How do you find its length?
A string is a array of characters usually of length one byte and can be found by using 'len' function.
**/
func lenOfStr(){
  printStr := "Hello World"
  fmt.Printf("The string '%s' has a length of %d\n", printStr, len(printStr))
}

/**
5. What’s the value of the expression (true && false) || (false && true) || !(false && false)?
(true && false) || (false && true) || !(false && false) = false || false || true =  true
**/



