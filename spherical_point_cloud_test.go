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
		var cloud SphericalPointCloud
		cloud.UnmarshalPacket(&packet)
		for i := range cloud.SphericalPoints {
			// all points in test data are strongest return
			require.Equal(t, cloud.SphericalPoints[i].LastReflection, false)
		}
	}
}

func TestSphericalPointCloud_UnmarshalExamplePacket(t *testing.T) {
	actual := &SphericalPointCloud{}
	actual.UnmarshalPacket(examplePacket())
	require.Equal(t, exampleSphericalPointCloud(), actual)
}

func TestSphericalPointCloud_UnmarshalExamplePacketLastReflection(t *testing.T) {
	actual := &SphericalPointCloud{}
	actual.UnmarshalPacket(examplePacketLastReflection())
	require.Equal(t, exampleSphericalPointCloudLastReflection(), actual)
}

func TestLastReflection(t *testing.T) {
	var cloud SphericalPointCloud
	cloud.UnmarshalPacket(examplePacketLastReflection())
	require.Equal(t, cloud.SphericalPoints[1].LastReflection, true)
	require.Equal(t, cloud.SphericalPoints[35].LastReflection, false)
}

func TestTimingOffset(t *testing.T) {
	const eps = 0.005
	var cloud SphericalPointCloud
	cloud.UnmarshalPacket(examplePacketLastReflection())
	require.InDelta(t, 2.304, cloud.SphericalPoints[1].TimingOffset, eps)
	require.InDelta(t, 34.560, cloud.SphericalPoints[15].TimingOffset, eps)
	require.InDelta(t, 89.856, cloud.SphericalPoints[31].TimingOffset, eps)
	require.InDelta(t, 642.816, cloud.SphericalPoints[len(cloud.SphericalPoints)-1].TimingOffset, eps)
	cloud.UnmarshalPacket(examplePacket())
	require.InDelta(t, 2.304, cloud.SphericalPoints[1].TimingOffset, eps)
	require.InDelta(t, 34.560, cloud.SphericalPoints[15].TimingOffset, eps)
	require.InDelta(t, 89.856, cloud.SphericalPoints[31].TimingOffset, eps)
	require.InDelta(t, 1306.37, cloud.SphericalPoints[len(cloud.SphericalPoints)-1].TimingOffset, eps)
}
