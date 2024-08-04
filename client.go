package pesto

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:19090"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// NewClient -
func NewClient(host, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	//ar, err := c.SignIn()
	//if err != nil {
	//	return nil, err
	//}

	//c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json; charset=UTF-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	pestoApiResponseHttpStatusCode := res.StatusCode

	switch pestoApiResponseHttpStatusCode {
	case http.StatusOK:
	 fmt.Println("CLIENT.GO [doRequest] - Pesto API Response HTTP CODE IS: http.StatusOK")
	case http.StatusFound:
	 fmt.Println("CLIENT.GO [doRequest] - Pesto API Response HTTP CODE IS: http.StatusFound")
	case http.StatusCreated:
	 fmt.Println("CLIENT.GO [doRequest] - Pesto API Response HTTP CODE IS: http.StatusCreated")
	case http.StatusNoContent:
	 fmt.Println("CLIENT.GO [doRequest] - Pesto API Response HTTP CODE IS: http.StatusNoContent")
	default:
		return nil, fmt.Errorf("CLIENT.GO [doRequest] - ERROR - Pesto API Response HTTP status is: %d, body: %s", res.StatusCode, body)
	}
	return body, err
}
