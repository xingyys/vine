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

package nop

import (
	"github.com/lack-io/vine/service/dao"
	"github.com/lack-io/vine/service/dao/clause"
	"github.com/lack-io/vine/service/dao/schema"
)

type noopDialect struct {
	opts dao.Options
}

func (d noopDialect) Init(options dao.Options) error {
	d.opts = options
	return nil
}

func (d noopDialect) Options() dao.Options {
	return d.opts
}

func (d noopDialect) NewDB() (*dao.DB, error) {
	return &dao.DB{}, nil
}

func (noopDialect) DefaultValueOf(field *schema.Field) clause.Expression {
	return clause.Expr{SQL: "DEFAULT"}
}

func (noopDialect) Migrator(*dao.DB) dao.Migrator {
	return nil
}

func (noopDialect) BindVarTo(writer clause.Writer, stmt *dao.Statement, v interface{}) {
	writer.WriteByte('?')
}

func (noopDialect) QuoteTo(writer clause.Writer, str string) {
	writer.WriteByte('`')
	writer.WriteString(str)
	writer.WriteByte('`')
}

func (noopDialect) Explain(sql string, vars ...interface{}) string {
	return ""
}

func (noopDialect) String() string {
	return "dummy"
}

func (noopDialect) DataTypeOf(*schema.Field) string {
	return ""
}

func NewDialect(opts ...dao.Option) dao.Dialect {
	return &noopDialect{}
}
