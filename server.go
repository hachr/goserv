package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    )

    func main() {
    	arg := os.Args
	length := len(arg)
	port := "8080"
	if length == 2 {
		port = arg[1]
	}
        fmt.Println("Serving files in the current directory on port", port)
	    http.Handle("/", http.FileServer(http.Dir(".")))
	        err := http.ListenAndServe(":" + port, nil)
	    if err != nil {
	            log.Fatal("ListenAndServe: ", err)
	        }
	}
