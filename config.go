package etcd

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "Etcd"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Etcd struct {
	Endpoints   []string
	Username    string
	Password    string
	DialTimeout uint
}
