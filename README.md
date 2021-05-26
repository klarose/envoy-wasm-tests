# envoy-wasm-tests

## Overview

This repo contains a series of envoy extensions compiled to wasm. They may be built
from different languages accordingly. The repo was bootstrapped by using the istio
[example extensions](https://github.com/istio-ecosystem/wasm-extensions/tree/6dd2be0310c7e1db793e472e558d320f53c4c70d/example).

These extensions are mostly used for testing/proving out issues/etc

## Extensions

- nop: This extension does nothing. It just proves that you can compile and load it.

## Development

To build, run `bazel build extensions/...`. The compiled .wasm files will be found
in `bazel-bin/extensions`.

Currently, our only test framework is the integration tests described in the istio
[walkthrough](https://github.com/istio-ecosystem/wasm-extensions/blob/master/doc/write-a-wasm-extension-with-cpp.md).

To run it, simply `go test -count=1 ./...`. In the future I will add proper unit tests and hook them into
a bazel target.

## License

See LICENSE.md
