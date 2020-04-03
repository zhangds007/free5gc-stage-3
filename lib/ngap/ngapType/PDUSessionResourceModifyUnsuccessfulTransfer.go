//go:binary-only-package

package ngapType

// Need to import "free5gc/lib/aper" if it uses "aper"

type PDUSessionResourceModifyUnsuccessfulTransfer struct {
	Cause                  Cause                                                                         `aper:"valueLB:0,valueUB:5"`
	CriticalityDiagnostics *CriticalityDiagnostics                                                       `aper:"valueExt,optional"`
	IEExtensions           *ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs `aper:"optional"`
}
