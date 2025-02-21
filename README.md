# kryptonim

!!! Please note that Auth header is required - you can turn Authorization off in .env !!!

## Overview

This project provides endpoints to get exchange rates and perform cryptocurrency exchanges.

## Endpoints

### GET /rates

Fetches the latest exchange rates for the specified currencies.

**Query Parameters:**
- `currencies` (required): A comma-separated list of currency codes.

**Headers:**
- `Auth` (required): The authorization token.

### GET /exchange

Performs a cryptocurrency exchange based on the provided parameters.

**Query Parameters:**
- `from` (required): The cryptocurrency to exchange from.
- `to` (required): The cryptocurrency to exchange to.
- `amount` (required): The amount of cryptocurrency to exchange.

**Headers:**
- `Auth` (required): The authorization token.

## Environment Variables

- `PORT`: The port on which the server runs.
- `AUTH_REQ`: Set to `ON` to require the `Auth` header, `OFF` to disable authorization.
- `AUTH_TOKEN`: The token required for authorization.
- `OPENEXCHANGERATES_URL`: The URL for fetching exchange rates.
- `OPENEXCHANGERATES_APP_ID`: The app ID for accessing the exchange rates API.

## Running the Server

1. Set up the environment variables in a `.env` file.
2. Run the server using the following command:

```sh
go run cmd/main.go
```
