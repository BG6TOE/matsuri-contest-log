//
//  Generated file. Do not edit.
//

// clang-format off

#include "generated_plugin_registrant.h"

#include <gui_backend/gui_backend_plugin.h>

void fl_register_plugins(FlPluginRegistry* registry) {
  g_autoptr(FlPluginRegistrar) gui_backend_registrar =
      fl_plugin_registry_get_registrar_for_plugin(registry, "GuiBackendPlugin");
  gui_backend_plugin_register_with_registrar(gui_backend_registrar);
}
