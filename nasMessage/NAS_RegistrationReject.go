// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type RegistrationReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationRejectMessageIdentity
	nasType.Cause5GMM
	*nasType.T3346Value
	*nasType.T3502Value
	*nasType.EAPMessage
	*nasType.RejectedNSSAI
	*nasType.CAGInformationList
	*nasType.ExtendedRejectedNSSAI
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ExtendedCAGInformationList
	*nasType.LowerBoundTimerValue
	ForbiddenTAIRoaming           *nasType.TAIList
	ForbiddenTAIRegionalProvision *nasType.TAIList
	*nasType.N3IWFIdentifier
	*nasType.TNANInformation
}

func NewRegistrationReject(iei uint8) (registrationReject *RegistrationReject) {
	registrationReject = &RegistrationReject{}
	return registrationReject
}

const (
	RegistrationRejectT3346ValueType                    uint8 = 0x5F
	RegistrationRejectT3502ValueType                    uint8 = 0x16
	RegistrationRejectEAPMessageType                    uint8 = 0x78
	RegistrationRejectRejectedNSSAIType                 uint8 = 0x69
	RegistrationRejectCAGInformationListType            uint8 = 0x75
	RegistrationRejectExtendedRejectedNSSAIType         uint8 = 0x68
	RegistrationRejectDisasterReturnWaitRangeType       uint8 = 0x2C
	RegistrationRejectExtendedCAGInformationListType    uint8 = 0x71
	RegistrationRejectLowerBoundTimerValueType          uint8 = 0x3A
	RegistrationRejectForbiddenTAIRoamingType           uint8 = 0x1D
	RegistrationRejectForbiddenTAIRegionalProvisionType uint8 = 0x1E
	RegistrationRejectN3IWFIdentifierType               uint8 = 0x3E
	RegistrationRejectTNANInformationType               uint8 = 0x4D
)

func (a *RegistrationReject) EncodeRegistrationReject(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RegistrationRejectMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
	if a.T3346Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3346Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3346Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3346Value.Octet)
	}
	if a.T3502Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3502Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3502Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3502Value.Octet)
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
	if a.RejectedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RejectedNSSAI.Buffer)
	}
	if a.CAGInformationList != nil {
		binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.CAGInformationList.Buffer)
	}
	if a.ExtendedRejectedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.ExtendedRejectedNSSAI.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:uint8(a.ExtendedRejectedNSSAI.GetLen())])
	}
	if a.DisasterReturnWaitRange != nil {
		binary.Write(buffer, binary.BigEndian, a.DisasterReturnWaitRange.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.DisasterReturnWaitRange.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:uint8(a.DisasterReturnWaitRange.GetLen())])
	}
	if a.ExtendedCAGInformationList != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedCAGInformationList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedCAGInformationList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Buffer)
	}
	if a.LowerBoundTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.LowerBoundTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, a.LowerBoundTimerValue.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet)
	}
	if a.ForbiddenTAIRoaming != nil {
		binary.Write(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Buffer)
	}
	if a.ForbiddenTAIRegionalProvision != nil {
		binary.Write(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Buffer)
	}
	if a.N3IWFIdentifier != nil {
		binary.Write(buffer, binary.BigEndian, a.N3IWFIdentifier.GetIei())
		binary.Write(buffer, binary.BigEndian, a.N3IWFIdentifier.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.N3IWFIdentifier.Buffer)
	}
	if a.TNANInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.TNANInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.TNANInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.TNANInformation.Buffer)
	}
}

func (a *RegistrationReject) DecodeRegistrationReject(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RegistrationRejectMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
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
		case RegistrationRejectT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len)
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet)
		case RegistrationRejectT3502ValueType:
			a.T3502Value = nasType.NewT3502Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3502Value.Len)
			a.T3502Value.SetLen(a.T3502Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3502Value.Octet)
		case RegistrationRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case RegistrationRejectRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len)
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer[:a.RejectedNSSAI.GetLen()])
		case RegistrationRejectCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len)
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()])
		case RegistrationRejectExtendedRejectedNSSAIType:
			a.ExtendedRejectedNSSAI = nasType.NewExtendedRejectedNSSAI(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.ExtendedRejectedNSSAI.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:l])
		case RegistrationRejectDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.DisasterReturnWaitRange.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:l])
		case RegistrationRejectExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len)
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()])
		case RegistrationRejectLowerBoundTimerValueType:
			a.LowerBoundTimerValue = nasType.NewLowerBoundTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Len)
			a.LowerBoundTimerValue.SetLen(a.LowerBoundTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet)
		case RegistrationRejectForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len)
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()])
		case RegistrationRejectForbiddenTAIRegionalProvisionType:
			a.ForbiddenTAIRegionalProvision = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Len)
			a.ForbiddenTAIRegionalProvision.SetLen(a.ForbiddenTAIRegionalProvision.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.Buffer[:a.ForbiddenTAIRegionalProvision.GetLen()])
		case RegistrationRejectN3IWFIdentifierType:
			a.N3IWFIdentifier = nasType.NewN3IWFIdentifier(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.N3IWFIdentifier.Len)
			a.N3IWFIdentifier.SetLen(a.N3IWFIdentifier.GetLen())
			binary.Read(buffer, binary.BigEndian, a.N3IWFIdentifier.Buffer[:a.N3IWFIdentifier.GetLen()])
		case RegistrationRejectTNANInformationType:
			a.TNANInformation = nasType.NewTNANInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.TNANInformation.Len)
			a.TNANInformation.SetLen(a.TNANInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.TNANInformation.Buffer[:a.TNANInformation.GetLen()])
		default:
		}
	}
}
