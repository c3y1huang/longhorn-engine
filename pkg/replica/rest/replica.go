package rest

import (
	"github.com/rancher/go-rancher/api"
	"github.com/rancher/go-rancher/client"
	"net/http"
)

// ListReplicas writes replicas to response
func (s *Server) ListReplicas(rw http.ResponseWriter, req *http.Request) error {
	apiContext := api.GetApiContext(req)
	resp := client.GenericCollection{}
	resp.Data = append(resp.Data, s.Replica(apiContext))

	apiContext.Write(&resp)
	return nil
}

// Replica returns new Replica with state and info
func (s *Server) Replica(apiContext *api.ApiContext) *Replica {
	state, info := s.s.Status()
	return NewReplica(apiContext, state, info, s.s.Replica())
}

// doOp writes Replica with state and info to reponse
func (s *Server) doOp(req *http.Request, err error) error {
	if err != nil {
		return err
	}

	apiContext := api.GetApiContext(req)
	apiContext.Write(s.Replica(apiContext))
	return nil
}
