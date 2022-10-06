///
//  Generated code. Do not modify.
//  source: proto/mclgui.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'mcl.pbenum.dart' as $0;
import 'mclgui.pbenum.dart';

export 'mclgui.pbenum.dart';

class QSOMessage extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QSOMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'dxCallsign')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'operator')
    ..aInt64(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'time')
    ..aInt64(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'freq')
    ..e<$0.Mode>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'mode', $pb.PbFieldType.OE, defaultOrMaker: $0.Mode.phone, valueOf: $0.Mode.valueOf, enumValues: $0.Mode.values)
    ..pPS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'exchangeSent')
    ..pPS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'exchangeRcvd')
    ..hasRequiredFields = false
  ;

  QSOMessage._() : super();
  factory QSOMessage({
    $core.String? dxCallsign,
    $core.String? operator,
    $fixnum.Int64? time,
    $fixnum.Int64? freq,
    $0.Mode? mode,
    $core.Iterable<$core.String>? exchangeSent,
    $core.Iterable<$core.String>? exchangeRcvd,
  }) {
    final _result = create();
    if (dxCallsign != null) {
      _result.dxCallsign = dxCallsign;
    }
    if (operator != null) {
      _result.operator = operator;
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
    return _result;
  }
  factory QSOMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QSOMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QSOMessage clone() => QSOMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QSOMessage copyWith(void Function(QSOMessage) updates) => super.copyWith((message) => updates(message as QSOMessage)) as QSOMessage; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QSOMessage create() => QSOMessage._();
  QSOMessage createEmptyInstance() => create();
  static $pb.PbList<QSOMessage> createRepeated() => $pb.PbList<QSOMessage>();
  @$core.pragma('dart2js:noInline')
  static QSOMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QSOMessage>(create);
  static QSOMessage? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get dxCallsign => $_getSZ(0);
  @$pb.TagNumber(1)
  set dxCallsign($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDxCallsign() => $_has(0);
  @$pb.TagNumber(1)
  void clearDxCallsign() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get operator => $_getSZ(1);
  @$pb.TagNumber(2)
  set operator($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasOperator() => $_has(1);
  @$pb.TagNumber(2)
  void clearOperator() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get time => $_getI64(2);
  @$pb.TagNumber(3)
  set time($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTime() => $_has(2);
  @$pb.TagNumber(3)
  void clearTime() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get freq => $_getI64(3);
  @$pb.TagNumber(4)
  set freq($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFreq() => $_has(3);
  @$pb.TagNumber(4)
  void clearFreq() => clearField(4);

  @$pb.TagNumber(5)
  $0.Mode get mode => $_getN(4);
  @$pb.TagNumber(5)
  set mode($0.Mode v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasMode() => $_has(4);
  @$pb.TagNumber(5)
  void clearMode() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<$core.String> get exchangeSent => $_getList(5);

  @$pb.TagNumber(7)
  $core.List<$core.String> get exchangeRcvd => $_getList(6);
}

class DraftQSOMessage extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DraftQSOMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'dxCallsign')
    ..hasRequiredFields = false
  ;

  DraftQSOMessage._() : super();
  factory DraftQSOMessage({
    $core.String? dxCallsign,
  }) {
    final _result = create();
    if (dxCallsign != null) {
      _result.dxCallsign = dxCallsign;
    }
    return _result;
  }
  factory DraftQSOMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DraftQSOMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DraftQSOMessage clone() => DraftQSOMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DraftQSOMessage copyWith(void Function(DraftQSOMessage) updates) => super.copyWith((message) => updates(message as DraftQSOMessage)) as DraftQSOMessage; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DraftQSOMessage create() => DraftQSOMessage._();
  DraftQSOMessage createEmptyInstance() => create();
  static $pb.PbList<DraftQSOMessage> createRepeated() => $pb.PbList<DraftQSOMessage>();
  @$core.pragma('dart2js:noInline')
  static DraftQSOMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DraftQSOMessage>(create);
  static DraftQSOMessage? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get dxCallsign => $_getSZ(0);
  @$pb.TagNumber(1)
  set dxCallsign($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDxCallsign() => $_has(0);
  @$pb.TagNumber(1)
  void clearDxCallsign() => clearField(1);
}

class ScoreResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ScoreResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..m<$core.String, $fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'categoryScore', entryClassName: 'ScoreResponse.CategoryScoreEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.O6, packageName: const $pb.PackageName('mcl'))
    ..hasRequiredFields = false
  ;

  ScoreResponse._() : super();
  factory ScoreResponse({
    $core.Map<$core.String, $fixnum.Int64>? categoryScore,
  }) {
    final _result = create();
    if (categoryScore != null) {
      _result.categoryScore.addAll(categoryScore);
    }
    return _result;
  }
  factory ScoreResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ScoreResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ScoreResponse clone() => ScoreResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ScoreResponse copyWith(void Function(ScoreResponse) updates) => super.copyWith((message) => updates(message as ScoreResponse)) as ScoreResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ScoreResponse create() => ScoreResponse._();
  ScoreResponse createEmptyInstance() => create();
  static $pb.PbList<ScoreResponse> createRepeated() => $pb.PbList<ScoreResponse>();
  @$core.pragma('dart2js:noInline')
  static ScoreResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ScoreResponse>(create);
  static ScoreResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$core.String, $fixnum.Int64> get categoryScore => $_getMap(0);
}

class Spot extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Spot', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'dxCallsign')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'deCallsign')
    ..aInt64(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'spotTimestamp')
    ..aInt64(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'spotFrequency')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'memo')
    ..hasRequiredFields = false
  ;

  Spot._() : super();
  factory Spot({
    $core.String? dxCallsign,
    $core.String? deCallsign,
    $fixnum.Int64? spotTimestamp,
    $fixnum.Int64? spotFrequency,
    $core.String? memo,
  }) {
    final _result = create();
    if (dxCallsign != null) {
      _result.dxCallsign = dxCallsign;
    }
    if (deCallsign != null) {
      _result.deCallsign = deCallsign;
    }
    if (spotTimestamp != null) {
      _result.spotTimestamp = spotTimestamp;
    }
    if (spotFrequency != null) {
      _result.spotFrequency = spotFrequency;
    }
    if (memo != null) {
      _result.memo = memo;
    }
    return _result;
  }
  factory Spot.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Spot.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Spot clone() => Spot()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Spot copyWith(void Function(Spot) updates) => super.copyWith((message) => updates(message as Spot)) as Spot; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Spot create() => Spot._();
  Spot createEmptyInstance() => create();
  static $pb.PbList<Spot> createRepeated() => $pb.PbList<Spot>();
  @$core.pragma('dart2js:noInline')
  static Spot getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Spot>(create);
  static Spot? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get dxCallsign => $_getSZ(0);
  @$pb.TagNumber(1)
  set dxCallsign($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDxCallsign() => $_has(0);
  @$pb.TagNumber(1)
  void clearDxCallsign() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get deCallsign => $_getSZ(1);
  @$pb.TagNumber(2)
  set deCallsign($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDeCallsign() => $_has(1);
  @$pb.TagNumber(2)
  void clearDeCallsign() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get spotTimestamp => $_getI64(2);
  @$pb.TagNumber(3)
  set spotTimestamp($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasSpotTimestamp() => $_has(2);
  @$pb.TagNumber(3)
  void clearSpotTimestamp() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get spotFrequency => $_getI64(3);
  @$pb.TagNumber(4)
  set spotFrequency($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasSpotFrequency() => $_has(3);
  @$pb.TagNumber(4)
  void clearSpotFrequency() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get memo => $_getSZ(4);
  @$pb.TagNumber(5)
  set memo($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasMemo() => $_has(4);
  @$pb.TagNumber(5)
  void clearMemo() => clearField(5);
}

class CallsignLookupResultCatrgory extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CallsignLookupResultCatrgory', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..pPS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'possibleCallsigns')
    ..hasRequiredFields = false
  ;

  CallsignLookupResultCatrgory._() : super();
  factory CallsignLookupResultCatrgory({
    $core.Iterable<$core.String>? possibleCallsigns,
  }) {
    final _result = create();
    if (possibleCallsigns != null) {
      _result.possibleCallsigns.addAll(possibleCallsigns);
    }
    return _result;
  }
  factory CallsignLookupResultCatrgory.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CallsignLookupResultCatrgory.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CallsignLookupResultCatrgory clone() => CallsignLookupResultCatrgory()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CallsignLookupResultCatrgory copyWith(void Function(CallsignLookupResultCatrgory) updates) => super.copyWith((message) => updates(message as CallsignLookupResultCatrgory)) as CallsignLookupResultCatrgory; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CallsignLookupResultCatrgory create() => CallsignLookupResultCatrgory._();
  CallsignLookupResultCatrgory createEmptyInstance() => create();
  static $pb.PbList<CallsignLookupResultCatrgory> createRepeated() => $pb.PbList<CallsignLookupResultCatrgory>();
  @$core.pragma('dart2js:noInline')
  static CallsignLookupResultCatrgory getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CallsignLookupResultCatrgory>(create);
  static CallsignLookupResultCatrgory? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get possibleCallsigns => $_getList(0);
}

class BandStatus extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BandStatus', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..m<$fixnum.Int64, $core.String>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'bandStatus', entryClassName: 'BandStatus.BandStatusEntry', keyFieldType: $pb.PbFieldType.O6, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('mcl'))
    ..hasRequiredFields = false
  ;

  BandStatus._() : super();
  factory BandStatus({
    $core.Map<$fixnum.Int64, $core.String>? bandStatus,
  }) {
    final _result = create();
    if (bandStatus != null) {
      _result.bandStatus.addAll(bandStatus);
    }
    return _result;
  }
  factory BandStatus.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BandStatus.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BandStatus clone() => BandStatus()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BandStatus copyWith(void Function(BandStatus) updates) => super.copyWith((message) => updates(message as BandStatus)) as BandStatus; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BandStatus create() => BandStatus._();
  BandStatus createEmptyInstance() => create();
  static $pb.PbList<BandStatus> createRepeated() => $pb.PbList<BandStatus>();
  @$core.pragma('dart2js:noInline')
  static BandStatus getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BandStatus>(create);
  static BandStatus? _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$fixnum.Int64, $core.String> get bandStatus => $_getMap(0);
}

class CallsignLookupResult extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CallsignLookupResult', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..m<$core.String, CallsignLookupResultCatrgory>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'databaseLookup', entryClassName: 'CallsignLookupResult.DatabaseLookupEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OM, valueCreator: CallsignLookupResultCatrgory.create, packageName: const $pb.PackageName('mcl'))
    ..m<$core.String, BandStatus>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'bandStatus', entryClassName: 'CallsignLookupResult.BandStatusEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OM, valueCreator: BandStatus.create, packageName: const $pb.PackageName('mcl'))
    ..hasRequiredFields = false
  ;

  CallsignLookupResult._() : super();
  factory CallsignLookupResult({
    $core.Map<$core.String, CallsignLookupResultCatrgory>? databaseLookup,
    $core.Map<$core.String, BandStatus>? bandStatus,
  }) {
    final _result = create();
    if (databaseLookup != null) {
      _result.databaseLookup.addAll(databaseLookup);
    }
    if (bandStatus != null) {
      _result.bandStatus.addAll(bandStatus);
    }
    return _result;
  }
  factory CallsignLookupResult.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CallsignLookupResult.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CallsignLookupResult clone() => CallsignLookupResult()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CallsignLookupResult copyWith(void Function(CallsignLookupResult) updates) => super.copyWith((message) => updates(message as CallsignLookupResult)) as CallsignLookupResult; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CallsignLookupResult create() => CallsignLookupResult._();
  CallsignLookupResult createEmptyInstance() => create();
  static $pb.PbList<CallsignLookupResult> createRepeated() => $pb.PbList<CallsignLookupResult>();
  @$core.pragma('dart2js:noInline')
  static CallsignLookupResult getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CallsignLookupResult>(create);
  static CallsignLookupResult? _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$core.String, CallsignLookupResultCatrgory> get databaseLookup => $_getMap(0);

  @$pb.TagNumber(2)
  $core.Map<$core.String, BandStatus> get bandStatus => $_getMap(1);
}

class RadioStatus extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RadioStatus', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'freq')
    ..e<RadioMode>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'mode', $pb.PbFieldType.OE, defaultOrMaker: RadioMode.UNKNOWN, valueOf: RadioMode.valueOf, enumValues: RadioMode.values)
    ..aOB(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'tx')
    ..hasRequiredFields = false
  ;

  RadioStatus._() : super();
  factory RadioStatus({
    $fixnum.Int64? freq,
    RadioMode? mode,
    $core.bool? tx,
  }) {
    final _result = create();
    if (freq != null) {
      _result.freq = freq;
    }
    if (mode != null) {
      _result.mode = mode;
    }
    if (tx != null) {
      _result.tx = tx;
    }
    return _result;
  }
  factory RadioStatus.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RadioStatus.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RadioStatus clone() => RadioStatus()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RadioStatus copyWith(void Function(RadioStatus) updates) => super.copyWith((message) => updates(message as RadioStatus)) as RadioStatus; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RadioStatus create() => RadioStatus._();
  RadioStatus createEmptyInstance() => create();
  static $pb.PbList<RadioStatus> createRepeated() => $pb.PbList<RadioStatus>();
  @$core.pragma('dart2js:noInline')
  static RadioStatus getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RadioStatus>(create);
  static RadioStatus? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get freq => $_getI64(0);
  @$pb.TagNumber(1)
  set freq($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFreq() => $_has(0);
  @$pb.TagNumber(1)
  void clearFreq() => clearField(1);

  @$pb.TagNumber(2)
  RadioMode get mode => $_getN(1);
  @$pb.TagNumber(2)
  set mode(RadioMode v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasMode() => $_has(1);
  @$pb.TagNumber(2)
  void clearMode() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get tx => $_getBF(2);
  @$pb.TagNumber(3)
  set tx($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTx() => $_has(2);
  @$pb.TagNumber(3)
  void clearTx() => clearField(3);
}

enum RadioCommands_Op {
  sendCwMessage, 
  setCwSpeed, 
  setRadioBandMode, 
  abort, 
  notSet
}

class RadioCommands extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, RadioCommands_Op> _RadioCommands_OpByTag = {
    1 : RadioCommands_Op.sendCwMessage,
    2 : RadioCommands_Op.setCwSpeed,
    3 : RadioCommands_Op.setRadioBandMode,
    4 : RadioCommands_Op.abort,
    0 : RadioCommands_Op.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RadioCommands', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..oo(0, [1, 2, 3, 4])
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sendCwMessage')
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'setCwSpeed', $pb.PbFieldType.O3)
    ..e<RadioMode>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'setRadioBandMode', $pb.PbFieldType.OE, defaultOrMaker: RadioMode.UNKNOWN, valueOf: RadioMode.valueOf, enumValues: RadioMode.values)
    ..aOB(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'abort')
    ..hasRequiredFields = false
  ;

  RadioCommands._() : super();
  factory RadioCommands({
    $core.String? sendCwMessage,
    $core.int? setCwSpeed,
    RadioMode? setRadioBandMode,
    $core.bool? abort,
  }) {
    final _result = create();
    if (sendCwMessage != null) {
      _result.sendCwMessage = sendCwMessage;
    }
    if (setCwSpeed != null) {
      _result.setCwSpeed = setCwSpeed;
    }
    if (setRadioBandMode != null) {
      _result.setRadioBandMode = setRadioBandMode;
    }
    if (abort != null) {
      _result.abort = abort;
    }
    return _result;
  }
  factory RadioCommands.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RadioCommands.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RadioCommands clone() => RadioCommands()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RadioCommands copyWith(void Function(RadioCommands) updates) => super.copyWith((message) => updates(message as RadioCommands)) as RadioCommands; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RadioCommands create() => RadioCommands._();
  RadioCommands createEmptyInstance() => create();
  static $pb.PbList<RadioCommands> createRepeated() => $pb.PbList<RadioCommands>();
  @$core.pragma('dart2js:noInline')
  static RadioCommands getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RadioCommands>(create);
  static RadioCommands? _defaultInstance;

  RadioCommands_Op whichOp() => _RadioCommands_OpByTag[$_whichOneof(0)]!;
  void clearOp() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.String get sendCwMessage => $_getSZ(0);
  @$pb.TagNumber(1)
  set sendCwMessage($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSendCwMessage() => $_has(0);
  @$pb.TagNumber(1)
  void clearSendCwMessage() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get setCwSpeed => $_getIZ(1);
  @$pb.TagNumber(2)
  set setCwSpeed($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSetCwSpeed() => $_has(1);
  @$pb.TagNumber(2)
  void clearSetCwSpeed() => clearField(2);

  @$pb.TagNumber(3)
  RadioMode get setRadioBandMode => $_getN(2);
  @$pb.TagNumber(3)
  set setRadioBandMode(RadioMode v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasSetRadioBandMode() => $_has(2);
  @$pb.TagNumber(3)
  void clearSetRadioBandMode() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get abort => $_getBF(3);
  @$pb.TagNumber(4)
  set abort($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAbort() => $_has(3);
  @$pb.TagNumber(4)
  void clearAbort() => clearField(4);
}
