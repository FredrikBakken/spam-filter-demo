import 'dart:ui';

import 'package:client/services/firestore_upload.dart';
import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';

class MessagingPage extends StatefulWidget {
  final String username;

  MessagingPage({Key key, @required this.username}) : super(key: key);

  @override
  _MessagingPageState createState() {
    return _MessagingPageState();
  }
}

class _MessagingPageState extends State<MessagingPage> {
  final _formKey = GlobalKey<FormState>();
  bool isValid = false;
  TextEditingController _message = new TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Spam Filter Demo | Messaging"),
      ),
      body: Padding(
        padding: const EdgeInsets.only(bottom: 80.0),
        child: Container(
          child: StreamBuilder(
            stream: FirebaseFirestore.instance
                .collection('messages')
                .orderBy('timestamp', descending: true)
                .snapshots(),
            builder: (context, snapshot) {
              if (!snapshot.hasData) {
                return Center(
                  child: CircularProgressIndicator(
                    valueColor: AlwaysStoppedAnimation<Color>(Colors.blue),
                  ),
                );
              } else {
                return ListView.builder(
                  padding: EdgeInsets.all(10.0),
                  itemBuilder: (context, index) => Container(
                    child: Row(
                      children: [
                        Text(
                          snapshot.data.documents[index]['username'] + ": ",
                          style: TextStyle(fontWeight: FontWeight.bold),
                        ),
                        Text(snapshot.data.documents[index]['message']
                            .toString()),
                      ],
                    ),
                  ),
                  // buildItem(context, snapshot.data.documents[index]),
                  itemCount: snapshot.data.documents.length,
                  reverse: true,
                );
              }
            },
          ),
        ),
      ),
      bottomSheet: Container(
        height: 80,
        width: double.infinity,
        decoration: BoxDecoration(color: Color(0xFFe9eaec).withOpacity(0.2)),
        child: Padding(
          padding: const EdgeInsets.only(left: 18, right: 18),
          child: Form(
            key: _formKey,
            child: Container(
              width: MediaQuery.of(context).size.width,
              child: Row(
                children: <Widget>[
                  Container(
                    width: (MediaQuery.of(context).size.width) * 0.75,
                    height: 40,
                    decoration: BoxDecoration(
                        color: Color(0xFFe9eaec),
                        borderRadius: BorderRadius.circular(20)),
                    child: Padding(
                      padding: const EdgeInsets.only(left: 12),
                      child: TextFormField(
                        cursorColor: Color(0xFF000000),
                        controller: _message,
                        decoration: InputDecoration(
                          border: InputBorder.none,
                          hintText: "Aa",
                        ),
                        textCapitalization: TextCapitalization.sentences,
                        validator: (value) {
                          if (value.isEmpty) {
                            return 'Please enter a message';
                          }
                          return null;
                        },
                      ),
                    ),
                  ),
                  Spacer(),
                  GestureDetector(
                    onTap: () {
                      if (_formKey.currentState.validate()) {
                        createRecord(widget.username, _message.text);
                        _formKey.currentState.reset();
                      }
                    },
                    child: Icon(
                      Icons.send,
                      size: 35,
                      color: Colors.blue,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}

/*
Widget buildItem(BuildContext context, DocumentSnapshot document) {
    if (document.data()['id'] == currentUserId) {
      return Container();
    } else {
      return Container(
        child: FlatButton(
          child: Row(
            children: <Widget>[
              Material(
                child: document.data()['photoUrl'] != null
                    ? CachedNetworkImage(
                        placeholder: (context, url) => Container(
                          child: CircularProgressIndicator(
                            strokeWidth: 1.0,
                            valueColor:
                                AlwaysStoppedAnimation<Color>(themeColor),
                          ),
                          width: 50.0,
                          height: 50.0,
                          padding: EdgeInsets.all(15.0),
                        ),
                        imageUrl: document.data()['photoUrl'],
                        width: 50.0,
                        height: 50.0,
                        fit: BoxFit.cover,
                      )
                    : Icon(
                        Icons.account_circle,
                        size: 50.0,
                        color: greyColor,
                      ),
                borderRadius: BorderRadius.all(Radius.circular(25.0)),
                clipBehavior: Clip.hardEdge,
              ),
              Flexible(
                child: Container(
                  child: Column(
                    children: <Widget>[
                      Container(
                        child: Text(
                          'Nickname: ${document.data()['nickname']}',
                          style: TextStyle(color: primaryColor),
                        ),
                        alignment: Alignment.centerLeft,
                        margin: EdgeInsets.fromLTRB(10.0, 0.0, 0.0, 5.0),
                      ),
                      Container(
                        child: Text(
                          'About me: ${document.data()['aboutMe'] ?? 'Not available'}',
                          style: TextStyle(color: primaryColor),
                        ),
                        alignment: Alignment.centerLeft,
                        margin: EdgeInsets.fromLTRB(10.0, 0.0, 0.0, 0.0),
                      )
                    ],
                  ),
                  margin: EdgeInsets.only(left: 20.0),
                ),
              ),
            ],
          ),
          onPressed: () {
            Navigator.push(
                context,
                MaterialPageRoute(
                    builder: (context) => Chat(
                          peerId: document.id,
                          peerAvatar: document.data()['photoUrl'],
                        )));
          },
          color: greyColor2,
          padding: EdgeInsets.fromLTRB(25.0, 10.0, 25.0, 10.0),
          shape:
              RoundedRectangleBorder(borderRadius: BorderRadius.circular(10.0)),
        ),
        margin: EdgeInsets.only(bottom: 10.0, left: 5.0, right: 5.0),
      );
    }
  }
}*/
