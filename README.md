# Facades in Go

## Installation

gofacades supports the last 2 versions of Go and requires a Go version with [modules](https://github.com/golang/go/wiki/Modules) support. So make sure to initialize a Go module:

```bash
go mod init github.com/my/repo
```

Then install gofacades

```bash
go get github.com/Aikintech/gofacades
```

## Quick start

Crypt (encryption)

```go
import (
    "fmt"

    facades "github.com/Aikintech/gofacades"
)

func main() {
    // Secret key must be 16 or 24 or 32 or 64 bytes
    var (
        secretKey = "Om8FLaOZc0Y2IVx58K9MGTgm8RCmmE0L"
        stringToBeEncrypted = "test-string"
    )
    var crypt = facades.Crypt(secretKey)

    encrypted, err := crypt.EncryptString(stringToBeEncrypted)
    if err != nil {
        panic(err)
    }

    // Decrypt string
    decrypted, err := crypt.DecryptString(encrypted)
    if err != nil {
        panic(err)
    }

    fmt.Println(decrypted)
}

```

Redis

```go
import (
    "fmt"

    r "github.com/gofiber/storage/redis/v3"
    facades "github.com/Aikintech/gofacades"
)

func main() {
    var prefix = "myapp"
    var key = "tel"
    var rdb = facades.Redis(r.Config{
        Host:      "127.0.0.1",
        Port:      6379,
        Username:  "",
        Password:  "",
        Database:  0,
    }, prefix)

    // Store value by key
    err := rdb.Set(key, []byte("0244123456"))
    if err != nil {
        panic(err)
    }

    // Retrieve stored value by key
    val, err := rdb.Get(key)
    if err != nil {
        panic(err)
    }

    // Cast val ([]byte) to string
    fmt.Println(string(val))
}

```
