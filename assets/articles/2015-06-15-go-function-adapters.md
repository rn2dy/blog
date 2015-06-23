In this post I will talk about function adapter in Go, the purpose and application.

Here is an example of using function adapter to decorate a integer generator for differnt output sequence

    package main

    import (
      "fmt"
    )

    type generator interface {
      gen() chan int
    }

    type squareSeq func() chan int

    func (q squareSeq) gen() chan int {
      out := make(chan int)
      in := q()
      go func() {
        for {
          k := <- in
          out <- k * k
        }
      }()
      return out
    }

    func intSeq() chan int {
      out := make(chan int)
      k := 1
      go func() {
        for {
          out <- k
          k++
        }
      }()
      return out
    }

    func mygen(g generator) {
      out := <- g.gen()
      k := 10
      for k > 0 {
        fmt.Println(<-out)
        k--
      }
    }

    func main() {
      out := squareSeq(intSeq).gen()

      for {
        k := <-out
        fmt.Println(k)
        if k == 5 {
          break
        }
      }
    }
