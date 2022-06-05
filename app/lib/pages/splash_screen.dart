/* -------------- External -------------- */
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

/* -------------- Types -------------- */
import 'package:trips/types/user.dart';

/* -------------- Services -------------- */
import 'package:trips/services/oauth.dart';

/* -------------- Providers -------------- */
import 'package:trips/providers/authentication.dart';

class SplashPage extends StatefulWidget {
  const SplashPage({Key? key}) : super(key: key);

  @override
  _SplashPageState createState() => _SplashPageState();
}

class _SplashPageState extends State<SplashPage> {
  @override
  void initState() {
    super.initState();

    Future.delayed(const Duration(milliseconds: 100)).then((_) async {
      try {
        User user = await OAuth.recoverSession();
        context.read<Authentication>().setUser(user);
        // onAuthenticated(_);
        Navigator.of(context).pushNamedAndRemoveUntil('/', (route) => false);
      } catch (err) {
        context.read<Authentication>().setUser(null);
        Navigator.of(context)
            .pushNamedAndRemoveUntil('/login', (route) => false);
        // onUnauthenticated(_);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(child: CircularProgressIndicator()),
    );
  }
}
