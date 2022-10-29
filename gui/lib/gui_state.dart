import 'dart:async';

import 'package:flutter/foundation.dart';
import 'package:flutter/services.dart';
import 'package:grpc/grpc.dart';
import 'package:gui_lib/gui_lib.dart';
import 'proto/google/protobuf/empty.pb.dart';
import 'proto/proto/mcl.pbgrpc.dart';
import 'proto/proto/mclgui.pbgrpc.dart';
import 'proto/proto/radio.pbgrpc.dart';

const platform = MethodChannel('mcl.matsu.dev/embed');

enum CallbackKind { kOnQsoUpdate, kOnQsoFieldSetUpdate, kRadioStatusUpdate }

class GuiState {
  GuiClient? _guiClient;
  RadioClient? _radioClient;
  RealtimeGuiClient? _realtimeGui;
  final bool _guiServerConnected = false;

  ActiveContest? activeContest;

  var displayExchangeSentFields = <QSOField>[];
  var displayExchangeRcvdFields = <QSOField>[];

  int callbackIdentifier = 0;

  Map<CallbackKind, Map<int, void Function()>> callbacks = {};

  List<QSO> activeQsos = [];
  Map<String, RadioStatus> radios = {};

  GuiState() {
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
    runMclGuiServer("tcp://127.0.0.1:62122");

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
    _realtimeGui = RealtimeGuiClient(channel);
    _radioClient = RadioClient(channel);

    Timer.periodic(const Duration(seconds: 1), (timer) {
      refreshRadios();
    });
  }

  GuiClient? getGuiClient() {
    return _guiClient;
  }

  RealtimeGuiClient get realtimeGui => _realtimeGui!;
  RadioClient get radioClient => _radioClient!;

  Future<void> updateExchangeFields() async {
    var fields = await _guiClient!.getQSOFields(Empty());

    if (!listEquals(displayExchangeSentFields, fields.exchangeSent) ||
        !listEquals(displayExchangeRcvdFields, fields.exchangeRcvd)) {
      displayExchangeSentFields = fields.exchangeSent;
      displayExchangeRcvdFields = fields.exchangeRcvd;

      invokeCallback(CallbackKind.kOnQsoFieldSetUpdate);
    }
  }

  Future<void> refreshRadios() async {
    if (_radioClient == null) {
      return;
    }

    radios = (await _radioClient!.listRadioStatus(Empty())).radios;

    invokeCallback(CallbackKind.kRadioStatusUpdate);
  }

  Future<void> refreshQsos() async {
    if (_guiClient == null) {
      return;
    }

    var qsos = await _guiClient!.getActiveQSOs(Empty());

    var loadedQsos = qsos.qso;
    loadedQsos.sort((qso1, qso2) => qso1.time.compareTo(qso2.time));

    activeQsos = loadedQsos;

    activeContest = await _guiClient!.getActiveContest(Empty());

    await updateExchangeFields();
    invokeCallback(CallbackKind.kOnQsoUpdate);
  }

  bool guiServerConnected() {
    return _guiClient != null;
  }
}

final GuiState state = GuiState();
