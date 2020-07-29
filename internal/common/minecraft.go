package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type profile struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Properties []texture `json:"properties"`
}

type uuid struct {
	ID   string `json:"id"`
	Name string `json:"username"`
}

type texture struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}

// GetPlayerName gets a Minecraft player name of a player UUID. It will return an empty string if
// nothing was found.
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
	var profile profile

	err = json.Unmarshal(body, &profile)

	if err != nil {
		log.Printf("Failed to parse JSON for \"%s\", because\n"+err.Error(), playerID)
		return ""
	}

	return profile.Name
}

// GetPlayerID will get the UUID of a given Minecraft player name.
// An empty string will be returned if nothing was found.
func GetPlayerID(playerName string) string {
	names := []string{playerName}
	serialized, _ := json.Marshal(names)
	resp, err := http.Post(
		"https://api.mojang.com/profiles/minecraft",
		"application/json",
		bytes.NewBuffer(serialized),
	)

	if err != nil {
		log.Printf("Failed to get \"%s\" player ID because\n%s\n", playerName, err.Error())
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result []uuid

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Printf("Failed to parse JSON for \"%s\", because\n"+err.Error(), playerName)
		return ""
	}

	if len(result) > 0 {
		return result[0].ID
	}
	return ""
}
