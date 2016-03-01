package vertebrae

type Vertebrae map[string]interface{}

// Add service
func (v Vertebrae) Add(name string, object interface{}) Vertebrae {
	v[name] = object
	return v
}

// Remove service
func (v Vertebrae) Remove(name string) Vertebrae {
	delete(v, name)
	return v
}

// Example:
//
// myApp := &MyApp{}
//
// container := make(Vertebrae)
// container.Add("my.service", myApp)
// container.Add("my.string", "Testing 123")
//
// fmt.Println(container["my.service"].(*MyApp).TestService(23))
// fmt.Println(container["my.string"].(string))
//
// type MyApp struct{}
//
// func (myApp *MyApp) TestService(param int) int {
// 	return 12 + param
// }
