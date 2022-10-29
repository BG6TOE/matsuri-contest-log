///
//  Generated code. Do not modify.
//  source: proto/mclgui.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
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
@$core.Deprecated('Use openFileRequestDescriptor instead')
const OpenFileRequest$json = const {
  '1': 'OpenFileRequest',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `OpenFileRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List openFileRequestDescriptor = $convert.base64Decode('Cg9PcGVuRmlsZVJlcXVlc3QSEgoEbmFtZRgBIAEoCVIEbmFtZQ==');
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
@$core.Deprecated('Use draftQSOMessageDescriptor instead')
const DraftQSOMessage$json = const {
  '1': 'DraftQSOMessage',
  '2': const [
    const {'1': 'uid', '3': 1, '4': 1, '5': 9, '10': 'uid'},
    const {'1': 'dx_callsign', '3': 2, '4': 1, '5': 9, '10': 'dxCallsign'},
    const {'1': 'mode', '3': 3, '4': 1, '5': 9, '10': 'mode'},
    const {'1': 'exchange_sent', '3': 4, '4': 3, '5': 11, '6': '.mcl.DraftQSOMessage.ExchangeSentEntry', '10': 'exchangeSent'},
    const {'1': 'exchange_rcvd', '3': 5, '4': 3, '5': 11, '6': '.mcl.DraftQSOMessage.ExchangeRcvdEntry', '10': 'exchangeRcvd'},
    const {'1': 'expect', '3': 6, '4': 1, '5': 9, '10': 'expect'},
  ],
  '3': const [DraftQSOMessage_ExchangeSentEntry$json, DraftQSOMessage_ExchangeRcvdEntry$json],
};

@$core.Deprecated('Use draftQSOMessageDescriptor instead')
const DraftQSOMessage_ExchangeSentEntry$json = const {
  '1': 'ExchangeSentEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

@$core.Deprecated('Use draftQSOMessageDescriptor instead')
const DraftQSOMessage_ExchangeRcvdEntry$json = const {
  '1': 'ExchangeRcvdEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `DraftQSOMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List draftQSOMessageDescriptor = $convert.base64Decode('Cg9EcmFmdFFTT01lc3NhZ2USEAoDdWlkGAEgASgJUgN1aWQSHwoLZHhfY2FsbHNpZ24YAiABKAlSCmR4Q2FsbHNpZ24SEgoEbW9kZRgDIAEoCVIEbW9kZRJLCg1leGNoYW5nZV9zZW50GAQgAygLMiYubWNsLkRyYWZ0UVNPTWVzc2FnZS5FeGNoYW5nZVNlbnRFbnRyeVIMZXhjaGFuZ2VTZW50EksKDWV4Y2hhbmdlX3JjdmQYBSADKAsyJi5tY2wuRHJhZnRRU09NZXNzYWdlLkV4Y2hhbmdlUmN2ZEVudHJ5UgxleGNoYW5nZVJjdmQSFgoGZXhwZWN0GAYgASgJUgZleHBlY3QaPwoRRXhjaGFuZ2VTZW50RW50cnkSEAoDa2V5GAEgASgJUgNrZXkSFAoFdmFsdWUYAiABKAlSBXZhbHVlOgI4ARo/ChFFeGNoYW5nZVJjdmRFbnRyeRIQCgNrZXkYASABKAlSA2tleRIUCgV2YWx1ZRgCIAEoCVIFdmFsdWU6AjgB');
