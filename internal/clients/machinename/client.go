// Package machinename determines the machine name of the underlying compute platform based on
// environment variables or a file (defaulting to the serial number file in the Raspberry Pi's
// device tree)
package machinename

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sargassum-world/godest"
	"github.com/sargassum-world/godest/clientcache"
	"github.com/sargassum-world/godest/env"
)

type Client struct {
	Config Config
	Logger godest.Logger
	Cache  *Cache
}

func NewClient(c Config, cache clientcache.Cache, l godest.Logger) *Client {
	return &Client{
		Config: c,
		Logger: l,
		Cache: &Cache{
			Cache: cache,
		},
	}
}

func (c *Client) GetName() (string, error) {
	if name, cacheHit := c.getNameFromCache(); cacheHit {
		c.Logger.Debugf("machine name was loaded from cache as %s", name)
		return name, nil // empty name indicates nonexistent name
	}
	return c.getNameFromSystem()
}

func (c *Client) getNameFromCache() (string, bool) {
	name, cacheHit, err := c.Cache.GetName()
	if err != nil {
		// Log the error but return as a cache miss so we can manually generate the name
		c.Logger.Error(errors.Wrap(err, "couldn't get the cache entry for the machine name"))
		return "", false // treat an unparseable cache entry like a cache miss
	}
	return name, cacheHit
}

func (c *Client) getNameFromSystem() (string, error) {
	name, err := c.getMachineName()
	if err != nil {
		c.Logger.Warnf(
			"falling back to 'unknown' as the machine name, which couldn't be determined: %s",
			err.Error(),
		)
		return "unknown", nil
	}
	if err := c.Cache.SetName(name, c.Config.CacheCost); err != nil {
		return "", errors.Wrapf(err, "couldn't cache machine name %s", name)
	}
	if name != "" {
		c.Logger.Infof("machine name was determined to be %s", name)
	}
	return name, nil
}

func (c *Client) getMachineName() (string, error) {
	name := env.GetString(envPrefix+"NAME", "")
	if name == "" {
		rawFile, err := os.ReadFile(filepath.Clean(c.Config.NameFile))
		if err != nil {
			return "", errors.Wrapf(err, "couldn't read machine name from file %s", c.Config.NameFile)
		}
		name = strings.TrimSpace(string(rawFile))
	}

	return name, nil
}
