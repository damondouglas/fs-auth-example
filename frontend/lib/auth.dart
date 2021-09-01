import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:flutter/painting.dart';
import 'package:flutter/rendering.dart';

import 'counts.dart';

class Login extends StatefulWidget {
  @override
  _LoginState createState() => _LoginState();
}

class _LoginState extends State<Login> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final FirebaseAuth auth = FirebaseAuth.instance;

  @override
  Widget build(BuildContext context) {
    var email = 'user1@example.com';
    var password = 'q1w2e3r4t5';
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
                    if (_formKey.currentState!.validate()) {
                      Navigator.of(context)
                          .push(MaterialPageRoute(builder: (context) {
                        return PasswordAuth(email: email, password: password);
                      }));
                    }
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

class PasswordAuth extends StatefulWidget {
  final String email;
  final String password;

  PasswordAuth({required this.email, required this.password});

  @override
  _PasswordAuthState createState() => _PasswordAuthState();
}

class _PasswordAuthState extends State<PasswordAuth> {
  @override
  Widget build(BuildContext context) {
    return FutureBuilder<UserCredential>(
        future: FirebaseAuth.instance.signInWithEmailAndPassword(
            email: widget.email, password: widget.password),
        builder:
            (BuildContext context, AsyncSnapshot<UserCredential> snapshot) {
          if (snapshot.hasError) {
            ScaffoldMessenger.of(context).showSnackBar(SnackBar(
              content: Text('${snapshot.error}'),
            ));
          }
          final children = <Widget>[];

          if (snapshot.connectionState != ConnectionState.done) {
            children.clear();
            children.add(CircularProgressIndicator());
          }
          if (snapshot.hasData) {
            children.clear();
            children.add(Counts());
          }
          return Row(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: children,
          );
        });
  }
}
