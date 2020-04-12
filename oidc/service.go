package oidc

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Service struct {
}

func (s *Service) Authenticate(idToken string) (string, error) {
	u, err := url.Parse("https://www.googleapis.com/oauth2/v1/tokeninfo")
	if err != nil {
		panic(err)
	}
	vs := url.Values{}
	vs.Add("id_token", idToken)
	u.RawQuery = vs.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var ti TokenInfo
		if err := json.NewDecoder(res.Body).Decode(&ti); err != nil {
			return "", err
		}
		return ti.Email, nil
	}
	io.Copy(os.Stdout, res.Body)
	return "", InvalidToken
}
