import 'package:flutter_test/flutter_test.dart';
import 'package:gui_backend/gui_backend.dart';
import 'package:gui_backend/gui_backend_platform_interface.dart';
import 'package:gui_backend/gui_backend_method_channel.dart';
import 'package:plugin_platform_interface/plugin_platform_interface.dart';

class MockGuiBackendPlatform
    with MockPlatformInterfaceMixin
    implements GuiBackendPlatform {

  @override
  Future<String?> getPlatformVersion() => Future.value('42');
}

void main() {
  final GuiBackendPlatform initialPlatform = GuiBackendPlatform.instance;

  test('$MethodChannelGuiBackend is the default instance', () {
    expect(initialPlatform, isInstanceOf<MethodChannelGuiBackend>());
  });

  test('getPlatformVersion', () async {
    GuiBackend guiBackendPlugin = GuiBackend();
    MockGuiBackendPlatform fakePlatform = MockGuiBackendPlatform();
    GuiBackendPlatform.instance = fakePlatform;

    expect(await guiBackendPlugin.getPlatformVersion(), '42');
  });
}
