/*
File     : FileShare.go
Author   : feimyy
E-mail   : feimyy@hotmail.com
Desc     : a sample FileShare server
*/

package main

import (
	"net/http"
	"os"
	"strings"
)

func Share(path string, port string) {

	handle := http.FileServer(http.Dir(path))

	err := http.ListenAndServe(":"+port, handle)
	if err != nil {
		println("ListenAddress :", err.Error())
	}
}
func main() {

	port := "80" // Default port: 80
	path := "."  // Default path is current dir

	if len(os.Args) > 1 && len(os.Args) <= 3 {
		path = strings.Join(os.Args[1:2], "")
		port = strings.Join(os.Args[2:3], "")
	} else {
		if len(os.Args) > 3 {
			println("Args Error")
			return
		}
	}
	println("FileServer Directory at :", path)
	println("Listening on port ", port)
	Share(path, port)

}
