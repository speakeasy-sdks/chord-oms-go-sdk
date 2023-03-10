package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/utils"
	"net/http"
	"strings"
)

type StoreCreditEvents struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewStoreCreditEvents(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *StoreCreditEvents {
	return &StoreCreditEvents{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// ListCurrentUserStoreCreditEvents - List current user's store credit events
// Lists the current user's store credit events.
func (s *StoreCreditEvents) ListCurrentUserStoreCreditEvents(ctx context.Context, request operations.ListCurrentUserStoreCreditEventsRequest) (*operations.ListCurrentUserStoreCreditEventsResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/store_credit_events/mine"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListCurrentUserStoreCreditEventsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCurrentUserStoreCreditEventsPaginationData
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PaginationData = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCurrentUserStoreCreditEvents401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListCurrentUserStoreCreditEvents401ApplicationJSONObject = out
		}
	}

	return res, nil
}
