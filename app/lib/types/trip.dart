class Trip {
  String partitionKey;
  String sortKey;
  String picture;
  String name;
  String description;
  int doneAt;
  int createdAt;
  List<String>? images;

  Trip({
    required this.partitionKey,
    required this.sortKey,
    required this.picture,
    required this.name,
    required this.description,
    required this.doneAt,
    required this.createdAt,
    this.images,
  });

  Map<String, dynamic> toJson() {
    Map<String, dynamic> map = {
      'partition_key': partitionKey,
      'sort_key': sortKey,
      'picture': picture,
      'name': name,
      'description': description,
      'done_at': doneAt,
      'created_at': createdAt,
      'images': images
    };

    return map;
  }

  static Trip fromJson(dynamic json) {
    Trip trip = Trip(
        partitionKey: json['partition_key'],
        sortKey: json['sort_key'],
        name: json['name'],
        description: json['description'],
        picture: json['picture'],
        createdAt: json['created_at'],
        doneAt: json['done_at'],
        images: []);

    final imgs = json['images'];
    if (imgs == null) return trip;
    if (imgs is List) {
      trip.images = imgs.cast<String>();
    } else {
      trip.images = imgs['images'] ?? [];
    }

    return trip;
  }
}

class TripArguments {
  String id;
  String banner;

  TripArguments({required this.id, required this.banner});
}
