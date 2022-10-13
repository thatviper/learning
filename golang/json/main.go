package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type user struct {
	name string
	job  string
}

type ReqresResponse struct {
	Id        string
	CreatedAt string
}

func main() {

	var rithish user = user{
		name: "Rithish",
		job:  "CEO",
	}

	fmt.Println(fmt.Sprintf(`%v is a %v`, rithish.name, rithish.job))

	userJson, err := json.Marshal(rithish)

	response, err := http.Post("https://reqres.in/api/users", "application/json", bytes.NewBuffer(userJson))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	var resBody ReqresResponse
	err = json.Unmarshal(responseData, &resBody)
	fmt.Println(resBody)

}
