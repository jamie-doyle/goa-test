// Package design ...
package design

import (
	. "goa.design/goa/dsl"
)

var Table = Type("Table", func() {
	Description("Describes a table")
	Field(1, "table_id", String, "A table ID", func() {
		Example("some-table-id")
	})
	Required("table_id")
})
