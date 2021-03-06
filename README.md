# Description
Migrations lib for Golang sql

[![GitHub license](https://img.shields.io/github/license/c0de4un/migrations)](https://github.com/c0de4un/migrations/blob/main/LICENSE)[![GitHub stars](https://img.shields.io/github/stars/c0de4un/migrations)](https://github.com/c0de4un/migrations/stargazers)[![GitHub issues](https://img.shields.io/github/issues/c0de4un/migrations)](https://github.com/c0de4un/migrations/issues)
![GitHub issues](https://img.shields.io/badge/language-Go-blue)

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
