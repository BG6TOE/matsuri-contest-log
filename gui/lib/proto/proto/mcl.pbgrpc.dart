///
//  Generated code. Do not modify.
//  source: proto/mcl.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'mcl.pb.dart' as $0;
import '../google/protobuf/empty.pb.dart' as $1;
export 'mcl.pb.dart';

class BinlogClient extends $grpc.Client {
  static final _$createContest =
      $grpc.ClientMethod<$0.CreateContestRequest, $0.StandardResponse>(
          '/mcl.Binlog/CreateContest',
          ($0.CreateContestRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));
  static final _$loadContest =
      $grpc.ClientMethod<$0.LoadContestRequest, $0.StandardResponse>(
          '/mcl.Binlog/LoadContest',
          ($0.LoadContestRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));
  static final _$getActiveContest =
      $grpc.ClientMethod<$1.Empty, $0.ActiveContest>(
          '/mcl.Binlog/GetActiveContest',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.ActiveContest.fromBuffer(value));
  static final _$push =
      $grpc.ClientMethod<$0.BinlogMessageSet, $0.StandardResponse>(
          '/mcl.Binlog/Push',
          ($0.BinlogMessageSet value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.StandardResponse.fromBuffer(value));
  static final _$retrieve =
      $grpc.ClientMethod<$0.RetrieveBinlogRequest, $0.BinlogMessageSet>(
          '/mcl.Binlog/Retrieve',
          ($0.RetrieveBinlogRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.BinlogMessageSet.fromBuffer(value));
  static final _$retrieveSnapshot =
      $grpc.ClientMethod<$0.RetrieveBinlogRequest, $0.SnapshotMessage>(
          '/mcl.Binlog/RetrieveSnapshot',
          ($0.RetrieveBinlogRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.SnapshotMessage.fromBuffer(value));

  BinlogClient($grpc.ClientChannel channel,
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

  $grpc.ResponseFuture<$0.ActiveContest> getActiveContest($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActiveContest, request, options: options);
  }

  $grpc.ResponseFuture<$0.StandardResponse> push($0.BinlogMessageSet request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$push, request, options: options);
  }

  $grpc.ResponseFuture<$0.BinlogMessageSet> retrieve(
      $0.RetrieveBinlogRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$retrieve, request, options: options);
  }

  $grpc.ResponseFuture<$0.SnapshotMessage> retrieveSnapshot(
      $0.RetrieveBinlogRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$retrieveSnapshot, request, options: options);
  }
}

abstract class BinlogServiceBase extends $grpc.Service {
  $core.String get $name => 'mcl.Binlog';

  BinlogServiceBase() {
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
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.ActiveContest>(
        'GetActiveContest',
        getActiveContest_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.ActiveContest value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.BinlogMessageSet, $0.StandardResponse>(
        'Push',
        push_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.BinlogMessageSet.fromBuffer(value),
        ($0.StandardResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.RetrieveBinlogRequest, $0.BinlogMessageSet>(
            'Retrieve',
            retrieve_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.RetrieveBinlogRequest.fromBuffer(value),
            ($0.BinlogMessageSet value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.RetrieveBinlogRequest, $0.SnapshotMessage>(
            'RetrieveSnapshot',
            retrieveSnapshot_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.RetrieveBinlogRequest.fromBuffer(value),
            ($0.SnapshotMessage value) => value.writeToBuffer()));
  }

  $async.Future<$0.StandardResponse> createContest_Pre($grpc.ServiceCall call,
      $async.Future<$0.CreateContestRequest> request) async {
    return createContest(call, await request);
  }

  $async.Future<$0.StandardResponse> loadContest_Pre($grpc.ServiceCall call,
      $async.Future<$0.LoadContestRequest> request) async {
    return loadContest(call, await request);
  }

  $async.Future<$0.ActiveContest> getActiveContest_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getActiveContest(call, await request);
  }

  $async.Future<$0.StandardResponse> push_Pre($grpc.ServiceCall call,
      $async.Future<$0.BinlogMessageSet> request) async {
    return push(call, await request);
  }

  $async.Future<$0.BinlogMessageSet> retrieve_Pre($grpc.ServiceCall call,
      $async.Future<$0.RetrieveBinlogRequest> request) async {
    return retrieve(call, await request);
  }

  $async.Future<$0.SnapshotMessage> retrieveSnapshot_Pre($grpc.ServiceCall call,
      $async.Future<$0.RetrieveBinlogRequest> request) async {
    return retrieveSnapshot(call, await request);
  }

  $async.Future<$0.StandardResponse> createContest(
      $grpc.ServiceCall call, $0.CreateContestRequest request);
  $async.Future<$0.StandardResponse> loadContest(
      $grpc.ServiceCall call, $0.LoadContestRequest request);
  $async.Future<$0.ActiveContest> getActiveContest(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$0.StandardResponse> push(
      $grpc.ServiceCall call, $0.BinlogMessageSet request);
  $async.Future<$0.BinlogMessageSet> retrieve(
      $grpc.ServiceCall call, $0.RetrieveBinlogRequest request);
  $async.Future<$0.SnapshotMessage> retrieveSnapshot(
      $grpc.ServiceCall call, $0.RetrieveBinlogRequest request);
}
