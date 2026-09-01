// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// ServiceLevelAuthenticationComplete 8.3.18
type ServiceLevelAuthenticationComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity
	nasType.ServiceLevelAAContainer
}

func NewServiceLevelAuthenticationComplete(iei uint8) (serviceLevelAuthenticationComplete *ServiceLevelAuthenticationComplete) {
	serviceLevelAuthenticationComplete = &ServiceLevelAuthenticationComplete{}
	return serviceLevelAuthenticationComplete
}

func (a *ServiceLevelAuthenticationComplete) EncodeServiceLevelAuthenticationComplete(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer)
}

func (a *ServiceLevelAuthenticationComplete) DecodeServiceLevelAuthenticationComplete(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len)
	a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer)
}
