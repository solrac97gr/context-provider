package ctxprovider

import "context"

type ContextProvider struct {
	context context.Context
}

func NewContextProvider(parent context.Context) *ContextProvider {
	return &ContextProvider{
		context: parent,
	}
}

func (c *ContextProvider) GetCtx() context.Context {
	return c.context
}

func (c *ContextProvider) SetCtx(ctx context.Context) error {
	c.context = ctx
	return nil
}

func (c *ContextProvider) ForkCtx() context.Context {
	return context.WithValue(c.context,
		"forked", true,
	)
}
