package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//GET
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://ctf.makemake.cc/task3/challenge", nil)
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("authorization", "Basic YnJ1bm86RTl1aEdLeWlmYTloWXpJblp1VlZGQk8=")
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error making request", err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error making request", err.Error())
		return
	}
	sb := string(body)
	log.Println(sb)

	//APPLY SOLUTION
	var jRes map[string]interface{}
	json.Unmarshal([]byte(sb), &jRes)
	challenge := jRes["challenge"].(map[string]interface{})

	solution := map[string]int{}

	for key, value := range challenge {
		res := int(value.(float64))
		solution[key] = minNumberOfSquares(res)
	}

	//POST SOLUTION
	postBody, _ := json.Marshal(solution)
	req, _ = http.NewRequest("POST", "https://ctf.makemake.cc/task3/submit", bytes.NewBuffer(postBody))
	token := jRes["token"].(string)
	req.Header.Set("token", token)
	res, err = client.Do(req)
	if err != nil {
		log.Println("Error making POST", err.Error())
		return
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error making request", err.Error())
		return
	}
	sb = string(body)
	log.Println(sb)

}
