// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/omec-project/nas/v2/nasType"
)

type SecurityModeComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.SecurityModeCompleteMessageIdentity
	*nasType.IMEISV
	*nasType.NASMessageContainer
	NonIMEISVPEI *nasType.MobileIdentity5GS
}

func NewSecurityModeComplete(iei uint8) (securityModeComplete *SecurityModeComplete) {
	securityModeComplete = &SecurityModeComplete{}
	return securityModeComplete
}

const (
	SecurityModeCompleteIMEISVType              uint8 = 0x77
	SecurityModeCompleteNASMessageContainerType uint8 = 0x71
	SecurityModeCompleteNonIMEISVPEIType        uint8 = 0x78
)

func (a *SecurityModeComplete) EncodeSecurityModeComplete(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SecurityModeCompleteMessageIdentity.Octet)
	if a.IMEISV != nil {
		binary.Write(buffer, binary.BigEndian, a.IMEISV.GetIei())
		binary.Write(buffer, binary.BigEndian, a.IMEISV.GetLen())
		binary.Write(buffer, binary.BigEndian, a.IMEISV.Octet[:a.IMEISV.GetLen()])
	}
	if a.NASMessageContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer)
	}
	if a.NonIMEISVPEI != nil {
		binary.Write(buffer, binary.BigEndian, a.NonIMEISVPEI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NonIMEISVPEI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NonIMEISVPEI.Buffer)
	}
}

func (a *SecurityModeComplete) DecodeSecurityModeComplete(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SecurityModeCompleteMessageIdentity.Octet)
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		binary.Read(buffer, binary.BigEndian, &ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		switch tmpIeiN {
		case SecurityModeCompleteIMEISVType:
			a.IMEISV = nasType.NewIMEISV(ieiN)
			if a.IMEISV != nil {
				binary.Read(buffer, binary.BigEndian, &a.IMEISV.Len)
				a.IMEISV.SetLen(a.IMEISV.GetLen())
				binary.Read(buffer, binary.BigEndian, a.IMEISV.Octet[:a.IMEISV.GetLen()])
			} else {
				return fmt.Errorf("IMEISV is nil")
			}
		case SecurityModeCompleteNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len)
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer[:a.NASMessageContainer.GetLen()])
		case SecurityModeCompleteNonIMEISVPEIType:
			a.NonIMEISVPEI = nasType.NewMobileIdentity5GS(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NonIMEISVPEI.Len)
			a.NonIMEISVPEI.SetLen(a.NonIMEISVPEI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NonIMEISVPEI.Buffer[:a.NonIMEISVPEI.GetLen()])
		default:
		}
	}
	return nil
}
