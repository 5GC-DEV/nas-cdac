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

type ServiceAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ServiceAcceptMessageIdentity
	*nasType.PDUSessionStatus
	*nasType.PDUSessionReactivationResult
	*nasType.PDUSessionReactivationResultErrorCause
	*nasType.EAPMessage
	*nasType.T3448Value
	*nasType.FiveGSAdditionalRequestResult
	ForbiddenTAIRoaming           *nasType.TAIList
	ForbiddenTAIRegionalProvision *nasType.TAIList
}

func NewServiceAccept(iei uint8) (serviceAccept *ServiceAccept) {
	serviceAccept = &ServiceAccept{}
	return serviceAccept
}

const (
	ServiceAcceptPDUSessionStatusType                       uint8 = 0x50
	ServiceAcceptPDUSessionReactivationResultType           uint8 = 0x26
	ServiceAcceptPDUSessionReactivationResultErrorCauseType uint8 = 0x72
	ServiceAcceptEAPMessageType                             uint8 = 0x78
	ServiceAcceptT3448ValueType                             uint8 = 0x6B
	ServiceAcceptFiveGSAdditionalRequestResultType          uint8 = 0x34
	ServiceAcceptForbiddenTAIRoamingType                    uint8 = 0x1D
	ServiceAcceptForbiddenTAIRegionalProvisionType          uint8 = 0x1E
)

func (a *ServiceAccept) EncodeServiceAccept(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ServiceAcceptMessageIdentity.Octet)
	if a.PDUSessionStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PDUSessionStatus.Buffer)
	}
	if a.PDUSessionReactivationResult != nil {
		binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PDUSessionReactivationResult.Buffer)
	}
	if a.PDUSessionReactivationResultErrorCause != nil {
		binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PDUSessionReactivationResultErrorCause.Buffer)
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
	if a.FiveGSAdditionalRequestResult != nil {
		binary.Write(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.FiveGSAdditionalRequestResult.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.Buffer[:uint8(a.FiveGSAdditionalRequestResult.GetLen())])
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

func (a *ServiceAccept) DecodeServiceAccept(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ServiceAcceptMessageIdentity.Octet)
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
		case ServiceAcceptPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len)
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()])
		case ServiceAcceptPDUSessionReactivationResultType:
			a.PDUSessionReactivationResult = nasType.NewPDUSessionReactivationResult(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResult.Len)
			a.PDUSessionReactivationResult.SetLen(a.PDUSessionReactivationResult.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResult.Buffer[:a.PDUSessionReactivationResult.GetLen()])
		case ServiceAcceptPDUSessionReactivationResultErrorCauseType:
			a.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResultErrorCause.Len)
			a.PDUSessionReactivationResultErrorCause.SetLen(a.PDUSessionReactivationResultErrorCause.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.Buffer[:a.PDUSessionReactivationResultErrorCause.GetLen()])
		case ServiceAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case ServiceAcceptT3448ValueType:
			a.T3448Value = nasType.NewT3448Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3448Value.Len)
			a.T3448Value.SetLen(a.T3448Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3448Value.Octet)
		case ServiceAcceptFiveGSAdditionalRequestResultType:
			a.FiveGSAdditionalRequestResult = nasType.NewFiveGSAdditionalRequestResult(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.FiveGSAdditionalRequestResult.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.Buffer[:l])
		case ServiceAcceptForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len)
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()])
		case ServiceAcceptForbiddenTAIRegionalProvisionType:
			a.ForbiddenTAIRegionalProvision = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Len)
			a.ForbiddenTAIRegionalProvision.SetLen(a.ForbiddenTAIRegionalProvision.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.Buffer[:a.ForbiddenTAIRegionalProvision.GetLen()])
		default:
		}
	}
}
