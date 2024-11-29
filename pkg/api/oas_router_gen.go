// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/submissions"
			origElem := elem
			if l := len("/submissions"); len(elem) >= l && elem[0:l] == "/submissions" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch r.Method {
				case "GET":
					s.handleListSubmissionsRequest([0]string{}, elemIsEscaped, w, r)
				case "POST":
					s.handleCreateSubmissionRequest([0]string{}, elemIsEscaped, w, r)
				default:
					s.notAllowed(w, r, "GET,POST")
				}

				return
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				origElem := elem
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "SubmissionID"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleGetSubmissionRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'c': // Prefix: "completed"
						origElem := elem
						if l := len("completed"); len(elem) >= l && elem[0:l] == "completed" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PATCH":
								s.handlePatchSubmissionCompletedRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PATCH")
							}

							return
						}

						elem = origElem
					case 'm': // Prefix: "model"
						origElem := elem
						if l := len("model"); len(elem) >= l && elem[0:l] == "model" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PATCH":
								s.handlePatchSubmissionModelRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PATCH")
							}

							return
						}

						elem = origElem
					case 's': // Prefix: "status/"
						origElem := elem
						if l := len("status/"); len(elem) >= l && elem[0:l] == "status/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'p': // Prefix: "publish"
							origElem := elem
							if l := len("publish"); len(elem) >= l && elem[0:l] == "publish" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "PATCH":
									s.handleActionSubmissionPublishRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PATCH")
								}

								return
							}

							elem = origElem
						case 'r': // Prefix: "re"
							origElem := elem
							if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'j': // Prefix: "ject"
								origElem := elem
								if l := len("ject"); len(elem) >= l && elem[0:l] == "ject" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "PATCH":
										s.handleActionSubmissionRejectRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "PATCH")
									}

									return
								}

								elem = origElem
							case 'q': // Prefix: "quest-changes"
								origElem := elem
								if l := len("quest-changes"); len(elem) >= l && elem[0:l] == "quest-changes" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "PATCH":
										s.handleActionSubmissionRequestChangesRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "PATCH")
									}

									return
								}

								elem = origElem
							case 'v': // Prefix: "voke"
								origElem := elem
								if l := len("voke"); len(elem) >= l && elem[0:l] == "voke" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "PATCH":
										s.handleActionSubmissionRevokeRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "PATCH")
									}

									return
								}

								elem = origElem
							}

							elem = origElem
						case 's': // Prefix: "submit"
							origElem := elem
							if l := len("submit"); len(elem) >= l && elem[0:l] == "submit" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "PATCH":
									s.handleActionSubmissionSubmitRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PATCH")
								}

								return
							}

							elem = origElem
						case 't': // Prefix: "trigger-"
							origElem := elem
							if l := len("trigger-"); len(elem) >= l && elem[0:l] == "trigger-" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'p': // Prefix: "publish"
								origElem := elem
								if l := len("publish"); len(elem) >= l && elem[0:l] == "publish" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "PATCH":
										s.handleActionSubmissionTriggerPublishRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "PATCH")
									}

									return
								}

								elem = origElem
							case 'v': // Prefix: "validate"
								origElem := elem
								if l := len("validate"); len(elem) >= l && elem[0:l] == "validate" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "PATCH":
										s.handleActionSubmissionTriggerValidateRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "PATCH")
									}

									return
								}

								elem = origElem
							}

							elem = origElem
						case 'v': // Prefix: "validate"
							origElem := elem
							if l := len("validate"); len(elem) >= l && elem[0:l] == "validate" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "PATCH":
									s.handleActionSubmissionValidateRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PATCH")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/submissions"
			origElem := elem
			if l := len("/submissions"); len(elem) >= l && elem[0:l] == "/submissions" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch method {
				case "GET":
					r.name = ListSubmissionsOperation
					r.summary = "Get list of submissions"
					r.operationID = "listSubmissions"
					r.pathPattern = "/submissions"
					r.args = args
					r.count = 0
					return r, true
				case "POST":
					r.name = CreateSubmissionOperation
					r.summary = "Create new submission"
					r.operationID = "createSubmission"
					r.pathPattern = "/submissions"
					r.args = args
					r.count = 0
					return r, true
				default:
					return
				}
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				origElem := elem
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "SubmissionID"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = GetSubmissionOperation
						r.summary = "Retrieve map with ID"
						r.operationID = "getSubmission"
						r.pathPattern = "/submissions/{SubmissionID}"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'c': // Prefix: "completed"
						origElem := elem
						if l := len("completed"); len(elem) >= l && elem[0:l] == "completed" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "PATCH":
								r.name = PatchSubmissionCompletedOperation
								r.summary = "Retrieve map with ID"
								r.operationID = "patchSubmissionCompleted"
								r.pathPattern = "/submissions/{SubmissionID}/completed"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'm': // Prefix: "model"
						origElem := elem
						if l := len("model"); len(elem) >= l && elem[0:l] == "model" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "PATCH":
								r.name = PatchSubmissionModelOperation
								r.summary = "Update model following role restrictions"
								r.operationID = "patchSubmissionModel"
								r.pathPattern = "/submissions/{SubmissionID}/model"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 's': // Prefix: "status/"
						origElem := elem
						if l := len("status/"); len(elem) >= l && elem[0:l] == "status/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'p': // Prefix: "publish"
							origElem := elem
							if l := len("publish"); len(elem) >= l && elem[0:l] == "publish" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "PATCH":
									r.name = ActionSubmissionPublishOperation
									r.summary = "Role Validator changes status from Publishing -> Published"
									r.operationID = "actionSubmissionPublish"
									r.pathPattern = "/submissions/{SubmissionID}/status/publish"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'r': // Prefix: "re"
							origElem := elem
							if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'j': // Prefix: "ject"
								origElem := elem
								if l := len("ject"); len(elem) >= l && elem[0:l] == "ject" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "PATCH":
										r.name = ActionSubmissionRejectOperation
										r.summary = "Role Reviewer changes status from Submitted -> Rejected"
										r.operationID = "actionSubmissionReject"
										r.pathPattern = "/submissions/{SubmissionID}/status/reject"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'q': // Prefix: "quest-changes"
								origElem := elem
								if l := len("quest-changes"); len(elem) >= l && elem[0:l] == "quest-changes" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "PATCH":
										r.name = ActionSubmissionRequestChangesOperation
										r.summary = "Role Reviewer changes status from Validated|Accepted|Submitted -> ChangesRequested"
										r.operationID = "actionSubmissionRequestChanges"
										r.pathPattern = "/submissions/{SubmissionID}/status/request-changes"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'v': // Prefix: "voke"
								origElem := elem
								if l := len("voke"); len(elem) >= l && elem[0:l] == "voke" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "PATCH":
										r.name = ActionSubmissionRevokeOperation
										r.summary = "Role Submitter changes status from Submitted|ChangesRequested -> UnderConstruction"
										r.operationID = "actionSubmissionRevoke"
										r.pathPattern = "/submissions/{SubmissionID}/status/revoke"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}

							elem = origElem
						case 's': // Prefix: "submit"
							origElem := elem
							if l := len("submit"); len(elem) >= l && elem[0:l] == "submit" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "PATCH":
									r.name = ActionSubmissionSubmitOperation
									r.summary = "Role Submitter changes status from UnderConstruction|ChangesRequested -> Submitted"
									r.operationID = "actionSubmissionSubmit"
									r.pathPattern = "/submissions/{SubmissionID}/status/submit"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 't': // Prefix: "trigger-"
							origElem := elem
							if l := len("trigger-"); len(elem) >= l && elem[0:l] == "trigger-" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'p': // Prefix: "publish"
								origElem := elem
								if l := len("publish"); len(elem) >= l && elem[0:l] == "publish" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "PATCH":
										r.name = ActionSubmissionTriggerPublishOperation
										r.summary = "Role Admin changes status from Validated -> Publishing"
										r.operationID = "actionSubmissionTriggerPublish"
										r.pathPattern = "/submissions/{SubmissionID}/status/trigger-publish"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'v': // Prefix: "validate"
								origElem := elem
								if l := len("validate"); len(elem) >= l && elem[0:l] == "validate" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "PATCH":
										r.name = ActionSubmissionTriggerValidateOperation
										r.summary = "Role Reviewer triggers validation and changes status from Submitted|Accepted -> Validating"
										r.operationID = "actionSubmissionTriggerValidate"
										r.pathPattern = "/submissions/{SubmissionID}/status/trigger-validate"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}

							elem = origElem
						case 'v': // Prefix: "validate"
							origElem := elem
							if l := len("validate"); len(elem) >= l && elem[0:l] == "validate" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "PATCH":
									r.name = ActionSubmissionValidateOperation
									r.summary = "Role Validator changes status from Validating -> Validated"
									r.operationID = "actionSubmissionValidate"
									r.pathPattern = "/submissions/{SubmissionID}/status/validate"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
