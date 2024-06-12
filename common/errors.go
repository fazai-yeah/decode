package common

import "errors"

// 自定义错误
var ErrNotEnoughParts = errors.New("filename does not contain enough parts")
