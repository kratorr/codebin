package grpc

import (
	"context"
	"fmt"
	"time"

	"codebin/internal/converter"
	authIr "codebin/internal/interceptor"
	"codebin/internal/model"
	pbSnippet "codebin/pkg/proto/snippet"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SnippetService interface {
	CreateSnippet(ctx context.Context, snippet *model.Snippet, user *model.User) (*model.Snippet, error)
	DeleteSnippetByID(ctx context.Context, id string) error
	ReadLatest(ctx context.Context) ([]*model.Snippet, error)
}

var _ pbSnippet.SnippetV1Server = &SnippetHandler{}

type SnippetHandler struct {
	pbSnippet.UnimplementedSnippetV1Server
	service SnippetService
}

type SnippetHandlerDeps struct {
	Service SnippetService
}

func (h *SnippetHandler) ReadLatest(ctx context.Context, request *pbSnippet.ReadLatestRequest) (*pbSnippet.ReadLatestResponse, error) {
	resp, _ := h.service.ReadLatest(ctx)

	snippets := make([]*pbSnippet.Snippet, len(resp))
	fmt.Println(len(snippets))
	for v := range resp {
		snippets[v] = converter.SnippetToProto(resp[v])
	}
	r := &pbSnippet.ReadLatestResponse{Snippets: snippets}
	return r, nil
}

func NewSnippetHandler(deps *SnippetHandlerDeps) *SnippetHandler {
	return &SnippetHandler{service: deps.Service}
}

func (h *SnippetHandler) Create(ctx context.Context, req *pbSnippet.CreateSnippetRequest) (*pbSnippet.CreateSnippetResponse, error) {
	snippet := converter.ProtoToSnippet(req.GetSnippet())
	user, authenticated := authIr.FromContext(ctx)
	var err error
	var newSnippet *model.Snippet
	if !authenticated {
		newSnippet, err = h.service.CreateSnippet(ctx, snippet, nil)
	} else {
		newSnippet, err = h.service.CreateSnippet(ctx, snippet, &model.User{ID: user.ID, Username: user.Username})
	}
	if err != nil {
		return nil, err
	}

	return &pbSnippet.CreateSnippetResponse{ShortUrl: newSnippet.ShortURL}, nil
}

func (h *SnippetHandler) List(ctx context.Context, req *emptypb.Empty) (*pbSnippet.ListSnippetResponse, error) {
	time.Sleep(time.Second * 300)
	return &pbSnippet.ListSnippetResponse{}, nil
}
