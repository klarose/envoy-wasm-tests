workspace(name = "envoy_wasm_tests")
# load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

# protobuf_deps()

# Pulls proxy wasm cpp SDK with a specific SHA
PROXY_WASM_CPP_SDK_SHA = "f5ecda129d1e45de36cb7898641ac225a50ce7f0"
PROXY_WASM_CPP_SDK_SHA256 = "0f675ef5c4f8fdcf2fce8152868c6c6fd33251a0deb4a8fc1ef721f9ed387dbc"

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "proxy_wasm_cpp_sdk",
    sha256 = PROXY_WASM_CPP_SDK_SHA256,
    strip_prefix = "proxy-wasm-cpp-sdk-" + PROXY_WASM_CPP_SDK_SHA,
    url = "https://github.com/proxy-wasm/proxy-wasm-cpp-sdk/archive/" + PROXY_WASM_CPP_SDK_SHA + ".tar.gz",
)

load("@proxy_wasm_cpp_sdk//bazel/dep:deps.bzl", "wasm_dependencies")

wasm_dependencies()

load("@proxy_wasm_cpp_sdk//bazel/dep:deps_extra.bzl", "wasm_dependencies_extra")

wasm_dependencies_extra()

### optional imports ###
# To import commonly used libraries from istio proxy, such as base64, json, and flatbuffer.
IO_ISTIO_PROXY_SHA = "a4e4b362e334f4275ac2154fb831ddf3cf45db03"
IO_ISTIO_PROXY_SHA256 = "23296eb7a31e42eb2845488f6869539332dc705c72a3b0ebd196bdaf6d16c23c"

http_archive(
    name = "io_istio_proxy",
    sha256 = IO_ISTIO_PROXY_SHA256,
    strip_prefix = "proxy-" + IO_ISTIO_PROXY_SHA,
    url = "https://github.com/istio/proxy/archive/" + IO_ISTIO_PROXY_SHA + ".tar.gz",
)

load("@envoy_wasm_tests//bazel:wasm.bzl", "wasm_libraries")

wasm_libraries()
