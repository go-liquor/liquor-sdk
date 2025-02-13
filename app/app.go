package app

import (
	"github.com/go-liquor/liquor-sdk/config"
	"github.com/go-liquor/liquor-sdk/logger"
	"github.com/go-liquor/liquor-sdk/server/http"
	"go.uber.org/fx"
)

// NewApp create a new app
func NewApp(modules ...fx.Option) {
	options := []fx.Option{
		config.ConfigModule,
		logger.LoggerModule,
		http.HttpModule,
	}
	options = append(options, modules...)
	app := fx.New(options...)
	app.Run()
}

func NewModule(moduleName string, in ...any) fx.Option {
	var ops = []fx.Option{}

	for _, i := range in {
		if opt, ok := i.(fx.Option); ok {
			ops = append(ops, opt)
			continue
		}
		ops = append(ops, fx.Provide(i))
	}

	return fx.Module(moduleName, ops...)
}

// RegisterServices register you services
func RegisterServices(services ...interface{}) fx.Option {
	return fx.Provide(
		services...,
	)
}

// RegisterProviders register you providers
func RegisterProviders(providers ...interface{}) fx.Option {
	return fx.Module("liquor-app-providers", fx.Provide(
		providers...,
	))
}
