# Chord OMS Go SDK

Go SDK for Chord [OMS API](https://chord.stoplight.io/docs/chord-oms/811626a87baa4-oms-api). The modern commerce platform built for growth. Created by experts in DTC innovation, Chord's headless commerce platform gives you the power to do anything. 

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/chord-oms-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
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

## Authentication

The API provides two means of authentication: one is through your user's API key, while the other is through an order's guest token.

### API key
You can use your API key to access all resources in the API. The API key must be passed in the Authorization header in the following form:

`Authorization: Bearer API_KEY`

As an admin, you can find your API token in the admin section under Users > Your e-email > API Access (at admin/users/<user_id>/edit)

<!-- Start API Key Example Usage -->
Example:

```bash
curl --header "Authorization: Bearer 1a6a9936ad150a2ee345c65331da7a3ccc2de" http://plant-staging.assembly-api.com/api/stores
```
<!-- End API Key Example Usage -->

### Order token
For allowing guests to manage their cart and place their order, you can use the order's guest token. This token is contained in the guest_token property of the order, and it allows you to perform certain checkout-related operations on the order such as managing line items, completing the checkout flow etc.

The order token must be passed in the X-Spree-Order-Token header in the following form:

`X-Spree-Order-Token: ORDER_TOKEN`

If you are already providing an API key, you don't need to also provide the order token (although you may do so). More information on authentication can be found [here](https://chord.stoplight.io/docs/chord-oms/ZG9jOjEwODE5NTQ-authentication)
<!-- Start SDK Available Operations -->
## SDK Available Operations

### Address books

* `GetUserAddressBook` - Get user address book
* `RemoveAddressFromUserAddressBook` - Remove address from user address book
* `UpdateUserAddressBook` - Update user address book

### Addresses

* `GetCheckoutAddress` - Get checkout address
* `GetOrderAddress` - Get order address
* `UpdateCheckoutAddress` - Update checkout address
* `UpdateOrderAddress` - Update order address

### Checkouts

* `AdvanceCheckout` - Advance checkout
* `CompleteCheckout` - Complete checkout
* `FinalizeCheckout` - Finalize the checkout
* `TransitionCheckout` - Transition checkout
* `UpdateCheckout` - Update checkout

### Classifications

* `UpdateClassification` - Update classification

### Configuration

* `GetConfig` - Get system configuration
* `GetMoneyConfig` - Get money configuration

### Countries

* `GetCountry` - Get country
* `ListCountries` - List countries

### Coupon codes

* `ApplyOrderCouponCode` - Apply order coupon code
* `CreateOrderCouponCode` - Create order coupon code
* `DeleteOrderCouponCode` - Delete order coupon code

### Credit cards

* `ListUserCreditCards` - List user credit cards
* `UpdateCreditCard` - Update credit card

### Gift Cards

* `RedeemGiftCard` - Redeems a Gift Card

### Images

* `CreateProductImage` - Create product image
* `CreateVariantImage` - Create variant image
* `DeleteProductImage` - Delete product image
* `DeleteVariantImage` - Delete variant image
* `GetProductImage` - Get product image
* `GetVariantImage` - Get variant image
* `ListProductImages` - List product images
* `ListVariantImages` - List variant images
* `UpdateProductImage` - Update product image
* `UpdateVariantImage` - Update variant image

### Inventory units

* `GetInventoryUnit` - Get inventory unit
* `UpdateInventoryUnit` - Update inventory unit

### Line items

* `CreateCheckoutLineItem` - Create checkout line item
* `CreateOrderLineItem` - Create order line item
* `CrossSellOrderLineItem` - Cross-sell order line items
* `DeleteCheckoutLineItem` - Delete checkout line item
* `DeleteOrderLineItem` - Delete order line item
* `UpdateCheckoutLineItem` - Update checkout line item
* `UpdateOrderLineItem` - Update order line item

### Option types

* `CreateOptionType` - Create option type
* `DeleteOptionType` - Delete option type
* `GetOptionType` - Get option type
* `ListOptionTypes` - List option types
* `UpdateOptionType` - Update option type

### Option values

* `CreateOptionTypeValue` - Create option type value
* `CreateOptionValue` - Create option value
* `DeleteOptionTypeValue` - Delete option type value
* `DeleteOptionValue` - Delete option value
* `GetOptionTypeValue` - Get option type value
* `GetOptionValue` - Get option value
* `ListOptionTypeValues` - List option type values
* `ListOptionValues` - List option values
* `UpdateOptionTypeValue` - Update option type value
* `UpdateOptionValue` - Update option value

### Orders

* `CancelOrder` - Cancel order
* `CreateOrder` - Create order
* `EmptyOrder` - Empty order
* `GetCurrentOrder` - Get current order
* `GetOrder` - Get order
* `ListCurrentUserOrders` - List current user's orders.
* `ListOrders` - List orders
* `UpdateOrder` - Update order

### Payments

* `AuthorizeCheckoutPayment` - Authorize checkout payment
* `AuthorizeOrderPayment` - Authorize order payment
* `CaptureCheckoutPayment` - Capture checkout payment
* `CaptureOrderPayment` - Capture order payment
* `CreateCheckoutPayment` - Create checkout payment
* `CreateOrderPayment` - Create order payment
* `CreditCheckoutPayment` - Credit checkout payment
* `CreditOrderPayment` - Credit order payment
* `GetCheckoutPayment` - Get checkout payment
* `GetOrderPayment` - Get order payment
* `ListCheckoutPayments` - List checkout payments
* `ListOrderPayments` - List order payments
* `PurchaseCheckoutPayment` - Purchase checkout payment
* `PurchaseOrderPayment` - Purchase order payment
* `UpdateCheckoutPayment` - Update checkout payment
* `UpdateOrderPayment` - Update order payment
* `VoidCheckoutPayment` - Void checkout payment
* `VoidOrderPayment` - Void order payment

### Product properties

* `CreateProductProperty` - Create product property
* `DeleteProductProperty` - Delete product property
* `GetProductProperty` - Get product property
* `ListProductProperties` - List product properties
* `UpdateProductProperty` - Update product property

### Products

* `CreateProduct` - Create product
* `DeleteProduct` - Delete product
* `GetProduct` - Get product
* `ListProductImages` - List product images
* `ListProducts` - List products
* `ListTaxonProducts` - List taxon products
* `UpdateProduct` - Update product

### Promotions

* `GetPromotion` - Get promotion

### Properties

* `CreateProperty` - Create property
* `DeleteProperty` - Delete property
* `GetProperty` - Get property
* `ListProperties` - List properties
* `UpdateProperty` - Update property

### Quick Checkout

* `QuickCheckout` - Quick Checkout

### Refunds

* `ListReturnAuthorizationRefunds` - List return authorization refunds

### Return authorizations

* `CancelCheckoutReturnAuthorization` - Cancel checkout return authorization
* `CancelOrderReturnAuthorization` - Cancel order return authorization
* `CreateCheckoutReturnAuthorization` - Create checkout return authorization
* `CreateOrderReturnAuthorization` - Create order return authorization
* `DeleteCheckoutReturnAuthorization` - Delete checkout return authorization
* `DeleteOrderReturnAuthorization` - Delete order return authorization
* `GetCheckoutReturnAuthorization` - Get checkout return authorization
* `GetOrderReturnAuthorization` - Get order return authorization
* `ListCheckoutReturnAuthorization` - List checkout return authorizations
* `ListOrderReturnAuthorizations` - List order return authorizations
* `ListReturnAuthorization` - List return authorizations
* `UpdateCheckoutReturnAuthorization` - Update checkout return authorization
* `UpdateOrderReturnAuthorization` - Update order return authorization

### Return items

* `ListReturnAuthorizationReturnItems` - List return authorization return items

### Shipments

* `AddShipmentItem` - Add shipment item
* `CreateShipment` - Create shipment
* `ListShipmentEstimatedRates` - List shipment estimated rates
* `ListShipments` - List shipments
* `ListUserShipments` - List user's shipments
* `ReadyShipment` - Ready shipment
* `RemoveShipmentID` - Remove shipment item
* `SelectShipmentShippingMethod` - Select shipment shipping method
* `ShipShipment` - Ship shipment
* `TransferShipmentItemToLocation` - Transfer shipment item to location
* `TransferShipmentItemToShipment` - Transfer shipment item to shipment
* `UpdateShipment` - Update shipment

### States

* `GetCountryState` - Get country state
* `GetState` - Get state
* `ListCountryStates` - List country states
* `ListStates` - List states

### Stock items

* `CreateStockLocationItem` - Create stock location item
* `DeleteStockItem` - Delete stock item
* `DeleteStockLocationItem` - Delete stock location item
* `GetStockLocationItem` - Get stock location item
* `ListStockLocationItems` - List stock location items
* `UpdateStockItem` - Update stock item
* `UpdateStockLocationItem` - Update stock location item

### Stock locations

* `CreateStockLocation` - Create stock location
* `DeleteStockLocation` - Delete stock location
* `GetStockLocation` - Get stock location
* `ListStockLocations` - List stock location
* `UpdateStockLocation` - Update stock location

### Stock movements

* `CreateStockLocationMovement` - Create stock location movement
* `GetStockLocationMovement` - Get stock location movement
* `ListStockLocationMovements` - List stock location movements

### Store credit events

* `ListCurrentUserStoreCreditEvents` - List current user's store credit events

### Stores

* `CreateStore` - Create store
* `DeleteStore` - Delete store
* `GetStore` - Get store
* `ListStores` - List stores
* `UpdateStore` - Update store

### Subscriptions

* `GetUsersUserIDSubscriptions` - List user subscriptions
* `GetUsersUserIDSubscriptionsID` - Get subscription
* `PatchUsersUserIDSubscriptionsID` - Update subscription
* `PostUsersUserIDSubscriptionsIDCancel` - Cancel a subscription
* `PostUsersUserIDSubscriptionsIDSkip` - Skip a subscription
* `PostUsersUserIDSubscriptions` - Create a subscription

### Taxonomies

* `CreateTaxonomy` - Create taxonomy
* `DeleteTaxonomy` - Delete taxonomy
* `GetTaxonomy` - Get taxonomy
* `GetTaxonomyJstree` - Get taxonomy jsTree
* `ListTaxonomies` - List taxonomies
* `UpdateTaxonomy` - Update taxonomy

### Taxons

* `CreateTaxonomyTaxon` - Create taxonomy taxon
* `DeleteTaxonomyTaxon` - Delete taxonomy taxon
* `GetTaxonJstree` - Get taxon jsTree
* `GetTaxonomyTaxon` - Get taxonomy taxon
* `ListTaxonomyTaxons` - List taxonomy taxons
* `ListTaxons` - List taxons
* `UpdateTaxonomyTaxon` - Update taxonomy taxon

### Users

* `CreateUser` - Create user
* `DeleteUser` - Delete user
* `GetUser` - Get user
* `GetUsersMe` - Get current user
* `ListUsers` - List users
* `UpdateUser` - Update user

### Variants

* `CreateProductVariant` - Create product variant
* `CreateVariant` - Create variant
* `DeleteProductVariant` - Delete product variant
* `DeleteVariant` - Delete variant
* `GetProductVariant` - Get product variant
* `GetVariant` - Get variant
* `GetVariantStockAvailability` - Return the variant availability
* `ListProductVariants` - List product variants
* `ListVariants` - List variants
* `UpdateProductVariant` - Update product variant
* `UpdateVariant` - Update variant

### Zones

* `CreateZone` - Create zone
* `DeleteZone` - Delete zone
* `GetZone` - Get zone
* `ListZones` - List zones
* `UpdateZone` - Update zone

<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
