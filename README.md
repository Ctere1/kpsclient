[![Go Report Card](https://goreportcard.com/badge/github.com/Ctere1/kpsclient)](https://goreportcard.com/report/github.com/Ctere1/kpsclient)
[![pkg.go.dev](https://pkg.go.dev/badge/github.com/Ctere1/kpsclient)](https://pkg.go.dev/github.com/Ctere1/kpsclient)

# About

A simple, standalone Go package for querying the Turkish Population and Citizenship Affairs (KPS) v2 services.

This package automates the WS-Trust (STS) request flow, sends HMAC-SHA1 signed SOAP requests using the SAML-based key from the STS, and parses service responses into a convenient Go struct.

> [!Note]   	
> The package uses HMAC-SHA1 as expected by NVI/KPS services.

> [!Warning]
> KPS services perform real identity validation; without test credentials, requests may fail or be denied.

## Features

- Automatic WS-Trust (Token Service / STS) authentication flow.
- Signs SOAP messages with HMAC-SHA1 (as required by KPS services).
- Parses SOAP service responses into a meaningful `Result` struct.
- Lightweight, independent and easy-to-use API.

## Installation

Use Go modules:

```bash
go get github.com/Ctere1/kpsclient
```

Or import directly in your code:

```go
import kpsclient "github.com/Ctere1/kpsclient"
```

## Quick Start

A usage example is available in `test/main.go`. Here's a brief summary:

```go
package main

import (
  "context"
  "time"
  "github.com/Ctere1/kpsclient"
)

func example() {
  client := kpsclient.New("USERNAME", "PASSWORD", nil)
  req := kpsclient.QueryRequest{
    TCNo:        "99999999999",
    FirstName:   "JOHN",
    LastName:    "DOE",
    BirthYear:   "1990",
    BirthMonth:  "01",
    BirthDay:    "01",
    SerialNumber: "ABC123456", // Optional: for new generation identity cards
  }
  ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
  defer cancel()
  res, err := client.DoQuery(ctx, req)
  if err != nil {
    // error handling
  }
  // Use res.Result structure as needed
  _ = res
}
```

## API Overview

- `func New(username, password string, httpClient *http.Client) *Client`
  - Creates a new `Client`. If `httpClient` is nil, a default client with a 30s timeout is used.

- `func (c *Client) DoQuery(ctx context.Context, req QueryRequest) (Result, error)`
  - Runs the STS authentication flow, sends a signed service request, and parses the response.

- `type QueryRequest` (input)
  - Fields:
    - `TCNo` (string): 11-digit Turkish Republic ID Number (required)
    - `FirstName` (string): Official first name, uppercase (required)
    - `LastName` (string): Official surname, uppercase (required)
    - `BirthYear` (string): Year of birth, format YYYY (required)
    - `BirthMonth` (string): Month of birth, format MM (optional)
    - `BirthDay` (string): Day of birth, format DD (optional)
    - `SerialNumber` (string): New generation identity card serial number (optional, used in certain verifications)

- `type Result` (output)
  - Fields:
    - `Status` (bool)
    - `Code` (1 success, 2 error/not found, 3 deceased)
    - `Message` (short message as string)
    - `Person` (`tc_vatandasi`, `yabanci`, `mavi`)
    - `Extra` (map of additional data)
    - `Raw` (raw SOAP/XML response)

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
