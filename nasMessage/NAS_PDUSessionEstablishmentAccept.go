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

type PDUSessionEstablishmentAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity
	nasType.SelectedSSCModeAndSelectedPDUSessionType
	nasType.AuthorizedQosRules
	nasType.SessionAMBR
	*nasType.Cause5GSM
	*nasType.PDUAddress
	*nasType.RQTimerValue
	*nasType.SNSSAI
	*nasType.AlwaysonPDUSessionIndication
	*nasType.MappedEPSBearerContexts
	*nasType.EAPMessage
	*nasType.AuthorizedQosFlowDescriptions
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.DNN
	*nasType.FiveGSMNetworkFeatureSupport
	*nasType.ServingPLMNRateControl
	*nasType.ATSSSContainer
	*nasType.ControlPlaneOnlyIndication
	*nasType.IPHeaderCompressionConfiguration
	*nasType.EthernetHeaderCompressionConfiguration
	*nasType.ServiceLevelAAContainer
	*nasType.ReceivedMBSContainer
	*nasType.N3QAI
}

func NewPDUSessionEstablishmentAccept(iei uint8) (pDUSessionEstablishmentAccept *PDUSessionEstablishmentAccept) {
	pDUSessionEstablishmentAccept = &PDUSessionEstablishmentAccept{}
	return pDUSessionEstablishmentAccept
}

const (
	PDUSessionEstablishmentAcceptCause5GSMType                            uint8 = 0x59
	PDUSessionEstablishmentAcceptPDUAddressType                           uint8 = 0x29
	PDUSessionEstablishmentAcceptRQTimerValueType                         uint8 = 0x56
	PDUSessionEstablishmentAcceptSNSSAIType                               uint8 = 0x22
	PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType         uint8 = 0x08
	PDUSessionEstablishmentAcceptMappedEPSBearerContextsType              uint8 = 0x75
	PDUSessionEstablishmentAcceptEAPMessageType                           uint8 = 0x78
	PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType        uint8 = 0x79
	PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionEstablishmentAcceptDNNType                                  uint8 = 0x25
	PDUSessionEstablishmentAcceptFiveGSMNetworkFeatureSupportType         uint8 = 0x17
	PDUSessionEstablishmentAcceptServingPLMNRateControlType               uint8 = 0x18
	PDUSessionEstablishmentAcceptATSSSContainerType                       uint8 = 0x77
	PDUSessionEstablishmentAcceptControlPlaneOnlyIndicationType           uint8 = 0x0C
	PDUSessionEstablishmentAcceptIPHeaderCompressionConfigurationType     uint8 = 0x66
	PDUSessionEstablishmentAcceptEthernetHeaderCompressionConfigType      uint8 = 0x1F
	PDUSessionEstablishmentAcceptServiceLevelAAContainerType              uint8 = 0x72
	PDUSessionEstablishmentAcceptReceivedMBSContainerType                 uint8 = 0x71
	PDUSessionEstablishmentAcceptN3QAIType                                uint8 = 0x70
)

func (a *PDUSessionEstablishmentAccept) EncodePDUSessionEstablishmentAccept(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SelectedSSCModeAndSelectedPDUSessionType.Octet)
	binary.Write(buffer, binary.BigEndian, a.AuthorizedQosRules.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.AuthorizedQosRules.Buffer)
	binary.Write(buffer, binary.BigEndian, a.SessionAMBR.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.SessionAMBR.Octet)
	if a.Cause5GSM != nil {
		binary.Write(buffer, binary.BigEndian, a.Cause5GSM.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
	}
	if a.PDUAddress != nil {
		binary.Write(buffer, binary.BigEndian, a.PDUAddress.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PDUAddress.GetLen())
		binary.Write(buffer, binary.BigEndian, a.PDUAddress.Octet[:a.PDUAddress.GetLen()])
	}
	if a.RQTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.RQTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.RQTimerValue.Octet)
	}
	if a.SNSSAI != nil {
		binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen())
		binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()])
	}
	if a.AlwaysonPDUSessionIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.AlwaysonPDUSessionIndication.Octet)
	}
	if a.MappedEPSBearerContexts != nil {
		binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetIei())
		binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Buffer)
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
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
	if a.DNN != nil {
		binary.Write(buffer, binary.BigEndian, a.DNN.GetIei())
		binary.Write(buffer, binary.BigEndian, a.DNN.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.DNN.Buffer)
	}
	if a.FiveGSMNetworkFeatureSupport != nil {
		binary.Write(buffer, binary.BigEndian, a.FiveGSMNetworkFeatureSupport.GetIei())
		binary.Write(buffer, binary.BigEndian, a.FiveGSMNetworkFeatureSupport.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.FiveGSMNetworkFeatureSupport.Buffer)
	}
	if a.ServingPLMNRateControl != nil {
		binary.Write(buffer, binary.BigEndian, a.ServingPLMNRateControl.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ServingPLMNRateControl.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ServingPLMNRateControl.Buffer)
	}
	if a.ATSSSContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.ATSSSContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ATSSSContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ATSSSContainer.Buffer)
	}
	if a.ControlPlaneOnlyIndication != nil {
		binary.Write(buffer, binary.BigEndian, &a.ControlPlaneOnlyIndication.Octet)
	}
	if a.IPHeaderCompressionConfiguration != nil {
		binary.Write(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.GetIei())
		binary.Write(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.IPHeaderCompressionConfiguration.Buffer)
	}
	if a.EthernetHeaderCompressionConfiguration != nil {
		binary.Write(buffer, binary.BigEndian, a.EthernetHeaderCompressionConfiguration.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EthernetHeaderCompressionConfiguration.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Octet)
	}
	if a.ServiceLevelAAContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer)
	}
	if a.ReceivedMBSContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.ReceivedMBSContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ReceivedMBSContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ReceivedMBSContainer.Buffer)
	}
	if a.N3QAI != nil {
		binary.Write(buffer, binary.BigEndian, a.N3QAI.GetIei())
		binary.Write(buffer, binary.BigEndian, a.N3QAI.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.N3QAI.Buffer)
	}
}

func (a *PDUSessionEstablishmentAccept) DecodePDUSessionEstablishmentAccept(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SelectedSSCModeAndSelectedPDUSessionType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosRules.Len)
	a.AuthorizedQosRules.SetLen(a.AuthorizedQosRules.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosRules.Buffer)
	binary.Read(buffer, binary.BigEndian, &a.SessionAMBR.Len)
	a.SessionAMBR.SetLen(a.SessionAMBR.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.SessionAMBR.Octet)
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
		case PDUSessionEstablishmentAcceptCause5GSMType:
			a.Cause5GSM = nasType.NewCause5GSM(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
		case PDUSessionEstablishmentAcceptPDUAddressType:
			a.PDUAddress = nasType.NewPDUAddress(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUAddress.Len)
			a.PDUAddress.SetLen(a.PDUAddress.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUAddress.Octet[:a.PDUAddress.GetLen()])
		case PDUSessionEstablishmentAcceptRQTimerValueType:
			a.RQTimerValue = nasType.NewRQTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RQTimerValue.Octet)
		case PDUSessionEstablishmentAcceptSNSSAIType:
			a.SNSSAI = nasType.NewSNSSAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len)
			a.SNSSAI.SetLen(a.SNSSAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()])
		case PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType:
			a.AlwaysonPDUSessionIndication = nasType.NewAlwaysonPDUSessionIndication(ieiN)
			a.AlwaysonPDUSessionIndication.Octet = ieiN
		case PDUSessionEstablishmentAcceptMappedEPSBearerContextsType:
			a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Len)
			a.MappedEPSBearerContexts.SetLen(a.MappedEPSBearerContexts.GetLen())
			binary.Read(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer[:a.MappedEPSBearerContexts.GetLen()])
		case PDUSessionEstablishmentAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType:
			a.AuthorizedQosFlowDescriptions = nasType.NewAuthorizedQosFlowDescriptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosFlowDescriptions.Len)
			a.AuthorizedQosFlowDescriptions.SetLen(a.AuthorizedQosFlowDescriptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.Buffer[:a.AuthorizedQosFlowDescriptions.GetLen()])
		case PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		case PDUSessionEstablishmentAcceptDNNType:
			a.DNN = nasType.NewDNN(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.DNN.Len)
			a.DNN.SetLen(a.DNN.GetLen())
			binary.Read(buffer, binary.BigEndian, a.DNN.Buffer[:a.DNN.GetLen()])
		case PDUSessionEstablishmentAcceptFiveGSMNetworkFeatureSupportType:
			a.FiveGSMNetworkFeatureSupport = nasType.NewFiveGSMNetworkFeatureSupport(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.FiveGSMNetworkFeatureSupport.Len)
			a.FiveGSMNetworkFeatureSupport.SetLen(a.FiveGSMNetworkFeatureSupport.GetLen())
			binary.Read(buffer, binary.BigEndian, a.FiveGSMNetworkFeatureSupport.Buffer[:a.FiveGSMNetworkFeatureSupport.GetLen()])
		case PDUSessionEstablishmentAcceptServingPLMNRateControlType:
			a.ServingPLMNRateControl = nasType.NewServingPLMNRateControl(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServingPLMNRateControl.Len)
			a.ServingPLMNRateControl.SetLen(a.ServingPLMNRateControl.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServingPLMNRateControl.Buffer[:a.ServingPLMNRateControl.GetLen()])
		case PDUSessionEstablishmentAcceptATSSSContainerType:
			a.ATSSSContainer = nasType.NewATSSSContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ATSSSContainer.Len)
			a.ATSSSContainer.SetLen(a.ATSSSContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ATSSSContainer.Buffer[:a.ATSSSContainer.GetLen()])
		case PDUSessionEstablishmentAcceptControlPlaneOnlyIndicationType:
			a.ControlPlaneOnlyIndication = nasType.NewControlPlaneOnlyIndication(ieiN)
			a.ControlPlaneOnlyIndication.Octet = ieiN
		case PDUSessionEstablishmentAcceptIPHeaderCompressionConfigurationType:
			a.IPHeaderCompressionConfiguration = nasType.NewIPHeaderCompressionConfiguration(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.IPHeaderCompressionConfiguration.Len)
			a.IPHeaderCompressionConfiguration.SetLen(a.IPHeaderCompressionConfiguration.GetLen())
			binary.Read(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.Buffer[:a.IPHeaderCompressionConfiguration.GetLen()])
		case PDUSessionEstablishmentAcceptEthernetHeaderCompressionConfigType:
			a.EthernetHeaderCompressionConfiguration = nasType.NewEthernetHeaderCompressionConfiguration(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Len)
			a.EthernetHeaderCompressionConfiguration.SetLen(a.EthernetHeaderCompressionConfiguration.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Octet)
		case PDUSessionEstablishmentAcceptServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len)
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()])
		case PDUSessionEstablishmentAcceptReceivedMBSContainerType:
			a.ReceivedMBSContainer = nasType.NewReceivedMBSContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ReceivedMBSContainer.Len)
			a.ReceivedMBSContainer.SetLen(a.ReceivedMBSContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ReceivedMBSContainer.Buffer[:a.ReceivedMBSContainer.GetLen()])
		case PDUSessionEstablishmentAcceptN3QAIType:
			a.N3QAI = nasType.NewN3QAI(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.N3QAI.Len)
			a.N3QAI.SetLen(a.N3QAI.GetLen())
			binary.Read(buffer, binary.BigEndian, a.N3QAI.Buffer[:a.N3QAI.GetLen()])
		default:
		}
	}
}
