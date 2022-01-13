# Find Exports

Parse a shell script from stdin and output all export declarations in an easy to read format.

## Usage Example

``` sh
$ npx sls environment | findexports
KMS_ARN="arn:aws:kms:us-east-1:4583122534343"
BUS_NAME="dev-bus-1"
LAMBDA_REGION="us-east-1"
AWS_ACCOUNT_ID="23749274927458"
```

## Install

```
$ go install github.com/icholy/findexports@latest
```

## Notes:

This output format is compatible with [env-cmd](https://www.npmjs.com/package/env-cmd).
