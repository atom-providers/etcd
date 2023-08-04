package etcd

import (
	"context"
	"time"

	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Etcd
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	if config.Username == "" {
		config.Username = "root"
	}

	return container.Container.Provide(func(ctx context.Context, logger *log.Logger) (*clientv3.Client, error) {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   config.Endpoints,
			Username:    config.Username,
			Password:    config.Password,
			DialTimeout: time.Duration(config.DialTimeout) * time.Second,
		})
		if err != nil {
			return nil, err
		}
		cli.WithLogger(logger.RawLogger)

		container.AddCloseAble(func() {
			cli.Close()
		})

		return cli, nil
	}, o.DiOptions()...)
}
