///
//  Generated code. Do not modify.
//  source: proto/common.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Continent extends $pb.ProtobufEnum {
  static const Continent unknown = Continent._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'unknown');
  static const Continent Africa = Continent._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'Africa');
  static const Continent Antarctica = Continent._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'Antarctica');
  static const Continent Asia = Continent._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'Asia');
  static const Continent Europe = Continent._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'Europe');
  static const Continent NorthAmerica = Continent._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'NorthAmerica');
  static const Continent Oceania = Continent._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'Oceania');
  static const Continent SouthAmerica = Continent._(7, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SouthAmerica');

  static const $core.List<Continent> values = <Continent> [
    unknown,
    Africa,
    Antarctica,
    Asia,
    Europe,
    NorthAmerica,
    Oceania,
    SouthAmerica,
  ];

  static final $core.Map<$core.int, Continent> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Continent? valueOf($core.int value) => _byValue[value];

  const Continent._($core.int v, $core.String n) : super(v, n);
}

