package http

import "go.uber.org/fx"

// NewRestModule creates a new REST module with handler and route registration.
//
// Parameters:
//   - name: Module name identifier
//   - providers: Handlers and middlewares to be provided
//   - route: Route registration function
//
// Returns:
//   - fx.Option: Fx module option for dependency injection
//
// Example:
//
//	NewRestModule(
//	    "users",
//	    NewUserHandler,
//	    RegisterUserRoutes,
//	)
func NewRestModule(name string, route any, providers ...any) fx.Option {
	return fx.Module("liquor-app-rest-"+name, fx.Provide(
		providers...,
	),
		fx.Invoke(route))
}

var HttpModule = fx.Module("liquor-app-http-server", fx.Provide(
	instanceServer,
),
	fx.Invoke(
		startServer,
		initialRoute,
	))
