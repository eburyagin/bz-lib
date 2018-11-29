package app

import (
	"bz-lib/cfg"

	nats "github.com/nats-io/go-nats"
)

type AppContext struct {
	Nats *nats.EncodedConn
	Conf *cfg.Config
}
