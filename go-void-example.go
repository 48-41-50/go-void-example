package main

import (
    "fmt"
    "strconv"
)

func main () {
    var x interface{}
    var y int
    var z string
    
    y = 1234
    fmt.Printf("y = %d\n", y);
    x = "nil (a string)"
    fmt.Printf("x = \"%s\"\n", x);
    x = y
    fmt.Printf("x = %d\n", x);
    x = &y
    fmt.Printf("x = %d\n", *(x.(*int)));
    *(x.(*int)) += 1
    fmt.Printf("x = %d\n", *(x.(*int)));
    fmt.Printf("y = %d\n", y);
    z = strconv.FormatInt(int64(*(x.(*int))), 10)
    fmt.Printf("z = \"%s (a string)\"\n", z);
    
}
