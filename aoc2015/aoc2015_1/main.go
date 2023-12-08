package main

import (
  "fmt"
  "bufio"
  "log"
  "os"
)

func main() {

  file, err := os.Open("/home/dedu/cod/AoC/aoc2015/aoc2015_1/input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  var x string
  var y int = 0

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    x = scanner.Text()
  }

  for i := 0; i <len(x); i++ {
    if string(x[i]) == "(" {
      y++
    } else if string(x[i]) == ")" {
      y--
    }

    if y == -1 {
      fmt.Println(i+1)
    }
  }

  fmt.Println(y)

 }
