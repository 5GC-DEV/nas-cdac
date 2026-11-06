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

type nasTypeConfigurationUpdateCompleteMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeConfigurationUpdateCompleteMessageIdentityTable = []nasTypeConfigurationUpdateCompleteMessageIdentityData{
	{nas.MsgTypeConfigurationUpdateComplete, nas.MsgTypeConfigurationUpdateComplete},
}

func TestNasTypeNewConfigurationUpdateCompleteMessageIdentity(t *testing.T) {
	a := nasType.NewConfigurationUpdateCompleteMessageIdentity()
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeGetSetConfigurationUpdateCompleteMessageIdentity(t *testing.T) {
	a := nasType.NewConfigurationUpdateCompleteMessageIdentity()
	for _, table := range nasTypeConfigurationUpdateCompleteMessageIdentityTable {
		a.SetMessageType(table.in)
		if !reflect.DeepEqual(table.out, a.GetMessageType()) {
			t.Errorf("Not equal: expected %v, got %v", table.out, a.GetMessageType())
		}
	}
}
