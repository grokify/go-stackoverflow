/*
 * Stack Exchange API
 *
 * Stack Exchange API
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package stackoverflow

type UsersResponse struct {
	Items          []User `json:"items,omitempty"`
	HasMore        bool   `json:"has_more,omitempty"`
	QuoteMax       int32  `json:"quote_max,omitempty"`
	QuotaRemaining int32  `json:"quota_remaining,omitempty"`
}
