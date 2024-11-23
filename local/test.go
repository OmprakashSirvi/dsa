package local

import "fmt"

func Test() {
	test := map[string]string{"test": "test"}

	pasPointer(&test)

	fmt.Printf("test : %v", test)
}

func pasPointer(test *map[string]string) {
	fmt.Printf("test : %v\n", test)
	v := map[string]string{"test": "not_Test"}

	*test = v
}
