/* -------------- External -------------- */
import 'package:flutter/material.dart';

/* -------------- Pages -------------- */
import 'package:trips/pages/authentication/login.dart';
import 'package:trips/pages/not_found.dart';
import 'package:trips/pages/home.dart';
import 'package:trips/pages/splash_screen.dart';
import 'package:trips/pages/trip/new_trip.dart';
import 'package:trips/pages/trip/trip.dart';

/* -------------- Types -------------- */
import 'package:trips/types/trip.dart';

class Routes {
  static Route<dynamic> generate(RouteSettings settings) {
    final args = settings.arguments;

    MaterialPageRoute goTo(Widget Function(BuildContext) route) {
      return MaterialPageRoute(builder: route);
    }

    switch (settings.name) {
      case '/':
        return goTo((_) => const HomePage());

      case '/login':
        return goTo((_) => const LoginPage());

      case '/splashscreen':
        return goTo((_) => const SplashPage());

      case '/trip':
        if (args == Null) return goTo((_) => const NotFoundPage());
        final tripArguments = args as TripArguments;
        return goTo((_) =>
            TripPage(tripId: tripArguments.id, banner: tripArguments.banner));

      case '/trip/new':
        return goTo((_) => const NewTrip());

      default:
        return goTo((_) => const NotFoundPage());
    }
  }
}
