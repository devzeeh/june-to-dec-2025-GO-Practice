package week3 //

import (
	"net/http"
)

// Simple HTTP server
type server struct {
	addr string
}

// ServeHTTP handles HTTP requests
func (s *server) ServeHTTP(W http.ResponseWriter, R *http.Request) {
	switch R.Method {
	case http.MethodGet:
		switch R.URL.Path {
		case "/":
			W.Write([]byte("Hello, Go Web!"))
			return
		case "/User":
			W.Write([]byte("User Page"))
			return
		}
	default:
		W.Write([]byte("Method not allowed"))
		return
	}
}

// NetHttp starts the HTTP server
func NetHttp() {
	s := &server{addr: ":8080"}
	err := http.ListenAndServe(s.addr, s)
	if err != nil {
		panic(err)
	}

}
