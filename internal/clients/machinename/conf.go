package machinename

import (
	"github.com/pkg/errors"
	"github.com/sargassum-world/godest/env"
)

const envPrefix = "MACHINENAME_"

type Config struct {
	NameFile string

	CacheCost float32
}

func GetConfig() (c Config, err error) {
	// This is a file path specific to the PlanktoScope OS
	const defaultNameFilePath = "/run/machine-name"
	c.NameFile = env.GetString(envPrefix+"NAMEFILE", defaultNameFilePath)

	const defaultCacheCost = 1.0
	c.CacheCost, err = env.GetFloat32(envPrefix+"CACHE_COST", defaultCacheCost)
	if err != nil {
		return Config{}, errors.Wrap(err, "couldn't make cache cost config")
	}
	return c, nil
}
