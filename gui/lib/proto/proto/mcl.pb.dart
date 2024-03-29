///
//  Generated code. Do not modify.
//  source: proto/mcl.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'common.pb.dart' as $2;

import 'mcl.pbenum.dart';

export 'mcl.pbenum.dart';

class Station extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Station', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'callsign')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'grid')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'region', $pb.PbFieldType.O3)
    ..m<$core.String, $core.String>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'customFields', entryClassName: 'Station.CustomFieldsEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('mcl'))
    ..hasRequiredFields = false
  ;

  Station._() : super();
  factory Station({
    $core.String? callsign,
    $core.String? grid,
    $core.int? region,
    $core.Map<$core.String, $core.String>? customFields,
  }) {
    final _result = create();
    if (callsign != null) {
      _result.callsign = callsign;
    }
    if (grid != null) {
      _result.grid = grid;
    }
    if (region != null) {
      _result.region = region;
    }
    if (customFields != null) {
      _result.customFields.addAll(customFields);
    }
    return _result;
  }
  factory Station.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Station.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Station clone() => Station()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Station copyWith(void Function(Station) updates) => super.copyWith((message) => updates(message as Station)) as Station; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Station create() => Station._();
  Station createEmptyInstance() => create();
  static $pb.PbList<Station> createRepeated() => $pb.PbList<Station>();
  @$core.pragma('dart2js:noInline')
  static Station getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Station>(create);
  static Station? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get callsign => $_getSZ(0);
  @$pb.TagNumber(1)
  set callsign($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCallsign() => $_has(0);
  @$pb.TagNumber(1)
  void clearCallsign() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get grid => $_getSZ(1);
  @$pb.TagNumber(2)
  set grid($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasGrid() => $_has(1);
  @$pb.TagNumber(2)
  void clearGrid() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get region => $_getIZ(2);
  @$pb.TagNumber(3)
  set region($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRegion() => $_has(2);
  @$pb.TagNumber(3)
  void clearRegion() => clearField(3);

  @$pb.TagNumber(4)
  $core.Map<$core.String, $core.String> get customFields => $_getMap(3);
}

class ActiveContest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActiveContest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'uuid')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOM<$2.ContestMetadata>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contest', subBuilder: $2.ContestMetadata.create)
    ..aOM<Station>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'station', subBuilder: Station.create)
    ..aInt64(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'beginTimestamp')
    ..aInt64(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'endTimestamp')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contestScript')
    ..hasRequiredFields = false
  ;

  ActiveContest._() : super();
  factory ActiveContest({
    $core.String? uuid,
    $core.String? name,
    $2.ContestMetadata? contest,
    Station? station,
    $fixnum.Int64? beginTimestamp,
    $fixnum.Int64? endTimestamp,
    $core.String? contestScript,
  }) {
    final _result = create();
    if (uuid != null) {
      _result.uuid = uuid;
    }
    if (name != null) {
      _result.name = name;
    }
    if (contest != null) {
      _result.contest = contest;
    }
    if (station != null) {
      _result.station = station;
    }
    if (beginTimestamp != null) {
      _result.beginTimestamp = beginTimestamp;
    }
    if (endTimestamp != null) {
      _result.endTimestamp = endTimestamp;
    }
    if (contestScript != null) {
      _result.contestScript = contestScript;
    }
    return _result;
  }
  factory ActiveContest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ActiveContest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ActiveContest clone() => ActiveContest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ActiveContest copyWith(void Function(ActiveContest) updates) => super.copyWith((message) => updates(message as ActiveContest)) as ActiveContest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ActiveContest create() => ActiveContest._();
  ActiveContest createEmptyInstance() => create();
  static $pb.PbList<ActiveContest> createRepeated() => $pb.PbList<ActiveContest>();
  @$core.pragma('dart2js:noInline')
  static ActiveContest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ActiveContest>(create);
  static ActiveContest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get uuid => $_getSZ(0);
  @$pb.TagNumber(1)
  set uuid($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUuid() => $_has(0);
  @$pb.TagNumber(1)
  void clearUuid() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $2.ContestMetadata get contest => $_getN(2);
  @$pb.TagNumber(3)
  set contest($2.ContestMetadata v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasContest() => $_has(2);
  @$pb.TagNumber(3)
  void clearContest() => clearField(3);
  @$pb.TagNumber(3)
  $2.ContestMetadata ensureContest() => $_ensure(2);

  @$pb.TagNumber(4)
  Station get station => $_getN(3);
  @$pb.TagNumber(4)
  set station(Station v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasStation() => $_has(3);
  @$pb.TagNumber(4)
  void clearStation() => clearField(4);
  @$pb.TagNumber(4)
  Station ensureStation() => $_ensure(3);

  @$pb.TagNumber(5)
  $fixnum.Int64 get beginTimestamp => $_getI64(4);
  @$pb.TagNumber(5)
  set beginTimestamp($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasBeginTimestamp() => $_has(4);
  @$pb.TagNumber(5)
  void clearBeginTimestamp() => clearField(5);

  @$pb.TagNumber(6)
  $fixnum.Int64 get endTimestamp => $_getI64(5);
  @$pb.TagNumber(6)
  set endTimestamp($fixnum.Int64 v) { $_setInt64(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasEndTimestamp() => $_has(5);
  @$pb.TagNumber(6)
  void clearEndTimestamp() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get contestScript => $_getSZ(6);
  @$pb.TagNumber(7)
  set contestScript($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasContestScript() => $_has(6);
  @$pb.TagNumber(7)
  void clearContestScript() => clearField(7);
}

class QSO extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QSO', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'uid')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'localCallsign')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'dxCallsign')
    ..aInt64(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'time')
    ..aInt64(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'freq')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'mode')
    ..m<$core.String, $core.String>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'exchangeSent', entryClassName: 'QSO.ExchangeSentEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('mcl'))
    ..m<$core.String, $core.String>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'exchangeRcvd', entryClassName: 'QSO.ExchangeRcvdEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('mcl'))
    ..e<QSOType>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: QSOType.qso, valueOf: QSOType.valueOf, enumValues: QSOType.values)
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'operator')
    ..hasRequiredFields = false
  ;

  QSO._() : super();
  factory QSO({
    $core.String? uid,
    $core.String? localCallsign,
    $core.String? dxCallsign,
    $fixnum.Int64? time,
    $fixnum.Int64? freq,
    $core.String? mode,
    $core.Map<$core.String, $core.String>? exchangeSent,
    $core.Map<$core.String, $core.String>? exchangeRcvd,
    QSOType? type,
    $core.String? operator,
  }) {
    final _result = create();
    if (uid != null) {
      _result.uid = uid;
    }
    if (localCallsign != null) {
      _result.localCallsign = localCallsign;
    }
    if (dxCallsign != null) {
      _result.dxCallsign = dxCallsign;
    }
    if (time != null) {
      _result.time = time;
    }
    if (freq != null) {
      _result.freq = freq;
    }
    if (mode != null) {
      _result.mode = mode;
    }
    if (exchangeSent != null) {
      _result.exchangeSent.addAll(exchangeSent);
    }
    if (exchangeRcvd != null) {
      _result.exchangeRcvd.addAll(exchangeRcvd);
    }
    if (type != null) {
      _result.type = type;
    }
    if (operator != null) {
      _result.operator = operator;
    }
    return _result;
  }
  factory QSO.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QSO.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QSO clone() => QSO()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QSO copyWith(void Function(QSO) updates) => super.copyWith((message) => updates(message as QSO)) as QSO; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QSO create() => QSO._();
  QSO createEmptyInstance() => create();
  static $pb.PbList<QSO> createRepeated() => $pb.PbList<QSO>();
  @$core.pragma('dart2js:noInline')
  static QSO getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QSO>(create);
  static QSO? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get uid => $_getSZ(0);
  @$pb.TagNumber(1)
  set uid($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUid() => $_has(0);
  @$pb.TagNumber(1)
  void clearUid() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get localCallsign => $_getSZ(1);
  @$pb.TagNumber(2)
  set localCallsign($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLocalCallsign() => $_has(1);
  @$pb.TagNumber(2)
  void clearLocalCallsign() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get dxCallsign => $_getSZ(2);
  @$pb.TagNumber(3)
  set dxCallsign($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDxCallsign() => $_has(2);
  @$pb.TagNumber(3)
  void clearDxCallsign() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get time => $_getI64(3);
  @$pb.TagNumber(4)
  set time($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTime() => $_has(3);
  @$pb.TagNumber(4)
  void clearTime() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get freq => $_getI64(4);
  @$pb.TagNumber(5)
  set freq($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasFreq() => $_has(4);
  @$pb.TagNumber(5)
  void clearFreq() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get mode => $_getSZ(5);
  @$pb.TagNumber(6)
  set mode($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasMode() => $_has(5);
  @$pb.TagNumber(6)
  void clearMode() => clearField(6);

  @$pb.TagNumber(7)
  $core.Map<$core.String, $core.String> get exchangeSent => $_getMap(6);

  @$pb.TagNumber(8)
  $core.Map<$core.String, $core.String> get exchangeRcvd => $_getMap(7);

  @$pb.TagNumber(9)
  QSOType get type => $_getN(8);
  @$pb.TagNumber(9)
  set type(QSOType v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasType() => $_has(8);
  @$pb.TagNumber(9)
  void clearType() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get operator => $_getSZ(9);
  @$pb.TagNumber(10)
  set operator($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasOperator() => $_has(9);
  @$pb.TagNumber(10)
  void clearOperator() => clearField(10);
}

class QSOOp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QSOOp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..e<QSOOperationType>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: QSOOperationType.add_or_update, valueOf: QSOOperationType.valueOf, enumValues: QSOOperationType.values)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'uid')
    ..hasRequiredFields = false
  ;

  QSOOp._() : super();
  factory QSOOp({
    QSOOperationType? type,
    $core.String? uid,
  }) {
    final _result = create();
    if (type != null) {
      _result.type = type;
    }
    if (uid != null) {
      _result.uid = uid;
    }
    return _result;
  }
  factory QSOOp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QSOOp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QSOOp clone() => QSOOp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QSOOp copyWith(void Function(QSOOp) updates) => super.copyWith((message) => updates(message as QSOOp)) as QSOOp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QSOOp create() => QSOOp._();
  QSOOp createEmptyInstance() => create();
  static $pb.PbList<QSOOp> createRepeated() => $pb.PbList<QSOOp>();
  @$core.pragma('dart2js:noInline')
  static QSOOp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QSOOp>(create);
  static QSOOp? _defaultInstance;

  @$pb.TagNumber(1)
  QSOOperationType get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(QSOOperationType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get uid => $_getSZ(1);
  @$pb.TagNumber(2)
  set uid($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUid() => $_has(1);
  @$pb.TagNumber(2)
  void clearUid() => clearField(2);
}

enum BinlogMessage_Message {
  qsoOp, 
  qso, 
  notSet
}

class BinlogMessage extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, BinlogMessage_Message> _BinlogMessage_MessageByTag = {
    4 : BinlogMessage_Message.qsoOp,
    5 : BinlogMessage_Message.qso,
    0 : BinlogMessage_Message.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BinlogMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..oo(0, [4, 5])
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'serial', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOM<QSOOp>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'qsoOp', subBuilder: QSOOp.create)
    ..aOM<QSO>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'qso', subBuilder: QSO.create)
    ..hasRequiredFields = false
  ;

  BinlogMessage._() : super();
  factory BinlogMessage({
    $fixnum.Int64? serial,
    QSOOp? qsoOp,
    QSO? qso,
  }) {
    final _result = create();
    if (serial != null) {
      _result.serial = serial;
    }
    if (qsoOp != null) {
      _result.qsoOp = qsoOp;
    }
    if (qso != null) {
      _result.qso = qso;
    }
    return _result;
  }
  factory BinlogMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BinlogMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BinlogMessage clone() => BinlogMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BinlogMessage copyWith(void Function(BinlogMessage) updates) => super.copyWith((message) => updates(message as BinlogMessage)) as BinlogMessage; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BinlogMessage create() => BinlogMessage._();
  BinlogMessage createEmptyInstance() => create();
  static $pb.PbList<BinlogMessage> createRepeated() => $pb.PbList<BinlogMessage>();
  @$core.pragma('dart2js:noInline')
  static BinlogMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BinlogMessage>(create);
  static BinlogMessage? _defaultInstance;

  BinlogMessage_Message whichMessage() => _BinlogMessage_MessageByTag[$_whichOneof(0)]!;
  void clearMessage() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $fixnum.Int64 get serial => $_getI64(0);
  @$pb.TagNumber(1)
  set serial($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSerial() => $_has(0);
  @$pb.TagNumber(1)
  void clearSerial() => clearField(1);

  @$pb.TagNumber(4)
  QSOOp get qsoOp => $_getN(1);
  @$pb.TagNumber(4)
  set qsoOp(QSOOp v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasQsoOp() => $_has(1);
  @$pb.TagNumber(4)
  void clearQsoOp() => clearField(4);
  @$pb.TagNumber(4)
  QSOOp ensureQsoOp() => $_ensure(1);

  @$pb.TagNumber(5)
  QSO get qso => $_getN(2);
  @$pb.TagNumber(5)
  set qso(QSO v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasQso() => $_has(2);
  @$pb.TagNumber(5)
  void clearQso() => clearField(5);
  @$pb.TagNumber(5)
  QSO ensureQso() => $_ensure(2);
}

class SnapshotMessage extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SnapshotMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..pc<QSO>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'qso', $pb.PbFieldType.PM, subBuilder: QSO.create)
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'serial', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  SnapshotMessage._() : super();
  factory SnapshotMessage({
    $core.Iterable<QSO>? qso,
    $fixnum.Int64? serial,
  }) {
    final _result = create();
    if (qso != null) {
      _result.qso.addAll(qso);
    }
    if (serial != null) {
      _result.serial = serial;
    }
    return _result;
  }
  factory SnapshotMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotMessage clone() => SnapshotMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotMessage copyWith(void Function(SnapshotMessage) updates) => super.copyWith((message) => updates(message as SnapshotMessage)) as SnapshotMessage; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SnapshotMessage create() => SnapshotMessage._();
  SnapshotMessage createEmptyInstance() => create();
  static $pb.PbList<SnapshotMessage> createRepeated() => $pb.PbList<SnapshotMessage>();
  @$core.pragma('dart2js:noInline')
  static SnapshotMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotMessage>(create);
  static SnapshotMessage? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<QSO> get qso => $_getList(0);

  @$pb.TagNumber(2)
  $fixnum.Int64 get serial => $_getI64(1);
  @$pb.TagNumber(2)
  set serial($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSerial() => $_has(1);
  @$pb.TagNumber(2)
  void clearSerial() => clearField(2);
}

class BinlogMessageSet extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BinlogMessageSet', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..pc<BinlogMessage>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'messages', $pb.PbFieldType.PM, subBuilder: BinlogMessage.create)
    ..hasRequiredFields = false
  ;

  BinlogMessageSet._() : super();
  factory BinlogMessageSet({
    $core.Iterable<BinlogMessage>? messages,
  }) {
    final _result = create();
    if (messages != null) {
      _result.messages.addAll(messages);
    }
    return _result;
  }
  factory BinlogMessageSet.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BinlogMessageSet.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BinlogMessageSet clone() => BinlogMessageSet()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BinlogMessageSet copyWith(void Function(BinlogMessageSet) updates) => super.copyWith((message) => updates(message as BinlogMessageSet)) as BinlogMessageSet; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BinlogMessageSet create() => BinlogMessageSet._();
  BinlogMessageSet createEmptyInstance() => create();
  static $pb.PbList<BinlogMessageSet> createRepeated() => $pb.PbList<BinlogMessageSet>();
  @$core.pragma('dart2js:noInline')
  static BinlogMessageSet getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BinlogMessageSet>(create);
  static BinlogMessageSet? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<BinlogMessage> get messages => $_getList(0);
}

class StandardResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StandardResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..e<ResultCode>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'resultCode', $pb.PbFieldType.OE, defaultOrMaker: ResultCode.success, valueOf: ResultCode.valueOf, enumValues: ResultCode.values)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errorMessage')
    ..hasRequiredFields = false
  ;

  StandardResponse._() : super();
  factory StandardResponse({
    ResultCode? resultCode,
    $core.String? errorMessage,
  }) {
    final _result = create();
    if (resultCode != null) {
      _result.resultCode = resultCode;
    }
    if (errorMessage != null) {
      _result.errorMessage = errorMessage;
    }
    return _result;
  }
  factory StandardResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StandardResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StandardResponse clone() => StandardResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StandardResponse copyWith(void Function(StandardResponse) updates) => super.copyWith((message) => updates(message as StandardResponse)) as StandardResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StandardResponse create() => StandardResponse._();
  StandardResponse createEmptyInstance() => create();
  static $pb.PbList<StandardResponse> createRepeated() => $pb.PbList<StandardResponse>();
  @$core.pragma('dart2js:noInline')
  static StandardResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StandardResponse>(create);
  static StandardResponse? _defaultInstance;

  @$pb.TagNumber(1)
  ResultCode get resultCode => $_getN(0);
  @$pb.TagNumber(1)
  set resultCode(ResultCode v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasResultCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearResultCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get errorMessage => $_getSZ(1);
  @$pb.TagNumber(2)
  set errorMessage($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasErrorMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearErrorMessage() => clearField(2);
}

class RetrieveBinlogRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RetrieveBinlogRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'serialStart', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'serialEnd', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  RetrieveBinlogRequest._() : super();
  factory RetrieveBinlogRequest({
    $fixnum.Int64? serialStart,
    $fixnum.Int64? serialEnd,
  }) {
    final _result = create();
    if (serialStart != null) {
      _result.serialStart = serialStart;
    }
    if (serialEnd != null) {
      _result.serialEnd = serialEnd;
    }
    return _result;
  }
  factory RetrieveBinlogRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RetrieveBinlogRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RetrieveBinlogRequest clone() => RetrieveBinlogRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RetrieveBinlogRequest copyWith(void Function(RetrieveBinlogRequest) updates) => super.copyWith((message) => updates(message as RetrieveBinlogRequest)) as RetrieveBinlogRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RetrieveBinlogRequest create() => RetrieveBinlogRequest._();
  RetrieveBinlogRequest createEmptyInstance() => create();
  static $pb.PbList<RetrieveBinlogRequest> createRepeated() => $pb.PbList<RetrieveBinlogRequest>();
  @$core.pragma('dart2js:noInline')
  static RetrieveBinlogRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RetrieveBinlogRequest>(create);
  static RetrieveBinlogRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get serialStart => $_getI64(0);
  @$pb.TagNumber(1)
  set serialStart($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSerialStart() => $_has(0);
  @$pb.TagNumber(1)
  void clearSerialStart() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get serialEnd => $_getI64(1);
  @$pb.TagNumber(2)
  set serialEnd($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSerialEnd() => $_has(1);
  @$pb.TagNumber(2)
  void clearSerialEnd() => clearField(2);
}

class LoadContestRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'LoadContestRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'databaseName')
    ..hasRequiredFields = false
  ;

  LoadContestRequest._() : super();
  factory LoadContestRequest({
    $core.String? databaseName,
  }) {
    final _result = create();
    if (databaseName != null) {
      _result.databaseName = databaseName;
    }
    return _result;
  }
  factory LoadContestRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LoadContestRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  LoadContestRequest clone() => LoadContestRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  LoadContestRequest copyWith(void Function(LoadContestRequest) updates) => super.copyWith((message) => updates(message as LoadContestRequest)) as LoadContestRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LoadContestRequest create() => LoadContestRequest._();
  LoadContestRequest createEmptyInstance() => create();
  static $pb.PbList<LoadContestRequest> createRepeated() => $pb.PbList<LoadContestRequest>();
  @$core.pragma('dart2js:noInline')
  static LoadContestRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LoadContestRequest>(create);
  static LoadContestRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get databaseName => $_getSZ(0);
  @$pb.TagNumber(1)
  set databaseName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDatabaseName() => $_has(0);
  @$pb.TagNumber(1)
  void clearDatabaseName() => clearField(1);
}

class CreateContestRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CreateContestRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'databaseName')
    ..aOM<ActiveContest>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contest', subBuilder: ActiveContest.create)
    ..hasRequiredFields = false
  ;

  CreateContestRequest._() : super();
  factory CreateContestRequest({
    $core.String? databaseName,
    ActiveContest? contest,
  }) {
    final _result = create();
    if (databaseName != null) {
      _result.databaseName = databaseName;
    }
    if (contest != null) {
      _result.contest = contest;
    }
    return _result;
  }
  factory CreateContestRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateContestRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateContestRequest clone() => CreateContestRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateContestRequest copyWith(void Function(CreateContestRequest) updates) => super.copyWith((message) => updates(message as CreateContestRequest)) as CreateContestRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CreateContestRequest create() => CreateContestRequest._();
  CreateContestRequest createEmptyInstance() => create();
  static $pb.PbList<CreateContestRequest> createRepeated() => $pb.PbList<CreateContestRequest>();
  @$core.pragma('dart2js:noInline')
  static CreateContestRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateContestRequest>(create);
  static CreateContestRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get databaseName => $_getSZ(0);
  @$pb.TagNumber(1)
  set databaseName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDatabaseName() => $_has(0);
  @$pb.TagNumber(1)
  void clearDatabaseName() => clearField(1);

  @$pb.TagNumber(2)
  ActiveContest get contest => $_getN(1);
  @$pb.TagNumber(2)
  set contest(ActiveContest v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasContest() => $_has(1);
  @$pb.TagNumber(2)
  void clearContest() => clearField(2);
  @$pb.TagNumber(2)
  ActiveContest ensureContest() => $_ensure(1);
}

class ParseContestRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ParseContestRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contestDescriptor')
    ..hasRequiredFields = false
  ;

  ParseContestRequest._() : super();
  factory ParseContestRequest({
    $core.String? contestDescriptor,
  }) {
    final _result = create();
    if (contestDescriptor != null) {
      _result.contestDescriptor = contestDescriptor;
    }
    return _result;
  }
  factory ParseContestRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ParseContestRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ParseContestRequest clone() => ParseContestRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ParseContestRequest copyWith(void Function(ParseContestRequest) updates) => super.copyWith((message) => updates(message as ParseContestRequest)) as ParseContestRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ParseContestRequest create() => ParseContestRequest._();
  ParseContestRequest createEmptyInstance() => create();
  static $pb.PbList<ParseContestRequest> createRepeated() => $pb.PbList<ParseContestRequest>();
  @$core.pragma('dart2js:noInline')
  static ParseContestRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ParseContestRequest>(create);
  static ParseContestRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get contestDescriptor => $_getSZ(0);
  @$pb.TagNumber(1)
  set contestDescriptor($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasContestDescriptor() => $_has(0);
  @$pb.TagNumber(1)
  void clearContestDescriptor() => clearField(1);
}

