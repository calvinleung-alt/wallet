
# Wallet

## Code Execution
```
go run cmd/main.go
```

## Description
These repository mainly consists of client exposed package named `wallet`.

And all the apis are lived in `api_op` prefixed files under `wallet` package.

For the server side simulation, it consists of several simplified commands lived in `internal/cmd` package.

Where data handling is located at `internal/data` and mocked behaviours of repository and transaction is applied.

`types` package contains all shared types between public api and internal usage.

## API

### Types

```Client```

#### Methods

```
func (c *Client) Deposit(*types.DepositInput) (*types.DepositOutput, error)

func (c *Client) Withdraw(*types.WithdrawInput) (*types.WithdrawOutput, error)

func (c *Client) Send(*types.SendInput) (*types.SendOutput, error)

func (c *Client) CheckBalance(*types.CheckBalanceInput) (*types.CheckBalanceOutput, error)
```

### Functional Requirement

User can be able to apply the features like Deposit, Withdraw, Send, CheckBalance

### Non-functional Requirement

In order to prevent double spending or undesired behaviour due to race condition and concurrency,

Transactional process will be applied and to ensure data operation linearizability.

### Development Time
Time spent on this test is around 5 hours.

### Features not implemented

Exact implementation of database connection and transactional behaviour is delayed for the sake of simplicity.

Also, it is supposed that a backend server should developed and up for serving the endpoints since the deliverable is mainly about the

client package for the reusable library.