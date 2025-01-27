package linodego

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// LongviewClient represents a LongviewClient object
type LongviewClient struct {
	ID int `json:"id"`
	// UpdatedStr string `json:"updated"`
	// Updated *time.Time `json:"-"`
}

// LongviewClientsPagedResponse represents a paginated LongviewClient API response
type LongviewClientsPagedResponse struct {
	*PageOptions
	Data []LongviewClient `json:"data"`
}

// endpoint gets the endpoint URL for LongviewClient
func (LongviewClientsPagedResponse) endpoint(c *Client, _ ...any) string {
	endpoint, err := c.LongviewClients.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

func (resp *LongviewClientsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(LongviewClientsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*LongviewClientsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListLongviewClients lists LongviewClients
func (c *Client) ListLongviewClients(ctx context.Context, opts *ListOptions) ([]LongviewClient, error) {
	response := LongviewClientsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetLongviewClient gets the template with the provided ID
func (c *Client) GetLongviewClient(ctx context.Context, id string) (*LongviewClient, error) {
	e, err := c.LongviewClients.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, id)
	r, err := c.R(ctx).SetResult(&LongviewClient{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*LongviewClient), nil
}
