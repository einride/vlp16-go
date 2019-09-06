package vlp16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimulateRead(t *testing.T) {
	sc, done := newPacketRecordingScanner(t, "testdata/recording.bin")
	defer done()
	for sc.Scan() {
		var data [lengthOfPacket]byte
		copy(data[:], sc.Bytes())
		var packet Packet
		packet.unmarshal(&data)
		var cloud PointCloud
		cloud.UnmarshalPacket(&packet)
		for i := range cloud.Points {
			// all points in test data are strongest return
			require.Equal(t, cloud.Points[i].LastReflection, false)
		}
	}
}

func TestSphericalPointCloud_UnmarshalExamplePacket(t *testing.T) {
	actual := &PointCloud{}
	actual.UnmarshalPacket(examplePacket())
	require.Equal(t, exampleSphericalPointCloud(), actual)
}

func TestSphericalPointCloud_UnmarshalExamplePacketLastReflection(t *testing.T) {
	actual := &PointCloud{}
	actual.UnmarshalPacket(examplePacketLastReflection())
	require.Equal(t, exampleSphericalPointCloudLastReflection(), actual)
}

func TestLastReflection(t *testing.T) {
	var cloud PointCloud
	cloud.UnmarshalPacket(examplePacketLastReflection())
	require.Equal(t, cloud.Points[1].LastReflection, true)
	require.Equal(t, cloud.Points[35].LastReflection, false)
}

func TestTimingOffset(t *testing.T) {
	const eps = 0.005
	var cloud PointCloud
	cloud.UnmarshalPacket(examplePacketLastReflection())
	require.InDelta(t, 2.304, cloud.Points[1].TimeOffset, eps)
	require.InDelta(t, 34.560, cloud.Points[15].TimeOffset, eps)
	require.InDelta(t, 89.856, cloud.Points[31].TimeOffset, eps)
	require.InDelta(t, 642.816, cloud.Points[len(cloud.Points)-1].TimeOffset, eps)
	cloud.UnmarshalPacket(examplePacket())
	require.InDelta(t, 2.304, cloud.Points[1].TimeOffset, eps)
	require.InDelta(t, 34.560, cloud.Points[15].TimeOffset, eps)
	require.InDelta(t, 89.856, cloud.Points[31].TimeOffset, eps)
	require.InDelta(t, 1306.37, cloud.Points[len(cloud.Points)-1].TimeOffset, eps)
}
