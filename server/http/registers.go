package http

import "go.uber.org/fx"

// RegisterRoutes register you routes
func RegisterRoutes(routes ...interface{}) fx.Option {
	return fx.Invoke(routes...)
}

// RegisterHandlers register you handlers
func RegisterHandlers(handlers ...interface{}) fx.Option {
	return fx.Provide(handlers...)
}

// NewRestModule creates a new REST module with handler and route registration.
//
// Parameters:
//   - name: Module name identifier
//   - handler: HTTP handler implementation
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
func NewRestModule(name string, handler any, route any) fx.Option {
	return fx.Module("liquor-app-rest-"+name, fx.Provide(
		handler,
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
