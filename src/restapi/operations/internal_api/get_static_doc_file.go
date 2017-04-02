package internal_api

import (
	"fmt"
	"io/ioutil"

	middleware "github.com/go-openapi/runtime/middleware"
)

func GetStaticDocFile(filename string) middleware.Responder {
	data, err := ioutil.ReadFile(fmt.Sprintf("./docs/%s", filename))

	if err != nil {
		return NewGetDocsNotFound()
	}

	payload := string(data)

	return NewGetDocsOK().WithPayload(payload)
}
