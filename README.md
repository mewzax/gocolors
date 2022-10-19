`go get github.com/mewzax/gocolors`

### Usage
```go
package main

import (
  "fmt"
  "github.com/mewzax/gocolors"
)

func main() {
  fmt.Println(gocolors.Colorize(gocolors.Green, "Hello World!"))
  fmt.Println(gocolors.Colorize(gocolor.RGB(66, 245, 233), "Custom RGB"))
}
