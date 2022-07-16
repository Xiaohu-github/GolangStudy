/*
*结构体的应用
*json 编解码
*orm 映射
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   string   `json:"yaer"`
	Price  int      `json:"rmb"`
	Actors []string `json:"Actors"`
}

func main() {
	move := Movie{"喜剧之王", "2000", 20, []string{"张柏芝", "周星驰"}}

	//将结构体转换为json，编码
	jsonStr, err := json.Marshal(move)
	if err != nil {
		fmt.Println("json Marshal err:", err)
		return
	}

	fmt.Printf("jsonStr: %s\n", jsonStr)

	//将json转换为结构体,解码
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json Unmarshal error", err)
	}

	fmt.Printf("Movie: %v\n", myMovie)
}
