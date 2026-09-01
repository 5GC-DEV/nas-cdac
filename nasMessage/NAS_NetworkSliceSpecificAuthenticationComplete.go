// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// NetworkSliceSpecificAuthenticationComplete 8.2.32
type NetworkSliceSpecificAuthenticationComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity
	nasType.SNSSAI
	nasType.EAPMessage
}

func NewNetworkSliceSpecificAuthenticationComplete(iei uint8) (networkSliceSpecificAuthenticationComplete *NetworkSliceSpecificAuthenticationComplete) {
	networkSliceSpecificAuthenticationComplete = &NetworkSliceSpecificAuthenticationComplete{}
	return networkSliceSpecificAuthenticationComplete
}

func (a *NetworkSliceSpecificAuthenticationComplete) EncodeNetworkSliceSpecificAuthenticationComplete(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen())
	binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()])
	binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}

func (a *NetworkSliceSpecificAuthenticationComplete) DecodeNetworkSliceSpecificAuthenticationComplete(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len)
	binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()])
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}
