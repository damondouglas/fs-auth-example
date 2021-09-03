///
//  Generated code. Do not modify.
//  source: counter/v1/counter.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'counter.pb.dart' as $0;
export 'counter.pb.dart';

class CounterServiceClient extends $grpc.Client {
  static final _$listCounts =
      $grpc.ClientMethod<$0.ListCountsRequest, $0.ListCountsResponse>(
          '/counter.v1.CounterService/ListCounts',
          ($0.ListCountsRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.ListCountsResponse.fromBuffer(value));
  static final _$updateCount =
      $grpc.ClientMethod<$0.UpdateCountRequest, $0.UpdateCountResponse>(
          '/counter.v1.CounterService/UpdateCount',
          ($0.UpdateCountRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.UpdateCountResponse.fromBuffer(value));
  static final _$streamCounts =
      $grpc.ClientMethod<$0.StreamCountsRequest, $0.Count>(
          '/counter.v1.CounterService/StreamCounts',
          ($0.StreamCountsRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Count.fromBuffer(value));

  CounterServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.ListCountsResponse> listCounts(
      $0.ListCountsRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listCounts, request, options: options);
  }

  $grpc.ResponseFuture<$0.UpdateCountResponse> updateCount(
      $0.UpdateCountRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateCount, request, options: options);
  }

  $grpc.ResponseStream<$0.Count> streamCounts($0.StreamCountsRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$streamCounts, $async.Stream.fromIterable([request]),
        options: options);
  }
}

abstract class CounterServiceBase extends $grpc.Service {
  $core.String get $name => 'counter.v1.CounterService';

  CounterServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.ListCountsRequest, $0.ListCountsResponse>(
        'ListCounts',
        listCounts_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ListCountsRequest.fromBuffer(value),
        ($0.ListCountsResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.UpdateCountRequest, $0.UpdateCountResponse>(
            'UpdateCount',
            updateCount_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.UpdateCountRequest.fromBuffer(value),
            ($0.UpdateCountResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.StreamCountsRequest, $0.Count>(
        'StreamCounts',
        streamCounts_Pre,
        false,
        true,
        ($core.List<$core.int> value) =>
            $0.StreamCountsRequest.fromBuffer(value),
        ($0.Count value) => value.writeToBuffer()));
  }

  $async.Future<$0.ListCountsResponse> listCounts_Pre($grpc.ServiceCall call,
      $async.Future<$0.ListCountsRequest> request) async {
    return listCounts(call, await request);
  }

  $async.Future<$0.UpdateCountResponse> updateCount_Pre($grpc.ServiceCall call,
      $async.Future<$0.UpdateCountRequest> request) async {
    return updateCount(call, await request);
  }

  $async.Stream<$0.Count> streamCounts_Pre($grpc.ServiceCall call,
      $async.Future<$0.StreamCountsRequest> request) async* {
    yield* streamCounts(call, await request);
  }

  $async.Future<$0.ListCountsResponse> listCounts(
      $grpc.ServiceCall call, $0.ListCountsRequest request);
  $async.Future<$0.UpdateCountResponse> updateCount(
      $grpc.ServiceCall call, $0.UpdateCountRequest request);
  $async.Stream<$0.Count> streamCounts(
      $grpc.ServiceCall call, $0.StreamCountsRequest request);
}
