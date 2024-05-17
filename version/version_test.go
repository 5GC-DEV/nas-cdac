// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package version_test

import (
	"testing"

	"github.com/5GC-DEV/nas/version"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, "2020-03-31-01", version.GetVersion())
}
