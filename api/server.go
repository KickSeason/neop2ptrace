package api

import (
	"fmt"
	"neop2ptrace/log"
	np "neop2ptrace/nodemap"
	"net/http"
)

var logger = log.NewLogger()

type ApiServer struct {
	host string
	port int
	srv  *http.Server
	nmap *np.NodeMap
}

func NewApiServer(host string, port int, nmap *np.NodeMap) *ApiServer {
	s := ApiServer{
		host: host,
		port: port,
		nmap: nmap,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/map", s.handler)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	s.srv = srv
	return &s
}

func (s *ApiServer) Start() {
	logger.Infof("start api srv listening: %d\n", s.port)
	if err := s.srv.ListenAndServe(); err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}
func (s *ApiServer) handler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	resp.Write([]byte(s.nmap.ToJson()))
	resp.WriteHeader(200)
}
