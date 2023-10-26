package v1

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

var (
	ErrOAuthBody = errors.New("unable to read body")
)

type GoogleUserResult struct {
	Id            string
	Email         string
	VerifiedEmail bool
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	Locale        string
}

type AuthHandler struct {
	oauth2Config *oauth2.Config
	oauth2Github *oauth2.Config
}

func NewAuthHandler(oauth2Google *oauth2.Config) *AuthHandler {
	oauth2GithubConfig := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint:     oauth2.Endpoint{},
		RedirectURL:  "",
		Scopes:       nil,
	}

	return &AuthHandler{
		oauth2Config: oauth2Google,
		oauth2Github: oauth2GithubConfig,
	}
}

func (ah *AuthHandler) SignUp(c *fiber.Ctx) error {
	return nil
}

func (ah *AuthHandler) GoogleAuth(c *fiber.Ctx) error {
	url := ah.oauth2Config.AuthCodeURL("rand") //TODO rand string and check in callback
	c.Status(fiber.StatusSeeOther)
	return c.Redirect(url, http.StatusTemporaryRedirect)

}

func (ah *AuthHandler) GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "rand" {
		return c.SendString("invalid state")
	}
	code := c.Query("code")
	token, err := ah.oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return errors.Wrap(err, "oauth exchange")
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return errors.Wrap(err, "get user info from google")
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			//TODO add logger
		}
	}(response.Body)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ErrOAuthBody
	}
	var googleUserRes map[string]interface{}

	if err := json.Unmarshal(data, &googleUserRes); err != nil {
		return errors.Wrap(err, "unmarshal google callback data")
	}

	//userBody := &GoogleUserResult{
	//	Id:            googleUserRes["id"].(string),
	//	Email:         googleUserRes["email"].(string),
	//	VerifiedEmail: googleUserRes["verified_email"].(bool),
	//	Name:          googleUserRes["name"].(string),
	//	GivenName:     googleUserRes["given_name"].(string),
	//	Picture:       googleUserRes["picture"].(string),
	//	Locale:        googleUserRes["locale"].(string),
	//}

	c.Redirect("http://localhost:3000/")
	return nil
}

func (ah *AuthHandler) GithubAuth(c *fiber.Ctx) error {
	return nil
}

func (ah *AuthHandler) GithubCallback(c *fiber.Ctx) error {
	return nil
}
