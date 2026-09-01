// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type PDUSessionReleaseCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONRELEASECOMMANDMessageIdentity
	nasType.Cause5GSM
	*nasType.BackoffTimerValue
	*nasType.EAPMessage
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.Fivegsmcongestionreattemptindicator
	*nasType.SpareHalfOctetAndAccessType
	*nasType.ServiceLevelAAContainer
	AlternativeSNSSAI *nasType.SNSSAI
}

func NewPDUSessionReleaseCommand(iei uint8) (pDUSessionReleaseCommand *PDUSessionReleaseCommand) {
	pDUSessionReleaseCommand = &PDUSessionReleaseCommand{}
	return pDUSessionReleaseCommand
}

const (
	PDUSessionReleaseCommandBackoffTimerValueType                    uint8 = 0x37
	PDUSessionReleaseCommandEAPMessageType                           uint8 = 0x78
	PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionReleaseCommandFivegsmcongestionreattemptindicatorType  uint8 = 0x61
	PDUSessionReleaseCommandAccessTypeType                           uint8 = 0x0D
	PDUSessionReleaseCommandServiceLevelAAContainerType              uint8 = 0x72
	PDUSessionReleaseCommandAlternativeSNSSAIType                    uint8 = 0x5A
)

func (a *PDUSessionReleaseCommand) EncodePDUSessionReleaseCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
	if a.BackoffTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer)
	}
	if a.Fivegsmcongestionreattemptindicator != nil {
		binary.Write(buffer, binary.BigEndian, a.Fivegsmcongestionreattemptindicator.GetIei())
		binary.Write(buffer, binary.BigEndian, a.Fivegsmcongestionreattemptindicator.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Octet)
	}
	if a.SpareHalfOctetAndAccessType != nil {
		binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndAccessType.Octet)
	}
	if a.ServiceLevelAAContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer)
	}
	if a.AlternativeSNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.Octet[:a.AlternativeSNSSAI.GetLen()])
	}
}

func (a *PDUSessionReleaseCommand) DecodePDUSessionReleaseCommand(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
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
		case PDUSessionReleaseCommandBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len)
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
		case PDUSessionReleaseCommandEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		case PDUSessionReleaseCommandFivegsmcongestionreattemptindicatorType:
			a.Fivegsmcongestionreattemptindicator = nasType.NewFivegsmcongestionreattemptindicator(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Len)
			a.Fivegsmcongestionreattemptindicator.SetLen(a.Fivegsmcongestionreattemptindicator.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Octet)
		case PDUSessionReleaseCommandAccessTypeType:
			a.SpareHalfOctetAndAccessType = nasType.NewSpareHalfOctetAndAccessType()
			a.SpareHalfOctetAndAccessType.Octet = ieiN
		case PDUSessionReleaseCommandServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len)
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()])
		case PDUSessionReleaseCommandAlternativeSNSSAIType:
			a.AlternativeSNSSAI = nasType.NewSNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AlternativeSNSSAI.Len)
			a.AlternativeSNSSAI.SetLen(a.AlternativeSNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AlternativeSNSSAI.Octet[:a.AlternativeSNSSAI.GetLen()])
		default:
		}
	}
}
