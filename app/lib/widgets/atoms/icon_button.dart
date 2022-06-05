/* -------------- External -------------- */
import 'package:flutter/material.dart';

class IconButton extends StatelessWidget {
  const IconButton(
      {Key? key,
      required this.onTap,
      required this.icon,
      this.padding,
      this.size})
      : super(key: key);

  final void Function() onTap;
  final IconData icon;
  final double? padding;
  final double? size;

  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      borderRadius: BorderRadius.circular(50),
      child: Material(
        child: InkWell(
          onTap: onTap,
          child: Padding(
            padding: EdgeInsets.all(padding ?? 12),
            child: Icon(
              icon,
              size: size ?? 24,
            ),
          ),
        ),
      ),
    );
  }
}
