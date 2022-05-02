/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"time"

	esmodel "github.com/devicechain-io/dc-event-sources/model"
)

// Event with token references resolved and info from assignment merged.
type Event struct {
	Source          string
	AltId           *string
	DeviceId        uint `gorm:"not null"`
	AssignmentId    uint `gorm:"not null"`
	DeviceGroupId   *uint
	CustomerId      *uint
	CustomerGroupId *uint
	AreaId          *uint
	AreaGroupId     *uint
	AssetId         *uint
	AssetGroupId    *uint
	OccurredTime    time.Time `gorm:"not null"`
	ProcessedTime   time.Time
	EventType       esmodel.EventType
}
