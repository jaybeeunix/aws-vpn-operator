package controller

import (
	"github.com/jaybeeunix/aws-vpn-operator/pkg/controller/vpngateway"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, vpngateway.Add)
}
