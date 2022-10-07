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
import '../google/protobuf/empty.pb.dart' as $1;
import 'mclgui.pb.dart' as $2;
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
  static final _$parseContest =
      $grpc.ClientMethod<$0.ParseContestRequest, $0.Contest>(
          '/mcl.Gui/ParseContest',
          ($0.ParseContestRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Contest.fromBuffer(value));
  static final _$getActiveContest =
      $grpc.ClientMethod<$1.Empty, $0.ActiveContest>(
          '/mcl.Gui/GetActiveContest',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.ActiveContest.fromBuffer(value));
  static final _$getQSOFields = $grpc.ClientMethod<$1.Empty, $2.QSOFields>(
      '/mcl.Gui/GetQSOFields',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.QSOFields.fromBuffer(value));
  static final _$logQSO = $grpc.ClientMethod<$2.QSOMessage, $0.QSO>(
      '/mcl.Gui/LogQSO',
      ($2.QSOMessage value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.QSO.fromBuffer(value));
  static final _$getActiveQSOs =
      $grpc.ClientMethod<$1.Empty, $0.SnapshotMessage>(
          '/mcl.Gui/GetActiveQSOs',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.SnapshotMessage.fromBuffer(value));
  static final _$deleteQSO = $grpc.ClientMethod<$0.QSO, $0.StandardResponse>(
      '/mcl.Gui/DeleteQSO',
      ($0.QSO value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.StandardResponse.fromBuffer(value));
  static final _$getScore = $grpc.ClientMethod<$1.Empty, $2.ScoreResponse>(
      '/mcl.Gui/GetScore',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.ScoreResponse.fromBuffer(value));

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

  $grpc.ResponseFuture<$0.Contest> parseContest($0.ParseContestRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$parseContest, request, options: options);
  }

  $grpc.ResponseFuture<$0.ActiveContest> getActiveContest($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActiveContest, request, options: options);
  }

  $grpc.ResponseFuture<$2.QSOFields> getQSOFields($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getQSOFields, request, options: options);
  }

  $grpc.ResponseFuture<$0.QSO> logQSO($2.QSOMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$logQSO, request, options: options);
  }

  $grpc.ResponseFuture<$0.SnapshotMessage> getActiveQSOs($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActiveQSOs, request, options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> deleteQSO($0.QSO request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteQSO, request, options: options);
  }

  $grpc.ResponseFuture<$2.ScoreResponse> getScore($1.Empty request,
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
    $addMethod($grpc.ServiceMethod<$0.ParseContestRequest, $0.Contest>(
        'ParseContest',
        parseContest_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.ParseContestRequest.fromBuffer(value),
        ($0.Contest value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.ActiveContest>(
        'GetActiveContest',
        getActiveContest_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.ActiveContest value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $2.QSOFields>(
        'GetQSOFields',
        getQSOFields_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($2.QSOFields value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.QSOMessage, $0.QSO>(
        'LogQSO',
        logQSO_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.QSOMessage.fromBuffer(value),
        ($0.QSO value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.SnapshotMessage>(
        'GetActiveQSOs',
        getActiveQSOs_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.SnapshotMessage value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QSO, $0.StandardResponse>(
        'DeleteQSO',
        deleteQSO_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.QSO.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $2.ScoreResponse>(
        'GetScore',
        getScore_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($2.ScoreResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.StandardResponse> createContest_Pre($grpc.ServiceCall call,
      $async.Future<$0.CreateContestRequest> request) async {
    return createContest(call, await request);
  }

  $async.Future<$0.StandardResponse> loadContest_Pre($grpc.ServiceCall call,
      $async.Future<$0.LoadContestRequest> request) async {
    return loadContest(call, await request);
  }

  $async.Future<$0.Contest> parseContest_Pre($grpc.ServiceCall call,
      $async.Future<$0.ParseContestRequest> request) async {
    return parseContest(call, await request);
  }

  $async.Future<$0.ActiveContest> getActiveContest_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getActiveContest(call, await request);
  }

  $async.Future<$2.QSOFields> getQSOFields_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getQSOFields(call, await request);
  }

  $async.Future<$0.QSO> logQSO_Pre(
      $grpc.ServiceCall call, $async.Future<$2.QSOMessage> request) async {
    return logQSO(call, await request);
  }

  $async.Future<$0.SnapshotMessage> getActiveQSOs_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getActiveQSOs(call, await request);
  }

  $async.Future<$0.StandardResponse> deleteQSO_Pre(
      $grpc.ServiceCall call, $async.Future<$0.QSO> request) async {
    return deleteQSO(call, await request);
  }

  $async.Future<$2.ScoreResponse> getScore_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getScore(call, await request);
  }

  $async.Future<$0.StandardResponse> createContest(
      $grpc.ServiceCall call, $0.CreateContestRequest request);
  $async.Future<$0.StandardResponse> loadContest(
      $grpc.ServiceCall call, $0.LoadContestRequest request);
  $async.Future<$0.Contest> parseContest(
      $grpc.ServiceCall call, $0.ParseContestRequest request);
  $async.Future<$0.ActiveContest> getActiveContest(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$2.QSOFields> getQSOFields(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$0.QSO> logQSO($grpc.ServiceCall call, $2.QSOMessage request);
  $async.Future<$0.SnapshotMessage> getActiveQSOs(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$0.StandardResponse> deleteQSO(
      $grpc.ServiceCall call, $0.QSO request);
  $async.Future<$2.ScoreResponse> getScore(
      $grpc.ServiceCall call, $1.Empty request);
}

class RealtimeGuiServerClient extends $grpc.Client {
  static final _$draftQSO =
      $grpc.ClientMethod<$2.DraftQSOMessage, $2.QSOMessage>(
          '/mcl.RealtimeGuiServer/DraftQSO',
          ($2.DraftQSOMessage value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.QSOMessage.fromBuffer(value));
  static final _$retrieveQSOUpdates =
      $grpc.ClientMethod<$1.Empty, $0.BinlogMessage>(
          '/mcl.RealtimeGuiServer/RetrieveQSOUpdates',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.BinlogMessage.fromBuffer(value));
  static final _$retrieveTelnet = $grpc.ClientMethod<$1.Empty, $2.Spot>(
      '/mcl.RealtimeGuiServer/RetrieveTelnet',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Spot.fromBuffer(value));
  static final _$sendSpotToTelnet =
      $grpc.ClientMethod<$2.Spot, $0.StandardResponse>(
          '/mcl.RealtimeGuiServer/SendSpotToTelnet',
          ($2.Spot value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));

  RealtimeGuiServerClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$2.QSOMessage> draftQSO($2.DraftQSOMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$draftQSO, request, options: options);
  }

  $grpc.ResponseStream<$0.BinlogMessage> retrieveQSOUpdates($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$retrieveQSOUpdates, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseStream<$2.Spot> retrieveTelnet($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$retrieveTelnet, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> sendSpotToTelnet($2.Spot request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$sendSpotToTelnet, request, options: options);
  }
}

abstract class RealtimeGuiServerServiceBase extends $grpc.Service {
  $core.String get $name => 'mcl.RealtimeGuiServer';

  RealtimeGuiServerServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.DraftQSOMessage, $2.QSOMessage>(
        'DraftQSO',
        draftQSO_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.DraftQSOMessage.fromBuffer(value),
        ($2.QSOMessage value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.BinlogMessage>(
        'RetrieveQSOUpdates',
        retrieveQSOUpdates_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.BinlogMessage value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $2.Spot>(
        'RetrieveTelnet',
        retrieveTelnet_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($2.Spot value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Spot, $0.StandardResponse>(
        'SendSpotToTelnet',
        sendSpotToTelnet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Spot.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
  }

  $async.Future<$2.QSOMessage> draftQSO_Pre(
      $grpc.ServiceCall call, $async.Future<$2.DraftQSOMessage> request) async {
    return draftQSO(call, await request);
  }

  $async.Stream<$0.BinlogMessage> retrieveQSOUpdates_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async* {
    yield* retrieveQSOUpdates(call, await request);
  }

  $async.Stream<$2.Spot> retrieveTelnet_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async* {
    yield* retrieveTelnet(call, await request);
  }

  $async.Future<$0.StandardResponse> sendSpotToTelnet_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Spot> request) async {
    return sendSpotToTelnet(call, await request);
  }

  $async.Future<$2.QSOMessage> draftQSO(
      $grpc.ServiceCall call, $2.DraftQSOMessage request);
  $async.Stream<$0.BinlogMessage> retrieveQSOUpdates(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Stream<$2.Spot> retrieveTelnet(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$0.StandardResponse> sendSpotToTelnet(
      $grpc.ServiceCall call, $2.Spot request);
}

class RadioClient extends $grpc.Client {
  static final _$getRadioMode = $grpc.ClientMethod<$1.Empty, $2.RadioStatus>(
      '/mcl.Radio/GetRadioMode',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.RadioStatus.fromBuffer(value));
  static final _$pollRadioMode = $grpc.ClientMethod<$1.Empty, $2.RadioStatus>(
      '/mcl.Radio/PollRadioMode',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.RadioStatus.fromBuffer(value));
  static final _$radioOp =
      $grpc.ClientMethod<$2.RadioCommands, $0.StandardResponse>(
          '/mcl.Radio/RadioOp',
          ($2.RadioCommands value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));

  RadioClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$2.RadioStatus> getRadioMode($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getRadioMode, request, options: options);
  }

  $grpc.ResponseStream<$2.RadioStatus> pollRadioMode($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$pollRadioMode, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> radioOp($2.RadioCommands request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$radioOp, request, options: options);
  }
}

abstract class RadioServiceBase extends $grpc.Service {
  $core.String get $name => 'mcl.Radio';

  RadioServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.Empty, $2.RadioStatus>(
        'GetRadioMode',
        getRadioMode_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($2.RadioStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $2.RadioStatus>(
        'PollRadioMode',
        pollRadioMode_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($2.RadioStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.RadioCommands, $0.StandardResponse>(
        'RadioOp',
        radioOp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.RadioCommands.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
  }

  $async.Future<$2.RadioStatus> getRadioMode_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getRadioMode(call, await request);
  }

  $async.Stream<$2.RadioStatus> pollRadioMode_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async* {
    yield* pollRadioMode(call, await request);
  }

  $async.Future<$0.StandardResponse> radioOp_Pre(
      $grpc.ServiceCall call, $async.Future<$2.RadioCommands> request) async {
    return radioOp(call, await request);
  }

  $async.Future<$2.RadioStatus> getRadioMode(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Stream<$2.RadioStatus> pollRadioMode(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$0.StandardResponse> radioOp(
      $grpc.ServiceCall call, $2.RadioCommands request);
}
