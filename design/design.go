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

var TablePayload = ResultType("TablePayload", func() {
	Extend(Table)
	Field(3, "description", String, "Describes the table being returned")
})


var _ = API("Table Info Service", func() {
	Title("Table Info Service")
	Description("Goa test service")

	Server("Table Info Service", func() {
		Services("TableInfo")
		Host("0.0.0.0", func() {
			URI("grpc://0.0.0.0:8000")
		})
	})
})

var _ = Service("TableInfo", func() {
	Method("get_info", func() {
		Payload(func() {
			Field(1, "tableName", String, "Table name", func() {
				Example("myGreatTable")
			})
			Required("tableName")
		})
		Result(TablePayload)
		GRPC(func() {
			Response(CodeOK)
		})
	})
})
