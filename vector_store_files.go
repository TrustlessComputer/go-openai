package openai

import (
	"context"
	"net/http"
)

type VectorStoreFileRequest struct {
	FileId string `json:"file_id"`
}

type VectorStoreFileResp struct {
	Id               string      `json:"id"`
	Object           string      `json:"object"`
	UsageBytes       int         `json:"usage_bytes"`
	CreatedAt        int         `json:"created_at"`
	VectorStoreId    string      `json:"vector_store_id"`
	Status           string      `json:"status"`
	LastError        interface{} `json:"last_error"`
	ChunkingStrategy struct {
		Type   string `json:"type"`
		Static struct {
			MaxChunkSizeTokens int `json:"max_chunk_size_tokens"`
			ChunkOverlapTokens int `json:"chunk_overlap_tokens"`
		} `json:"static"`
	} `json:"chunking_strategy"`

	httpHeader
}

func (c *Client) CreateVectorStoreFile(ctx context.Context, request VectorStoreFileResp) (response VectorStoreFileResp, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(vectorStoreSuffix), withBody(request),
		withBetaAssistantVersion(c.config.AssistantVersion))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
