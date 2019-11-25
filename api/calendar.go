package api

/*
	@file	calendat.go
	@author	LoperLee
	@date	2019-11-20
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	@author	LoperLee
	@brief	kakao i bot request json
*/
type KakaoAction struct {
	Intent struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"intent"`
	UserRequest struct {
		Timezone string `json:"timezone"`
		Params   struct {
			IgnoreMe string `json:"ignoreMe"`
		} `json:"params"`
		Block struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"block"`
		Utterance string      `json:"utterance"`
		Lang      interface{} `json:"lang"`
		User      struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			Properties struct {
			} `json:"properties"`
		} `json:"user"`
	} `json:"userRequest"`
	Bot struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"bot"`
	Action struct {
		Name        string      `json:"name"`
		ClientExtra interface{} `json:"clientExtra"`
		Params      struct {
			Date string `json:"date"`
		} `json:"params"`
		ID           string `json:"id"`
		DetailParams struct {
			Date struct {
				Origin    string `json:"origin"`
				Value     string `json:"value"`
				GroupName string `json:"groupName"`
			} `json:"date"`
		} `json:"detailParams"`
	} `json:"action"`
}

func ArticlesSelectAll(c *gin.Context) {
	var json KakaoAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// params := c.QueryParams()
	// ret, err := model.GetArticles(c.Request().Context())
	// if err != nil {
	// 	return c.NoContent(http.StatusBadGateway)
	// }
	// return c.JSON(http.StatusOK, ret)
}
