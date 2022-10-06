
import 'package:file_selector/file_selector.dart';
import 'package:flutter/material.dart';

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
      body: Column(
        children: <Widget>[
          TextButton(
              onPressed: () async {
                await openFile(acceptedTypeGroups: <XTypeGroup>[contestDbTypes]);
              },
              child: Text("Open an existing contest database")),
          
          TextButton(
              onPressed: () async {
                await openFile(acceptedTypeGroups: <XTypeGroup>[contestDescriptor]);
              },
              child: Text("Create an new contest database for a contest")),
          
          TextButton(
              onPressed: () async {
                await getSavePath(suggestedName: 'NewContest.mcldb.sqlite3', acceptedTypeGroups: <XTypeGroup>[contestDbTypes]);
              },
              child: Text("Create and open contest")),
        ],
      ),
    );
  }
}
