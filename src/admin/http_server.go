package admin

import (
	"net/http"
)

type HttpServer struct {
	homeDir string
	port    string
}

/*
  homeDir is the directory that is the root of the admin site.
  port should be a string that looks like ":8080" or whatever port to serve on.
*/
func NewHttpServer(homeDir, port string) *HttpServer {
	return &HttpServer{homeDir: homeDir, port: port}
}

func (self *HttpServer) ListenAndServe() {
	http.ListenAndServe(self.port, http.FileServer(http.Dir(self.homeDir)))
}
