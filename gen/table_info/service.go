// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// TableInfo service
//
// Command:
// $ goa gen github.com/jamie-doyle/goa-test/design

package tableinfo

import (
	"context"

	tableinfoviews "github.com/jamie-doyle/goa-test/gen/table_info/views"
)

// Service is the TableInfo service interface.
type Service interface {
	// GetInfo implements get_info.
	GetInfo(context.Context, *GetInfoPayload) (res *Tablepayload, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "TableInfo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"get_info"}

// GetInfoPayload is the payload type of the TableInfo service get_info method.
type GetInfoPayload struct {
	// Table name
	TableName string
}

// Tablepayload is the result type of the TableInfo service get_info method.
type Tablepayload struct {
	// Describes the table being returned
	Description *string
	// A table ID
	TableID string
}

// NewTablepayload initializes result type Tablepayload from viewed result type
// Tablepayload.
func NewTablepayload(vres *tableinfoviews.Tablepayload) *Tablepayload {
	var res *Tablepayload
	switch vres.View {
	case "default", "":
		res = newTablepayload(vres.Projected)
	}
	return res
}

// NewViewedTablepayload initializes viewed result type Tablepayload from
// result type Tablepayload using the given view.
func NewViewedTablepayload(res *Tablepayload, view string) *tableinfoviews.Tablepayload {
	var vres *tableinfoviews.Tablepayload
	switch view {
	case "default", "":
		p := newTablepayloadView(res)
		vres = &tableinfoviews.Tablepayload{p, "default"}
	}
	return vres
}

// newTablepayload converts projected type Tablepayload to service type
// Tablepayload.
func newTablepayload(vres *tableinfoviews.TablepayloadView) *Tablepayload {
	res := &Tablepayload{
		Description: vres.Description,
	}
	return res
}

// newTablepayloadView projects result type Tablepayload to projected type
// TablepayloadView using the "default" view.
func newTablepayloadView(res *Tablepayload) *tableinfoviews.TablepayloadView {
	vres := &tableinfoviews.TablepayloadView{
		Description: res.Description,
	}
	return vres
}
