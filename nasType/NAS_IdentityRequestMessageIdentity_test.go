// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"testing"

	"github.com/5GC-DEV/nas"
	"github.com/5GC-DEV/nas/nasType"

	"github.com/stretchr/testify/assert"
)

type nasTypeIdentityRequestMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeIdentityRequestMessageIdentityTable = []nasTypeIdentityRequestMessageIdentityData{
	{nas.MsgTypeIdentityRequest, nas.MsgTypeIdentityRequest},
}

func TestNasTypeNewIdentityRequestMessageIdentity(t *testing.T) {
	a := nasType.NewIdentityRequestMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeIdentityRequestMessageIdentity(t *testing.T) {
	a := nasType.NewIdentityRequestMessageIdentity()
	for _, table := range nasTypeIdentityRequestMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
