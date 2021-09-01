import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';

import 'auth.dart';

class Home extends StatelessWidget {
  final FirebaseApp firebaseApp;
  Home(this.firebaseApp);

  @override
  Widget build(BuildContext context) {
    return Login();
  }
}
