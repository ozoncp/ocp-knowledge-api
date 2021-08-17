package api

import (
	"context"

	api "github.com/ozoncp/ocp-knowledge-api/pkg/ocp-knowledge-api"
	"github.com/rs/zerolog/log"
)

type KnowledgeApi struct {
	api.UnimplementedOcpKnowledgeApiServer
}

func NewKnowledgeApi() api.OcpKnowledgeApiServer {
	return &KnowledgeApi{}
}

func (r *KnowledgeApi) CreateKnowledgeV1(ctx context.Context, request *api.CreateKnowledgeV1Request) (*api.CreateKnowledgeV1Response, error) {
	log.Info().Msgf("Create Knowledge %v", request.GetKnowledge())
	return &api.CreateKnowledgeV1Response{}, nil
}

func (r *KnowledgeApi) DescribeKnowledgeV1(ctx context.Context, request *api.DescribeKnowledgeV1Request) (*api.DescribeKnowledgeV1Response, error) {
	log.Info().Msgf("Describe Knowledge ID - %v", request.GetId())
	return &api.DescribeKnowledgeV1Response{}, nil
}

func (r *KnowledgeApi) ListKnowledgeV1(ctx context.Context, request *api.ListKnowledgeV1Request) (*api.ListKnowledgeV1Response, error) {
	log.Info().Msgf("List Knowledge. Offset - %v; Limit - %v", request.GetOffset(), request.GetLimit())
	return &api.ListKnowledgeV1Response{}, nil
}

func (r *KnowledgeApi) RemoveKnowledgeV1(ctx context.Context, request *api.RemoveKnowledgeV1Request) (*api.RemoveKnowledgeV1Response, error) {
	log.Info().Msgf("Remove knowledge ID - %v", request.GetId())
	return &api.RemoveKnowledgeV1Response{}, nil
}
