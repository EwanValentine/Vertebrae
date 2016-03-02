# Vertebrae

## Simple DI container, written in Go

```go
// Example:
myApp := &MyApp{}
container := make(Vertebrae)

container.Add("my.service", myApp)
container.Add("my.string", "Testing 123")

fmt.Println(container["my.service"].(*MyApp).TestService(23))
fmt.Println(container["my.string"].(string))

type MyApp struct{}

func (myApp *MyApp) TestService(param int) int {
	return 12 + param
}
```