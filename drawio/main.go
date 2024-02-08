package main

func main() {
	//
}

// MyInterface only has one method, notice the signature return value
type MyInterface interface {
	foo() bool
}

// MyStruct1 will implement the foo() bool function so it will have an "extends" association with MyInterface
type MyStruct1 struct {
}

func (s1 *MyStruct1) foo() bool {
	return true
}

// MyStruct2 will be directly composed of MyStruct1 so it will have a composition relationship with it
type MyStruct2 struct {
	MyStruct1
}

// MyStruct3 will have a foo() function but the return value is not a bool, so it will not have any relationship with MyInterface
type MyStruct3 struct {
	Foo MyStruct1
}

func (s3 *MyStruct3) foo() {

}
