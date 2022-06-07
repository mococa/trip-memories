/* -------------- External -------------- */
import 'package:flutter/material.dart';
import 'package:flutter_signin_button/flutter_signin_button.dart';
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

  redirect(User user) {
    context.read<Authentication>().setUser(user);
    Navigator.of(context).pushNamedAndRemoveUntil('/', (route) => false);
  }

  Future<void> _signIn() async {
    setState(() {
      _isLoading = true;
    });

    try {
      final googleUser = await OAuth.getGoogleAccount();
      if (googleUser == null) throw ('user is null');
      User user = await OAuth.authenticate(googleUser);
      redirect(user);
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
    return Container(
        decoration: BoxDecoration(
          image: DecorationImage(
            colorFilter: ColorFilter.mode(
                Colors.black.withOpacity(0.35), BlendMode.dstATop),
            image: const AssetImage(
                "assets/images/background/authentication-bg.jpg"),
            fit: BoxFit.cover,
          ),
        ),
        child: Scaffold(
          backgroundColor: Colors.transparent,
          body: ListView(
              reverse: true,
              padding: const EdgeInsets.symmetric(vertical: 18, horizontal: 12),
              children: [
                const SizedBox(height: 18),
                //
                buildGoogleSignInButton(),
                //
                const SizedBox(height: 24),
                //
                const Text.rich(
                  TextSpan(
                    text: "Memories\nof your trips\n",
                    children: [
                      TextSpan(
                        text: "Everlasting",
                        style: TextStyle(
                            fontWeight: FontWeight.bold, fontSize: 54),
                      ),
                    ],
                    style: TextStyle(fontSize: 48, color: Colors.white),
                  ),
                )
              ]),
        ));
  }

  SignInButtonBuilder buildGoogleSignInButton() {
    return SignInButtonBuilder(
      backgroundColor: Colors.white,
      onPressed: _isLoading ? () {} : _signIn,
      text: "Sign In with Google",
      fontSize: 16,
      textColor: Colors.black54,
      image: ClipRRect(
        borderRadius: BorderRadius.circular(8.0),
        child: const Image(
          image: AssetImage(
            'assets/logos/google_light.png',
            package: 'flutter_signin_button',
          ),
          height: 44.0,
        ),
      ),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(24),
      ),
    );
  }
}
