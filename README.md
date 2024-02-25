# Facades in Go

## Installation

go-facades supports the last 2 versions of Go and requires a Go version with [modules](https://github.com/golang/go/wiki/Modules) support. So make sure to initialize a Go module:

```bash
go mod init github.com/my/repo
```

Then install go-facades

```bash
go get github.com/Aikintech/go-facades
```

## Quick start

Crypt (encryption)

```go
import (
    "fmt"

    facades "github.com/Aikintech/go-facades"
)

func main() {
    // Secret key must be 16 or 24 or 32 or 64 characters long
    var secretKey = "Om8FLaOZc0Y2IVx58K9MGTgm8RCmmE0L"
    var stringToBeEncrypted = "test-string"

    encrypted, err := facades.Crypt(secretKey).EncryptString(stringToBeEncrypted)
    if err != nil {
        panic(err)
    }

    // Decrypt encoded/encrypted string
    decrypted, err := facades.Crypt(secretKey).DecryptString(encrypted)
    if err != nil {
        panic(err)
    }

    fmt.Println(decrypted)
}

```
