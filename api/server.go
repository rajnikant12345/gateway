package api

import (
	"net/http"

	"github.com/MiteshSharma/gateway/util"
	"github.com/urfave/negroni"
)

type Server struct {
	Router *negroni.Negroni
}

var ServerObj *Server

func InitServer() {
	ServerObj = &Server{}
	ServerObj.Router = InitApi()
}

func StartServer() {
	go func() {
		err := http.ListenAndServe(utils.GatewayConfigParam.ServerConfig.Port, ServerObj.Router)
		if err != nil {
			log.WithField("err", err).Fatal("Server starting failed.")
			return
		}
	}()
}

func StopServer() {

}
