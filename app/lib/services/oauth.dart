/* -------------- External -------------- */
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:provider/provider.dart';
import 'package:shared_preferences/shared_preferences.dart';

/* -------------- Providers -------------- */
import 'package:trips/providers/authentication.dart';

/* -------------- Types -------------- */
import 'package:trips/types/user.dart';

/* -------------- Services -------------- */
import 'package:trips/services/api.dart';

class OAuth {
  static final _googleSignIn = GoogleSignIn();

  static Future<GoogleSignInAccount?> getGoogleAccount() async =>
      await _googleSignIn.signIn().catchError((error) {
        throw Error();
      });

  static Future<User> login(String googleId) async {
    final res = await Api().get("/auth?user_pk=user%23$googleId");
    dynamic user = jsonDecode(res.body);

    return User.fromJson(user);
  }

  static Future<User> signUp(User user) async {
    final res = await Api().post("/auth", jsonEncode(user.toJson()));

    dynamic createdUser = jsonDecode(res.body);

    return User.fromJson(createdUser);
  }

  static Future<User> authenticate(GoogleSignInAccount gAccount) async {
    try {
      User user = await login(gAccount.id);
      await createSession(user);
      return user;
    } catch (err) {
      User user = User(
        email: gAccount.email,
        googleId: gAccount.id,
        name: gAccount.displayName ?? '',
        picture: gAccount.photoUrl ?? '',
      );

      try {
        User createdUser = await signUp(user);
        await createSession(createdUser);
        return createdUser;
      } catch (err) {
        rethrow;
      }
    }
  }

  static Future<User> recoverSession() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String? googleId = prefs.getString('googleId');
    if (googleId == null) throw ('id not stored');
    var user = await login(googleId);
    return user;
  }

  static Future<void> createSession(User user) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString('googleId', user.googleId);
  }

  static Future<void> clearSession(BuildContext context) async {
    Navigator.of(context).pushNamedAndRemoveUntil('/login', (route) => false);
    context.read<Authentication>().setUser(null);

    await _googleSignIn.disconnect();

    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.clear();
  }
}
