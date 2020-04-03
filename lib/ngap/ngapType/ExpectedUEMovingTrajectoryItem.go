//go:binary-only-package

package ngapType

// Need to import "free5gc/lib/aper" if it uses "aper"

type ExpectedUEMovingTrajectoryItem struct {
	NGRANCGI         NGRANCGI                                                        `aper:"valueLB:0,valueUB:2"`
	TimeStayedInCell *int64                                                          `aper:"valueLB:0,valueUB:4095,optional"`
	IEExtensions     *ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs `aper:"optional"`
}
