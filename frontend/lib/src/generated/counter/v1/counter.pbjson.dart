///
//  Generated code. Do not modify.
//  source: counter/v1/counter.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use countDescriptor instead')
const Count$json = const {
  '1': 'Count',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'count', '3': 2, '4': 1, '5': 3, '10': 'count'},
  ],
};

/// Descriptor for `Count`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List countDescriptor = $convert.base64Decode('CgVDb3VudBISCgRuYW1lGAEgASgJUgRuYW1lEhQKBWNvdW50GAIgASgDUgVjb3VudA==');
@$core.Deprecated('Use listCountsRequestDescriptor instead')
const ListCountsRequest$json = const {
  '1': 'ListCountsRequest',
  '2': const [
    const {'1': 'limit', '3': 1, '4': 1, '5': 3, '10': 'limit'},
  ],
};

/// Descriptor for `ListCountsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listCountsRequestDescriptor = $convert.base64Decode('ChFMaXN0Q291bnRzUmVxdWVzdBIUCgVsaW1pdBgBIAEoA1IFbGltaXQ=');
@$core.Deprecated('Use listCountsResponseDescriptor instead')
const ListCountsResponse$json = const {
  '1': 'ListCountsResponse',
  '2': const [
    const {'1': 'items', '3': 1, '4': 3, '5': 11, '6': '.counter.v1.Count', '10': 'items'},
  ],
};

/// Descriptor for `ListCountsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listCountsResponseDescriptor = $convert.base64Decode('ChJMaXN0Q291bnRzUmVzcG9uc2USJwoFaXRlbXMYASADKAsyES5jb3VudGVyLnYxLkNvdW50UgVpdGVtcw==');
@$core.Deprecated('Use streamCountsRequestDescriptor instead')
const StreamCountsRequest$json = const {
  '1': 'StreamCountsRequest',
};

/// Descriptor for `StreamCountsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List streamCountsRequestDescriptor = $convert.base64Decode('ChNTdHJlYW1Db3VudHNSZXF1ZXN0');
@$core.Deprecated('Use updateCountRequestDescriptor instead')
const UpdateCountRequest$json = const {
  '1': 'UpdateCountRequest',
  '2': const [
    const {'1': 'count', '3': 1, '4': 1, '5': 3, '10': 'count'},
  ],
};

/// Descriptor for `UpdateCountRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateCountRequestDescriptor = $convert.base64Decode('ChJVcGRhdGVDb3VudFJlcXVlc3QSFAoFY291bnQYASABKANSBWNvdW50');
@$core.Deprecated('Use updateCountResponseDescriptor instead')
const UpdateCountResponse$json = const {
  '1': 'UpdateCountResponse',
  '2': const [
    const {'1': 'count', '3': 1, '4': 1, '5': 3, '10': 'count'},
  ],
};

/// Descriptor for `UpdateCountResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateCountResponseDescriptor = $convert.base64Decode('ChNVcGRhdGVDb3VudFJlc3BvbnNlEhQKBWNvdW50GAEgASgDUgVjb3VudA==');
