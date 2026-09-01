// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type ServiceReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ServiceRejectMessageIdentity
	nasType.Cause5GMM
	*nasType.PDUSessionStatus
	*nasType.T3346Value
	*nasType.EAPMessage
	*nasType.T3448Value
	*nasType.CAGInformationList
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ExtendedCAGInformationList
	*nasType.LowerBoundTimerValue
	ForbiddenTAIRoaming           *nasType.TAIList
	ForbiddenTAIRegionalProvision *nasType.TAIList
}

func NewServiceReject(iei uint8) (serviceReject *ServiceReject) {
	serviceReject = &ServiceReject{}
	return serviceReject
}

const (
	ServiceRejectPDUSessionStatusType              uint8 = 0x50
	ServiceRejectT3346ValueType                    uint8 = 0x5F
	ServiceRejectEAPMessageType                    uint8 = 0x78
	ServiceRejectT3448ValueType                    uint8 = 0x6B
	ServiceRejectCAGInformationListType            uint8 = 0x75
	ServiceRejectDisasterReturnWaitRangeType       uint8 = 0x2C
	ServiceRejectExtendedCAGInformationListType    uint8 = 0x71
	ServiceRejectLowerBoundTimerValueType          uint8 = 0x3A
	ServiceRejectForbiddenTAIRoamingType           uint8 = 0x1D
	ServiceRejectForbiddenTAIRegionalProvisionType uint8 = 0x1E
)

func (a *ServiceReject) EncodeServiceReject(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ServiceRejectMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
	if a.PDUSessionStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PDUSessionStatus.Buffer)
	}
	if a.T3346Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3346Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3346Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3346Value.Octet)
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
	if a.T3448Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3448Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3448Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3448Value.Octet)
	}
	if a.CAGInformationList != nil {
		binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.CAGInformationList.Buffer)
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
}

func (a *ServiceReject) DecodeServiceReject(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ServiceRejectMessageIdentity.Octet)
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
		case ServiceRejectPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len)
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()])
		case ServiceRejectT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len)
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet)
		case ServiceRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case ServiceRejectT3448ValueType:
			a.T3448Value = nasType.NewT3448Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3448Value.Len)
			a.T3448Value.SetLen(a.T3448Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3448Value.Octet)
		case ServiceRejectCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len)
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()])
		case ServiceRejectDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.DisasterReturnWaitRange.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:l])
		case ServiceRejectExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len)
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()])
		case ServiceRejectLowerBoundTimerValueType:
			a.LowerBoundTimerValue = nasType.NewLowerBoundTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Len)
			a.LowerBoundTimerValue.SetLen(a.LowerBoundTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet)
		case ServiceRejectForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len)
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()])
		case ServiceRejectForbiddenTAIRegionalProvisionType:
			a.ForbiddenTAIRegionalProvision = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Len)
			a.ForbiddenTAIRegionalProvision.SetLen(a.ForbiddenTAIRegionalProvision.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.Buffer[:a.ForbiddenTAIRegionalProvision.GetLen()])
		default:
		}
	}
}
