// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RelayKeyAccept 8.2.35
type RelayKeyAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RELAYKEYACCEPTMessageIdentity
	nasType.ProSeRelayTransactionIdentity
	nasType.RelayKeyResponseParameters
	*nasType.EAPMessage
}

func NewRelayKeyAccept(iei uint8) (relayKeyAccept *RelayKeyAccept) {
	relayKeyAccept = &RelayKeyAccept{}
	return relayKeyAccept
}

const (
	RelayKeyAcceptEAPMessageType uint8 = 0x78
)

func (a *RelayKeyAccept) EncodeRelayKeyAccept(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RELAYKEYACCEPTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.RelayKeyResponseParameters.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.RelayKeyResponseParameters.Buffer)
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
}

func (a *RelayKeyAccept) DecodeRelayKeyAccept(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RELAYKEYACCEPTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RelayKeyResponseParameters.Len)
	a.RelayKeyResponseParameters.SetLen(a.RelayKeyResponseParameters.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.RelayKeyResponseParameters.Buffer)
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
		case RelayKeyAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		default:
		}
	}
}
