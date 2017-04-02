package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"
	swag "github.com/go-openapi/swag"

	"github.com/pocka/api.somesim/restapi/operations"
	"github.com/pocka/api.somesim/restapi/operations/flower"
	"github.com/pocka/api.somesim/restapi/operations/internal_api"
	"github.com/pocka/api.somesim/restapi/operations/item"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name somesim --spec ../../../../../../work/docs/swagger.yml

var runtimeConfig = struct {
	DBAddress string `long:"db-address" short:"d" description:"Address for db server" default:"db"`
	Password string `long:"password" description:"Password for admin user" default:"pass"`
}{}

func configureFlags(api *operations.SomesimAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Runtime Config",
			LongDescription: "",
			Options: &runtimeConfig,
		},
	}
}

func configureAPI(api *operations.SomesimAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.Middleware = func(builder middleware.Builder) http.Handler {
		return api.Context().RoutesHandler(builder)
	}

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HTMLProducer = runtime.TextProducer()

	api.InternalAPIGetDocsHandler = internal_api.GetDocsHandlerFunc(func(params internal_api.GetDocsParams) middleware.Responder {
		return internal_api.GetStaticDocFile("index.html")
	})
	api.InternalAPIGetDocsSwaggerYmlHandler = internal_api.GetDocsSwaggerYmlHandlerFunc(func(params internal_api.GetDocsSwaggerYmlParams) middleware.Responder {
		return internal_api.GetStaticDocFile("swagger.yml")
	})
	api.FlowerGetFlowersHandler = flower.GetFlowersHandlerFunc(func(params flower.GetFlowersParams) middleware.Responder {
		return middleware.NotImplemented("operation flower.GetFlowers has not yet been implemented")
	})
	api.FlowerGetFlowersNameHandler = flower.GetFlowersNameHandlerFunc(func(params flower.GetFlowersNameParams) middleware.Responder {
		return middleware.NotImplemented("operation flower.GetFlowersName has not yet been implemented")
	})
	api.ItemGetItemsHandler = item.GetItemsHandlerFunc(func(params item.GetItemsParams) middleware.Responder {
		return middleware.NotImplemented("operation item.GetItems has not yet been implemented")
	})
	api.ItemGetItemsNameHandler = item.GetItemsNameHandlerFunc(func(params item.GetItemsNameParams) middleware.Responder {
		return middleware.NotImplemented("operation item.GetItemsName has not yet been implemented")
	})
	api.ItemGetItemsNameImagesHandler = item.GetItemsNameImagesHandlerFunc(func(params item.GetItemsNameImagesParams) middleware.Responder {
		return middleware.NotImplemented("operation item.GetItemsNameImages has not yet been implemented")
	})
	api.ItemGetItemsNameImagesImageNameHandler = item.GetItemsNameImagesImageNameHandlerFunc(func(params item.GetItemsNameImagesImageNameParams) middleware.Responder {
		return middleware.NotImplemented("operation item.GetItemsNameImagesImageName has not yet been implemented")
	})
	api.FlowerPostFlowersHandler = flower.PostFlowersHandlerFunc(func(params flower.PostFlowersParams) middleware.Responder {
		return middleware.NotImplemented("operation flower.PostFlowers has not yet been implemented")
	})
	api.ItemPostItemsHandler = item.PostItemsHandlerFunc(func(params item.PostItemsParams) middleware.Responder {
		return middleware.NotImplemented("operation item.PostItems has not yet been implemented")
	})
	api.ItemPostItemsNameImagesHandler = item.PostItemsNameImagesHandlerFunc(func(params item.PostItemsNameImagesParams) middleware.Responder {
		return middleware.NotImplemented("operation item.PostItemsNameImages has not yet been implemented")
	})
	api.FlowerPutFlowersNameHandler = flower.PutFlowersNameHandlerFunc(func(params flower.PutFlowersNameParams) middleware.Responder {
		return middleware.NotImplemented("operation flower.PutFlowersName has not yet been implemented")
	})
	api.ItemPutItemsNameHandler = item.PutItemsNameHandlerFunc(func(params item.PutItemsNameParams) middleware.Responder {
		return middleware.NotImplemented("operation item.PutItemsName has not yet been implemented")
	})
	api.ItemPutItemsNameImagesImageNameHandler = item.PutItemsNameImagesImageNameHandlerFunc(func(params item.PutItemsNameImagesImageNameParams) middleware.Responder {
		return middleware.NotImplemented("operation item.PutItemsNameImagesImageName has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
