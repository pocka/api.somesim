package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"
	//"encoding/base64"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"
	swag "github.com/go-openapi/swag"

	"github.com/pocka/api.somesim/restapi/operations"
	"github.com/pocka/api.somesim/restapi/operations/flower"
	"github.com/pocka/api.somesim/restapi/operations/internal_api"
	"github.com/pocka/api.somesim/restapi/operations/item"
	"github.com/pocka/api.somesim/restapi/operations/user"
	"github.com/pocka/api.somesim/models"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name somesim --spec ../../../../../../work/docs/swagger.yml

var runtimeConfig = struct {
	DBAddress string `long:"db-address" short:"d" description:"Address for db server" default:"db"`
	AdminPassword string `long:"admin-password" description:"Password for admin user" default:"pass"`
	AdminUser string `long:"admin-user" description:"Administrator user name" default:"admin"`
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

	api.BasicAuthAuth = func(user string, pass string) (*models.Principal, error) {
		if user == runtimeConfig.AdminUser && pass == runtimeConfig.AdminPassword {
			principal := models.Principal{
				User: user,
				TokenType: models.NoToken,
				ExpiresIn: nil,
				AuthType: "basic",
			}

			return &principal, nil
		}
		return nil, errors.New(401, "Invalid username or password")
	}

	api.BearerAuth = func(value string) (*models.Principal, error) {
		// BASIC auth
		if strings.Index(value, "Basic ") == 0 {
			// If BASIC auth failed, api tries to this.
			// (Because BASIC auth has 'Authorization' header too.
			return nil, errors.New(401, "Invalid username or password")
		}

		// token auth
		if strings.Index(value, "Bearer ") == 0 {
			return nil, errors.New(501, "Token auth is not implemented")
		}

		return nil, errors.New(400, "Invalid Authorization header")
	}

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HTMLProducer = runtime.TextProducer()

	// Internal APIs
	api.InternalAPIGetDocsHandler = internal_api.GetDocsHandlerFunc(func(params internal_api.GetDocsParams) middleware.Responder {
		return internal_api.GetStaticDocFile("index.html")
	})
	api.InternalAPIGetDocsSwaggerYmlHandler = internal_api.GetDocsSwaggerYmlHandlerFunc(func(params internal_api.GetDocsSwaggerYmlParams) middleware.Responder {
		return internal_api.GetStaticDocFile("swagger.yml")
	})

	// Flower APIs
	api.FlowerGetFlowersHandler = flower.GetFlowersHandlerFunc(func(params flower.GetFlowersParams) middleware.Responder {
		return middleware.NotImplemented("operation flower.GetFlowers has not yet been implemented")
	})
	api.FlowerGetFlowersNameHandler = flower.GetFlowersNameHandlerFunc(func(params flower.GetFlowersNameParams) middleware.Responder {
		return middleware.NotImplemented("operation flower.GetFlowersName has not yet been implemented")
	})
	api.FlowerPostFlowersHandler = flower.PostFlowersHandlerFunc(func(params flower.PostFlowersParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation flower.PostFlowers has not yet been implemented")
	})
	api.FlowerPutFlowersNameHandler = flower.PutFlowersNameHandlerFunc(func(params flower.PutFlowersNameParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation flower.PutFlowersName has not yet been implemented")
	})

	// Item APIs
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
	api.ItemPostItemsHandler = item.PostItemsHandlerFunc(func(params item.PostItemsParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation item.PostItems has not yet been implemented")
	})
	api.ItemPostItemsNameImagesHandler = item.PostItemsNameImagesHandlerFunc(func(params item.PostItemsNameImagesParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation item.PostItemsNameImages has not yet been implemented")
	})
	api.ItemPutItemsNameHandler = item.PutItemsNameHandlerFunc(func(params item.PutItemsNameParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation item.PutItemsName has not yet been implemented")
	})
	api.ItemPutItemsNameImagesImageNameHandler = item.PutItemsNameImagesImageNameHandlerFunc(func(params item.PutItemsNameImagesImageNameParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation item.PutItemsNameImagesImageName has not yet been implemented")
	})

	// User APIs
	api.UserGetTokensHandler = user.GetTokensHandlerFunc(func(params user.GetTokensParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.GetTokensHandler has not yet been implemented")
	})
	api.UserGetTokensAccessTokenHandler = user.GetTokensAccessTokenHandlerFunc(func(params user.GetTokensAccessTokenParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.GetTokensAccessTokenHandler has not yet been implemented")
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
