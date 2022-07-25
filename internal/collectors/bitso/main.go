package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	sig := getSignature()
	res, err := rest(sig)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	for _, b := range res.Payload.Balances {
		fmt.Println(b.Currency + " -> " + b.Total)
	}
}

type response struct {
	Payload struct {
		Balances []struct {
			Currency string `json:"currency"`
			Total    string `json:"total"`
		} `json:"balances"`
	} `json:"payload"`
}

func getSignature() string {
	secret := ""
	data := "1655266896GET/v3/balance/"

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func rest(sig string) (*response, error) {
	r, err := http.NewRequest("GET", "https://api.bitso.com/v3/balance", nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Authorization", "Bitso qBZOUNWGXC:1655266896:"+sig)

	c := http.Client{}

	c.Get("https://api.bitso.com/v3/balance/")

	res, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	respon := response{}

	// fmt.Println(res.StatusCode)

	// bd, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(bd))

	err = json.NewDecoder(res.Body).Decode(&respon)
	if err != nil {
		return nil, err
	}
	return &respon, nil
}
