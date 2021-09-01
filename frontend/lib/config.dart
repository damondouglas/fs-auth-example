import 'dart:convert';

import 'package:flutter/services.dart';

const _configAssetKey = 'assets/config.json';
const _address = 'address';
const _http = 'http';
const _port = 'port';
const _tcp = 'tcp';

Future<String> _loadConfig() {
  return rootBundle.loadString(_configAssetKey);
}

Future<Config> loadConfig() async {
  var data = await _loadConfig();
  Map<String, dynamic> configMap = jsonDecode(data);
  return Config(
    http: HttpConfig.fromMap(configMap[_http]),
    tcp: TcpConfig.fromMap(configMap[_tcp])
  );
}

class Config {
  final HttpConfig http;
  final TcpConfig tcp;
  Config({required this.http, required this.tcp});
}

class HttpConfig {
  final String address;
  final int port;
  HttpConfig.fromMap(Map<String, dynamic> map) :
      address = map[_address],
      port = map[_port];
}

class TcpConfig {
  final String address;
  TcpConfig.fromMap(Map<String, dynamic> map) :
      address = map[_address];
}
