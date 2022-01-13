# Find Exports

Parse a shell script and output all export declarations in an easy to read format.

## Usage Example

``` sh
$ findexports ~/.bashrc
PATH=$PATH:/usr/local/go/bin
PATH=$PATH:/home/icholy/go/bin
PATH=$PATH:/usr/local/swift/swift-5.5.2-RELEASE-ubuntu20.04/usr/bin
PATH=$PATH:/usr/local/lua-language-server/bin
NVM_DIR="$HOME/.nvm"
```

## Install

```
$ go install github.com/icholy/findexports@latest
```

## Notes:

This output format is compatible with [env-cmd](https://www.npmjs.com/package/env-cmd).
