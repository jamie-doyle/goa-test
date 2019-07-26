// Package design ...
package design

import (
	. "github.com/jamie-doyle/goa-test/common/design"
	. "goa.design/goa/dsl"
)

var TablePayload = ResultType("TablePayload", func() {
	Extend(Table)
	Field(3, "description", String, "Describes the table being returned")
})
