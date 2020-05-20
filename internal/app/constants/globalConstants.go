package consts

type CodeType int32

const (
	StatusInactive = "INACTIVE"
	StatusActive   = "ACTIVE"

	CodeTypeRegister      CodeType = 0
	CodeTypeLogin         CodeType = 1
	CodeTypeResetPassword CodeType = 2

	MaximumNumberOfGroups       int = 10
	MaximumNumberOfGroupMembers int = 500
)
