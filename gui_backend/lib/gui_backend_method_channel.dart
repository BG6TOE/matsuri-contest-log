
import 'package:flutter/foundation.dart';
import 'package:flutter/services.dart';

import 'gui_backend_platform_interface.dart';

/// An implementation of [GuiBackendPlatform] that uses method channels.
class MethodChannelGuiBackend extends GuiBackendPlatform {
  /// The method channel used to interact with the native platform.
  @visibleForTesting
  final methodChannel = const MethodChannel('gui_backend');

  @override
  Future<String?> getPlatformVersion() async {
    final version = await methodChannel.invokeMethod<String>('getPlatformVersion');
    return version;
  }

  @override
  Future<int?> runServer() async {
    final res = await methodChannel.invokeMethod<int>('runServer', 'tcp://127.0.0.1:62122');
    return res;
  }
}
