package main

import (
	"io"
	"net/http"
        "go.uber.org/zap"
)

func main() {
logger, _ := zap.NewProduction()
defer logger.Sync()
helloHandler := func(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello, world!\n")
    logger.Info("Hello World Function Call")
}

http.HandleFunc("/hello", helloHandler)
http.ListenAndServe(":8080", nil)
}
