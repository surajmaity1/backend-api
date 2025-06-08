package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Environment string

const (
	Development Environment = "development"
)

func (e Environment) Validate() Environment {
	if e != Development {
		logrus.Warn(fmt.Sprintf("%s is not a valid environment", e))
		logrus.Warn("By default, setting it to development")
		e = Development
	}
	return e
}
