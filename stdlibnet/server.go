package stdlibnet

import (
	"fmt"
	"net"
	"net/http"
)

type Server struct {
	Port           string
	ConnectionType string
}

func (s Server) Run() {
	fmt.Println("listening on port 8080")
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Errorf("failed to create server")
	}

	ln.Accept()
}

func (s Server) DoSomething() {
	http.Handle()
}
