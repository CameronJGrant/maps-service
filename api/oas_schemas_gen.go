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

// Ref: #/components/schemas/Map
type Map struct {
	ID          OptInt64  `json:"ID"`
	DisplayName OptString `json:"DisplayName"`
	Creator     OptString `json:"Creator"`
	GameID      OptInt32  `json:"GameID"`
	Date        OptInt64  `json:"Date"`
}

// GetID returns the value of ID.
func (s *Map) GetID() OptInt64 {
	return s.ID
}

// GetDisplayName returns the value of DisplayName.
func (s *Map) GetDisplayName() OptString {
	return s.DisplayName
}

// GetCreator returns the value of Creator.
func (s *Map) GetCreator() OptString {
	return s.Creator
}

// GetGameID returns the value of GameID.
func (s *Map) GetGameID() OptInt32 {
	return s.GameID
}

// GetDate returns the value of Date.
func (s *Map) GetDate() OptInt64 {
	return s.Date
}

// SetID sets the value of ID.
func (s *Map) SetID(val OptInt64) {
	s.ID = val
}

// SetDisplayName sets the value of DisplayName.
func (s *Map) SetDisplayName(val OptString) {
	s.DisplayName = val
}

// SetCreator sets the value of Creator.
func (s *Map) SetCreator(val OptString) {
	s.Creator = val
}

// SetGameID sets the value of GameID.
func (s *Map) SetGameID(val OptInt32) {
	s.GameID = val
}

// SetDate sets the value of Date.
func (s *Map) SetDate(val OptInt64) {
	s.Date = val
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
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

// NewOptMap returns new OptMap with value set to v.
func NewOptMap(v Map) OptMap {
	return OptMap{
		Value: v,
		Set:   true,
	}
}

// OptMap is optional Map.
type OptMap struct {
	Value Map
	Set   bool
}

// IsSet returns true if OptMap was set.
func (o OptMap) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMap) Reset() {
	var v Map
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMap) SetTo(v Map) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMap) Get() (v Map, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMap) Or(d Map) Map {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRankFilter returns new OptRankFilter with value set to v.
func NewOptRankFilter(v RankFilter) OptRankFilter {
	return OptRankFilter{
		Value: v,
		Set:   true,
	}
}

// OptRankFilter is optional RankFilter.
type OptRankFilter struct {
	Value RankFilter
	Set   bool
}

// IsSet returns true if OptRankFilter was set.
func (o OptRankFilter) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRankFilter) Reset() {
	var v RankFilter
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRankFilter) SetTo(v RankFilter) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRankFilter) Get() (v RankFilter, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRankFilter) Or(d RankFilter) RankFilter {
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

// NewOptTimeFilter returns new OptTimeFilter with value set to v.
func NewOptTimeFilter(v TimeFilter) OptTimeFilter {
	return OptTimeFilter{
		Value: v,
		Set:   true,
	}
}

// OptTimeFilter is optional TimeFilter.
type OptTimeFilter struct {
	Value TimeFilter
	Set   bool
}

// IsSet returns true if OptTimeFilter was set.
func (o OptTimeFilter) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTimeFilter) Reset() {
	var v TimeFilter
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTimeFilter) SetTo(v TimeFilter) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTimeFilter) Get() (v TimeFilter, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTimeFilter) Or(d TimeFilter) TimeFilter {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUser returns new OptUser with value set to v.
func NewOptUser(v User) OptUser {
	return OptUser{
		Value: v,
		Set:   true,
	}
}

// OptUser is optional User.
type OptUser struct {
	Value User
	Set   bool
}

// IsSet returns true if OptUser was set.
func (o OptUser) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUser) Reset() {
	var v User
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUser) SetTo(v User) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUser) Get() (v User, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUser) Or(d User) User {
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

// Ref: #/components/schemas/Rank
type Rank struct {
	ID        OptInt64   `json:"ID"`
	User      OptUser    `json:"User"`
	StyleID   OptInt32   `json:"StyleID"`
	ModeID    OptInt32   `json:"ModeID"`
	GameID    OptInt32   `json:"GameID"`
	Rank      OptFloat64 `json:"Rank"`
	Skill     OptFloat64 `json:"Skill"`
	UpdatedAt OptInt64   `json:"UpdatedAt"`
}

// GetID returns the value of ID.
func (s *Rank) GetID() OptInt64 {
	return s.ID
}

// GetUser returns the value of User.
func (s *Rank) GetUser() OptUser {
	return s.User
}

// GetStyleID returns the value of StyleID.
func (s *Rank) GetStyleID() OptInt32 {
	return s.StyleID
}

// GetModeID returns the value of ModeID.
func (s *Rank) GetModeID() OptInt32 {
	return s.ModeID
}

// GetGameID returns the value of GameID.
func (s *Rank) GetGameID() OptInt32 {
	return s.GameID
}

// GetRank returns the value of Rank.
func (s *Rank) GetRank() OptFloat64 {
	return s.Rank
}

// GetSkill returns the value of Skill.
func (s *Rank) GetSkill() OptFloat64 {
	return s.Skill
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Rank) GetUpdatedAt() OptInt64 {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Rank) SetID(val OptInt64) {
	s.ID = val
}

// SetUser sets the value of User.
func (s *Rank) SetUser(val OptUser) {
	s.User = val
}

// SetStyleID sets the value of StyleID.
func (s *Rank) SetStyleID(val OptInt32) {
	s.StyleID = val
}

// SetModeID sets the value of ModeID.
func (s *Rank) SetModeID(val OptInt32) {
	s.ModeID = val
}

// SetGameID sets the value of GameID.
func (s *Rank) SetGameID(val OptInt32) {
	s.GameID = val
}

// SetRank sets the value of Rank.
func (s *Rank) SetRank(val OptFloat64) {
	s.Rank = val
}

// SetSkill sets the value of Skill.
func (s *Rank) SetSkill(val OptFloat64) {
	s.Skill = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Rank) SetUpdatedAt(val OptInt64) {
	s.UpdatedAt = val
}

// Ref: #/components/schemas/RankFilter
type RankFilter struct {
	StyleID OptInt32 `json:"StyleID"`
	GameID  OptInt32 `json:"GameID"`
	ModeID  OptInt32 `json:"ModeID"`
	Sort    OptInt64 `json:"Sort"`
}

// GetStyleID returns the value of StyleID.
func (s *RankFilter) GetStyleID() OptInt32 {
	return s.StyleID
}

// GetGameID returns the value of GameID.
func (s *RankFilter) GetGameID() OptInt32 {
	return s.GameID
}

// GetModeID returns the value of ModeID.
func (s *RankFilter) GetModeID() OptInt32 {
	return s.ModeID
}

// GetSort returns the value of Sort.
func (s *RankFilter) GetSort() OptInt64 {
	return s.Sort
}

// SetStyleID sets the value of StyleID.
func (s *RankFilter) SetStyleID(val OptInt32) {
	s.StyleID = val
}

// SetGameID sets the value of GameID.
func (s *RankFilter) SetGameID(val OptInt32) {
	s.GameID = val
}

// SetModeID sets the value of ModeID.
func (s *RankFilter) SetModeID(val OptInt32) {
	s.ModeID = val
}

// SetSort sets the value of Sort.
func (s *RankFilter) SetSort(val OptInt64) {
	s.Sort = val
}

// Ref: #/components/schemas/Time
type Time struct {
	ID      OptInt64 `json:"ID"`
	Time    OptInt64 `json:"Time"`
	User    OptUser  `json:"User"`
	Map     OptMap   `json:"Map"`
	Date    OptInt64 `json:"Date"`
	StyleID OptInt32 `json:"StyleID"`
	ModeID  OptInt32 `json:"ModeID"`
	GameID  OptInt32 `json:"GameID"`
}

// GetID returns the value of ID.
func (s *Time) GetID() OptInt64 {
	return s.ID
}

// GetTime returns the value of Time.
func (s *Time) GetTime() OptInt64 {
	return s.Time
}

// GetUser returns the value of User.
func (s *Time) GetUser() OptUser {
	return s.User
}

// GetMap returns the value of Map.
func (s *Time) GetMap() OptMap {
	return s.Map
}

// GetDate returns the value of Date.
func (s *Time) GetDate() OptInt64 {
	return s.Date
}

// GetStyleID returns the value of StyleID.
func (s *Time) GetStyleID() OptInt32 {
	return s.StyleID
}

// GetModeID returns the value of ModeID.
func (s *Time) GetModeID() OptInt32 {
	return s.ModeID
}

// GetGameID returns the value of GameID.
func (s *Time) GetGameID() OptInt32 {
	return s.GameID
}

// SetID sets the value of ID.
func (s *Time) SetID(val OptInt64) {
	s.ID = val
}

// SetTime sets the value of Time.
func (s *Time) SetTime(val OptInt64) {
	s.Time = val
}

// SetUser sets the value of User.
func (s *Time) SetUser(val OptUser) {
	s.User = val
}

// SetMap sets the value of Map.
func (s *Time) SetMap(val OptMap) {
	s.Map = val
}

// SetDate sets the value of Date.
func (s *Time) SetDate(val OptInt64) {
	s.Date = val
}

// SetStyleID sets the value of StyleID.
func (s *Time) SetStyleID(val OptInt32) {
	s.StyleID = val
}

// SetModeID sets the value of ModeID.
func (s *Time) SetModeID(val OptInt32) {
	s.ModeID = val
}

// SetGameID sets the value of GameID.
func (s *Time) SetGameID(val OptInt32) {
	s.GameID = val
}

// Ref: #/components/schemas/TimeFilter
type TimeFilter struct {
	ID      OptInt64 `json:"ID"`
	Time    OptInt64 `json:"Time"`
	UserID  OptInt64 `json:"UserID"`
	MapID   OptInt64 `json:"MapID"`
	StyleID OptInt32 `json:"StyleID"`
	ModeID  OptInt32 `json:"ModeID"`
	GameID  OptInt32 `json:"GameID"`
}

// GetID returns the value of ID.
func (s *TimeFilter) GetID() OptInt64 {
	return s.ID
}

// GetTime returns the value of Time.
func (s *TimeFilter) GetTime() OptInt64 {
	return s.Time
}

// GetUserID returns the value of UserID.
func (s *TimeFilter) GetUserID() OptInt64 {
	return s.UserID
}

// GetMapID returns the value of MapID.
func (s *TimeFilter) GetMapID() OptInt64 {
	return s.MapID
}

// GetStyleID returns the value of StyleID.
func (s *TimeFilter) GetStyleID() OptInt32 {
	return s.StyleID
}

// GetModeID returns the value of ModeID.
func (s *TimeFilter) GetModeID() OptInt32 {
	return s.ModeID
}

// GetGameID returns the value of GameID.
func (s *TimeFilter) GetGameID() OptInt32 {
	return s.GameID
}

// SetID sets the value of ID.
func (s *TimeFilter) SetID(val OptInt64) {
	s.ID = val
}

// SetTime sets the value of Time.
func (s *TimeFilter) SetTime(val OptInt64) {
	s.Time = val
}

// SetUserID sets the value of UserID.
func (s *TimeFilter) SetUserID(val OptInt64) {
	s.UserID = val
}

// SetMapID sets the value of MapID.
func (s *TimeFilter) SetMapID(val OptInt64) {
	s.MapID = val
}

// SetStyleID sets the value of StyleID.
func (s *TimeFilter) SetStyleID(val OptInt32) {
	s.StyleID = val
}

// SetModeID sets the value of ModeID.
func (s *TimeFilter) SetModeID(val OptInt32) {
	s.ModeID = val
}

// SetGameID sets the value of GameID.
func (s *TimeFilter) SetGameID(val OptInt32) {
	s.GameID = val
}

// Ref: #/components/schemas/User
type User struct {
	ID       OptInt64  `json:"ID"`
	Username OptString `json:"Username"`
	StateID  OptInt32  `json:"StateID"`
}

// GetID returns the value of ID.
func (s *User) GetID() OptInt64 {
	return s.ID
}

// GetUsername returns the value of Username.
func (s *User) GetUsername() OptString {
	return s.Username
}

// GetStateID returns the value of StateID.
func (s *User) GetStateID() OptInt32 {
	return s.StateID
}

// SetID sets the value of ID.
func (s *User) SetID(val OptInt64) {
	s.ID = val
}

// SetUsername sets the value of Username.
func (s *User) SetUsername(val OptString) {
	s.Username = val
}

// SetStateID sets the value of StateID.
func (s *User) SetStateID(val OptInt32) {
	s.StateID = val
}
