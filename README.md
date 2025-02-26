# GO WORKOS integration example using AuthKit and Fiber
An example application demonstrating how to authenticate users with AuthKit (https://workos.com) and the WorkOS GO SDK.

> Refer to the [User Management](https://workos.com/docs/user-management) documentation for reference.

## Prerequisites

You will need a [WorkOS account](https://dashboard.workos.com/signup).

## Requirements

GO 1.23 (tested)

## Running the example

1. In the [WorkOS dashboard](https://dashboard.workos.com), head to the Redirects tab and create a [sign-in callback redirect](https://workos.com/docs/user-management/1-configure-your-project/configure-a-redirect-uri) for `http://localhost:8080/auth/callback`.

2. After creating the redirect URI, navigate to the API keys tab and copy the _Client ID_ and the _Secret Key_. Rename the `.env.example` file to `.env` and supply your Client ID and API key as environment variables.

3. Verify your `.env` file has the following variables filled.

   ```bash
   WORKOS_CLIENT_ID=<YOUR_CLIENT_ID>
   WORKOS_API_KEY=<YOUR_API_SECRET_KEY>
   WORKOS_AUTHKIT_URL=<YOUR_AUTHKIT_URL>
   ```

4. Install the dependencies

   ```bash
   go get ./...
   ```

5. Run the following command and navigate to [http://localhost:8080](http://localhost:8080).

   ```bash
   go run cmd/web/main.go
   ```
