/* -------------- External -------------- */
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

/* -------------- Providers -------------- */
import 'package:trips/providers/authentication.dart';

/* -------------- Services -------------- */
import 'package:trips/services/trips.dart';
import 'package:trips/services/oauth.dart';

/* -------------- Types -------------- */
import 'package:trips/types/trip.dart';

/* -------------- Components -------------- */
import '../widgets/atoms/image_button.dart';
import '../widgets/atoms/welcome_text.dart';
import '../widgets/molecules/top_bar.dart';
import 'package:trips/widgets/atoms/page.dart' as app;

final places = ['floripa', 'sao-paulo', 'minas-gerais', 'rio'];

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  State<HomePage> createState() => _HomePage();
}

class _HomePage extends State<HomePage> {
  List<Trip> _trips = [];

  void handleLogout(BuildContext context) async {
    await OAuth.clearSession(context);
  }

  void handleCreateTrip(BuildContext context) async {
    Navigator.pushNamed(context, '/trip/new');
  }

  void _handleGetTrips() async {
    final foundTrips = await Trips().find();

    setState(() {
      _trips = foundTrips;
    });
  }

  @override
  void initState() {
    _handleGetTrips();

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return app.Page(
        padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
        children: [
          TopBar(
            onCreateMemory: () => handleCreateTrip(context),
            onLogout: () => handleLogout(context),
          ),
          WelcomeText(name: context.watch<Authentication>().user?.name ?? ''),
          const SizedBox(
            height: 20,
          ),
          TextField(
            decoration: InputDecoration(
                prefixIcon: const Icon(Icons.search),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
                hintText: 'Search'),
            style: const TextStyle(fontSize: 18.0),
          ),
          const SizedBox(
            height: 54.0,
          ),
          const Text(
            'Saved Memories',
            style: TextStyle(fontWeight: FontWeight.w600, fontSize: 20),
          ),
          const SizedBox(height: 8),
          GridView.count(
            crossAxisCount: 2,
            mainAxisSpacing: 8,
            crossAxisSpacing: 8,
            shrinkWrap: true,
            physics:
                const ScrollPhysics(parent: NeverScrollableScrollPhysics()),
            children: _trips.map(buildTripButton).toList(),
          )
        ]);
  }

  Widget buildTripButton(Trip trip) {
    return Hero(
      tag: 'trip-banner-${trip.id}',
      child: ImageButton(
          image: AssetImage(trip.banner),
          onTap: () {
            Navigator.of(context).pushNamed('/trip',
                arguments: TripArguments(id: trip.id, banner: trip.banner));
          }),
    );
  }
}
