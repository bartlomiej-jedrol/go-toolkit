package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	iLog "github.com/bartlomiej-jedrol/go-toolkit/log"
)

var service = "go-toolkit"

func SendHTTPRequest(enpointURL string, APIKey string, requestBody []byte) ([]byte, error) {
	function := "SendRequest"
	requestDataReader := bytes.NewReader(requestBody)

	URL, err := url.Parse(enpointURL)
	if err != nil {
		iLog.Error("failed to parse URL", nil, err, service, function)
		return nil, err
	}

	header := http.Header{}
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", fmt.Sprintf("Bearer %s", APIKey))

	request := http.Request{
		Method: http.MethodPost,
		URL:    URL,
		Header: header,
		Body:   io.NopCloser(requestDataReader),
	}

	// iLog.Info("request method", request.Method, nil, service, function)
	// iLog.Info("request url", request.URL.String(), nil, service, function)
	// iLog.Info("request headers", request.Header, nil, service, function)
	// iLog.Info("request body", string(requestBody), nil, service, function)

	client := http.Client{}
	response, err := client.Do(&request)
	if err != nil {
		iLog.Error("failed to send request", nil, err, service, function)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		iLog.Error("received non-OK HTTP status", response.Status, nil, service, function)
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		iLog.Error("failed to read response body", nil, err, service, function)
		return nil, err
	}
	return responseBody, nil
}
