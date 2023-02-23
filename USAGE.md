<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/chord-oms-go-sdk"
    "github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/operations"
)

func main() {
    s := sdk.New()
    
    req := operations.GetUserAddressBookRequest{
        Security: operations.GetUserAddressBookSecurity{
            APIKey: shared.SchemeAPIKey{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        },
        PathParams: operations.GetUserAddressBookPathParams{
            UserID: 548814,
        },
    }
    
    res, err := s.AddressBooks.GetUserAddressBook(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddressBook != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->