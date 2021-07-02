package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	Name      string         `json:"name"`
	Age       int            `json:"age,omitempty"` // 空的字段不进行序列化
	Weight    int            `json:"weight"`
	Profile   *PersonProfile `json:"profile"`
	Addresses *[]Address     `json:"addresses"`
}

type PersonProfile struct {
	Hobby string
}

type Address struct {
	Country string
	Region  string
	Detail  string
}

func TestMarshal() ([]byte, error) {
	profile := PersonProfile{
		Hobby: "swimming",
	}

	addressOne := Address{
		Country: "China",
		Region:  "ChangSha",
	}

	addressTwo := Address{
		Country: "Us",
		Region:  "芝加哥",
	}

	meAddresses := &[]Address{addressOne, addressTwo}

	me := Person{
		//Name: "Ethan",
		Age:       18,
		Weight:    75,
		Profile:   &profile,
		Addresses: meAddresses,
	}

	b, err := json.Marshal(me)

	if err != nil {
		return nil, err
	}
	fmt.Printf("%s", b)
	return b, nil
}

func TestUnmarshal(MarshalData []byte) {
	var me Person
	err := json.Unmarshal(MarshalData, &me)
	if err != nil {
		panic(err)
	}
	spew.Dump(me)
}

func main() {
	marshalData, err := TestMarshal()
	if err != nil {
		panic(err)
	}
	TestUnmarshal(marshalData)
}
