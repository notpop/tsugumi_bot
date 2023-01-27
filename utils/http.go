package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	url        string
	headers    map[string]string
	httpClient *http.Client
}

func New(url string, headers map[string]string) *Client {
	client := &Client{url, headers, &http.Client{}}
	return client
}

func (client *Client) DoRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseUrl, err := url.Parse(client.url)
	if err != nil {
		log.Println("BaseUrl Parse Error: ", err)
		return nil, err
	}

	apiUrl, err := url.Parse(urlPath)
	if err != nil {
		log.Println("Url Parse Error: ", err)
		return nil, err
	}

	endpoint := baseUrl.ResolveReference(apiUrl).String()

	log.Printf("action=doRequest endpoint=%s", endpoint)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Http NewRequest Error: ", err)
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range client.headers {
		req.Header.Add(key, value)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		log.Println("HttpClient Do Error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil ReadAll Error: ", err)
		return nil, err
	}

	return body, nil
}
