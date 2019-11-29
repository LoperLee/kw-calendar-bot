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
type kakaoSkill struct {
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
			Date int `json:"date"`
		} `json:"params"`
		ID           string `json:"id"`
		DetailParams struct {
			Date struct {
				Origin    int    `json:"origin"`
				Value     int    `json:"value"`
				GroupName string `json:"groupName"`
			} `json:"date"`
		} `json:"detailParams"`
	} `json:"action"`
}

func GetCalendar(con *gin.Context) {
	var json kakaoSkill
	if err := con.ShouldBindJSON(&json); err != nil {
		con.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	con.JSON(http.StatusOK, gin.H{"message": "Hello!", "date": json.Action.Params.Date})
	return
}
