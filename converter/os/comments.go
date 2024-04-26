package os

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/kokoichi206/searcher/converter/model"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
)

func (c *client) IndexComment(m *model.Comment) error {
	doc, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal comment: %w", err)
	}

	req := opensearchapi.IndexRequest{
		Index: "comments",
		Body:  bytes.NewReader(doc),
	}

	insertResponse, err := req.Do(context.Background(), c.osc)
	if err != nil {
		return fmt.Errorf("failed to insert comment: %w", err)
	}
	defer insertResponse.Body.Close()

	// io.Copy(os.Stdout, insertResponse.Body)

	return nil
}
