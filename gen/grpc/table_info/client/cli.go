// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// TableInfo gRPC client CLI support package
//
// Command:
// $ goa gen github.com/jamie-doyle/goa-test/design

package client

import (
	"encoding/json"
	"fmt"

	table_infopb "github.com/jamie-doyle/goa-test/gen/grpc/table_info/pb"
	tableinfo "github.com/jamie-doyle/goa-test/gen/table_info"
)

// BuildGetInfoPayload builds the payload for the TableInfo get_info endpoint
// from CLI flags.
func BuildGetInfoPayload(tableInfoGetInfoMessage string) (*tableinfo.GetInfoPayload, error) {
	var err error
	var message table_infopb.GetInfoRequest
	{
		if tableInfoGetInfoMessage != "" {
			err = json.Unmarshal([]byte(tableInfoGetInfoMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"tableName\": \"myGreatTable\"\n   }'")
			}
		}
	}
	v := &tableinfo.GetInfoPayload{
		TableName: message.TableName,
	}
	return v, nil
}
