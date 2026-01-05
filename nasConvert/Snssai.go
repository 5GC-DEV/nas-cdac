// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"encoding/hex"

	"github.com/5GC-DEV/nas-cdac/logger"
	"github.com/5GC-DEV/nas-cdac/nasType"
	"github.com/5GC-DEV/openapi-cdac/models"
)

func SnssaiToModels(nasSnssai *nasType.SNSSAI) (snssai models.Snssai) {
	sD := nasSnssai.GetSD()
	snssai.Sd = hex.EncodeToString(sD[:])
	snssai.Sst = int32(nasSnssai.GetSST())
	return
}

func SnssaiToNas(snssai models.Snssai) []uint8 {
	var buf []uint8

	logger.ConvertLog.Infof("[SnssaiToNas] Input: Sst=%d, Sd='%s'", snssai.Sst, snssai.Sd)

	// Check if SD is valid hex
	if snssai.Sd != "" {
		if _, err := hex.DecodeString(snssai.Sd); err != nil {
			logger.ConvertLog.Warnf("[SnssaiToNas] Invalid SD hex string: %s", snssai.Sd)
		}
	}

	if snssai.Sd == "" {
		buf = append(buf, 0x01) // Length = 1 (SST only)
		buf = append(buf, uint8(snssai.Sst))
		logger.ConvertLog.Infof("[SnssaiToNas] No SD: output = %x", buf)
	} else {
		buf = append(buf, 0x04) // Length = 4 (SST + 3-byte SD)
		buf = append(buf, uint8(snssai.Sst))
		logger.ConvertLog.Infof("[SnssaiToNas] SST byte: %d -> 0x%02x", snssai.Sst, uint8(snssai.Sst))

		if byteArray, err := hex.DecodeString(snssai.Sd); err != nil {
			logger.ConvertLog.Warnf("[SnssaiToNas] decode snssai.sd failed: %+v", err)
		} else {
			buf = append(buf, byteArray...)
			logger.ConvertLog.Infof("[SnssaiToNas] SD bytes: %x", byteArray)
		}
		logger.ConvertLog.Infof("[SnssaiToNas] With SD: output = %x", buf)
	}

	return buf
}

func RejectedSnssaiToNas(snssai models.Snssai, rejectCause uint8) []uint8 {
	var rejectedSnssai []uint8

	if snssai.Sd == "" {
		rejectedSnssai = append(rejectedSnssai, (0x01<<4)+rejectCause)
		rejectedSnssai = append(rejectedSnssai, uint8(snssai.Sst))
	} else {
		rejectedSnssai = append(rejectedSnssai, (0x04<<4)+rejectCause)
		rejectedSnssai = append(rejectedSnssai, uint8(snssai.Sst))
		if sDBytes, err := hex.DecodeString(snssai.Sd); err != nil {
			logger.ConvertLog.Warnf("decode snssai.sd failed: %+v", err)
		} else {
			rejectedSnssai = append(rejectedSnssai, sDBytes...)
		}
	}

	return rejectedSnssai
}
