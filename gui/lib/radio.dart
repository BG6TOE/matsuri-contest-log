import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:mcl_gui/proto/proto/mcl.pbgrpc.dart';
import 'package:mcl_gui/proto/proto/radio.pbgrpc.dart';
import 'package:mcl_gui/support.dart';
import 'package:tuple/tuple.dart';

import 'gui_state.dart';
import 'theme.dart';

class RadioStatusPanel extends StatefulWidget {
  @override
  State<RadioStatusPanel> createState() => _RadioStatusPanelState();
}

class _RadioStatusPanelState extends State<RadioStatusPanel> {
  int callbackId = 0;
  List<Tuple2<String, RadioStatus>> radios = [];

  void showRadioConfigDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: const Text('Radio Settings'),
          content: SingleChildScrollView(
            child: ListBody(
              children: const <Widget>[
                Text('This action cannot be undone'),
                Text('If in doubt press the cancel button to go back'),
              ],
            ),
          ),
          actions: <Widget>[
            TextButton(
              child: const Text('CANCEL'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
            TextButton(
              child: const Text('ERASE'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

  @override
  void initState() {
    super.initState();
    callbackId = state.registerCallback(
        CallbackKind.kRadioStatusUpdate, _updateRadioList);
    _updateRadioList();
  }

  @override
  void dispose() {
    state.deregisterCallback(CallbackKind.kRadioStatusUpdate, callbackId);

    super.dispose();
  }

  void _updateRadioList() {
    setState(() {
      radios = state.radios.entries.map((e) => Tuple2(e.key, e.value)).toList();
      radios.sort((a, b) => a.item1.compareTo(b.item1));
    });
  }

  @override
  Widget build(BuildContext context) {
    var radioList = ListBody(
      children: radios
          .map((e) => ListTile(
                title: Text(
                    "Rx: ${e.item2.rx.mode.toString()} ${e.item2.rx.frequency} / Tx: ${e.item2.tx.mode.toString()} ${e.item2.tx.frequency}"),
                subtitle: Text("${e.item1} ${e.item2.uri}"),
              ))
          .toList(),
    );

    return Column(
        // Column is also a layout widget. It takes a list of children and
        // arranges them vertically. By default, it sizes itself to fit its
        // children horizontally, and tries to be as tall as its parent.
        //
        // Invoke "debug painting" (press "p" in the console, choose the
        // "Toggle Debug Paint" action from the Flutter Inspector in Android
        // Studio, or the "Toggle Debug Paint" command in Visual Studio Code)
        // to see the wireframe for each widget.
        //
        // Column has various properties to control how it sizes itself and
        // how it positions its children. Here we use mainAxisAlignment to
        // center the children vertically; the main axis here is the vertical
        // axis because Columns are vertical (the cross axis would be
        // horizontal).
        mainAxisAlignment: MainAxisAlignment.end,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          const SizedBox(width: 1, height: 8),
          radioList,
          TextButton(
              onPressed: () {
                showRadioConfigDialog(context);
              },
              child: Text("Add New Radio")),
        ]);
  }
}
