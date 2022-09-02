

# gocolors

`go get github.com/mewzax/gocolors`


### Preview
![Capture](https://user-images.githubusercontent.com/75091300/176025830-10ab082b-44c9-437a-ad26-ad8aec6e8b85.JPG)

### Usage
```go
package main

import (
  "fmt"
  "github.com/mewzax/gocolors"
)

func main() {
  fmt.Println(gocolors.ColorizeForeground(gocolors.Green, "Hello World!"))
}
