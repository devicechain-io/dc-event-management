/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"database/sql"
	"time"

	esmodel "github.com/devicechain-io/dc-event-sources/model"
)

// Event with token references resolved and info from assignment merged.
type Event struct {
	DeviceId        uint
	EventType       esmodel.EventType
	OccurredTime    time.Time
	Source          string
	AltId           sql.NullString
	AssignmentId    uint
	DeviceGroupId   *uint
	CustomerId      *uint
	CustomerGroupId *uint
	AreaId          *uint
	AreaGroupId     *uint
	AssetId         *uint
	AssetGroupId    *uint
	ProcessedTime   time.Time
}

// Location event fields.
type LocationEvent struct {
	DeviceId     uint              `gorm:"not null"`
	EventType    esmodel.EventType `gorm:"not null"`
	OccurredTime time.Time         `gorm:"not null"`
	Event        Event             `gorm:"foreignKey:DeviceId,EventType,OccurredTime;References:DeviceId,EventType,OccurredTime"`
	Latitude     sql.NullFloat64   `gorm:"type:decimal(10,8);"`
	Longitude    sql.NullFloat64   `gorm:"type:decimal(11,8);"`
	Elevation    sql.NullFloat64   `gorm:"type:decimal(10,8);"`
}

// Information required to create a location event.
type LocationEventCreateRequest struct {
	Event
	Latitude  *float64
	Longitude *float64
	Elevation *float64
}

// Measurement event fields.
type MeasurementEvent struct {
	DeviceId     uint            `gorm:"not null"`
	OccurredTime time.Time       `gorm:"not null"`
	Event        Event           `gorm:"foreignKey:DeviceId,OccurredTime;References:DeviceId,OccurredTime"`
	Latitude     sql.NullFloat64 `gorm:"type:decimal(10,8);"`
	Longitude    sql.NullFloat64 `gorm:"type:decimal(11,8);"`
	Elevation    sql.NullFloat64 `gorm:"type:decimal(10,8);"`
}

// Information required to create a measurement event.
type MeasurementEventCreateRequest struct {
	Event
	Latitude  *float64
	Longitude *float64
	Elevation *float64
}
