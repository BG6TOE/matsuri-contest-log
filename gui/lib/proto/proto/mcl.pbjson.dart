///
//  Generated code. Do not modify.
//  source: proto/mcl.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use qSOTypeDescriptor instead')
const QSOType$json = const {
  '1': 'QSOType',
  '2': const [
    const {'1': 'qso', '2': 0},
    const {'1': 'xqso', '2': 1},
    const {'1': 'draft', '2': 2},
  ],
};

/// Descriptor for `QSOType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List qSOTypeDescriptor = $convert.base64Decode('CgdRU09UeXBlEgcKA3FzbxAAEggKBHhxc28QARIJCgVkcmFmdBAC');
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
    const {'1': 'invalid_access', '2': 2},
    const {'1': 'exists', '2': 3},
    const {'1': 'internal', '2': 4},
  ],
};

/// Descriptor for `ResultCode`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List resultCodeDescriptor = $convert.base64Decode('CgpSZXN1bHRDb2RlEgsKB3N1Y2Nlc3MQABISCg5pbnZhbGlkX2JpbmxvZxABEhIKDmludmFsaWRfYWNjZXNzEAISCgoGZXhpc3RzEAMSDAoIaW50ZXJuYWwQBA==');
@$core.Deprecated('Use stationDescriptor instead')
const Station$json = const {
  '1': 'Station',
  '2': const [
    const {'1': 'callsign', '3': 1, '4': 1, '5': 9, '10': 'callsign'},
    const {'1': 'grid', '3': 2, '4': 1, '5': 9, '10': 'grid'},
    const {'1': 'region', '3': 3, '4': 1, '5': 5, '10': 'region'},
    const {'1': 'custom_fields', '3': 4, '4': 3, '5': 11, '6': '.mcl.Station.CustomFieldsEntry', '10': 'customFields'},
  ],
  '3': const [Station_CustomFieldsEntry$json],
};

@$core.Deprecated('Use stationDescriptor instead')
const Station_CustomFieldsEntry$json = const {
  '1': 'CustomFieldsEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `Station`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List stationDescriptor = $convert.base64Decode('CgdTdGF0aW9uEhoKCGNhbGxzaWduGAEgASgJUghjYWxsc2lnbhISCgRncmlkGAIgASgJUgRncmlkEhYKBnJlZ2lvbhgDIAEoBVIGcmVnaW9uEkMKDWN1c3RvbV9maWVsZHMYBCADKAsyHi5tY2wuU3RhdGlvbi5DdXN0b21GaWVsZHNFbnRyeVIMY3VzdG9tRmllbGRzGj8KEUN1c3RvbUZpZWxkc0VudHJ5EhAKA2tleRgBIAEoCVIDa2V5EhQKBXZhbHVlGAIgASgJUgV2YWx1ZToCOAE=');
@$core.Deprecated('Use activeContestDescriptor instead')
const ActiveContest$json = const {
  '1': 'ActiveContest',
  '2': const [
    const {'1': 'uuid', '3': 1, '4': 1, '5': 9, '10': 'uuid'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'contest', '3': 3, '4': 1, '5': 11, '6': '.mcl.ContestMetadata', '10': 'contest'},
    const {'1': 'station', '3': 4, '4': 1, '5': 11, '6': '.mcl.Station', '10': 'station'},
    const {'1': 'begin_timestamp', '3': 5, '4': 1, '5': 3, '10': 'beginTimestamp'},
    const {'1': 'end_timestamp', '3': 6, '4': 1, '5': 3, '10': 'endTimestamp'},
    const {'1': 'contest_script', '3': 7, '4': 1, '5': 9, '10': 'contestScript'},
  ],
};

/// Descriptor for `ActiveContest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activeContestDescriptor = $convert.base64Decode('Cg1BY3RpdmVDb250ZXN0EhIKBHV1aWQYASABKAlSBHV1aWQSEgoEbmFtZRgCIAEoCVIEbmFtZRIuCgdjb250ZXN0GAMgASgLMhQubWNsLkNvbnRlc3RNZXRhZGF0YVIHY29udGVzdBImCgdzdGF0aW9uGAQgASgLMgwubWNsLlN0YXRpb25SB3N0YXRpb24SJwoPYmVnaW5fdGltZXN0YW1wGAUgASgDUg5iZWdpblRpbWVzdGFtcBIjCg1lbmRfdGltZXN0YW1wGAYgASgDUgxlbmRUaW1lc3RhbXASJQoOY29udGVzdF9zY3JpcHQYByABKAlSDWNvbnRlc3RTY3JpcHQ=');
@$core.Deprecated('Use qSODescriptor instead')
const QSO$json = const {
  '1': 'QSO',
  '2': const [
    const {'1': 'uid', '3': 1, '4': 1, '5': 9, '10': 'uid'},
    const {'1': 'local_callsign', '3': 2, '4': 1, '5': 9, '10': 'localCallsign'},
    const {'1': 'dx_callsign', '3': 3, '4': 1, '5': 9, '10': 'dxCallsign'},
    const {'1': 'time', '3': 4, '4': 1, '5': 3, '10': 'time'},
    const {'1': 'freq', '3': 5, '4': 1, '5': 3, '10': 'freq'},
    const {'1': 'mode', '3': 6, '4': 1, '5': 9, '10': 'mode'},
    const {'1': 'exchange_sent', '3': 7, '4': 3, '5': 11, '6': '.mcl.QSO.ExchangeSentEntry', '10': 'exchangeSent'},
    const {'1': 'exchange_rcvd', '3': 8, '4': 3, '5': 11, '6': '.mcl.QSO.ExchangeRcvdEntry', '10': 'exchangeRcvd'},
    const {'1': 'type', '3': 9, '4': 1, '5': 14, '6': '.mcl.QSOType', '10': 'type'},
    const {'1': 'operator', '3': 10, '4': 1, '5': 9, '10': 'operator'},
  ],
  '3': const [QSO_ExchangeSentEntry$json, QSO_ExchangeRcvdEntry$json],
};

@$core.Deprecated('Use qSODescriptor instead')
const QSO_ExchangeSentEntry$json = const {
  '1': 'ExchangeSentEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

@$core.Deprecated('Use qSODescriptor instead')
const QSO_ExchangeRcvdEntry$json = const {
  '1': 'ExchangeRcvdEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `QSO`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List qSODescriptor = $convert.base64Decode('CgNRU08SEAoDdWlkGAEgASgJUgN1aWQSJQoObG9jYWxfY2FsbHNpZ24YAiABKAlSDWxvY2FsQ2FsbHNpZ24SHwoLZHhfY2FsbHNpZ24YAyABKAlSCmR4Q2FsbHNpZ24SEgoEdGltZRgEIAEoA1IEdGltZRISCgRmcmVxGAUgASgDUgRmcmVxEhIKBG1vZGUYBiABKAlSBG1vZGUSPwoNZXhjaGFuZ2Vfc2VudBgHIAMoCzIaLm1jbC5RU08uRXhjaGFuZ2VTZW50RW50cnlSDGV4Y2hhbmdlU2VudBI/Cg1leGNoYW5nZV9yY3ZkGAggAygLMhoubWNsLlFTTy5FeGNoYW5nZVJjdmRFbnRyeVIMZXhjaGFuZ2VSY3ZkEiAKBHR5cGUYCSABKA4yDC5tY2wuUVNPVHlwZVIEdHlwZRIaCghvcGVyYXRvchgKIAEoCVIIb3BlcmF0b3IaPwoRRXhjaGFuZ2VTZW50RW50cnkSEAoDa2V5GAEgASgJUgNrZXkSFAoFdmFsdWUYAiABKAlSBXZhbHVlOgI4ARo/ChFFeGNoYW5nZVJjdmRFbnRyeRIQCgNrZXkYASABKAlSA2tleRIUCgV2YWx1ZRgCIAEoCVIFdmFsdWU6AjgB');
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
@$core.Deprecated('Use loadContestRequestDescriptor instead')
const LoadContestRequest$json = const {
  '1': 'LoadContestRequest',
  '2': const [
    const {'1': 'database_name', '3': 1, '4': 1, '5': 9, '10': 'databaseName'},
  ],
};

/// Descriptor for `LoadContestRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loadContestRequestDescriptor = $convert.base64Decode('ChJMb2FkQ29udGVzdFJlcXVlc3QSIwoNZGF0YWJhc2VfbmFtZRgBIAEoCVIMZGF0YWJhc2VOYW1l');
@$core.Deprecated('Use createContestRequestDescriptor instead')
const CreateContestRequest$json = const {
  '1': 'CreateContestRequest',
  '2': const [
    const {'1': 'database_name', '3': 1, '4': 1, '5': 9, '10': 'databaseName'},
    const {'1': 'contest', '3': 2, '4': 1, '5': 11, '6': '.mcl.ActiveContest', '10': 'contest'},
  ],
};

/// Descriptor for `CreateContestRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createContestRequestDescriptor = $convert.base64Decode('ChRDcmVhdGVDb250ZXN0UmVxdWVzdBIjCg1kYXRhYmFzZV9uYW1lGAEgASgJUgxkYXRhYmFzZU5hbWUSLAoHY29udGVzdBgCIAEoCzISLm1jbC5BY3RpdmVDb250ZXN0Ugdjb250ZXN0');
@$core.Deprecated('Use parseContestRequestDescriptor instead')
const ParseContestRequest$json = const {
  '1': 'ParseContestRequest',
  '2': const [
    const {'1': 'contest_descriptor', '3': 1, '4': 1, '5': 9, '10': 'contestDescriptor'},
  ],
};

/// Descriptor for `ParseContestRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List parseContestRequestDescriptor = $convert.base64Decode('ChNQYXJzZUNvbnRlc3RSZXF1ZXN0Ei0KEmNvbnRlc3RfZGVzY3JpcHRvchgBIAEoCVIRY29udGVzdERlc2NyaXB0b3I=');
