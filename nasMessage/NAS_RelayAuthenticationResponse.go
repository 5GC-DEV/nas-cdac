// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RelayAuthenticationResponse 8.2.38
type RelayAuthenticationResponse struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RELAYAUTHENTICATIONRESPONSEMessageIdentity
	nasType.ProSeRelayTransactionIdentity
	nasType.EAPMessage
}

func NewRelayAuthenticationResponse(iei uint8) (relayAuthenticationResponse *RelayAuthenticationResponse) {
	relayAuthenticationResponse = &RelayAuthenticationResponse{}
	return relayAuthenticationResponse
}

func (a *RelayAuthenticationResponse) EncodeRelayAuthenticationResponse(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RELAYAUTHENTICATIONRESPONSEMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}

func (a *RelayAuthenticationResponse) DecodeRelayAuthenticationResponse(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RELAYAUTHENTICATIONRESPONSEMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
}
