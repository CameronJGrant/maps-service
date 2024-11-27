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
		if s.ID.Set {
			e.FieldStart("ID")
			s.ID.Encode(e)
		}
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

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ID":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
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

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
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

// Encode encodes SubmissionCreate as json.
func (o OptSubmissionCreate) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes SubmissionCreate from json.
func (o *OptSubmissionCreate) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSubmissionCreate to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSubmissionCreate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSubmissionCreate) UnmarshalJSON(data []byte) error {
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
		if s.ID.Set {
			e.FieldStart("ID")
			s.ID.Encode(e)
		}
	}
	{
		if s.DisplayName.Set {
			e.FieldStart("DisplayName")
			s.DisplayName.Encode(e)
		}
	}
	{
		if s.Creator.Set {
			e.FieldStart("Creator")
			s.Creator.Encode(e)
		}
	}
	{
		if s.GameID.Set {
			e.FieldStart("GameID")
			s.GameID.Encode(e)
		}
	}
	{
		if s.Date.Set {
			e.FieldStart("Date")
			s.Date.Encode(e)
		}
	}
	{
		if s.Submitter.Set {
			e.FieldStart("Submitter")
			s.Submitter.Encode(e)
		}
	}
	{
		if s.AssetID.Set {
			e.FieldStart("AssetID")
			s.AssetID.Encode(e)
		}
	}
	{
		if s.AssetVersion.Set {
			e.FieldStart("AssetVersion")
			s.AssetVersion.Encode(e)
		}
	}
	{
		if s.Completed.Set {
			e.FieldStart("Completed")
			s.Completed.Encode(e)
		}
	}
	{
		if s.SubmissionType.Set {
			e.FieldStart("SubmissionType")
			s.SubmissionType.Encode(e)
		}
	}
	{
		if s.TargetAssetID.Set {
			e.FieldStart("TargetAssetID")
			s.TargetAssetID.Encode(e)
		}
	}
	{
		if s.StatusID.Set {
			e.FieldStart("StatusID")
			s.StatusID.Encode(e)
		}
	}
}

var jsonFieldsNameOfSubmission = [12]string{
	0:  "ID",
	1:  "DisplayName",
	2:  "Creator",
	3:  "GameID",
	4:  "Date",
	5:  "Submitter",
	6:  "AssetID",
	7:  "AssetVersion",
	8:  "Completed",
	9:  "SubmissionType",
	10: "TargetAssetID",
	11: "StatusID",
}

// Decode decodes Submission from json.
func (s *Submission) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Submission to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ID":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ID\"")
			}
		case "DisplayName":
			if err := func() error {
				s.DisplayName.Reset()
				if err := s.DisplayName.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"DisplayName\"")
			}
		case "Creator":
			if err := func() error {
				s.Creator.Reset()
				if err := s.Creator.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Creator\"")
			}
		case "GameID":
			if err := func() error {
				s.GameID.Reset()
				if err := s.GameID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"GameID\"")
			}
		case "Date":
			if err := func() error {
				s.Date.Reset()
				if err := s.Date.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Date\"")
			}
		case "Submitter":
			if err := func() error {
				s.Submitter.Reset()
				if err := s.Submitter.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Submitter\"")
			}
		case "AssetID":
			if err := func() error {
				s.AssetID.Reset()
				if err := s.AssetID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"AssetID\"")
			}
		case "AssetVersion":
			if err := func() error {
				s.AssetVersion.Reset()
				if err := s.AssetVersion.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"AssetVersion\"")
			}
		case "Completed":
			if err := func() error {
				s.Completed.Reset()
				if err := s.Completed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Completed\"")
			}
		case "SubmissionType":
			if err := func() error {
				s.SubmissionType.Reset()
				if err := s.SubmissionType.Decode(d); err != nil {
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
			if err := func() error {
				s.StatusID.Reset()
				if err := s.StatusID.Decode(d); err != nil {
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
		if s.DisplayName.Set {
			e.FieldStart("DisplayName")
			s.DisplayName.Encode(e)
		}
	}
	{
		if s.Creator.Set {
			e.FieldStart("Creator")
			s.Creator.Encode(e)
		}
	}
	{
		if s.GameID.Set {
			e.FieldStart("GameID")
			s.GameID.Encode(e)
		}
	}
	{
		if s.Submitter.Set {
			e.FieldStart("Submitter")
			s.Submitter.Encode(e)
		}
	}
	{
		if s.AssetID.Set {
			e.FieldStart("AssetID")
			s.AssetID.Encode(e)
		}
	}
	{
		if s.AssetVersion.Set {
			e.FieldStart("AssetVersion")
			s.AssetVersion.Encode(e)
		}
	}
	{
		if s.SubmissionType.Set {
			e.FieldStart("SubmissionType")
			s.SubmissionType.Encode(e)
		}
	}
	{
		if s.TargetAssetID.Set {
			e.FieldStart("TargetAssetID")
			s.TargetAssetID.Encode(e)
		}
	}
}

var jsonFieldsNameOfSubmissionCreate = [8]string{
	0: "DisplayName",
	1: "Creator",
	2: "GameID",
	3: "Submitter",
	4: "AssetID",
	5: "AssetVersion",
	6: "SubmissionType",
	7: "TargetAssetID",
}

// Decode decodes SubmissionCreate from json.
func (s *SubmissionCreate) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SubmissionCreate to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "DisplayName":
			if err := func() error {
				s.DisplayName.Reset()
				if err := s.DisplayName.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"DisplayName\"")
			}
		case "Creator":
			if err := func() error {
				s.Creator.Reset()
				if err := s.Creator.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Creator\"")
			}
		case "GameID":
			if err := func() error {
				s.GameID.Reset()
				if err := s.GameID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"GameID\"")
			}
		case "Submitter":
			if err := func() error {
				s.Submitter.Reset()
				if err := s.Submitter.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"Submitter\"")
			}
		case "AssetID":
			if err := func() error {
				s.AssetID.Reset()
				if err := s.AssetID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"AssetID\"")
			}
		case "AssetVersion":
			if err := func() error {
				s.AssetVersion.Reset()
				if err := s.AssetVersion.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"AssetVersion\"")
			}
		case "SubmissionType":
			if err := func() error {
				s.SubmissionType.Reset()
				if err := s.SubmissionType.Decode(d); err != nil {
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
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SubmissionCreate")
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
