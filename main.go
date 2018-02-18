package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	medium "github.com/medium/medium-sdk-go"
	yaml "gopkg.in/yaml.v2"
)

// CheckErr to handle errors
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type apikeys struct {
	Token string
}

func (a *apikeys) getAPIKeys(filename string) *apikeys {
	yamlFile, err := ioutil.ReadFile(filename)
	CheckErr(err)
	err = yaml.Unmarshal(yamlFile, a)
	CheckErr(err)
	return a
}

// GetPubs makes a GET request to the API to Get Users Publications
func GetPubs(uid string) (b []byte) {
	// var a apikeys
	// var token string
	// a.getAPIKeys(".secrets.yaml")
	// token = a.Token

	//log.Printf("Requesting %v", uri)
	//Authorization: Bearer 181d415f34379af07b2c11d144dfbe35d
	// Make a GET request to https://medium.com/@ev/latest with Accept: application/json header.
	//You can pass count parameter to limit the posts. Much easier to parse than RSS.
	//keys := "Bearer " + token
	c := &http.Client{}
	r, err := http.NewRequest("GET", "https://medium.com/@"+uid+"/latest", nil)
	CheckErr(err)
	r.Header.Add("Accept", "application/json")
	resp, err := c.Do(r)
	CheckErr(err)
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	CheckErr(err)
	return b
}

// func getAssetInfo(assetID string) (aI AssetInfo) {
// 	//log.Println(assetID)
// 	uri := "/workbenches/assets/" + assetID + "/info"
// 	body := APIReq(uri)
// 	aI = AssetInfo{}
// 	err := json.Unmarshal(body, &aI)
// 	CheckErr(err)

// 	//log.Println(body)
// 	return aI
// }

func prettyPrintJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	//log.Println("Medium Blog Bridge Coming Online ....")
	var k apikeys
	k.getAPIKeys(".secrets.yaml")
	token := k.Token

	m := medium.NewClientWithAccessToken(token)

	u, err := m.GetUser("")

	CheckErr(err)

	//log.Println(u.Username)

	p := GetPubs(u.Username)

	log.Println(string(p))

	// r := clientRequest{
	// 	method: "GET",
	// 	path:   fmt.Sprintf("/v1/users/%s/publications", userID),
	// }
	// p := &Publications{}
	// err := m.request(r, p)
	// return p, err

	// // Construct the request
	// path := fmt.Sprintf("https://api.medium.com/v1/users/%s/publications", u.ID)
	// req, err := http.NewRequest("GET", path, bytes.NewReader(body))
	// if err != nil {
	// 	return Error{fmt.Sprintf("Could not create request: %s", err), defaultCode}
	// }

	// req.Header.Add("Content-Type", ct)
	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Accept-Charset", "utf-8")
	// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.AccessToken))

	// // Create the HTTP client
	// client := &http.Client{
	// 	Transport: m.Transport,
	// 	Timeout:   m.Timeout,
	// }

	// // Make the request
	// res, err := client.Do(req)
	// if err != nil {
	// 	return Error{fmt.Sprintf("Failed to make request: %s", err), defaultCode}
	// }
	// defer res.Body.Close()

	// // Parse the response
	// c, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return Error{fmt.Sprintf("Could not read response: %s", err), defaultCode}
	// }

	// var env envelope
	// if err := json.Unmarshal(c, &env); err != nil {
	// 	return Error{fmt.Sprintf("Could not parse response: %s", err), defaultCode}
	// }

	// if http.StatusOK <= res.StatusCode && res.StatusCode < http.StatusMultipleChoices {
	// 	if env.Data != nil {
	// 		c, _ = json.Marshal(env.Data)
	// 	}
	// 	return json.Unmarshal(c, &result)
	// }
	// e := env.Errors[0]
	// return Error{e.Message, e.Code}

}
