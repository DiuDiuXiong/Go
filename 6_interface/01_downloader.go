package main

import (
	"Go/6_interface/infra"
	"Go/6_interface/test"
)

func getRetrieverInfra() infra.Retriever {
	return infra.Retriever{}
}

func getRetrieverTest() test.Retriever {
	return test.Retriever{}
}

func getRetriever(t string) retriever {
	switch t {
	case "infra":
		return &infra.Retriever{}
	case "test":
		return &test.Retriever{}
	default:
		return &test.Retriever{}
	}
}

/**
Here we want something that can .Get
1. so both retriever from infra & test can be called
2. so when we change the type, we don't need to change a lot of code from main() function, 低耦合
3. So here we need interface
``
type <interface_name> interface {
	<function_name>(<input_type>) (<return_type>, <...>)
}
``
And interface can be used as a type.

4. check the declaration down, and compare getRetriever with two other get retriever function, now only need to change
param of getRetriever
5. Note!! In Golang, both Retriever didn't confirm they implement retriever interface, but the code works.
- Also it is not dynamically decide type during run type, the actual calling function is decided during compilation
- As long as they implement functions of interface its okay, they (two retriever) don't have to specify they are implement what on code

6. The special feature is because in Golang, interface is defined by people who want to use it, not people who implement it
	For infra/test retriever, in original OOP language (Java):
	1. developer of infra/test has to first define retriever interface
	2. then they implement interface
	3. user (downloader here) needs to know the interface in order to call the function

	For Golang, it is the one who use functions implements interface (implement in download file), and other implementation try
	to satisfy requirements here. This enables type decided at compile time.

7. The type of r (retriever below), from the code below, can be seen that, is explicitly determined, which is because interface actually
contains :
	1. pointers to those objects
	2. object type,
so it can be changed during code. Also since it only contains pointer, thus pass by value is fine
8. interface contain type assertion (forced type transformation强制类型转换):
	- strict version: <interface_object>.(<actual type>) can return the actual object, if <actual type> is not the
		actual type (some other possible type), will cause runtime error
	- loss version:
		- if <interRetriever>, <ok_var> := <interface_object>.(<actual type>); <ok_var> {
			...
		}
9. insertion can type switch:
- switch `r.(type)`
10. `interface{}` represents any format!! (don't need to implement anything) check imooc 6-4 14:12 for more example
11. In order to combine more than one interface, just:
- type <composition_interface_name> interface {
	<interface_1>
	<interface_2>
}
And who ever implements all interface above is able to be use as <composition_interface_name>. Check RetrieverComposition for example.
*/

type retriever interface {
	Get(string) (string, error)
}

type anotherInterface interface {
	PrintMessage(string)
}

/*
*
This is for composition
*/
type RetrieverComposition interface {
	retriever
	anotherInterface
}

func session(s RetrieverComposition) {
	if msg, err := s.Get("https://google.com"); err == nil {
		s.PrintMessage(msg)
	}
}

/*
func main() {
	var r retriever = getRetriever("test")
	fmt.Printf("%T %v\n", r, r)
	if resp, err := r.Get("https://www.google.com"); err == nil {
		fmt.Println(resp)
	}

	r = getRetriever("infra")
	fmt.Printf("%T %v\n", r, r)
	if _, err := r.Get("https://www.google.com"); err == nil {
		//fmt.Println(resp)
	}

	fmt.Println(&r)
	if r, ok := r.(*infra.Retriever); ok {
		fmt.Println(&r)
	}
	//r.(*infra.Retriever).Get("") // error

	var pythonList []interface{}
	pythonList = append(pythonList, 1)
	pythonList = append(pythonList, "b")
	pythonList = append(pythonList, true)
	fmt.Println(pythonList)

	// try composition
	c := test.Retriever{}
	session(&c)
	//session(&r) error, r cannot

}
*/
