/* -------------- External -------------- */
import 'package:flutter/material.dart';

class Page extends StatelessWidget {
  const Page({
    Key? key,
    this.padding,
    this.crossAxisAlignment,
    this.onRefresh,
    required this.children,
  }) : super(key: key);

  final EdgeInsetsGeometry? padding;
  final CrossAxisAlignment? crossAxisAlignment;
  final List<Widget> children;
  final Future<void> Function()? onRefresh;

  void nothing() {}

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: SafeArea(
      child: onRefresh != null
          ? RefreshIndicator(onRefresh: onRefresh!, child: _buildPageContent())
          : _buildPageContent(),
    ));
  }

  ListView _buildPageContent() {
    return ListView(padding: padding ?? const EdgeInsets.all(24), children: [
      Column(
        crossAxisAlignment: crossAxisAlignment ?? CrossAxisAlignment.start,
        children: children,
      ),
    ]);
  }
}
