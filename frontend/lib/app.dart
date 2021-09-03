import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
import 'package:fs_auth_example/config.dart';

import 'auth.dart';
import 'counts.dart';
import 'constants.dart' as constants;

final Future<FirebaseApp> _initialization = Firebase.initializeApp();

class App extends StatefulWidget {
  @override
  _AppState createState() => _AppState();
}

class _AppState extends State<App> {
  @override
  Widget build(BuildContext context) {
    return FutureBuilder<ConfigLoader>(
        future: _initialize(),
        builder: (BuildContext context, AsyncSnapshot<ConfigLoader> snapshot) {
          if (snapshot.hasError) {
            print(snapshot.error);
          }
          if (snapshot.connectionState != ConnectionState.done) {
            return Container();
          }
          if (snapshot.data == null) {
            return Container();
          }
          return MaterialApp(
            debugShowCheckedModeBanner: false,
            theme: ThemeData.dark(),
            routes: {
              constants.kRouteCounts: (_) => Scaffold(
                    body: ConfigLoaderProvider(
                      configLoader: snapshot.data!,
                      child: Counts(),
                    ),
                  ),
              constants.kRouteHome: (_) => Scaffold(
                    body: Login(),
                  ),
            },
          );
        });
  }
}

Future<ConfigLoader> _initialize() async {
  await _initialization;
  final configMap = await loadConfigMap();
  return Future.value(ConfigLoader(configMap));
}
