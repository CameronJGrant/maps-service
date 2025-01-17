// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// ActionSubmissionPublishParams is parameters of actionSubmissionPublish operation.
type ActionSubmissionPublishParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackActionSubmissionPublishParams(packed middleware.Parameters) (params ActionSubmissionPublishParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeActionSubmissionPublishParams(args [1]string, argsEscaped bool, r *http.Request) (params ActionSubmissionPublishParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ActionSubmissionRejectParams is parameters of actionSubmissionReject operation.
type ActionSubmissionRejectParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackActionSubmissionRejectParams(packed middleware.Parameters) (params ActionSubmissionRejectParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeActionSubmissionRejectParams(args [1]string, argsEscaped bool, r *http.Request) (params ActionSubmissionRejectParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ActionSubmissionRequestChangesParams is parameters of actionSubmissionRequestChanges operation.
type ActionSubmissionRequestChangesParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackActionSubmissionRequestChangesParams(packed middleware.Parameters) (params ActionSubmissionRequestChangesParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeActionSubmissionRequestChangesParams(args [1]string, argsEscaped bool, r *http.Request) (params ActionSubmissionRequestChangesParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ActionSubmissionRevokeParams is parameters of actionSubmissionRevoke operation.
type ActionSubmissionRevokeParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackActionSubmissionRevokeParams(packed middleware.Parameters) (params ActionSubmissionRevokeParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeActionSubmissionRevokeParams(args [1]string, argsEscaped bool, r *http.Request) (params ActionSubmissionRevokeParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ActionSubmissionSubmitParams is parameters of actionSubmissionSubmit operation.
type ActionSubmissionSubmitParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackActionSubmissionSubmitParams(packed middleware.Parameters) (params ActionSubmissionSubmitParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeActionSubmissionSubmitParams(args [1]string, argsEscaped bool, r *http.Request) (params ActionSubmissionSubmitParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ActionSubmissionTriggerPublishParams is parameters of actionSubmissionTriggerPublish operation.
type ActionSubmissionTriggerPublishParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackActionSubmissionTriggerPublishParams(packed middleware.Parameters) (params ActionSubmissionTriggerPublishParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeActionSubmissionTriggerPublishParams(args [1]string, argsEscaped bool, r *http.Request) (params ActionSubmissionTriggerPublishParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ActionSubmissionTriggerValidateParams is parameters of actionSubmissionTriggerValidate operation.
type ActionSubmissionTriggerValidateParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackActionSubmissionTriggerValidateParams(packed middleware.Parameters) (params ActionSubmissionTriggerValidateParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeActionSubmissionTriggerValidateParams(args [1]string, argsEscaped bool, r *http.Request) (params ActionSubmissionTriggerValidateParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ActionSubmissionValidateParams is parameters of actionSubmissionValidate operation.
type ActionSubmissionValidateParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackActionSubmissionValidateParams(packed middleware.Parameters) (params ActionSubmissionValidateParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeActionSubmissionValidateParams(args [1]string, argsEscaped bool, r *http.Request) (params ActionSubmissionValidateParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteScriptParams is parameters of deleteScript operation.
type DeleteScriptParams struct {
	// The unique identifier for a script.
	ScriptID int64
}

func unpackDeleteScriptParams(packed middleware.Parameters) (params DeleteScriptParams) {
	{
		key := middleware.ParameterKey{
			Name: "ScriptID",
			In:   "path",
		}
		params.ScriptID = packed[key].(int64)
	}
	return params
}

func decodeDeleteScriptParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteScriptParams, _ error) {
	// Decode path: ScriptID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ScriptID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ScriptID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ScriptID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteScriptPolicyParams is parameters of deleteScriptPolicy operation.
type DeleteScriptPolicyParams struct {
	// The unique identifier for a script policy.
	ScriptPolicyID int64
}

func unpackDeleteScriptPolicyParams(packed middleware.Parameters) (params DeleteScriptPolicyParams) {
	{
		key := middleware.ParameterKey{
			Name: "ScriptPolicyID",
			In:   "path",
		}
		params.ScriptPolicyID = packed[key].(int64)
	}
	return params
}

func decodeDeleteScriptPolicyParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteScriptPolicyParams, _ error) {
	// Decode path: ScriptPolicyID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ScriptPolicyID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ScriptPolicyID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ScriptPolicyID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetScriptParams is parameters of getScript operation.
type GetScriptParams struct {
	// The unique identifier for a script.
	ScriptID int64
}

func unpackGetScriptParams(packed middleware.Parameters) (params GetScriptParams) {
	{
		key := middleware.ParameterKey{
			Name: "ScriptID",
			In:   "path",
		}
		params.ScriptID = packed[key].(int64)
	}
	return params
}

func decodeGetScriptParams(args [1]string, argsEscaped bool, r *http.Request) (params GetScriptParams, _ error) {
	// Decode path: ScriptID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ScriptID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ScriptID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ScriptID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetScriptPolicyParams is parameters of getScriptPolicy operation.
type GetScriptPolicyParams struct {
	// The unique identifier for a script policy.
	ScriptPolicyID int64
}

func unpackGetScriptPolicyParams(packed middleware.Parameters) (params GetScriptPolicyParams) {
	{
		key := middleware.ParameterKey{
			Name: "ScriptPolicyID",
			In:   "path",
		}
		params.ScriptPolicyID = packed[key].(int64)
	}
	return params
}

func decodeGetScriptPolicyParams(args [1]string, argsEscaped bool, r *http.Request) (params GetScriptPolicyParams, _ error) {
	// Decode path: ScriptPolicyID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ScriptPolicyID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ScriptPolicyID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ScriptPolicyID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetScriptPolicyFromHashParams is parameters of getScriptPolicyFromHash operation.
type GetScriptPolicyFromHashParams struct {
	FromScriptHash string
}

func unpackGetScriptPolicyFromHashParams(packed middleware.Parameters) (params GetScriptPolicyFromHashParams) {
	{
		key := middleware.ParameterKey{
			Name: "FromScriptHash",
			In:   "path",
		}
		params.FromScriptHash = packed[key].(string)
	}
	return params
}

func decodeGetScriptPolicyFromHashParams(args [1]string, argsEscaped bool, r *http.Request) (params GetScriptPolicyFromHashParams, _ error) {
	// Decode path: FromScriptHash.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "FromScriptHash",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FromScriptHash = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    16,
					MinLengthSet: true,
					MaxLength:    16,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.FromScriptHash)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "FromScriptHash",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetSubmissionParams is parameters of getSubmission operation.
type GetSubmissionParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackGetSubmissionParams(packed middleware.Parameters) (params GetSubmissionParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeGetSubmissionParams(args [1]string, argsEscaped bool, r *http.Request) (params GetSubmissionParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListSubmissionsParams is parameters of listSubmissions operation.
type ListSubmissionsParams struct {
	Page   Pagination
	Filter OptSubmissionFilter
}

func unpackListSubmissionsParams(packed middleware.Parameters) (params ListSubmissionsParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		params.Page = packed[key].(Pagination)
	}
	{
		key := middleware.ParameterKey{
			Name: "filter",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Filter = v.(OptSubmissionFilter)
		}
	}
	return params
}

func decodeListSubmissionsParams(args [0]string, argsEscaped bool, r *http.Request) (params ListSubmissionsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "Page", Required: true}, {Name: "Limit", Required: true}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return params.Page.DecodeURI(d)
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Page.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: filter.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "filter",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{Name: "ID", Required: true}, {Name: "DisplayName", Required: false}, {Name: "Creator", Required: false}, {Name: "GameID", Required: false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFilterVal SubmissionFilter
				if err := func() error {
					return paramsDotFilterVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Filter.SetTo(paramsDotFilterVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Filter.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "filter",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SetSubmissionCompletedParams is parameters of setSubmissionCompleted operation.
type SetSubmissionCompletedParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
}

func unpackSetSubmissionCompletedParams(packed middleware.Parameters) (params SetSubmissionCompletedParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	return params
}

func decodeSetSubmissionCompletedParams(args [1]string, argsEscaped bool, r *http.Request) (params SetSubmissionCompletedParams, _ error) {
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateScriptParams is parameters of updateScript operation.
type UpdateScriptParams struct {
	// The unique identifier for a script.
	ScriptID int64
}

func unpackUpdateScriptParams(packed middleware.Parameters) (params UpdateScriptParams) {
	{
		key := middleware.ParameterKey{
			Name: "ScriptID",
			In:   "path",
		}
		params.ScriptID = packed[key].(int64)
	}
	return params
}

func decodeUpdateScriptParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateScriptParams, _ error) {
	// Decode path: ScriptID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ScriptID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ScriptID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ScriptID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateScriptPolicyParams is parameters of updateScriptPolicy operation.
type UpdateScriptPolicyParams struct {
	// The unique identifier for a script policy.
	ScriptPolicyID int64
}

func unpackUpdateScriptPolicyParams(packed middleware.Parameters) (params UpdateScriptPolicyParams) {
	{
		key := middleware.ParameterKey{
			Name: "ScriptPolicyID",
			In:   "path",
		}
		params.ScriptPolicyID = packed[key].(int64)
	}
	return params
}

func decodeUpdateScriptPolicyParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateScriptPolicyParams, _ error) {
	// Decode path: ScriptPolicyID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ScriptPolicyID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ScriptPolicyID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ScriptPolicyID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateSubmissionModelParams is parameters of updateSubmissionModel operation.
type UpdateSubmissionModelParams struct {
	// The unique identifier for a submission.
	SubmissionID int64
	ModelID      int64
	VersionID    int64
}

func unpackUpdateSubmissionModelParams(packed middleware.Parameters) (params UpdateSubmissionModelParams) {
	{
		key := middleware.ParameterKey{
			Name: "SubmissionID",
			In:   "path",
		}
		params.SubmissionID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "ModelID",
			In:   "query",
		}
		params.ModelID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "VersionID",
			In:   "query",
		}
		params.VersionID = packed[key].(int64)
	}
	return params
}

func decodeUpdateSubmissionModelParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateSubmissionModelParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: SubmissionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "SubmissionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.SubmissionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "SubmissionID",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: ModelID.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ModelID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ModelID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ModelID",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: VersionID.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "VersionID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.VersionID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "VersionID",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
