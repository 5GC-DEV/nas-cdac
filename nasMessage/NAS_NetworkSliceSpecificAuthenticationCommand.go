// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// NetworkSliceSpecificAuthenticationCommand 8.2.31
type NetworkSliceSpecificAuthenticationCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity
	nasType.SNSSAI
	nasType.EAPMessage
}

func NewNetworkSliceSpecificAuthenticationCommand(iei uint8) (networkSliceSpecificAuthenticationCommand *NetworkSliceSpecificAuthenticationCommand) {
	networkSliceSpecificAuthenticationCommand = &NetworkSliceSpecificAuthenticationCommand{}
	return networkSliceSpecificAuthenticationCommand
}

func (a *NetworkSliceSpecificAuthenticationCommand) EncodeNetworkSliceSpecificAuthenticationCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen())
	binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()])
	binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}

func (a *NetworkSliceSpecificAuthenticationCommand) DecodeNetworkSliceSpecificAuthenticationCommand(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len)
	binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()])
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}
