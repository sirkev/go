package main

import "fmt"
func sum(c chan int, co chan int){
    for{
       x:= <-c
       x:= <-c
//prints sum for whats happening on the queue
       fmt.Println(x, "+",y, "=", x+y
       co <-x +y 
}
}


func main(){
    c := make(chan int,6)
    co := make(chan int,3)
    go sum(c,co)
  //input output queue
    c <- 10
    c <- 15
    c <- 99
    c <- 1
    
    r := <-co
    fmt.println(r)
    r := <-co
    fmt.println(r)
    r :=<-co
    fmt.println(r)

}
}
