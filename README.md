> About learn golang...

# flag

> 解析参数

- 用法

```go
package main

import (
	"flag"
	"fmt"
)

var (
	port = flag.Int("port", 5300, "server port")
)

func main() {
	flag.Parse()

	fmt.Printf("server port is: %d\n", *port)
}
```

- 运行

```go
➜  _flag git:(main) ✗ go run main.go
server port is: 5300
➜  _flag git:(main) ✗ go run main.go -port 3000
server port is: 3000
```

- help

```go
➜  _flag git:(main) ✗ go run main.go -h
-port int
server port (default 5300)
```


