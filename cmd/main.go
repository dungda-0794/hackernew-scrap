package main

import (
	"hackernew-scrap/infrastructure"
)

func main() {
	_, err := infrastructure.ConnectDatabase()
	if err != nil {
		infrastructure.ErrLog.Fatal(err)
		return
	}
}
