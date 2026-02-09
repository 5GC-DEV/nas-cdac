// SPDX-License-Identifier: Apache-2.0

/*
 * NSSF NS Selection
 *
 * Models manually added to support Rejected NSSAI (Strictly based on TS 24.501 Table 9.11.3.46.1)
 */
package models

type RejectCause string

const (
	// Maps to Cause Value 0 (0000)
	RejectCause_S_NSSAI_NOT_AVAILABLE_IN_CURRENT_PLMN_OR_SNPN RejectCause = "S_NSSAI_NOT_AVAILABLE_IN_CURRENT_PLMN_OR_SNPN"

	// Maps to Cause Value 1 (0001)
	RejectCause_S_NSSAI_NOT_AVAILABLE_IN_TA RejectCause = "S_NSSAI_NOT_AVAILABLE_IN_TA"

	// Maps to Cause Value 2 (0010)
	RejectCause_S_NSSAI_NOT_AVAILABLE_DUE_TO_FAILED_OR_REVOKED_NSAA RejectCause = "S_NSSAI_NOT_AVAILABLE_DUE_TO_FAILED_OR_REVOKED_NSAA"
)
