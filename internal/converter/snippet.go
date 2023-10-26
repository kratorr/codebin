package converter

import (
	"codebin/internal/model"
	pbSnippet "codebin/pkg/proto/snippet"
	"fmt"
)

func ProtoToSnippet(protoSnippet *pbSnippet.Snippet) *model.Snippet {
	return &model.Snippet{
		Title: protoSnippet.GetTitle(),
		Text:  protoSnippet.GetText(),
		Language: model.Language{
			ID: protoSnippet.GetLanguage(),
		},
	}
}

func SnippetToProto(snippet *model.Snippet) *pbSnippet.Snippet {
	fmt.Println(snippet)
	return &pbSnippet.Snippet{
		Title:    snippet.Title,
		Text:     snippet.ShortURL, // TODO just for test
		Language: snippet.Language.ID,
	}
}
