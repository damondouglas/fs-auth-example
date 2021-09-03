import 'dart:convert';

import 'package:firebase_auth/firebase_auth.dart';
import 'package:fs_auth_example/config.dart';
import 'package:fs_auth_example/src/generated/counter/v1/counter.pbgrpc.dart';
import 'package:grpc/grpc.dart';
import 'package:grpc/grpc_or_grpcweb.dart';

final _authorization = 'authorization';
final _tokenType = 'Bearer';

class CounterService {
  final ConfigLoader _configLoader;
  CounterService(this._configLoader);

  Config get _http => _configLoader.http;
  Config get _tcp => _configLoader.tcp;

  GrpcOrGrpcWebClientChannel get _channel => GrpcOrGrpcWebClientChannel.toSeparateEndpoints(
    grpcHost: _tcp.host,
    grpcPort: _tcp.port,
    grpcTransportSecure: true,
    grpcWebHost: _http.host,
    grpcWebPort: _http.port,
    grpcWebTransportSecure: true,
  );

  CounterServiceClient client(IdTokenResult idTokenResult) {
    return CounterServiceClient(
      _channel,
      options: CallOptions(metadata: {
        _authorization: '$_tokenType ${idTokenResult.token}',
      }),
    );
  }
}
