// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RelayAuthenticationRequest 8.2.37
type RelayAuthenticationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RELAYAUTHENTICATIONREQUESTMessageIdentity
	nasType.ProSeRelayTransactionIdentity
	nasType.EAPMessage
}

func NewRelayAuthenticationRequest(iei uint8) (relayAuthenticationRequest *RelayAuthenticationRequest) {
	relayAuthenticationRequest = &RelayAuthenticationRequest{}
	return relayAuthenticationRequest
}

func (a *RelayAuthenticationRequest) EncodeRelayAuthenticationRequest(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RELAYAUTHENTICATIONREQUESTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}

func (a *RelayAuthenticationRequest) DecodeRelayAuthenticationRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RELAYAUTHENTICATIONREQUESTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}
