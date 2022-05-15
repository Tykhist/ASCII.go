package main

import (
  // Module for printing with Println method
  "fmt"
  // Module for current time with Now method
  t "time"
)

func main () {
  // Prints art of Golang Gopher, a tree and a dog, and then the time of printing
  fmt.Println("      `.-::::::-.`                            ")
  fmt.Println("  .:-::::::::::::::-:.                        ")
  fmt.Println("  `_:::    ::    :::_`                        ")
  fmt.Println("   .:( ^   :: ^   ):.       ^                 ")
  fmt.Println("   `:::   (..)   :::.      (=)                ")
  fmt.Println("   `:::::::UU:::::::`     (===)               ")
  fmt.Println("   .::::::::::::::::.    (=====)              ")
  fmt.Println("   O::::::::::::::::O   (=======)             ")
  fmt.Println("   -::::::::::::::::-     |===|      __      _")
  fmt.Println("   `::::::::::::::::`     |===|    o'')}____//")
  fmt.Println("    .::::::::::::::.      |===|     `_/      )")
  fmt.Println("      oO:::::::Oo        /=====\\     (_(_/-(_/")
  // Gives current time
  fmt.Println(t.Now())
}
