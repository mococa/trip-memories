class User {
  String googleId;
  int? createdAt;
  String picture;
  String name;
  String email;

  User({
    required this.googleId,
    this.createdAt,
    required this.picture,
    required this.name,
    required this.email,
  });

  Map<String, dynamic> toJson() {
    Map<String, dynamic> map = {
      'google_id': googleId,
      'email': email,
      'name': name,
      'picture': picture,
    };

    if (createdAt != null) map['created_at'] = createdAt ?? 0;

    return map;
  }

  static User fromJson(dynamic json) => User(
        email: json['email'],
        googleId: json['google_id'],
        name: json['name'],
        picture: json['picture'],
        createdAt: json['created_at'],
      );
}
