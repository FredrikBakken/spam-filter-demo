import 'package:client/ui/views/login_page.dart';
import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp();
  runApp(ClientDemo());
}

class ClientDemo extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      initialRoute: '/login',
      routes: {
        // When navigating to the "/" route, build the FirstScreen widget.
        '/login': (context) => LoginForm(),
        // When navigating to the "/second" route, build the SecondScreen widget.
        // '/messaging': (context) => MessagingPage(),
      },
    );
  }
}

class FirstRoute extends StatefulWidget {
  FirstRoute({Key key, this.title}) : super(key: key);
  final String title;

  @override
  _FirstRouteState createState() => _FirstRouteState();
}

class _FirstRouteState extends State<FirstRoute> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text("Spam Filter Demo"),
        ),
        body: FutureBuilder(
          future: getData(),
          builder: (context, AsyncSnapshot<DocumentSnapshot> snapshot) {
            print(snapshot);
            print(snapshot.data);
            if (snapshot.connectionState == ConnectionState.done) {
              return Column(
                children: [
                  Container(
                    height: 27,
                    child: Text(
                      "Name: ${snapshot.data.data()['name']}",
                      overflow: TextOverflow.fade,
                      style: TextStyle(fontSize: 20),
                    ),
                  ),
                ],
              );
            } else if (snapshot.connectionState == ConnectionState.none) {
              return Text("No data");
            }
            return CircularProgressIndicator();
          },
        ));
  }

  Future<DocumentSnapshot> getData() async {
    return await FirebaseFirestore.instance
        .collection("users")
        .doc("docID")
        .get();
  }
}
