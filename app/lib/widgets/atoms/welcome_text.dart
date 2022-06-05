/* -------------- External -------------- */
import 'package:flutter/material.dart';

class WelcomeText extends StatelessWidget {
  const WelcomeText({Key? key, required this.name}) : super(key: key);

  final String name;

  @override
  Widget build(BuildContext context) {
    return Text.rich(
      TextSpan(
          text: "Welcome,\n",
          style: const TextStyle(fontWeight: FontWeight.bold),
          children: [
            TextSpan(
                text: name,
                style: const TextStyle(fontWeight: FontWeight.normal))
          ]),
      style: const TextStyle(fontSize: 50),
    );
  }
}
