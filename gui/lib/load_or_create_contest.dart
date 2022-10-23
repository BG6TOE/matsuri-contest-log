import 'package:file_selector/file_selector.dart';
import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:mcl_gui/proto/proto/mcl.pbgrpc.dart';
import 'package:mcl_gui/proto/proto/common.pb.dart';
import 'package:mcl_gui/qso.dart';
import 'package:omni_datetime_picker/omni_datetime_picker.dart';
import 'package:protobuf/protobuf.dart';
import 'package:tuple/tuple.dart';

import 'gui_state.dart';

class CreateContestForm extends StatefulWidget {
  final ContestMetadata info;

  const CreateContestForm({super.key, required this.info});

  @override
  State<CreateContestForm> createState() => _CreateContestFormState();
}

class _CreateContestFormState extends State<CreateContestForm> {
  TextEditingController callsignController = TextEditingController();
  List<TextEditingController> customFieldsController =
      <TextEditingController>[];
  List<DateTime>? contestStartEndTime;
  List<Tuple2<String, ContestFieldInfo>> customFields = [];

  _CreateContestFormState();

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    disposeControllers();
    callsignController.dispose();

    super.dispose();
  }

  void disposeControllers() {
    for (int i = 0; i < customFieldsController.length; i++) {
      customFieldsController[i].dispose();
    }

    customFieldsController = [];
  }

  String _formatBeginEndTime() {
    if (contestStartEndTime == null) {
      return "Set contest begin and end time";
    } else {
      return "${contestStartEndTime![0].toUtc().toIso8601String()} -> ${contestStartEndTime![1].toUtc().toIso8601String()}";
    }
  }

  @override
  Widget build(BuildContext context) {
    final List<Widget> fields = <Widget>[];

    if (customFieldsController.length == 0) {
      for (var info in widget.info.fieldInfos.entries) {
        if (info.value.fieldType == "tx_const" || info.value.fieldType == "info") {
          customFields.add(Tuple2<String, ContestFieldInfo>(info.key, info.value));
        }
      }
      customFieldsController = customFields.map((e) => TextEditingController()).toList();
    }

    fields.add(Text("Contest: ${widget.info.contestName}"));
    fields.add(Text("Version: ${widget.info.version}"));

    fields.add(QSOExchTextField(
      "Station Callsign",
      controller: callsignController,
    ));

    for (int i = 0; i < customFields.length; i++) {
      fields.add(const SizedBox(width: 1, height: 8));
      fields.add(QSOExchTextField(
        customFields[i].item2.displayName,
        controller: customFieldsController[i],
      ));
    }

    fields.add(Row(
      children: [
        TextButton(
            onPressed: () async {
              contestStartEndTime = await showOmniDateTimeRangePicker(
                context: this.context,
                type: OmniDateTimePickerType.dateAndTime,
                primaryColor: Colors.cyan,
                backgroundColor: Colors.grey[900],
                calendarTextColor: Colors.white,
                tabTextColor: Colors.white,
                unselectedTabBackgroundColor: Colors.grey[700],
                buttonTextColor: Colors.white,
                timeSpinnerTextStyle:
                    const TextStyle(color: Colors.white70, fontSize: 18),
                timeSpinnerHighlightedTextStyle:
                    const TextStyle(color: Colors.white, fontSize: 24),
                is24HourMode: true,
                isShowSeconds: false,
                startInitialDate: DateTime.now().toUtc(),
                startFirstDate:
                    DateTime(1600).subtract(const Duration(days: 3652)),
                startLastDate: DateTime.now().add(
                  const Duration(days: 3652),
                ),
                endInitialDate: DateTime.now(),
                endFirstDate:
                    DateTime(1600).subtract(const Duration(days: 3652)),
                endLastDate: DateTime.now().add(
                  const Duration(days: 3652),
                ),
                borderRadius: const Radius.circular(16),
              );
              setState(() {});
            },
            child: Text(_formatBeginEndTime())),
      ],
    ));

    fields.add(TextButton(
        onPressed: () async {
          final filename = await getSavePath(acceptedTypeGroups: <XTypeGroup>[
            XTypeGroup(
              label: 'Contest Database (*.mclcontest.db)',
              extensions: <String>['mclcontest.db'],
            )
          ]);

          var customFieldsMap = <String, String>{};
          for (int i = 0; i < customFields.length; i++) {
            customFieldsMap[customFields[i].item1] = customFieldsController[i].text;
          }

          ActiveContest contest = ActiveContest(contest: widget.info, station: Station(callsign: callsignController.text, customFields: customFieldsMap));
          if (contestStartEndTime != null) {
            contest.beginTimestamp = Int64(
                contestStartEndTime![0].toUtc().millisecondsSinceEpoch ~/ 1000);
            contest.endTimestamp = Int64(
                contestStartEndTime![1].toUtc().millisecondsSinceEpoch ~/ 1000);
          }
          if (filename != null) {
            await state.getGuiClient()!.createContest(
                CreateContestRequest(databaseName: filename, contest: contest));
            await state
                .getGuiClient()!
                .loadContest(LoadContestRequest(databaseName: filename));
            Navigator.pushReplacementNamed(this.context, '/contest');
          }
        },
        child: Text("Save to...")));

    return Column(children: fields);
  }
}

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
  ContestMetadata? contestInfo;

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    const XTypeGroup contestDbTypes = XTypeGroup(
      label: 'Contest Database (*.mclcontest.db)',
      extensions: <String>['mclcontest.db'],
    );
    const XTypeGroup contestDescriptor = XTypeGroup(
      label: 'Contest Descriptor (*.mclcontest.lua)',
      extensions: <String>['mclcontest.lua'],
    );
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      body: Row(
        children: <Widget>[
          Expanded(
            child: Column(
              children: <Widget>[
                SizedBox(
                  height: 8,
                  width: 1,
                ),
                TextButton(
                    onPressed: () async {
                      final filename = await openFile(
                          acceptedTypeGroups: <XTypeGroup>[contestDbTypes]);
                      if (filename != null) {
                        await state.getGuiClient()!.loadContest(
                            LoadContestRequest(databaseName: filename.path));
                        Navigator.pushReplacementNamed(
                            this.context, '/contest');
                      }
                    },
                    child: Text("Open an existing contest database")),
                SizedBox(
                  height: 8,
                  width: 1,
                ),
                Divider(
                  height: 1,
                ),
                SizedBox(
                  height: 8,
                  width: 1,
                ),
                TextButton(
                    onPressed: () async {
                      final file = await openFile(
                          acceptedTypeGroups: <XTypeGroup>[contestDescriptor]);
                      if (file != null) {
                        contestInfo = await state.getGuiClient()!.parseContest(
                            ParseContestRequest(contestDescriptor: await file.readAsString()));
                        setState(() {});
                      }
                    },
                    child: Text("Create an new contest database")),
                SizedBox(
                  height: 8,
                  width: 1,
                ),
                contestInfo != null ? CreateContestForm(info: contestInfo!) : Text(""),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
