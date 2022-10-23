///
//  Generated code. Do not modify.
//  source: proto/common.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use continentDescriptor instead')
const Continent$json = const {
  '1': 'Continent',
  '2': const [
    const {'1': 'unknown', '2': 0},
    const {'1': 'Africa', '2': 1},
    const {'1': 'Antarctica', '2': 2},
    const {'1': 'Asia', '2': 3},
    const {'1': 'Europe', '2': 4},
    const {'1': 'NorthAmerica', '2': 5},
    const {'1': 'Oceania', '2': 6},
    const {'1': 'SouthAmerica', '2': 7},
  ],
};

/// Descriptor for `Continent`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List continentDescriptor = $convert.base64Decode('CglDb250aW5lbnQSCwoHdW5rbm93bhAAEgoKBkFmcmljYRABEg4KCkFudGFyY3RpY2EQAhIICgRBc2lhEAMSCgoGRXVyb3BlEAQSEAoMTm9ydGhBbWVyaWNhEAUSCwoHT2NlYW5pYRAGEhAKDFNvdXRoQW1lcmljYRAH');
@$core.Deprecated('Use contestFieldInfoDescriptor instead')
const ContestFieldInfo$json = const {
  '1': 'ContestFieldInfo',
  '2': const [
    const {'1': 'DisplayName', '3': 1, '4': 1, '5': 9, '10': 'DisplayName'},
    const {'1': 'Description', '3': 2, '4': 1, '5': 9, '10': 'Description'},
    const {'1': 'FieldType', '3': 3, '4': 1, '5': 9, '10': 'FieldType'},
    const {'1': 'AdifName', '3': 4, '4': 1, '5': 9, '10': 'AdifName'},
    const {'1': 'ValueRegex', '3': 5, '4': 1, '5': 9, '10': 'ValueRegex'},
    const {'1': 'ValidValues', '3': 6, '4': 3, '5': 9, '10': 'ValidValues'},
  ],
};

/// Descriptor for `ContestFieldInfo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List contestFieldInfoDescriptor = $convert.base64Decode('ChBDb250ZXN0RmllbGRJbmZvEiAKC0Rpc3BsYXlOYW1lGAEgASgJUgtEaXNwbGF5TmFtZRIgCgtEZXNjcmlwdGlvbhgCIAEoCVILRGVzY3JpcHRpb24SHAoJRmllbGRUeXBlGAMgASgJUglGaWVsZFR5cGUSGgoIQWRpZk5hbWUYBCABKAlSCEFkaWZOYW1lEh4KClZhbHVlUmVnZXgYBSABKAlSClZhbHVlUmVnZXgSIAoLVmFsaWRWYWx1ZXMYBiADKAlSC1ZhbGlkVmFsdWVz');
@$core.Deprecated('Use contestMetadataDescriptor instead')
const ContestMetadata$json = const {
  '1': 'ContestMetadata',
  '2': const [
    const {'1': 'ApiVersion', '3': 1, '4': 1, '5': 9, '10': 'ApiVersion'},
    const {'1': 'Version', '3': 2, '4': 1, '5': 9, '10': 'Version'},
    const {'1': 'ContestName', '3': 3, '4': 1, '5': 9, '10': 'ContestName'},
    const {'1': 'AvailableModes', '3': 4, '4': 3, '5': 9, '10': 'AvailableModes'},
    const {'1': 'AvailableBands', '3': 5, '4': 3, '5': 9, '10': 'AvailableBands'},
    const {'1': 'ExchangeSent', '3': 6, '4': 3, '5': 9, '10': 'ExchangeSent'},
    const {'1': 'ExchangeRcvd', '3': 7, '4': 3, '5': 9, '10': 'ExchangeRcvd'},
    const {'1': 'Multipliers', '3': 8, '4': 3, '5': 9, '10': 'Multipliers'},
    const {'1': 'FieldInfos', '3': 10, '4': 3, '5': 11, '6': '.mcl.ContestMetadata.FieldInfosEntry', '10': 'FieldInfos'},
  ],
  '3': const [ContestMetadata_FieldInfosEntry$json],
};

@$core.Deprecated('Use contestMetadataDescriptor instead')
const ContestMetadata_FieldInfosEntry$json = const {
  '1': 'FieldInfosEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.mcl.ContestFieldInfo', '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `ContestMetadata`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List contestMetadataDescriptor = $convert.base64Decode('Cg9Db250ZXN0TWV0YWRhdGESHgoKQXBpVmVyc2lvbhgBIAEoCVIKQXBpVmVyc2lvbhIYCgdWZXJzaW9uGAIgASgJUgdWZXJzaW9uEiAKC0NvbnRlc3ROYW1lGAMgASgJUgtDb250ZXN0TmFtZRImCg5BdmFpbGFibGVNb2RlcxgEIAMoCVIOQXZhaWxhYmxlTW9kZXMSJgoOQXZhaWxhYmxlQmFuZHMYBSADKAlSDkF2YWlsYWJsZUJhbmRzEiIKDEV4Y2hhbmdlU2VudBgGIAMoCVIMRXhjaGFuZ2VTZW50EiIKDEV4Y2hhbmdlUmN2ZBgHIAMoCVIMRXhjaGFuZ2VSY3ZkEiAKC011bHRpcGxpZXJzGAggAygJUgtNdWx0aXBsaWVycxJECgpGaWVsZEluZm9zGAogAygLMiQubWNsLkNvbnRlc3RNZXRhZGF0YS5GaWVsZEluZm9zRW50cnlSCkZpZWxkSW5mb3MaVAoPRmllbGRJbmZvc0VudHJ5EhAKA2tleRgBIAEoCVIDa2V5EisKBXZhbHVlGAIgASgLMhUubWNsLkNvbnRlc3RGaWVsZEluZm9SBXZhbHVlOgI4AQ==');
