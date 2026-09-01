// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RelayKeyRequest 8.2.34
type RelayKeyRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RELAYKEYREQUESTMessageIdentity
	nasType.ProSeRelayTransactionIdentity
	nasType.RelayKeyRequestParameters
}

func NewRelayKeyRequest(iei uint8) (relayKeyRequest *RelayKeyRequest) {
	relayKeyRequest = &RelayKeyRequest{}
	return relayKeyRequest
}

func (a *RelayKeyRequest) EncodeRelayKeyRequest(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RELAYKEYREQUESTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.RelayKeyRequestParameters.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.RelayKeyRequestParameters.Buffer)
}

func (a *RelayKeyRequest) DecodeRelayKeyRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RELAYKEYREQUESTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RelayKeyRequestParameters.Len)
	a.RelayKeyRequestParameters.SetLen(a.RelayKeyRequestParameters.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.RelayKeyRequestParameters.Buffer)
}
