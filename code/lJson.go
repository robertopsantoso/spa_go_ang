package main

import (
	"fmt" // to printout
	"encoding/json" // to marshal(parse) from struct and unmarshal(stringify) to struct
	"os" // to open file
	"io/ioutil" // to read data from file
)

type Notebook struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Info Info `json:"info"`
}

type Info struct {
	Creadate string // case-sensitive
	LastUpdate string
}

func main () {
	// predetermined Object
	info1 := Info{Creadate:"2020-12-07", LastUpdate:"2020-12-07"}
	book1 := Notebook{Id:1, Title:"Title 1", Content:"This is the content of first note", Info:info1}
	book2 := `{"id":2, "title":"Title 2", "content":"This is the content of second note", "info":{"Creadate":"2020-12-07", "LastUpdate":"2020-12-07"}}`

	obj1, err1 := json.MarshalIndent(book1, "", "	")
	if err1 != nil {
		fmt.Println(err1)
	}
	var obj2 Notebook
	err2 := json.Unmarshal([]byte(book2), &obj2)
	if err2 != nil {
		fmt.Println(err2)
	}
	// not predetermined Object
	var obj3 map[string]interface{}
	err3 := json.Unmarshal([]byte(book2), &obj3)
	if err3 != nil {
		fmt.Println(err3)
	}

	// open json file
	jsonFile, err4 := os.Open("assets/note.json")
	if err4 != nil {
		fmt.Println(err4)
	}
	defer jsonFile.Close()

	// read from file
	book3, err5 := ioutil.ReadAll(jsonFile)
	if err5 != nil {
		fmt.Println(err5)
	}
	var obj4 Notebook
	json.Unmarshal(book3, &obj4)

	// from array of struct
	var notes = []Notebook {
		Notebook{Id:4, Title:"Title 4", Content:"This is the content of fourth note", Info:info1},
		Notebook{Id:5, Title:"Title 5", Content:"This is the content of fifth note", Info:info1},
	}
	obj5, err6 := json.MarshalIndent(notes, "", "	")
	if err6 != nil {
		fmt.Println(err6)
	}

	// open json file
	arrJsonFile, err7 := os.Open("assets/arrJson.json")
	if err7 != nil {
		fmt.Println(err7)
	}
	defer arrJsonFile.Close()

	// read from file
	book4, err8 := ioutil.ReadAll(arrJsonFile)
	if err8 != nil {
		fmt.Println(err8)
	}
	var arr1 []Notebook
	json.Unmarshal(book4, &arr1)

	fmt.Println("From Struct to Obj")
	fmt.Printf("%v\n", book1)
	fmt.Printf("%v\n", string(obj1))
	fmt.Println("\nFrom Obj to Struct")
	fmt.Printf("%v\n", book2)
	fmt.Println("\nFrom Obj2")
	fmt.Printf("%v\n", obj2)
	fmt.Printf("%v\n", obj2.Info.LastUpdate)
	fmt.Println("\nFrom non predetermined Obj")
	fmt.Printf("%v\n", obj3)
	fmt.Printf("%v\n", obj3["info"].(map[string]interface{})["LastUpdate"])
	fmt.Println("\nFrom json file")
	fmt.Printf("%v\n", string(book3))
	fmt.Printf("%v\n", obj4)
	fmt.Printf("%v\n", obj4.Info.LastUpdate)
	fmt.Println("\nFrom struct array")
	fmt.Printf("%v\n", notes)
	fmt.Printf("%v\n", string(obj5))
	fmt.Println("\nFrom array of object file to struct")
	fmt.Printf("%v\n", arr1)
}