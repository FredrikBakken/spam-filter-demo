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
                        snapshot.data.documents[index]['ham-or-spam'] == ""
                            ? SizedBox(
                                width: 24.0,
                                height: 24.0,
                                child: Center(
                                  child: SizedBox(
                                    width: 20.0,
                                    height: 20.0,
                                    child: CircularProgressIndicator(),
                                  ),
                                ),
                              )
                            : snapshot.data.documents[index]['ham-or-spam'] ==
                                    "false"
                                ? SizedBox(
                                    width: 24.0,
                                    height: 24.0,
                                    child: Icon(
                                      Icons.check,
                                      color: Colors.green,
                                    ),
                                  )
                                : SizedBox(
                                    width: 24.0,
                                    height: 24.0,
                                    child: Icon(
                                      Icons.warning,
                                      color: Colors.red,
                                    ),
                                  ),
                        SizedBox(width: 12.0),
                        Container(
                          width: MediaQuery.of(context).size.width * 0.8,
                          child: Row(
                            children: [
                              Text(
                                snapshot.data.documents[index]['username'] +
                                    ": ",
                                style: TextStyle(fontWeight: FontWeight.bold),
                              ),
                              SizedBox(width: 8.0),
                              Expanded(
                                  child: snapshot.data.documents[index]
                                              ['ham-or-spam'] ==
                                          "true"
                                      ? Text(
                                          "(This is a spam message) " +
                                              snapshot.data
                                                  .documents[index]['message']
                                                  .toString(),
                                          textAlign: TextAlign.left,
                                        )
                                      : Text(
                                          snapshot
                                              .data.documents[index]['message']
                                              .toString(),
                                          textAlign: TextAlign.left,
                                        )),
                            ],
                          ),
                        ),
                      ],
                    ),
                  ),
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
                    width: (MediaQuery.of(context).size.width - 80),
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
