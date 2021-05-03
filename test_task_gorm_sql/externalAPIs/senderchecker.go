package senderchecker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type serviceResponce struct {
	Source string `json:"source,omitempty"`
	Email  string `json:"email,omitempty"`
	Local  string `json:"local,omitempty"`
	Domain string `json:"domain,omitempty"`
	Type   string `json:"type,omitempty"`
	Qc     int    `json:"qc,omitempty"`
}

var errExpected = errors.New("DISPOSABLE email address")

func SenderChecker(email string) error {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatal("API_KEY env is not set")
	}
	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		log.Fatal("SECRET_KEY env is not set")
	}

	body := bytes.NewBuffer([]byte(fmt.Sprint(`[ "`, email, `" ]`)))

	cli := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://cleaner.dadata.ru/api/v1/clean/email",
		body,
	)
	if err != nil {
		log.Println("dadata create request error:", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", fmt.Sprint("Token ", apiKey))
	req.Header.Add("X-Secret", secretKey)

	resp, err := cli.Do(req)
	if err != nil {
		log.Println("dadata get response error:", err)
		return err
	}
	if (resp.StatusCode / 100) != 2 {
		return errors.New("unable to sent request to dadata")
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("response body read error:", err)
		return err
	}

	var results []serviceResponce

	err = json.Unmarshal(payload, &results)
	if err != nil {
		log.Println("results unmarshall error:", err)
		return err
	}

	if results[0].Qc == 3 {
		return errExpected
	}

	return nil
}
