///
//  Generated code. Do not modify.
//  source: proto/radio.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'radio.pb.dart' as $4;
import 'mcl.pb.dart' as $0;
import '../google/protobuf/empty.pb.dart' as $1;
export 'radio.pb.dart';

class RadioClient extends $grpc.Client {
  static final _$getRadioMode =
      $grpc.ClientMethod<$4.RadioSelector, $4.RadioStatus>(
          '/mcl.Radio/GetRadioMode',
          ($4.RadioSelector value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $4.RadioStatus.fromBuffer(value));
  static final _$pollRadioMode =
      $grpc.ClientMethod<$4.RadioSelector, $4.RadioStatus>(
          '/mcl.Radio/PollRadioMode',
          ($4.RadioSelector value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $4.RadioStatus.fromBuffer(value));
  static final _$radioOp =
      $grpc.ClientMethod<$4.RadioCommand, $0.StandardResponse>(
          '/mcl.Radio/RadioOp',
          ($4.RadioCommand value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));
  static final _$listRadioStatus =
      $grpc.ClientMethod<$1.Empty, $4.ActiveRadioList>(
          '/mcl.Radio/ListRadioStatus',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.ActiveRadioList.fromBuffer(value));
  static final _$listAudioDevices =
      $grpc.ClientMethod<$1.Empty, $4.AudioDeviceList>(
          '/mcl.Radio/ListAudioDevices',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.AudioDeviceList.fromBuffer(value));
  static final _$listSupportedRadios =
      $grpc.ClientMethod<$1.Empty, $4.SupportedRadioList>(
          '/mcl.Radio/ListSupportedRadios',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.SupportedRadioList.fromBuffer(value));
  static final _$setupRadio =
      $grpc.ClientMethod<$4.RadioConfig, $0.StandardResponse>(
          '/mcl.Radio/SetupRadio',
          ($4.RadioConfig value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));
  static final _$shutdownRadio =
      $grpc.ClientMethod<$4.RadioConfig, $0.StandardResponse>(
          '/mcl.Radio/ShutdownRadio',
          ($4.RadioConfig value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));

  RadioClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$4.RadioStatus> getRadioMode($4.RadioSelector request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getRadioMode, request, options: options);
  }

  $grpc.ResponseStream<$4.RadioStatus> pollRadioMode($4.RadioSelector request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$pollRadioMode, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> radioOp($4.RadioCommand request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$radioOp, request, options: options);
  }

  $grpc.ResponseFuture<$4.ActiveRadioList> listRadioStatus($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listRadioStatus, request, options: options);
  }

  $grpc.ResponseFuture<$4.AudioDeviceList> listAudioDevices($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listAudioDevices, request, options: options);
  }

  $grpc.ResponseFuture<$4.SupportedRadioList> listSupportedRadios(
      $1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listSupportedRadios, request, options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> setupRadio($4.RadioConfig request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$setupRadio, request, options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> shutdownRadio(
      $4.RadioConfig request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$shutdownRadio, request, options: options);
  }
}

abstract class RadioServiceBase extends $grpc.Service {
  $core.String get $name => 'mcl.Radio';

  RadioServiceBase() {
    $addMethod($grpc.ServiceMethod<$4.RadioSelector, $4.RadioStatus>(
        'GetRadioMode',
        getRadioMode_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.RadioSelector.fromBuffer(value),
        ($4.RadioStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.RadioSelector, $4.RadioStatus>(
        'PollRadioMode',
        pollRadioMode_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $4.RadioSelector.fromBuffer(value),
        ($4.RadioStatus value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.RadioCommand, $0.StandardResponse>(
        'RadioOp',
        radioOp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.RadioCommand.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $4.ActiveRadioList>(
        'ListRadioStatus',
        listRadioStatus_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($4.ActiveRadioList value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $4.AudioDeviceList>(
        'ListAudioDevices',
        listAudioDevices_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($4.AudioDeviceList value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $4.SupportedRadioList>(
        'ListSupportedRadios',
        listSupportedRadios_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($4.SupportedRadioList value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.RadioConfig, $0.StandardResponse>(
        'SetupRadio',
        setupRadio_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.RadioConfig.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.RadioConfig, $0.StandardResponse>(
        'ShutdownRadio',
        shutdownRadio_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.RadioConfig.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
  }

  $async.Future<$4.RadioStatus> getRadioMode_Pre(
      $grpc.ServiceCall call, $async.Future<$4.RadioSelector> request) async {
    return getRadioMode(call, await request);
  }

  $async.Stream<$4.RadioStatus> pollRadioMode_Pre(
      $grpc.ServiceCall call, $async.Future<$4.RadioSelector> request) async* {
    yield* pollRadioMode(call, await request);
  }

  $async.Future<$0.StandardResponse> radioOp_Pre(
      $grpc.ServiceCall call, $async.Future<$4.RadioCommand> request) async {
    return radioOp(call, await request);
  }

  $async.Future<$4.ActiveRadioList> listRadioStatus_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return listRadioStatus(call, await request);
  }

  $async.Future<$4.AudioDeviceList> listAudioDevices_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return listAudioDevices(call, await request);
  }

  $async.Future<$4.SupportedRadioList> listSupportedRadios_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return listSupportedRadios(call, await request);
  }

  $async.Future<$0.StandardResponse> setupRadio_Pre(
      $grpc.ServiceCall call, $async.Future<$4.RadioConfig> request) async {
    return setupRadio(call, await request);
  }

  $async.Future<$0.StandardResponse> shutdownRadio_Pre(
      $grpc.ServiceCall call, $async.Future<$4.RadioConfig> request) async {
    return shutdownRadio(call, await request);
  }

  $async.Future<$4.RadioStatus> getRadioMode(
      $grpc.ServiceCall call, $4.RadioSelector request);
  $async.Stream<$4.RadioStatus> pollRadioMode(
      $grpc.ServiceCall call, $4.RadioSelector request);
  $async.Future<$0.StandardResponse> radioOp(
      $grpc.ServiceCall call, $4.RadioCommand request);
  $async.Future<$4.ActiveRadioList> listRadioStatus(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$4.AudioDeviceList> listAudioDevices(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$4.SupportedRadioList> listSupportedRadios(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$0.StandardResponse> setupRadio(
      $grpc.ServiceCall call, $4.RadioConfig request);
  $async.Future<$0.StandardResponse> shutdownRadio(
      $grpc.ServiceCall call, $4.RadioConfig request);
}
