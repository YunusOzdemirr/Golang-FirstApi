package dataloaders

import (
	"encoding/json"

	"golang-firstapi/models"

	util "golang-firstapi/utils"
)

func LoadUsers() []models.User {
	bytes, _ := util.ReadFile("../json/users.json")
	var users []models.User
	json.Unmarshal([]byte(bytes), &users)
	return users
}

func LoadInterests() []models.Interest {
	bytes, _ := util.ReadFile("../json/interest.json")
	var interests []models.Interest
	json.Unmarshal([]byte(bytes), &interests)
	return interests
}

func LoadInterestMappings() []models.InterestMapping {
	bytes, _ := util.ReadFile("../json/userInterestMappings.json")
	var interestMapping []models.InterestMapping
	json.Unmarshal([]byte(bytes), &interestMapping)
	return interestMapping
}
