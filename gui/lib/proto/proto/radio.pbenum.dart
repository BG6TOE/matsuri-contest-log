///
//  Generated code. Do not modify.
//  source: proto/radio.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class RadioMode extends $pb.ProtobufEnum {
  static const RadioMode UNKNOWN = RadioMode._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'UNKNOWN');
  static const RadioMode CW = RadioMode._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'CW');
  static const RadioMode CWR = RadioMode._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'CWR');
  static const RadioMode LSB = RadioMode._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'LSB');
  static const RadioMode USB = RadioMode._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'USB');
  static const RadioMode AM = RadioMode._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'AM');
  static const RadioMode FM = RadioMode._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'FM');
  static const RadioMode DATAL = RadioMode._(7, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DATAL');
  static const RadioMode DATAU = RadioMode._(8, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'DATAU');

  static const $core.List<RadioMode> values = <RadioMode> [
    UNKNOWN,
    CW,
    CWR,
    LSB,
    USB,
    AM,
    FM,
    DATAL,
    DATAU,
  ];

  static final $core.Map<$core.int, RadioMode> _byValue = $pb.ProtobufEnum.initByValue(values);
  static RadioMode? valueOf($core.int value) => _byValue[value];

  const RadioMode._($core.int v, $core.String n) : super(v, n);
}

