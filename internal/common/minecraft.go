package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Profile struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Properties []Texture `json:"properties"`
}

type UUID struct {
	ID   string `json:"id"`
	Name string `json:"username"`
}

type Texture struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}

// an empty string will return if an error occurred or there was no result
func GetPlayerName(playerID string) string {
	resp, err := http.Get(
		"https://sessionserver.mojang.com/session/minecraft/profile/" + playerID,
	)

	if err != nil {
		log.Printf("Failed to get player name for \"%s\", because\n"+err.Error(), playerID)
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var profile Profile

	err = json.Unmarshal(body, &profile)

	if err != nil {
		log.Printf("Failed to parse JSON for \"%s\", because\n"+err.Error(), playerID)
		return ""
	}

	return profile.Name
}

func GetPlayerID(playerName string) string {
	names := []string{playerName}
	serialized, _ := json.Marshal(names)
	resp, err := http.Post(
		"https://api.mojang.com/profiles/minecraft",
		"application/json",
		bytes.NewBuffer(serialized),
	)

	if err != nil {
		log.Printf("Failed to get \"%s\" player ID because\n"+err.Error(), playerName)
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result []UUID

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Printf("Failed to parse JSON for \"%s\", because\n"+err.Error(), playerName)
		return ""
	}

	if len(result) > 0 {
		return result[0].ID
	} else {
		return ""
	}
}
