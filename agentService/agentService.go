package agentservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	nats "practice/nats"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	AddInstance(w http.ResponseWriter, r *http.Request)
	DeleteInstance(w http.ResponseWriter, r *http.Request)
}

func NewRouter(router Router) http.Handler {
	r := chi.NewRouter()

	r.Get("/addInstance", router.AddInstance)
	r.Post("/deleteInstance", router.DeleteInstance)

	return r
}

type Controller struct{}

var nc = nats.Connection()

func (c *Controller) AddInstance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("agent received add")
	nats.Nc.Publish("instance", []byte("add"))
}

func (c *Controller) DeleteInstance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("agent received del")
	id := Id{}
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &id)
	nats.Nc.Publish("instance", []byte("del "+strconv.Itoa(id.ID)))
}
