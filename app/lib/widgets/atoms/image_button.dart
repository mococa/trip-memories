/* -------------- External -------------- */
import 'package:flutter/material.dart';

class ImageButton extends StatelessWidget {
  const ImageButton(
      {Key? key,
      this.borderRadius,
      required this.image,
      required this.onTap,
      this.onLongPress})
      : super(key: key);

  final ImageProvider<Object> image;
  final void Function() onTap;
  final void Function()? onLongPress;
  final double? borderRadius;
  final double defaultBorderRadius = 8.0;

  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      borderRadius: BorderRadius.circular(borderRadius ?? defaultBorderRadius),
      child: Material(
        child: Ink.image(
          image: image,
          fit: BoxFit.cover,
          child: InkWell(
            onTap: onTap,
            onLongPress: onLongPress,
          ),
        ),
      ),
    );
  }
}
