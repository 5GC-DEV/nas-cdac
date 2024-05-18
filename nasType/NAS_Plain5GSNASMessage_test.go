// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"testing"

	"github.com/5gc-dev/nas/nasType"

	"github.com/stretchr/testify/assert"
)

func TestNasTypeNewPlain5GSNASMessage(t *testing.T) {
	a := nasType.NewPlain5GSNASMessage()
	assert.NotNil(t, a)
}
