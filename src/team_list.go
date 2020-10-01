package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	FlagSubIP   string   `json:"flag_sub_ip"`
	FlagSubPORT int      `json:"flag_sub_port"`
	Regex       string   `json:"regex"`
	Teams       []string `json:"teams"`
}

func team_list(data []byte) ([]string, string, int, string) {

	// change variable result according to the structure of the team json

	var result Config
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
	}

	return result.Teams, result.FlagSubIP, result.FlagSubPORT, result.Regex

}
