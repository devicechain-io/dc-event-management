/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package test

import (
	"context"
	"time"

	emmodel "github.com/devicechain-io/dc-event-management/model"
	"github.com/devicechain-io/dc-microservice/config"
	"github.com/devicechain-io/dc-microservice/core"
	"github.com/stretchr/testify/mock"
)

// Microservice information used for testing.
var EventManagementMicroservice = &core.Microservice{
	StartTime:                    time.Now(),
	InstanceId:                   "devicechain",
	TenantId:                     "tenant1",
	TenantName:                   "Tenant 1",
	MicroserviceId:               "event-management",
	MicroserviceName:             "Event Management",
	FunctionalArea:               "event-management",
	InstanceConfiguration:        config.InstanceConfiguration{},
	MicroserviceConfigurationRaw: make([]byte, 0),
}

/**
 * Mock for event management API.
 */

type MockApi struct {
	mock.Mock
}

func (api *MockApi) CreateLocationEvent(ctx context.Context, request *emmodel.LocationEventCreateRequest) (*emmodel.LocationEvent, error) {
	args := api.Mock.Called()
	return args.Get(0).(*emmodel.LocationEvent), args.Error(1)
}
