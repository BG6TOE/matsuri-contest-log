#ifndef FLUTTER_PLUGIN_GUI_BACKEND_PLUGIN_H_
#define FLUTTER_PLUGIN_GUI_BACKEND_PLUGIN_H_

#include <flutter/method_channel.h>
#include <flutter/plugin_registrar_windows.h>

#include <memory>

namespace gui_backend {

class GuiBackendPlugin : public flutter::Plugin {
 public:
  static void RegisterWithRegistrar(flutter::PluginRegistrarWindows *registrar);

  GuiBackendPlugin();

  virtual ~GuiBackendPlugin();

  // Disallow copy and assign.
  GuiBackendPlugin(const GuiBackendPlugin&) = delete;
  GuiBackendPlugin& operator=(const GuiBackendPlugin&) = delete;

 private:
  // Called when a method is called on this plugin's channel from Dart.
  void HandleMethodCall(
      const flutter::MethodCall<flutter::EncodableValue> &method_call,
      std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result);
};

}  // namespace gui_backend

#endif  // FLUTTER_PLUGIN_GUI_BACKEND_PLUGIN_H_
