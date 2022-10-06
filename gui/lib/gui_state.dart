import 'dart:async';

import 'package:flutter/services.dart';
import 'package:grpc/grpc.dart';
import 'package:gui_backend/gui_backend.dart';
import 'proto/google/protobuf/empty.pb.dart';
import 'proto/proto/mcl.pbgrpc.dart';
import 'proto/proto/mclgui.pbgrpc.dart';

const platform = MethodChannel('mcl.matsu.dev/embed');

class RadioState {}

enum CallbackKind { kOnQsoUpdate }

class GuiState {
  final RadioState _radioState;
  GuiClient? _guiClient;
  final bool _guiServerConnected = false;

  var exchangeSentFields = <String>[];
  var exchangeRcvdFields = <String>[];

  int callbackIdentifier = 0;

  Map<CallbackKind, Map<int, void Function()>> callbacks = {};
  List<QSO> activeQsos = [];

  GuiState() : _radioState = RadioState() {
    exchangeSentFields = <String>['RST Snt'];
    exchangeRcvdFields = <String>['RST Rcv', 'Exch Rcv'];

    for (var val in CallbackKind.values) {
      callbacks[val] = <int, void Function()>{};
    }
  }

  void startServer() {
    unawaited(_startGuiServer.call());
  }

  void _invoke(int key, void Function() all) {
    all.call();
  }

  void invokeCallback(CallbackKind kind) {
    callbacks[kind]?.forEach(_invoke);
  }

  int registerCallback(CallbackKind kind, void Function() callback) {
    callbackIdentifier++;
    callbacks[kind]![callbackIdentifier] = callback;
    return callbackIdentifier;
  }

  void deregisterCallback(CallbackKind kind, int identifier) {
    callbacks[kind]!.remove(identifier);
  }

  Future<void> _startGuiServer() async {
    await GuiBackend().runServer();

    final uri = Uri.parse("tcp://127.0.0.1:62122");
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

  void refreshQsos() async {
    if (_guiClient == null) {
      return;
    }

    var qsos = await _guiClient!.getActiveQSOs(Empty());

    var loadedQsos = qsos.qso;
    loadedQsos.sort((qso1, qso2) => qso1.time.compareTo(qso2.time));

    activeQsos = loadedQsos;

    invokeCallback(CallbackKind.kOnQsoUpdate);
  }

  bool guiServerConnected() {
    return _guiClient != null;
  }
}

final GuiState state = GuiState();
