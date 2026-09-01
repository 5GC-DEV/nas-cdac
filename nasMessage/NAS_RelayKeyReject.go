// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RelayKeyReject 8.2.36
type RelayKeyReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RELAYKEYREJECTMessageIdentity
	nasType.ProSeRelayTransactionIdentity
	*nasType.EAPMessage
}

func NewRelayKeyReject(iei uint8) (relayKeyReject *RelayKeyReject) {
	relayKeyReject = &RelayKeyReject{}
	return relayKeyReject
}

const (
	RelayKeyRejectEAPMessageType uint8 = 0x78
)

func (a *RelayKeyReject) EncodeRelayKeyReject(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RELAYKEYREJECTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
}

func (a *RelayKeyReject) DecodeRelayKeyReject(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RELAYKEYREJECTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
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
		case RelayKeyRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		default:
		}
	}
}
