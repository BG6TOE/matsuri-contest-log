///
//  Generated code. Do not modify.
//  source: proto/mclgui.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'mcl.pb.dart' as $0;
import 'mclgui.pb.dart' as $1;
import '../google/protobuf/empty.pb.dart' as $2;
export 'mclgui.pb.dart';

class GuiClient extends $grpc.Client {
  static final _$createContest =
      $grpc.ClientMethod<$0.CreateContestRequest, $0.StandardResponse>(
          '/mcl.Gui/CreateContest',
          ($0.CreateContestRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));
  static final _$loadContest =
      $grpc.ClientMethod<$0.LoadContestRequest, $0.StandardResponse>(
          '/mcl.Gui/LoadContest',
          ($0.LoadContestRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));
  static final _$logQSO = $grpc.ClientMethod<$1.QSOMessage, $0.QSO>(
      '/mcl.Gui/LogQSO',
      ($1.QSOMessage value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.QSO.fromBuffer(value));
  static final _$getActiveQSOs =
      $grpc.ClientMethod<$2.Empty, $0.SnapshotMessage>(
          '/mcl.Gui/GetActiveQSOs',
          ($2.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.SnapshotMessage.fromBuffer(value));
  static final _$deleteQSO = $grpc.ClientMethod<$0.QSO, $0.StandardResponse>(
      '/mcl.Gui/DeleteQSO',
      ($0.QSO value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.StandardResponse.fromBuffer(value));
  static final _$getScore = $grpc.ClientMethod<$2.Empty, $1.ScoreResponse>(
      '/mcl.Gui/GetScore',
      ($2.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.ScoreResponse.fromBuffer(value));

  GuiClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.StandardResponse> createContest(
      $0.CreateContestRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createContest, request, options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> loadContest(
      $0.LoadContestRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$loadContest, request, options: options);
  }

  $grpc.ResponseFuture<$0.QSO> logQSO($1.QSOMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$logQSO, request, options: options);
  }

  $grpc.ResponseFuture<$0.SnapshotMessage> getActiveQSOs($2.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActiveQSOs, request, options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> deleteQSO($0.QSO request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteQSO, request, options: options);
  }

  $grpc.ResponseFuture<$1.ScoreResponse> getScore($2.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getScore, request, options: options);
  }
}

abstract class GuiServiceBase extends $grpc.Service {
  $core.String get $name => 'mcl.Gui';

  GuiServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.CreateContestRequest, $0.StandardResponse>(
            'CreateContest',
            createContest_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.CreateContestRequest.fromBuffer(value),
            ($0.StandardResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.LoadContestRequest, $0.StandardResponse>(
        'LoadContest',
        loadContest_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.LoadContestRequest.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.QSOMessage, $0.QSO>(
        'LogQSO',
        logQSO_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.QSOMessage.fromBuffer(value),
        ($0.QSO value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Empty, $0.SnapshotMessage>(
        'GetActiveQSOs',
        getActiveQSOs_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Empty.fromBuffer(value),
        ($0.SnapshotMessage value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QSO, $0.StandardResponse>(
        'DeleteQSO',
        deleteQSO_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.QSO.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Empty, $1.ScoreResponse>(
        'GetScore',
        getScore_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Empty.fromBuffer(value),
        ($1.ScoreResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.StandardResponse> createContest_Pre($grpc.ServiceCall call,
      $async.Future<$0.CreateContestRequest> request) async {
    return createContest(call, await request);
  }

  $async.Future<$0.StandardResponse> loadContest_Pre($grpc.ServiceCall call,
      $async.Future<$0.LoadContestRequest> request) async {
    return loadContest(call, await request);
  }

  $async.Future<$0.QSO> logQSO_Pre(
      $grpc.ServiceCall call, $async.Future<$1.QSOMessage> request) async {
    return logQSO(call, await request);
  }

  $async.Future<$0.SnapshotMessage> getActiveQSOs_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Empty> request) async {
    return getActiveQSOs(call, await request);
  }

  $async.Future<$0.StandardResponse> deleteQSO_Pre(
      $grpc.ServiceCall call, $async.Future<$0.QSO> request) async {
    return deleteQSO(call, await request);
  }

  $async.Future<$1.ScoreResponse> getScore_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Empty> request) async {
    return getScore(call, await request);
  }

  $async.Future<$0.StandardResponse> createContest(
      $grpc.ServiceCall call, $0.CreateContestRequest request);
  $async.Future<$0.StandardResponse> loadContest(
      $grpc.ServiceCall call, $0.LoadContestRequest request);
  $async.Future<$0.QSO> logQSO($grpc.ServiceCall call, $1.QSOMessage request);
  $async.Future<$0.SnapshotMessage> getActiveQSOs(
      $grpc.ServiceCall call, $2.Empty request);
  $async.Future<$0.StandardResponse> deleteQSO(
      $grpc.ServiceCall call, $0.QSO request);
  $async.Future<$1.ScoreResponse> getScore(
      $grpc.ServiceCall call, $2.Empty request);
}

class RealtimeGuiServerClient extends $grpc.Client {
  static final _$draftQSO =
      $grpc.ClientMethod<$1.DraftQSOMessage, $1.QSOMessage>(
          '/mcl.RealtimeGuiServer/DraftQSO',
          ($1.DraftQSOMessage value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.QSOMessage.fromBuffer(value));
  static final _$retrieveQSOUpdates =
      $grpc.ClientMethod<$2.Empty, $0.BinlogMessage>(
          '/mcl.RealtimeGuiServer/RetrieveQSOUpdates',
          ($2.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.BinlogMessage.fromBuffer(value));
  static final _$retrieveTelnet = $grpc.ClientMethod<$2.Empty, $1.Spot>(
      '/mcl.RealtimeGuiServer/RetrieveTelnet',
      ($2.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Spot.fromBuffer(value));
  static final _$sendSpotToTelnet =
      $grpc.ClientMethod<$1.Spot, $0.StandardResponse>(
          '/mcl.RealtimeGuiServer/SendSpotToTelnet',
          ($1.Spot value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));

  RealtimeGuiServerClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.QSOMessage> draftQSO($1.DraftQSOMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$draftQSO, request, options: options);
  }

  $grpc.ResponseStream<$0.BinlogMessage> retrieveQSOUpdates($2.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$retrieveQSOUpdates, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseStream<$1.Spot> retrieveTelnet($2.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$retrieveTelnet, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> sendSpotToTelnet($1.Spot request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$sendSpotToTelnet, request, options: options);
  }
}

abstract class RealtimeGuiServerServiceBase extends $grpc.Service {
  $core.String get $name => 'mcl.RealtimeGuiServer';

  RealtimeGuiServerServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.DraftQSOMessage, $1.QSOMessage>(
        'DraftQSO',
        draftQSO_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.DraftQSOMessage.fromBuffer(value),
        ($1.QSOMessage value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Empty, $0.BinlogMessage>(
        'RetrieveQSOUpdates',
        retrieveQSOUpdates_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $2.Empty.fromBuffer(value),
        ($0.BinlogMessage value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Empty, $1.Spot>(
        'RetrieveTelnet',
        retrieveTelnet_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $2.Empty.fromBuffer(value),
        ($1.Spot value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Spot, $0.StandardResponse>(
        'SendSpotToTelnet',
        sendSpotToTelnet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Spot.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.QSOMessage> draftQSO_Pre(
      $grpc.ServiceCall call, $async.Future<$1.DraftQSOMessage> request) async {
    return draftQSO(call, await request);
  }

  $async.Stream<$0.BinlogMessage> retrieveQSOUpdates_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Empty> request) async* {
    yield* retrieveQSOUpdates(call, await request);
  }

  $async.Stream<$1.Spot> retrieveTelnet_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Empty> request) async* {
    yield* retrieveTelnet(call, await request);
  }

  $async.Future<$0.StandardResponse> sendSpotToTelnet_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Spot> request) async {
    return sendSpotToTelnet(call, await request);
  }

  $async.Future<$1.QSOMessage> draftQSO(
      $grpc.ServiceCall call, $1.DraftQSOMessage request);
  $async.Stream<$0.BinlogMessage> retrieveQSOUpdates(
      $grpc.ServiceCall call, $2.Empty request);
  $async.Stream<$1.Spot> retrieveTelnet(
      $grpc.ServiceCall call, $2.Empty request);
  $async.Future<$0.StandardResponse> sendSpotToTelnet(
      $grpc.ServiceCall call, $1.Spot request);
}

class RadioClient extends $grpc.Client {
  static final _$getRadioMode = $grpc.ClientMethod<$2.Empty, $1.RadioStatus>(
      '/mcl.Radio/GetRadioMode',
      ($2.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.RadioStatus.fromBuffer(value));
  static final _$pollRadioMode = $grpc.ClientMethod<$2.Empty, $1.RadioStatus>(
      '/mcl.Radio/PollRadioMode',
      ($2.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.RadioStatus.fromBuffer(value));
  static final _$radioOp =
      $grpc.ClientMethod<$1.RadioCommands, $0.StandardResponse>(
          '/mcl.Radio/RadioOp',
          ($1.RadioCommands value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));

  RadioClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.RadioStatus> getRadioMode($2.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getRadioMode, request, options: options);
  }

  $grpc.ResponseStream<$1.RadioStatus> pollRadioMode($2.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$pollRadioMode, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> radioOp($1.RadioCommands request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$radioOp, request, options: options);
  }
}

abstract class RadioServiceBase extends $grpc.Service {
  $core.String get $name => 'mcl.Radio';

  RadioServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.Empty, $1.RadioStatus>(
        'GetRadioMode',
        getRadioMode_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Empty.fromBuffer(value),
        ($1.RadioStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Empty, $1.RadioStatus>(
        'PollRadioMode',
        pollRadioMode_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $2.Empty.fromBuffer(value),
        ($1.RadioStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.RadioCommands, $0.StandardResponse>(
        'RadioOp',
        radioOp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.RadioCommands.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.RadioStatus> getRadioMode_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Empty> request) async {
    return getRadioMode(call, await request);
  }

  $async.Stream<$1.RadioStatus> pollRadioMode_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Empty> request) async* {
    yield* pollRadioMode(call, await request);
  }

  $async.Future<$0.StandardResponse> radioOp_Pre(
      $grpc.ServiceCall call, $async.Future<$1.RadioCommands> request) async {
    return radioOp(call, await request);
  }

  $async.Future<$1.RadioStatus> getRadioMode(
      $grpc.ServiceCall call, $2.Empty request);
  $async.Stream<$1.RadioStatus> pollRadioMode(
      $grpc.ServiceCall call, $2.Empty request);
  $async.Future<$0.StandardResponse> radioOp(
      $grpc.ServiceCall call, $1.RadioCommands request);
}
