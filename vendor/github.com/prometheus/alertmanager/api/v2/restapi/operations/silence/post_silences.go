// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostSilencesHandlerFunc turns a function with the right signature into a post silences handler
type PostSilencesHandlerFunc func(PostSilencesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostSilencesHandlerFunc) Handle(params PostSilencesParams) middleware.Responder {
	return fn(params)
}

// PostSilencesHandler interface for that can handle valid post silences params
type PostSilencesHandler interface {
	Handle(PostSilencesParams) middleware.Responder
}

// NewPostSilences creates a new http.Handler for the post silences operation
func NewPostSilences(ctx *middleware.Context, handler PostSilencesHandler) *PostSilences {
	return &PostSilences{Context: ctx, Handler: handler}
}

/*PostSilences swagger:route POST /silences silence postSilences

Post a new silence or update an existing one

*/
type PostSilences struct {
	Context *middleware.Context
	Handler PostSilencesHandler
}

func (o *PostSilences) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostSilencesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostSilencesOKBody post silences o k body
// swagger:model PostSilencesOKBody
type PostSilencesOKBody struct {

	// silence ID
	SilenceID string `json:"silenceID,omitempty"`
}

// Validate validates this post silences o k body
func (o *PostSilencesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostSilencesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSilencesOKBody) UnmarshalBinary(b []byte) error {
	var res PostSilencesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}