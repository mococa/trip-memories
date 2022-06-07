/* -------------- External -------------- */
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart' as material;

/* -------------- Widgets -------------- */
import 'package:trips/widgets/atoms/icon_button.dart';

class TopBar extends StatelessWidget {
  const TopBar({Key? key, required this.onCreateMemory, required this.onLogout})
      : super(key: key);

  final void Function() onCreateMemory;
  final void Function() onLogout;

  @override
  Widget build(BuildContext context) {
    return (Row(
      mainAxisAlignment: MainAxisAlignment.end,
      children: [
        material.TextButton(
            onPressed: onLogout,
            child: const Text(
              'Logout',
              style: TextStyle(color: material.Colors.black),
            )),
        IconButton(onTap: onCreateMemory, icon: material.Icons.add),
      ],
    ));
  }
}
