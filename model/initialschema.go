/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"time"

	esmodel "github.com/devicechain-io/dc-event-sources/model"
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Creates the initial schema migration for this functional area.
func NewInitialSchema() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220420000000",
		Migrate: func(tx *gorm.DB) error {
			// Location event fields.
			type LocationEvent struct {
				DeviceId     uint      `gorm:"not null"`
				OccurredTime time.Time `gorm:"not null"`
				Event        Event     `gorm:"foreignKey:DeviceId,OccurredTime;References:DeviceId,OccurredTime"`
				Latitude     float64   `gorm:"type:decimal(10,8);"`
				Longitude    float64   `gorm:"type:decimal(11,8);"`
				Elevation    float64   `gorm:"type:decimal(10,8);"`
			}

			// Base event fields.
			type Event struct {
				DeviceId        uint      `gorm:"primaryKey"`
				OccurredTime    time.Time `gorm:"primaryKey"`
				Source          string
				AltId           *string
				AssignmentId    uint `gorm:"not null"`
				DeviceGroupId   *uint
				CustomerId      *uint
				CustomerGroupId *uint
				AreaId          *uint
				AreaGroupId     *uint
				AssetId         *uint
				AssetGroupId    *uint
				ProcessedTime   time.Time
				EventType       esmodel.EventType `gorm:"not null"`
			}

			err := tx.AutoMigrate(&Event{}, &LocationEvent{})
			if err != nil {
				return err
			}

			// Convert to a hypertable.
			err = tx.Raw("SELECT create_hypertable('event-management.events', 'occurred_time');").Row().Err()
			if err != nil {
				return err
			}

			// Add index on device id.
			tx.Exec("CREATE INDEX ON \"event-management\".\"events\" (device_id, occurred_time DESC);")
			if tx.Error != nil {
				return tx.Error
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
