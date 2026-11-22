package openfga

import (
	"context"
	"fmt"

	openfga "github.com/openfga/go-sdk/client"

	"openhealth/internal/domain"
)

type Client struct {
	openfga *openfga.OpenFgaClient
}

// NewClient creates a new OpenFGA client.
func NewClient(uri, storeID, modelID string) (*Client, error) {
	client, err := openfga.NewSdkClient(&openfga.ClientConfiguration{
		ApiUrl:               uri,
		StoreId:              storeID,
		AuthorizationModelId: modelID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create fga client: %w", err)
	}

	return &Client{
		openfga: client,
	}, nil
}

// AddTuple - writes a single relationship tuple to OpenFGA.
func (c *Client) AddTuple(ctx context.Context, tuple domain.Tuple) error {
	body := openfga.ClientWriteRequest{
		Writes: []openfga.ClientTupleKey{
			{
				User:     tuple.User(),
				Relation: tuple.Relation,
				Object:   tuple.Object(),
			},
		},
	}

	_, err := c.openfga.Write(ctx).Body(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to add tuple %s - %w", tuple.ToString(), err)
	}

	return nil
}

// RemoveTuple - deletes a single relationship tuple from OpenFGA.
func (c *Client) RemoveTuple(ctx context.Context, tuple domain.Tuple) error {
	body := openfga.ClientWriteRequest{
		Deletes: []openfga.ClientTupleKeyWithoutCondition{
			{
				User:     tuple.User(),
				Relation: tuple.Relation,
				Object:   tuple.Object(),
			},
		},
	}

	_, err := c.openfga.Write(ctx).Body(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to remove tuple %s - %w", tuple.ToString(), err)
	}

	return nil
}

// Check - asks OpenFGA whether the tuple's user has the tuple's relation on the tuple's object.
func (c *Client) Check(ctx context.Context, tuple domain.Tuple) (bool, error) {
	body := openfga.ClientCheckRequest{
		User:     tuple.User(),
		Relation: tuple.Relation,
		Object:   tuple.Object(),
	}

	resp, err := c.openfga.Check(ctx).Body(body).Execute()
	if err != nil {
		return false, fmt.Errorf("failed to check tuple %s - %w", tuple.ToString(), err)
	}

	return resp.GetAllowed(), nil
}
