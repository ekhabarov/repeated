## Reproducing an issue with protoc-gen-struct-transformer

### Step
* Install necessary protoc plugins: `protoc-gen-gogofaster`, `protoc-gen-struct-transformer`
* Run protoc
```shell
protoc \
  -I ${GOPATH}/src/github.com/bold-commerce \
  -I ${GOPATH}/src/github.com/gogo \
  -I . \
  --gogofaster_out=Mprotoc-gen-struct-transformer/options/annotations.proto=github.com/bold-commerce/protoc-gen-struct-transformer/options:./pb \
  --struct-transformer_out=package=transform,goimports=true:. ./service.proto
```
* Build `transform` package.
