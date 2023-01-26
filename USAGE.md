<!-- Start SDK Example Usage -->
```go
package main

import (
    "github.com/speakeasy-sdks/chord-oms-go-sdk"
    "github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/operations"
)

func main() {
    s := sdk.New()
    
    req := operations.GetUserAddressBookRequest{
        Security: operations.GetUserAddressBookSecurity{
            APIKey: shared.SchemeAPIKey{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        PathParams: operations.GetUserAddressBookPathParams{
            UserID: "sit",
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