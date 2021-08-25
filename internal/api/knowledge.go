package api

import (
	"context"

	"github.com/ozoncp/ocp-knowledge-api/internal/models"
	"github.com/ozoncp/ocp-knowledge-api/internal/repo"
	api "github.com/ozoncp/ocp-knowledge-api/pkg/ocp-knowledge-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KnowledgeApi struct {
	api.UnimplementedOcpKnowledgeApiServer
	repo repo.KnowledgeRepo
}

func NewKnowledgeApi(repo repo.KnowledgeRepo) api.OcpKnowledgeApiServer {
	return &KnowledgeApi{
		repo: repo,
	}
}

func (r *KnowledgeApi) CreateKnowledgeV1(ctx context.Context, request *api.CreateKnowledgeV1Request) (*api.CreateKnowledgeV1Response, error) {
	log.Info().Msgf("Create Knowledge %v", request.GetKnowledge())

	knowledge := models.Knowledge{
		UserId: request.Knowledge.UserId,
		Topic:  request.Knowledge.Topic,
		Text:   request.Knowledge.Text,
	}

	resultId, err := r.repo.AddKnowledge(knowledge)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.CreateKnowledgeV1Response{
		Id: resultId,
	}, nil
}

func (r *KnowledgeApi) DescribeKnowledgeV1(ctx context.Context, request *api.DescribeKnowledgeV1Request) (*api.DescribeKnowledgeV1Response, error) {
	log.Info().Msgf("Describe Knowledge ID - %v", request.GetId())

	knowledge, err := r.repo.DescribeKnowledge(request.Id)

	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if knowledge == nil {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	result := api.Knowledge{

		UserId: knowledge.UserId,
		Topic:  knowledge.Topic,
		Text:   knowledge.Text,
	}

	return &api.DescribeKnowledgeV1Response{Knowledge: &result}, nil
}

func (r *KnowledgeApi) ListKnowledgeV1(ctx context.Context, request *api.ListKnowledgeV1Request) (*api.ListKnowledgeV1Response, error) {
	log.Info().Msgf("List Knowledge. Offset - %v; Limit - %v", request.GetOffset(), request.GetLimit())

	knowledge, err := r.repo.ListKnowledge(request.Limit, request.Offset)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*api.Knowledge, 0, len(knowledge))
	for _, k := range knowledge {
		result = append(
			result,
			&api.Knowledge{
				UserId: k.UserId,
				Topic:  k.Topic,
				Text:   k.Text,
			})
	}

	return &api.ListKnowledgeV1Response{Knowledge: result}, nil
}

func (r *KnowledgeApi) RemoveKnowledgeV1(ctx context.Context, request *api.RemoveKnowledgeV1Request) (*api.RemoveKnowledgeV1Response, error) {
	log.Info().Msgf("Remove knowledge ID - %v", request.GetId())

	isRemoved, err := r.repo.RemoveKnowledge(request.Id)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.RemoveKnowledgeV1Response{IsRemoved: isRemoved}, nil
}
