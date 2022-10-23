import 'dart:math';

import 'package:file_selector/file_selector.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:mcl_gui/proto/proto/mcl.pbgrpc.dart';
import 'package:mcl_gui/proto/proto/mclgui.pbgrpc.dart';
import 'package:mcl_gui/support.dart';

import 'gui_state.dart';

class Settings extends StatefulWidget {
  @override
  State<Settings> createState() => _SettingsState();
}

class _SettingsState extends State<Settings> {
  @override
  Widget build(BuildContext context) {
    return TextButton(onPressed: () async {
          final filename = await getSavePath(acceptedTypeGroups: <XTypeGroup>[
            XTypeGroup(
              label: 'Amateur Data Interchange Format (*.adif, *.adi)',
              extensions: <String>['adif', 'adi'],
            ),
          ], suggestedName: "${state.activeContest?.contest.contestName}.adif");

          if (filename != null) {
            await state.getGuiClient()?.exportToAdif(OpenFileRequest(name: filename));
          }
    }, child: Text("Export ADIF"));
  }
}
