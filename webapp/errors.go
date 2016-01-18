// Copyright 2015 Eleme Inc. All rights reserved.

package webapp

import (
	"fmt"
	"net/http"
)

// WebError is errors for web operations.
type WebError struct {
	// HTTP status code
	Code int `json:"code"`
	// Message
	Msg string `json:"msg"`
}

// Errors
var (
	// Common
	ErrBadRequest = NewWebError(http.StatusBadRequest, "Bad request")
	ErrNotNull    = NewWebError(http.StatusBadRequest, "Null value")
	ErrPrimaryKey = NewWebError(http.StatusForbidden, "Primarykey voilated")
	ErrUnique     = NewWebError(http.StatusForbidden, "Value should be unique")
	ErrNotFound   = NewWebError(http.StatusNotFound, "Not found")
	// Project
	ErrProjectID            = NewWebError(http.StatusBadRequest, "Bad project id")
	ErrProjectName          = NewWebError(http.StatusBadRequest, "Bad project name")
	ErrProjectNotFound      = NewWebError(http.StatusNotFound, "Project not found")
	ErrDuplicateProjectName = NewWebError(http.StatusForbidden, "Duplicate project name")
	ErrDuplicateProjectUser = NewWebError(http.StatusForbidden, "Duplicate user to project")
	// User
	ErrUserID            = NewWebError(http.StatusBadRequest, "Bad user id")
	ErrUserName          = NewWebError(http.StatusBadRequest, "Bad user name")
	ErrUserEmail         = NewWebError(http.StatusBadRequest, "Bad user email")
	ErrUserPhone         = NewWebError(http.StatusBadRequest, "Bad user phone")
	ErrUserNotFound      = NewWebError(http.StatusNotFound, "User not found")
	ErrDuplicateUserName = NewWebError(http.StatusForbidden, "Duplicate user name")
	// Rule
	ErrRuleID                   = NewWebError(http.StatusBadRequest, "Bad rule id")
	ErrRulePattern              = NewWebError(http.StatusBadRequest, "Bad rule pattern")
	ErrRuleWhen                 = NewWebError(http.StatusBadRequest, "Bad rule condition")
	ErrDuplicateRulePattern     = NewWebError(http.StatusForbidden, "Duplicate rule pattern")
	ErrRuleNotFound             = NewWebError(http.StatusNotFound, "Rule not found")
	ErrRuleNoCondition          = NewWebError(http.StatusBadRequest, "No condition specified")
	ErrRuleThresholdMaxRequired = NewWebError(http.StatusBadRequest, "ThresholdMax is required")
	ErrRuleThresholdMinRequired = NewWebError(http.StatusBadRequest, "ThresholdMin is required")
	// Metric
	ErrMetricNotFound = NewWebError(http.StatusNotFound, "Metric not found")
)

// NewWebError creates a WebError.
func NewWebError(code int, text string) *WebError {
	return &WebError{code, text}
}

// Error returns the string format of the WebError.
func (err *WebError) Error() string {
	return fmt.Sprintf("[%d]: %s", err.Code, err.Msg)
}

// NewUnexceptedWebError returns an unexcepted WebError.
func NewUnexceptedWebError(err error) *WebError {
	return NewWebError(http.StatusInternalServerError, err.Error())
}
