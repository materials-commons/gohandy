package gohandy

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type EzClient struct {
	*http.Client
	body io.Reader
}

func NewInsecureClient() *EzClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &EzClient{Client: &http.Client{Transport: tr}}
}

func NewClient() *EzClient {
	return &EzClient{Client: &http.Client{}}
}

func (c *EzClient) JsonGet(url string, out interface{}) (int, error) {
	resp, err := c.Client.Get(url)
	if err != nil {
		return 0, err
	}
	return decodeJsonResponse(resp, out)
}

func (c *EzClient) JsonString(j string) *EzClient {
	c.body = strings.NewReader(j)
	return c
}

func (c *EzClient) Json(j interface{}) *EzClient {
	b, err := json.Marshal(j)
	if err == nil {
		c.body = bytes.NewReader(b)
	}

	return c
}

func (c *EzClient) JsonPost(url string, out interface{}) (int, error) {
	resp, err := c.Client.Post(url, "application/json", c.body)
	if err != nil {
		return 0, err
	}
	return decodeJsonResponse(resp, out)
}

func decodeJsonResponse(resp *http.Response, out interface{}) (int, error) {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		return resp.StatusCode, errors.New(string(body))
	}

	fmt.Println(string(body))
	err := json.Unmarshal(body, out)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
	return resp.StatusCode, nil
}

func (c *EzClient) PostFile(url, filepath, formName string, params map[string]string) (int, error) {
	// Setup body
	body := bytes.NewBufferString("")
	writer := multipart.NewWriter(body)
	defer writer.Close()

	// Create file form
	part, err := writer.CreateFormFile(formName, filepath)
	if err != nil {
		return 0, err
	}

	// Read file into form
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	part.Write(fileContents)

	// Add boundary
	boundary := writer.Boundary()
	closeStr := fmt.Sprintf("\r\n--%s--\r\n", boundary)

	// Add additional params
	if params != nil {
		for key, value := range params {
			writer.WriteField(key, value)
		}
	}

	// Create the request and set headers
	closeBuf := bytes.NewBufferString(closeStr)
	reader := io.MultiReader(body, file, closeBuf)
	req, err := http.NewRequest("POST", url, reader)
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = int64(body.Len()) + int64(closeBuf.Len())

	// Post the file
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	// Return status and any error code
	if resp.StatusCode > 299 {
		return resp.StatusCode, errors.New(string(b))
	}

	return resp.StatusCode, nil
}
