package ctxprovider

import (
	"context"
	"errors"
)

type CtxField string

type ContextProvider interface {
	GetCtx() context.Context
	SetCtx(ctx context.Context) error
	ForkCtx() context.Context
}

type ContextProviderImpl struct {
	context context.Context
}

func NewContextProvider(parent context.Context) *ContextProviderImpl {
	return &ContextProviderImpl{
		context: parent,
	}
}

func (c *ContextProviderImpl) GetCtx() context.Context {
	return c.context
}

func (c *ContextProviderImpl) SetCtx(ctx context.Context) error {
	if ctx == nil {
		return errors.New("invalid: ctx can't be nil")
	}
	c.context = ctx
	return nil
}

func (c *ContextProviderImpl) ForkCtx() context.Context {
	var forked CtxField = "forked"

	return context.WithValue(c.context,
		forked, true,
	)
}

func (c *ContextProviderImpl) SetField(key CtxField, value interface{}) {
	c.context = context.WithValue(c.context, key, value)
}
