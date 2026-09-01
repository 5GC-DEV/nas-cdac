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

type PDUSessionModificationCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONMODIFICATIONCOMMANDMessageIdentity
	*nasType.Cause5GSM
	*nasType.SessionAMBR
	*nasType.RQTimerValue
	*nasType.AlwaysonPDUSessionIndication
	*nasType.AuthorizedQosRules
	*nasType.MappedEPSBearerContexts
	*nasType.AuthorizedQosFlowDescriptions
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.ATSSSContainer
	*nasType.IPHeaderCompressionConfiguration
	*nasType.ServingPLMNRateControl
	*nasType.EthernetHeaderCompressionConfiguration
	*nasType.ReceivedMBSContainer
	*nasType.ServiceLevelAAContainer
	AlternativeSNSSAI *nasType.SNSSAI
	*nasType.N3QAI
}

func NewPDUSessionModificationCommand(iei uint8) (pDUSessionModificationCommand *PDUSessionModificationCommand) {
	pDUSessionModificationCommand = &PDUSessionModificationCommand{}
	return pDUSessionModificationCommand
}

const (
	PDUSessionModificationCommandCause5GSMType                            uint8 = 0x59
	PDUSessionModificationCommandSessionAMBRType                          uint8 = 0x2A
	PDUSessionModificationCommandRQTimerValueType                         uint8 = 0x56
	PDUSessionModificationCommandAlwaysonPDUSessionIndicationType         uint8 = 0x08
	PDUSessionModificationCommandAuthorizedQosRulesType                   uint8 = 0x7A
	PDUSessionModificationCommandMappedEPSBearerContextsType              uint8 = 0x7F
	PDUSessionModificationCommandAuthorizedQosFlowDescriptionsType        uint8 = 0x79
	PDUSessionModificationCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionModificationCommandATSSSContainerType                       uint8 = 0x77
	PDUSessionModificationCommandIPHeaderCompressionConfigurationType     uint8 = 0x66
	PDUSessionModificationCommandServingPLMNRateControlType               uint8 = 0x1E
	PDUSessionModificationCommandEthernetHeaderCompressionConfigType      uint8 = 0x1F
	PDUSessionModificationCommandReceivedMBSContainerType                 uint8 = 0x71
	PDUSessionModificationCommandServiceLevelAAContainerType              uint8 = 0x72
	PDUSessionModificationCommandAlternativeSNSSAIType                    uint8 = 0x5A
	PDUSessionModificationCommandN3QAIType                                uint8 = 0x70
)

func (a *PDUSessionModificationCommand) EncodePDUSessionModificationCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONCOMMANDMessageIdentity.Octet)
	if a.Cause5GSM != nil {
		binary.Write(buffer, binary.BigEndian, a.Cause5GSM.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
	}
	if a.SessionAMBR != nil {
		binary.Write(buffer, binary.BigEndian, a.SessionAMBR.GetIei())
		binary.Write(buffer, binary.BigEndian, a.SessionAMBR.GetLen())
		binary.Write(buffer, binary.BigEndian, a.SessionAMBR.Octet[:a.SessionAMBR.GetLen()])
	}
	if a.RQTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.RQTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.RQTimerValue.Octet)
	}
	if a.AlwaysonPDUSessionIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.AlwaysonPDUSessionIndication.Octet)
	}
	if a.AuthorizedQosRules != nil {
		binary.Write(buffer, binary.BigEndian, a.AuthorizedQosRules.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AuthorizedQosRules.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AuthorizedQosRules.Buffer)
	}
	if a.MappedEPSBearerContexts != nil {
		binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetIei())
		binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Buffer)
	}
	if a.AuthorizedQosFlowDescriptions != nil {
		binary.Write(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AuthorizedQosFlowDescriptions.Buffer)
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer)
	}
	if a.ATSSSContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.ATSSSContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ATSSSContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ATSSSContainer.Buffer)
	}
	if a.IPHeaderCompressionConfiguration != nil {
		binary.Write(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.GetIei())
		binary.Write(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.IPHeaderCompressionConfiguration.Buffer)
	}
	if a.ServingPLMNRateControl != nil {
		binary.Write(buffer, binary.BigEndian, a.ServingPLMNRateControl.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ServingPLMNRateControl.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ServingPLMNRateControl.Buffer)
	}
	if a.EthernetHeaderCompressionConfiguration != nil {
		binary.Write(buffer, binary.BigEndian, a.EthernetHeaderCompressionConfiguration.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EthernetHeaderCompressionConfiguration.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Octet)
	}
	if a.ReceivedMBSContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.ReceivedMBSContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ReceivedMBSContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ReceivedMBSContainer.Buffer)
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
	if a.N3QAI != nil {
		binary.Write(buffer, binary.BigEndian, a.N3QAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.N3QAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.N3QAI.Buffer)
	}
}

func (a *PDUSessionModificationCommand) DecodePDUSessionModificationCommand(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONCOMMANDMessageIdentity.Octet)
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
		case PDUSessionModificationCommandCause5GSMType:
			a.Cause5GSM = nasType.NewCause5GSM(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
		case PDUSessionModificationCommandSessionAMBRType:
			a.SessionAMBR = nasType.NewSessionAMBR(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SessionAMBR.Len)
			a.SessionAMBR.SetLen(a.SessionAMBR.GetLen())
			binary.Read(buffer, binary.BigEndian, a.SessionAMBR.Octet[:a.SessionAMBR.GetLen()])
		case PDUSessionModificationCommandRQTimerValueType:
			a.RQTimerValue = nasType.NewRQTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RQTimerValue.Octet)
		case PDUSessionModificationCommandAlwaysonPDUSessionIndicationType:
			a.AlwaysonPDUSessionIndication = nasType.NewAlwaysonPDUSessionIndication(ieiN)
			a.AlwaysonPDUSessionIndication.Octet = ieiN
		case PDUSessionModificationCommandAuthorizedQosRulesType:
			a.AuthorizedQosRules = nasType.NewAuthorizedQosRules(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosRules.Len)
			a.AuthorizedQosRules.SetLen(a.AuthorizedQosRules.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AuthorizedQosRules.Buffer[:a.AuthorizedQosRules.GetLen()])
		case PDUSessionModificationCommandMappedEPSBearerContextsType:
			a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Len)
			a.MappedEPSBearerContexts.SetLen(a.MappedEPSBearerContexts.GetLen())
			binary.Read(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer[:a.MappedEPSBearerContexts.GetLen()])
		case PDUSessionModificationCommandAuthorizedQosFlowDescriptionsType:
			a.AuthorizedQosFlowDescriptions = nasType.NewAuthorizedQosFlowDescriptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosFlowDescriptions.Len)
			a.AuthorizedQosFlowDescriptions.SetLen(a.AuthorizedQosFlowDescriptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.Buffer[:a.AuthorizedQosFlowDescriptions.GetLen()])
		case PDUSessionModificationCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		case PDUSessionModificationCommandATSSSContainerType:
			a.ATSSSContainer = nasType.NewATSSSContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ATSSSContainer.Len)
			a.ATSSSContainer.SetLen(a.ATSSSContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ATSSSContainer.Buffer[:a.ATSSSContainer.GetLen()])
		case PDUSessionModificationCommandIPHeaderCompressionConfigurationType:
			a.IPHeaderCompressionConfiguration = nasType.NewIPHeaderCompressionConfiguration(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.IPHeaderCompressionConfiguration.Len)
			a.IPHeaderCompressionConfiguration.SetLen(a.IPHeaderCompressionConfiguration.GetLen())
			binary.Read(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.Buffer[:a.IPHeaderCompressionConfiguration.GetLen()])
		case PDUSessionModificationCommandServingPLMNRateControlType:
			a.ServingPLMNRateControl = nasType.NewServingPLMNRateControl(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServingPLMNRateControl.Len)
			a.ServingPLMNRateControl.SetLen(a.ServingPLMNRateControl.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServingPLMNRateControl.Buffer[:a.ServingPLMNRateControl.GetLen()])
		case PDUSessionModificationCommandEthernetHeaderCompressionConfigType:
			a.EthernetHeaderCompressionConfiguration = nasType.NewEthernetHeaderCompressionConfiguration(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Len)
			a.EthernetHeaderCompressionConfiguration.SetLen(a.EthernetHeaderCompressionConfiguration.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Octet)
		case PDUSessionModificationCommandReceivedMBSContainerType:
			a.ReceivedMBSContainer = nasType.NewReceivedMBSContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ReceivedMBSContainer.Len)
			a.ReceivedMBSContainer.SetLen(a.ReceivedMBSContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ReceivedMBSContainer.Buffer[:a.ReceivedMBSContainer.GetLen()])
		case PDUSessionModificationCommandServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len)
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()])
		case PDUSessionModificationCommandAlternativeSNSSAIType:
			a.AlternativeSNSSAI = nasType.NewSNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AlternativeSNSSAI.Len)
			a.AlternativeSNSSAI.SetLen(a.AlternativeSNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AlternativeSNSSAI.Octet[:a.AlternativeSNSSAI.GetLen()])
		case PDUSessionModificationCommandN3QAIType:
			a.N3QAI = nasType.NewN3QAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.N3QAI.Len)
			a.N3QAI.SetLen(a.N3QAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.N3QAI.Buffer[:a.N3QAI.GetLen()])
		default:
		}
	}
}
