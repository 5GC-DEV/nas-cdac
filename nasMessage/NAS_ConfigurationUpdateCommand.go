// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type ConfigurationUpdateCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ConfigurationUpdateCommandMessageIdentity
	*nasType.ConfigurationUpdateIndication
	*nasType.GUTI5G
	*nasType.TAIList
	*nasType.AllowedNSSAI
	*nasType.ServiceAreaList
	*nasType.FullNameForNetwork
	*nasType.ShortNameForNetwork
	*nasType.LocalTimeZone
	*nasType.UniversalTimeAndLocalTimeZone
	*nasType.NetworkDaylightSavingTime
	*nasType.LADNInformation
	*nasType.MICOIndication
	*nasType.NetworkSlicingIndication
	*nasType.ConfiguredNSSAI
	*nasType.RejectedNSSAI
	*nasType.OperatordefinedAccessCategoryDefinitions
	*nasType.SMSIndication
	*nasType.CAGInformationList
	*nasType.UERadioCapabilityID
	*nasType.UERadioCapabilityIDDeletionIndicationIE
	*nasType.TruncatedFiveGSTMSIConfiguration
	*nasType.ExtendedRejectedNSSAI
	*nasType.ServiceLevelAAContainer
	*nasType.NSSRGInformation
	*nasType.RegistrationWaitRange
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ListOfPLMNsForDisasterCondition
	*nasType.ExtendedCAGInformationList
	*nasType.NSAGInformation
	*nasType.ExtendedLADNInformation
	*nasType.RegistrationResult5GS
	*nasType.AdditionalConfigurationIndication
	*nasType.UpdatedPEIPSAssistanceInformation
	*nasType.PriorityIndicator
	*nasType.RANTimingSynchronization
	*nasType.AlternativeNSSAI
	*nasType.SNSSAILocationValidityInformation
	*nasType.SNSSAITimeValidityInformation
	*nasType.DiscontinuousCoverageMaxTimeOffset
	*nasType.PartiallyAllowedNSSAI
	*nasType.PartiallyRejectedNSSAI
	*nasType.FeatureAuthorizationIndication
}

func NewConfigurationUpdateCommand(iei uint8) (configurationUpdateCommand *ConfigurationUpdateCommand) {
	configurationUpdateCommand = &ConfigurationUpdateCommand{}
	return configurationUpdateCommand
}

const (
	ConfigurationUpdateCommandConfigurationUpdateIndicationType            uint8 = 0x0D
	ConfigurationUpdateCommandGUTI5GType                                   uint8 = 0x77
	ConfigurationUpdateCommandTAIListType                                  uint8 = 0x54
	ConfigurationUpdateCommandAllowedNSSAIType                             uint8 = 0x15
	ConfigurationUpdateCommandServiceAreaListType                          uint8 = 0x27
	ConfigurationUpdateCommandFullNameForNetworkType                       uint8 = 0x43
	ConfigurationUpdateCommandShortNameForNetworkType                      uint8 = 0x45
	ConfigurationUpdateCommandLocalTimeZoneType                            uint8 = 0x46
	ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType            uint8 = 0x47
	ConfigurationUpdateCommandNetworkDaylightSavingTimeType                uint8 = 0x49
	ConfigurationUpdateCommandLADNInformationType                          uint8 = 0x79
	ConfigurationUpdateCommandMICOIndicationType                           uint8 = 0x0B
	ConfigurationUpdateCommandNetworkSlicingIndicationType                 uint8 = 0x09
	ConfigurationUpdateCommandConfiguredNSSAIType                          uint8 = 0x31
	ConfigurationUpdateCommandRejectedNSSAIType                            uint8 = 0x11
	ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType uint8 = 0x76
	ConfigurationUpdateCommandSMSIndicationType                            uint8 = 0x0F
	ConfigurationUpdateCommandCAGInformationListType                       uint8 = 0x75
	ConfigurationUpdateCommandUERadioCapabilityIDType                      uint8 = 0x67
	ConfigurationUpdateCommandUERadioCapabilityIDDeletionIndicationType    uint8 = 0x0A
	ConfigurationUpdateCommandTruncatedFiveGSTMSIConfigurationType         uint8 = 0x1B
	ConfigurationUpdateCommandExtendedRejectedNSSAIType                    uint8 = 0x68
	ConfigurationUpdateCommandServiceLevelAAContainerType                  uint8 = 0x72
	ConfigurationUpdateCommandNSSRGInformationType                         uint8 = 0x70
	ConfigurationUpdateCommandRegistrationWaitRangeType                    uint8 = 0x14
	ConfigurationUpdateCommandDisasterReturnWaitRangeType                  uint8 = 0x2C
	ConfigurationUpdateCommandListOfPLMNsForDisasterConditionType          uint8 = 0x13
	ConfigurationUpdateCommandExtendedCAGInformationListType               uint8 = 0x71
	ConfigurationUpdateCommandNSAGInformationType                          uint8 = 0x73
	ConfigurationUpdateCommandExtendedLADNInformationType                  uint8 = 0x78
	ConfigurationUpdateCommandRegistrationResult5GSType                    uint8 = 0x44
	ConfigurationUpdateCommandAdditionalConfigurationIndicationType        uint8 = 0x0C
	ConfigurationUpdateCommandUpdatedPEIPSAssistanceInformationType        uint8 = 0x1F
	ConfigurationUpdateCommandPriorityIndicatorType                        uint8 = 0x0E
	ConfigurationUpdateCommandRANTimingSynchronizationType                 uint8 = 0x4B
	ConfigurationUpdateCommandAlternativeNSSAIType                         uint8 = 0x4C
	ConfigurationUpdateCommandSNSSAILocationValidityInformationType        uint8 = 0x7D
	ConfigurationUpdateCommandSNSSAITimeValidityInformationType            uint8 = 0x5B
	ConfigurationUpdateCommandDiscontinuousCoverageMaxTimeOffsetType       uint8 = 0x4F
	ConfigurationUpdateCommandPartiallyAllowedNSSAIType                    uint8 = 0x74
	ConfigurationUpdateCommandPartiallyRejectedNSSAIType                   uint8 = 0x7A
	ConfigurationUpdateCommandFeatureAuthorizationIndicationType           uint8 = 0x5C
)

func (a *ConfigurationUpdateCommand) EncodeConfigurationUpdateCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ConfigurationUpdateCommandMessageIdentity.Octet)
	if a.ConfigurationUpdateIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.ConfigurationUpdateIndication.Octet)
	}
	if a.GUTI5G != nil {
		binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetIei())
		binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetLen())
		binary.Write(buffer, binary.BigEndian, a.GUTI5G.Octet[:a.GUTI5G.GetLen()])
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
	if a.ServiceAreaList != nil {
		binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ServiceAreaList.Buffer)
	}
	if a.FullNameForNetwork != nil {
		binary.Write(buffer, binary.BigEndian, a.FullNameForNetwork.GetIei())
		binary.Write(buffer, binary.BigEndian, a.FullNameForNetwork.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.FullNameForNetwork.Buffer)
	}
	if a.ShortNameForNetwork != nil {
		binary.Write(buffer, binary.BigEndian, a.ShortNameForNetwork.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ShortNameForNetwork.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ShortNameForNetwork.Buffer)
	}
	if a.LocalTimeZone != nil {
		binary.Write(buffer, binary.BigEndian, a.LocalTimeZone.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.LocalTimeZone.Octet)
	}
	if a.UniversalTimeAndLocalTimeZone != nil {
		binary.Write(buffer, binary.BigEndian, a.UniversalTimeAndLocalTimeZone.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.UniversalTimeAndLocalTimeZone.Octet)
	}
	if a.NetworkDaylightSavingTime != nil {
		binary.Write(buffer, binary.BigEndian, a.NetworkDaylightSavingTime.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NetworkDaylightSavingTime.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NetworkDaylightSavingTime.Octet)
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
	if a.ConfiguredNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Buffer)
	}
	if a.RejectedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RejectedNSSAI.Buffer)
	}
	if a.OperatordefinedAccessCategoryDefinitions != nil {
		binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Buffer)
	}
	if a.SMSIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.SMSIndication.Octet)
	}
	if a.CAGInformationList != nil {
		binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.CAGInformationList.Buffer)
	}
	if a.UERadioCapabilityID != nil {
		binary.Write(buffer, binary.BigEndian, a.UERadioCapabilityID.GetIei())
		binary.Write(buffer, binary.BigEndian, uint8(a.UERadioCapabilityID.GetLen()))
		binary.Write(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:uint8(a.UERadioCapabilityID.GetLen())])
	}
	if a.UERadioCapabilityIDDeletionIndicationIE != nil {
		binary.Write(buffer, binary.BigEndian, &a.UERadioCapabilityIDDeletionIndicationIE.Octet)
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
	if a.ExtendedLADNInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedLADNInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedLADNInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedLADNInformation.Buffer)
	}
	if a.RegistrationResult5GS != nil {
		binary.Write(buffer, binary.BigEndian, a.RegistrationResult5GS.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RegistrationResult5GS.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet)
	}
	if a.AdditionalConfigurationIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.AdditionalConfigurationIndication.Octet)
	}
	if a.UpdatedPEIPSAssistanceInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.UpdatedPEIPSAssistanceInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UpdatedPEIPSAssistanceInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UpdatedPEIPSAssistanceInformation.Buffer)
	}
	if a.PriorityIndicator != nil {
		binary.Write(buffer, binary.BigEndian, &a.PriorityIndicator.Octet)
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
	if a.SNSSAILocationValidityInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.SNSSAILocationValidityInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.SNSSAILocationValidityInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.SNSSAILocationValidityInformation.Buffer)
	}
	if a.SNSSAITimeValidityInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.SNSSAITimeValidityInformation.Buffer)
	}
	if a.DiscontinuousCoverageMaxTimeOffset != nil {
		binary.Write(buffer, binary.BigEndian, a.DiscontinuousCoverageMaxTimeOffset.GetIei())
		binary.Write(buffer, binary.BigEndian, a.DiscontinuousCoverageMaxTimeOffset.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Octet)
	}
	if a.PartiallyAllowedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.PartiallyAllowedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PartiallyAllowedNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PartiallyAllowedNSSAI.Buffer)
	}
	if a.PartiallyRejectedNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.PartiallyRejectedNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PartiallyRejectedNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PartiallyRejectedNSSAI.Buffer)
	}
	if a.FeatureAuthorizationIndication != nil {
		binary.Write(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.GetIei())
		binary.Write(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.FeatureAuthorizationIndication.Buffer)
	}
}

func (a *ConfigurationUpdateCommand) DecodeConfigurationUpdateCommand(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ConfigurationUpdateCommandMessageIdentity.Octet)
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
		case ConfigurationUpdateCommandConfigurationUpdateIndicationType:
			a.ConfigurationUpdateIndication = nasType.NewConfigurationUpdateIndication(ieiN)
			a.ConfigurationUpdateIndication.Octet = ieiN
		case ConfigurationUpdateCommandGUTI5GType:
			a.GUTI5G = nasType.NewGUTI5G(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.GUTI5G.Len)
			a.GUTI5G.SetLen(a.GUTI5G.GetLen())
			binary.Read(buffer, binary.BigEndian, a.GUTI5G.Octet[:a.GUTI5G.GetLen()])
		case ConfigurationUpdateCommandTAIListType:
			a.TAIList = nasType.NewTAIList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.TAIList.Len)
			a.TAIList.SetLen(a.TAIList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.TAIList.Buffer[:a.TAIList.GetLen()])
		case ConfigurationUpdateCommandAllowedNSSAIType:
			a.AllowedNSSAI = nasType.NewAllowedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AllowedNSSAI.Len)
			a.AllowedNSSAI.SetLen(a.AllowedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AllowedNSSAI.Buffer[:a.AllowedNSSAI.GetLen()])
		case ConfigurationUpdateCommandServiceAreaListType:
			a.ServiceAreaList = nasType.NewServiceAreaList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServiceAreaList.Len)
			a.ServiceAreaList.SetLen(a.ServiceAreaList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServiceAreaList.Buffer[:a.ServiceAreaList.GetLen()])
		case ConfigurationUpdateCommandFullNameForNetworkType:
			a.FullNameForNetwork = nasType.NewFullNameForNetwork(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.FullNameForNetwork.Len)
			a.FullNameForNetwork.SetLen(a.FullNameForNetwork.GetLen())
			binary.Read(buffer, binary.BigEndian, a.FullNameForNetwork.Buffer[:a.FullNameForNetwork.GetLen()])
		case ConfigurationUpdateCommandShortNameForNetworkType:
			a.ShortNameForNetwork = nasType.NewShortNameForNetwork(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ShortNameForNetwork.Len)
			a.ShortNameForNetwork.SetLen(a.ShortNameForNetwork.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ShortNameForNetwork.Buffer[:a.ShortNameForNetwork.GetLen()])
		case ConfigurationUpdateCommandLocalTimeZoneType:
			a.LocalTimeZone = nasType.NewLocalTimeZone(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.LocalTimeZone.Octet)
		case ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType:
			a.UniversalTimeAndLocalTimeZone = nasType.NewUniversalTimeAndLocalTimeZone(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UniversalTimeAndLocalTimeZone.Octet)
		case ConfigurationUpdateCommandNetworkDaylightSavingTimeType:
			a.NetworkDaylightSavingTime = nasType.NewNetworkDaylightSavingTime(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NetworkDaylightSavingTime.Len)
			a.NetworkDaylightSavingTime.SetLen(a.NetworkDaylightSavingTime.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.NetworkDaylightSavingTime.Octet)
		case ConfigurationUpdateCommandLADNInformationType:
			a.LADNInformation = nasType.NewLADNInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.LADNInformation.Len)
			a.LADNInformation.SetLen(a.LADNInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.LADNInformation.Buffer[:a.LADNInformation.GetLen()])
		case ConfigurationUpdateCommandMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case ConfigurationUpdateCommandNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case ConfigurationUpdateCommandConfiguredNSSAIType:
			a.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Len)
			a.ConfiguredNSSAI.SetLen(a.ConfiguredNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ConfiguredNSSAI.Buffer[:a.ConfiguredNSSAI.GetLen()])
		case ConfigurationUpdateCommandRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len)
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer[:a.RejectedNSSAI.GetLen()])
		case ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType:
			a.OperatordefinedAccessCategoryDefinitions = nasType.NewOperatordefinedAccessCategoryDefinitions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Len)
			a.OperatordefinedAccessCategoryDefinitions.SetLen(a.OperatordefinedAccessCategoryDefinitions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.Buffer[:a.OperatordefinedAccessCategoryDefinitions.GetLen()])
		case ConfigurationUpdateCommandSMSIndicationType:
			a.SMSIndication = nasType.NewSMSIndication(ieiN)
			a.SMSIndication.Octet = ieiN
		case ConfigurationUpdateCommandCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len)
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()])
		case ConfigurationUpdateCommandUERadioCapabilityIDType:
			a.UERadioCapabilityID = nasType.NewUERadioCapabilityID(ieiN)
			var lenN0 uint8
			binary.Read(buffer, binary.BigEndian, &lenN0)
			a.UERadioCapabilityID.SetLen(uint16(lenN0))
			binary.Read(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:lenN0])
		case ConfigurationUpdateCommandUERadioCapabilityIDDeletionIndicationType:
			a.UERadioCapabilityIDDeletionIndicationIE = nasType.NewUERadioCapabilityIDDeletionIndicationIE(ieiN)
			a.UERadioCapabilityIDDeletionIndicationIE.Octet = ieiN
		case ConfigurationUpdateCommandTruncatedFiveGSTMSIConfigurationType:
			a.TruncatedFiveGSTMSIConfiguration = nasType.NewTruncatedFiveGSTMSIConfiguration(ieiN)
			var lenN1 uint8
			binary.Read(buffer, binary.BigEndian, &lenN1)
			a.TruncatedFiveGSTMSIConfiguration.SetLen(uint16(lenN1))
			binary.Read(buffer, binary.BigEndian, a.TruncatedFiveGSTMSIConfiguration.Buffer[:lenN1])
		case ConfigurationUpdateCommandExtendedRejectedNSSAIType:
			a.ExtendedRejectedNSSAI = nasType.NewExtendedRejectedNSSAI(ieiN)
			var lenN2 uint8
			binary.Read(buffer, binary.BigEndian, &lenN2)
			a.ExtendedRejectedNSSAI.SetLen(uint16(lenN2))
			binary.Read(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:lenN2])
		case ConfigurationUpdateCommandServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len)
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()])
		case ConfigurationUpdateCommandNSSRGInformationType:
			a.NSSRGInformation = nasType.NewNSSRGInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NSSRGInformation.Len)
			a.NSSRGInformation.SetLen(a.NSSRGInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NSSRGInformation.Buffer[:a.NSSRGInformation.GetLen()])
		case ConfigurationUpdateCommandRegistrationWaitRangeType:
			a.RegistrationWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var lenN3 uint8
			binary.Read(buffer, binary.BigEndian, &lenN3)
			a.RegistrationWaitRange.SetLen(uint16(lenN3))
			binary.Read(buffer, binary.BigEndian, a.RegistrationWaitRange.Buffer[:lenN3])
		case ConfigurationUpdateCommandDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var lenN3a uint8
			binary.Read(buffer, binary.BigEndian, &lenN3a)
			a.DisasterReturnWaitRange.SetLen(uint16(lenN3a))
			binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:lenN3a])
		case ConfigurationUpdateCommandListOfPLMNsForDisasterConditionType:
			a.ListOfPLMNsForDisasterCondition = nasType.NewListOfPLMNsForDisasterCondition(ieiN)
			var lenN4 uint8
			binary.Read(buffer, binary.BigEndian, &lenN4)
			a.ListOfPLMNsForDisasterCondition.SetLen(uint16(lenN4))
			binary.Read(buffer, binary.BigEndian, a.ListOfPLMNsForDisasterCondition.Buffer[:lenN4])
		case ConfigurationUpdateCommandExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len)
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()])
		case ConfigurationUpdateCommandNSAGInformationType:
			a.NSAGInformation = nasType.NewNSAGInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NSAGInformation.Len)
			a.NSAGInformation.SetLen(a.NSAGInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NSAGInformation.Buffer[:a.NSAGInformation.GetLen()])
		case ConfigurationUpdateCommandExtendedLADNInformationType:
			a.ExtendedLADNInformation = nasType.NewExtendedLADNInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedLADNInformation.Len)
			a.ExtendedLADNInformation.SetLen(a.ExtendedLADNInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedLADNInformation.Buffer[:a.ExtendedLADNInformation.GetLen()])
		case ConfigurationUpdateCommandRegistrationResult5GSType:
			a.RegistrationResult5GS = nasType.NewRegistrationResult5GS(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Len)
			a.RegistrationResult5GS.SetLen(a.RegistrationResult5GS.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet)
		case ConfigurationUpdateCommandAdditionalConfigurationIndicationType:
			a.AdditionalConfigurationIndication = nasType.NewAdditionalConfigurationIndication(ieiN)
			a.AdditionalConfigurationIndication.Octet = ieiN
		case ConfigurationUpdateCommandUpdatedPEIPSAssistanceInformationType:
			a.UpdatedPEIPSAssistanceInformation = nasType.NewUpdatedPEIPSAssistanceInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UpdatedPEIPSAssistanceInformation.Len)
			a.UpdatedPEIPSAssistanceInformation.SetLen(a.UpdatedPEIPSAssistanceInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.UpdatedPEIPSAssistanceInformation.Buffer[:a.UpdatedPEIPSAssistanceInformation.GetLen()])
		case ConfigurationUpdateCommandPriorityIndicatorType:
			a.PriorityIndicator = nasType.NewPriorityIndicator(ieiN)
			a.PriorityIndicator.Octet = ieiN
		case ConfigurationUpdateCommandRANTimingSynchronizationType:
			a.RANTimingSynchronization = nasType.NewRANTimingSynchronization(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RANTimingSynchronization.Len)
			a.RANTimingSynchronization.SetLen(a.RANTimingSynchronization.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RANTimingSynchronization.Buffer[:a.RANTimingSynchronization.GetLen()])
		case ConfigurationUpdateCommandAlternativeNSSAIType:
			a.AlternativeNSSAI = nasType.NewAlternativeNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AlternativeNSSAI.Len)
			a.AlternativeNSSAI.SetLen(a.AlternativeNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AlternativeNSSAI.Buffer[:a.AlternativeNSSAI.GetLen()])
		case ConfigurationUpdateCommandSNSSAILocationValidityInformationType:
			a.SNSSAILocationValidityInformation = nasType.NewSNSSAILocationValidityInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SNSSAILocationValidityInformation.Len)
			a.SNSSAILocationValidityInformation.SetLen(a.SNSSAILocationValidityInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.SNSSAILocationValidityInformation.Buffer[:a.SNSSAILocationValidityInformation.GetLen()])
		case ConfigurationUpdateCommandSNSSAITimeValidityInformationType:
			a.SNSSAITimeValidityInformation = nasType.NewSNSSAITimeValidityInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SNSSAITimeValidityInformation.Len)
			a.SNSSAITimeValidityInformation.SetLen(a.SNSSAITimeValidityInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.Buffer[:a.SNSSAITimeValidityInformation.GetLen()])
		case ConfigurationUpdateCommandDiscontinuousCoverageMaxTimeOffsetType:
			a.DiscontinuousCoverageMaxTimeOffset = nasType.NewDiscontinuousCoverageMaxTimeOffset(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Len)
			a.DiscontinuousCoverageMaxTimeOffset.SetLen(a.DiscontinuousCoverageMaxTimeOffset.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Octet)
		case ConfigurationUpdateCommandPartiallyAllowedNSSAIType:
			a.PartiallyAllowedNSSAI = nasType.NewPartiallyAllowedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PartiallyAllowedNSSAI.Len)
			a.PartiallyAllowedNSSAI.SetLen(a.PartiallyAllowedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PartiallyAllowedNSSAI.Buffer[:a.PartiallyAllowedNSSAI.GetLen()])
		case ConfigurationUpdateCommandPartiallyRejectedNSSAIType:
			a.PartiallyRejectedNSSAI = nasType.NewPartiallyRejectedNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PartiallyRejectedNSSAI.Len)
			a.PartiallyRejectedNSSAI.SetLen(a.PartiallyRejectedNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PartiallyRejectedNSSAI.Buffer[:a.PartiallyRejectedNSSAI.GetLen()])
		case ConfigurationUpdateCommandFeatureAuthorizationIndicationType:
			a.FeatureAuthorizationIndication = nasType.NewFeatureAuthorizationIndication(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.FeatureAuthorizationIndication.Len)
			a.FeatureAuthorizationIndication.SetLen(a.FeatureAuthorizationIndication.GetLen())
			binary.Read(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.Buffer[:a.FeatureAuthorizationIndication.GetLen()])
		default:
		}
	}
}
