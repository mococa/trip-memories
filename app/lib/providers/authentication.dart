/* -------------- External -------------- */
import 'package:flutter/material.dart';

/* -------------- Types -------------- */
import 'package:trips/types/user.dart';

class Authentication with ChangeNotifier {
  User? _user;

  User? get user => _user;

  setUser(User? user) {
    _user = user;

    notifyListeners();
  }
}
