// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_ai.proto

package server

import (
	"context"

	"aicoreops_ai/aicoreops_ai"
	"aicoreops_ai/internal/logic"
	"aicoreops_ai/internal/svc"
)

type AicoreopsAiServer struct {
	svcCtx *svc.ServiceContext
	aicoreops_ai.UnimplementedAicoreopsAiServer
}

func NewAicoreopsAiServer(svcCtx *svc.ServiceContext) *AicoreopsAiServer {
	return &AicoreopsAiServer{
		svcCtx: svcCtx,
	}
}

func (s *AicoreopsAiServer) Ping(ctx context.Context, in *aicoreops_ai.Request) (*aicoreops_ai.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
