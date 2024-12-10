// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// ActionSubmissionPublishOK is response for ActionSubmissionPublish operation.
type ActionSubmissionPublishOK struct{}

// ActionSubmissionRejectOK is response for ActionSubmissionReject operation.
type ActionSubmissionRejectOK struct{}

// ActionSubmissionRequestChangesOK is response for ActionSubmissionRequestChanges operation.
type ActionSubmissionRequestChangesOK struct{}

// ActionSubmissionRevokeOK is response for ActionSubmissionRevoke operation.
type ActionSubmissionRevokeOK struct{}

// ActionSubmissionSubmitOK is response for ActionSubmissionSubmit operation.
type ActionSubmissionSubmitOK struct{}

// ActionSubmissionTriggerPublishOK is response for ActionSubmissionTriggerPublish operation.
type ActionSubmissionTriggerPublishOK struct{}

// ActionSubmissionTriggerValidateOK is response for ActionSubmissionTriggerValidate operation.
type ActionSubmissionTriggerValidateOK struct{}

// ActionSubmissionValidateOK is response for ActionSubmissionValidate operation.
type ActionSubmissionValidateOK struct{}

type CookieAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *CookieAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *CookieAuth) SetAPIKey(val string) {
	s.APIKey = val
}

// DeleteScriptOK is response for DeleteScript operation.
type DeleteScriptOK struct{}

// DeleteScriptPolicyOK is response for DeleteScriptPolicy operation.
type DeleteScriptPolicyOK struct{}

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

// Ref: #/components/schemas/Id
type ID struct {
	ID int64 `json:"ID"`
}

// GetID returns the value of ID.
func (s *ID) GetID() int64 {
	return s.ID
}

// SetID sets the value of ID.
func (s *ID) SetID(val int64) {
	s.ID = val
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

// Ref: #/components/schemas/Script
type Script struct {
	ID           int64  `json:"ID"`
	Hash         string `json:"Hash"`
	Source       string `json:"Source"`
	SubmissionID int64  `json:"SubmissionID"`
}

// GetID returns the value of ID.
func (s *Script) GetID() int64 {
	return s.ID
}

// GetHash returns the value of Hash.
func (s *Script) GetHash() string {
	return s.Hash
}

// GetSource returns the value of Source.
func (s *Script) GetSource() string {
	return s.Source
}

// GetSubmissionID returns the value of SubmissionID.
func (s *Script) GetSubmissionID() int64 {
	return s.SubmissionID
}

// SetID sets the value of ID.
func (s *Script) SetID(val int64) {
	s.ID = val
}

// SetHash sets the value of Hash.
func (s *Script) SetHash(val string) {
	s.Hash = val
}

// SetSource sets the value of Source.
func (s *Script) SetSource(val string) {
	s.Source = val
}

// SetSubmissionID sets the value of SubmissionID.
func (s *Script) SetSubmissionID(val int64) {
	s.SubmissionID = val
}

// Ref: #/components/schemas/ScriptCreate
type ScriptCreate struct {
	Source       string   `json:"Source"`
	SubmissionID OptInt64 `json:"SubmissionID"`
}

// GetSource returns the value of Source.
func (s *ScriptCreate) GetSource() string {
	return s.Source
}

// GetSubmissionID returns the value of SubmissionID.
func (s *ScriptCreate) GetSubmissionID() OptInt64 {
	return s.SubmissionID
}

// SetSource sets the value of Source.
func (s *ScriptCreate) SetSource(val string) {
	s.Source = val
}

// SetSubmissionID sets the value of SubmissionID.
func (s *ScriptCreate) SetSubmissionID(val OptInt64) {
	s.SubmissionID = val
}

// Ref: #/components/schemas/ScriptPolicy
type ScriptPolicy struct {
	ID             int64  `json:"ID"`
	FromScriptHash string `json:"FromScriptHash"`
	ToScriptID     int64  `json:"ToScriptID"`
	Policy         int32  `json:"Policy"`
}

// GetID returns the value of ID.
func (s *ScriptPolicy) GetID() int64 {
	return s.ID
}

// GetFromScriptHash returns the value of FromScriptHash.
func (s *ScriptPolicy) GetFromScriptHash() string {
	return s.FromScriptHash
}

// GetToScriptID returns the value of ToScriptID.
func (s *ScriptPolicy) GetToScriptID() int64 {
	return s.ToScriptID
}

// GetPolicy returns the value of Policy.
func (s *ScriptPolicy) GetPolicy() int32 {
	return s.Policy
}

// SetID sets the value of ID.
func (s *ScriptPolicy) SetID(val int64) {
	s.ID = val
}

// SetFromScriptHash sets the value of FromScriptHash.
func (s *ScriptPolicy) SetFromScriptHash(val string) {
	s.FromScriptHash = val
}

// SetToScriptID sets the value of ToScriptID.
func (s *ScriptPolicy) SetToScriptID(val int64) {
	s.ToScriptID = val
}

// SetPolicy sets the value of Policy.
func (s *ScriptPolicy) SetPolicy(val int32) {
	s.Policy = val
}

// Ref: #/components/schemas/ScriptPolicyCreate
type ScriptPolicyCreate struct {
	FromScriptID int64 `json:"FromScriptID"`
	ToScriptID   int64 `json:"ToScriptID"`
	Policy       int32 `json:"Policy"`
}

// GetFromScriptID returns the value of FromScriptID.
func (s *ScriptPolicyCreate) GetFromScriptID() int64 {
	return s.FromScriptID
}

// GetToScriptID returns the value of ToScriptID.
func (s *ScriptPolicyCreate) GetToScriptID() int64 {
	return s.ToScriptID
}

// GetPolicy returns the value of Policy.
func (s *ScriptPolicyCreate) GetPolicy() int32 {
	return s.Policy
}

// SetFromScriptID sets the value of FromScriptID.
func (s *ScriptPolicyCreate) SetFromScriptID(val int64) {
	s.FromScriptID = val
}

// SetToScriptID sets the value of ToScriptID.
func (s *ScriptPolicyCreate) SetToScriptID(val int64) {
	s.ToScriptID = val
}

// SetPolicy sets the value of Policy.
func (s *ScriptPolicyCreate) SetPolicy(val int32) {
	s.Policy = val
}

// Ref: #/components/schemas/ScriptPolicyUpdate
type ScriptPolicyUpdate struct {
	ID           int64    `json:"ID"`
	FromScriptID OptInt64 `json:"FromScriptID"`
	ToScriptID   OptInt64 `json:"ToScriptID"`
	Policy       OptInt32 `json:"Policy"`
}

// GetID returns the value of ID.
func (s *ScriptPolicyUpdate) GetID() int64 {
	return s.ID
}

// GetFromScriptID returns the value of FromScriptID.
func (s *ScriptPolicyUpdate) GetFromScriptID() OptInt64 {
	return s.FromScriptID
}

// GetToScriptID returns the value of ToScriptID.
func (s *ScriptPolicyUpdate) GetToScriptID() OptInt64 {
	return s.ToScriptID
}

// GetPolicy returns the value of Policy.
func (s *ScriptPolicyUpdate) GetPolicy() OptInt32 {
	return s.Policy
}

// SetID sets the value of ID.
func (s *ScriptPolicyUpdate) SetID(val int64) {
	s.ID = val
}

// SetFromScriptID sets the value of FromScriptID.
func (s *ScriptPolicyUpdate) SetFromScriptID(val OptInt64) {
	s.FromScriptID = val
}

// SetToScriptID sets the value of ToScriptID.
func (s *ScriptPolicyUpdate) SetToScriptID(val OptInt64) {
	s.ToScriptID = val
}

// SetPolicy sets the value of Policy.
func (s *ScriptPolicyUpdate) SetPolicy(val OptInt32) {
	s.Policy = val
}

// Ref: #/components/schemas/ScriptUpdate
type ScriptUpdate struct {
	ID           int64     `json:"ID"`
	Source       OptString `json:"Source"`
	SubmissionID OptInt64  `json:"SubmissionID"`
}

// GetID returns the value of ID.
func (s *ScriptUpdate) GetID() int64 {
	return s.ID
}

// GetSource returns the value of Source.
func (s *ScriptUpdate) GetSource() OptString {
	return s.Source
}

// GetSubmissionID returns the value of SubmissionID.
func (s *ScriptUpdate) GetSubmissionID() OptInt64 {
	return s.SubmissionID
}

// SetID sets the value of ID.
func (s *ScriptUpdate) SetID(val int64) {
	s.ID = val
}

// SetSource sets the value of Source.
func (s *ScriptUpdate) SetSource(val OptString) {
	s.Source = val
}

// SetSubmissionID sets the value of SubmissionID.
func (s *ScriptUpdate) SetSubmissionID(val OptInt64) {
	s.SubmissionID = val
}

// Ref: #/components/schemas/Submission
type Submission struct {
	ID             int64    `json:"ID"`
	DisplayName    string   `json:"DisplayName"`
	Creator        string   `json:"Creator"`
	GameID         int32    `json:"GameID"`
	Date           int64    `json:"Date"`
	Submitter      int64    `json:"Submitter"`
	AssetID        int64    `json:"AssetID"`
	AssetVersion   int64    `json:"AssetVersion"`
	Completed      bool     `json:"Completed"`
	SubmissionType int32    `json:"SubmissionType"`
	TargetAssetID  OptInt64 `json:"TargetAssetID"`
	StatusID       int32    `json:"StatusID"`
}

// GetID returns the value of ID.
func (s *Submission) GetID() int64 {
	return s.ID
}

// GetDisplayName returns the value of DisplayName.
func (s *Submission) GetDisplayName() string {
	return s.DisplayName
}

// GetCreator returns the value of Creator.
func (s *Submission) GetCreator() string {
	return s.Creator
}

// GetGameID returns the value of GameID.
func (s *Submission) GetGameID() int32 {
	return s.GameID
}

// GetDate returns the value of Date.
func (s *Submission) GetDate() int64 {
	return s.Date
}

// GetSubmitter returns the value of Submitter.
func (s *Submission) GetSubmitter() int64 {
	return s.Submitter
}

// GetAssetID returns the value of AssetID.
func (s *Submission) GetAssetID() int64 {
	return s.AssetID
}

// GetAssetVersion returns the value of AssetVersion.
func (s *Submission) GetAssetVersion() int64 {
	return s.AssetVersion
}

// GetCompleted returns the value of Completed.
func (s *Submission) GetCompleted() bool {
	return s.Completed
}

// GetSubmissionType returns the value of SubmissionType.
func (s *Submission) GetSubmissionType() int32 {
	return s.SubmissionType
}

// GetTargetAssetID returns the value of TargetAssetID.
func (s *Submission) GetTargetAssetID() OptInt64 {
	return s.TargetAssetID
}

// GetStatusID returns the value of StatusID.
func (s *Submission) GetStatusID() int32 {
	return s.StatusID
}

// SetID sets the value of ID.
func (s *Submission) SetID(val int64) {
	s.ID = val
}

// SetDisplayName sets the value of DisplayName.
func (s *Submission) SetDisplayName(val string) {
	s.DisplayName = val
}

// SetCreator sets the value of Creator.
func (s *Submission) SetCreator(val string) {
	s.Creator = val
}

// SetGameID sets the value of GameID.
func (s *Submission) SetGameID(val int32) {
	s.GameID = val
}

// SetDate sets the value of Date.
func (s *Submission) SetDate(val int64) {
	s.Date = val
}

// SetSubmitter sets the value of Submitter.
func (s *Submission) SetSubmitter(val int64) {
	s.Submitter = val
}

// SetAssetID sets the value of AssetID.
func (s *Submission) SetAssetID(val int64) {
	s.AssetID = val
}

// SetAssetVersion sets the value of AssetVersion.
func (s *Submission) SetAssetVersion(val int64) {
	s.AssetVersion = val
}

// SetCompleted sets the value of Completed.
func (s *Submission) SetCompleted(val bool) {
	s.Completed = val
}

// SetSubmissionType sets the value of SubmissionType.
func (s *Submission) SetSubmissionType(val int32) {
	s.SubmissionType = val
}

// SetTargetAssetID sets the value of TargetAssetID.
func (s *Submission) SetTargetAssetID(val OptInt64) {
	s.TargetAssetID = val
}

// SetStatusID sets the value of StatusID.
func (s *Submission) SetStatusID(val int32) {
	s.StatusID = val
}

// Ref: #/components/schemas/SubmissionCreate
type SubmissionCreate struct {
	DisplayName   string   `json:"DisplayName"`
	Creator       string   `json:"Creator"`
	GameID        int32    `json:"GameID"`
	AssetID       int64    `json:"AssetID"`
	AssetVersion  int64    `json:"AssetVersion"`
	TargetAssetID OptInt64 `json:"TargetAssetID"`
}

// GetDisplayName returns the value of DisplayName.
func (s *SubmissionCreate) GetDisplayName() string {
	return s.DisplayName
}

// GetCreator returns the value of Creator.
func (s *SubmissionCreate) GetCreator() string {
	return s.Creator
}

// GetGameID returns the value of GameID.
func (s *SubmissionCreate) GetGameID() int32 {
	return s.GameID
}

// GetAssetID returns the value of AssetID.
func (s *SubmissionCreate) GetAssetID() int64 {
	return s.AssetID
}

// GetAssetVersion returns the value of AssetVersion.
func (s *SubmissionCreate) GetAssetVersion() int64 {
	return s.AssetVersion
}

// GetTargetAssetID returns the value of TargetAssetID.
func (s *SubmissionCreate) GetTargetAssetID() OptInt64 {
	return s.TargetAssetID
}

// SetDisplayName sets the value of DisplayName.
func (s *SubmissionCreate) SetDisplayName(val string) {
	s.DisplayName = val
}

// SetCreator sets the value of Creator.
func (s *SubmissionCreate) SetCreator(val string) {
	s.Creator = val
}

// SetGameID sets the value of GameID.
func (s *SubmissionCreate) SetGameID(val int32) {
	s.GameID = val
}

// SetAssetID sets the value of AssetID.
func (s *SubmissionCreate) SetAssetID(val int64) {
	s.AssetID = val
}

// SetAssetVersion sets the value of AssetVersion.
func (s *SubmissionCreate) SetAssetVersion(val int64) {
	s.AssetVersion = val
}

// SetTargetAssetID sets the value of TargetAssetID.
func (s *SubmissionCreate) SetTargetAssetID(val OptInt64) {
	s.TargetAssetID = val
}

// Ref: #/components/schemas/SubmissionFilter
type SubmissionFilter struct {
	ID          int64     `json:"ID"`
	DisplayName OptString `json:"DisplayName"`
	Creator     OptString `json:"Creator"`
	GameID      OptInt32  `json:"GameID"`
	Date        OptInt64  `json:"Date"`
}

// GetID returns the value of ID.
func (s *SubmissionFilter) GetID() int64 {
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
func (s *SubmissionFilter) SetID(val int64) {
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

// UpdateScriptOK is response for UpdateScript operation.
type UpdateScriptOK struct{}

// UpdateScriptPolicyOK is response for UpdateScriptPolicy operation.
type UpdateScriptPolicyOK struct{}
