# Description
Migrations lib for Golang sql

# Requirements
* Go 1.18+
* MySQL

# Test
*
```bash
    $go test
```

# Examples
## Up to MAX version
```go
package main

import (
    "github.com/c0de4un/migrations"
    "fmt"
)

func main() {
    err := migrations.Up("/configs/db.xml")
    if err != nil {
        fmt.Error(err)
    }
}
```

## Down to specific version
```go
package main

import (
    "github.com/c0de4un/migrations"
    "fmt"
)

func main() {
    err := migrations.Down("/configs/db.xml", 13)
    if err != nil {
        fmt.Error(err)
    }
}
```
