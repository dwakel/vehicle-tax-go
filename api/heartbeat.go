// Package classification Microservice API.
//
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     Contact:
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//
// swagger:meta
package api

import (
	"fmt"
	"log"
	"net/http"
)

type Heartbeat struct {
	l *log.Logger
}

func NewHeartbeat(l *log.Logger) *Heartbeat {
	return &Heartbeat{l}
}

// swagger:route GET / heartbeat apiHealth
// Gets Health Status of API
//responses:
//	200:

// PingHeartBeatController HeartBeat returns liveness status
func (this *Heartbeat) Heartbeat(rw http.ResponseWriter, r *http.Request) {
	this.l.Println("I'm Alive")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "I'm Alive!!!!!")
}
