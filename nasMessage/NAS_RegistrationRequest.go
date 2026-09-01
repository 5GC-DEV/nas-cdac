// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type RegistrationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationRequestMessageIdentity
	nasType.NgksiAndRegistrationType5GS
	nasType.MobileIdentity5GS
	*nasType.NoncurrentNativeNASKeySetIdentifier
	*nasType.Capability5GMM
	*nasType.UESecurityCapability
	*nasType.RequestedNSSAI
	*nasType.LastVisitedRegisteredTAI
	*nasType.S1UENetworkCapability
	*nasType.UplinkDataStatus
	*nasType.PDUSessionStatus
	*nasType.MICOIndication
	*nasType.UEStatus
	*nasType.AdditionalGUTI
	*nasType.AllowedPDUSessionStatus
	*nasType.UesUsageSetting
	*nasType.RequestedDRXParameters
	*nasType.EPSNASMessageContainer
	*nasType.LADNIndication
	*nasType.PayloadContainer
	*nasType.NetworkSlicingIndication
	*nasType.UpdateType5GS
	*nasType.MobileStationClassmark2
	*nasType.SupportedCodecs
	*nasType.NASMessageContainer
	*nasType.EPSBearerContextStatus
	*nasType.ExtendedDRXParameters
	*nasType.UERadioCapabilityID
	*nasType.RequestedMappedNSSAI
	*nasType.AdditionalInformationRequested
	*nasType.WUSAssistanceInformation
	*nasType.N5GCIndication
	*nasType.NBN1ModeDRXParameters
	*nasType.UERequestType
	*nasType.PagingRestriction
	*nasType.ServiceLevelAAContainer
	*nasType.NID
	*nasType.PLMNIdentityWithDisasterCondition
	*nasType.PEIPSAssistanceInformation
	*nasType.TimeDuration
	*nasType.Non3GPPPathSwitchingInformation
	*nasType.AUN3Indication
	*nasType.T3512Value
}

func NewRegistrationRequest(iei uint8) (registrationRequest *RegistrationRequest) {
	registrationRequest = &RegistrationRequest{}
	return registrationRequest
}

const (
	RegistrationRequestNoncurrentNativeNASKeySetIdentifierType uint8 = 0x0C
	RegistrationRequestCapability5GMMType                      uint8 = 0x10
	RegistrationRequestUESecurityCapabilityType                uint8 = 0x2E
	RegistrationRequestRequestedNSSAIType                      uint8 = 0x2F
	RegistrationRequestLastVisitedRegisteredTAIType            uint8 = 0x52
	RegistrationRequestS1UENetworkCapabilityType               uint8 = 0x17
	RegistrationRequestUplinkDataStatusType                    uint8 = 0x40
	RegistrationRequestPDUSessionStatusType                    uint8 = 0x50
	RegistrationRequestMICOIndicationType                      uint8 = 0x0B
	RegistrationRequestUEStatusType                            uint8 = 0x2B
	RegistrationRequestAdditionalGUTIType                      uint8 = 0x77
	RegistrationRequestAllowedPDUSessionStatusType             uint8 = 0x25
	RegistrationRequestUesUsageSettingType                     uint8 = 0x18
	RegistrationRequestRequestedDRXParametersType              uint8 = 0x51
	RegistrationRequestEPSNASMessageContainerType              uint8 = 0x70
	RegistrationRequestLADNIndicationType                      uint8 = 0x74
	RegistrationRequestPayloadContainerType                    uint8 = 0x7B
	RegistrationRequestNetworkSlicingIndicationType            uint8 = 0x09
	RegistrationRequestUpdateType5GSType                       uint8 = 0x53
	RegistrationRequestNASMessageContainerType                 uint8 = 0x71
	RegistrationRequestEPSBearerContextStatusType              uint8 = 0x60
	RegistrationRequestExtendedDRXParametersType               uint8 = 0x6E
	RegistrationRequestUERadioCapabilityIDType                 uint8 = 0x67
	RegistrationRequestRequestedMappedNSSAIType                uint8 = 0x35
	RegistrationRequestAdditionalInformationRequestedType      uint8 = 0x48
	RegistrationRequestWUSAssistanceInformationType            uint8 = 0x1A
	RegistrationRequestN5GCIndicationType                      uint8 = 0x0A
	RegistrationRequestNBN1ModeDRXParametersType               uint8 = 0x30
	RegistrationRequestUERequestTypeType                       uint8 = 0x29
	RegistrationRequestPagingRestrictionType                   uint8 = 0x28
	RegistrationRequestServiceLevelAAContainerType             uint8 = 0x72
	RegistrationRequestNIDType                                 uint8 = 0x32
	RegistrationRequestPLMNIdentityWithDisasterConditionType   uint8 = 0x16
	RegistrationRequestPEIPSAssistanceInformationType          uint8 = 0x2A
	RegistrationRequestTimeDurationType                        uint8 = 0x3C
	RegistrationRequestNon3GPPPathSwitchingInformationType     uint8 = 0x3F
	RegistrationRequestAUN3IndicationType                      uint8 = 0x56
	RegistrationRequestMobileStationClassmark2Type             uint8 = 0x41
	RegistrationRequestSupportedCodecsType                     uint8 = 0x42
	RegistrationRequestRequestedT3512ValueType                 uint8 = 0x3B
)

func (a *RegistrationRequest) EncodeRegistrationRequest(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RegistrationRequestMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.NgksiAndRegistrationType5GS.Octet)
	binary.Write(buffer, binary.BigEndian, a.MobileIdentity5GS.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.MobileIdentity5GS.Buffer)
	if a.NoncurrentNativeNASKeySetIdentifier != nil {
		binary.Write(buffer, binary.BigEndian, &a.NoncurrentNativeNASKeySetIdentifier.Octet)
	}
	if a.Capability5GMM != nil {
		binary.Write(buffer, binary.BigEndian, a.Capability5GMM.GetIei())
		binary.Write(buffer, binary.BigEndian, a.Capability5GMM.GetLen())
		binary.Write(buffer, binary.BigEndian, a.Capability5GMM.Octet[:a.Capability5GMM.GetLen()])
	}
	if a.UESecurityCapability != nil {
		binary.Write(buffer, binary.BigEndian, a.UESecurityCapability.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UESecurityCapability.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UESecurityCapability.Buffer)
	}
	if a.RequestedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.RequestedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RequestedNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RequestedNSSAI.Buffer)
	}
	if a.LastVisitedRegisteredTAI != nil {
		binary.Write(buffer, binary.BigEndian, a.LastVisitedRegisteredTAI.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.LastVisitedRegisteredTAI.Octet)
	}
	if a.S1UENetworkCapability != nil {
		binary.Write(buffer, binary.BigEndian, a.S1UENetworkCapability.GetIei())
		binary.Write(buffer, binary.BigEndian, a.S1UENetworkCapability.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.S1UENetworkCapability.Buffer)
	}
	if a.UplinkDataStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UplinkDataStatus.Buffer)
	}
	if a.PDUSessionStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PDUSessionStatus.Buffer)
	}
	if a.MICOIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.MICOIndication.Octet)
	}
	if a.UEStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.UEStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UEStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UEStatus.Octet)
	}
	if a.AdditionalGUTI != nil {
		binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.GetLen())
		binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.Octet[:a.AdditionalGUTI.GetLen()])
	}
	if a.AllowedPDUSessionStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Buffer)
	}
	if a.UesUsageSetting != nil {
		binary.Write(buffer, binary.BigEndian, a.UesUsageSetting.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UesUsageSetting.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UesUsageSetting.Octet)
	}
	if a.RequestedDRXParameters != nil {
		binary.Write(buffer, binary.BigEndian, a.RequestedDRXParameters.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RequestedDRXParameters.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RequestedDRXParameters.Octet)
	}
	if a.EPSNASMessageContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.EPSNASMessageContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EPSNASMessageContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EPSNASMessageContainer.Buffer)
	}
	if a.LADNIndication != nil {
		binary.Write(buffer, binary.BigEndian, a.LADNIndication.GetIei())
		binary.Write(buffer, binary.BigEndian, a.LADNIndication.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.LADNIndication.Buffer)
	}
	if a.PayloadContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PayloadContainer.Buffer)
	}
	if a.NetworkSlicingIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.NetworkSlicingIndication.Octet)
	}
	if a.UpdateType5GS != nil {
		binary.Write(buffer, binary.BigEndian, a.UpdateType5GS.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UpdateType5GS.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UpdateType5GS.Octet)
	}
	if a.MobileStationClassmark2 != nil {
		binary.Write(buffer, binary.BigEndian, a.MobileStationClassmark2.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.MobileStationClassmark2.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.MobileStationClassmark2.Buffer[:uint8(a.MobileStationClassmark2.GetLen())])
	}
	if a.SupportedCodecs != nil {
		binary.Write(buffer, binary.BigEndian, a.SupportedCodecs.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.SupportedCodecs.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.SupportedCodecs.Buffer[:uint8(a.SupportedCodecs.GetLen())])
	}
	if a.NASMessageContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer)
	}
	if a.EPSBearerContextStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.EPSBearerContextStatus.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.Buffer[:uint8(a.EPSBearerContextStatus.GetLen())])
	}
	if a.ExtendedDRXParameters != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedDRXParameters.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.ExtendedDRXParameters.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.ExtendedDRXParameters.Buffer[:uint8(a.ExtendedDRXParameters.GetLen())])
	}
	if a.UERadioCapabilityID != nil {
		binary.Write(buffer, binary.BigEndian, a.UERadioCapabilityID.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.UERadioCapabilityID.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:uint8(a.UERadioCapabilityID.GetLen())])
	}
	if a.RequestedMappedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.RequestedMappedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.RequestedMappedNSSAI.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.RequestedMappedNSSAI.Buffer[:uint8(a.RequestedMappedNSSAI.GetLen())])
	}
	if a.AdditionalInformationRequested != nil {
		binary.Write(buffer, binary.BigEndian, a.AdditionalInformationRequested.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.AdditionalInformationRequested.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.AdditionalInformationRequested.Buffer[:uint8(a.AdditionalInformationRequested.GetLen())])
	}
	if a.WUSAssistanceInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.WUSAssistanceInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.WUSAssistanceInformation.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.WUSAssistanceInformation.Buffer[:uint8(a.WUSAssistanceInformation.GetLen())])
	}
	if a.N5GCIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.N5GCIndication.Octet)
	}
	if a.NBN1ModeDRXParameters != nil {
		binary.Write(buffer, binary.BigEndian, a.NBN1ModeDRXParameters.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.NBN1ModeDRXParameters.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.NBN1ModeDRXParameters.Buffer[:uint8(a.NBN1ModeDRXParameters.GetLen())])
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
	if a.ServiceLevelAAContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer)
	}
	if a.NID != nil {
		binary.Write(buffer, binary.BigEndian, a.NID.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.NID.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.NID.Buffer[:uint8(a.NID.GetLen())])
	}
	if a.PLMNIdentityWithDisasterCondition != nil {
		binary.Write(buffer, binary.BigEndian, a.PLMNIdentityWithDisasterCondition.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.PLMNIdentityWithDisasterCondition.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.PLMNIdentityWithDisasterCondition.Buffer[:uint8(a.PLMNIdentityWithDisasterCondition.GetLen())])
	}
	if a.PEIPSAssistanceInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.PEIPSAssistanceInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.PEIPSAssistanceInformation.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.PEIPSAssistanceInformation.Buffer[:uint8(a.PEIPSAssistanceInformation.GetLen())])
	}
	if a.TimeDuration != nil {
		binary.Write(buffer, binary.BigEndian, a.TimeDuration.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.TimeDuration.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.TimeDuration.Buffer[:uint8(a.TimeDuration.GetLen())])
	}
	if a.Non3GPPPathSwitchingInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.Non3GPPPathSwitchingInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.Non3GPPPathSwitchingInformation.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.Non3GPPPathSwitchingInformation.Buffer[:uint8(a.Non3GPPPathSwitchingInformation.GetLen())])
	}
	if a.AUN3Indication != nil {
		binary.Write(buffer, binary.BigEndian, a.AUN3Indication.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AUN3Indication.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AUN3Indication.Octet)
	}
	if a.T3512Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3512Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3512Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3512Value.Octet)
	}
}

func (a *RegistrationRequest) DecodeRegistrationRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RegistrationRequestMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.NgksiAndRegistrationType5GS.Octet)
	binary.Read(buffer, binary.BigEndian, &a.MobileIdentity5GS.Len)
	a.MobileIdentity5GS.SetLen(a.MobileIdentity5GS.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.MobileIdentity5GS.Buffer)
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
		case RegistrationRequestNoncurrentNativeNASKeySetIdentifierType:
			a.NoncurrentNativeNASKeySetIdentifier = nasType.NewNoncurrentNativeNASKeySetIdentifier(ieiN)
			a.NoncurrentNativeNASKeySetIdentifier.Octet = ieiN
		case RegistrationRequestCapability5GMMType:
			a.Capability5GMM = nasType.NewCapability5GMM(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Capability5GMM.Len)
			a.Capability5GMM.SetLen(a.Capability5GMM.GetLen())
			binary.Read(buffer, binary.BigEndian, a.Capability5GMM.Octet[:a.Capability5GMM.GetLen()])
		case RegistrationRequestUESecurityCapabilityType:
			a.UESecurityCapability = nasType.NewUESecurityCapability(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UESecurityCapability.Len)
			a.UESecurityCapability.SetLen(a.UESecurityCapability.GetLen())
			binary.Read(buffer, binary.BigEndian, a.UESecurityCapability.Buffer[:a.UESecurityCapability.GetLen()])
		case RegistrationRequestRequestedNSSAIType:
			a.RequestedNSSAI = nasType.NewRequestedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RequestedNSSAI.Len)
			a.RequestedNSSAI.SetLen(a.RequestedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RequestedNSSAI.Buffer[:a.RequestedNSSAI.GetLen()])
		case RegistrationRequestLastVisitedRegisteredTAIType:
			a.LastVisitedRegisteredTAI = nasType.NewLastVisitedRegisteredTAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.LastVisitedRegisteredTAI.Octet)
		case RegistrationRequestS1UENetworkCapabilityType:
			a.S1UENetworkCapability = nasType.NewS1UENetworkCapability(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.S1UENetworkCapability.Len)
			a.S1UENetworkCapability.SetLen(a.S1UENetworkCapability.GetLen())
			binary.Read(buffer, binary.BigEndian, a.S1UENetworkCapability.Buffer[:a.S1UENetworkCapability.GetLen()])
		case RegistrationRequestUplinkDataStatusType:
			a.UplinkDataStatus = nasType.NewUplinkDataStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UplinkDataStatus.Len)
			a.UplinkDataStatus.SetLen(a.UplinkDataStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.UplinkDataStatus.Buffer[:a.UplinkDataStatus.GetLen()])
		case RegistrationRequestPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len)
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()])
		case RegistrationRequestMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case RegistrationRequestUEStatusType:
			a.UEStatus = nasType.NewUEStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UEStatus.Len)
			a.UEStatus.SetLen(a.UEStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.UEStatus.Octet)
		case RegistrationRequestAdditionalGUTIType:
			a.AdditionalGUTI = nasType.NewAdditionalGUTI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AdditionalGUTI.Len)
			a.AdditionalGUTI.SetLen(a.AdditionalGUTI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AdditionalGUTI.Octet[:a.AdditionalGUTI.GetLen()])
		case RegistrationRequestAllowedPDUSessionStatusType:
			a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Len)
			a.AllowedPDUSessionStatus.SetLen(a.AllowedPDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.Buffer[:a.AllowedPDUSessionStatus.GetLen()])
		case RegistrationRequestUesUsageSettingType:
			a.UesUsageSetting = nasType.NewUesUsageSetting(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UesUsageSetting.Len)
			a.UesUsageSetting.SetLen(a.UesUsageSetting.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.UesUsageSetting.Octet)
		case RegistrationRequestRequestedDRXParametersType:
			a.RequestedDRXParameters = nasType.NewRequestedDRXParameters(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RequestedDRXParameters.Len)
			a.RequestedDRXParameters.SetLen(a.RequestedDRXParameters.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.RequestedDRXParameters.Octet)
		case RegistrationRequestEPSNASMessageContainerType:
			a.EPSNASMessageContainer = nasType.NewEPSNASMessageContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EPSNASMessageContainer.Len)
			a.EPSNASMessageContainer.SetLen(a.EPSNASMessageContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EPSNASMessageContainer.Buffer[:a.EPSNASMessageContainer.GetLen()])
		case RegistrationRequestLADNIndicationType:
			a.LADNIndication = nasType.NewLADNIndication(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.LADNIndication.Len)
			a.LADNIndication.SetLen(a.LADNIndication.GetLen())
			binary.Read(buffer, binary.BigEndian, a.LADNIndication.Buffer[:a.LADNIndication.GetLen()])
		case RegistrationRequestPayloadContainerType:
			a.PayloadContainer = nasType.NewPayloadContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Len)
			a.PayloadContainer.SetLen(a.PayloadContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PayloadContainer.Buffer[:a.PayloadContainer.GetLen()])
		case RegistrationRequestNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case RegistrationRequestUpdateType5GSType:
			a.UpdateType5GS = nasType.NewUpdateType5GS(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UpdateType5GS.Len)
			a.UpdateType5GS.SetLen(a.UpdateType5GS.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.UpdateType5GS.Octet)
		case RegistrationRequestNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len)
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer[:a.NASMessageContainer.GetLen()])
		case RegistrationRequestEPSBearerContextStatusType:
			a.EPSBearerContextStatus = nasType.NewEPSBearerContextStatus(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.EPSBearerContextStatus.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.EPSBearerContextStatus.Buffer[:lenN])
		case RegistrationRequestExtendedDRXParametersType:
			a.ExtendedDRXParameters = nasType.NewExtendedDRXParameters(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.ExtendedDRXParameters.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.ExtendedDRXParameters.Buffer[:lenN])
		case RegistrationRequestUERadioCapabilityIDType:
			a.UERadioCapabilityID = nasType.NewUERadioCapabilityID(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.UERadioCapabilityID.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:lenN])
		case RegistrationRequestRequestedMappedNSSAIType:
			a.RequestedMappedNSSAI = nasType.NewRequestedMappedNSSAI(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.RequestedMappedNSSAI.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.RequestedMappedNSSAI.Buffer[:lenN])
		case RegistrationRequestAdditionalInformationRequestedType:
			a.AdditionalInformationRequested = nasType.NewAdditionalInformationRequested(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.AdditionalInformationRequested.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.AdditionalInformationRequested.Buffer[:lenN])
		case RegistrationRequestWUSAssistanceInformationType:
			a.WUSAssistanceInformation = nasType.NewWUSAssistanceInformation(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.WUSAssistanceInformation.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.WUSAssistanceInformation.Buffer[:lenN])
		case RegistrationRequestN5GCIndicationType:
			a.N5GCIndication = nasType.NewN5GCIndication(ieiN)
			a.N5GCIndication.Octet = ieiN
		case RegistrationRequestNBN1ModeDRXParametersType:
			a.NBN1ModeDRXParameters = nasType.NewNBN1ModeDRXParameters(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.NBN1ModeDRXParameters.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.NBN1ModeDRXParameters.Buffer[:lenN])
		case RegistrationRequestUERequestTypeType:
			a.UERequestType = nasType.NewUERequestType(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.UERequestType.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.UERequestType.Buffer[:lenN])
		case RegistrationRequestPagingRestrictionType:
			a.PagingRestriction = nasType.NewPagingRestriction(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.PagingRestriction.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.PagingRestriction.Buffer[:lenN])
		case RegistrationRequestServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len)
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()])
		case RegistrationRequestNIDType:
			a.NID = nasType.NewNID(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.NID.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.NID.Buffer[:lenN])
		case RegistrationRequestPLMNIdentityWithDisasterConditionType:
			a.PLMNIdentityWithDisasterCondition = nasType.NewPLMNIdentityWithDisasterCondition(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.PLMNIdentityWithDisasterCondition.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.PLMNIdentityWithDisasterCondition.Buffer[:lenN])
		case RegistrationRequestPEIPSAssistanceInformationType:
			a.PEIPSAssistanceInformation = nasType.NewPEIPSAssistanceInformation(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.PEIPSAssistanceInformation.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.PEIPSAssistanceInformation.Buffer[:lenN])
		case RegistrationRequestTimeDurationType:
			a.TimeDuration = nasType.NewTimeDuration(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.TimeDuration.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.TimeDuration.Buffer[:lenN])
		case RegistrationRequestNon3GPPPathSwitchingInformationType:
			a.Non3GPPPathSwitchingInformation = nasType.NewNon3GPPPathSwitchingInformation(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.Non3GPPPathSwitchingInformation.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.Non3GPPPathSwitchingInformation.Buffer[:lenN])
		case RegistrationRequestAUN3IndicationType:
			a.AUN3Indication = nasType.NewAUN3Indication(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AUN3Indication.Len)
			a.AUN3Indication.SetLen(a.AUN3Indication.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.AUN3Indication.Octet)
		case RegistrationRequestMobileStationClassmark2Type:
			a.MobileStationClassmark2 = nasType.NewMobileStationClassmark2(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.MobileStationClassmark2.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.MobileStationClassmark2.Buffer[:lenN])
		case RegistrationRequestSupportedCodecsType:
			a.SupportedCodecs = nasType.NewSupportedCodecs(ieiN)
			var lenN uint8
			binary.Read(buffer, binary.BigEndian, &lenN)
			a.SupportedCodecs.SetLen(uint16(lenN))
			binary.Read(buffer, binary.BigEndian, a.SupportedCodecs.Buffer[:lenN])
		case RegistrationRequestRequestedT3512ValueType:
			a.T3512Value = nasType.NewT3512Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3512Value.Len)
			a.T3512Value.SetLen(a.T3512Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3512Value.Octet)
		default:
		}
	}
}
