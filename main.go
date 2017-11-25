package main

import (
    "log"
    "net/http"
    "os"
//    "github.com/gin-gonic/gin"
//    "strings"
    "fmt"
)


func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.error("$PORT must be set")
		port := 8080
	}

    http.HandleFunc("/", sayhelloName) // set router
    err := http.ListenAndServe(":"+port, nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

// gin-gonic sounds cool
//	router := gin.New()
//	router.Use(gin.Logger())
//	router.LoadHTMLGlob("templates/*.tmpl.html")
//	router.Static("/static", "static")
//
//	router.GET("/", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.tmpl.html", nil)
//	})
//
//	router.Run(":" + port)
}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
//    r.ParseForm()  // parse arguments, you have to call this by yourself
//    fmt.Println(r.Form)  // print form information in server side
//    fmt.Println("path", r.URL.Path)
//    fmt.Println("scheme", r.URL.Scheme)
//    fmt.Println(r.Form["url_long"])
//    for k, v := range r.Form {
//        fmt.Println("key:", k)
//        fmt.Println("val:", strings.Join(v, ""))
//    }
    fmt.Fprintf(w, "Hello & go bye") // send data to client side
}
