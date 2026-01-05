// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"encoding/hex"
	"fmt"

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

	// DEBUG
	fmt.Printf("[DEBUG] SnssaiToNas called: Sst=%d (0x%02x), Sd=%s\n",
		snssai.Sst, snssai.Sst, snssai.Sd)

	if snssai.Sd == "" {
		buf = append(buf, 0x01)
		buf = append(buf, uint8(snssai.Sst))
		fmt.Printf("[DEBUG] Output (no SD): %x\n", buf)
	} else {
		buf = append(buf, 0x04)
		sstByte := uint8(snssai.Sst)
		buf = append(buf, sstByte)
		fmt.Printf("[DEBUG] SST byte: %d -> 0x%02x\n", snssai.Sst, sstByte)

		if byteArray, err := hex.DecodeString(snssai.Sd); err != nil {
			logger.ConvertLog.Warnf("decode snssai.sd failed: %+v", err)
		} else {
			buf = append(buf, byteArray...)
		}
		fmt.Printf("[DEBUG] Output (with SD): %x\n", buf)
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
