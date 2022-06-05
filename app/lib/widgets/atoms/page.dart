/* -------------- External -------------- */
import 'package:flutter/material.dart';

class Page extends StatelessWidget {
  const Page({
    Key? key,
    this.padding,
    this.crossAxisAlignment,
    required this.children,
  }) : super(key: key);

  final EdgeInsetsGeometry? padding;
  final CrossAxisAlignment? crossAxisAlignment;
  final List<Widget> children;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: SafeArea(
            child: SingleChildScrollView(
      padding: padding ?? const EdgeInsets.all(24),
      child: Column(
        crossAxisAlignment: crossAxisAlignment ?? CrossAxisAlignment.start,
        children: children,
      ),
    )));
  }
}
