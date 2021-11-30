package req

import (
	advent_of_go "advent-of-go"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeRequest(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2020/day/%v/input", day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", advent_of_go.Session)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}
