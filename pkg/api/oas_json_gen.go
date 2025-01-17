// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("code")
		e.Int64(s.Code)
	}
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
}

var jsonFieldsNameOfError = [2]string{
	0: "code",
	1: "message",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.Code = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfError) {
					name = jsonFieldsNameOfError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ID) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ID) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("ID")
		e.Int64(s.ID)
	}
}

var jsonFieldsNameOfID = [1]string{
	0: "ID",
}

// Decode decodes ID from json.
func (s *ID) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ID to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ID":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ID\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ID")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfID) {
					name = jsonFieldsNameOfID[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int32 as json.
func (o OptInt32) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int32(int32(o.Value))
}

// Decode decodes int32 from json.
func (o *OptInt32) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt32 to nil")
	}
	o.Set = true
	v, err := d.Int32()
	if err != nil {
		return err
	}
	o.Value = int32(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt32) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt32) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int64 as json.
func (o OptInt64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int64(int64(o.Value))
}

// Decode decodes int64 from json.
func (o *OptInt64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt64 to nil")
	}
	o.Set = true
	v, err := d.Int64()
	if err != nil {
		return err
	}
	o.Value = int64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Script) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Script) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("ID")
		e.Int64(s.ID)
	}
	{
		e.FieldStart("Hash")
		e.Str(s.Hash)
	}
	{
		e.FieldStart("Source")
		e.Str(s.Source)
	}
	{
		e.FieldStart("SubmissionID")
		e.Int64(s.SubmissionID)
	}
}

var jsonFieldsNameOfScript = [4]string{
	0: "ID",
	1: "Hash",
	2: "Source",
	3: "SubmissionID",
}

// Decode decodes Script from json.
func (s *Script) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Script to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ID":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ID\"")
			}
		case "Hash":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Hash = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Hash\"")
			}
		case "Source":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Source = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Source\"")
			}
		case "SubmissionID":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int64()
				s.SubmissionID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"SubmissionID\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Script")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfScript) {
					name = jsonFieldsNameOfScript[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Script) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Script) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ScriptCreate) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ScriptCreate) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("Source")
		e.Str(s.Source)
	}
	{
		if s.SubmissionID.Set {
			e.FieldStart("SubmissionID")
			s.SubmissionID.Encode(e)
		}
	}
}

var jsonFieldsNameOfScriptCreate = [2]string{
	0: "Source",
	1: "SubmissionID",
}

// Decode decodes ScriptCreate from json.
func (s *ScriptCreate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ScriptCreate to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "Source":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Source = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Source\"")
			}
		case "SubmissionID":
			if err := func() error {
				s.SubmissionID.Reset()
				if err := s.SubmissionID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"SubmissionID\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ScriptCreate")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfScriptCreate) {
					name = jsonFieldsNameOfScriptCreate[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ScriptCreate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ScriptCreate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ScriptPolicy) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ScriptPolicy) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("ID")
		e.Int64(s.ID)
	}
	{
		e.FieldStart("FromScriptHash")
		e.Str(s.FromScriptHash)
	}
	{
		e.FieldStart("ToScriptID")
		e.Int64(s.ToScriptID)
	}
	{
		e.FieldStart("Policy")
		e.Int32(s.Policy)
	}
}

var jsonFieldsNameOfScriptPolicy = [4]string{
	0: "ID",
	1: "FromScriptHash",
	2: "ToScriptID",
	3: "Policy",
}

// Decode decodes ScriptPolicy from json.
func (s *ScriptPolicy) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ScriptPolicy to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ID":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ID\"")
			}
		case "FromScriptHash":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.FromScriptHash = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"FromScriptHash\"")
			}
		case "ToScriptID":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int64()
				s.ToScriptID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ToScriptID\"")
			}
		case "Policy":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int32()
				s.Policy = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Policy\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ScriptPolicy")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfScriptPolicy) {
					name = jsonFieldsNameOfScriptPolicy[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ScriptPolicy) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ScriptPolicy) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ScriptPolicyCreate) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ScriptPolicyCreate) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("FromScriptID")
		e.Int64(s.FromScriptID)
	}
	{
		e.FieldStart("ToScriptID")
		e.Int64(s.ToScriptID)
	}
	{
		e.FieldStart("Policy")
		e.Int32(s.Policy)
	}
}

var jsonFieldsNameOfScriptPolicyCreate = [3]string{
	0: "FromScriptID",
	1: "ToScriptID",
	2: "Policy",
}

// Decode decodes ScriptPolicyCreate from json.
func (s *ScriptPolicyCreate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ScriptPolicyCreate to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "FromScriptID":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.FromScriptID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"FromScriptID\"")
			}
		case "ToScriptID":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int64()
				s.ToScriptID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ToScriptID\"")
			}
		case "Policy":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int32()
				s.Policy = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Policy\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ScriptPolicyCreate")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfScriptPolicyCreate) {
					name = jsonFieldsNameOfScriptPolicyCreate[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ScriptPolicyCreate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ScriptPolicyCreate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ScriptPolicyUpdate) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ScriptPolicyUpdate) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("ID")
		e.Int64(s.ID)
	}
	{
		if s.FromScriptID.Set {
			e.FieldStart("FromScriptID")
			s.FromScriptID.Encode(e)
		}
	}
	{
		if s.ToScriptID.Set {
			e.FieldStart("ToScriptID")
			s.ToScriptID.Encode(e)
		}
	}
	{
		if s.Policy.Set {
			e.FieldStart("Policy")
			s.Policy.Encode(e)
		}
	}
}

var jsonFieldsNameOfScriptPolicyUpdate = [4]string{
	0: "ID",
	1: "FromScriptID",
	2: "ToScriptID",
	3: "Policy",
}

// Decode decodes ScriptPolicyUpdate from json.
func (s *ScriptPolicyUpdate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ScriptPolicyUpdate to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ID":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ID\"")
			}
		case "FromScriptID":
			if err := func() error {
				s.FromScriptID.Reset()
				if err := s.FromScriptID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"FromScriptID\"")
			}
		case "ToScriptID":
			if err := func() error {
				s.ToScriptID.Reset()
				if err := s.ToScriptID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ToScriptID\"")
			}
		case "Policy":
			if err := func() error {
				s.Policy.Reset()
				if err := s.Policy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Policy\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ScriptPolicyUpdate")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfScriptPolicyUpdate) {
					name = jsonFieldsNameOfScriptPolicyUpdate[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ScriptPolicyUpdate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ScriptPolicyUpdate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ScriptUpdate) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ScriptUpdate) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("ID")
		e.Int64(s.ID)
	}
	{
		if s.Source.Set {
			e.FieldStart("Source")
			s.Source.Encode(e)
		}
	}
	{
		if s.SubmissionID.Set {
			e.FieldStart("SubmissionID")
			s.SubmissionID.Encode(e)
		}
	}
}

var jsonFieldsNameOfScriptUpdate = [3]string{
	0: "ID",
	1: "Source",
	2: "SubmissionID",
}

// Decode decodes ScriptUpdate from json.
func (s *ScriptUpdate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ScriptUpdate to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ID":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ID\"")
			}
		case "Source":
			if err := func() error {
				s.Source.Reset()
				if err := s.Source.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Source\"")
			}
		case "SubmissionID":
			if err := func() error {
				s.SubmissionID.Reset()
				if err := s.SubmissionID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"SubmissionID\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ScriptUpdate")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfScriptUpdate) {
					name = jsonFieldsNameOfScriptUpdate[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ScriptUpdate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ScriptUpdate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Submission) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Submission) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("ID")
		e.Int64(s.ID)
	}
	{
		e.FieldStart("DisplayName")
		e.Str(s.DisplayName)
	}
	{
		e.FieldStart("Creator")
		e.Str(s.Creator)
	}
	{
		e.FieldStart("GameID")
		e.Int32(s.GameID)
	}
	{
		e.FieldStart("CreatedAt")
		e.Int64(s.CreatedAt)
	}
	{
		e.FieldStart("UpdatedAt")
		e.Int64(s.UpdatedAt)
	}
	{
		e.FieldStart("Submitter")
		e.Int64(s.Submitter)
	}
	{
		e.FieldStart("AssetID")
		e.Int64(s.AssetID)
	}
	{
		e.FieldStart("AssetVersion")
		e.Int64(s.AssetVersion)
	}
	{
		e.FieldStart("Completed")
		e.Bool(s.Completed)
	}
	{
		e.FieldStart("SubmissionType")
		e.Int32(s.SubmissionType)
	}
	{
		if s.TargetAssetID.Set {
			e.FieldStart("TargetAssetID")
			s.TargetAssetID.Encode(e)
		}
	}
	{
		e.FieldStart("StatusID")
		e.Int32(s.StatusID)
	}
}

var jsonFieldsNameOfSubmission = [13]string{
	0:  "ID",
	1:  "DisplayName",
	2:  "Creator",
	3:  "GameID",
	4:  "CreatedAt",
	5:  "UpdatedAt",
	6:  "Submitter",
	7:  "AssetID",
	8:  "AssetVersion",
	9:  "Completed",
	10: "SubmissionType",
	11: "TargetAssetID",
	12: "StatusID",
}

// Decode decodes Submission from json.
func (s *Submission) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Submission to nil")
	}
	var requiredBitSet [2]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ID":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ID\"")
			}
		case "DisplayName":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.DisplayName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"DisplayName\"")
			}
		case "Creator":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Creator = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Creator\"")
			}
		case "GameID":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int32()
				s.GameID = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"GameID\"")
			}
		case "CreatedAt":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Int64()
				s.CreatedAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"CreatedAt\"")
			}
		case "UpdatedAt":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := d.Int64()
				s.UpdatedAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"UpdatedAt\"")
			}
		case "Submitter":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				v, err := d.Int64()
				s.Submitter = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Submitter\"")
			}
		case "AssetID":
			requiredBitSet[0] |= 1 << 7
			if err := func() error {
				v, err := d.Int64()
				s.AssetID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"AssetID\"")
			}
		case "AssetVersion":
			requiredBitSet[1] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.AssetVersion = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"AssetVersion\"")
			}
		case "Completed":
			requiredBitSet[1] |= 1 << 1
			if err := func() error {
				v, err := d.Bool()
				s.Completed = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Completed\"")
			}
		case "SubmissionType":
			requiredBitSet[1] |= 1 << 2
			if err := func() error {
				v, err := d.Int32()
				s.SubmissionType = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"SubmissionType\"")
			}
		case "TargetAssetID":
			if err := func() error {
				s.TargetAssetID.Reset()
				if err := s.TargetAssetID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"TargetAssetID\"")
			}
		case "StatusID":
			requiredBitSet[1] |= 1 << 4
			if err := func() error {
				v, err := d.Int32()
				s.StatusID = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"StatusID\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Submission")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [2]uint8{
		0b11111111,
		0b00010111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSubmission) {
					name = jsonFieldsNameOfSubmission[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Submission) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Submission) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SubmissionCreate) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SubmissionCreate) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("DisplayName")
		e.Str(s.DisplayName)
	}
	{
		e.FieldStart("Creator")
		e.Str(s.Creator)
	}
	{
		e.FieldStart("GameID")
		e.Int32(s.GameID)
	}
	{
		e.FieldStart("AssetID")
		e.Int64(s.AssetID)
	}
	{
		e.FieldStart("AssetVersion")
		e.Int64(s.AssetVersion)
	}
	{
		if s.TargetAssetID.Set {
			e.FieldStart("TargetAssetID")
			s.TargetAssetID.Encode(e)
		}
	}
}

var jsonFieldsNameOfSubmissionCreate = [6]string{
	0: "DisplayName",
	1: "Creator",
	2: "GameID",
	3: "AssetID",
	4: "AssetVersion",
	5: "TargetAssetID",
}

// Decode decodes SubmissionCreate from json.
func (s *SubmissionCreate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SubmissionCreate to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "DisplayName":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.DisplayName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"DisplayName\"")
			}
		case "Creator":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Creator = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Creator\"")
			}
		case "GameID":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int32()
				s.GameID = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"GameID\"")
			}
		case "AssetID":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int64()
				s.AssetID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"AssetID\"")
			}
		case "AssetVersion":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Int64()
				s.AssetVersion = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"AssetVersion\"")
			}
		case "TargetAssetID":
			if err := func() error {
				s.TargetAssetID.Reset()
				if err := s.TargetAssetID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"TargetAssetID\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SubmissionCreate")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSubmissionCreate) {
					name = jsonFieldsNameOfSubmissionCreate[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SubmissionCreate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SubmissionCreate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
