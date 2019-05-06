package main

import (
	"fmt"

	knot "github.com/eaciit/knot/knot.v1"
	"github.com/raditzlawliet/go-jwt-example/controller"
)

func init() {

}

func main() {
	api := controller.Api{}

	ks := new(knot.Server)
	ks.Address = fmt.Sprintf(":%v", 4000)

	ks.Route("/stop", func(wc *knot.WebContext) interface{} {
		wc.Server.Stop()
		return ""
	})

	ks.Register(&api, "")
	ks.Listen()

}
