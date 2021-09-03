import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:flutter/painting.dart';
import 'package:fs_auth_example/api.dart';
import 'package:fs_auth_example/src/generated/counter/v1/counter.pbgrpc.dart';

import 'config.dart';

class Counts extends StatefulWidget {
  Counts();
  @override
  _CountsState createState() => _CountsState();
}

class _CountsState extends State<Counts> {
  @override
  Widget build(BuildContext context) {
    final configLoader = ConfigLoaderProvider.of(context).configLoader;
    final user = FirebaseAuth.instance.currentUser;
    if (user == null) {
      Navigator.of(context).pushNamed('/');
    }
    return SafeArea(
      child: Center(
        child: FutureBuilder<IdTokenResult>(
          future: user!.getIdTokenResult(),
          builder: (BuildContext context, AsyncSnapshot<IdTokenResult> snapshot) {
            if (snapshot.hasError) {
              return Text('${snapshot.error}');
            }
            if (snapshot.connectionState != ConnectionState.done) {
              return CircularProgressIndicator();
            }
            return _CountList(CounterService(configLoader).client(snapshot.data!));
          },
        ),
      ),
    );
  }
}

class _CountList extends StatefulWidget {
  final CounterServiceClient counterServiceClient;
  _CountList(this.counterServiceClient);
  @override
  _CountListState createState() => _CountListState();
}

class _CountListState extends State<_CountList> {
  Map<String, Count> counts = {};
  Count myCount = Count();
  @override
  Widget build(BuildContext context) {
    final countStream = widget.counterServiceClient.streamCounts(StreamCountsRequest());
    final user = FirebaseAuth.instance.currentUser!;
    countStream.listen((Count count) {
      setState(() {
        counts[count.name] = count;
        if (user.email == count.name) {
          myCount = count;
        }
      });
    });
    return Stack(
      children: [
        ListView(
          children: counts.values.map((Count count) => ListTile(
            title: Text('${count.name}: ${count.count}'),
          )).toList()
        ),
        Align(
          alignment: Alignment.bottomRight,
          child: FloatingActionButton(
            child: Icon(Icons.add),
            onPressed: () {
              widget.counterServiceClient.updateCount(UpdateCountRequest(
                count: myCount.count + 1,
              ));
            },
          ),
        ),
      ],
    );
  }
}
