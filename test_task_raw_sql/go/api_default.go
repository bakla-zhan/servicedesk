/*
 * ServiceDesk
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	senderchecker "servicedesk/externalAPIs"
)

var database Repository

func DBinit() {
	db := DBConnect()
	database = db
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	req := &Request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Println("unparse json error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// проверка email автора
	err = senderchecker.SenderChecker(req.Email)
	if err != nil {
		log.Println("senderchecker:", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = database.CreateItem(req)
	if err != nil {
		log.Println("DB create item error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	idStr := vars["requestId"]

	err := database.DeleteItem(idStr)
	if err != nil {
		log.Println("DB delete item error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	idStr := vars["requestId"]

	res, err := database.GetItem(idStr)
	if err != nil && err != ErrNotFound {
		log.Println("get item by ID error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err == ErrNotFound {
		log.Println("item not found:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ListRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	res, err := database.GetItemsList()
	if err != nil {
		log.Println("get items error:", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func UpdateRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	idStr := vars["requestId"]
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID parse error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &Request{}

	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Println("unparse json error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Id = int64(ID)

	err = database.UpdateItem(req)
	if err != nil {
		log.Println("DB update item error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
