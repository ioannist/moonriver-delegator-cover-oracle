# Oracle service for Moonriver Staking Cover contract
Collects staking data from Moonriver and reports to the contract at https://github.com/ioannist/moonriver-delegator-cover-contract

## How it works
* The service checks if there is a new era (round) to report, and if it has already reported the last completed era
* It starts collecting staking data from the parachain and signs and sends the data to the OracleMaster contract.

## Requirements
* The service is a binary executable compiled from Go. You can run it under systemd. TODO
* The service is also available as an AWS Lambda microservice. To run it this way you will need to deploy it to your AWS account.


## Build AWS mircroservice
* Make sure your system has AWS credentials stored 
* Edit "--profile=mb" to use your AWS profile, or remove it if you are only using one AWS account in your system
* Make sure Golang > 1.19 is installed
* Execute the following to build and deploy the service
```shell
./deploy
```

## Configuration options
Go to your AWS Lambda in the AWS console and save the following Env variables in Configuration:
`ETH_URL`: the https endpoint for Ethereum calls on Moonriver **Required**.
`GAS_LIMIT`: 10000000 will do **Required**.
`ORACLE_MASTER`: 0x09b3941c7c75928770a10FbabAd706cBddf559Ee **Required**.
`ORACLE_MEMBER`: the private key of your account that you will use to sign oracle reports **Required**.
`RPC_URL`: the rpc endpoint for Moonriver **Required**.
