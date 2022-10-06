import 'package:plugin_platform_interface/plugin_platform_interface.dart';

import 'gui_backend_method_channel.dart';

abstract class GuiBackendPlatform extends PlatformInterface {
  /// Constructs a GuiBackendPlatform.
  GuiBackendPlatform() : super(token: _token);

  static final Object _token = Object();

  static GuiBackendPlatform _instance = MethodChannelGuiBackend();

  /// The default instance of [GuiBackendPlatform] to use.
  ///
  /// Defaults to [MethodChannelGuiBackend].
  static GuiBackendPlatform get instance => _instance;

  /// Platform-specific implementations should set this with their own
  /// platform-specific class that extends [GuiBackendPlatform] when
  /// they register themselves.
  static set instance(GuiBackendPlatform instance) {
    PlatformInterface.verifyToken(instance, _token);
    _instance = instance;
  }

  Future<String?> getPlatformVersion() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }

  Future<int?> runServer() {
    throw UnimplementedError('runServer() has not been implemented.');
  }
}
