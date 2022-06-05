/* -------------- External -------------- */
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

/* -------------- Providers -------------- */
import 'package:trips/providers/authentication.dart';

/* -------------- Types -------------- */
import 'package:trips/types/user.dart';

/* -------------- Services -------------- */
import 'package:trips/services/oauth.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({Key? key}) : super(key: key);

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  bool _isLoading = false;

  Future<void> _signIn() async {
    setState(() {
      _isLoading = true;
    });

    try {
      final googleUser = await OAuth.googleSignIn();
      if (googleUser == null) throw ('user is null');
      User user = await OAuth.authenticate(googleUser);

      context.read<Authentication>().setUser(user);
      Navigator.of(context).pushNamedAndRemoveUntil('/', (route) => false);
    } catch (err) {
      print(err);
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Trip Memos')),
      body: ListView(
        padding: const EdgeInsets.symmetric(vertical: 18, horizontal: 12),
        children: [
          const SizedBox(height: 18),
          ElevatedButton(
            onPressed: _isLoading ? null : _signIn,
            child: Text(_isLoading ? 'Loading' : 'Google Sign in'),
          ),
        ],
      ),
    );
  }
}
