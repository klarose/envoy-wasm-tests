#include <string>

#include "proxy_wasm_intrinsics.h"

//! @class StaticHeaderRootContext is the root of the StaticHeader
//! plugin.
//!
//! The StaticHeader plugin adds a single header. Fun stuff!
class StaticHeaderRootContext : public RootContext {
 public:
  explicit StaticHeaderRootContext(uint32_t id, std::string_view root_id)
      : RootContext(id, root_id) {}

  bool onConfigure(size_t) override;

 private:
};

class StaticHeaderContext : public Context {
 public:
  explicit StaticHeaderContext(uint32_t id, RootContext* root) : Context(id, root) {}

  FilterHeadersStatus onResponseHeaders(uint32_t, bool) override;

 private:
  inline StaticHeaderRootContext* rootContext() {
    return dynamic_cast<StaticHeaderRootContext*>(this->root());
  }
};
