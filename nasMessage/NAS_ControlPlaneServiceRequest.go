// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// ControlPlaneServiceRequest 8.2.30
type ControlPlaneServiceRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.CONTROLPLANESERVICEREQUESTMessageIdentity
	nasType.ControlPlaneServiceTypeAndNgksi
	*nasType.PDUSessionStatus
	*nasType.UplinkDataStatus
	*nasType.NASMessageContainer
	*nasType.AllowedPDUSessionStatus
	*nasType.UERequestType
	*nasType.PagingRestriction
	*nasType.CIoTSmallDataContainer
	*nasType.PduSessionID2Value
	*nasType.AdditionalInformation
	*nasType.ReleaseAssistanceIndication
}

func NewControlPlaneServiceRequest(iei uint8) (controlPlaneServiceRequest *ControlPlaneServiceRequest) {
	controlPlaneServiceRequest = &ControlPlaneServiceRequest{}
	return controlPlaneServiceRequest
}

const (
	ControlPlaneServiceRequestPDUSessionStatusType            uint8 = 0x50
	ControlPlaneServiceRequestUplinkDataStatusType            uint8 = 0x40
	ControlPlaneServiceRequestNASMessageContainerType         uint8 = 0x71
	ControlPlaneServiceRequestAllowedPDUSessionStatusType     uint8 = 0x25
	ControlPlaneServiceRequestUERequestTypeType               uint8 = 0x29
	ControlPlaneServiceRequestPagingRestrictionType           uint8 = 0x28
	ControlPlaneServiceRequestCIoTSmallDataContainerType      uint8 = 0x6F
	ControlPlaneServiceRequestPduSessionIDType                uint8 = 0x12
	ControlPlaneServiceRequestAdditionalInformationType       uint8 = 0x24
	ControlPlaneServiceRequestReleaseAssistanceIndicationType uint8 = 0x0F
)

func (a *ControlPlaneServiceRequest) EncodeControlPlaneServiceRequest(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.CONTROLPLANESERVICEREQUESTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ControlPlaneServiceTypeAndNgksi.Octet)
	if a.PDUSessionStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PDUSessionStatus.Buffer)
	}
	if a.UplinkDataStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UplinkDataStatus.Buffer)
	}
	if a.NASMessageContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer)
	}
	if a.AllowedPDUSessionStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Buffer)
	}
	if a.UERequestType != nil {
		binary.Write(buffer, binary.BigEndian, a.UERequestType.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.UERequestType.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.UERequestType.Buffer[:uint8(a.UERequestType.GetLen())])
	}
	if a.PagingRestriction != nil {
		binary.Write(buffer, binary.BigEndian, a.PagingRestriction.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.PagingRestriction.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.PagingRestriction.Buffer[:uint8(a.PagingRestriction.GetLen())])
	}
	if a.CIoTSmallDataContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.CIoTSmallDataContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.CIoTSmallDataContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.CIoTSmallDataContainer.Buffer)
	}
	if a.PduSessionID2Value != nil {
		binary.Write(buffer, binary.BigEndian, a.PduSessionID2Value.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet)
	}
	if a.AdditionalInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AdditionalInformation.Buffer)
	}
	if a.ReleaseAssistanceIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.ReleaseAssistanceIndication.Octet)
	}
}

func (a *ControlPlaneServiceRequest) DecodeControlPlaneServiceRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.CONTROLPLANESERVICEREQUESTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ControlPlaneServiceTypeAndNgksi.Octet)
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
		case ControlPlaneServiceRequestPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len)
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer)
		case ControlPlaneServiceRequestUplinkDataStatusType:
			a.UplinkDataStatus = nasType.NewUplinkDataStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UplinkDataStatus.Len)
			a.UplinkDataStatus.SetLen(a.UplinkDataStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.UplinkDataStatus.Buffer)
		case ControlPlaneServiceRequestNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len)
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer)
		case ControlPlaneServiceRequestAllowedPDUSessionStatusType:
			a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Len)
			a.AllowedPDUSessionStatus.SetLen(a.AllowedPDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.Buffer)
		case ControlPlaneServiceRequestUERequestTypeType:
			a.UERequestType = nasType.NewUERequestType(ieiN)
			var lenN0 uint8
			binary.Read(buffer, binary.BigEndian, &lenN0)
			a.UERequestType.SetLen(uint16(lenN0))
			binary.Read(buffer, binary.BigEndian, a.UERequestType.Buffer[:lenN0])
		case ControlPlaneServiceRequestPagingRestrictionType:
			a.PagingRestriction = nasType.NewPagingRestriction(ieiN)
			var lenN1 uint8
			binary.Read(buffer, binary.BigEndian, &lenN1)
			a.PagingRestriction.SetLen(uint16(lenN1))
			binary.Read(buffer, binary.BigEndian, a.PagingRestriction.Buffer[:lenN1])
		case ControlPlaneServiceRequestCIoTSmallDataContainerType:
			a.CIoTSmallDataContainer = nasType.NewCIoTSmallDataContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.CIoTSmallDataContainer.Len)
			a.CIoTSmallDataContainer.SetLen(a.CIoTSmallDataContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.CIoTSmallDataContainer.Buffer[:a.CIoTSmallDataContainer.GetLen()])
		case ControlPlaneServiceRequestPduSessionIDType:
			a.PduSessionID2Value = nasType.NewPduSessionID2Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet)
		case ControlPlaneServiceRequestAdditionalInformationType:
			a.AdditionalInformation = nasType.NewAdditionalInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AdditionalInformation.Len)
			a.AdditionalInformation.SetLen(a.AdditionalInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AdditionalInformation.Buffer[:a.AdditionalInformation.GetLen()])
		case ControlPlaneServiceRequestReleaseAssistanceIndicationType:
			a.ReleaseAssistanceIndication = nasType.NewReleaseAssistanceIndication(ieiN)
			a.ReleaseAssistanceIndication.Octet = ieiN
		default:
		}
	}
}
