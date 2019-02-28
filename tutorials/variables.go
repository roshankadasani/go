package main

import "fmt"

func main()  {
  var num float64
  num = 20.2
  fmt.Println(num)
  fmt.Printf("num is of type %T\n", num)

  var a = "initial"
  fmt.Println(a)

  var b,c int = 1,2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)
}
