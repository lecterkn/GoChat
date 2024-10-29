package channel_permission

import (
	"fmt"
)

type ChannelPermission int

const (
	ReadOnly = iota
	Writable
	Private
)

type attr struct {
	V    ChannelPermission
	Code string
}

const size = 3

var channelPermissions = [size]attr{
	{ReadOnly, "readOnly"},
	{Writable, "writable"},
	{Private, "private"},
}

func (cp ChannelPermission) ToCode() string {
	for _, permission := range channelPermissions {
		if permission.V == cp {
			return permission.Code
		}
	}
	return ""
}

func GetChannelPermissionFromCode(code string) (ChannelPermission, error) {
	for _, permission := range channelPermissions {
		if permission.Code == code {
			return permission.V, nil
		}
	}
	return ReadOnly, fmt.Errorf("invalid channel permission code")
}
