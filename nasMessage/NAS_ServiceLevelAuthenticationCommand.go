// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// ServiceLevelAuthenticationCommand 8.3.17
type ServiceLevelAuthenticationCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity
	nasType.ServiceLevelAAContainer
}

func NewServiceLevelAuthenticationCommand(iei uint8) (serviceLevelAuthenticationCommand *ServiceLevelAuthenticationCommand) {
	serviceLevelAuthenticationCommand = &ServiceLevelAuthenticationCommand{}
	return serviceLevelAuthenticationCommand
}

func (a *ServiceLevelAuthenticationCommand) EncodeServiceLevelAuthenticationCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer)
}

func (a *ServiceLevelAuthenticationCommand) DecodeServiceLevelAuthenticationCommand(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len)
	a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer)
}
