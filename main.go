package main

import "kw-calendar-bot/route"

/*
	@file	main.go
	@author	LoperLee
	@date	2019-11-20
*/

func main() {
	route.StartAPIServer().Run(":25623")
}
