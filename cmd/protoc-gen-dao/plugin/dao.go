// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

	"github.com/lack-io/vine/cmd/generator"
)

var TagString = "gen"

const (
	// message tag
	_ignore = "ignore"

	// field common tag
	_required = "required"
	_default  = "default"
	_in       = "in"
	_enum     = "enum"
	_notIn    = "not_in"

	// string tag
	_minLen   = "min_len"
	_maxLen   = "max_len"
	_prefix   = "prefix"
	_suffix   = "suffix"
	_contains = "contains"
	_number   = "number"
	_email    = "email"
	_ip       = "ip"
	_ipv4     = "ipv4"
	_ipv6     = "ipv6"
	_crontab  = "crontab"
	_uuid     = "uuid"
	_uri      = "uri"
	_domain   = "domain"
	_pattern  = "pattern"

	// int32, int64, uint32, uint64, float32, float64 tag
	_ne  = "ne"
	_eq  = "eq"
	_lt  = "lt"
	_lte = "lte"
	_gt  = "gt"
	_gte = "gte"

	// bytes tag
	_maxBytes = "max_bytes"
	_minBytes = "min_bytes"

	// repeated tag: required, min_len, max_len
	// message tag: required
)

type Tag struct {
	Key   string
	Value string
}

// Paths for packages used by code generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
	isPkgPath      = "github.com/lack-io/vine/util/is"
	stringsPkgPath = "strings"
)

// dao is an implementation of the Go protocol buffer compiler's
// plugin architecture. It generates bindings for dao support.
type dao struct {
	gen *generator.Generator
}

func New() *dao {
	return &dao{}
}

// Name returns the name of this plugin, "dao".
func (g *dao) Name() string {
	return "dao"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.
var (
	isPkg      string
	stringsPkg string
	pkgImports map[generator.GoPackageName]bool
)

// Init initializes the plugin.
func (g *dao) Init(gen *generator.Generator) {
	g.gen = gen
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *dao) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *dao) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *dao) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *dao) Generate(file *generator.FileDescriptor) {
	if len(file.Comments()) == 0 {
		return
	}

	isPkg = string(g.gen.AddImport(isPkgPath))
	stringsPkg = string(g.gen.AddImport(stringsPkgPath))

	g.P("// Reference imports to suppress errors if they are not otherwise used.")
	g.P("var _ ", isPkg, ".Empty")
	g.P("var _ ", stringsPkg, ".Builder")
	g.P()
	for i, msg := range file.Messages() {
		g.generateMessage(file, msg, i)
	}
}

// GenerateImports generates the import declaration for this file.
func (g *dao) GenerateImports(file *generator.FileDescriptor, imports map[generator.GoImportPath]generator.GoPackageName) {
	if len(file.Comments()) == 0 {
		return
	}

	// We need to keep track of imported packages to make sure we don't produce
	// a name collision when generating types.
	pkgImports = make(map[generator.GoPackageName]bool)
	for _, name := range imports {
		pkgImports[name] = true
	}
}

func (g *dao) generateMessage(file *generator.FileDescriptor, msg *generator.MessageDescriptor, index int) {
	if msg.Proto.Options != nil && *(msg.Proto.Options.MapEntry) {
		return
	}
	g.P("func (m *", msg.Proto.Name, ") Validate() error {")
	g.P(`return m.validate("")`)
	g.P("}")
	g.P()
	g.P("func (m *", msg.Proto.Name, ") validate(prefix string) error {")
	if g.ignoredMessage(msg) {
		g.P("return nil")
	} else {
		g.P("errs := make([]error, 0)")
		for _, field := range msg.Fields {
			if field.Proto.IsRepeated() {
				g.generateRepeatedField(field)
				continue
			}
			if field.Proto.IsMessage() {
				g.generateMessageField(field)
				continue
			}
			switch *field.Proto.Type {
			case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
				descriptor.FieldDescriptorProto_TYPE_FLOAT,
				descriptor.FieldDescriptorProto_TYPE_FIXED32,
				descriptor.FieldDescriptorProto_TYPE_FIXED64,
				descriptor.FieldDescriptorProto_TYPE_INT32,
				descriptor.FieldDescriptorProto_TYPE_INT64:
				g.generateNumberField(field)
			case descriptor.FieldDescriptorProto_TYPE_STRING:
				g.generateStringField(field)
			case descriptor.FieldDescriptorProto_TYPE_BYTES:
				g.generateBytesField(field)
			}
		}
		g.P(fmt.Sprintf("return %s.MargeErr(errs...)", isPkg))
	}
	g.P("}")
	g.P()
}

func (g *dao) generateNumberField(field *generator.FieldDescriptor) {
	fieldName := generator.CamelCase(*field.Proto.Name)
	tags := g.extractTags(field.Comments)
	if len(tags) == 0 {
		return
	}
	if _, ok := tags[_required]; ok {
		g.P("if int64(m.", fieldName, ") == 0 {")
		g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is required\", prefix))", *field.Proto.JsonName))
		if len(tags) > 1 {
			g.P("} else {")
		}
	} else {
		if tag, ok := tags[_default]; ok {
			g.P("if int64(m.", fieldName, ") == 0 {")
			g.P("m.", fieldName, " = ", tag.Value)
			g.P("}")
		}
		g.P("if int64(m.", fieldName, ") != 0 {")
	}
	for _, tag := range tags {
		switch tag.Key {
		case _enum, _in:
			value := strings.TrimPrefix(tag.Value, "[")
			value = strings.TrimSuffix(value, "]")
			g.P(fmt.Sprintf("if %s.In([]interface{}{%s}, m.%s) {", isPkg, value, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must in '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _notIn:
			value := strings.TrimPrefix(tag.Value, "[")
			value = strings.TrimSuffix(value, "]")
			g.P(fmt.Sprintf("if %s.NotIn([]interface{}{%s}, m.%s) {", isPkg, value, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must not in '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _eq:
			g.P("if !(m.", fieldName, " == ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must equal to '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _ne:
			g.P("if !(m.", fieldName, " != ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must not equal to '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _lt:
			g.P("if !(m.", fieldName, " < ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must less than '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _lte:
			g.P("if !(m.", fieldName, " <= ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must less than or equal to '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _gt:
			g.P("if !(m.", fieldName, " > ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must great than '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _gte:
			g.P("if !(m.", fieldName, " >= ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must great than or equal to '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		}
	}
	g.P("}")
}

func (g *dao) generateStringField(field *generator.FieldDescriptor) {
	fieldName := generator.CamelCase(*field.Proto.Name)
	tags := g.extractTags(field.Comments)
	if len(tags) == 0 {
		return
	}
	if _, ok := tags[_required]; ok {
		g.P("if len(m.", fieldName, ") == 0 {")
		g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is required\", prefix))", *field.Proto.JsonName))
		if len(tags) > 1 {
			g.P("} else {")
		}
	} else {
		if tag, ok := tags[_default]; ok {
			g.P("if len(m.", fieldName, ") == 0 {")
			g.P("m.", fieldName, " = ", tag.Value)
			g.P("}")
		}
		g.P("if len(m.", fieldName, ") != 0 {")
	}
	for _, tag := range tags {
		fieldName := generator.CamelCase(*field.Proto.Name)
		switch tag.Key {
		case _enum, _in:
			value := fullStringSlice(tag.Value)
			g.P(fmt.Sprintf("if %s.In([]string{%s}, m.%s) {", isPkg, value, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must in '[%s]'\", prefix))", *field.Proto.JsonName, strings.ReplaceAll(value, "\"", "")))
			g.P("}")
		case _notIn:
			value := fullStringSlice(tag.Value)
			g.P(fmt.Sprintf("if %s.NotIn([]string{%s}, m.%s) {", isPkg, value, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must not in '[%s]'\", prefix))", *field.Proto.JsonName, strings.ReplaceAll(value, "\"", "")))
			g.P("}")
		case _minLen:
			g.P("if !(len(m.", fieldName, ") >= ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' length must less than '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _maxLen:
			g.P("if !(len(m.", fieldName, ") <= ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' length must great than '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _prefix:
			value := TrimString(tag.Value, "\"")
			g.P("if !strings.HasPrefix(m.", fieldName, ", \"", value, "\") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must start with '%s'\", prefix))", *field.Proto.JsonName, value))
			g.P("}")
		case _suffix:
			value := TrimString(tag.Value, "\"")
			g.P("if !strings.HasSuffix(m.", fieldName, ", \"", value, "\") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must end with '%s'\", prefix))", *field.Proto.JsonName, value))
			g.P("}")
		case _contains:
			value := TrimString(tag.Value, "\"")
			g.P("if !strings.Contains(m.", fieldName, ", \"", value, "\") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' must contain '%s'\", prefix))", *field.Proto.JsonName, value))
			g.P("}")
		case _number:
			g.P(fmt.Sprintf("if !%s.Number(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid number\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _email:
			g.P(fmt.Sprintf("if !%s.Email(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid email\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _ip:
			g.P(fmt.Sprintf("if !%s.IP(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid ip\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _ipv4:
			g.P(fmt.Sprintf("if !%s.IPv4(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid ipv4\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _ipv6:
			g.P(fmt.Sprintf("if !%s.IPv6(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid ipv6\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _crontab:
			g.P(fmt.Sprintf("if !%s.Crontab(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid crontab\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _uuid:
			g.P(fmt.Sprintf("if !%s.Uuid(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid uuid\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _uri:
			g.P(fmt.Sprintf("if !%s.URL(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid url\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _domain:
			g.P(fmt.Sprintf("if !%s.Domain(m.%s) {", isPkg, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is not a valid domain\", prefix))", *field.Proto.JsonName))
			g.P("}")
		case _pattern:
			value := TrimString(tag.Value, "`")
			g.P(fmt.Sprintf("if !%s.Re(`%s`, m.%s) {", isPkg, value, fieldName))
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(`field '%%s%s' is not a valid pattern '%s'`, prefix))", *field.Proto.JsonName, value))
			g.P("}")
		}
	}
	g.P("}")
}

func (g *dao) generateBytesField(field *generator.FieldDescriptor) {
	fieldName := generator.CamelCase(*field.Proto.Name)
	tags := g.extractTags(field.Comments)
	if len(tags) == 0 {
		return
	}
	if _, ok := tags[_required]; ok {
		g.P("if len(m.", fieldName, ") == 0 {")
		g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is required\", prefix))", *field.Proto.JsonName))
		if len(tags) > 1 {
			g.P("} else {")
		}
	} else {
		g.P("if len(m.", fieldName, ") != 0 {")
	}
	for _, tag := range tags {
		fieldName := generator.CamelCase(*field.Proto.Name)
		switch tag.Key {
		case _minBytes:
			g.P("if !(len(m.", fieldName, ") <= ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' length must less than '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _maxBytes:
			g.P("if !(len(m.", fieldName, ") >= ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' length must great than '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		}
	}
	g.P("}")
}

func (g *dao) generateRepeatedField(field *generator.FieldDescriptor) {
	fieldName := generator.CamelCase(*field.Proto.Name)
	tags := g.extractTags(field.Comments)
	if len(tags) == 0 {
		return
	}
	if _, ok := tags[_required]; ok {
		g.P("if len(m.", fieldName, ") == 0 {")
		g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is required\", prefix))", *field.Proto.JsonName))
		if len(tags) > 1 {
			g.P("} else {")
		}
	} else {
		g.P("if len(m.", fieldName, ") != 0 {")
	}
	for _, tag := range tags {
		fieldName := generator.CamelCase(*field.Proto.Name)
		switch tag.Key {
		case _minLen:
			g.P("if !(len(m.", fieldName, ") <= ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' length must less than '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		case _maxLen:
			g.P("if !(len(m.", fieldName, ") >= ", tag.Value, ") {")
			g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' length must great than '%s'\", prefix))", *field.Proto.JsonName, tag.Value))
			g.P("}")
		}
	}
	g.P("}")
}

func (g *dao) generateMessageField(field *generator.FieldDescriptor) {
	fieldName := generator.CamelCase(*field.Proto.Name)
	tags := g.extractTags(field.Comments)
	if len(tags) == 0 {
		return
	}
	if _, ok := tags[_required]; ok {
		g.P("if m.", fieldName, " == nil {")
		g.P(fmt.Sprintf("errs = append(errs, fmt.Errorf(\"field '%%s%s' is required\", prefix))", *field.Proto.JsonName))
		g.P("} else {")
		g.P(fmt.Sprintf("errs = append(errs, m.%s.validate(prefix+\"%s.\"))", fieldName, *field.Proto.JsonName))
		g.P("}")
	}
}

func (g *dao) ignoredMessage(msg *generator.MessageDescriptor) bool {
	tags := g.extractTags(msg.Comments)
	for _, c := range tags {
		if c.Key == _ignore {
			return true
		}
	}
	return false
}

func TrimString(s string, c string) string {
	s = strings.TrimPrefix(s, c)
	s = strings.TrimSuffix(s, c)
	return s
}

func fullStringSlice(s string) string {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	parts := strings.Split(s, ",")
	out := make([]string, 0)
	for _, a := range parts {
		a = strings.TrimSpace(a)
		if len(a) == 0 {
			continue
		}
		if !strings.HasPrefix(a, "\"") {
			a = "\"" + a
		}
		if !strings.HasSuffix(a, "\"") {
			a = a + "\""
		}
		out = append(out, a)
	}
	return strings.Join(out, ",")
}

func (g *dao) extractTags(comments []*generator.Comment) map[string]*Tag {
	if comments == nil || len(comments) == 0 {
		return nil
	}
	tags := make(map[string]*Tag, 0)
	for _, c := range comments {
		if c.Tag != TagString || len(c.Text) == 0 {
			continue
		}
		if strings.HasPrefix(c.Text, _pattern) {
			if i := strings.Index(c.Text, "="); i == -1 {
				g.gen.Fail("invalid pattern format")
			} else {
				pa := string(c.Text[i+1])
				pe := string(c.Text[len(c.Text)-1])
				if pa != "`" || pe != "`" {
					g.gen.Fail(fmt.Sprintf("invalid pattern value, pa=%s, pe=%s", pa, pe))
				}
				key := strings.TrimSpace(c.Text[:i])
				value := strings.TrimSpace(c.Text[i+1:])
				if len(value) == 0 {
					g.gen.Fail(fmt.Sprintf("tag '%s' missing value", key))
				}
				tags[key] = &Tag{
					Key:   key,
					Value: value,
				}
			}
			continue
		}
		parts := strings.Split(c.Text, ";")
		for _, p := range parts {
			tag := new(Tag)
			p = strings.TrimSpace(p)
			if i := strings.Index(p, "="); i > 0 {
				tag.Key = strings.TrimSpace(p[:i])
				v := strings.TrimSpace(p[i+1:])
				if v == "" {
					g.gen.Fail(fmt.Sprintf("tag '%s' missing value", tag.Key))
				}
				tag.Value = v
			} else {
				tag.Key = p
			}
			tags[tag.Key] = tag
		}
	}

	return tags
}