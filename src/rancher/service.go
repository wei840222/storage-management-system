package rancher

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// GetServiceEndpoint get endpoint by service name
func GetServiceEndpoint(serviceName string) string {
	url := "https://acl.csie.ntut.edu.tw:3000/v3/project/c-6rfmm:p-vtj8t/dnsRecords/storage-manage-system-dev:" + serviceName

	// Create a Bearer string by appending string access token
	const bearer = "Bearer " + "token-9s52k:z5m5cwwxft7blkszf8nw2j5th9mlxgjzlmwswk5pgl2bgv87sdvcwc"

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("GetServiceEndpoint %s", string(body))
	addresse := gjson.Get(string(body), "publicEndpoints.0.addresses.0")
	port := gjson.Get(string(body), "publicEndpoints.0.port")
	log.Print(addresse, port)
	log.Printf("GetServiceEndpoint %s %s", addresse, port)
	return "http://" + addresse.String() + ":" + port.String()
}
