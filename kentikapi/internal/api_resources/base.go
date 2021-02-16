package api_resources

import (
	"context"
	"encoding/json"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/validation"
)

// BaseAPI provides marshall/unmarshall + validation functionality for all resource APIs
type BaseAPI struct {
	Transport api_connection.Transport
}

// GetAndValidate retrieves json at "url" unmarshalls and validates against required fields defined in struct tags of "output"
// output must be pointer to object
func (b BaseAPI) GetAndValidate(ctx context.Context, url string, output interface{}) error {
	responseBody, err := b.Transport.Get(ctx, url)
	if err != nil {
		return err
	}
	
	if err = json.Unmarshal(responseBody, &output); err != nil {
		return err
	}

	if err = validation.CheckResponseRequiredFields("get", output); err != nil {
		return err
	}

	return nil
}