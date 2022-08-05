package cogs

type Command interface {
	Invokes() []string
	Description() string
	ModRequired() bool
	Exec(ctx *Context) error
}
