/* -------------- External -------------- */
import 'dart:convert';
import 'dart:io';
import 'package:http/http.dart' as http;

/* -------------- Services -------------- */
import 'package:flutter/services.dart';

/* -------------- Types -------------- */
import 'package:trips/types/trip.dart';
import 'package:trips/types/user.dart' as u;

/* -------------- Utils -------------- */
import 'package:trips/utils/constants.dart';

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

  Future< /*List<Trip>*/ void> findByUserId(String googleId) async {
    final String url = "$apiUrl/trips?user_pk=user#$googleId";
    http.Response res = await http.get(Uri.parse(url));

    dynamic trips = jsonDecode(res.body);

    print(trips);
  }

  static Future<List<String?>> uploadImages(
      u.User user, List<File> imgs, String tripId) async {
    List<String?> urls = await Future.wait(imgs.map((img) async {
      // var upload = await supabase.storage.from('pictures').upload(
      //     "${user.id}/$tripId/${const Uuid().v4()}.${img.path.split('.').last}",
      //     img);
      // return upload.data;
      return "";
    }));

    return urls;
  }

  static Future createTrip(u.User user, String title, String description,
      DateTime doneAt, List<File> imgs) async {
    // PostgrestResponse<dynamic> newTrip = await supabase.from('trips').insert({
    //   'user': "${user.id}",
    //   'done_at': doneAt.toString(),
    //   'name': title,
    //   'description': description,
    // }).execute();
    // if (newTrip.error != null) {
    //   print(newTrip.error);
    //   throw ("Error!");
    // }
    // final uploaded_imgs = await uploadImages(user, imgs, newTrip.data[0]['id']);
    // final res = await supabase.storage
    //     .from('pictures')
    //     .list(path: 'd698867d-822e-4f1c-a073-8c73152167b0');
    // final data = res.data;
    // print(data?.map((e) => e.name).toList());
    // // print(newTrip.data);
    // // print(uploaded_imgs);
  }

  Future<Trip> findById(String id) async {
    var trips = await find();
    await Future.delayed(const Duration(milliseconds: 500));

    return trips.where((trip) => trip.id == id).first;
  }
}
