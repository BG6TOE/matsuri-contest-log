import 'dart:async';
import 'dart:ui';

import 'package:file_selector/file_selector.dart';
import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'gui_state.dart';
import 'proto/proto/mcl.pbenum.dart';
import 'proto/proto/mclgui.pbgrpc.dart';
import 'qso.dart';
import 'title.dart';

class LoadOrCreateContestPage extends StatefulWidget {
  const LoadOrCreateContestPage({super.key});

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  @override
  State<LoadOrCreateContestPage> createState() =>
      _LoadOrCreateContestPageState();
}

class _LoadOrCreateContestPageState extends State<LoadOrCreateContestPage> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final XTypeGroup typeGroup = const XTypeGroup(
      label: 'Contest Database',
      extensions: <String>['mcldb.sqlite3'],
    );
    final XTypeGroup contestDescriptor = const XTypeGroup(
      label: 'Contest Descriptor',
      extensions: <String>['mclcontest.lua'],
    );
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      body: Column(
        children: <Widget>[
          TextButton(
              onPressed: () async {
                await openFile(acceptedTypeGroups: <XTypeGroup>[typeGroup]);
              },
              child: Text("Open Database")),
        ],
      ),
    );
  }
}
