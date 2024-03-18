package oauthGoogle

import (
	"authservice/app/config"
	"authservice/internal"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOauth struct {
	oauthConfig *oauth2.Config
}

type GoogleOauthToken struct {
	Access_token string
	Id_token     string
}

type GoogleInterface interface {
	GetAuthURL() string
	GetGoogleOauthToken(code string) (*GoogleOauthToken, error)
	GetGoogleUser(access_token string, id_token string) (*internal.User, error)
}

func New() GoogleInterface {
	return &GoogleOauth{
		oauthConfig: &oauth2.Config{
			RedirectURL:  config.GOOGLE_URL,
			ClientID:     config.CLIENT_ID,
			ClientSecret: config.CLIENT_SECRET,
			Scopes:       config.SCOPES,
			Endpoint:     google.Endpoint,
		},
	}
}

// GetAuthURL implements GoogleInterface.
func (google *GoogleOauth) GetAuthURL() string {
	return google.oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// GetGoogleOauthToken implements GoogleInterface.
func (google *GoogleOauth) GetGoogleOauthToken(code string) (*GoogleOauthToken, error) {

	const token = "https://oauth2.googleapis.com/token"
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", google.oauthConfig.ClientID)
	values.Add("client_secret", google.oauthConfig.ClientSecret)
	values.Add("redirect_uri", google.oauthConfig.RedirectURL)

	query := values.Encode()

	req, err := http.NewRequest("POST", token, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleOauthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleOauthTokenRes); err != nil {
		return nil, err
	}

	access_token, ok1 := GoogleOauthTokenRes["access_token"].(string)
	id_token, ok2 := GoogleOauthTokenRes["id_token"].(string)
	if !ok1 || !ok2 {
		return nil, errors.New("invalid token data")
	}

	tokenBody := &GoogleOauthToken{
		Access_token: access_token,
		Id_token:     id_token,
	}

	return tokenBody, nil
}

// GetGoogleUser implements GoogleInterface.
func (google *GoogleOauth) GetGoogleUser(access_token string, id_token string) (*internal.User, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", access_token)

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", id_token))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return nil, err
	}

	userBody := &internal.User{
		VerifiedEmail:         true,
		RegistrationType: "Google",
	}

	if email, ok := GoogleUserRes["email"].(string); ok {
		userBody.Email = email
	}

	if name, ok := GoogleUserRes["name"].(string); ok {
		userBody.FullName = name
	}

	if picture, ok := GoogleUserRes["picture"].(string); ok {
		userBody.PhotoProfile = picture
	}

	return userBody, nil
}
