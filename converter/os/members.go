package os

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kokoichi206/searcher/converter/model"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
)

func (c *client) IndexMember(m *model.Member) error {
	doc, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal member: %w", err)
	}

	req := opensearchapi.IndexRequest{
		Index: "members",
		Body:  bytes.NewReader(doc),
	}

	insertResponse, err := req.Do(context.Background(), c.osc)
	if err != nil {
		return fmt.Errorf("failed to insert member: %w", err)
	}
	defer insertResponse.Body.Close()

	io.Copy(os.Stdout, insertResponse.Body)

	return nil
}

func (c *client) IndexBlog(m *model.Blog) error {
	doc, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal blog: %w", err)
	}

	req := opensearchapi.IndexRequest{
		Index: "blogs",
		Body:  bytes.NewReader(doc),
	}

	insertResponse, err := req.Do(context.Background(), c.osc)
	if err != nil {
		return fmt.Errorf("failed to insert blog: %w", err)
	}
	defer insertResponse.Body.Close()

	io.Copy(os.Stdout, insertResponse.Body)

	return nil
}
