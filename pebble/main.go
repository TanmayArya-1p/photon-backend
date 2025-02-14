package pebble

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

var (
	serverURL = os.Getenv("PEBBLE_SERVER_URL")
)

func createFormData(fields map[string]string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range fields {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}

func Login(uid, pwd string) (string, error) {
	fields := map[string]string{
		"uid":      uid,
		"password": pwd,
	}
	body, contentType, err := createFormData(fields)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(serverURL+"/user/login", contentType, body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func Register(username, password string) (string, error) {
	fields := map[string]string{
		"username": username,
		"password": password,
	}
	body, contentType, err := createFormData(fields)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(serverURL+"/user/create", contentType, body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func CreateSession(uid string, secret string, sesKey string) (string, error) {
	fields := map[string]string{
		"key": sesKey,
	}
	body, contentType, err := createFormData(fields)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", serverURL+"/session", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("uid", uid)
	req.Header.Set("secret", secret)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func JoinSession(uid string, secret string, sesID, sesKey string) (string, error) {
	fields := map[string]string{
		"key": sesKey,
	}
	body, contentType, err := createFormData(fields)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("PUT", serverURL+"/session/join?sid="+sesID, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("uid", uid)
	req.Header.Set("secret", secret)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func LeaveSession(uid string, secret string, sid string) (string, error) {
	req, err := http.NewRequest("DELETE", serverURL+"/session/leave?sid="+sid, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("uid", uid)
	req.Header.Set("secret", secret)
	req.Header.Set("Content-Type", "multipart/form-data")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func SessionMetadata(uid string, secret string, sid string) (string, error) {
	req, err := http.NewRequest("PUT", serverURL+"/session?sid="+sid, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("uid", uid)
	req.Header.Set("secret", secret)
	req.Header.Set("Content-Type", "multipart/form-data")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
