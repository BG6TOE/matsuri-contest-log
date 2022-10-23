///
//  Generated code. Do not modify.
//  source: proto/mcl.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'mcl.pb.dart' as $2;
import '../google/protobuf/empty.pb.dart' as $1;
import 'mcl.pbjson.dart';

export 'mcl.pb.dart';

abstract class BinlogServiceBase extends $pb.GeneratedService {
  $async.Future<$2.StandardResponse> createContest($pb.ServerContext ctx, $2.CreateContestRequest request);
  $async.Future<$2.StandardResponse> loadContest($pb.ServerContext ctx, $2.LoadContestRequest request);
  $async.Future<$2.ActiveContest> getActiveContest($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$2.StandardResponse> push($pb.ServerContext ctx, $2.BinlogMessageSet request);
  $async.Future<$2.BinlogMessageSet> retrieve($pb.ServerContext ctx, $2.RetrieveBinlogRequest request);
  $async.Future<$2.SnapshotMessage> retrieveSnapshot($pb.ServerContext ctx, $2.RetrieveBinlogRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateContest': return $2.CreateContestRequest();
      case 'LoadContest': return $2.LoadContestRequest();
      case 'GetActiveContest': return $1.Empty();
      case 'Push': return $2.BinlogMessageSet();
      case 'Retrieve': return $2.RetrieveBinlogRequest();
      case 'RetrieveSnapshot': return $2.RetrieveBinlogRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateContest': return this.createContest(ctx, request as $2.CreateContestRequest);
      case 'LoadContest': return this.loadContest(ctx, request as $2.LoadContestRequest);
      case 'GetActiveContest': return this.getActiveContest(ctx, request as $1.Empty);
      case 'Push': return this.push(ctx, request as $2.BinlogMessageSet);
      case 'Retrieve': return this.retrieve(ctx, request as $2.RetrieveBinlogRequest);
      case 'RetrieveSnapshot': return this.retrieveSnapshot(ctx, request as $2.RetrieveBinlogRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => BinlogServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => BinlogServiceBase$messageJson;
}

