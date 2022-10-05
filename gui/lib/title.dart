import 'dart:async';

import 'package:flutter/material.dart';

import 'gui_state.dart';

class GuiTitle extends StatefulWidget {
  const GuiTitle({super.key});

  @override
  State<GuiTitle> createState() => _GuiTitleState();
}

class _GuiTitleState extends State<GuiTitle> {
  DateTime _current = DateTime.now();

  void refreshTime() {
    _current = DateTime.now();
    Timer(Duration(milliseconds: 1000 - _current.millisecond), () {
      if (mounted) {
        setState(() {
          refreshTime();
        });
      }
    });
  }

  @override
  void initState() {
    super.initState();
    refreshTime();
  }

  String formatDateTime(DateTime dateTime) {
    String year = dateTime.year.toString().padLeft(4, '0');
    String month = dateTime.month.toString().padLeft(2, '0');
    String day = dateTime.day.toString().padLeft(2, '0');
    String hour = dateTime.hour.toString().padLeft(2, '0');
    String minute = dateTime.minute.toString().padLeft(2, '0');
    String second = dateTime.second.toString().padLeft(2, '0');
    String timezone = dateTime.timeZoneName;

    return "$year-$month-$day $hour:$minute:$second $timezone";
  }

  @override
  Widget build(BuildContext context) {
    final String currentUtc = formatDateTime(_current.toUtc());

    final serverReadyIndicator = state.guiServerConnected()
        ? Container(
            color: Colors.green.shade900,
            child: Row(children: <Widget>[
              const Icon(
                Icons.import_export,
                size: 20,
                color: Colors.white,
              ),
              const Icon(
                Icons.radio,
                size: 20,
                color: Colors.white,
              ),
              const Icon(
                Icons.sync,
                size: 20,
                color: Colors.white,
              ),
            ]))
        : Container(
            color: Colors.amber.shade900,
            child: Row(children: <Widget>[
              const Icon(
                Icons.import_export,
                size: 20,
                color: Colors.white,
              ),
              const Icon(
                Icons.radio,
                size: 20,
                color: Colors.white,
              ),
              const Icon(
                Icons.sync,
                size: 20,
                color: Colors.white,
              ),
            ]));

    return Container(
        color: Colors.blue.shade400,
        height: 20,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            serverReadyIndicator,
            const Spacer(),
            Text(currentUtc, style: TextStyle(color: Colors.white)),
          ],
        ));
  }
}
