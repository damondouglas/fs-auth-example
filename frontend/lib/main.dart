import 'package:flutter/foundation.dart' as foundation;
import 'package:flutter/material.dart';

import 'app.dart';

void main() {
  print(foundation.kIsWeb);
  WidgetsFlutterBinding.ensureInitialized();
  runApp(App());
}
