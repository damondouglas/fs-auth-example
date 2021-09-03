import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

const _configAssetKey = 'assets/config.json';
const _host = 'host';
const _http = 'http';
const _port = 'port';
const _tcp = 'tcp';

/**
 * ConfigLoaderProvider injects inheritable configuration data into the application.
 */
class ConfigLoaderProvider extends InheritedWidget {
  final Widget child;
  final ConfigLoader configLoader;
  ConfigLoaderProvider({required this.child, required this.configLoader}) : super(child: child);

  static ConfigLoaderProvider of(BuildContext context) {
    final ConfigLoaderProvider? result = context.dependOnInheritedWidgetOfExactType<ConfigLoaderProvider>();
    assert(result != null, 'No ConfigLoaderProvider found in context');
    return result!;
  }

  @override
  bool updateShouldNotify(ConfigLoaderProvider old) => old.configLoader != this.configLoader;
}

Future<String> _loadConfig() {
  return rootBundle.loadString(_configAssetKey);
}

Future<Map<String, dynamic>> loadConfigMap() async {
  var data = await _loadConfig();
  return jsonDecode(data);
}

/**
 * ConfigLoader loads application configuration data.
 */
class ConfigLoader {
  final Map<String, dynamic> _data;
  ConfigLoader(this._data);

  Map<String, dynamic> get _httpData => _data[_http];

  Map<String, dynamic> get _tcpData => _data[_tcp];

  Config get http => Config(_httpData);

  Config get tcp => Config(_tcpData);
}

/**
 * Config stores application configuration data.
 */
class Config {
  final Map<String, dynamic> _data;
  Config(this._data);

  String get host => _data[_host];
  int get port => _data[_port];
}
