package reflectoinPractice

import(
	"fmt"
	"reflect"
)

type interfaceA interface {
	ToString(message string) bool
}

func GenerateFunction(){
	type BigNumber int
	reflect.MakeFunc()
}