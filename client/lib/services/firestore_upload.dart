import 'package:cloud_firestore/cloud_firestore.dart';

void createRecord(String username, String message) async {
  await FirebaseFirestore.instance.collection("messages").add({
    'message': message,
    'username': username,
    'timestamp': DateTime.now().microsecondsSinceEpoch,
    'ham-or-spam': "",
  });
}
