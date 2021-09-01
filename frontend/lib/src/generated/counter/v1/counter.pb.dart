///
//  Generated code. Do not modify.
//  source: counter/v1/counter.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class Count extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Count', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'counter.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'count')
    ..hasRequiredFields = false
  ;

  Count._() : super();
  factory Count({
    $core.String? name,
    $fixnum.Int64? count,
  }) {
    final _result = create();
    if (name != null) {
      _result.name = name;
    }
    if (count != null) {
      _result.count = count;
    }
    return _result;
  }
  factory Count.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Count.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Count clone() => Count()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Count copyWith(void Function(Count) updates) => super.copyWith((message) => updates(message as Count)) as Count; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Count create() => Count._();
  Count createEmptyInstance() => create();
  static $pb.PbList<Count> createRepeated() => $pb.PbList<Count>();
  @$core.pragma('dart2js:noInline')
  static Count getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Count>(create);
  static Count? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get count => $_getI64(1);
  @$pb.TagNumber(2)
  set count($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCount() => $_has(1);
  @$pb.TagNumber(2)
  void clearCount() => clearField(2);
}

class ListCountsRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ListCountsRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'counter.v1'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'limit')
    ..hasRequiredFields = false
  ;

  ListCountsRequest._() : super();
  factory ListCountsRequest({
    $fixnum.Int64? limit,
  }) {
    final _result = create();
    if (limit != null) {
      _result.limit = limit;
    }
    return _result;
  }
  factory ListCountsRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListCountsRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListCountsRequest clone() => ListCountsRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListCountsRequest copyWith(void Function(ListCountsRequest) updates) => super.copyWith((message) => updates(message as ListCountsRequest)) as ListCountsRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListCountsRequest create() => ListCountsRequest._();
  ListCountsRequest createEmptyInstance() => create();
  static $pb.PbList<ListCountsRequest> createRepeated() => $pb.PbList<ListCountsRequest>();
  @$core.pragma('dart2js:noInline')
  static ListCountsRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListCountsRequest>(create);
  static ListCountsRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get limit => $_getI64(0);
  @$pb.TagNumber(1)
  set limit($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLimit() => $_has(0);
  @$pb.TagNumber(1)
  void clearLimit() => clearField(1);
}

class ListCountsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ListCountsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'counter.v1'), createEmptyInstance: create)
    ..pc<Count>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'items', $pb.PbFieldType.PM, subBuilder: Count.create)
    ..hasRequiredFields = false
  ;

  ListCountsResponse._() : super();
  factory ListCountsResponse({
    $core.Iterable<Count>? items,
  }) {
    final _result = create();
    if (items != null) {
      _result.items.addAll(items);
    }
    return _result;
  }
  factory ListCountsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListCountsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListCountsResponse clone() => ListCountsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListCountsResponse copyWith(void Function(ListCountsResponse) updates) => super.copyWith((message) => updates(message as ListCountsResponse)) as ListCountsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListCountsResponse create() => ListCountsResponse._();
  ListCountsResponse createEmptyInstance() => create();
  static $pb.PbList<ListCountsResponse> createRepeated() => $pb.PbList<ListCountsResponse>();
  @$core.pragma('dart2js:noInline')
  static ListCountsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListCountsResponse>(create);
  static ListCountsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Count> get items => $_getList(0);
}

class StreamCountsRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StreamCountsRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'counter.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  StreamCountsRequest._() : super();
  factory StreamCountsRequest() => create();
  factory StreamCountsRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StreamCountsRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StreamCountsRequest clone() => StreamCountsRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StreamCountsRequest copyWith(void Function(StreamCountsRequest) updates) => super.copyWith((message) => updates(message as StreamCountsRequest)) as StreamCountsRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StreamCountsRequest create() => StreamCountsRequest._();
  StreamCountsRequest createEmptyInstance() => create();
  static $pb.PbList<StreamCountsRequest> createRepeated() => $pb.PbList<StreamCountsRequest>();
  @$core.pragma('dart2js:noInline')
  static StreamCountsRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StreamCountsRequest>(create);
  static StreamCountsRequest? _defaultInstance;
}

class UpdateCountRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateCountRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'counter.v1'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'count')
    ..hasRequiredFields = false
  ;

  UpdateCountRequest._() : super();
  factory UpdateCountRequest({
    $fixnum.Int64? count,
  }) {
    final _result = create();
    if (count != null) {
      _result.count = count;
    }
    return _result;
  }
  factory UpdateCountRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateCountRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateCountRequest clone() => UpdateCountRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateCountRequest copyWith(void Function(UpdateCountRequest) updates) => super.copyWith((message) => updates(message as UpdateCountRequest)) as UpdateCountRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateCountRequest create() => UpdateCountRequest._();
  UpdateCountRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateCountRequest> createRepeated() => $pb.PbList<UpdateCountRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdateCountRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateCountRequest>(create);
  static UpdateCountRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get count => $_getI64(0);
  @$pb.TagNumber(1)
  set count($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCount() => $_has(0);
  @$pb.TagNumber(1)
  void clearCount() => clearField(1);
}

class UpdateCountResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateCountResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'counter.v1'), createEmptyInstance: create)
    ..aInt64(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'count')
    ..hasRequiredFields = false
  ;

  UpdateCountResponse._() : super();
  factory UpdateCountResponse({
    $fixnum.Int64? count,
  }) {
    final _result = create();
    if (count != null) {
      _result.count = count;
    }
    return _result;
  }
  factory UpdateCountResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateCountResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateCountResponse clone() => UpdateCountResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateCountResponse copyWith(void Function(UpdateCountResponse) updates) => super.copyWith((message) => updates(message as UpdateCountResponse)) as UpdateCountResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateCountResponse create() => UpdateCountResponse._();
  UpdateCountResponse createEmptyInstance() => create();
  static $pb.PbList<UpdateCountResponse> createRepeated() => $pb.PbList<UpdateCountResponse>();
  @$core.pragma('dart2js:noInline')
  static UpdateCountResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateCountResponse>(create);
  static UpdateCountResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get count => $_getI64(0);
  @$pb.TagNumber(1)
  set count($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCount() => $_has(0);
  @$pb.TagNumber(1)
  void clearCount() => clearField(1);
}

class CounterServiceApi {
  $pb.RpcClient _client;
  CounterServiceApi(this._client);

  $async.Future<ListCountsResponse> listCounts($pb.ClientContext? ctx, ListCountsRequest request) {
    var emptyResponse = ListCountsResponse();
    return _client.invoke<ListCountsResponse>(ctx, 'CounterService', 'ListCounts', request, emptyResponse);
  }
  $async.Future<UpdateCountResponse> updateCount($pb.ClientContext? ctx, UpdateCountRequest request) {
    var emptyResponse = UpdateCountResponse();
    return _client.invoke<UpdateCountResponse>(ctx, 'CounterService', 'UpdateCount', request, emptyResponse);
  }
  $async.Future<Count> streamCounts($pb.ClientContext? ctx, StreamCountsRequest request) {
    var emptyResponse = Count();
    return _client.invoke<Count>(ctx, 'CounterService', 'StreamCounts', request, emptyResponse);
  }
}

