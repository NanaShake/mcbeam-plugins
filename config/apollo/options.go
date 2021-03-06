package apollo

import (
	"context"
	"github.com/micro/micro/v3/service/config"
	agollo "github.com/philchia/agollo/v4"
)

type appConfigKey struct{}
type namesapceKey struct{}

func WithConfig(apollo *agollo.Conf) config.Option {
	return func(o *config.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, appConfigKey{}, apollo)
	}
}

func Namespace(ns string) config.Option {
	return func(o *config.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, namesapceKey{}, ns)
	}
}