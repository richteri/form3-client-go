package integration

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Setup test data exported from the Postman collection.
func setup(baseURL string) {
	createAccount1(baseURL)
	createAccount2(baseURL)
	createAccount3(baseURL)
}

func createAccount1(baseURL string) {
	url := fmt.Sprintf("%s/v1/organisation/accounts", baseURL)
	method := "POST"

	payload := strings.NewReader(`{
  "data": {
    "id": "aab1f5b8-5334-47df-b334-9568ec59ec32",
    "organisation_id": "dabe0971-b3a8-465f-aa6c-30bf8ba98c89",
    "type": "accounts",
    "attributes": {
       "country": "GB",
        "base_currency": "GBP",
        "bank_id": "400302",
        "bank_id_code": "GBDSC",
        "account_number": "10000004",
        "customer_id": "234",
        "iban": "GB28NWBK40030212764204",
        "bic": "NWBKGB42",
        "account_classification": "Personal"
    }
  }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func createAccount2(baseURL string) {
	url := fmt.Sprintf("%s/v1/organisation/accounts", baseURL)
	method := "POST"

	payload := strings.NewReader(`{
  "data": {
    "id": "e36ceed3-db72-4340-a38f-324de1c38565",
    "organisation_id": "dabe0971-b3a8-465f-aa6c-30bf8ba98c89",
    "type": "accounts",
    "attributes": {
       "country": "GB",
        "base_currency": "GBP",
        "bank_id": "400302",
        "bank_id_code": "GBDSC",
        "customer_id": "234",
        "bic": "NWBKGB42",
        "name": [
            "Samantha Holder"
        ],
        "alternative_names": [
            "Sam Holder"
        ],
        "account_classification": "Personal",
        "joint_account": false,
        "account_matching_opt_out": false,
        "secondary_identification": "A1B2C3D4"
    }
  }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func createAccount3(baseURL string) {
	url := fmt.Sprintf("%s/v1/organisation/accounts", baseURL)
	method := "POST"

	payload := strings.NewReader(`{
  "data": {
    "id": "748396c9-a253-4b1e-8175-a6e34ed7cbdc",
    "organisation_id": "dabe0971-b3a8-465f-aa6c-30bf8ba98c89",
    "type": "accounts",
    "attributes": {
    	"country": "GB",
    	"base_currency": "GBP",
    	"name": [
    		"James Bond"
    	],
        "bank_id": "400305",
        "bank_id_code": "GBDSC",
        "bic": "LHVBEE22",
        "account_classification": "Personal",
        "private_identification": {
            "birth_date": "1920-11-11",
            "birth_country": "GB",
            "identification": "MI6008",
            "address": [
                "11 Up and Down Street"
            ],
            "country": "GB",
            "city": "London"
        }
    },
    "relationships": {
        "master_account": {
            "data": [
                {
                    "type": "accounts",
                    "id": "b7954905-bcff-43bd-b64f-7e622eba6cb8"
                }
            ]
        }
    }
  }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
