import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:flutter/painting.dart';
import 'package:flutter/rendering.dart';

import 'constants.dart' as constants;

/**
 * Login performs authentication flow using a form to acquire email/password
 * credentials from a user.
 */
class Login extends StatefulWidget {
  @override
  _LoginState createState() => _LoginState();
}

class _LoginState extends State<Login> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final FirebaseAuth auth = FirebaseAuth.instance;

  @override
  Widget build(BuildContext context) {
    var email = '';
    var password = '';
    return Center(
      child: Container(
        height: 400,
        width: 200,
        child: Form(
          key: _formKey,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              TextFormField(
                initialValue: email,
                decoration: InputDecoration(
                  hintText: 'Enter your email',
                ),
                validator: (String? value) {
                  return (value == null || value.isEmpty)
                      ? 'Email is required'
                      : null;
                },
                onChanged: (String? value) {
                  if (value == null || value.isEmpty) {
                    return;
                  }
                  email = value;
                },
              ),
              TextFormField(
                initialValue: password,
                obscureText: true,
                decoration: InputDecoration(
                  hintText: 'Enter your password',
                ),
                validator: (String? value) {
                  return (value == null || value.isEmpty)
                      ? 'Password is required'
                      : null;
                },
                onChanged: (String? value) {
                  if (value == null || value.isEmpty) {
                    return;
                  }
                  password = value;
                },
              ),
              Padding(
                padding: EdgeInsets.symmetric(vertical: 16.0),
                child: ElevatedButton(
                  onPressed: () {
                    if (!_formKey.currentState!.validate()) {
                      return;
                    }
                    _login(email, password)
                    .then((_) {
                      Navigator.of(context).pushNamed(constants.kRouteCounts);
                    }).catchError((err) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(content: Text('$err'))
                      );
                    });
                  },
                  child: Text('Submit'),
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}

Future<UserCredential> _login(String email, String password) =>
    FirebaseAuth.instance.signInWithEmailAndPassword(email: email, password: password);

