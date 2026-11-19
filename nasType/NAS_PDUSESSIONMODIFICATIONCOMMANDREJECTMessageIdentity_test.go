// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"reflect"
	"testing"

	"github.com/5GC-DEV/nas-cdac/v2"
	"github.com/5GC-DEV/nas-cdac/v2/nasType"
)

func TestNasTypeNewPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

type nasTypePDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionModificationCommandReject, nas.MsgTypePDUSessionModificationCommandReject},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
