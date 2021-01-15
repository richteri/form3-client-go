package integration

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func shutdown(baseURL string) {
	deleteAccount(baseURL, "aab1f5b8-5334-47df-b334-9568ec59ec32")
	deleteAccount(baseURL, "e36ceed3-db72-4340-a38f-324de1c38565")
	deleteAccount(baseURL, "748396c9-a253-4b1e-8175-a6e34ed7cbdc")
}

func deleteAccount(baseURL string, id string) {
	url := fmt.Sprintf("%s/v1/organisation/accounts/%s?version=0", baseURL, id)
	method := "DELETE"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

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
