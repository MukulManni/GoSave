package main

func intializeRoutes() {
	r.GET("/", homePage)

	r.POST("/post", saveMsg)

	r.GET("/:id", getMsg)
}
