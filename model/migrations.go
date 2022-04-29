/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
)

var (
	Migrations = []*gormigrate.Migration{
		NewInitialSchema(),
	}
)
