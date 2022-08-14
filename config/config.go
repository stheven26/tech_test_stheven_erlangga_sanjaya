package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tkanos/gonfig"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

type Details struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type Data struct {
	ID                     string    `json:"_id"`
	Index                  uint      `json:"index"`
	Guid                   string    `json:"guid"`
	IsActive               bool      `json:"isActive"`
	Details                []Details `json:"details"`
	Greeting               string    `json:"greeting"`
	FavoriteTransportation string    `json:"favoriteTransportation"`
}

type SampleData struct {
	Data []Data
}

// type SampleData struct {
// 	Data []struct {
// 		ID       string `json:"_id"`
// 		Index    uint   `json:"index"`
// 		Guid     string `json:"guid"`
// 		IsActive bool   `json:"isActive"`
// 		Details  []struct {
// 			Name    string `json:"name"`
// 			Balance int    `json:"balance"`
// 		} `json:"details"`
// 		Greeting               string `json:"greeting"`
// 		FavoriteTransportation string `json:"favoriteTransportation"`
// 	}
// }

func Configuration() Config {
	config := Config{}
	gonfig.GetConf("./config/config.json", &config)
	return config
}

func ParsingJson() SampleData {
	jsonFile, err := os.Open("./config/sample.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened sample.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var data SampleData

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &data)

	return data
}
