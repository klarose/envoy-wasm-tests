#include <string>

#include "proxy_wasm_intrinsics.h"

//! @class NopRootContext is the root of the Nop
//! plugin.
//!
//! The Nop plugin does nothing. Enjoy!
class NopRootContext : public RootContext {
 public:
  explicit NopRootContext(uint32_t id, std::string_view root_id)
      : RootContext(id, root_id) {}

  bool onConfigure(size_t) override;

 private:
};

class NopContext : public Context {
 public:
  explicit NopContext(uint32_t id, RootContext* root) : Context(id, root) {}

  FilterHeadersStatus onResponseHeaders(uint32_t, bool) override;

 private:
  inline NopRootContext* rootContext() {
    return dynamic_cast<NopRootContext*>(this->root());
  }
};
