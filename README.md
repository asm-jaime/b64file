# b64file
load a base64 string from a file/save a base64 string to a file.

# example
```go
package main

import(
  "log"
  b64f "github.com/asm-jaime/b64file"
)

func main() {
  b64str, err := b64f.FileToB64('./data.correct.jpeg')
  log.Println(err)
  log.Println(b64str)
  err = b64f.B64ToFile('./data.result.jpeg', b64str)
  log.Println(err)
}
```

# License
CC0
