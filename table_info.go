package tableinfoservice

import (
	"context"
	"log"

	tableinfo "github.com/jamie-doyle/goa-test/gen/table_info"
)

// TableInfo service example implementation.
// The example methods log the requests and return zero values.
type tableInfosrvc struct {
	logger *log.Logger
}

// NewTableInfo returns the TableInfo service implementation.
func NewTableInfo(logger *log.Logger) tableinfo.Service {
	return &tableInfosrvc{logger}
}

// GetInfo implements get_info.
func (s *tableInfosrvc) GetInfo(ctx context.Context, p *tableinfo.GetInfoPayload) (res *tableinfo.Tablepayload, err error) {
	desc := "This is returned 😁"

	res = &tableinfo.Tablepayload{
		TableID:     "This isn't returned 😢",
		Description: &desc,
	}
	s.logger.Print("tableInfo.get_info")
	return
}
