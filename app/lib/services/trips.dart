/* -------------- External -------------- */
import 'dart:convert';
import 'package:http_parser/http_parser.dart';
import 'dart:io';
import 'package:http/http.dart' as http;

/* -------------- Services -------------- */
import 'package:flutter/services.dart';
import 'package:mime/mime.dart';

/* -------------- Types -------------- */
import 'package:trips/types/trip.dart';
import 'package:trips/types/user.dart' as u;

/* -------------- Utils -------------- */
import 'package:trips/utils/constants.dart';

/* -------------- Services -------------- */
import 'package:trips/services/api.dart';

class Trips {
  Future<List<Trip>> find() async {
    final String response =
        await rootBundle.loadString('assets/mocks/trips.json');
    Iterable data = json.decode(response);
    Iterable<Trip> trips = data.map((json) => Trip(
        id: json['id'],
        banner: json['banner'],
        description: json['description'],
        images: json['images'].cast<String>(),
        title: json['title']));

    return trips.toList();
  }

  static Future<void> findUserTrips(String googleId) async {
    final res = await Api(secure: true).get("/trips");

    dynamic trips = jsonDecode(res.body);

    print(trips);
  }

  static Future<String?> uploadImages(
      u.User user, List<File> imgs, String tripId) async {
    final request =
        http.MultipartRequest("POST", Uri.parse("$apiUrl/trips/upload-images"));

    final files = await Future.wait(imgs.map((image) async {
      final mimeTypeData =
          lookupMimeType(image.path, headerBytes: [0xFF, 0xD8])!.split('/');

      final file = await http.MultipartFile.fromPath('image', image.path,
          contentType: MediaType(mimeTypeData[0], mimeTypeData[1]));

      return file;
    }));

    request.headers.addAll({"Authorization": "user#${user.googleId}"});

    request.files.addAll(files);

    request.fields['trip'] = tripId;

    final streamedResponse = await request.send();

    final response = await http.Response.fromStream(streamedResponse);

    print(response.body);
    return response.body;
  }

  static Future createTrip(u.User user, String title, String description,
      DateTime doneAt, List<File> imgs) async {
    final res = await Api(secure: true).post(
        "/trips",
        jsonEncode({
          'partition_key': "user#${user.googleId}",
          'name': title,
          'description': description,
          'done_at': doneAt.millisecond,
        }));

    print(res.body);
  }

  Future<Trip> findById(String id) async {
    var trips = await find();
    await Future.delayed(const Duration(milliseconds: 500));

    return trips.where((trip) => trip.id == id).first;
  }
}
