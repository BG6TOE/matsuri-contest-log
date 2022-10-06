
import 'gui_backend_platform_interface.dart';

class GuiBackend {
  Future<String?> getPlatformVersion() {
    return GuiBackendPlatform.instance.getPlatformVersion();
  }

  Future<int?> runServer() {
    return GuiBackendPlatform.instance.runServer();
  }
}
