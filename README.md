# request

### API
- `func Get(url string) ([]byte, error)`: common function to parse url to html string(with uft8 encoding)
- `func Fetch(url string) ([]byte, error)`: this function is upper lever for Get, so it can be used for crawler
- `func Download(url,filepath string)  error`: download the file from the network to local filepath, if filepah not exists, will `mkdir -p` the Dir

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

### download

```go
package main

import "github.com/scott-x/request"

func main() {
	request.Download("https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=1157181355,2562840370&fm=26&gp=0.jpg","./imgs/a.jpg")
}
```

### tools

```bash
go install github.com/scott-x/request/cmd/request
request -url http://xxx.xx #determine if the url supports crawler
```