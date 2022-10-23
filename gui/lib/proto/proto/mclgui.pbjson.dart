///
//  Generated code. Do not modify.
//  source: proto/mclgui.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use radioModeDescriptor instead')
const RadioMode$json = const {
  '1': 'RadioMode',
  '2': const [
    const {'1': 'UNKNOWN', '2': 0},
    const {'1': 'CW', '2': 1},
    const {'1': 'CWR', '2': 2},
    const {'1': 'LSB', '2': 3},
    const {'1': 'USB', '2': 4},
    const {'1': 'AM', '2': 5},
    const {'1': 'FM', '2': 6},
    const {'1': 'DATAL', '2': 7},
    const {'1': 'DATAU', '2': 8},
  ],
};

/// Descriptor for `RadioMode`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List radioModeDescriptor = $convert.base64Decode('CglSYWRpb01vZGUSCwoHVU5LTk9XThAAEgYKAkNXEAESBwoDQ1dSEAISBwoDTFNCEAMSBwoDVVNCEAQSBgoCQU0QBRIGCgJGTRAGEgkKBURBVEFMEAcSCQoFREFUQVUQCA==');
@$core.Deprecated('Use qSOFieldDescriptor instead')
const QSOField$json = const {
  '1': 'QSOField',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `QSOField`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List qSOFieldDescriptor = $convert.base64Decode('CghRU09GaWVsZBIUCgV0aXRsZRgBIAEoCVIFdGl0bGUSFAoFdmFsdWUYAiABKAlSBXZhbHVl');
@$core.Deprecated('Use draftQSOMessageDescriptor instead')
const DraftQSOMessage$json = const {
  '1': 'DraftQSOMessage',
  '2': const [
    const {'1': 'dx_callsign', '3': 1, '4': 1, '5': 9, '10': 'dxCallsign'},
  ],
};

/// Descriptor for `DraftQSOMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List draftQSOMessageDescriptor = $convert.base64Decode('Cg9EcmFmdFFTT01lc3NhZ2USHwoLZHhfY2FsbHNpZ24YASABKAlSCmR4Q2FsbHNpZ24=');
@$core.Deprecated('Use scoreResponseDescriptor instead')
const ScoreResponse$json = const {
  '1': 'ScoreResponse',
  '2': const [
    const {'1': 'category_score', '3': 1, '4': 3, '5': 11, '6': '.mcl.ScoreResponse.CategoryScoreEntry', '10': 'categoryScore'},
  ],
  '3': const [ScoreResponse_CategoryScoreEntry$json],
};

@$core.Deprecated('Use scoreResponseDescriptor instead')
const ScoreResponse_CategoryScoreEntry$json = const {
  '1': 'CategoryScoreEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 3, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `ScoreResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List scoreResponseDescriptor = $convert.base64Decode('Cg1TY29yZVJlc3BvbnNlEkwKDmNhdGVnb3J5X3Njb3JlGAEgAygLMiUubWNsLlNjb3JlUmVzcG9uc2UuQ2F0ZWdvcnlTY29yZUVudHJ5Ug1jYXRlZ29yeVNjb3JlGkAKEkNhdGVnb3J5U2NvcmVFbnRyeRIQCgNrZXkYASABKAlSA2tleRIUCgV2YWx1ZRgCIAEoA1IFdmFsdWU6AjgB');
@$core.Deprecated('Use qSOFieldsDescriptor instead')
const QSOFields$json = const {
  '1': 'QSOFields',
  '2': const [
    const {'1': 'exchange_sent', '3': 1, '4': 3, '5': 11, '6': '.mcl.QSOField', '10': 'exchangeSent'},
    const {'1': 'exchange_rcvd', '3': 2, '4': 3, '5': 11, '6': '.mcl.QSOField', '10': 'exchangeRcvd'},
  ],
};

/// Descriptor for `QSOFields`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List qSOFieldsDescriptor = $convert.base64Decode('CglRU09GaWVsZHMSMgoNZXhjaGFuZ2Vfc2VudBgBIAMoCzINLm1jbC5RU09GaWVsZFIMZXhjaGFuZ2VTZW50EjIKDWV4Y2hhbmdlX3JjdmQYAiADKAsyDS5tY2wuUVNPRmllbGRSDGV4Y2hhbmdlUmN2ZA==');
@$core.Deprecated('Use spotDescriptor instead')
const Spot$json = const {
  '1': 'Spot',
  '2': const [
    const {'1': 'dx_callsign', '3': 1, '4': 1, '5': 9, '10': 'dxCallsign'},
    const {'1': 'de_callsign', '3': 2, '4': 1, '5': 9, '10': 'deCallsign'},
    const {'1': 'spot_timestamp', '3': 3, '4': 1, '5': 3, '10': 'spotTimestamp'},
    const {'1': 'spot_frequency', '3': 4, '4': 1, '5': 3, '10': 'spotFrequency'},
    const {'1': 'memo', '3': 5, '4': 1, '5': 9, '10': 'memo'},
  ],
};

/// Descriptor for `Spot`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List spotDescriptor = $convert.base64Decode('CgRTcG90Eh8KC2R4X2NhbGxzaWduGAEgASgJUgpkeENhbGxzaWduEh8KC2RlX2NhbGxzaWduGAIgASgJUgpkZUNhbGxzaWduEiUKDnNwb3RfdGltZXN0YW1wGAMgASgDUg1zcG90VGltZXN0YW1wEiUKDnNwb3RfZnJlcXVlbmN5GAQgASgDUg1zcG90RnJlcXVlbmN5EhIKBG1lbW8YBSABKAlSBG1lbW8=');
@$core.Deprecated('Use callsignLookupResultCatrgoryDescriptor instead')
const CallsignLookupResultCatrgory$json = const {
  '1': 'CallsignLookupResultCatrgory',
  '2': const [
    const {'1': 'possible_callsigns', '3': 1, '4': 3, '5': 9, '10': 'possibleCallsigns'},
  ],
};

/// Descriptor for `CallsignLookupResultCatrgory`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List callsignLookupResultCatrgoryDescriptor = $convert.base64Decode('ChxDYWxsc2lnbkxvb2t1cFJlc3VsdENhdHJnb3J5Ei0KEnBvc3NpYmxlX2NhbGxzaWducxgBIAMoCVIRcG9zc2libGVDYWxsc2lnbnM=');
@$core.Deprecated('Use bandStatusDescriptor instead')
const BandStatus$json = const {
  '1': 'BandStatus',
  '2': const [
    const {'1': 'band_status', '3': 1, '4': 3, '5': 11, '6': '.mcl.BandStatus.BandStatusEntry', '10': 'bandStatus'},
  ],
  '3': const [BandStatus_BandStatusEntry$json],
};

@$core.Deprecated('Use bandStatusDescriptor instead')
const BandStatus_BandStatusEntry$json = const {
  '1': 'BandStatusEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 3, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `BandStatus`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bandStatusDescriptor = $convert.base64Decode('CgpCYW5kU3RhdHVzEkAKC2JhbmRfc3RhdHVzGAEgAygLMh8ubWNsLkJhbmRTdGF0dXMuQmFuZFN0YXR1c0VudHJ5UgpiYW5kU3RhdHVzGj0KD0JhbmRTdGF0dXNFbnRyeRIQCgNrZXkYASABKANSA2tleRIUCgV2YWx1ZRgCIAEoCVIFdmFsdWU6AjgB');
@$core.Deprecated('Use callsignLookupResultDescriptor instead')
const CallsignLookupResult$json = const {
  '1': 'CallsignLookupResult',
  '2': const [
    const {'1': 'database_lookup', '3': 1, '4': 3, '5': 11, '6': '.mcl.CallsignLookupResult.DatabaseLookupEntry', '10': 'databaseLookup'},
    const {'1': 'band_status', '3': 2, '4': 3, '5': 11, '6': '.mcl.CallsignLookupResult.BandStatusEntry', '10': 'bandStatus'},
  ],
  '3': const [CallsignLookupResult_DatabaseLookupEntry$json, CallsignLookupResult_BandStatusEntry$json],
};

@$core.Deprecated('Use callsignLookupResultDescriptor instead')
const CallsignLookupResult_DatabaseLookupEntry$json = const {
  '1': 'DatabaseLookupEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.mcl.CallsignLookupResultCatrgory', '10': 'value'},
  ],
  '7': const {'7': true},
};

@$core.Deprecated('Use callsignLookupResultDescriptor instead')
const CallsignLookupResult_BandStatusEntry$json = const {
  '1': 'BandStatusEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.mcl.BandStatus', '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `CallsignLookupResult`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List callsignLookupResultDescriptor = $convert.base64Decode('ChRDYWxsc2lnbkxvb2t1cFJlc3VsdBJWCg9kYXRhYmFzZV9sb29rdXAYASADKAsyLS5tY2wuQ2FsbHNpZ25Mb29rdXBSZXN1bHQuRGF0YWJhc2VMb29rdXBFbnRyeVIOZGF0YWJhc2VMb29rdXASSgoLYmFuZF9zdGF0dXMYAiADKAsyKS5tY2wuQ2FsbHNpZ25Mb29rdXBSZXN1bHQuQmFuZFN0YXR1c0VudHJ5UgpiYW5kU3RhdHVzGmQKE0RhdGFiYXNlTG9va3VwRW50cnkSEAoDa2V5GAEgASgJUgNrZXkSNwoFdmFsdWUYAiABKAsyIS5tY2wuQ2FsbHNpZ25Mb29rdXBSZXN1bHRDYXRyZ29yeVIFdmFsdWU6AjgBGk4KD0JhbmRTdGF0dXNFbnRyeRIQCgNrZXkYASABKAlSA2tleRIlCgV2YWx1ZRgCIAEoCzIPLm1jbC5CYW5kU3RhdHVzUgV2YWx1ZToCOAE=');
@$core.Deprecated('Use radioStatusDescriptor instead')
const RadioStatus$json = const {
  '1': 'RadioStatus',
  '2': const [
    const {'1': 'freq', '3': 1, '4': 1, '5': 3, '10': 'freq'},
    const {'1': 'mode', '3': 2, '4': 1, '5': 14, '6': '.mcl.RadioMode', '10': 'mode'},
    const {'1': 'tx', '3': 3, '4': 1, '5': 8, '10': 'tx'},
  ],
};

/// Descriptor for `RadioStatus`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List radioStatusDescriptor = $convert.base64Decode('CgtSYWRpb1N0YXR1cxISCgRmcmVxGAEgASgDUgRmcmVxEiIKBG1vZGUYAiABKA4yDi5tY2wuUmFkaW9Nb2RlUgRtb2RlEg4KAnR4GAMgASgIUgJ0eA==');
@$core.Deprecated('Use radioCommandsDescriptor instead')
const RadioCommands$json = const {
  '1': 'RadioCommands',
  '2': const [
    const {'1': 'send_cw_message', '3': 1, '4': 1, '5': 9, '9': 0, '10': 'sendCwMessage'},
    const {'1': 'set_cw_speed', '3': 2, '4': 1, '5': 5, '9': 0, '10': 'setCwSpeed'},
    const {'1': 'set_radio_band_mode', '3': 3, '4': 1, '5': 14, '6': '.mcl.RadioMode', '9': 0, '10': 'setRadioBandMode'},
    const {'1': 'abort', '3': 4, '4': 1, '5': 8, '9': 0, '10': 'abort'},
  ],
  '8': const [
    const {'1': 'op'},
  ],
};

/// Descriptor for `RadioCommands`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List radioCommandsDescriptor = $convert.base64Decode('Cg1SYWRpb0NvbW1hbmRzEigKD3NlbmRfY3dfbWVzc2FnZRgBIAEoCUgAUg1zZW5kQ3dNZXNzYWdlEiIKDHNldF9jd19zcGVlZBgCIAEoBUgAUgpzZXRDd1NwZWVkEj8KE3NldF9yYWRpb19iYW5kX21vZGUYAyABKA4yDi5tY2wuUmFkaW9Nb2RlSABSEHNldFJhZGlvQmFuZE1vZGUSFgoFYWJvcnQYBCABKAhIAFIFYWJvcnRCBAoCb3A=');
