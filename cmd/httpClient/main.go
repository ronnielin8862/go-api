package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {

	url := "http://47.241.20.47:8090/wallet/getaccountbalance"
	body := "{\n    \"account_identifier\": {\n        \"address\": \"TFU3d1TKMvmXcAHyRGfVToyfNtAfdUmH5g\"\n    }, \n    \"block_identifier\": {\n        \"hash\": \"00000000020e931413765b06f89f534484721d9ffb273413c58e84fad8f91ab5\",\n        \"number\":34509588\n    },\n    \"visible\": true\n}"

	var jsonStr = []byte(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Headers:", resp.Header)
	fmt.Println("Body:", body)
}
