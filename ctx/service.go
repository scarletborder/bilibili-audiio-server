package ctx

import (
	bilibili "github.com/CuteReimu/bilibili/v2"
)

type SrvCtx struct {
	Version string

	B23_client *bilibili.Client
}

func NewSrvCtx() *SrvCtx {
	b23_client := bilibili.NewAnonymousClient()

	return &SrvCtx{
		B23_client: b23_client,
	}
}
