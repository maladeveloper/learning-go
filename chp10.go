package main

import(
  "fmt"
  "time"
)

func main(){
  sleepSecs := 2
  sleepFunc(sleepSecs)
  fmt.Printf("Printing after sleeping for %d seconds\n", sleepSecs)
}

/**
1. How do you specify the direction of a channel type?
The direction of the channel is specified during the initialisation of the channel
as either the sender `func pinger(c chan <- string)` or as the receiver `func printer(c <- chan string)`
if this is not specified then the channel is bidirectional such that it can send and receive messages.

2. Write your own Sleep function using time.After.
**/
func sleepFunc( sleepSecs int){
  select{
    case <- time.After(time.Duration(sleepSecs) * time.Second):
      fmt.Println("Timout called")
  }
}

/**
3. What is a buffered channel? How would you create one with a capacity of 20?
A buffered channel is a channel that is non-blocking so long as its capacity has not been reached, 
and which at that point, becomes indicernible from a regular channel.
A buffered channel with the capacity of 20 can be created like this: 
`c := make(chan int, 20)`
**/
