// SPDX-License-Identifier: Apache-2.0

/*
 * NSSF NS Selection
 *
 * Models manually added to support Rejected NSSAI (based on TS 29.531/29.571)
 *
 */
package models

type RejectedSnssai struct {
	RejectedSnssai *Snssai     `json:"snssai"`
	RejectCause    RejectCause `json:"rejectCause,omitempty"`
}
