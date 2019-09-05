package vlp16

import "time"

const PointsPerCloud = BlocksPerPacket * ChannelsPerBlock

// TODO: Clever allocation of memory

type SphericalPointCloud struct {
	Timestamp       time.Duration
	DeltaTimestamp  time.Duration
	DeltaAzimuth    float32
	DeltaElevation  float32
	LowestElevation float32
	Columns         []Column
}

type Column struct {
	TimestampOffset time.Duration
	Azimuth         float32
	Rows            []Row
}

type Row struct {
	SphericalPoints []SphericalPoint
}

type SphericalPoint struct {
	Radius           float32
	Reflectivity     uint8
	IsLastReflection bool
}

func (s *SphericalPointCloud) UnmarshalPacket(packet *Packet) {
	// duration is in nanoseconds and Velodyne timestamp in microseconds
	s.Timestamp = time.Duration(packet.Timestamp) * time.Microsecond
	s.DeltaTimestamp = time.Duration(int(singleFiringTime*1000)) * time.Millisecond
	// TODO: Implement precision azimuth calculations
	s.DeltaAzimuth = 0
	s.DeltaElevation = deltaElevation
	s.LowestElevation = lowestElevation
	for i := range packet.Blocks {
		block := &packet.Blocks[i]
		for j := range block.Channels {
			channel := &block.Channels[j]
			columnIndex, rowIndex, pointIndex := indices(i, j, packet.ReturnMode)
			switch j {
			case 0:
				s.Columns[columnIndex].Azimuth = deg2Rad(float32(block.Azimuth) * azimuthFactor)
			case 16:
				s.Columns[columnIndex].Azimuth =
					deg2Rad(float32(interpolateAzimuth(i, packet)) * azimuthFactor)
			}
			s.Columns[columnIndex].Rows[rowIndex].SphericalPoints[pointIndex] := &SphericalPoint{
				Radius:       float32(channel.Distance) * distanceFactor,
				Reflectivity: packet.Blocks[i].Channels[j].Reflectivity,
				IsLastReflection: packet.ReturnMode == ReturnModeLastReturn ||
					// dual return mode: even number blocks (0,2,4,...) contain last return
					packet.ReturnMode == ReturnModeDualReturn && i%2 == 0,
			}
		}
	}
}

func indices(blockIndex int, channelIndex int, returnMode ReturnMode) (int, int, int) {
	if returnMode == ReturnModeDualReturn {
		return blockIndex - (blockIndex % 2) + (channelIndex / numberOfRows),
			rowIndex(channelIndex),
			channelIndex % numberOfRows
	} else {
		return blockIndex*2 + channelIndex/numberOfRows, rowIndex(channelIndex), 0
	}
}

func rowIndex(channelID int) int {
	if channelID > numberOfRows-1 { // Account for second firing
		channelID -= numberOfRows
	}
	return channelID/2 + channelID%2*8
}
