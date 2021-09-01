///
//  Generated code. Do not modify.
//  source: counter/v1/counter.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'counter.pb.dart' as $0;
import 'counter.pbjson.dart';

export 'counter.pb.dart';

abstract class CounterServiceBase extends $pb.GeneratedService {
  $async.Future<$0.ListCountsResponse> listCounts($pb.ServerContext ctx, $0.ListCountsRequest request);
  $async.Future<$0.UpdateCountResponse> updateCount($pb.ServerContext ctx, $0.UpdateCountRequest request);
  $async.Future<$0.Count> streamCounts($pb.ServerContext ctx, $0.StreamCountsRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'ListCounts': return $0.ListCountsRequest();
      case 'UpdateCount': return $0.UpdateCountRequest();
      case 'StreamCounts': return $0.StreamCountsRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'ListCounts': return this.listCounts(ctx, request as $0.ListCountsRequest);
      case 'UpdateCount': return this.updateCount(ctx, request as $0.UpdateCountRequest);
      case 'StreamCounts': return this.streamCounts(ctx, request as $0.StreamCountsRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => CounterServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => CounterServiceBase$messageJson;
}

