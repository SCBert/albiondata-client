package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
}

func (p Page) toString() string {
	return toJson(p)
}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func loadItems(idx int) {
	//itemList := make(map[int]string)
	//x int
	pages := getPages()
	fmt.Println(pages.ID)
	/*for x, y := range pages {
		//itemList
		if id = p. {
			fmt.Println(p.toString())
		}
	}*/

}

func getPages() []Page {
	raw, err := ioutil.ReadFile("albion.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Page
	json.Unmarshal(raw, &c)
	return c
}