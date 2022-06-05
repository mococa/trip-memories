/* -------------- External -------------- */
import 'dart:io';
import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'package:lightbox/lightbox.dart';
import 'package:provider/provider.dart';

/* -------------- Providers -------------- */
import 'package:trips/providers/authentication.dart';

/* -------------- Services -------------- */
import 'package:trips/services/trips.dart';

/* -------------- Types -------------- */
import 'package:trips/types/user.dart';

/* -------------- Widgets -------------- */
import 'package:trips/widgets/atoms/icon_button.dart' as wdgt;
import 'package:trips/widgets/atoms/image_button.dart';
import 'package:trips/widgets/atoms/page.dart' as app;

class NewTrip extends StatefulWidget {
  const NewTrip({Key? key}) : super(key: key);

  @override
  State<NewTrip> createState() => _NewTripState();
}

class _NewTripState extends State<NewTrip> {
  final nameController = TextEditingController();
  final descriptionController = TextEditingController();
  List<File> images = [];
  DateTime doneAt = DateTime.now();
  bool loading = false;

  Future handleAddImage() async {
    try {
      final img = await ImagePicker().pickImage(source: ImageSource.gallery);

      if (img == null) return;

      setState(() {
        images = [...images, File(img.path)];
      });
    } catch (err) {
      print(err);
    }
  }

  handleRemoveImage(File file) {
    ScaffoldMessenger.of(context).showSnackBar(
      const SnackBar(
        content: Text('Image removed successfully'),
      ),
    );

    setState(() {
      images = images.where((img) => img.path != file.path).toList();
    });
  }

  Future handleCreateTrip() async {
    setState(() {
      loading = true;
    });

    User? user = Provider.of<Authentication>(context, listen: false).user;

    if (user == null || images.isEmpty) return;

    try {
      await Trips.createTrip(user, nameController.text,
          descriptionController.text, doneAt, images);
    } catch (err) {
      print(err);
    }

    setState(() {
      loading = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    return app.Page(
      children: [
        const Text(
          'Enregister a memory',
          style: TextStyle(fontSize: 32, fontWeight: FontWeight.w500),
        ),
        label('Title'),
        TextField(
          controller: nameController,
          decoration:
              const InputDecoration(filled: true, hintText: 'Rio de Janeiro'),
        ),
        label('Description'),
        TextFormField(
          controller: descriptionController,
          minLines: 5,
          maxLines: 5,
          keyboardType: TextInputType.multiline,
          decoration: const InputDecoration(
              filled: true, hintText: "We've had our best trip ever"),
        ),
        Wrap(
            crossAxisAlignment: WrapCrossAlignment.center,
            spacing: 4,
            children: [
              label('Date'),
              Column(
                children: const [
                  SizedBox(height: 16),
                  Text('(DD/MM/YYYY)', style: TextStyle(color: Colors.black54)),
                  SizedBox(height: 4)
                ],
              )
            ]),
        datePicker(),
        const SizedBox(height: 10),
        SizedBox(
          width: MediaQuery.of(context).size.width,
          height: 100,
          child: SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: Wrap(
                crossAxisAlignment: WrapCrossAlignment.start,
                spacing: 10,
                children: images.map(buildPicture).toList()
                  ..add(buildAddPicture())),
          ),
        ),
        const SizedBox(height: 10),
        ElevatedButton(
            onPressed: loading ? null : handleCreateTrip,
            style: ElevatedButton.styleFrom(
              minimumSize: const Size.fromHeight(40),
            ),
            child: Text(loading ? 'Creating' : 'Create'))
      ],
    );
  }

  Widget buildPicture(File file) {
    return SizedBox(
      width: 100,
      height: 100,
      child: ImageButton(
        borderRadius: 4,
        image: Image.file(file).image,
        onTap: () {
          Navigator.push(context, LightBoxRoute(
            builder: (BuildContext context) {
              return LightBox(
                imageType: ImageType.FILE,
                images: images.map((e) => e.path).toList(),
              );
            },
          ));
        },
        onLongPress: () => handleRemoveImage(file),
      ),
    );
  }

  Widget buildAddPicture() {
    return SizedBox(
      width: 100,
      height: 100,
      child: wdgt.IconButton(
        icon: Icons.add_a_photo,
        onTap: handleAddImage,
      ),
    );
  }

  Widget label(String txt) {
    return Column(children: [
      const SizedBox(height: 16),
      Text(
        txt,
        style: const TextStyle(fontSize: 16),
      ),
      const SizedBox(height: 4)
    ]);
  }

  Widget datePicker() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          '${doneAt.day.toString().padLeft(2, '0')}/${doneAt.month.toString().padLeft(2, '0')}/${doneAt.year}',
          style: const TextStyle(
            fontSize: 16,
          ),
        ),
        ElevatedButton(
            onPressed: () async {
              DateTime? newDate = await showDatePicker(
                  context: context,
                  initialDate: doneAt,
                  firstDate: DateTime(1900),
                  lastDate: DateTime.now());

              if (newDate != null) {
                setState(() {
                  doneAt = newDate;
                });
              }
            },
            child: const Text('Select date'))
      ],
    );
  }
}
