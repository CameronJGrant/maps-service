// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Represents error object.
// Ref: #/components/schemas/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSubmissionFilter returns new OptSubmissionFilter with value set to v.
func NewOptSubmissionFilter(v SubmissionFilter) OptSubmissionFilter {
	return OptSubmissionFilter{
		Value: v,
		Set:   true,
	}
}

// OptSubmissionFilter is optional SubmissionFilter.
type OptSubmissionFilter struct {
	Value SubmissionFilter
	Set   bool
}

// IsSet returns true if OptSubmissionFilter was set.
func (o OptSubmissionFilter) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSubmissionFilter) Reset() {
	var v SubmissionFilter
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSubmissionFilter) SetTo(v SubmissionFilter) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSubmissionFilter) Get() (v SubmissionFilter, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSubmissionFilter) Or(d SubmissionFilter) SubmissionFilter {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Pagination
type Pagination struct {
	Page  int32 `json:"Page"`
	Limit int32 `json:"Limit"`
}

// GetPage returns the value of Page.
func (s *Pagination) GetPage() int32 {
	return s.Page
}

// GetLimit returns the value of Limit.
func (s *Pagination) GetLimit() int32 {
	return s.Limit
}

// SetPage sets the value of Page.
func (s *Pagination) SetPage(val int32) {
	s.Page = val
}

// SetLimit sets the value of Limit.
func (s *Pagination) SetLimit(val int32) {
	s.Limit = val
}

// PatchSubmissionCompletedOK is response for PatchSubmissionCompleted operation.
type PatchSubmissionCompletedOK struct{}

// PatchSubmissionModelOK is response for PatchSubmissionModel operation.
type PatchSubmissionModelOK struct{}

// PatchSubmissionStatusOK is response for PatchSubmissionStatus operation.
type PatchSubmissionStatusOK struct{}

// Ref: #/components/schemas/Submission
type Submission struct {
	ID             OptInt64  `json:"ID"`
	DisplayName    OptString `json:"DisplayName"`
	Creator        OptString `json:"Creator"`
	GameID         OptInt32  `json:"GameID"`
	Date           OptInt64  `json:"Date"`
	Submitter      OptInt64  `json:"Submitter"`
	AssetID        OptInt64  `json:"AssetID"`
	AssetVersion   OptInt64  `json:"AssetVersion"`
	Completed      OptBool   `json:"Completed"`
	SubmissionType OptInt32  `json:"SubmissionType"`
	TargetAssetID  OptInt64  `json:"TargetAssetID"`
	StatusID       OptInt32  `json:"StatusID"`
}

// GetID returns the value of ID.
func (s *Submission) GetID() OptInt64 {
	return s.ID
}

// GetDisplayName returns the value of DisplayName.
func (s *Submission) GetDisplayName() OptString {
	return s.DisplayName
}

// GetCreator returns the value of Creator.
func (s *Submission) GetCreator() OptString {
	return s.Creator
}

// GetGameID returns the value of GameID.
func (s *Submission) GetGameID() OptInt32 {
	return s.GameID
}

// GetDate returns the value of Date.
func (s *Submission) GetDate() OptInt64 {
	return s.Date
}

// GetSubmitter returns the value of Submitter.
func (s *Submission) GetSubmitter() OptInt64 {
	return s.Submitter
}

// GetAssetID returns the value of AssetID.
func (s *Submission) GetAssetID() OptInt64 {
	return s.AssetID
}

// GetAssetVersion returns the value of AssetVersion.
func (s *Submission) GetAssetVersion() OptInt64 {
	return s.AssetVersion
}

// GetCompleted returns the value of Completed.
func (s *Submission) GetCompleted() OptBool {
	return s.Completed
}

// GetSubmissionType returns the value of SubmissionType.
func (s *Submission) GetSubmissionType() OptInt32 {
	return s.SubmissionType
}

// GetTargetAssetID returns the value of TargetAssetID.
func (s *Submission) GetTargetAssetID() OptInt64 {
	return s.TargetAssetID
}

// GetStatusID returns the value of StatusID.
func (s *Submission) GetStatusID() OptInt32 {
	return s.StatusID
}

// SetID sets the value of ID.
func (s *Submission) SetID(val OptInt64) {
	s.ID = val
}

// SetDisplayName sets the value of DisplayName.
func (s *Submission) SetDisplayName(val OptString) {
	s.DisplayName = val
}

// SetCreator sets the value of Creator.
func (s *Submission) SetCreator(val OptString) {
	s.Creator = val
}

// SetGameID sets the value of GameID.
func (s *Submission) SetGameID(val OptInt32) {
	s.GameID = val
}

// SetDate sets the value of Date.
func (s *Submission) SetDate(val OptInt64) {
	s.Date = val
}

// SetSubmitter sets the value of Submitter.
func (s *Submission) SetSubmitter(val OptInt64) {
	s.Submitter = val
}

// SetAssetID sets the value of AssetID.
func (s *Submission) SetAssetID(val OptInt64) {
	s.AssetID = val
}

// SetAssetVersion sets the value of AssetVersion.
func (s *Submission) SetAssetVersion(val OptInt64) {
	s.AssetVersion = val
}

// SetCompleted sets the value of Completed.
func (s *Submission) SetCompleted(val OptBool) {
	s.Completed = val
}

// SetSubmissionType sets the value of SubmissionType.
func (s *Submission) SetSubmissionType(val OptInt32) {
	s.SubmissionType = val
}

// SetTargetAssetID sets the value of TargetAssetID.
func (s *Submission) SetTargetAssetID(val OptInt64) {
	s.TargetAssetID = val
}

// SetStatusID sets the value of StatusID.
func (s *Submission) SetStatusID(val OptInt32) {
	s.StatusID = val
}

// Ref: #/components/schemas/SubmissionFilter
type SubmissionFilter struct {
	ID          OptInt64  `json:"ID"`
	DisplayName OptString `json:"DisplayName"`
	Creator     OptString `json:"Creator"`
	GameID      OptInt32  `json:"GameID"`
	Date        OptInt64  `json:"Date"`
}

// GetID returns the value of ID.
func (s *SubmissionFilter) GetID() OptInt64 {
	return s.ID
}

// GetDisplayName returns the value of DisplayName.
func (s *SubmissionFilter) GetDisplayName() OptString {
	return s.DisplayName
}

// GetCreator returns the value of Creator.
func (s *SubmissionFilter) GetCreator() OptString {
	return s.Creator
}

// GetGameID returns the value of GameID.
func (s *SubmissionFilter) GetGameID() OptInt32 {
	return s.GameID
}

// GetDate returns the value of Date.
func (s *SubmissionFilter) GetDate() OptInt64 {
	return s.Date
}

// SetID sets the value of ID.
func (s *SubmissionFilter) SetID(val OptInt64) {
	s.ID = val
}

// SetDisplayName sets the value of DisplayName.
func (s *SubmissionFilter) SetDisplayName(val OptString) {
	s.DisplayName = val
}

// SetCreator sets the value of Creator.
func (s *SubmissionFilter) SetCreator(val OptString) {
	s.Creator = val
}

// SetGameID sets the value of GameID.
func (s *SubmissionFilter) SetGameID(val OptInt32) {
	s.GameID = val
}

// SetDate sets the value of Date.
func (s *SubmissionFilter) SetDate(val OptInt64) {
	s.Date = val
}
