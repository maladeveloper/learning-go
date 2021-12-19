package main

import "fmt"

func main(){
  fmt.Println("Nothing to print.")
}

/**
1. Why do we use packages?
To reduce code complexity through seperation of functionality and to reduce
compile times as packages are already pre-compiled

2. What is the difference between an identifier that starts with a capital letter and
one that doesn’t (e.g., Average versus average)?
The identifier that starts with a capital letter is public in that other packages and 
programs may be able to referance/call it.

2. What is the difference between an identifier that starts with a capital letter and
one that doesn’t (e.g., Average versus average)?

3. What is a package alias? How do you make one?
During the import specify the variable that it is to be called.
**/
