package koop

type CallableEntity interface {
	GetSelfInvoke() bool
	GetCommand() string
	GetArgs() []string
	IsRaw() bool
}
