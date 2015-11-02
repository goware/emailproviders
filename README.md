# List of public email providers for Golang

[![GoDoc](https://godoc.org/github.com/goware/emailproviders?status.svg)](https://godoc.org/github.com/goware/emailproviders)

## Example

```go
import "github.com/goware/emailproviders"

func main() {
    emailproviders.Exists("example@gmail.com")   // true
    emailproviders.Exists("example@company.com") // false
}
```

## License
Licensed under the [MIT License](./LICENSE).
