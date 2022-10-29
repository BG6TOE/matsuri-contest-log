///
//  Generated code. Do not modify.
//  source: proto/radio.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'radio.pbenum.dart';

export 'radio.pbenum.dart';

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

class RadioVFOConfig extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RadioVFOConfig', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..e<RadioMode>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'mode', $pb.PbFieldType.OE, defaultOrMaker: RadioMode.UNKNOWN, valueOf: RadioMode.valueOf, enumValues: RadioMode.values)
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'frequency')
    ..hasRequiredFields = false
  ;

  RadioVFOConfig._() : super();
  factory RadioVFOConfig({
    RadioMode? mode,
    $fixnum.Int64? frequency,
  }) {
    final _result = create();
    if (mode != null) {
      _result.mode = mode;
    }
    if (frequency != null) {
      _result.frequency = frequency;
    }
    return _result;
  }
  factory RadioVFOConfig.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RadioVFOConfig.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RadioVFOConfig clone() => RadioVFOConfig()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RadioVFOConfig copyWith(void Function(RadioVFOConfig) updates) => super.copyWith((message) => updates(message as RadioVFOConfig)) as RadioVFOConfig; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RadioVFOConfig create() => RadioVFOConfig._();
  RadioVFOConfig createEmptyInstance() => create();
  static $pb.PbList<RadioVFOConfig> createRepeated() => $pb.PbList<RadioVFOConfig>();
  @$core.pragma('dart2js:noInline')
  static RadioVFOConfig getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RadioVFOConfig>(create);
  static RadioVFOConfig? _defaultInstance;

  @$pb.TagNumber(1)
  RadioMode get mode => $_getN(0);
  @$pb.TagNumber(1)
  set mode(RadioMode v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasMode() => $_has(0);
  @$pb.TagNumber(1)
  void clearMode() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get frequency => $_getI64(1);
  @$pb.TagNumber(2)
  set frequency($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFrequency() => $_has(1);
  @$pb.TagNumber(2)
  void clearFrequency() => clearField(2);
}

class RadioModeConfig extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RadioModeConfig', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOM<RadioVFOConfig>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'rx', subBuilder: RadioVFOConfig.create)
    ..aOM<RadioVFOConfig>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'tx', subBuilder: RadioVFOConfig.create)
    ..hasRequiredFields = false
  ;

  RadioModeConfig._() : super();
  factory RadioModeConfig({
    RadioVFOConfig? rx,
    RadioVFOConfig? tx,
  }) {
    final _result = create();
    if (rx != null) {
      _result.rx = rx;
    }
    if (tx != null) {
      _result.tx = tx;
    }
    return _result;
  }
  factory RadioModeConfig.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RadioModeConfig.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RadioModeConfig clone() => RadioModeConfig()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RadioModeConfig copyWith(void Function(RadioModeConfig) updates) => super.copyWith((message) => updates(message as RadioModeConfig)) as RadioModeConfig; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RadioModeConfig create() => RadioModeConfig._();
  RadioModeConfig createEmptyInstance() => create();
  static $pb.PbList<RadioModeConfig> createRepeated() => $pb.PbList<RadioModeConfig>();
  @$core.pragma('dart2js:noInline')
  static RadioModeConfig getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RadioModeConfig>(create);
  static RadioModeConfig? _defaultInstance;

  @$pb.TagNumber(1)
  RadioVFOConfig get rx => $_getN(0);
  @$pb.TagNumber(1)
  set rx(RadioVFOConfig v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRx() => $_has(0);
  @$pb.TagNumber(1)
  void clearRx() => clearField(1);
  @$pb.TagNumber(1)
  RadioVFOConfig ensureRx() => $_ensure(0);

  @$pb.TagNumber(2)
  RadioVFOConfig get tx => $_getN(1);
  @$pb.TagNumber(2)
  set tx(RadioVFOConfig v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTx() => $_has(1);
  @$pb.TagNumber(2)
  void clearTx() => clearField(2);
  @$pb.TagNumber(2)
  RadioVFOConfig ensureTx() => $_ensure(1);
}

enum RadioCommand_Op {
  sendCwMessage, 
  setCwSpeed, 
  setRadioBandMode, 
  abortTx, 
  notSet
}

class RadioCommand extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, RadioCommand_Op> _RadioCommand_OpByTag = {
    2 : RadioCommand_Op.sendCwMessage,
    3 : RadioCommand_Op.setCwSpeed,
    4 : RadioCommand_Op.setRadioBandMode,
    5 : RadioCommand_Op.abortTx,
    0 : RadioCommand_Op.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RadioCommand', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5])
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'channel', $pb.PbFieldType.O3)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sendCwMessage')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'setCwSpeed', $pb.PbFieldType.O3)
    ..aOM<RadioModeConfig>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'setRadioBandMode', subBuilder: RadioModeConfig.create)
    ..aOB(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'abortTx')
    ..hasRequiredFields = false
  ;

  RadioCommand._() : super();
  factory RadioCommand({
    $core.int? channel,
    $core.String? sendCwMessage,
    $core.int? setCwSpeed,
    RadioModeConfig? setRadioBandMode,
    $core.bool? abortTx,
  }) {
    final _result = create();
    if (channel != null) {
      _result.channel = channel;
    }
    if (sendCwMessage != null) {
      _result.sendCwMessage = sendCwMessage;
    }
    if (setCwSpeed != null) {
      _result.setCwSpeed = setCwSpeed;
    }
    if (setRadioBandMode != null) {
      _result.setRadioBandMode = setRadioBandMode;
    }
    if (abortTx != null) {
      _result.abortTx = abortTx;
    }
    return _result;
  }
  factory RadioCommand.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RadioCommand.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RadioCommand clone() => RadioCommand()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RadioCommand copyWith(void Function(RadioCommand) updates) => super.copyWith((message) => updates(message as RadioCommand)) as RadioCommand; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RadioCommand create() => RadioCommand._();
  RadioCommand createEmptyInstance() => create();
  static $pb.PbList<RadioCommand> createRepeated() => $pb.PbList<RadioCommand>();
  @$core.pragma('dart2js:noInline')
  static RadioCommand getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RadioCommand>(create);
  static RadioCommand? _defaultInstance;

  RadioCommand_Op whichOp() => _RadioCommand_OpByTag[$_whichOneof(0)]!;
  void clearOp() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.int get channel => $_getIZ(0);
  @$pb.TagNumber(1)
  set channel($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasChannel() => $_has(0);
  @$pb.TagNumber(1)
  void clearChannel() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get sendCwMessage => $_getSZ(1);
  @$pb.TagNumber(2)
  set sendCwMessage($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSendCwMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearSendCwMessage() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get setCwSpeed => $_getIZ(2);
  @$pb.TagNumber(3)
  set setCwSpeed($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasSetCwSpeed() => $_has(2);
  @$pb.TagNumber(3)
  void clearSetCwSpeed() => clearField(3);

  @$pb.TagNumber(4)
  RadioModeConfig get setRadioBandMode => $_getN(3);
  @$pb.TagNumber(4)
  set setRadioBandMode(RadioModeConfig v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasSetRadioBandMode() => $_has(3);
  @$pb.TagNumber(4)
  void clearSetRadioBandMode() => clearField(4);
  @$pb.TagNumber(4)
  RadioModeConfig ensureSetRadioBandMode() => $_ensure(3);

  @$pb.TagNumber(5)
  $core.bool get abortTx => $_getBF(4);
  @$pb.TagNumber(5)
  set abortTx($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasAbortTx() => $_has(4);
  @$pb.TagNumber(5)
  void clearAbortTx() => clearField(5);
}

class RadioSelector extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RadioSelector', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'channel', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  RadioSelector._() : super();
  factory RadioSelector({
    $core.int? channel,
  }) {
    final _result = create();
    if (channel != null) {
      _result.channel = channel;
    }
    return _result;
  }
  factory RadioSelector.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RadioSelector.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RadioSelector clone() => RadioSelector()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RadioSelector copyWith(void Function(RadioSelector) updates) => super.copyWith((message) => updates(message as RadioSelector)) as RadioSelector; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RadioSelector create() => RadioSelector._();
  RadioSelector createEmptyInstance() => create();
  static $pb.PbList<RadioSelector> createRepeated() => $pb.PbList<RadioSelector>();
  @$core.pragma('dart2js:noInline')
  static RadioSelector getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RadioSelector>(create);
  static RadioSelector? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get channel => $_getIZ(0);
  @$pb.TagNumber(1)
  set channel($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasChannel() => $_has(0);
  @$pb.TagNumber(1)
  void clearChannel() => clearField(1);
}

class AudioDevice extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AudioDevice', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'deviceName')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sampleRate')
    ..hasRequiredFields = false
  ;

  AudioDevice._() : super();
  factory AudioDevice({
    $core.String? deviceName,
    $core.String? sampleRate,
  }) {
    final _result = create();
    if (deviceName != null) {
      _result.deviceName = deviceName;
    }
    if (sampleRate != null) {
      _result.sampleRate = sampleRate;
    }
    return _result;
  }
  factory AudioDevice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AudioDevice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AudioDevice clone() => AudioDevice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AudioDevice copyWith(void Function(AudioDevice) updates) => super.copyWith((message) => updates(message as AudioDevice)) as AudioDevice; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AudioDevice create() => AudioDevice._();
  AudioDevice createEmptyInstance() => create();
  static $pb.PbList<AudioDevice> createRepeated() => $pb.PbList<AudioDevice>();
  @$core.pragma('dart2js:noInline')
  static AudioDevice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AudioDevice>(create);
  static AudioDevice? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get deviceName => $_getSZ(0);
  @$pb.TagNumber(1)
  set deviceName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDeviceName() => $_has(0);
  @$pb.TagNumber(1)
  void clearDeviceName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get sampleRate => $_getSZ(1);
  @$pb.TagNumber(2)
  set sampleRate($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSampleRate() => $_has(1);
  @$pb.TagNumber(2)
  void clearSampleRate() => clearField(2);
}

class RadioConfig extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RadioConfig', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'channel', $pb.PbFieldType.O3)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'model')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'uri')
    ..aOM<AudioDevice>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'audioInput', subBuilder: AudioDevice.create)
    ..aOM<AudioDevice>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'audioOutput', subBuilder: AudioDevice.create)
    ..hasRequiredFields = false
  ;

  RadioConfig._() : super();
  factory RadioConfig({
    $core.int? channel,
    $core.String? model,
    $core.String? uri,
    AudioDevice? audioInput,
    AudioDevice? audioOutput,
  }) {
    final _result = create();
    if (channel != null) {
      _result.channel = channel;
    }
    if (model != null) {
      _result.model = model;
    }
    if (uri != null) {
      _result.uri = uri;
    }
    if (audioInput != null) {
      _result.audioInput = audioInput;
    }
    if (audioOutput != null) {
      _result.audioOutput = audioOutput;
    }
    return _result;
  }
  factory RadioConfig.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RadioConfig.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RadioConfig clone() => RadioConfig()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RadioConfig copyWith(void Function(RadioConfig) updates) => super.copyWith((message) => updates(message as RadioConfig)) as RadioConfig; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RadioConfig create() => RadioConfig._();
  RadioConfig createEmptyInstance() => create();
  static $pb.PbList<RadioConfig> createRepeated() => $pb.PbList<RadioConfig>();
  @$core.pragma('dart2js:noInline')
  static RadioConfig getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RadioConfig>(create);
  static RadioConfig? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get channel => $_getIZ(0);
  @$pb.TagNumber(1)
  set channel($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasChannel() => $_has(0);
  @$pb.TagNumber(1)
  void clearChannel() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get model => $_getSZ(1);
  @$pb.TagNumber(2)
  set model($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasModel() => $_has(1);
  @$pb.TagNumber(2)
  void clearModel() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get uri => $_getSZ(2);
  @$pb.TagNumber(3)
  set uri($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasUri() => $_has(2);
  @$pb.TagNumber(3)
  void clearUri() => clearField(3);

  @$pb.TagNumber(4)
  AudioDevice get audioInput => $_getN(3);
  @$pb.TagNumber(4)
  set audioInput(AudioDevice v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasAudioInput() => $_has(3);
  @$pb.TagNumber(4)
  void clearAudioInput() => clearField(4);
  @$pb.TagNumber(4)
  AudioDevice ensureAudioInput() => $_ensure(3);

  @$pb.TagNumber(5)
  AudioDevice get audioOutput => $_getN(4);
  @$pb.TagNumber(5)
  set audioOutput(AudioDevice v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasAudioOutput() => $_has(4);
  @$pb.TagNumber(5)
  void clearAudioOutput() => clearField(5);
  @$pb.TagNumber(5)
  AudioDevice ensureAudioOutput() => $_ensure(4);
}

class AudioDeviceList extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'AudioDeviceList', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..pc<AudioDevice>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'inputDevices', $pb.PbFieldType.PM, subBuilder: AudioDevice.create)
    ..pc<AudioDevice>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'outputDevices', $pb.PbFieldType.PM, subBuilder: AudioDevice.create)
    ..hasRequiredFields = false
  ;

  AudioDeviceList._() : super();
  factory AudioDeviceList({
    $core.Iterable<AudioDevice>? inputDevices,
    $core.Iterable<AudioDevice>? outputDevices,
  }) {
    final _result = create();
    if (inputDevices != null) {
      _result.inputDevices.addAll(inputDevices);
    }
    if (outputDevices != null) {
      _result.outputDevices.addAll(outputDevices);
    }
    return _result;
  }
  factory AudioDeviceList.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AudioDeviceList.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AudioDeviceList clone() => AudioDeviceList()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AudioDeviceList copyWith(void Function(AudioDeviceList) updates) => super.copyWith((message) => updates(message as AudioDeviceList)) as AudioDeviceList; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AudioDeviceList create() => AudioDeviceList._();
  AudioDeviceList createEmptyInstance() => create();
  static $pb.PbList<AudioDeviceList> createRepeated() => $pb.PbList<AudioDeviceList>();
  @$core.pragma('dart2js:noInline')
  static AudioDeviceList getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AudioDeviceList>(create);
  static AudioDeviceList? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<AudioDevice> get inputDevices => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<AudioDevice> get outputDevices => $_getList(1);
}

class SupportedRadioList extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SupportedRadioList', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'mcl'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'model')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'displayName')
    ..hasRequiredFields = false
  ;

  SupportedRadioList._() : super();
  factory SupportedRadioList({
    $core.String? model,
    $core.String? displayName,
  }) {
    final _result = create();
    if (model != null) {
      _result.model = model;
    }
    if (displayName != null) {
      _result.displayName = displayName;
    }
    return _result;
  }
  factory SupportedRadioList.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SupportedRadioList.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SupportedRadioList clone() => SupportedRadioList()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SupportedRadioList copyWith(void Function(SupportedRadioList) updates) => super.copyWith((message) => updates(message as SupportedRadioList)) as SupportedRadioList; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SupportedRadioList create() => SupportedRadioList._();
  SupportedRadioList createEmptyInstance() => create();
  static $pb.PbList<SupportedRadioList> createRepeated() => $pb.PbList<SupportedRadioList>();
  @$core.pragma('dart2js:noInline')
  static SupportedRadioList getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SupportedRadioList>(create);
  static SupportedRadioList? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get model => $_getSZ(0);
  @$pb.TagNumber(1)
  set model($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasModel() => $_has(0);
  @$pb.TagNumber(1)
  void clearModel() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get displayName => $_getSZ(1);
  @$pb.TagNumber(2)
  set displayName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDisplayName() => $_has(1);
  @$pb.TagNumber(2)
  void clearDisplayName() => clearField(2);
}

