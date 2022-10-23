///
//  Generated code. Do not modify.
//  source: proto/mclgui.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'mcl.pb.dart' as $2;
import 'common.pb.dart' as $0;
import '../google/protobuf/empty.pb.dart' as $1;
import 'mclgui.pb.dart' as $3;
import 'mclgui.pbjson.dart';

export 'mclgui.pb.dart';

abstract class GuiServiceBase extends $pb.GeneratedService {
  $async.Future<$2.StandardResponse> createContest($pb.ServerContext ctx, $2.CreateContestRequest request);
  $async.Future<$2.StandardResponse> loadContest($pb.ServerContext ctx, $2.LoadContestRequest request);
  $async.Future<$0.ContestMetadata> parseContest($pb.ServerContext ctx, $2.ParseContestRequest request);
  $async.Future<$2.ActiveContest> getActiveContest($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$3.QSOFields> getQSOFields($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$2.QSO> logQSO($pb.ServerContext ctx, $2.QSO request);
  $async.Future<$2.SnapshotMessage> getActiveQSOs($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$2.StandardResponse> deleteQSO($pb.ServerContext ctx, $2.QSO request);
  $async.Future<$3.ScoreResponse> getScore($pb.ServerContext ctx, $1.Empty request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateContest': return $2.CreateContestRequest();
      case 'LoadContest': return $2.LoadContestRequest();
      case 'ParseContest': return $2.ParseContestRequest();
      case 'GetActiveContest': return $1.Empty();
      case 'GetQSOFields': return $1.Empty();
      case 'LogQSO': return $2.QSO();
      case 'GetActiveQSOs': return $1.Empty();
      case 'DeleteQSO': return $2.QSO();
      case 'GetScore': return $1.Empty();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateContest': return this.createContest(ctx, request as $2.CreateContestRequest);
      case 'LoadContest': return this.loadContest(ctx, request as $2.LoadContestRequest);
      case 'ParseContest': return this.parseContest(ctx, request as $2.ParseContestRequest);
      case 'GetActiveContest': return this.getActiveContest(ctx, request as $1.Empty);
      case 'GetQSOFields': return this.getQSOFields(ctx, request as $1.Empty);
      case 'LogQSO': return this.logQSO(ctx, request as $2.QSO);
      case 'GetActiveQSOs': return this.getActiveQSOs(ctx, request as $1.Empty);
      case 'DeleteQSO': return this.deleteQSO(ctx, request as $2.QSO);
      case 'GetScore': return this.getScore(ctx, request as $1.Empty);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => GuiServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => GuiServiceBase$messageJson;
}

abstract class RealtimeGuiServerServiceBase extends $pb.GeneratedService {
  $async.Future<$2.QSO> draftQSO($pb.ServerContext ctx, $3.DraftQSOMessage request);
  $async.Future<$2.BinlogMessage> retrieveQSOUpdates($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$3.Spot> retrieveTelnet($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$2.StandardResponse> sendSpotToTelnet($pb.ServerContext ctx, $3.Spot request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'DraftQSO': return $3.DraftQSOMessage();
      case 'RetrieveQSOUpdates': return $1.Empty();
      case 'RetrieveTelnet': return $1.Empty();
      case 'SendSpotToTelnet': return $3.Spot();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'DraftQSO': return this.draftQSO(ctx, request as $3.DraftQSOMessage);
      case 'RetrieveQSOUpdates': return this.retrieveQSOUpdates(ctx, request as $1.Empty);
      case 'RetrieveTelnet': return this.retrieveTelnet(ctx, request as $1.Empty);
      case 'SendSpotToTelnet': return this.sendSpotToTelnet(ctx, request as $3.Spot);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => RealtimeGuiServerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => RealtimeGuiServerServiceBase$messageJson;
}

abstract class RadioServiceBase extends $pb.GeneratedService {
  $async.Future<$3.RadioStatus> getRadioMode($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$3.RadioStatus> pollRadioMode($pb.ServerContext ctx, $1.Empty request);
  $async.Future<$2.StandardResponse> radioOp($pb.ServerContext ctx, $3.RadioCommands request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetRadioMode': return $1.Empty();
      case 'PollRadioMode': return $1.Empty();
      case 'RadioOp': return $3.RadioCommands();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetRadioMode': return this.getRadioMode(ctx, request as $1.Empty);
      case 'PollRadioMode': return this.pollRadioMode(ctx, request as $1.Empty);
      case 'RadioOp': return this.radioOp(ctx, request as $3.RadioCommands);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => RadioServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => RadioServiceBase$messageJson;
}

