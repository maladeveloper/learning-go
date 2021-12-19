package main

import "fmt"

/**
1. What is whitespace?
A space, tab or new line or carriage return (ENTER).

2. What is a comment? What are the two ways of writing a comment?
A comment is a line of code that is not executed and can be written with // for a single line
or /** for a multiline.

3. Our program began with package main. What would the files in the fmt package
begin with?
The files in the fmt package would begin with 'package fmt'.

4. We used the Println function defined in the fmt package. If you wanted to use
the Exit function from the os package, what would you need to do?
import "os" os.Exit()

5. Modify the program we wrote so that instead of printing Hello, World it prints
Hello, my name is followed by your name.
**/
func main(){
  fmt.Println("Hello, my name is Malavan Srikumar")
}
