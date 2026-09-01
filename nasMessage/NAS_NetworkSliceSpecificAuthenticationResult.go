// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// NetworkSliceSpecificAuthenticationResult 8.2.33
type NetworkSliceSpecificAuthenticationResult struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity
	nasType.SNSSAI
	nasType.EAPMessage
}

func NewNetworkSliceSpecificAuthenticationResult(iei uint8) (networkSliceSpecificAuthenticationResult *NetworkSliceSpecificAuthenticationResult) {
	networkSliceSpecificAuthenticationResult = &NetworkSliceSpecificAuthenticationResult{}
	return networkSliceSpecificAuthenticationResult
}

func (a *NetworkSliceSpecificAuthenticationResult) EncodeNetworkSliceSpecificAuthenticationResult(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen())
	binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()])
	binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}

func (a *NetworkSliceSpecificAuthenticationResult) DecodeNetworkSliceSpecificAuthenticationResult(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONRESULTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len)
	binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()])
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}
