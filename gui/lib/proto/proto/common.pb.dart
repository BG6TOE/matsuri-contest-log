///
//  Generated code. Do not modify.
//  source: proto/common.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'common.pbenum.dart';

class ContestFieldInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ContestFieldInfo', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'DisplayName', protoName: 'DisplayName')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Description', protoName: 'Description')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FieldType', protoName: 'FieldType')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AdifName', protoName: 'AdifName')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ValueRegex', protoName: 'ValueRegex')
    ..pPS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ValidValues', protoName: 'ValidValues')
    ..hasRequiredFields = false
  ;

  ContestFieldInfo._() : super();
  factory ContestFieldInfo({
    $core.String? displayName,
    $core.String? description,
    $core.String? fieldType,
    $core.String? adifName,
    $core.String? valueRegex,
    $core.Iterable<$core.String>? validValues,
  }) {
    final _result = create();
    if (displayName != null) {
      _result.displayName = displayName;
    }
    if (description != null) {
      _result.description = description;
    }
    if (fieldType != null) {
      _result.fieldType = fieldType;
    }
    if (adifName != null) {
      _result.adifName = adifName;
    }
    if (valueRegex != null) {
      _result.valueRegex = valueRegex;
    }
    if (validValues != null) {
      _result.validValues.addAll(validValues);
    }
    return _result;
  }
  factory ContestFieldInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ContestFieldInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ContestFieldInfo clone() => ContestFieldInfo()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ContestFieldInfo copyWith(void Function(ContestFieldInfo) updates) => super.copyWith((message) => updates(message as ContestFieldInfo)) as ContestFieldInfo; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ContestFieldInfo create() => ContestFieldInfo._();
  ContestFieldInfo createEmptyInstance() => create();
  static $pb.PbList<ContestFieldInfo> createRepeated() => $pb.PbList<ContestFieldInfo>();
  @$core.pragma('dart2js:noInline')
  static ContestFieldInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ContestFieldInfo>(create);
  static ContestFieldInfo? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get displayName => $_getSZ(0);
  @$pb.TagNumber(1)
  set displayName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDisplayName() => $_has(0);
  @$pb.TagNumber(1)
  void clearDisplayName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get description => $_getSZ(1);
  @$pb.TagNumber(2)
  set description($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDescription() => $_has(1);
  @$pb.TagNumber(2)
  void clearDescription() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get fieldType => $_getSZ(2);
  @$pb.TagNumber(3)
  set fieldType($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFieldType() => $_has(2);
  @$pb.TagNumber(3)
  void clearFieldType() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get adifName => $_getSZ(3);
  @$pb.TagNumber(4)
  set adifName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAdifName() => $_has(3);
  @$pb.TagNumber(4)
  void clearAdifName() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get valueRegex => $_getSZ(4);
  @$pb.TagNumber(5)
  set valueRegex($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasValueRegex() => $_has(4);
  @$pb.TagNumber(5)
  void clearValueRegex() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<$core.String> get validValues => $_getList(5);
}

class ContestMetadata extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ContestMetadata', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ApiVersion', protoName: 'ApiVersion')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Version', protoName: 'Version')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ContestName', protoName: 'ContestName')
    ..pPS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AvailableModes', protoName: 'AvailableModes')
    ..pPS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AvailableBands', protoName: 'AvailableBands')
    ..pPS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ExchangeSent', protoName: 'ExchangeSent')
    ..pPS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ExchangeRcvd', protoName: 'ExchangeRcvd')
    ..pPS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Multipliers', protoName: 'Multipliers')
    ..m<$core.String, ContestFieldInfo>(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FieldInfos', protoName: 'FieldInfos', entryClassName: 'ContestMetadata.FieldInfosEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OM, valueCreator: ContestFieldInfo.create, packageName: const $pb.PackageName('mcl'))
    ..hasRequiredFields = false
  ;

  ContestMetadata._() : super();
  factory ContestMetadata({
    $core.String? apiVersion,
    $core.String? version,
    $core.String? contestName,
    $core.Iterable<$core.String>? availableModes,
    $core.Iterable<$core.String>? availableBands,
    $core.Iterable<$core.String>? exchangeSent,
    $core.Iterable<$core.String>? exchangeRcvd,
    $core.Iterable<$core.String>? multipliers,
    $core.Map<$core.String, ContestFieldInfo>? fieldInfos,
  }) {
    final _result = create();
    if (apiVersion != null) {
      _result.apiVersion = apiVersion;
    }
    if (version != null) {
      _result.version = version;
    }
    if (contestName != null) {
      _result.contestName = contestName;
    }
    if (availableModes != null) {
      _result.availableModes.addAll(availableModes);
    }
    if (availableBands != null) {
      _result.availableBands.addAll(availableBands);
    }
    if (exchangeSent != null) {
      _result.exchangeSent.addAll(exchangeSent);
    }
    if (exchangeRcvd != null) {
      _result.exchangeRcvd.addAll(exchangeRcvd);
    }
    if (multipliers != null) {
      _result.multipliers.addAll(multipliers);
    }
    if (fieldInfos != null) {
      _result.fieldInfos.addAll(fieldInfos);
    }
    return _result;
  }
  factory ContestMetadata.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ContestMetadata.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ContestMetadata clone() => ContestMetadata()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ContestMetadata copyWith(void Function(ContestMetadata) updates) => super.copyWith((message) => updates(message as ContestMetadata)) as ContestMetadata; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ContestMetadata create() => ContestMetadata._();
  ContestMetadata createEmptyInstance() => create();
  static $pb.PbList<ContestMetadata> createRepeated() => $pb.PbList<ContestMetadata>();
  @$core.pragma('dart2js:noInline')
  static ContestMetadata getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ContestMetadata>(create);
  static ContestMetadata? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get apiVersion => $_getSZ(0);
  @$pb.TagNumber(1)
  set apiVersion($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasApiVersion() => $_has(0);
  @$pb.TagNumber(1)
  void clearApiVersion() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get version => $_getSZ(1);
  @$pb.TagNumber(2)
  set version($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasVersion() => $_has(1);
  @$pb.TagNumber(2)
  void clearVersion() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get contestName => $_getSZ(2);
  @$pb.TagNumber(3)
  set contestName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasContestName() => $_has(2);
  @$pb.TagNumber(3)
  void clearContestName() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$core.String> get availableModes => $_getList(3);

  @$pb.TagNumber(5)
  $core.List<$core.String> get availableBands => $_getList(4);

  @$pb.TagNumber(6)
  $core.List<$core.String> get exchangeSent => $_getList(5);

  @$pb.TagNumber(7)
  $core.List<$core.String> get exchangeRcvd => $_getList(6);

  @$pb.TagNumber(8)
  $core.List<$core.String> get multipliers => $_getList(7);

  @$pb.TagNumber(10)
  $core.Map<$core.String, ContestFieldInfo> get fieldInfos => $_getMap(8);
}

