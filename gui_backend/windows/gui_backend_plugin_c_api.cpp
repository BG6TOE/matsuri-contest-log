#include "include/gui_backend/gui_backend_plugin_c_api.h"

#include <flutter/plugin_registrar_windows.h>

#include "gui_backend_plugin.h"

void GuiBackendPluginCApiRegisterWithRegistrar(
    FlutterDesktopPluginRegistrarRef registrar) {
  gui_backend::GuiBackendPlugin::RegisterWithRegistrar(
      flutter::PluginRegistrarManager::GetInstance()
          ->GetRegistrar<flutter::PluginRegistrarWindows>(registrar));
}
