package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTimeFormatConvert(t *testing.T) {
	input := "1:15 pm"
	expected := "T131500"

	actual := TimeFormatConvert(input)

	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("Actual: %#v\n Expected %#v\n", actual, expected)
	}
}
