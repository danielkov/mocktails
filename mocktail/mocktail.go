package mocktail

import (
	"reflect"
	"fmt"
)

func setMockValue(tag string) {
	fmt.Println(tag)
}

func Mock(s interface{}) interface{} {
	v := reflect.ValueOf(s)
	n := v.NumField()

	for i := 0; i < n; i++ {
		tag := v.Type().Field(i).Tag.Get("mock")

		if tag == "" || tag == "-" {
			continue
		}

		setMockValue(tag)

	}
}
