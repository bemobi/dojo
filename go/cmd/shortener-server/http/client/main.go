package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response []struct {
	TestPlan struct {
		Slug          string `json:"slug"`
		Description   string `json:"description"`
		CheckInterval int    `json:"checkInterval"`
	} `json:"testPlan"`
	Score     int    `json:"score"`
	MaxScore  int    `json:"maxScore"`
	Timestamp string `json:"timestamp"`
}

func main() {
	req, err := http.NewRequest("GET", "http://okay.bemobi.com.br/api/overview/test-plan", nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("user", "123qwe")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data Response
	json.NewDecoder(res.Body).Decode(&data)

	dataFmt, _ := json.MarshalIndent(data, "", "\t")

	fmt.Println(string(dataFmt))
}
