///
//  Generated code. Do not modify.
//  source: proto/mcl.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Mode extends $pb.ProtobufEnum {
  static const Mode phone = Mode._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'phone');
  static const Mode phone_lsb = Mode._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'phone_lsb');
  static const Mode phone_usb = Mode._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'phone_usb');
  static const Mode phone_am = Mode._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'phone_am');
  static const Mode phone_fm = Mode._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'phone_fm');
  static const Mode cw = Mode._(128, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'cw');
  static const Mode data = Mode._(129, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'data');
  static const Mode data_rtty = Mode._(130, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'data_rtty');
  static const Mode data_psk = Mode._(131, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'data_psk');
  static const Mode data_ft8 = Mode._(132, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'data_ft8');
  static const Mode data_jt65 = Mode._(133, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'data_jt65');

  static const $core.List<Mode> values = <Mode> [
    phone,
    phone_lsb,
    phone_usb,
    phone_am,
    phone_fm,
    cw,
    data,
    data_rtty,
    data_psk,
    data_ft8,
    data_jt65,
  ];

  static final $core.Map<$core.int, Mode> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Mode? valueOf($core.int value) => _byValue[value];

  const Mode._($core.int v, $core.String n) : super(v, n);
}

class QSOType extends $pb.ProtobufEnum {
  static const QSOType qso = QSOType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'qso');
  static const QSOType xqso = QSOType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'xqso');

  static const $core.List<QSOType> values = <QSOType> [
    qso,
    xqso,
  ];

  static final $core.Map<$core.int, QSOType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static QSOType? valueOf($core.int value) => _byValue[value];

  const QSOType._($core.int v, $core.String n) : super(v, n);
}

class QSOOperationType extends $pb.ProtobufEnum {
  static const QSOOperationType add_or_update = QSOOperationType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'add_or_update');
  static const QSOOperationType delete = QSOOperationType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'delete');

  static const $core.List<QSOOperationType> values = <QSOOperationType> [
    add_or_update,
    delete,
  ];

  static final $core.Map<$core.int, QSOOperationType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static QSOOperationType? valueOf($core.int value) => _byValue[value];

  const QSOOperationType._($core.int v, $core.String n) : super(v, n);
}

class ResultCode extends $pb.ProtobufEnum {
  static const ResultCode success = ResultCode._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'success');
  static const ResultCode invalid_binlog = ResultCode._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'invalid_binlog');
  static const ResultCode invalid_access = ResultCode._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'invalid_access');
  static const ResultCode exists = ResultCode._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'exists');
  static const ResultCode internal = ResultCode._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'internal');

  static const $core.List<ResultCode> values = <ResultCode> [
    success,
    invalid_binlog,
    invalid_access,
    exists,
    internal,
  ];

  static final $core.Map<$core.int, ResultCode> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ResultCode? valueOf($core.int value) => _byValue[value];

  const ResultCode._($core.int v, $core.String n) : super(v, n);
}

