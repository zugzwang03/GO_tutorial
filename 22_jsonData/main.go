package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json data in depth")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	someCourses := []course{
		{"ReactJS Bootcamp", 299, "Someplatform", "abc123", []string{"web dev", "js"}},
		{"MERN Bootcamp", 298, "Someplatform", "bcd654", []string{"full stack dev", "js"}},
		{"Angular Bootcamp", 249, "Someplatform", "htc567", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(someCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "Someplatform",
                "Tags": ["web dev","js"]
    }
	`)

	var myCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// add data as key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and Value is %v and type of data is %T\n", k, v, v)
	}
}
