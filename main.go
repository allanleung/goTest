package main

import (
     "fmt"
     "math"
     "os"
)

func main() {
  var myNumber = 1.2345
  roundedUp := math.Ceil(myNumber)
  //Call the Floor function from the "math" package.
  roundedDown := math.Floor(myNumber)
  // Call the Println function from the "fmt" package.
  fmt.Println("This is rounding up and down from 1.234")
  fmt.Println(roundedUp, roundedDown)
  
  /*
  var a int
  a = 1
  var b, c int
  b = 2
  c = 3
  var d = 5
  e := 6
  
  fmt.Println(a, b, c, d, e)

  */

  var myInt int
  var myString string
  var myBool bool
  fmt.Println("myInt", myInt)       // 0
  fmt.Println("myString", myString) //""
  fmt.Println("myBool", myBool)     //false

  beforeLoop := 888
  for i := 1; i <= 5; i++ {
    //inLoop := 999
    // if its not in use, please delete it
    fmt.Println(i)
  }

  fmt.Println(beforeLoop) // OK
  //fmt.Println(i)          // Error!
  //fmt.Println(inLoop)     // Error!

  fileInfo, error := os.Stat("existent.txt")
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println(fileInfo.Size())
    }
    fileInfo, error = os.Stat("nonexistent.txt")
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println(fileInfo.Size())
    }

    // Pointer related 

    var aValue float64 = 1.23
    var aPointer *float64 = &aValue
    fmt.Println("aPointer:", aPointer) // prints something like "aPointer: 0x1040a128"
    fmt.Println("*aPointer:", *aPointer) // prints "*aPointer: 1.23"

    // Array and testing 

     var months [3]string
    months[0] = "Apr"
    months[1] = "May"
    months[2] = "Jun"
    var salesByMonth [3]float64
    salesByMonth[0] = 1710.26
    salesByMonth[1] = 2245.97
    salesByMonth[2] = 3032.40
    fmt.Println(months[0], salesByMonth[0])
    fmt.Println(months[1], salesByMonth[1])
    fmt.Println(months[2], salesByMonth[2])


    
    // Array with Slice 
    // Two basic slice 
    a := [5]int{0, 1, 2, 3, 4}
    //a[2] = 88 // Make second ele 88
    s1 := a[0:3] // This reads from index 0 to 3 but NOT including index 3
    s2 := a[1:5] // From 1 to 4, STOP on 5

    // [0 1 88 3 4] [0 1 88] [88 3 4] [1 88 3]

    s2[0] = 9999

    // The 0 index of the s2 slice. Since s2 range is [1 2 3 4] and the 0 index is 1
    // 1 got replaced by 9999
    // [0 999 88 3 4] [0 999 88] [999 88 3 4] [999 88 3]

    s3 := a[1:4] // Should print [1 2 3]

    //s1 = s1[0:4]
    //s2 = s2[0:4]

    fmt.Println(a, s1, s2, s3)
    fmt.Println("len(s1):", len(s1), "cap(s1):", cap(s1))
    fmt.Println("len(s2):", len(s2), "cap(s2):", cap(s2))


    // slice example 
    s := []int{1, 2, 33333}
    s = append(s, 4, 5)
    s = append(s, 6, 7, 8)
    fmt.Println(s)

}