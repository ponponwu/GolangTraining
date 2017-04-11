package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Hoppo float64 `json:"h_oppo"`
	Aoppo float64 `json:"a_oppo"`
	Doppo float64 `json:"d_oppo"`
	Bytes []byte
}

func main() {
	bs, _ := json.Marshal({1,2,3})
	fmt.Println(string(bs))
	m := Message{}
	m.Bytes = bs
	fmt.Println(m)
}
