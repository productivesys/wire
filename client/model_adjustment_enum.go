/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AdjustmentEnum : Adjustment  * `01` - Pricing Error * `03` - Extension Error * `04` - Item Not Accepted (Damaged) * `05` - Item Not Accepted (Quality) * `06` - Quantity Contested 07   Incorrect Product * `11` - Returns (Damaged) * `12` - Returns (Quality) * `59` - Item Not Received * `75` - Total Order Not Received * `81` - Credit as Agreed * `CM` - Covered by Credit Memo 
type AdjustmentEnum string

// List of AdjustmentEnum
const (
	_1 AdjustmentEnum = "1"
	_3 AdjustmentEnum = "3"
	_4 AdjustmentEnum = "4"
	_5 AdjustmentEnum = "5"
	_6 AdjustmentEnum = "6"
	_11 AdjustmentEnum = "11"
	_12 AdjustmentEnum = "12"
	_59 AdjustmentEnum = "59"
	_75 AdjustmentEnum = "75"
	_81 AdjustmentEnum = "81"
	CM AdjustmentEnum = "CM"
)