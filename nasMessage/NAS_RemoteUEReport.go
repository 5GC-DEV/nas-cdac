// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RemoteUEReport 8.3.19
type RemoteUEReport struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.REMOTEUEREPORTMessageIdentity
	*nasType.RemoteUEContextList                              // Remote UE context connected (IEI 0x76)
	RemoteUEContextDisconnected  *nasType.RemoteUEContextList // Remote UE context disconnected (IEI 0x70)
}

func NewRemoteUEReport(iei uint8) (remoteUEReport *RemoteUEReport) {
	remoteUEReport = &RemoteUEReport{}
	return remoteUEReport
}

const (
	RemoteUEReportRemoteUEContextConnectedType    uint8 = 0x76
	RemoteUEReportRemoteUEContextDisconnectedType uint8 = 0x70
)

func (a *RemoteUEReport) EncodeRemoteUEReport(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.REMOTEUEREPORTMessageIdentity.Octet)
	if a.RemoteUEContextList != nil {
		binary.Write(buffer, binary.BigEndian, a.RemoteUEContextList.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RemoteUEContextList.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RemoteUEContextList.Buffer)
	}
	if a.RemoteUEContextDisconnected != nil {
		binary.Write(buffer, binary.BigEndian, a.RemoteUEContextDisconnected.GetIei())
		binary.Write(buffer, binary.BigEndian, a.RemoteUEContextDisconnected.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.RemoteUEContextDisconnected.Buffer)
	}
}

func (a *RemoteUEReport) DecodeRemoteUEReport(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.REMOTEUEREPORTMessageIdentity.Octet)
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
		case RemoteUEReportRemoteUEContextConnectedType:
			a.RemoteUEContextList = nasType.NewRemoteUEContextList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RemoteUEContextList.Len)
			a.RemoteUEContextList.SetLen(a.RemoteUEContextList.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RemoteUEContextList.Buffer)
		case RemoteUEReportRemoteUEContextDisconnectedType:
			a.RemoteUEContextDisconnected = nasType.NewRemoteUEContextList(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.RemoteUEContextDisconnected.Len)
			a.RemoteUEContextDisconnected.SetLen(a.RemoteUEContextDisconnected.GetLen())
			binary.Read(buffer, binary.BigEndian, a.RemoteUEContextDisconnected.Buffer)
		default:
		}
	}
}
