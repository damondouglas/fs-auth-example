import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';

class Counts extends StatefulWidget {
  @override
  _CountsState createState() => _CountsState();
}

class _CountsState extends State<Counts> {
  final FirebaseAuth auth = FirebaseAuth.instance;

  @override
  Widget build(BuildContext context) {
    if (auth.currentUser == null) {
      Navigator.of(context).pop();
    }
    return Container();
  }
}

// class Api {
//   final ChannelMode mode;
//   late ClientChannel _channel;
//
//   Api({required this.mode});
//
//   Future<CounterServiceApi> _api() async {
//     await init();
//     return CounterServiceApi(_channel as RpcClient);
//   }
//
//   Future init() async {
//     var _config = await loadConfig();
//     switch(mode) {
//       case ChannelMode.tcp:
//         _channel = _tcp(_config);
//         break;
//       case ChannelMode.http:
//         _channel = _http(_config);
//         break;
//     }
//     return Future.value();
//   }
//   ClientChannel _tcp(Config config) => ClientChannel(config.address,
//     port: config.port,
//     options: ChannelOptions(
//       credentials: ChannelCredentials.secure(),
//     ),
//   );
//
//   ClientChannel _http(Config config) => GrpcWebClientChannel.xhr(
//       Uri.parse('${config.scheme}://${config.address}:${config.port}')
//   ) as ClientChannel;
// }
//
// enum ChannelMode {
//   http,
//   tcp,
// }
