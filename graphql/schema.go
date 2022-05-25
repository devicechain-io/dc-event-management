/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	_ "embed"

	gqlcore "github.com/devicechain-io/dc-microservice/graphql"
	"github.com/devicechain-io/dc-microservice/rdb"
)

//go:embed schema.graphql
var SchemaContent string

type SchemaResolver struct{}

// Get rdb manager from context.
func (s *SchemaResolver) GetRdbManager(ctx context.Context) *rdb.RdbManager {
	return ctx.Value(gqlcore.ContextRdbKey).(*rdb.RdbManager)
}
