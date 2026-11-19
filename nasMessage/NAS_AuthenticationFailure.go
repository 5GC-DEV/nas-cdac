// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/5GC-DEV/nas-cdac/v2/nasType"
)

type AuthenticationFailure struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationFailureMessageIdentity
	nasType.Cause5GMM
	*nasType.AuthenticationFailureParameter
}

func NewAuthenticationFailure(iei uint8) (authenticationFailure *AuthenticationFailure) {
	authenticationFailure = &AuthenticationFailure{}
	return authenticationFailure
}

const (
	AuthenticationFailureAuthenticationFailureParameterType uint8 = 0x30
)

func (a *AuthenticationFailure) EncodeAuthenticationFailure(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.AuthenticationFailureMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return
	}
	if a.AuthenticationFailureParameter != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureParameter.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureParameter.Octet[:a.GetLen()]); err != nil {
			return
		}
	}
}

func (a *AuthenticationFailure) DecodeAuthenticationFailure(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationFailureMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return
		}
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		switch tmpIeiN {
		case AuthenticationFailureAuthenticationFailureParameterType:
			a.AuthenticationFailureParameter = nasType.NewAuthenticationFailureParameter(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
				return
			}
			a.SetLen(a.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AuthenticationFailureParameter.Octet[:a.GetLen()]); err != nil {
				return
			}
		default:
		}
	}
}
