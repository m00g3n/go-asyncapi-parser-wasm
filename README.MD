# AsyncAPI + WebAssembly + GoLang POC

## Requirements

- golang 1.11+
- docker

## Usage

### Build wasm file 

```bash
GO111MODULE=on GOOS=js GOARCH=wasm go build -o ./html/main.wasm
```

```
docker run --rm -p 8080:80 --name some-nginx \
    -v $(pwd)/html:/usr/share/nginx/html:ro \
    -v $(pwd)/nginx/mime.types:/etc/nginx/mime.types:ro \
    -v $(pwd)/nginx/nginx.conf:/etc/nginx/nginx.conf:ro \
    -d nginx
```

### Visit http://localhost:8080 and open browser console, next call parse function with the following command:

```javascript
JSON.parse(parse(`asyncapi: '2.0.0-rc1'
id: 'urn:com:ec:stage:server'
info:
  title: ec-all-events
  version: v1
  description: EC Events v1
channels:
  customer.created.v1:
    subscribe:
      summary: Customer Register Event v2
      message:
        payload:
          type: object
          required:
          - storeUid
          - customerId
          - customerUid
          properties:
            customerUid:
              type: string
              example: adam@hybris.com
              description: Email of a Customer
              title: Customer uid
            storeUid:
              type: string
              example: 4caad296
              description: Unique id of a Store
              title: StoreUid
            customerId:
              type: string
              example: 0ed118f9474e
              description: Unique id of a Customer
              title: Customer uid
  quote.orderplaced.v1:
    subscribe:
      summary: Order Placed for Quote Event v1
      message:
        payload:
          type: object
          required:
          - quoteCode
          - orderCode
          properties:
            quoteCode:
              type: string
              example: '76272727'
              description: Code of Quote
              title: Quote Code
            orderCode:
              type: string
              example: '76272727'
              description: Code of Order
              title: Order Code
  quote.expiresoonevent.v1:
    subscribe:
      summary: Quote Expiring Soon Event v1
      message:
        payload:
          type: object
          required:
          - quoteCode
          - quoteUserUid
          - quoteUserType
          - storeUid
          - customerUid
          properties:
            customerUid:
              type: string
              example: adam@hybris.com
              description: Email Address of Customer
              title: Customer ID
            storeUid:
              type: string
              example: 4caad296
              description: ID of the Store (BaseStore)
              title: Store ID
            quoteCode:
              type: string
              example: '76272727'
              description: Code of Quote
              title: Quote Code
            quoteUserUid:
              type: string
              example: adam@hybris.com
              description: ID Of User Canceling Quote
              title: ID Of User
            quoteUserType:
              type: string
              example: BUYER
              description: User Type
              title: User Type`, true).Document)
```
