package main

import (
	"Chinglish/core"
	"Chinglish/global"
	"Chinglish/initialize"
)
func main() {
	initialize.Redis()
	initialize.Mysql()
	defer global.GVA_DB.Close()

	core.RunWindowsServer()
}