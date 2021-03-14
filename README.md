# request

### API
- `func Get(url string) ([]byte, error)`: common function to parse url to html string(with uft8 encoding)
- `func Fetch(url string) ([]byte, error)`: this function is upper lever for Get, so it can be used for crawler

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

### tools

```bash
go install github.com/scott-x/request/cmd/request
request -url http://xxx.xx #determine if the url supports crawler
```