/* -------------- External -------------- */
import 'package:flutter/material.dart';
import 'package:lightbox/lightbox.dart';

/* -------------- Services -------------- */
import 'package:trips/services/trips.dart';

/* -------------- Types -------------- */
import 'package:trips/types/trip.dart';

/* -------------- Widgets -------------- */
import 'package:trips/widgets/atoms/image_button.dart';

class TripPage extends StatefulWidget {
  const TripPage({Key? key, required this.tripId, required this.banner})
      : super(key: key);

  final String tripId;
  final String banner;

  @override
  State<TripPage> createState() => _TripPageState();
}

class _TripPageState extends State<TripPage> {
  Trip? trip;

  _handleGetTrip(String id) async {
    final myTrip = await Trips().findById(id);

    setState(() {
      trip = myTrip;
    });
  }

  @override
  void initState() {
    _handleGetTrip(widget.tripId);

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final double sheetHeight = (MediaQuery.of(context).size.height - 250) /
        MediaQuery.of(context).size.height;

    if (trip != null) {
      WidgetsBinding.instance.addPostFrameCallback((timeStamp) {
        showModalBottomSheet(
            context: context,
            barrierColor: null,
            isDismissible: true,
            backgroundColor: Colors.transparent,
            isScrollControlled: true,
            enableDrag: false,
            builder: (BuildContext acontext) {
              return DraggableScrollableSheet(
                initialChildSize: sheetHeight,
                minChildSize: 0.35,
                maxChildSize: 0.95,
                // expand: true,
                builder: tripContent,
              );
            }).whenComplete(() => Navigator.pop(context));
      });
    }

    return Scaffold(
      body: Column(
        children: [
          Hero(
            tag: 'trip-banner-${widget.tripId}',
            child: Image(
              image: AssetImage(widget.banner),
              fit: BoxFit.cover,
              height: 260,
              width: MediaQuery.of(context).size.width,
            ),
          ),
          Expanded(
            flex: 1,
            child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: trip == null
                    ? [
                        const CircularProgressIndicator(
                          color: Colors.amber,
                        )
                      ]
                    : []),
          )
        ],
      ),
    );
  }

  Widget buildPicture(String imgUrl) {
    return SizedBox(
        width: 100,
        height: 100,
        child: ImageButton(
            borderRadius: 4,
            image: Image.network(imgUrl).image,
            onTap: () {
              Navigator.push(context, LightBoxRoute(
                builder: (BuildContext context) {
                  return LightBox(
                    imageType: ImageType.URL,
                    images: const [
                      'https://user-images.githubusercontent.com/42270511/114654248-b8381a00-9d24-11eb-9a9d-727efe4ce1d4.png'
                    ],
                  );
                },
              ));
            }));
  }

  Widget tripImages() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Text("Photos",
            style: TextStyle(fontSize: 16, fontWeight: FontWeight.w500)),
        const SizedBox(height: 4),
        SizedBox(
          width: MediaQuery.of(context).size.width,
          height: 100,
          child: SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: Wrap(
              crossAxisAlignment: WrapCrossAlignment.start,
              spacing: 10,
              children: const [
                'https://user-images.githubusercontent.com/42270511/114654248-b8381a00-9d24-11eb-9a9d-727efe4ce1d4.png'
              ].map(buildPicture).toList(),
            ),
          ),
        ),
      ],
    );
  }

  Widget tripContent(_, controller) => Container(
        decoration: const BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.vertical(top: Radius.circular(20)),
        ),
        padding: const EdgeInsets.all(24),
        // height: MediaQuery.of(context).size.height - 240,
        child: ListView(
          controller: controller,
          children: [
            Column(children: [
              Text(trip?.title ?? '',
                  style: const TextStyle(
                      fontWeight: FontWeight.w600, fontSize: 24)),
              const SizedBox(height: 4),
              Text(trip?.description ?? ''),
              const SizedBox(height: 8),
              tripImages()
            ]),
          ],
        ),
      );
}
