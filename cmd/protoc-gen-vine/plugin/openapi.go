// Copyright 2021 lack
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

	"github.com/lack-io/vine/cmd/generator"
)

type ComponentKind int32

const (
	Auth ComponentKind = iota
	Request
	Response
	Error
)

type Component struct {
	Name    string
	Kind    ComponentKind
	Service string
	Proto   *generator.MessageDescriptor
}

func (g *vine) generateOpenAPI(svc *generator.ServiceDescriptor) {
	srvName := svc.Proto.GetName()
	srvTags := g.extractTags(svc.Comments)
	if _, ok := srvTags[_openapi]; !ok {
		return
	}
	g.P(`Openapi: "3.0.1",`)
	g.P("Info: &registry.OpenAPIInfo{")
	g.P(`Title: "`, srvName, `Service",`)
	desc := extractDesc(svc.Comments)
	if len(desc) == 0 {
		desc = []string{"OpenAPI3.0 for " + srvName}
	}
	g.P(`Description: "'`, strings.Join(desc, " "), `'",`)
	term, ok := srvTags[_termURL]
	if ok {
		g.P(fmt.Sprintf(`TermsOfService: "%s",`, term.Value))
	}
	contactName, ok1 := srvTags[_contactName]
	contactEmail, ok2 := srvTags[_contactEmail]
	if ok1 && ok2 {
		g.P("Contact: &registry.OpenAPIContact{")
		g.P(fmt.Sprintf(`Name: "%s",`, contactName.Value))
		g.P(fmt.Sprintf(`Email: "%s",`, contactEmail.Value))
		g.P("},")
	}
	licenseName, ok1 := srvTags[_licenseName]
	licenseUrl, ok2 := srvTags[_licenseUrl]
	if ok1 && ok2 {
		g.P("License: &registry.OpenAPILicense{")
		g.P(fmt.Sprintf(`Name: "%s",`, licenseName.Value))
		g.P(fmt.Sprintf(`Url: "%s",`, licenseUrl.Value))
		g.P("},")
	}
	g.P("},")
	externalDocDesc, ok1 := srvTags[_externalDocDesc]
	externalDocUrl, ok2 := srvTags[_externalDocUrl]
	if ok1 && ok2 {
		g.P("ExternalDocs: &registry.OpenAPIExternalDocs{")
		g.P(fmt.Sprintf(`Description: "%s",`, externalDocDesc.Value))
		g.P(fmt.Sprintf(`Url: "%s",`, externalDocUrl.Value))
		g.P("},")
	}
	g.P("Servers: []string{},")
	g.P("Tags: []*registry.OpenAPITag{")
	g.P("&registry.OpenAPITag{")
	g.P(fmt.Sprintf(`Name: "%s",`, srvName))
	g.P(fmt.Sprintf(`Description: "%s",`, strings.Join(desc, " ")))
	if externalDocDesc != nil && externalDocUrl != nil {
		g.P("ExternalDocs: &registry.OpenAPIExternalDocs{")
		g.P(fmt.Sprintf(`Description: "%s",`, externalDocDesc.Value))
		g.P(fmt.Sprintf(`Url: "%s",`, externalDocUrl.Value))
		g.P("},")
	}
	g.P("},")
	g.P("},")

	g.P(`Paths: map[string]*registry.OpenAPIPath{`)
	for _, meth := range svc.Methods {
		g.generateMethodOpenAPI(svc, meth)
	}
	g.P("},")
	g.P(`Components: &registry.OpenAPIComponents{`)
	g.generateComponents(srvName)
	g.P("},")
}

func (g *vine) generateMethodOpenAPI(svc *generator.ServiceDescriptor, method *generator.MethodDescriptor) {
	srvName := svc.Proto.GetName()
	methodName := method.Proto.GetName()
	tags := g.extractTags(method.Comments)
	if len(tags) == 0 {
		return
	}
	var meth string
	var path string
	if v, ok := tags[_get]; ok {
		meth = v.Key
		path = v.Value
	} else if v, ok = tags[_post]; ok {
		meth = v.Key
		path = v.Value
	} else if v, ok = tags[_patch]; ok {
		meth = v.Key
		path = v.Value
	} else if v, ok = tags[_put]; ok {
		meth = v.Key
		path = v.Value
	} else if v, ok = tags[_delete]; ok {
		meth = v.Key
		path = v.Value
	} else {
		return
	}

	pathParams := g.extractPathParams(path)

	summary, _ := tags[_summary]
	g.P(fmt.Sprintf(`"%s": &registry.OpenAPIPath{`, path))
	g.P(fmt.Sprintf(`%s: &registry.OpenAPIPathDocs{`, generator.CamelCase(meth)))
	g.P(fmt.Sprintf(`Tags: []string{"%s"},`, srvName))
	if summary != nil {
		g.P(fmt.Sprintf(`Summary: "%s",`, summary.Value))
	}
	desc := extractDesc(method.Comments)
	if len(desc) == 0 {
		desc = []string{srvName + " " + methodName}
	}
	g.P(fmt.Sprintf(`Description: "%s",`, strings.Join(desc, " ")))
	g.P(fmt.Sprintf(`OperationId: "%s", `, srvName+methodName))
	msg := g.extractMessage(method.Proto.GetInputType())
	if msg == nil {
		g.gen.Fail("%s not found", method.Proto.GetInputType())
		return
	}
	mname := g.extractImportMessageName(msg)
	g.schemas[mname] = &Component{
		Name:    mname,
		Kind:    Request,
		Service: srvName,
		Proto:   msg,
	}

	if len(pathParams) > 0 || meth == _get {
		g.P("Parameters: []*registry.PathParameters{")
		g.generateParameters(srvName, msg, pathParams)
		g.P("},")
	}
	if meth != _get {
		g.P("RequestBody: &registry.PathRequestBody{")
		desc := extractDesc(msg.Comments)
		if len(desc) == 0 {
			desc = []string{methodName + " " + msg.Proto.GetName()}
		}
		g.P(fmt.Sprintf(`Description: "%s",`, strings.Join(desc, " ")))
		g.P("Content: &registry.PathRequestBodyContent{")
		g.P("ApplicationJson: &registry.ApplicationContent{")
		g.P("Schema: &registry.Schema{")
		g.P(fmt.Sprintf(`Ref: "#/components/schemas/%s",`, mname))
		g.P("},")
		g.P("},")
		g.P("},")
		g.P("},")
	}
	msg = g.extractMessage(method.Proto.GetOutputType())
	if msg == nil {
		g.gen.Fail("%s not found", method.Proto.GetOutputType())
		return
	}
	mname = g.extractImportMessageName(msg)
	g.schemas[mname] = &Component{
		Name:    mname,
		Kind:    Response,
		Service: srvName,
		Proto:   msg,
	}
	g.P("Responses: map[string]*registry.PathResponse{")
	g.generateResponse(msg, tags)
	g.P("},")
	g.P(`Security: []*registry.PathSecurity{`)
	g.generateSecurity(tags)
	g.P("},")
	g.P("},")
	g.P("},")
}

func (g *vine) generateParameters(srvName string, msg *generator.MessageDescriptor, paths []string) {
	if msg == nil {
		return
	}

	generateField := func(g *vine, field *generator.FieldDescriptor, in string) {
		tags := g.extractTags(field.Comments)
		g.gen.P("&registry.PathParameters{")
		g.P(fmt.Sprintf(`Name: "%s",`, field.Proto.GetJsonName()))
		g.P(fmt.Sprintf(`In: "%s",`, in))
		desc := extractDesc(field.Comments)
		if len(desc) == 0 {
			desc = []string{msg.Proto.GetName() + " field " + field.Proto.GetJsonName()}
		}
		g.P(fmt.Sprintf(`Description: "%s",`, strings.Join(desc, " ")))
		if in == "path" {
			g.P("Required: true,")
		} else if len(tags) > 0 {
			if _, ok := tags[_required]; ok {
				g.P("Required: true,")
			}
		}
		g.P(`Style: "form",`)
		g.P("Explode: true,")
		g.P("Schema: &registry.Schema{")
		fieldTags := g.extractTags(field.Comments)
		g.generateSchema(srvName, field, fieldTags, true)
		g.P("},")
		g.P("},")
	}

	fields := make([]string, 0)
	for _, p := range paths {
		field := g.extractMessageField(srvName, p, msg)
		generateField(g, field, "path")
		fields = append(fields, field.Proto.GetJsonName())
	}

	for _, field := range msg.Fields {
		for _, f := range fields {
			if f == field.Proto.GetJsonName() {
				continue
			}
		}
		generateField(g, field, "query")
	}
}

func (g *vine) generateResponse(msg *generator.MessageDescriptor, tags map[string]*Tag) {
	printer := func(code int32, desc, schema string) {
		g.P(fmt.Sprintf(`"%d": &registry.PathResponse{`, code))
		g.P(fmt.Sprintf(`Description: "%s",`, desc))
		g.P(`Content: &registry.PathRequestBodyContent{`)
		g.P(`ApplicationJson: &registry.ApplicationContent{`)
		g.P(fmt.Sprintf(`Schema: &registry.Schema{Ref: "#/components/errors/%s"},`, schema))
		g.P("},")
		g.P("},")
		g.P("},")
	}

	// 200 result
	printer(200, "successful response (stream response)", msg.Proto.GetName())

	t, ok := tags[_result]
	if !ok {
		return
	}

	if _, ok := tags[_security]; ok {
		printer(401, "Unauthorized", "VineError")
		printer(403, "Forbidden", "VineError")
	}

	s := strings.TrimPrefix(t.Value, "[")
	s = strings.TrimSuffix(s, "]")
	parts := strings.Split(s, ",")
	if len(parts) > 0 {
		g.errors["VineError"] = &Component{
			Name: "VineError",
			Kind: Error,
		}
	}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		code, _ := strconv.ParseInt(part, 10, 64)
		if !(code >= 200 && code <= 599) {
			g.gen.Fail("invalid result code: %s", part)
			return
		}
		switch code {
		case 400:
			printer(400, "BadRequest", "VineError")
		case 404:
			printer(404, "NotFound", "VineError")
		case 405:
			printer(405, "MethodNotAllowed", "VineError")
		case 408:
			printer(408, "Timeout", "VineError")
		case 409:
			printer(409, "Conflict", "VineError")
		case 500:
			printer(500, "InternalServerError", "VineError")
		case 501:
			printer(501, "NotImplemented", "VineError")
		case 502:
			printer(502, "BadGateway", "VineError")
		case 503:
			printer(503, "ServiceUnavailable", "VineError")
		case 504:
			printer(504, "GatewayTimeout", "VineError")
		}
	}
}

func (g *vine) generateSecurity(tags map[string]*Tag) {
	if len(tags) == 0 {
		return
	}

	t, ok := tags[_security]
	if !ok {
		return
	}

	cp := &Component{Kind: Auth}
	g.P(`&registry.PathSecurity{`)
	parts := strings.Split(t.Value, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		switch p {
		case "bearer":
			g.P("Bearer: []string{},")
			cp.Name = "Bearer"
		case "apiKeys":
			g.P("ApiKeys: []string{},")
			cp.Name = "apiKeys"
		case "basic":
			g.P("Basic: []string{},")
			cp.Name = "basic"
		default:
			g.gen.Fail("invalid security type: ", p)
			return
		}
	}
	g.security[cp.Name] = cp
	g.P("},")
}

func (g *vine) generateComponents(srvName string) {
	g.P(`SecuritySchemas: &registry.SecuritySchemas{`)
	for _, c := range g.security {
		switch c.Name {
		case "Bearer":
			g.P(`Basic: &registry.BearerSecurity{Type: "http", Schema: "bearer"}`)
		case "ApiKeys":
			g.P(`ApiKeys: &registry.APIKeysSecurity{Type: "apiKey", In: "header", Name: "X-API-Key"},`)
		case "Basic":
			g.P(`Basic: &registry.BasicSecurity{Type: "http", Schema: "basic"}`)
		}
	}
	g.P("},")

	fn := func(schemas map[string]*Component) {
		for name, c := range schemas {
			switch c.Kind {
			case Request:
				g.P(fmt.Sprintf(`"%s": &registry.Model{`, name))
				g.P(`Type: "object",`)
				g.P(`Properties: map[string]*registry.Schema{`)
				requirements := []string{}
				for _, field := range c.Proto.Fields {
					tags := g.extractTags(field.Comments)
					if _, ok := tags[_required]; ok {
						requirements = append(requirements, `"`+field.Proto.GetJsonName()+`"`)
					}
					g.P(fmt.Sprintf(`"%s": &registry.Schema{`, field.Proto.GetJsonName()))
					g.generateSchema(srvName, field, tags, false)
					g.P("},")
				}
				g.P("},")
				if len(requirements) > 0 {
					g.P(fmt.Sprintf(`Required: []string{%s},`, strings.Join(requirements, ",")))
				}
				g.P("},")
			case Response:
				g.P(fmt.Sprintf(`"%s": &registry.Model{`, name))
				g.P(`Type: "object",`)
				g.P(`Properties: map[string]*registry.Schema{`)
				for _, field := range c.Proto.Fields {
					tags := g.extractTags(field.Comments)
					g.P(fmt.Sprintf(`"%s": &registry.Schema{`, field.Proto.GetJsonName()))
					g.generateSchema(srvName, field, tags, false)
					g.P("},")
				}
				g.P("},")
				g.P("},")
			case Error:

			}
		}
	}

	g.P(`Schemas: map[string]*registry.Model{`)
	fn(g.schemas)
	fn(g.extSchemas)
	g.P("},")
	g.P(`Errors: map[string]*registry.Model{`)
	for name, _ := range g.errors {
		if name == "VineError" {
			g.P(`"VineError": &registry.Model{
					Type: "object",
					Properties: map[string]*registry.Schema{
						"id":       &registry.Schema{Type: "string", Description: "the name from component"},
						"code":     &registry.Schema{Type: "integer", Format: "int32", Description: "the code from http"},
						"detail":   &registry.Schema{Type: "string", Description: "the detail message for error"},
						"status":   &registry.Schema{Type: "string", Description: "a text for the HTTP status code"},
						"position": &registry.Schema{Type: "string", Description: "the code position for error"},
						"child":    &registry.Schema{Type: "object", Description: "more message", Ref: "#/components/errors/Child"},
						"stacks":   &registry.Schema{Type: "array", Description: "external message", Items: []*registry.Schema{&registry.Schema{Type: "object", Ref: "#/components/errors/Stack"}}},
					},
				},
				"Child": &registry.Model{
					Type: "object",
					Properties: map[string]*registry.Schema{
						"code":   &registry.Schema{Type: "integer", Description: "context status code", Format: "int32"},
						"detail": &registry.Schema{Type: "string", Description: "context error message"},
					},
				},
				"Stack": &registry.Model{
					Type: "object",
					Properties: map[string]*registry.Schema{
						"code":     &registry.Schema{Type: "integer", Format: "int32", Description: "more status code"},
						"detail":   &registry.Schema{Type: "string", Description: "more message"},
						"position": &registry.Schema{Type: "string", Description: "the position for more message"},
					},
				},`)
		}
	}
	g.P("},")
}

func (g *vine) generateSchema(srvName string, field *generator.FieldDescriptor, tags map[string]*Tag, allowRequired bool) {
	generateNumber := func(g *vine, field *generator.FieldDescriptor, tags map[string]*Tag) {
		if _, ok := tags[_required]; ok {
			if allowRequired {
				g.P(`Required: true,`)
			}
		}
		for key, tag := range tags {
			switch key {
			case _enum, _in:
				g.P(fmt.Sprintf(`Enum: []string{%s},`, fullStringSlice(tag.Value)))
			case _gt:
				g.P("ExclusiveMinimum: true")
				g.P(fmt.Sprintf(`Minimum: %s,`, tag.Value))
			case _gte:
				g.P(fmt.Sprintf(`Minimum: %s,`, tag.Value))
			case _lt:
				g.P("ExclusiveMaximum: true")
				g.P(fmt.Sprintf(`Maximum: %s,`, tag.Value))
			case _lte:
				g.P(fmt.Sprintf(`Maximum: %s,`, tag.Value))
			case _readOnly:
				g.P(`ReadOnly: true,`)
			case _writeOnly:
				g.P(`WriteOnly: true,`)
			case _default:
				g.P(fmt.Sprintf(`Default: "%s",`, TrimString(tag.Value, `"`)))
			case _example:
				g.P(fmt.Sprintf(`Example: "%s",`, TrimString(tag.Value, `"`)))
			}
		}
	}

	generateString := func(g *vine, field *generator.FieldDescriptor, tags map[string]*Tag) {
		if _, ok := tags[_required]; ok {
			if allowRequired {
				g.P(`Required: true,`)
			}
		}
		for key, tag := range tags {
			switch key {
			case _enum, _in:
				g.P(fmt.Sprintf(`Enum: []string{%s},`, fullStringSlice(tag.Value)))
			case _minLen:
				g.P(fmt.Sprintf(`MinLength: %s,`, tag.Value))
			case _maxLen:
				g.P(fmt.Sprintf(`MinLength: %s,`, tag.Value))
			case _date:
				g.P(`Format: "date",`)
			case _dateTime:
				g.P(`Format: "date-time",`)
			case _password:
				g.P(`Format: "password",`)
			case _byte:
				g.P(`Format: "byte",`)
			case _binary:
				g.P(`Format: "binary",`)
			case _email:
				g.P(`Format: "email",`)
			case _uuid:
				g.P(`Format: "uuid",`)
			case _uri:
				g.P(`Format: "uri",`)
			case _hostname:
				g.P(`Format: "hostname",`)
			case _ip, _ipv4:
				g.P(`Format: "ipv4",`)
			case _ipv6:
				g.P(`Format: "ipv6",`)
			case _readOnly:
				g.P(`ReadOnly: true,`)
			case _writeOnly:
				g.P(`WriteOnly: true,`)
			case _pattern:
				g.P(fmt.Sprintf(`Pattern: "'%s'",`, TrimString(tag.Value, "`")))
			case _default:
				g.P(fmt.Sprintf(`Default: "%s",`, TrimString(tag.Value, `"`)))
			case _example:
				g.P(fmt.Sprintf(`Example: "%s",`, TrimString(tag.Value, `"`)))
			}
		}
	}

	// generate map
	if field.Proto.IsRepeated() && strings.HasSuffix(field.Proto.GetTypeName(), "Entry") {
		if _, ok := tags[_required]; ok {
			if allowRequired {
				g.P(`Required: true,`)
			}
		}
		g.P(`Type: "object",`)
		g.P(`AdditionalProperties: &registry.Schema{`)
		msg := g.extractMessage(field.Proto.GetTypeName())
		if msg == nil {
			g.gen.Fail("message<%s> not found", field.Proto.GetTypeName())
			return
		}
		var valueField *generator.FieldDescriptor
		for _, fd := range msg.Fields {
			if fd.Proto.GetName() == "value" {
				valueField = fd
			}
		}
		if valueField != nil {
			mname := g.extractImportMessageName(msg)
			g.extSchemas[mname] = &Component{
				Name:    mname,
				Kind:    Request,
				Service: srvName,
				Proto:   msg,
			}
			g.generateSchema(srvName, valueField, g.extractTags(valueField.Comments), allowRequired)
		} else {
			// inner MapEntry
			name := field.Proto.GetTypeName()
			if index := strings.LastIndex(name, "."); index > 0 {
				name = name[index+1:]
			}
			for _, m := range g.gen.File().Messages() {
				if m.Proto.GetName() == name {
					for _, f := range m.Fields {
						if f.Proto.GetName() == "value" {
							g.generateSchema(srvName, f, g.extractTags(f.Comments), allowRequired)
						}
					}
				}
			}
		}
		g.P(`},`)
		return
	}

	if field.Proto.IsRepeated() && !strings.HasSuffix(field.Proto.GetTypeName(), "Entry") {
		if _, ok := tags[_required]; ok {
			if allowRequired {
				g.P(`Required: true,`)
			}
		}
		g.P(`Type: "array",`)
		switch field.Proto.GetType() {
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
			descriptor.FieldDescriptorProto_TYPE_FLOAT,
			descriptor.FieldDescriptorProto_TYPE_INT64,
			descriptor.FieldDescriptorProto_TYPE_INT32,
			descriptor.FieldDescriptorProto_TYPE_FIXED64,
			descriptor.FieldDescriptorProto_TYPE_FIXED32,
			descriptor.FieldDescriptorProto_TYPE_STRING:
			g.P(`Items: []*registry.Schema{`)
			g.P(`&registry.Schema{Type: "integer"},`)
			g.P("},")
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			g.P(`Items: []*registry.Schema{`)
			msg := g.extractMessage(field.Proto.GetTypeName())
			if msg == nil {
				g.gen.Fail("message<%s> not found", field.Proto.GetTypeName())
				return
			}
			mname := g.extractImportMessageName(msg)
			g.extSchemas[mname] = &Component{
				Name:    mname,
				Kind:    Request,
				Service: srvName,
				Proto:   msg,
			}
			g.P(fmt.Sprintf(`&registry.Schema{Type: "object", Ref: "#/components/schemas/%s"},`, mname))
			g.P("},")
		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			g.P(`Items: []*registry.Schema{`)
			g.P(`&registry.Schema{Type: "boolean"},`)
			g.P("},")
		}
		return
	}

	switch field.Proto.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		g.P(`Type: "number",`)
		g.P(`Format: "double",`)
		generateNumber(g, field, tags)
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		g.P(`Type: "number",`)
		g.P(`Format: "float",`)
		generateNumber(g, field, tags)
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		g.P(`Type: "integer",`)
		g.P(`Format: "int64",`)
		generateNumber(g, field, tags)
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		g.P(`Type: "integer",`)
		g.P(`Format: "int32",`)
		generateNumber(g, field, tags)
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		g.P(`Type: "integer",`)
		g.P(`Format: "int32",`)
		tags[_gte] = &Tag{Key: _gte, Value: "0"}
		generateNumber(g, field, tags)
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		g.P(`Type: "integer",`)
		g.P(`Format: "int32",`)
		tags[_gte] = &Tag{Key: _gte, Value: "0"}
		generateNumber(g, field, tags)
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		g.P(`Type: "string",`)
		generateString(g, field, tags)
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		g.P(`Type: "object",`)
		if _, ok := tags[_required]; ok {
			if allowRequired {
				g.P(`Required: true,`)
			}
		}
		msg := g.extractMessage(field.Proto.GetTypeName())
		if msg == nil {
			g.gen.Fail("message<%s> not found", field.Proto.GetTypeName())
			return
		}
		mname := g.extractImportMessageName(msg)
		g.extSchemas[mname] = &Component{
			Name:    mname,
			Kind:    Request,
			Service: srvName,
			Proto:   msg,
		}
		g.P(fmt.Sprintf(`Ref: "#/components/schemas/%s",`, mname))
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		g.P(`Type: "boolean"`)
		if _, ok := tags[_required]; ok {
			if allowRequired {
				g.P(`Required: true,`)
			}
		}
	}
}

func (g *vine) extractMessageField(srvName, fname string, msg *generator.MessageDescriptor) *generator.FieldDescriptor {
	name := fname
	index := strings.Index(fname, ".")
	if index > 0 {
		name = fname[:index]
	}
	for _, field := range msg.Fields {
		switch {
		case *field.Proto.JsonName == fname && !field.Proto.IsMessage():
			return field
		case *field.Proto.JsonName == name && index > 0 && field.Proto.IsMessage():
			submsg := g.extractMessage(field.Proto.GetTypeName())
			if submsg == nil {
				g.gen.Fail("message<%s> not found", field.Proto.GetTypeName())
				return nil
			}
			mname := g.extractImportMessageName(submsg)
			g.schemas[mname] = &Component{
				Name:    mname,
				Kind:    Request,
				Service: srvName,
				Proto:   submsg,
			}
			return g.extractMessageField(srvName, fname[index+1:], submsg)
		}
	}
	g.gen.Fail("%s not found", fname)
	return nil
}

// extractMessage extract MessageDescriptor by name
func (g *vine) extractMessage(name string) *generator.MessageDescriptor {
	obj := g.gen.ObjectNamed(name)
	for _, f := range g.gen.AllFiles() {
		for _, m := range f.Messages() {
			if m.Proto.GoImportPath() == obj.GoImportPath() {
				for _, item := range obj.TypeName() {
					if item == m.Proto.GetName() {
						return m
					}
				}
			}
		}
	}
	return nil
}

// extractPathParams extract parameters by router path. e.g /{id}/{name}
func (g *vine) extractPathParams(path string) []string {
	paths := []string{}

	var cur int
	for i, c := range path {
		if c == '{' {
			cur = i
			continue
		}
		if c == '}' {
			if cur+1 >= i {
				g.gen.Fail("invalid path")
				return nil
			}
			paths = append(paths, path[cur+1:i])
			cur = 0
		}
	}
	if cur != 0 {
		g.gen.Fail("invalid path")
		return nil
	}
	return paths
}

func (g *vine) extractImportMessageName(msg *generator.MessageDescriptor) string {
	pkg := msg.Proto.GoImportPath().String()
	pkg = TrimString(pkg, `"`)
	if index := strings.LastIndex(pkg, "/"); index > 0 {
		pkg = pkg[index+1:]
	}
	return pkg + "." + msg.Proto.GetName()
}