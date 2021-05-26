#include "plugin.h"

// Boilerplate code to register the extension implementation.
static RegisterContextFactory register_StaticHeader(CONTEXT_FACTORY(StaticHeaderContext),
                                           ROOT_FACTORY(StaticHeaderRootContext));

bool StaticHeaderRootContext::onConfigure(size_t) { return true; }

FilterHeadersStatus StaticHeaderContext::onResponseHeaders(uint32_t, bool) {
  LOG_DEBUG("StaticHeader: Got some headers");
  replaceResponseHeader("Dummy-Header", "some value");
  return FilterHeadersStatus::Continue;
}
