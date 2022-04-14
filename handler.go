package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		nil,
	)
}

func saveMsg(c *gin.Context) {

	text := c.PostForm("data")

	for {

		key := rand.Intn(9999-1000) + 1000

		_, ok := list[key]

		if ok == false {
			list[key] = data{key, text}
			fmt.Println("Created key: ", key)
		}

		c.Redirect(http.StatusFound, strconv.Itoa(key))
		break
	}
}

func getMsg(c *gin.Context) {
	sid := c.Param("id")

	id, _ := strconv.Atoi(sid)

	data, ok := list[id]

	if ok == true {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"text": data.content,
			},
		)
	} else {
		c.HTML(
			http.StatusNotFound,
			"notfound.html",
			nil,
		)
	}
}
