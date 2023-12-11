package main

import (
  "fmt"
  "bufio"
  "log"
  "os"
)

func main() {

  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  var inputText string

  for scanner.Scan() {
    inputText = scanner.Text()
  }

  var x int = 0
  var y int = 0
 
  var xCoordinates = make([]int, (len(inputText)+1))
  var yCoordinates = make([]int, (len(inputText)+1))

  xCoordinates[len(inputText)] = 0
  yCoordinates[len(inputText)] = 0

  for i := 0; i < len(inputText); i++ {

    if string(inputText[i]) == ">" {
      x++
    } else if string(inputText[i]) == "<" {
      x--
    } else if string(inputText[i]) == "^" {
      y++
    } else if string(inputText[i]) == "v" {
      y--
    }
    
    xCoordinates[i] = x
    yCoordinates[i] = y
  }

  var uniqueHouses int = len(xCoordinates)

  for i := 0; i <len(xCoordinates); i++ {

    isDuplicate := false

    for j := i; j <len(xCoordinates); j++ {

      if (xCoordinates[i] == xCoordinates[j] && 
          yCoordinates[i] == yCoordinates[j] && 
          i != j) {

        isDuplicate = true
      }
    }

    if isDuplicate {
      uniqueHouses--
    }
  }

  fmt.Println("Number of houses with at least one present: ", uniqueHouses)
}
