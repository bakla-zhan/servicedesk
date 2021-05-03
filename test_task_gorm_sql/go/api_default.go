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

	database.CreateItem(req)

	w.WriteHeader(http.StatusCreated)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	idStr := vars["requestId"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID parse error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	database.DeleteItem(id)

	w.WriteHeader(http.StatusNoContent)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	idStr := vars["requestId"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID parse error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := database.GetItem(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ListRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	res := database.GetItemsList()

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

	database.UpdateItem(req)

	w.WriteHeader(http.StatusOK)
}
