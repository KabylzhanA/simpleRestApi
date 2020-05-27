package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simpleRestApi/model"
)

func (a *App) loadCourse(w http.ResponseWriter, r *http.Request) {
	body,error := ioutil.ReadAll(r.Body)
	if error!=nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,error)
		return
	}
	result, err := model.LoadCourse(string(body),a.DB)
	fmt.Println(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(result)
	return
}
