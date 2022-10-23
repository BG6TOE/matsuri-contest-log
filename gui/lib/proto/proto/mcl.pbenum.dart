///
//  Generated code. Do not modify.
//  source: proto/mcl.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

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

