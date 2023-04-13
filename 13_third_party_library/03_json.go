package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	ID         string
	Name       string
	Quantity   int
	TotalPrice float64
	unseen     int
}

func printStructOriginal() {
	o := Order{
		ID:         "1234",
		Name:       "learn go",
		Quantity:   3,
		TotalPrice: 30,
		unseen:     0,
	}

	fmt.Printf("%v\n", o)  // %v to display all values
	fmt.Printf("%+v\n", o) // %+v name/value
}

func jsonPrint() {
	o := Order{
		ID:         "1234",
		Name:       "learn go",
		Quantity:   3,
		TotalPrice: 30,
		unseen:     0,
	}

	b, _ := json.Marshal(o) // turn struct into a byte sequence that can be transmitted over internet (second output is error)
	fmt.Printf("%s\n", b)   // however, unseen is not presented, this is because it is lowercase thus private
}

type OrderSelfDefined struct {
	ID         string  `json:"id"`
	Name       string  `json:"name,omitempty"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
}

type ComposeType struct {
	MyOrder OrderSelfDefined `json:"item"`
	ID      string           `json:"id"`
}

func jsonSelfDefinedName() {
	o := OrderSelfDefined{
		ID:         "1234",
		Name:       "learn go",
		Quantity:   3,
		TotalPrice: 30,
	}

	o2 := OrderSelfDefined{
		ID:         "1234",
		Quantity:   3,
		TotalPrice: 30,
	}

	b, _ := json.Marshal(o)
	fmt.Printf("%s\n", b)

	b2, _ := json.Marshal(o2)
	fmt.Printf("%s\n", b2)
}

func checkCompose() {
	o := OrderSelfDefined{
		ID:         "1234",
		Name:       "learn go",
		Quantity:   3,
		TotalPrice: 30,
	}

	o2 := ComposeType{
		MyOrder: o,
		ID:      "learn go",
	}

	b2, _ := json.Marshal(o2)
	fmt.Printf("%s\n", b2)
}

func unmarshal() {
	s := `{"item":{"id":"1234","name":"learn go","quantity":3,"totalPrice":30},"id":"learn go"}`
	var o ComposeType
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

/*
1. json.Marshal(<struct>) to get json as a sequence of bytes
2. after each struct variable definition, add
	<var name> <var type> `json:"<name>"`
to define the name you wish that field have when converted to json.
3. add omitempty within "" so if that variable is not defined (zero value), will not be output when called from json.Marshal
	<var name> <var type> `json:"<name>,omitempty"`
4. For compose, its json compose json, even if pointer type, json.Marshal(...) will also output same thing
5. For compose, also slice of object will present in json as list
6. For unmarshal
	- Create a object of same type, which will be used for store unmarshal data form json
	- err := json.Unmarshal([]byte(<json_string>, &<storageObjectCreated>))
*/

func main() {
	// jsonPrint()
	//  jsonSelfDefinedName()
	// checkCompose()
	unmarshal()
}
