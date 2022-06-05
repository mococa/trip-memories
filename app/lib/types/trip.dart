class Trip {
  String id;
  String title;
  String description;
  String banner;
  List<String> images;

  Trip(
      {required this.id,
      required this.title,
      required this.description,
      required this.banner,
      required this.images});
}

class TripArguments {
  String id;
  String banner;

  TripArguments({required this.id, required this.banner});
}
