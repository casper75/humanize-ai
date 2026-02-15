## Installation

```bash
go get github.com/casper75/humaize-ai
```

## Usage

```go
package main

import (
  "fmt"
  "github.com/casper75/humaize-ai"
)

func main() {
  hm := NewHumanizer()
  test = "Hello\u200b\xa0World!  "
  result := hm.HumanizeString(test)
  fmt.Println(result)
}