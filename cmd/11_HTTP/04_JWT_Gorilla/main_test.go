package main

import (
	"encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type JwtToken struct {
	Token string `json:"token"`
}

func TestHandleLoginReq(t *testing.T) {
	resp, errPOST := http.PostForm(
		"http://localhost:"+strconv.Itoa(HTTPPort)+URLLogin,

		url.Values{
			"username": {"x"},
			"password": {"y"},
		},
	)
	require.NoError(t,
		errPOST,
		fmt.Sprintf("POST error:%s", errPOST.Error()),
	)

	defer resp.Body.Close()

	var jwtToken JwtToken

	require.NoError(t,
		json.NewDecoder(resp.Body).Decode(&jwtToken),
	)

	fmt.Println("token: ", jwtToken.Token)
}

func TestHandleRestrictedRoutes(t *testing.T) {
	client := &http.Client{}

	request, errCreateRequest := http.NewRequest(
		"GET",
		"http://localhost:"+strconv.Itoa(HTTPPort)+URLRestricted,
		nil,
	)
	require.NoError(t, errCreateRequest)

	token := "xxx"

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Accept", "application/json")

	response, errPerformRequest := client.Do(request)
	require.NoError(t, errPerformRequest)

	fmt.Println("response: ", response)
}
