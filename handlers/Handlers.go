package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-firstapi/dataloaders"
	"golang-firstapi/models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
func Handler(w http.ResponseWriter, r *http.Request) {
	page := models.Page{Id: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	users := dataloaders.LoadUsers()
	interests := dataloaders.LoadInterests()
	interestMappings := dataloaders.LoadInterestMappings()
	var newUsers []models.User
	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.Id == interestMapping.UserId {
				for _, interest := range interests {
					if interestMapping.InterestId == interest.Id {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
		fmt.Println(user.FirstName)
	}
	viewModel := models.UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
}
