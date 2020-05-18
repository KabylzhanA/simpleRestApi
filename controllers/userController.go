package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"simpleRestApi/model"
	"strconv"
)

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {

	newUser := &model.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	if err = json.Unmarshal(body, &newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if err = newUser.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	if user, _ := newUser.GetUser(a.DB); user != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "User Created Before")
		return
	}

	newUser, err = newUser.SaveUser(a.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	json.NewEncoder(w).Encode(newUser)
	return
}
func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := model.GetAllUsers(a.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(allUsers)
	return
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	userDelete, err := model.GetUserById(id, a.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	result, err := userDelete.DeleteUser(a.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(result)
	return
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	userUpdate := &model.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if _, err := model.GetUserById(id, a.DB); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if err = json.Unmarshal(body, &userUpdate); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	if _, err = userUpdate.UpdateUser(id, a.DB); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(userUpdate)
	return
}
