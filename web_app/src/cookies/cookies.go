package cookies

import (
	"net/http"
	"web_app/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func SaveAuthData(w http.ResponseWriter, id, token string) error {
	data := map[string]string{
		"id": id,
		"token": token,
	}

	encryptedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name: "data",
		Value: encryptedData,
		Path: "/",
		HttpOnly: true,
	})

	return nil
}

func ReadCookies(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}