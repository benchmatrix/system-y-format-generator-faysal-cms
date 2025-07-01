package main

import (
	"fmt"

	//"system-y-format-generator-faysal-cms/DbConnector"
	routing "system-y-format-generator-faysal-cms/Routing"
)

func main() {
	fmt.Println("Format Generator Faysal Running")

	routing.StartServer()
}
