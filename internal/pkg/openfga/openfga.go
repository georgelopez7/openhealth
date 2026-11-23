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

// NewClient resolves the named store and its latest authorization model before
// creating the OpenFGA client.
func NewClient(ctx context.Context, uri, storeName string) (*Client, error) {
	adminClient, err := openfga.NewSdkClient(&openfga.ClientConfiguration{ApiUrl: uri})
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenFGA admin client: %w", err)
	}

	storeID, err := resolveStoreID(ctx, adminClient, storeName)
	if err != nil {
		return nil, err
	}

	modelOptions := openfga.ClientReadLatestAuthorizationModelOptions{StoreId: &storeID}
	latestModel, err := adminClient.ReadLatestAuthorizationModel(ctx).Options(modelOptions).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to read latest authorization model for store %q: %w", storeName, err)
	}
	if latestModel == nil || latestModel.AuthorizationModel == nil || latestModel.AuthorizationModel.GetId() == "" {
		return nil, fmt.Errorf("OpenFGA store %q has no authorization model", storeName)
	}

	client, err := openfga.NewSdkClient(&openfga.ClientConfiguration{
		ApiUrl:               uri,
		StoreId:              storeID,
		AuthorizationModelId: latestModel.AuthorizationModel.GetId(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenFGA client: %w", err)
	}

	return &Client{
		openfga: client,
	}, nil
}

func resolveStoreID(ctx context.Context, client *openfga.OpenFgaClient, storeName string) (string, error) {
	if storeName == "" {
		return "", fmt.Errorf("OpenFGA store name is required")
	}

	var continuationToken string
	var storeID string
	for {
		pageSize := int32(100)
		options := openfga.ClientListStoresOptions{PageSize: &pageSize}
		if continuationToken != "" {
			options.ContinuationToken = &continuationToken
		}

		response, err := client.ListStores(ctx).Options(options).Execute()
		if err != nil {
			return "", fmt.Errorf("failed to list OpenFGA stores: %w", err)
		}

		for _, store := range response.GetStores() {
			if store.GetName() != storeName {
				continue
			}
			if storeID != "" {
				return "", fmt.Errorf("multiple OpenFGA stores are named %q", storeName)
			}
			storeID = store.GetId()
		}

		continuationToken = response.GetContinuationToken()
		if continuationToken == "" {
			break
		}
	}

	if storeID == "" {
		return "", fmt.Errorf("OpenFGA store %q not found", storeName)
	}
	return storeID, nil
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
