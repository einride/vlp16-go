package vlp16

import (
	"time"
)

const (
	// ColumnsPerPacket is the number of columns in point cloud matrix.
	ColumnsPerPacket = BlocksPerPacket * 2
	// RowsPerColumnSingleReturn is the number of rows in point cloud in single return mode.
	RowsPerColumnSingleReturn = 16
	// FullFiringTime is the total time for laser firings plus recharge (55.296 µs).
	FullFiringTime = 55296 * time.Nanosecond
	// SingleFiringTime is the time for one laser firing (2.304 µs).
	SingleFiringTime = 2304 * time.Nanosecond
	// RechargeTime is the recharge time between laser firings (18.432 µs).
	RechargeTime = 18432 * time.Nanosecond
	// LowestElevation is the elevation angle of first row of measurements.
	LowestElevationDegrees = -15
	// DeltaElevation is the angle difference between two rows.
	DeltaElevationDegrees = 2
)

// compile-time assertion on full firing time.
var _ [FullFiringTime]struct{} = [RowsPerColumnSingleReturn*SingleFiringTime + RechargeTime]struct{}{}
