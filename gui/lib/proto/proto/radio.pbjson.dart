///
//  Generated code. Do not modify.
//  source: proto/radio.proto
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
@$core.Deprecated('Use radioVFOConfigDescriptor instead')
const RadioVFOConfig$json = const {
  '1': 'RadioVFOConfig',
  '2': const [
    const {'1': 'mode', '3': 1, '4': 1, '5': 14, '6': '.mcl.RadioMode', '10': 'mode'},
    const {'1': 'frequency', '3': 2, '4': 1, '5': 3, '10': 'frequency'},
  ],
};

/// Descriptor for `RadioVFOConfig`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List radioVFOConfigDescriptor = $convert.base64Decode('Cg5SYWRpb1ZGT0NvbmZpZxIiCgRtb2RlGAEgASgOMg4ubWNsLlJhZGlvTW9kZVIEbW9kZRIcCglmcmVxdWVuY3kYAiABKANSCWZyZXF1ZW5jeQ==');
@$core.Deprecated('Use radioModeConfigDescriptor instead')
const RadioModeConfig$json = const {
  '1': 'RadioModeConfig',
  '2': const [
    const {'1': 'rx', '3': 1, '4': 1, '5': 11, '6': '.mcl.RadioVFOConfig', '10': 'rx'},
    const {'1': 'tx', '3': 2, '4': 1, '5': 11, '6': '.mcl.RadioVFOConfig', '10': 'tx'},
  ],
};

/// Descriptor for `RadioModeConfig`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List radioModeConfigDescriptor = $convert.base64Decode('Cg9SYWRpb01vZGVDb25maWcSIwoCcngYASABKAsyEy5tY2wuUmFkaW9WRk9Db25maWdSAnJ4EiMKAnR4GAIgASgLMhMubWNsLlJhZGlvVkZPQ29uZmlnUgJ0eA==');
@$core.Deprecated('Use radioCommandDescriptor instead')
const RadioCommand$json = const {
  '1': 'RadioCommand',
  '2': const [
    const {'1': 'channel', '3': 1, '4': 1, '5': 5, '10': 'channel'},
    const {'1': 'send_cw_message', '3': 2, '4': 1, '5': 9, '9': 0, '10': 'sendCwMessage'},
    const {'1': 'set_cw_speed', '3': 3, '4': 1, '5': 5, '9': 0, '10': 'setCwSpeed'},
    const {'1': 'set_radio_band_mode', '3': 4, '4': 1, '5': 11, '6': '.mcl.RadioModeConfig', '9': 0, '10': 'setRadioBandMode'},
    const {'1': 'abort_tx', '3': 5, '4': 1, '5': 8, '9': 0, '10': 'abortTx'},
  ],
  '8': const [
    const {'1': 'op'},
  ],
};

/// Descriptor for `RadioCommand`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List radioCommandDescriptor = $convert.base64Decode('CgxSYWRpb0NvbW1hbmQSGAoHY2hhbm5lbBgBIAEoBVIHY2hhbm5lbBIoCg9zZW5kX2N3X21lc3NhZ2UYAiABKAlIAFINc2VuZEN3TWVzc2FnZRIiCgxzZXRfY3dfc3BlZWQYAyABKAVIAFIKc2V0Q3dTcGVlZBJFChNzZXRfcmFkaW9fYmFuZF9tb2RlGAQgASgLMhQubWNsLlJhZGlvTW9kZUNvbmZpZ0gAUhBzZXRSYWRpb0JhbmRNb2RlEhsKCGFib3J0X3R4GAUgASgISABSB2Fib3J0VHhCBAoCb3A=');
@$core.Deprecated('Use radioSelectorDescriptor instead')
const RadioSelector$json = const {
  '1': 'RadioSelector',
  '2': const [
    const {'1': 'channel', '3': 1, '4': 1, '5': 5, '10': 'channel'},
  ],
};

/// Descriptor for `RadioSelector`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List radioSelectorDescriptor = $convert.base64Decode('Cg1SYWRpb1NlbGVjdG9yEhgKB2NoYW5uZWwYASABKAVSB2NoYW5uZWw=');
@$core.Deprecated('Use audioDeviceDescriptor instead')
const AudioDevice$json = const {
  '1': 'AudioDevice',
  '2': const [
    const {'1': 'device_name', '3': 1, '4': 1, '5': 9, '10': 'deviceName'},
    const {'1': 'sample_rate', '3': 2, '4': 1, '5': 9, '10': 'sampleRate'},
  ],
};

/// Descriptor for `AudioDevice`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List audioDeviceDescriptor = $convert.base64Decode('CgtBdWRpb0RldmljZRIfCgtkZXZpY2VfbmFtZRgBIAEoCVIKZGV2aWNlTmFtZRIfCgtzYW1wbGVfcmF0ZRgCIAEoCVIKc2FtcGxlUmF0ZQ==');
@$core.Deprecated('Use radioConfigDescriptor instead')
const RadioConfig$json = const {
  '1': 'RadioConfig',
  '2': const [
    const {'1': 'channel', '3': 1, '4': 1, '5': 5, '10': 'channel'},
    const {'1': 'model', '3': 2, '4': 1, '5': 9, '10': 'model'},
    const {'1': 'uri', '3': 3, '4': 1, '5': 9, '10': 'uri'},
    const {'1': 'audio_input', '3': 4, '4': 1, '5': 11, '6': '.mcl.AudioDevice', '10': 'audioInput'},
    const {'1': 'audio_output', '3': 5, '4': 1, '5': 11, '6': '.mcl.AudioDevice', '10': 'audioOutput'},
  ],
};

/// Descriptor for `RadioConfig`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List radioConfigDescriptor = $convert.base64Decode('CgtSYWRpb0NvbmZpZxIYCgdjaGFubmVsGAEgASgFUgdjaGFubmVsEhQKBW1vZGVsGAIgASgJUgVtb2RlbBIQCgN1cmkYAyABKAlSA3VyaRIxCgthdWRpb19pbnB1dBgEIAEoCzIQLm1jbC5BdWRpb0RldmljZVIKYXVkaW9JbnB1dBIzCgxhdWRpb19vdXRwdXQYBSABKAsyEC5tY2wuQXVkaW9EZXZpY2VSC2F1ZGlvT3V0cHV0');
@$core.Deprecated('Use audioDeviceListDescriptor instead')
const AudioDeviceList$json = const {
  '1': 'AudioDeviceList',
  '2': const [
    const {'1': 'input_devices', '3': 1, '4': 3, '5': 11, '6': '.mcl.AudioDevice', '10': 'inputDevices'},
    const {'1': 'output_devices', '3': 2, '4': 3, '5': 11, '6': '.mcl.AudioDevice', '10': 'outputDevices'},
  ],
};

/// Descriptor for `AudioDeviceList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List audioDeviceListDescriptor = $convert.base64Decode('Cg9BdWRpb0RldmljZUxpc3QSNQoNaW5wdXRfZGV2aWNlcxgBIAMoCzIQLm1jbC5BdWRpb0RldmljZVIMaW5wdXREZXZpY2VzEjcKDm91dHB1dF9kZXZpY2VzGAIgAygLMhAubWNsLkF1ZGlvRGV2aWNlUg1vdXRwdXREZXZpY2Vz');
@$core.Deprecated('Use supportedRadioListDescriptor instead')
const SupportedRadioList$json = const {
  '1': 'SupportedRadioList',
  '2': const [
    const {'1': 'model', '3': 1, '4': 1, '5': 9, '10': 'model'},
    const {'1': 'display_name', '3': 2, '4': 1, '5': 9, '10': 'displayName'},
  ],
};

/// Descriptor for `SupportedRadioList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List supportedRadioListDescriptor = $convert.base64Decode('ChJTdXBwb3J0ZWRSYWRpb0xpc3QSFAoFbW9kZWwYASABKAlSBW1vZGVsEiEKDGRpc3BsYXlfbmFtZRgCIAEoCVILZGlzcGxheU5hbWU=');
