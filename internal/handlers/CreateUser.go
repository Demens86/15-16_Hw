package handlers

import (
	"encoding/json"
	"net/http"
	"prac/models"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := new(models.UserRequest)
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println(userRequest)

	ctx := r.Context()
	user, err := h.useCase.CreateUser(ctx, userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(user)

	//handler
	//userData, err := json.Marshal(user)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//_, err = w.Write(userData)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
}
