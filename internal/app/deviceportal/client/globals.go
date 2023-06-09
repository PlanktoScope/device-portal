// Package client contains client code for external APIs
package client

import (
	"github.com/pkg/errors"
	"github.com/sargassum-world/godest"
	"github.com/sargassum-world/godest/clientcache"

	"github.com/PlanktoScope/device-portal/internal/app/deviceportal/conf"
	"github.com/PlanktoScope/device-portal/internal/clients/machinename"
)

type BaseGlobals struct {
	Cache clientcache.Cache

	Logger godest.Logger
}

type Globals struct {
	Config conf.Config
	Base   *BaseGlobals

	MachineName *machinename.Client
}

func NewBaseGlobals(config conf.Config, l godest.Logger) (g *BaseGlobals, err error) {
	g = &BaseGlobals{}
	if g.Cache, err = clientcache.NewRistrettoCache(config.Cache); err != nil {
		return nil, errors.Wrap(err, "couldn't set up client cache")
	}
	g.Logger = l
	return g, nil
}

func NewGlobals(config conf.Config, l godest.Logger) (g *Globals, err error) {
	g = &Globals{
		Config: config,
	}
	g.Base, err = NewBaseGlobals(config, l)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't set up base globals")
	}

	machineNameConfig, err := machinename.GetConfig()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't set up machine-name config")
	}
	g.MachineName = machinename.NewClient(machineNameConfig, g.Base.Cache, l)
	return g, nil
}
