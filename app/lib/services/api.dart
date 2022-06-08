/* -------------- External -------------- */
import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

/* -------------- Constants -------------- */
import 'package:trips/utils/constants.dart';

class Api {
  Api({this.secure}) : super();

  final bool? secure;

  Future<http.Response> post(String endpoint, Object body) async {
    Map<String, String> headers = {
      'Content-Type': 'application/json; charset=UTF-8'
    };

    if (secure == true) {
      SharedPreferences prefs = await SharedPreferences.getInstance();
      String? userId = prefs.getString("googleId");

      if (userId == null) throw ("Not logged in");

      headers.addAll({
        "Authorization": "user#$userId",
      });
    }

    return http.post(Uri.parse("$apiUrl$endpoint"),
        body: jsonEncode(body), headers: headers);
  }

  Future<http.Response> get(String endpoint) async {
    Map<String, String> headers = {
      'Content-Type': 'application/json; charset=UTF-8'
    };

    if (secure == true) {
      SharedPreferences prefs = await SharedPreferences.getInstance();
      String? userId = prefs.getString("googleId");

      if (userId == null) throw ("Not logged in");

      headers.addAll({
        "Authorization": "user#$userId",
      });
    }
    print(Uri.parse("$apiUrl$endpoint"));
    return http.get(Uri.parse("$apiUrl$endpoint"), headers: headers);
  }
}
