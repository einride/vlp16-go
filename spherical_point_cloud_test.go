package vlp16

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSphericalPointCloud_UnmarshalExamplePacket(t *testing.T) {
	actual := &PointCloud{}
	actual.UnmarshalPacket(examplePacket())
	require.True(t, isEqual(examplePointCloud(), actual))
}

func isEqual(p *PointCloud, pc *PointCloud) bool {
	delta := 1e-5
	if len(p.Points) != len(pc.Points) || len(p.Azimuths) != len(pc.Azimuths) {
		return false
	}
	isEqual := true
	isEqual = isEqual && (p.TimeSinceTopOfHour == pc.TimeSinceTopOfHour)
	for i := range p.Azimuths {
		isEqual = isEqual && (math.Abs(p.Azimuths[i].Radians()-pc.Azimuths[i].Radians()) < delta)
	}
	for i := range p.Points {
		isEqual = isEqual &&
			(math.Abs(p.Points[i].Distance.Metres()-pc.Points[i].Distance.Metres()) < delta) &&
			p.Points[i].Column == pc.Points[i].Column &&
			p.Points[i].Row == pc.Points[i].Row &&
			p.Points[i].Reflectivity == pc.Points[i].Reflectivity &&
			p.Points[i].IsLastReflection == pc.Points[i].IsLastReflection
	}
	return isEqual
}
