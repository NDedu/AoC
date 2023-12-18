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

  var inputText string
  var niceStringsCountPart1 int = 0
  var niceStringsCountPart2 int = 0


  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    inputText = scanner.Text()

    var vowelsCheck int = 0
    var duplicateLetterCheck int = 0
    var stringCheck int = 0

    var repeatingLetterP2Check int = 0
    var twicePairsP2Check int = 0

    for i := 0; i < len(inputText); i++ {

      if (string(inputText[i]) == "a" ||
          string(inputText[i]) == "e" ||
          string(inputText[i]) == "i" ||
          string(inputText[i]) == "o" ||
          string(inputText[i]) == "u") {

        vowelsCheck++
      }
    }
  
    for i := 0; i < (len(inputText)-1); i++ {

      if (string(inputText[i]) == string(inputText[i+1])) {

        duplicateLetterCheck++
      }
    }
  
    for i := 0; i < (len(inputText)-1); i++ {

      if ((string(inputText[i]) == "a" && string(inputText[i+1]) == "b") ||
          (string(inputText[i]) == "c" && string(inputText[i+1]) == "d") ||
          (string(inputText[i]) == "p" && string(inputText[i+1]) == "q") ||
          (string(inputText[i]) == "x" && string(inputText[i+1]) == "y")) {
        
        stringCheck++
      }
    }
 
    if (vowelsCheck > 2 && duplicateLetterCheck > 0 && stringCheck == 0) {

      niceStringsCountPart1++
    }
    
    // Part 2
    for i := 0; i < (len(inputText)-2); i++ {

      if string(inputText[i]) == string(inputText[i+2]) {

        repeatingLetterP2Check++   
      }
    }

    for i := 0; i < (len(inputText)-3); i++ {

      for j := i; j < len(inputText)-3; j++ {

        if (string(inputText[i]) == string(inputText[j+2]) &&
            string(inputText[i+1]) == string(inputText[j+3])) {

          twicePairsP2Check++
        }
      }
    }

    if (repeatingLetterP2Check > 0 && twicePairsP2Check > 0) {

      niceStringsCountPart2++
    }
  }

  fmt.Println("Number of nice strings in part 1: ", niceStringsCountPart1)
  fmt.Println("Number of nice strings in part 2: ", niceStringsCountPart2)
 }
