# request

```go
package main

import (
	"fmt"
	"github.com/scott-x/request"
)
const url = "http://www.baidu.com"

func main() {
	content, _ := request.Get(url)
	fmt.Println(string(content))
}
```