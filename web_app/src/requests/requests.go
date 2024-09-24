package requests

import (
	"io"
	"net/http"
	"web_app/src/cookies"
)

// Insere o token de acesso na requisição
func RequestAuthenticated(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.ReadCookies(r)
	request.Header.Add("Authorization", "Bearer " + cookie["token"])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}