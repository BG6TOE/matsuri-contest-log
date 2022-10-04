import 'dart:async';
import 'dart:developer';

import 'package:flutter/services.dart';
import 'package:grpc/grpc.dart';
import 'package:mcl_gui/proto/proto/mclgui.pbgrpc.dart';

const platform = MethodChannel('mcl.matsu.dev/embed');

class RadioState {}

class GuiState {
  final RadioState _radioState;
  GuiClient? _guiClient;
  final bool _guiServerConnected = false;

  var exchangeSentFields = <String>[];
  var exchangeRcvdFields = <String>[];

  GuiState() : _radioState = RadioState() {
    exchangeSentFields = <String>['RST Snt'];
    exchangeRcvdFields = <String>['RST Rcv', 'Exch Rcv'];

    unawaited(_startGuiServer.call());
  }

  Future<void> _startGuiServer() async {
    await platform.invokeMethod('start', "database.sqlite");

    final uri = Uri.parse("tcp://127.0.0.1:62122");

    print(uri.host);
    print(uri.port);

    final channel = ClientChannel(
      uri.host,
      port: uri.port,
      options: ChannelOptions(
        credentials: const ChannelCredentials.insecure(),
        codecRegistry:
            CodecRegistry(codecs: const [GzipCodec(), IdentityCodec()]),
      ),
    );

    _guiClient = GuiClient(channel);
  }

  GuiClient? getGuiClient() {
    return _guiClient;
  }

  bool guiServerConnected() {
    return _guiClient != null;
  }
}

final GuiState state = GuiState();
