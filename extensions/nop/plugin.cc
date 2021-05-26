#include "plugin.h"

// Boilerplate code to register the extension implementation.
static RegisterContextFactory register_Nop(
    CONTEXT_FACTORY(NopContext),
    ROOT_FACTORY(NopRootContext));

bool NopRootContext::onConfigure(size_t) { return true; }
