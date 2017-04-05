package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pocka/api.somesim/models"
	"github.com/pocka/api.somesim/restapi/operations/flower"
	"github.com/pocka/api.somesim/restapi/operations/internal_api"
	"github.com/pocka/api.somesim/restapi/operations/item"
	"github.com/pocka/api.somesim/restapi/operations/user"
)

// NewSomesimAPI creates a new Somesim instance
func NewSomesimAPI(spec *loads.Document) *SomesimAPI {
	return &SomesimAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
		ServeError:      errors.ServeError,
		JSONConsumer:    runtime.JSONConsumer(),
		JSONProducer:    runtime.JSONProducer(),
		HTMLProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("html producer has not yet been implemented")
		}),
		TxtProducer: runtime.TextProducer(),
		InternalAPIGetDocsHandler: internal_api.GetDocsHandlerFunc(func(params internal_api.GetDocsParams) middleware.Responder {
			return middleware.NotImplemented("operation InternalAPIGetDocs has not yet been implemented")
		}),
		InternalAPIGetDocsSwaggerYmlHandler: internal_api.GetDocsSwaggerYmlHandlerFunc(func(params internal_api.GetDocsSwaggerYmlParams) middleware.Responder {
			return middleware.NotImplemented("operation InternalAPIGetDocsSwaggerYml has not yet been implemented")
		}),
		FlowerGetFlowersHandler: flower.GetFlowersHandlerFunc(func(params flower.GetFlowersParams) middleware.Responder {
			return middleware.NotImplemented("operation FlowerGetFlowers has not yet been implemented")
		}),
		FlowerGetFlowersNameHandler: flower.GetFlowersNameHandlerFunc(func(params flower.GetFlowersNameParams) middleware.Responder {
			return middleware.NotImplemented("operation FlowerGetFlowersName has not yet been implemented")
		}),
		ItemGetItemsHandler: item.GetItemsHandlerFunc(func(params item.GetItemsParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemGetItems has not yet been implemented")
		}),
		ItemGetItemsNameHandler: item.GetItemsNameHandlerFunc(func(params item.GetItemsNameParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemGetItemsName has not yet been implemented")
		}),
		ItemGetItemsNameImagesHandler: item.GetItemsNameImagesHandlerFunc(func(params item.GetItemsNameImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemGetItemsNameImages has not yet been implemented")
		}),
		ItemGetItemsNameImagesImageNameHandler: item.GetItemsNameImagesImageNameHandlerFunc(func(params item.GetItemsNameImagesImageNameParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemGetItemsNameImagesImageName has not yet been implemented")
		}),
		UserGetTokensHandler: user.GetTokensHandlerFunc(func(params user.GetTokensParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UserGetTokens has not yet been implemented")
		}),
		UserGetTokensAccessTokenHandler: user.GetTokensAccessTokenHandlerFunc(func(params user.GetTokensAccessTokenParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UserGetTokensAccessToken has not yet been implemented")
		}),
		FlowerPostFlowersHandler: flower.PostFlowersHandlerFunc(func(params flower.PostFlowersParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation FlowerPostFlowers has not yet been implemented")
		}),
		ItemPostItemsHandler: item.PostItemsHandlerFunc(func(params item.PostItemsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation ItemPostItems has not yet been implemented")
		}),
		ItemPostItemsNameImagesHandler: item.PostItemsNameImagesHandlerFunc(func(params item.PostItemsNameImagesParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation ItemPostItemsNameImages has not yet been implemented")
		}),
		FlowerPutFlowersNameHandler: flower.PutFlowersNameHandlerFunc(func(params flower.PutFlowersNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation FlowerPutFlowersName has not yet been implemented")
		}),
		ItemPutItemsNameHandler: item.PutItemsNameHandlerFunc(func(params item.PutItemsNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation ItemPutItemsName has not yet been implemented")
		}),
		ItemPutItemsNameImagesImageNameHandler: item.PutItemsNameImagesImageNameHandlerFunc(func(params item.PutItemsNameImagesImageNameParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation ItemPutItemsNameImagesImageName has not yet been implemented")
		}),

		// Applies when the Authorization header is set with the Basic scheme
		BasicAuthAuth: func(user string, pass string) (*models.Principal, error) {
			return nil, errors.NotImplemented("basic auth  (BasicAuth) has not yet been implemented")
		},

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
	}
}

/*SomesimAPI # 概要
MoE用染色シミュレータ「そめしむ」のバックエンドAPI.
*/
type SomesimAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer
	// HTMLProducer registers a producer for a "text/html" mime type
	HTMLProducer runtime.Producer
	// TxtProducer registers a producer for a "text/plain" mime type
	TxtProducer runtime.Producer

	// BasicAuthAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	BasicAuthAuth func(string, string) (*models.Principal, error)

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (*models.Principal, error)

	// InternalAPIGetDocsHandler sets the operation handler for the get docs operation
	InternalAPIGetDocsHandler internal_api.GetDocsHandler
	// InternalAPIGetDocsSwaggerYmlHandler sets the operation handler for the get docs swagger yml operation
	InternalAPIGetDocsSwaggerYmlHandler internal_api.GetDocsSwaggerYmlHandler
	// FlowerGetFlowersHandler sets the operation handler for the get flowers operation
	FlowerGetFlowersHandler flower.GetFlowersHandler
	// FlowerGetFlowersNameHandler sets the operation handler for the get flowers name operation
	FlowerGetFlowersNameHandler flower.GetFlowersNameHandler
	// ItemGetItemsHandler sets the operation handler for the get items operation
	ItemGetItemsHandler item.GetItemsHandler
	// ItemGetItemsNameHandler sets the operation handler for the get items name operation
	ItemGetItemsNameHandler item.GetItemsNameHandler
	// ItemGetItemsNameImagesHandler sets the operation handler for the get items name images operation
	ItemGetItemsNameImagesHandler item.GetItemsNameImagesHandler
	// ItemGetItemsNameImagesImageNameHandler sets the operation handler for the get items name images image name operation
	ItemGetItemsNameImagesImageNameHandler item.GetItemsNameImagesImageNameHandler
	// UserGetTokensHandler sets the operation handler for the get tokens operation
	UserGetTokensHandler user.GetTokensHandler
	// UserGetTokensAccessTokenHandler sets the operation handler for the get tokens access token operation
	UserGetTokensAccessTokenHandler user.GetTokensAccessTokenHandler
	// FlowerPostFlowersHandler sets the operation handler for the post flowers operation
	FlowerPostFlowersHandler flower.PostFlowersHandler
	// ItemPostItemsHandler sets the operation handler for the post items operation
	ItemPostItemsHandler item.PostItemsHandler
	// ItemPostItemsNameImagesHandler sets the operation handler for the post items name images operation
	ItemPostItemsNameImagesHandler item.PostItemsNameImagesHandler
	// FlowerPutFlowersNameHandler sets the operation handler for the put flowers name operation
	FlowerPutFlowersNameHandler flower.PutFlowersNameHandler
	// ItemPutItemsNameHandler sets the operation handler for the put items name operation
	ItemPutItemsNameHandler item.PutItemsNameHandler
	// ItemPutItemsNameImagesImageNameHandler sets the operation handler for the put items name images image name operation
	ItemPutItemsNameImagesImageNameHandler item.PutItemsNameImagesImageNameHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *SomesimAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SomesimAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SomesimAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SomesimAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SomesimAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SomesimAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SomesimAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SomesimAPI
func (o *SomesimAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.HTMLProducer == nil {
		unregistered = append(unregistered, "HTMLProducer")
	}

	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.BasicAuthAuth == nil {
		unregistered = append(unregistered, "BasicAuthAuth")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.InternalAPIGetDocsHandler == nil {
		unregistered = append(unregistered, "internal_api.GetDocsHandler")
	}

	if o.InternalAPIGetDocsSwaggerYmlHandler == nil {
		unregistered = append(unregistered, "internal_api.GetDocsSwaggerYmlHandler")
	}

	if o.FlowerGetFlowersHandler == nil {
		unregistered = append(unregistered, "flower.GetFlowersHandler")
	}

	if o.FlowerGetFlowersNameHandler == nil {
		unregistered = append(unregistered, "flower.GetFlowersNameHandler")
	}

	if o.ItemGetItemsHandler == nil {
		unregistered = append(unregistered, "item.GetItemsHandler")
	}

	if o.ItemGetItemsNameHandler == nil {
		unregistered = append(unregistered, "item.GetItemsNameHandler")
	}

	if o.ItemGetItemsNameImagesHandler == nil {
		unregistered = append(unregistered, "item.GetItemsNameImagesHandler")
	}

	if o.ItemGetItemsNameImagesImageNameHandler == nil {
		unregistered = append(unregistered, "item.GetItemsNameImagesImageNameHandler")
	}

	if o.UserGetTokensHandler == nil {
		unregistered = append(unregistered, "user.GetTokensHandler")
	}

	if o.UserGetTokensAccessTokenHandler == nil {
		unregistered = append(unregistered, "user.GetTokensAccessTokenHandler")
	}

	if o.FlowerPostFlowersHandler == nil {
		unregistered = append(unregistered, "flower.PostFlowersHandler")
	}

	if o.ItemPostItemsHandler == nil {
		unregistered = append(unregistered, "item.PostItemsHandler")
	}

	if o.ItemPostItemsNameImagesHandler == nil {
		unregistered = append(unregistered, "item.PostItemsNameImagesHandler")
	}

	if o.FlowerPutFlowersNameHandler == nil {
		unregistered = append(unregistered, "flower.PutFlowersNameHandler")
	}

	if o.ItemPutItemsNameHandler == nil {
		unregistered = append(unregistered, "item.PutItemsNameHandler")
	}

	if o.ItemPutItemsNameImagesImageNameHandler == nil {
		unregistered = append(unregistered, "item.PutItemsNameImagesImageNameHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SomesimAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SomesimAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "BasicAuth":
			_ = scheme
			result[name] = security.BasicAuth(func(username, password string) (interface{}, error) {
				return o.BasicAuthAuth(username, password)
			})

		case "Bearer":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.BearerAuth(token)
			})

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *SomesimAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SomesimAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		case "text/html":
			result["text/html"] = o.HTMLProducer

		case "text/plain":
			result["text/plain"] = o.TxtProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SomesimAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the somesim API
func (o *SomesimAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SomesimAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/docs"] = internal_api.NewGetDocs(o.context, o.InternalAPIGetDocsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/docs/swagger.yml"] = internal_api.NewGetDocsSwaggerYml(o.context, o.InternalAPIGetDocsSwaggerYmlHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flowers"] = flower.NewGetFlowers(o.context, o.FlowerGetFlowersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flowers/{name}"] = flower.NewGetFlowersName(o.context, o.FlowerGetFlowersNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/items"] = item.NewGetItems(o.context, o.ItemGetItemsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/items/{name}"] = item.NewGetItemsName(o.context, o.ItemGetItemsNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/items/{name}/images"] = item.NewGetItemsNameImages(o.context, o.ItemGetItemsNameImagesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/items/{name}/images/{image_name}"] = item.NewGetItemsNameImagesImageName(o.context, o.ItemGetItemsNameImagesImageNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tokens"] = user.NewGetTokens(o.context, o.UserGetTokensHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tokens/access_token"] = user.NewGetTokensAccessToken(o.context, o.UserGetTokensAccessTokenHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flowers"] = flower.NewPostFlowers(o.context, o.FlowerPostFlowersHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/items"] = item.NewPostItems(o.context, o.ItemPostItemsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/items/{name}/images"] = item.NewPostItemsNameImages(o.context, o.ItemPostItemsNameImagesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flowers/{name}"] = flower.NewPutFlowersName(o.context, o.FlowerPutFlowersNameHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/items/{name}"] = item.NewPutItemsName(o.context, o.ItemPutItemsNameHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/items/{name}/images/{image_name}"] = item.NewPutItemsNameImagesImageName(o.context, o.ItemPutItemsNameImagesImageNameHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SomesimAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *SomesimAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
