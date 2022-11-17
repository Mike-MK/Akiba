package mpesa

import (
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
)

func GetMpesaAuthToken() (string,error){
	key := "DBKT6n6IXsxGt1vwueGB1pygGrlII1PX"
    secret :=  "jGJA828Jt97lP9k0"

	ks := key+":"+secret
	b64Key := b64.StdEncoding.EncodeToString([]byte(ks))
	// fmt.Println(b64Key)

	client := &http.Client{}
	auth_url := "https://sandbox.safaricom.co.ke/oauth/v1/generate"

	req, err := http.NewRequest("GET", auth_url, nil)
	if err != nil {
		return "",err
	}

	authHeader := "Basic "+ b64Key
	req.Header.Add("Authorization", authHeader)
	q := req.URL.Query()
    q.Add("grant_type", "client_credentials")
    req.URL.RawQuery = q.Encode()

	
	resp, err := client.Do(req)
	if err != nil {
		return "",err
	}

	var body map[string]string
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return "",err
	}

	// fmt.Println(body)
	defer resp.Body.Close()


	return body["access_token"],nil
}