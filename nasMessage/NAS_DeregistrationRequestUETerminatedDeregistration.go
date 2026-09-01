// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type DeregistrationRequestUETerminatedDeregistration struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DeregistrationRequestMessageIdentity
	nasType.SpareHalfOctetAndDeregistrationType
	*nasType.Cause5GMM
	*nasType.T3346Value
	*nasType.RejectedNSSAI
	*nasType.CAGInformationList
	*nasType.ExtendedRejectedNSSAI
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ExtendedCAGInformationList
	*nasType.LowerBoundTimerValue
	ForbiddenTAIRoaming           *nasType.TAIList
	ForbiddenTAIRegionalProvision *nasType.TAIList
}

func NewDeregistrationRequestUETerminatedDeregistration(iei uint8) (deregistrationRequestUETerminatedDeregistration *DeregistrationRequestUETerminatedDeregistration) {
	deregistrationRequestUETerminatedDeregistration = &DeregistrationRequestUETerminatedDeregistration{}
	return deregistrationRequestUETerminatedDeregistration
}

const (
	DeregistrationRequestUETerminatedDeregistrationCause5GMMType                     uint8 = 0x58
	DeregistrationRequestUETerminatedDeregistrationT3346ValueType                    uint8 = 0x5F
	DeregistrationRequestUETerminatedDeregistrationRejectedNSSAIType                 uint8 = 0x6D
	DeregistrationRequestUETerminatedDeregistrationCAGInformationListType            uint8 = 0x75
	DeregistrationRequestUETerminatedDeregistrationExtendedRejectedNSSAIType         uint8 = 0x68
	DeregistrationRequestUETerminatedDeregistrationDisasterReturnWaitRangeType       uint8 = 0x2C
	DeregistrationRequestUETerminatedDeregistrationExtendedCAGInformationListType    uint8 = 0x71
	DeregistrationRequestUETerminatedDeregistrationLowerBoundTimerValueType          uint8 = 0x3A
	DeregistrationRequestUETerminatedDeregistrationForbiddenTAIRoamingType           uint8 = 0x1D
	DeregistrationRequestUETerminatedDeregistrationForbiddenTAIRegionalProvisionType uint8 = 0x1E
)

func (a *DeregistrationRequestUETerminatedDeregistration) EncodeDeregistrationRequestUETerminatedDeregistration(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.DeregistrationRequestMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndDeregistrationType.Octet)
	if a.Cause5GMM != nil {
		binary.Write(buffer, binary.BigEndian, a.Cause5GMM.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
	}
	if a.T3346Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3346Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3346Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3346Value.Octet)
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
}

func (a *DeregistrationRequestUETerminatedDeregistration) DecodeDeregistrationRequestUETerminatedDeregistration(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.DeregistrationRequestMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndDeregistrationType.Octet)
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
		case DeregistrationRequestUETerminatedDeregistrationCause5GMMType:
			a.Cause5GMM = nasType.NewCause5GMM(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
		case DeregistrationRequestUETerminatedDeregistrationT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len)
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet)
		case DeregistrationRequestUETerminatedDeregistrationRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len)
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer[:a.RejectedNSSAI.GetLen()])
		case DeregistrationRequestUETerminatedDeregistrationCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len)
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()])
		case DeregistrationRequestUETerminatedDeregistrationExtendedRejectedNSSAIType:
			a.ExtendedRejectedNSSAI = nasType.NewExtendedRejectedNSSAI(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.ExtendedRejectedNSSAI.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:l])
		case DeregistrationRequestUETerminatedDeregistrationDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.DisasterReturnWaitRange.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:l])
		case DeregistrationRequestUETerminatedDeregistrationExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len)
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()])
		case DeregistrationRequestUETerminatedDeregistrationLowerBoundTimerValueType:
			a.LowerBoundTimerValue = nasType.NewLowerBoundTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Len)
			a.LowerBoundTimerValue.SetLen(a.LowerBoundTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet)
		case DeregistrationRequestUETerminatedDeregistrationForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len)
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()])
		case DeregistrationRequestUETerminatedDeregistrationForbiddenTAIRegionalProvisionType:
			a.ForbiddenTAIRegionalProvision = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Len)
			a.ForbiddenTAIRegionalProvision.SetLen(a.ForbiddenTAIRegionalProvision.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.Buffer[:a.ForbiddenTAIRegionalProvision.GetLen()])
		default:
		}
	}
}
