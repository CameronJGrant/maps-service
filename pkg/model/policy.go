package model

import "time"

type Policy int32

const (
	ScriptPolicyAllowed Policy = 0
	ScriptPolicyBlocked Policy = 1
	ScriptPolicyDelete  Policy = 2
	ScriptPolicyReplace Policy = 3
)

type ScriptPolicy struct {
	ID int64 `gorm:"primaryKey"`
	// Hash of the source code that leads to this policy.
	// If this is a replacement mapping, the original source may not be pointed to by any policy.
	// The original source should still exist in the scripts table, which can be located by the same hash.
	FromScriptHash uint64
	// The ID of the replacement source (ScriptPolicyReplace)
	// or verbatim source (ScriptPolicyAllowed)
	// or 0 (other)
	ToScriptID int64
	Policy     Policy
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
