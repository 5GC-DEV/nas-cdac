// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type RegistrationAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationAcceptMessageIdentity
	nasType.RegistrationResult5GS
	*nasType.GUTI5G
	*nasType.EquivalentPlmns
	*nasType.TAIList
	*nasType.AllowedNSSAI
	*nasType.RejectedNSSAI
	*nasType.ConfiguredNSSAI
	*nasType.NetworkFeatureSupport5GS
	*nasType.PDUSessionStatus
	*nasType.PDUSessionReactivationResult
	*nasType.PDUSessionReactivationResultErrorCause
	*nasType.LADNInformation
	*nasType.MICOIndication
	*nasType.NetworkSlicingIndication
	*nasType.ServiceAreaList
	*nasType.T3512Value
	*nasType.Non3GppDeregistrationTimerValue
	*nasType.T3502Value
	*nasType.EmergencyNumberList
	*nasType.ExtendedEmergencyNumberList
	*nasType.SORTransparentContainer
	*nasType.EAPMessage
	*nasType.NSSAIInclusionMode
	*nasType.OperatordefinedAccessCategoryDefinitions
	*nasType.NegotiatedDRXParameters
	*nasType.Non3GPPNWProvidedPolicies
	*nasType.EPSBearerContextStatus
	*nasType.ExtendedDRXParameters
	*nasType.UERadioCapabilityID
	*nasType.UERadioCapabilityIDDeletionIndicationIE
	*nasType.CipheringKeyData
	*nasType.CAGInformationList
	*nasType.TruncatedFiveGSTMSIConfiguration
	*nasType.ExtendedRejectedNSSAI
	*nasType.ServiceLevelAAContainer
	*nasType.FiveGSAdditionalRequestResult
	*nasType.NSSRGInformation
	*nasType.RegistrationWaitRange
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ListOfPLMNsForDisasterCondition
	*nasType.ExtendedCAGInformationList
	*nasType.NSAGInformation
	*nasType.PendingNSSAI
	NegotiatedNBN1ModeDRXParameters      *nasType.NBN1ModeDRXParameters
	NegotiatedWUSAssistanceInformation   *nasType.WUSAssistanceInformation
	NegotiatedPEIPSAssistanceInformation *nasType.PEIPSAssistanceInformation
	ForbiddenTAIRoaming                  *nasType.TAIList
	ForbiddenTAIRegionalProvision        *nasType.TAIList
	*nasType.EquivalentSNPNs
	*nasType.NID
	*nasType.RANTimingSynchronization
	*nasType.AlternativeNSSAI
	*nasType.DiscontinuousCoverageMaxTimeOffset
	*nasType.SNSSAITimeValidityInformation
	UnavailabilityPeriodDuration *nasType.LowerBoundTimerValue
	*nasType.FeatureAuthorizationIndication
}

func NewRegistrationAccept(iei uint8) (registrationAccept *RegistrationAccept) {
	registrationAccept = &RegistrationAccept{}
	return registrationAccept
}

const (
	RegistrationAcceptGUTI5GType                                   uint8 = 0x77
	RegistrationAcceptEquivalentPlmnsType                          uint8 = 0x4A
	RegistrationAcceptTAIListType                                  uint8 = 0x54
	RegistrationAcceptAllowedNSSAIType                             uint8 = 0x15
	RegistrationAcceptRejectedNSSAIType                            uint8 = 0x11
	RegistrationAcceptConfiguredNSSAIType                          uint8 = 0x31
	RegistrationAcceptNetworkFeatureSupport5GSType                 uint8 = 0x21
	RegistrationAcceptPDUSessionStatusType                         uint8 = 0x50
	RegistrationAcceptPDUSessionReactivationResultType             uint8 = 0x26
	RegistrationAcceptPDUSessionReactivationResultErrorCauseType   uint8 = 0x72
	RegistrationAcceptLADNInformationType                          uint8 = 0x79
	RegistrationAcceptMICOIndicationType                           uint8 = 0x0B
	RegistrationAcceptNetworkSlicingIndicationType                 uint8 = 0x09
	RegistrationAcceptServiceAreaListType                          uint8 = 0x27
	RegistrationAcceptT3512ValueType                               uint8 = 0x5E
	RegistrationAcceptNon3GppDeregistrationTimerValueType          uint8 = 0x5D
	RegistrationAcceptT3502ValueType                               uint8 = 0x16
	RegistrationAcceptEmergencyNumberListType                      uint8 = 0x34
	RegistrationAcceptExtendedEmergencyNumberListType              uint8 = 0x7A
	RegistrationAcceptSORTransparentContainerType                  uint8 = 0x73
	RegistrationAcceptEAPMessageType                               uint8 = 0x78
	RegistrationAcceptNSSAIInclusionModeType                       uint8 = 0x0A
	RegistrationAcceptOperatordefinedAccessCategoryDefinitionsType uint8 = 0x76
	RegistrationAcceptNegotiatedDRXParametersType                  uint8 = 0x51
	RegistrationAcceptEPSBearerContextStatusType                   uint8 = 0x60
	RegistrationAcceptExtendedDRXParametersType                    uint8 = 0x6E
	RegistrationAcceptUERadioCapabilityIDType                      uint8 = 0x67
	RegistrationAcceptUERadioCapabilityIDDeletionIndicationType    uint8 = 0x0E
	RegistrationAcceptCipheringKeyDataType                         uint8 = 0x74
	RegistrationAcceptCAGInformationListType                       uint8 = 0x75
	RegistrationAcceptTruncatedFiveGSTMSIConfigurationType         uint8 = 0x1B
	RegistrationAcceptExtendedRejectedNSSAIType                    uint8 = 0x68
	RegistrationAcceptServiceLevelAAContainerType                  uint8 = 0x7B
	RegistrationAcceptFiveGSAdditionalRequestResultType            uint8 = 0x35
	RegistrationAcceptNSSRGInformationType                         uint8 = 0x70
	RegistrationAcceptRegistrationWaitRangeType                    uint8 = 0x14
	RegistrationAcceptDisasterReturnWaitRangeType                  uint8 = 0x2C
	RegistrationAcceptListOfPLMNsForDisasterConditionType          uint8 = 0x13
	RegistrationAcceptExtendedCAGInformationListType               uint8 = 0x71
	RegistrationAcceptNSAGInformationType                          uint8 = 0x7C
	RegistrationAcceptNon3GPPNWProvidedPoliciesType                uint8 = 0x0D
	RegistrationAcceptPendingNSSAIType                             uint8 = 0x39
	RegistrationAcceptNegotiatedNBN1ModeDRXParametersType          uint8 = 0x29
	RegistrationAcceptNegotiatedWUSAssistanceInformationType       uint8 = 0x1C
	RegistrationAcceptNegotiatedPEIPSAssistanceInformationType     uint8 = 0x33
	RegistrationAcceptForbiddenTAIRoamingType                      uint8 = 0x1D
	RegistrationAcceptForbiddenTAIRegionalProvisionType            uint8 = 0x1E
	RegistrationAcceptEquivalentSNPNsType                          uint8 = 0x3D
	RegistrationAcceptNIDType                                      uint8 = 0x32
	RegistrationAcceptRANTimingSynchronizationType                 uint8 = 0x4B
	RegistrationAcceptAlternativeNSSAIType                         uint8 = 0x4C
	RegistrationAcceptDiscontinuousCoverageMaxTimeOffsetType       uint8 = 0x4F
	RegistrationAcceptSNSSAITimeValidityInformationType            uint8 = 0x5B
	RegistrationAcceptUnavailabilityPeriodDurationType             uint8 = 0x3C
	RegistrationAcceptFeatureAuthorizationIndicationType           uint8 = 0x5C
)

func (a *RegistrationAccept) EncodeRegistrationAccept(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RegistrationAcceptMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.RegistrationResult5GS.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet)
	if a.GUTI5G != nil {
		binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetIei())
		binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetLen())
		binary.Write(buffer, binary.BigEndian, a.GUTI5G.Octet[:a.GUTI5G.GetLen()])
	}
	if a.EquivalentPlmns != nil {
		binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.GetLen())
		binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.Octet[:a.EquivalentPlmns.GetLen()])
	}
	if a.TAIList != nil {
		binary.Write(buffer, binary.BigEndian, a.TAIList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.TAIList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.TAIList.Buffer)
	}
	if a.AllowedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AllowedNSSAI.Buffer)
	}
	if a.RejectedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RejectedNSSAI.Buffer)
	}
	if a.ConfiguredNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Buffer)
	}
	if a.NetworkFeatureSupport5GS != nil {
		binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.GetLen())
		binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.Octet[:a.NetworkFeatureSupport5GS.GetLen()])
	}
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
	if a.LADNInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.LADNInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.LADNInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.LADNInformation.Buffer)
	}
	if a.MICOIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.MICOIndication.Octet)
	}
	if a.NetworkSlicingIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.NetworkSlicingIndication.Octet)
	}
	if a.ServiceAreaList != nil {
		binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ServiceAreaList.Buffer)
	}
	if a.T3512Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3512Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3512Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3512Value.Octet)
	}
	if a.Non3GppDeregistrationTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.Non3GppDeregistrationTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, a.Non3GppDeregistrationTimerValue.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.Non3GppDeregistrationTimerValue.Octet)
	}
	if a.T3502Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3502Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3502Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3502Value.Octet)
	}
	if a.EmergencyNumberList != nil {
		binary.Write(buffer, binary.BigEndian, a.EmergencyNumberList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EmergencyNumberList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EmergencyNumberList.Buffer)
	}
	if a.ExtendedEmergencyNumberList != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedEmergencyNumberList.Buffer)
	}
	if a.SORTransparentContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.SORTransparentContainer.Buffer)
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
	if a.NSSAIInclusionMode != nil {
		binary.Write(buffer, binary.BigEndian, &a.NSSAIInclusionMode.Octet)
	}
	if a.OperatordefinedAccessCategoryDefinitions != nil {
		binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Buffer)
	}
	if a.NegotiatedDRXParameters != nil {
		binary.Write(buffer, binary.BigEndian, a.NegotiatedDRXParameters.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NegotiatedDRXParameters.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NegotiatedDRXParameters.Octet)
	}
	if a.Non3GPPNWProvidedPolicies != nil {
		binary.Write(buffer, binary.BigEndian, &a.Non3GPPNWProvidedPolicies.Octet)
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
	if a.UERadioCapabilityIDDeletionIndicationIE != nil {
		binary.Write(buffer, binary.BigEndian, &a.UERadioCapabilityIDDeletionIndicationIE.Octet)
	}
	if a.CipheringKeyData != nil {
		binary.Write(buffer, binary.BigEndian, a.CipheringKeyData.GetIei())
		binary.Write(buffer, binary.BigEndian, a.CipheringKeyData.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.CipheringKeyData.Buffer)
	}
	if a.CAGInformationList != nil {
		binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.CAGInformationList.Buffer)
	}
	if a.TruncatedFiveGSTMSIConfiguration != nil {
		binary.Write(buffer, binary.BigEndian, a.TruncatedFiveGSTMSIConfiguration.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.TruncatedFiveGSTMSIConfiguration.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.TruncatedFiveGSTMSIConfiguration.Buffer[:uint8(a.TruncatedFiveGSTMSIConfiguration.GetLen())])
	}
	if a.ExtendedRejectedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.ExtendedRejectedNSSAI.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:uint8(a.ExtendedRejectedNSSAI.GetLen())])
	}
	if a.ServiceLevelAAContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer)
	}
	if a.FiveGSAdditionalRequestResult != nil {
		binary.Write(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.FiveGSAdditionalRequestResult.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.Buffer[:uint8(a.FiveGSAdditionalRequestResult.GetLen())])
	}
	if a.NSSRGInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.NSSRGInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NSSRGInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NSSRGInformation.Buffer)
	}
	if a.RegistrationWaitRange != nil {
		binary.Write(buffer, binary.BigEndian, a.RegistrationWaitRange.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.RegistrationWaitRange.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.RegistrationWaitRange.Buffer[:uint8(a.RegistrationWaitRange.GetLen())])
	}
	if a.DisasterReturnWaitRange != nil {
		binary.Write(buffer, binary.BigEndian, a.DisasterReturnWaitRange.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.DisasterReturnWaitRange.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:uint8(a.DisasterReturnWaitRange.GetLen())])
	}
	if a.ListOfPLMNsForDisasterCondition != nil {
		binary.Write(buffer, binary.BigEndian, a.ListOfPLMNsForDisasterCondition.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.ListOfPLMNsForDisasterCondition.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.ListOfPLMNsForDisasterCondition.Buffer[:uint8(a.ListOfPLMNsForDisasterCondition.GetLen())])
	}
	if a.ExtendedCAGInformationList != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedCAGInformationList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedCAGInformationList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Buffer)
	}
	if a.NSAGInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.NSAGInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NSAGInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NSAGInformation.Buffer)
	}
	if a.PendingNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.PendingNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PendingNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PendingNSSAI.Buffer)
	}
	if a.NegotiatedNBN1ModeDRXParameters != nil {
		binary.Write(buffer, binary.BigEndian, a.NegotiatedNBN1ModeDRXParameters.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NegotiatedNBN1ModeDRXParameters.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NegotiatedNBN1ModeDRXParameters.Buffer)
	}
	if a.NegotiatedWUSAssistanceInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.NegotiatedWUSAssistanceInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.NegotiatedWUSAssistanceInformation.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.NegotiatedWUSAssistanceInformation.Buffer[:uint8(a.NegotiatedWUSAssistanceInformation.GetLen())])
	}
	if a.NegotiatedPEIPSAssistanceInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.NegotiatedPEIPSAssistanceInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.NegotiatedPEIPSAssistanceInformation.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.NegotiatedPEIPSAssistanceInformation.Buffer[:uint8(a.NegotiatedPEIPSAssistanceInformation.GetLen())])
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
	if a.EquivalentSNPNs != nil {
		binary.Write(buffer, binary.BigEndian, a.EquivalentSNPNs.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EquivalentSNPNs.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EquivalentSNPNs.Buffer)
	}
	if a.NID != nil {
		binary.Write(buffer, binary.BigEndian, a.NID.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.NID.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.NID.Buffer[:uint8(a.NID.GetLen())])
	}
	if a.RANTimingSynchronization != nil {
		binary.Write(buffer, binary.BigEndian, a.RANTimingSynchronization.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RANTimingSynchronization.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RANTimingSynchronization.Buffer)
	}
	if a.AlternativeNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.AlternativeNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AlternativeNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AlternativeNSSAI.Buffer)
	}
	if a.DiscontinuousCoverageMaxTimeOffset != nil {
		binary.Write(buffer, binary.BigEndian, a.DiscontinuousCoverageMaxTimeOffset.GetIei())
		binary.Write(buffer, binary.BigEndian, a.DiscontinuousCoverageMaxTimeOffset.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Octet)
	}
	if a.SNSSAITimeValidityInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.SNSSAITimeValidityInformation.Buffer)
	}
	if a.UnavailabilityPeriodDuration != nil {
		binary.Write(buffer, binary.BigEndian, a.UnavailabilityPeriodDuration.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UnavailabilityPeriodDuration.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UnavailabilityPeriodDuration.Octet)
	}
	if a.FeatureAuthorizationIndication != nil {
		binary.Write(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.GetIei())
		binary.Write(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.FeatureAuthorizationIndication.Buffer)
	}
}

func (a *RegistrationAccept) DecodeRegistrationAccept(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RegistrationAcceptMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Len)
	a.RegistrationResult5GS.SetLen(a.RegistrationResult5GS.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet)
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
		case RegistrationAcceptGUTI5GType:
			a.GUTI5G = nasType.NewGUTI5G(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.GUTI5G.Len)
			a.GUTI5G.SetLen(a.GUTI5G.GetLen())
			binary.Read(buffer, binary.BigEndian, a.GUTI5G.Octet[:a.GUTI5G.GetLen()])
		case RegistrationAcceptEquivalentPlmnsType:
			a.EquivalentPlmns = nasType.NewEquivalentPlmns(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EquivalentPlmns.Len)
			a.EquivalentPlmns.SetLen(a.EquivalentPlmns.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EquivalentPlmns.Octet[:a.EquivalentPlmns.GetLen()])
		case RegistrationAcceptTAIListType:
			a.TAIList = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.TAIList.Len)
			a.TAIList.SetLen(a.TAIList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.TAIList.Buffer[:a.TAIList.GetLen()])
		case RegistrationAcceptAllowedNSSAIType:
			a.AllowedNSSAI = nasType.NewAllowedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AllowedNSSAI.Len)
			a.AllowedNSSAI.SetLen(a.AllowedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AllowedNSSAI.Buffer[:a.AllowedNSSAI.GetLen()])
		case RegistrationAcceptRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len)
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer[:a.RejectedNSSAI.GetLen()])
		case RegistrationAcceptConfiguredNSSAIType:
			a.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Len)
			a.ConfiguredNSSAI.SetLen(a.ConfiguredNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ConfiguredNSSAI.Buffer[:a.ConfiguredNSSAI.GetLen()])
		case RegistrationAcceptNetworkFeatureSupport5GSType:
			a.NetworkFeatureSupport5GS = nasType.NewNetworkFeatureSupport5GS(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NetworkFeatureSupport5GS.Len)
			a.NetworkFeatureSupport5GS.SetLen(a.NetworkFeatureSupport5GS.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.Octet[:a.NetworkFeatureSupport5GS.GetLen()])
		case RegistrationAcceptPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len)
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()])
		case RegistrationAcceptPDUSessionReactivationResultType:
			a.PDUSessionReactivationResult = nasType.NewPDUSessionReactivationResult(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResult.Len)
			a.PDUSessionReactivationResult.SetLen(a.PDUSessionReactivationResult.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResult.Buffer[:a.PDUSessionReactivationResult.GetLen()])
		case RegistrationAcceptPDUSessionReactivationResultErrorCauseType:
			a.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResultErrorCause.Len)
			a.PDUSessionReactivationResultErrorCause.SetLen(a.PDUSessionReactivationResultErrorCause.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.Buffer[:a.PDUSessionReactivationResultErrorCause.GetLen()])
		case RegistrationAcceptLADNInformationType:
			a.LADNInformation = nasType.NewLADNInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.LADNInformation.Len)
			a.LADNInformation.SetLen(a.LADNInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.LADNInformation.Buffer[:a.LADNInformation.GetLen()])
		case RegistrationAcceptMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case RegistrationAcceptNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case RegistrationAcceptServiceAreaListType:
			a.ServiceAreaList = nasType.NewServiceAreaList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServiceAreaList.Len)
			a.ServiceAreaList.SetLen(a.ServiceAreaList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServiceAreaList.Buffer[:a.ServiceAreaList.GetLen()])
		case RegistrationAcceptT3512ValueType:
			a.T3512Value = nasType.NewT3512Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3512Value.Len)
			a.T3512Value.SetLen(a.T3512Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3512Value.Octet)
		case RegistrationAcceptNon3GppDeregistrationTimerValueType:
			a.Non3GppDeregistrationTimerValue = nasType.NewNon3GppDeregistrationTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Non3GppDeregistrationTimerValue.Len)
			a.Non3GppDeregistrationTimerValue.SetLen(a.Non3GppDeregistrationTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.Non3GppDeregistrationTimerValue.Octet)
		case RegistrationAcceptT3502ValueType:
			a.T3502Value = nasType.NewT3502Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3502Value.Len)
			a.T3502Value.SetLen(a.T3502Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3502Value.Octet)
		case RegistrationAcceptEmergencyNumberListType:
			a.EmergencyNumberList = nasType.NewEmergencyNumberList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EmergencyNumberList.Len)
			a.EmergencyNumberList.SetLen(a.EmergencyNumberList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EmergencyNumberList.Buffer[:a.EmergencyNumberList.GetLen()])
		case RegistrationAcceptExtendedEmergencyNumberListType:
			a.ExtendedEmergencyNumberList = nasType.NewExtendedEmergencyNumberList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedEmergencyNumberList.Len)
			a.ExtendedEmergencyNumberList.SetLen(a.ExtendedEmergencyNumberList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.Buffer[:a.ExtendedEmergencyNumberList.GetLen()])
		case RegistrationAcceptSORTransparentContainerType:
			a.SORTransparentContainer = nasType.NewSORTransparentContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SORTransparentContainer.Len)
			a.SORTransparentContainer.SetLen(a.SORTransparentContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.SORTransparentContainer.Buffer[:a.SORTransparentContainer.GetLen()])
		case RegistrationAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case RegistrationAcceptNSSAIInclusionModeType:
			a.NSSAIInclusionMode = nasType.NewNSSAIInclusionMode(ieiN)
			a.NSSAIInclusionMode.Octet = ieiN
		case RegistrationAcceptOperatordefinedAccessCategoryDefinitionsType:
			a.OperatordefinedAccessCategoryDefinitions = nasType.NewOperatordefinedAccessCategoryDefinitions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Len)
			a.OperatordefinedAccessCategoryDefinitions.SetLen(a.OperatordefinedAccessCategoryDefinitions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.Buffer[:a.OperatordefinedAccessCategoryDefinitions.GetLen()])
		case RegistrationAcceptNegotiatedDRXParametersType:
			a.NegotiatedDRXParameters = nasType.NewNegotiatedDRXParameters(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NegotiatedDRXParameters.Len)
			a.NegotiatedDRXParameters.SetLen(a.NegotiatedDRXParameters.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.NegotiatedDRXParameters.Octet)
		case RegistrationAcceptEPSBearerContextStatusType:
			a.EPSBearerContextStatus = nasType.NewEPSBearerContextStatus(ieiN)
			var lenN0 uint8
			binary.Read(buffer, binary.BigEndian, &lenN0)
			a.EPSBearerContextStatus.SetLen(uint16(lenN0))
			binary.Read(buffer, binary.BigEndian, a.EPSBearerContextStatus.Buffer[:lenN0])
		case RegistrationAcceptExtendedDRXParametersType:
			a.ExtendedDRXParameters = nasType.NewExtendedDRXParameters(ieiN)
			var lenN1 uint8
			binary.Read(buffer, binary.BigEndian, &lenN1)
			a.ExtendedDRXParameters.SetLen(uint16(lenN1))
			binary.Read(buffer, binary.BigEndian, a.ExtendedDRXParameters.Buffer[:lenN1])
		case RegistrationAcceptUERadioCapabilityIDType:
			a.UERadioCapabilityID = nasType.NewUERadioCapabilityID(ieiN)
			var lenN2 uint8
			binary.Read(buffer, binary.BigEndian, &lenN2)
			a.UERadioCapabilityID.SetLen(uint16(lenN2))
			binary.Read(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:lenN2])
		case RegistrationAcceptUERadioCapabilityIDDeletionIndicationType:
			a.UERadioCapabilityIDDeletionIndicationIE = nasType.NewUERadioCapabilityIDDeletionIndicationIE(ieiN)
			a.UERadioCapabilityIDDeletionIndicationIE.Octet = ieiN
		case RegistrationAcceptCipheringKeyDataType:
			a.CipheringKeyData = nasType.NewCipheringKeyData(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.CipheringKeyData.Len)
			a.CipheringKeyData.SetLen(a.CipheringKeyData.GetLen())
			binary.Read(buffer, binary.BigEndian, a.CipheringKeyData.Buffer[:a.CipheringKeyData.GetLen()])
		case RegistrationAcceptCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len)
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()])
		case RegistrationAcceptTruncatedFiveGSTMSIConfigurationType:
			a.TruncatedFiveGSTMSIConfiguration = nasType.NewTruncatedFiveGSTMSIConfiguration(ieiN)
			var lenN3 uint8
			binary.Read(buffer, binary.BigEndian, &lenN3)
			a.TruncatedFiveGSTMSIConfiguration.SetLen(uint16(lenN3))
			binary.Read(buffer, binary.BigEndian, a.TruncatedFiveGSTMSIConfiguration.Buffer[:lenN3])
		case RegistrationAcceptExtendedRejectedNSSAIType:
			a.ExtendedRejectedNSSAI = nasType.NewExtendedRejectedNSSAI(ieiN)
			var lenN4 uint8
			binary.Read(buffer, binary.BigEndian, &lenN4)
			a.ExtendedRejectedNSSAI.SetLen(uint16(lenN4))
			binary.Read(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:lenN4])
		case RegistrationAcceptServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len)
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()])
		case RegistrationAcceptFiveGSAdditionalRequestResultType:
			a.FiveGSAdditionalRequestResult = nasType.NewFiveGSAdditionalRequestResult(ieiN)
			var lenN5 uint8
			binary.Read(buffer, binary.BigEndian, &lenN5)
			a.FiveGSAdditionalRequestResult.SetLen(uint16(lenN5))
			binary.Read(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.Buffer[:lenN5])
		case RegistrationAcceptNSSRGInformationType:
			a.NSSRGInformation = nasType.NewNSSRGInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NSSRGInformation.Len)
			a.NSSRGInformation.SetLen(a.NSSRGInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NSSRGInformation.Buffer[:a.NSSRGInformation.GetLen()])
		case RegistrationAcceptRegistrationWaitRangeType:
			a.RegistrationWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var lenN6 uint8
			binary.Read(buffer, binary.BigEndian, &lenN6)
			a.RegistrationWaitRange.SetLen(uint16(lenN6))
			binary.Read(buffer, binary.BigEndian, a.RegistrationWaitRange.Buffer[:lenN6])
		case RegistrationAcceptDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var lenN6a uint8
			binary.Read(buffer, binary.BigEndian, &lenN6a)
			a.DisasterReturnWaitRange.SetLen(uint16(lenN6a))
			binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:lenN6a])
		case RegistrationAcceptListOfPLMNsForDisasterConditionType:
			a.ListOfPLMNsForDisasterCondition = nasType.NewListOfPLMNsForDisasterCondition(ieiN)
			var lenN7 uint8
			binary.Read(buffer, binary.BigEndian, &lenN7)
			a.ListOfPLMNsForDisasterCondition.SetLen(uint16(lenN7))
			binary.Read(buffer, binary.BigEndian, a.ListOfPLMNsForDisasterCondition.Buffer[:lenN7])
		case RegistrationAcceptExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len)
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()])
		case RegistrationAcceptNSAGInformationType:
			a.NSAGInformation = nasType.NewNSAGInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NSAGInformation.Len)
			a.NSAGInformation.SetLen(a.NSAGInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NSAGInformation.Buffer[:a.NSAGInformation.GetLen()])
		case RegistrationAcceptNon3GPPNWProvidedPoliciesType:
			a.Non3GPPNWProvidedPolicies = nasType.NewNon3GPPNWProvidedPolicies(ieiN)
			a.Non3GPPNWProvidedPolicies.Octet = ieiN
		case RegistrationAcceptPendingNSSAIType:
			a.PendingNSSAI = nasType.NewPendingNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PendingNSSAI.Len)
			a.PendingNSSAI.SetLen(a.PendingNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PendingNSSAI.Buffer[:a.PendingNSSAI.GetLen()])
		case RegistrationAcceptNegotiatedNBN1ModeDRXParametersType:
			a.NegotiatedNBN1ModeDRXParameters = nasType.NewNBN1ModeDRXParameters(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NegotiatedNBN1ModeDRXParameters.Len)
			a.NegotiatedNBN1ModeDRXParameters.SetLen(a.NegotiatedNBN1ModeDRXParameters.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NegotiatedNBN1ModeDRXParameters.Buffer[:a.NegotiatedNBN1ModeDRXParameters.GetLen()])
		case RegistrationAcceptNegotiatedWUSAssistanceInformationType:
			a.NegotiatedWUSAssistanceInformation = nasType.NewWUSAssistanceInformation(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.NegotiatedWUSAssistanceInformation.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.NegotiatedWUSAssistanceInformation.Buffer[:l])
		case RegistrationAcceptNegotiatedPEIPSAssistanceInformationType:
			a.NegotiatedPEIPSAssistanceInformation = nasType.NewPEIPSAssistanceInformation(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.NegotiatedPEIPSAssistanceInformation.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.NegotiatedPEIPSAssistanceInformation.Buffer[:l])
		case RegistrationAcceptForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len)
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()])
		case RegistrationAcceptForbiddenTAIRegionalProvisionType:
			a.ForbiddenTAIRegionalProvision = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Len)
			a.ForbiddenTAIRegionalProvision.SetLen(a.ForbiddenTAIRegionalProvision.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.Buffer[:a.ForbiddenTAIRegionalProvision.GetLen()])
		case RegistrationAcceptEquivalentSNPNsType:
			a.EquivalentSNPNs = nasType.NewEquivalentSNPNs(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EquivalentSNPNs.Len)
			a.EquivalentSNPNs.SetLen(a.EquivalentSNPNs.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EquivalentSNPNs.Buffer[:a.EquivalentSNPNs.GetLen()])
		case RegistrationAcceptNIDType:
			a.NID = nasType.NewNID(ieiN)
			var l uint8
			binary.Read(buffer, binary.BigEndian, &l)
			a.NID.SetLen(uint16(l))
			binary.Read(buffer, binary.BigEndian, a.NID.Buffer[:l])
		case RegistrationAcceptRANTimingSynchronizationType:
			a.RANTimingSynchronization = nasType.NewRANTimingSynchronization(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RANTimingSynchronization.Len)
			a.RANTimingSynchronization.SetLen(a.RANTimingSynchronization.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RANTimingSynchronization.Buffer[:a.RANTimingSynchronization.GetLen()])
		case RegistrationAcceptAlternativeNSSAIType:
			a.AlternativeNSSAI = nasType.NewAlternativeNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AlternativeNSSAI.Len)
			a.AlternativeNSSAI.SetLen(a.AlternativeNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AlternativeNSSAI.Buffer[:a.AlternativeNSSAI.GetLen()])
		case RegistrationAcceptDiscontinuousCoverageMaxTimeOffsetType:
			a.DiscontinuousCoverageMaxTimeOffset = nasType.NewDiscontinuousCoverageMaxTimeOffset(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Len)
			a.DiscontinuousCoverageMaxTimeOffset.SetLen(a.DiscontinuousCoverageMaxTimeOffset.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Octet)
		case RegistrationAcceptSNSSAITimeValidityInformationType:
			a.SNSSAITimeValidityInformation = nasType.NewSNSSAITimeValidityInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SNSSAITimeValidityInformation.Len)
			a.SNSSAITimeValidityInformation.SetLen(a.SNSSAITimeValidityInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.Buffer[:a.SNSSAITimeValidityInformation.GetLen()])
		case RegistrationAcceptUnavailabilityPeriodDurationType:
			a.UnavailabilityPeriodDuration = nasType.NewLowerBoundTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UnavailabilityPeriodDuration.Len)
			a.UnavailabilityPeriodDuration.SetLen(a.UnavailabilityPeriodDuration.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.UnavailabilityPeriodDuration.Octet)
		case RegistrationAcceptFeatureAuthorizationIndicationType:
			a.FeatureAuthorizationIndication = nasType.NewFeatureAuthorizationIndication(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.FeatureAuthorizationIndication.Len)
			a.FeatureAuthorizationIndication.SetLen(a.FeatureAuthorizationIndication.GetLen())
			binary.Read(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.Buffer[:a.FeatureAuthorizationIndication.GetLen()])
		default:
		}
	}
}
