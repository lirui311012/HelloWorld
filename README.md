# HelloWorld
just say hello

1. Clone the repository and initialize the module
```bash
git clone git@github.com:lirui311012/HelloWorld.git

cd HelloWorld && go mod init github.com/lirui311012/HelloWorld
```


2. format the code && check the code
```bash
go fmt ./... && go vet ./...
```

3. test the code
```bash
go test -v .
go test -v ./...
go run ./cmd/helloworld -name jack
```

## Install

import code

```bash
go get github.com/lirui311012/HelloWorld@latest
```

install cmd

```bash
go install github.com/lirui311012/HelloWorld/cmd/helloworld@latest

helloworld -name jack2
```

## Example

Here's a simple example as follows:

```go
package main

import (
  "fmt"
  hello "github.com/lirui311012/HelloWorld"
)

func main() {
  result := hello.Hello("jack")
  fmt.Println(result)
}
```

# Upload

