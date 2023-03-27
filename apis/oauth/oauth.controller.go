package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"goFiber/main/config"
	"goFiber/main/services/user"
	"goFiber/main/utils"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func GoogleLoginAPI(c *fiber.Ctx) error {
	token := config.GetGoogleToken()
	url := config.GetGoogleLoginURL(token)
	c.Render("login", fiber.Map{
		"title": "google login",
		"url":   url,
	})

	return nil
}

func GoogleLoginCallBack(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := config.GoogleOauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err.Error(),
			"data":    nil,
		})
		return err
	}
	response, err := http.Get(config.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err.Error(),
			"data":    nil,
		})
		return err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err.Error(),
			"data":    nil,
		})
		return err
	}
	fmt.Println(string(contents))
	jsonMap := make(map[string]interface{})
	json.Unmarshal(contents, &jsonMap)

	email := fmt.Sprintf("%v", jsonMap["email"])
	result, err := user.CreateUserService(email, "", "", true)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err.Error(),
			"data":    nil,
		})
		return err
	}
	userByEmail, isUserExist, _ := user.IsUserExistByEmail(email)

	if isUserExist == false {
		return err
	}

	t, _, err := utils.CreateJWTToken(userByEmail)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err.Error(),
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": "true",
		"message": "",
		"data":    result,
		"token":   t,
	})

	return nil
}
