/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package config

import (
	"github.com/devicechain-io/dc-microservice/config"
)

const (
	KAFKA_TOPIC_PERSISTED_EVENTS = "persisted-events"
)

type EventManagementConfiguration struct {
	RdbConfiguration config.MicroserviceRdbConfiguration
}

// Creates the default device management configuration
func NewDeviceManagementConfiguration() *EventManagementConfiguration {
	return &EventManagementConfiguration{
		RdbConfiguration: config.MicroserviceRdbConfiguration{
			SqlDebug: true,
		},
	}
}
