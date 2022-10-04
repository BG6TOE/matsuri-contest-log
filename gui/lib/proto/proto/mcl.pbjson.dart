///
//  Generated code. Do not modify.
//  source: proto/mcl.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use modeDescriptor instead')
const Mode$json = const {
  '1': 'Mode',
  '2': const [
    const {'1': 'phone', '2': 0},
    const {'1': 'phone_lsb', '2': 1},
    const {'1': 'phone_usb', '2': 2},
    const {'1': 'phone_am', '2': 3},
    const {'1': 'phone_fm', '2': 4},
    const {'1': 'cw', '2': 128},
    const {'1': 'data', '2': 129},
    const {'1': 'data_rtty', '2': 130},
    const {'1': 'data_psk', '2': 131},
    const {'1': 'data_ft8', '2': 132},
    const {'1': 'data_jt65', '2': 133},
  ],
};

/// Descriptor for `Mode`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List modeDescriptor = $convert.base64Decode('CgRNb2RlEgkKBXBob25lEAASDQoJcGhvbmVfbHNiEAESDQoJcGhvbmVfdXNiEAISDAoIcGhvbmVfYW0QAxIMCghwaG9uZV9mbRAEEgcKAmN3EIABEgkKBGRhdGEQgQESDgoJZGF0YV9ydHR5EIIBEg0KCGRhdGFfcHNrEIMBEg0KCGRhdGFfZnQ4EIQBEg4KCWRhdGFfanQ2NRCFAQ==');
@$core.Deprecated('Use qSOTypeDescriptor instead')
const QSOType$json = const {
  '1': 'QSOType',
  '2': const [
    const {'1': 'qso', '2': 0},
    const {'1': 'xqso', '2': 1},
  ],
};

/// Descriptor for `QSOType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List qSOTypeDescriptor = $convert.base64Decode('CgdRU09UeXBlEgcKA3FzbxAAEggKBHhxc28QAQ==');
@$core.Deprecated('Use qSOOperationTypeDescriptor instead')
const QSOOperationType$json = const {
  '1': 'QSOOperationType',
  '2': const [
    const {'1': 'add_or_update', '2': 0},
    const {'1': 'delete', '2': 2},
  ],
};

/// Descriptor for `QSOOperationType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List qSOOperationTypeDescriptor = $convert.base64Decode('ChBRU09PcGVyYXRpb25UeXBlEhEKDWFkZF9vcl91cGRhdGUQABIKCgZkZWxldGUQAg==');
@$core.Deprecated('Use resultCodeDescriptor instead')
const ResultCode$json = const {
  '1': 'ResultCode',
  '2': const [
    const {'1': 'success', '2': 0},
    const {'1': 'invalid_binlog', '2': 1},
  ],
};

/// Descriptor for `ResultCode`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List resultCodeDescriptor = $convert.base64Decode('CgpSZXN1bHRDb2RlEgsKB3N1Y2Nlc3MQABISCg5pbnZhbGlkX2JpbmxvZxAB');
@$core.Deprecated('Use stationDescriptor instead')
const Station$json = const {
  '1': 'Station',
  '2': const [
    const {'1': 'callsign', '3': 1, '4': 1, '5': 9, '10': 'callsign'},
    const {'1': 'grid', '3': 2, '4': 1, '5': 9, '10': 'grid'},
    const {'1': 'region', '3': 3, '4': 1, '5': 5, '10': 'region'},
    const {'1': 'cq_zone', '3': 4, '4': 1, '5': 5, '10': 'cqZone'},
    const {'1': 'itu_zone', '3': 5, '4': 1, '5': 5, '10': 'ituZone'},
    const {'1': 'dxcc_id', '3': 6, '4': 1, '5': 5, '10': 'dxccId'},
  ],
};

/// Descriptor for `Station`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List stationDescriptor = $convert.base64Decode('CgdTdGF0aW9uEhoKCGNhbGxzaWduGAEgASgJUghjYWxsc2lnbhISCgRncmlkGAIgASgJUgRncmlkEhYKBnJlZ2lvbhgDIAEoBVIGcmVnaW9uEhcKB2NxX3pvbmUYBCABKAVSBmNxWm9uZRIZCghpdHVfem9uZRgFIAEoBVIHaXR1Wm9uZRIXCgdkeGNjX2lkGAYgASgFUgZkeGNjSWQ=');
@$core.Deprecated('Use contestDescriptor instead')
const Contest$json = const {
  '1': 'Contest',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'uid', '3': 2, '4': 1, '5': 9, '10': 'uid'},
    const {'1': 'filename', '3': 3, '4': 1, '5': 9, '10': 'filename'},
    const {'1': 'category', '3': 4, '4': 1, '5': 9, '10': 'category'},
    const {'1': 'begin_timestamp', '3': 5, '4': 1, '5': 3, '10': 'beginTimestamp'},
    const {'1': 'end_timestamp', '3': 6, '4': 1, '5': 3, '10': 'endTimestamp'},
  ],
};

/// Descriptor for `Contest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List contestDescriptor = $convert.base64Decode('CgdDb250ZXN0EhIKBG5hbWUYASABKAlSBG5hbWUSEAoDdWlkGAIgASgJUgN1aWQSGgoIZmlsZW5hbWUYAyABKAlSCGZpbGVuYW1lEhoKCGNhdGVnb3J5GAQgASgJUghjYXRlZ29yeRInCg9iZWdpbl90aW1lc3RhbXAYBSABKANSDmJlZ2luVGltZXN0YW1wEiMKDWVuZF90aW1lc3RhbXAYBiABKANSDGVuZFRpbWVzdGFtcA==');
@$core.Deprecated('Use contestManifestDescriptor instead')
const ContestManifest$json = const {
  '1': 'ContestManifest',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'display_name', '3': 2, '4': 1, '5': 9, '10': 'displayName'},
    const {'1': 'exchange_data', '3': 3, '4': 3, '5': 9, '10': 'exchangeData'},
  ],
};

/// Descriptor for `ContestManifest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List contestManifestDescriptor = $convert.base64Decode('Cg9Db250ZXN0TWFuaWZlc3QSEgoEbmFtZRgBIAEoCVIEbmFtZRIhCgxkaXNwbGF5X25hbWUYAiABKAlSC2Rpc3BsYXlOYW1lEiMKDWV4Y2hhbmdlX2RhdGEYAyADKAlSDGV4Y2hhbmdlRGF0YQ==');
@$core.Deprecated('Use qSODescriptor instead')
const QSO$json = const {
  '1': 'QSO',
  '2': const [
    const {'1': 'uid', '3': 1, '4': 1, '5': 9, '10': 'uid'},
    const {'1': 'local_callsign', '3': 2, '4': 1, '5': 9, '10': 'localCallsign'},
    const {'1': 'dx_callsign', '3': 3, '4': 1, '5': 9, '10': 'dxCallsign'},
    const {'1': 'time', '3': 4, '4': 1, '5': 3, '10': 'time'},
    const {'1': 'freq', '3': 5, '4': 1, '5': 3, '10': 'freq'},
    const {'1': 'mode', '3': 6, '4': 1, '5': 14, '6': '.mcl.Mode', '10': 'mode'},
    const {'1': 'is_satellite', '3': 7, '4': 1, '5': 8, '10': 'isSatellite'},
    const {'1': 'exch_sent', '3': 8, '4': 3, '5': 9, '10': 'exchSent'},
    const {'1': 'exch_rcvd', '3': 9, '4': 3, '5': 9, '10': 'exchRcvd'},
    const {'1': 'type', '3': 10, '4': 1, '5': 14, '6': '.mcl.QSOType', '10': 'type'},
    const {'1': 'operator', '3': 11, '4': 1, '5': 9, '10': 'operator'},
  ],
};

/// Descriptor for `QSO`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List qSODescriptor = $convert.base64Decode('CgNRU08SEAoDdWlkGAEgASgJUgN1aWQSJQoObG9jYWxfY2FsbHNpZ24YAiABKAlSDWxvY2FsQ2FsbHNpZ24SHwoLZHhfY2FsbHNpZ24YAyABKAlSCmR4Q2FsbHNpZ24SEgoEdGltZRgEIAEoA1IEdGltZRISCgRmcmVxGAUgASgDUgRmcmVxEh0KBG1vZGUYBiABKA4yCS5tY2wuTW9kZVIEbW9kZRIhCgxpc19zYXRlbGxpdGUYByABKAhSC2lzU2F0ZWxsaXRlEhsKCWV4Y2hfc2VudBgIIAMoCVIIZXhjaFNlbnQSGwoJZXhjaF9yY3ZkGAkgAygJUghleGNoUmN2ZBIgCgR0eXBlGAogASgOMgwubWNsLlFTT1R5cGVSBHR5cGUSGgoIb3BlcmF0b3IYCyABKAlSCG9wZXJhdG9y');
@$core.Deprecated('Use qSOOpDescriptor instead')
const QSOOp$json = const {
  '1': 'QSOOp',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.mcl.QSOOperationType', '10': 'type'},
    const {'1': 'uid', '3': 2, '4': 1, '5': 9, '10': 'uid'},
  ],
};

/// Descriptor for `QSOOp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List qSOOpDescriptor = $convert.base64Decode('CgVRU09PcBIpCgR0eXBlGAEgASgOMhUubWNsLlFTT09wZXJhdGlvblR5cGVSBHR5cGUSEAoDdWlkGAIgASgJUgN1aWQ=');
@$core.Deprecated('Use binlogMessageDescriptor instead')
const BinlogMessage$json = const {
  '1': 'BinlogMessage',
  '2': const [
    const {'1': 'serial', '3': 1, '4': 1, '5': 4, '10': 'serial'},
    const {'1': 'qso_op', '3': 4, '4': 1, '5': 11, '6': '.mcl.QSOOp', '9': 0, '10': 'qsoOp'},
    const {'1': 'qso', '3': 5, '4': 1, '5': 11, '6': '.mcl.QSO', '9': 0, '10': 'qso'},
  ],
  '8': const [
    const {'1': 'message'},
  ],
};

/// Descriptor for `BinlogMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List binlogMessageDescriptor = $convert.base64Decode('Cg1CaW5sb2dNZXNzYWdlEhYKBnNlcmlhbBgBIAEoBFIGc2VyaWFsEiMKBnFzb19vcBgEIAEoCzIKLm1jbC5RU09PcEgAUgVxc29PcBIcCgNxc28YBSABKAsyCC5tY2wuUVNPSABSA3Fzb0IJCgdtZXNzYWdl');
@$core.Deprecated('Use snapshotMessageDescriptor instead')
const SnapshotMessage$json = const {
  '1': 'SnapshotMessage',
  '2': const [
    const {'1': 'qso', '3': 1, '4': 3, '5': 11, '6': '.mcl.QSO', '10': 'qso'},
    const {'1': 'serial', '3': 2, '4': 1, '5': 4, '10': 'serial'},
  ],
};

/// Descriptor for `SnapshotMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotMessageDescriptor = $convert.base64Decode('Cg9TbmFwc2hvdE1lc3NhZ2USGgoDcXNvGAEgAygLMggubWNsLlFTT1IDcXNvEhYKBnNlcmlhbBgCIAEoBFIGc2VyaWFs');
@$core.Deprecated('Use binlogMessageSetDescriptor instead')
const BinlogMessageSet$json = const {
  '1': 'BinlogMessageSet',
  '2': const [
    const {'1': 'messages', '3': 1, '4': 3, '5': 11, '6': '.mcl.BinlogMessage', '10': 'messages'},
  ],
};

/// Descriptor for `BinlogMessageSet`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List binlogMessageSetDescriptor = $convert.base64Decode('ChBCaW5sb2dNZXNzYWdlU2V0Ei4KCG1lc3NhZ2VzGAEgAygLMhIubWNsLkJpbmxvZ01lc3NhZ2VSCG1lc3NhZ2Vz');
@$core.Deprecated('Use standardResponseDescriptor instead')
const StandardResponse$json = const {
  '1': 'StandardResponse',
  '2': const [
    const {'1': 'result_code', '3': 1, '4': 1, '5': 14, '6': '.mcl.ResultCode', '10': 'resultCode'},
    const {'1': 'error_message', '3': 2, '4': 1, '5': 9, '10': 'errorMessage'},
  ],
};

/// Descriptor for `StandardResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List standardResponseDescriptor = $convert.base64Decode('ChBTdGFuZGFyZFJlc3BvbnNlEjAKC3Jlc3VsdF9jb2RlGAEgASgOMg8ubWNsLlJlc3VsdENvZGVSCnJlc3VsdENvZGUSIwoNZXJyb3JfbWVzc2FnZRgCIAEoCVIMZXJyb3JNZXNzYWdl');
@$core.Deprecated('Use retrieveBinlogRequestDescriptor instead')
const RetrieveBinlogRequest$json = const {
  '1': 'RetrieveBinlogRequest',
  '2': const [
    const {'1': 'serial_start', '3': 1, '4': 1, '5': 4, '10': 'serialStart'},
    const {'1': 'serial_end', '3': 2, '4': 1, '5': 4, '10': 'serialEnd'},
  ],
};

/// Descriptor for `RetrieveBinlogRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List retrieveBinlogRequestDescriptor = $convert.base64Decode('ChVSZXRyaWV2ZUJpbmxvZ1JlcXVlc3QSIQoMc2VyaWFsX3N0YXJ0GAEgASgEUgtzZXJpYWxTdGFydBIdCgpzZXJpYWxfZW5kGAIgASgEUglzZXJpYWxFbmQ=');
