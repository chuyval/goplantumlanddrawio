# Example generating DrawIO diagrams

Golang Example
```golang
package main

func main() {
	//
}

//MyInterface only has one method, notice the signature return value
type MyInterface interface {
	foo() bool
}

//MyStruct1 will implement the foo() bool function so it will have an "extends" association with MyInterface
type MyStruct1 struct {
}

func (s1 *MyStruct1) foo() bool {
	return true
}

//MyStruct2 will be directly composed of MyStruct1 so it will have a composition relationship with it
type MyStruct2 struct {
	MyStruct1
}

//MyStruct3 will have a foo() function but the return value is not a bool, so it will not have any relationship with MyInterface
type MyStruct3 struct {
	Foo MyStruct1
}

func (s3 *MyStruct3) foo() {

}
```

Generate PlantUML
```bash
goplantuml drawio 
```

Generates:
```
@startuml
namespace main {
    interface MyInterface  {
        - foo() bool

    }
    class MyStruct1 << (S,Aquamarine) >> {
        - foo() bool

    }
    class MyStruct2 << (S,Aquamarine) >> {
    }
    class MyStruct3 << (S,Aquamarine) >> {
        + Foo MyStruct1

        - foo() 

    }
}
"main.MyStruct1" *-- "main.MyStruct2"

"main.MyInterface" <|-- "main.MyStruct1"

@enduml
```

Diagram using PlantUML Online Server: [UML Diagram](https://www.plantuml.com/plantuml/uml/SoWkIImgAStDuSfBp4qjBaXCJbN8JSpCKwZcKW22pBoIrA8qaA1lfIUS3PA40bs5jFny3Ks5fFpy72wmQ2sOJCv9B2u6QWekAIfDBZ5KiB5Hq0ZMSImiJSnDBChCIzLKiBCZsHWZ7CRWr61Co5vj1RVyV8GpkJ0S17DCBPSIA5Zqu4gL5BHqTHLG1OsKk63A4A6bf61JWw7I4AZI8JKl1HZm0000) 

Generate PlantUML with Options
```bash
goplantuml -aggregate-private-members -show-aggregations -show-aliases -show-compositions -show-connection-labels -show-implementations -show-options-as-note drawio
```

Generates
```
@startuml
legend
<u><b>Legend</b></u>
Render Aggregations: true
Render Fields: true
Render Methods: true
Private Aggregations: true
end legend
namespace main {
    interface MyInterface  {
        - foo() bool

    }
    class MyStruct1 << (S,Aquamarine) >> {
        - foo() bool

    }
    class MyStruct2 << (S,Aquamarine) >> {
    }
    class MyStruct3 << (S,Aquamarine) >> {
        + Foo MyStruct1

        - foo() 

    }
}
"main.MyStruct1" *-- "extends""main.MyStruct2"

"main.MyInterface" <|-- "implements""main.MyStruct1"

"main.MyStruct3""uses" o-- "main.MyStruct1"

@enduml
```

Diagram using PlantUML Online Server: [UML Diagram](https://www.plantuml.com/plantuml/uml/SoWkIImgAStDuSfBp4qjBaXCJbN8JSpCKwZcKW22pBoIrA8qaA1lfIUS3PA40bs5jFny3Ks5fFpy72wmQ2sOJCv9B2u6QWekAIfDBZ5KiB5Hq0ZMSImiJSnDBChCIzLKiBCZsHWZ7CRWr61Co5vj1RVyV8GpkJ0S17DCBPSIA5Zqu4gL5BHqTHLG1OsKk63A4A6bf61JWw7I4AZI8JKl1HZm0000)

Check Rendered UML Diagram at this site: https://www.plantuml.com/

Render Repositories here: https://www.dumels.com/
Just type the github URL
