package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/decengame/xp-token-tatum/model"
	"github.com/joho/godotenv"
)

// TATUM-API-KEY the TATUM API KEY
var TATUM_API_KEY string

var (
	BASE_ENDPOINT       string
	GAME_WALLET_ADDRESS string
	PRIVATE_KEY         string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured reading env file. Err: %s", err)
	}
	TATUM_API_KEY = os.Getenv("TATUM_API_KEY")
	BASE_ENDPOINT = os.Getenv("BASE_ENDPOINT")
	GAME_WALLET_ADDRESS = os.Getenv("GAME_WALLET_ADDRESS")
	PRIVATE_KEY = os.Getenv("PRIVATE_KEY")
}

func main() {
	var err error
	nonce := 9
	reqUrl := BASE_ENDPOINT + "/blockchain/token/deploy"
	data := []byte(`{
    "chain": "MATIC",
    "symbol": "DECENXP` + strconv.Itoa(nonce) + `",
    "name": "DecenXPTokenTest` + strconv.Itoa(nonce) + `",
    "totalCap": "1",
    "supply": "1",
    "digits": 1,
    "address": "` + GAME_WALLET_ADDRESS + `",
    "fromPrivateKey": "0x` + PRIVATE_KEY + `",
    "nonce": ` + strconv.Itoa(nonce) + `
  }`)
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", TATUM_API_KEY)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	if res.StatusCode != 200 {
		fmt.Printf("Request: %+v\n", req)
		fmt.Println(string(data))
		fmt.Printf("Response: %+v\n", res)
		fmt.Println(string(body))
		log.Fatalln("Error: ", res.StatusCode, res.Status)
	}
	var deployResponse model.DeployResponse
	err = json.Unmarshal(body, &deployResponse)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	reqUrl = BASE_ENDPOINT + "/polygon/transaction/" + deployResponse.TxID
	req, err = http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("x-api-key", TATUM_API_KEY)
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	if res.StatusCode != 200 {
		fmt.Printf("Request: %+v\n", req)
		fmt.Println(string(data))
		fmt.Printf("Response: %+v\n", res)
		fmt.Println(string(body))
		log.Fatalln("Error: ", res.StatusCode, res.Status)
	}
	var txResponse model.TransactionResponse
	err = json.Unmarshal(body, &txResponse)
	if err != nil {
		log.Println(string(body))
		log.Fatalln("Error: ", err.Error())
	}
	if !txResponse.Status {
		log.Fatalln("Error on the deployment. Check the transaction", deployResponse.TxID)
	}
	log.Println("Contract", txResponse.ContractAddress, " successfully deployed within the transaction", deployResponse.TxID)
}
