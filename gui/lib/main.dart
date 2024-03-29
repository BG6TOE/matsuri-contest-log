import 'dart:async';
import 'dart:ui';

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:mcl_gui/log.dart';
import 'package:mcl_gui/proto/proto/mcl.pb.dart';
import 'package:mcl_gui/proto/proto/mcl.pbjson.dart';
import 'load_or_create_contest.dart';
import 'gui_state.dart';
import 'proto/proto/mcl.pbenum.dart';
import 'proto/proto/mclgui.pbgrpc.dart';
import 'qso.dart';
import 'settings.dart';
import 'title.dart';
import 'radio.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  state.startServer();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Matsuri Contest Log',
      theme: ThemeData(
          // This is the theme of your application.
          //
          // Try running your application with "flutter run". You'll see the
          // application has a blue toolbar. Then, without quitting the app, try
          // changing the primarySwatch below to Colors.green and then invoke
          // "hot reload" (press "r" in the console where you ran "flutter run",
          // or simply save your changes to "hot reload" in a Flutter IDE).
          // Notice that the counter didn't reset back to zero; the application
          // is not restarted.
          primarySwatch: Colors.blue,
          useMaterial3: true),
      initialRoute: '/open-contest',
      routes: {
        '/contest': (context) =>
            const GuiMainPage(title: 'Flutter Demo Home Page'),
        '/open-contest': (context) => const LoadOrCreateContestPage(),
      },
    );
  }
}

class GuiMainPage extends StatefulWidget {
  const GuiMainPage({super.key, required this.title});

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<GuiMainPage> createState() => _GuiMainPageState();
}

class _GuiMainPageState extends State<GuiMainPage> {
  int _currentPage = 3;
  QsoTable _qsoTable = QsoTable();
  QsoPanel _qsoPanel = QsoPanel();
  RadioStatusPanel _radioStatus = RadioStatusPanel();

  @override
  void initState() {
    super.initState();

    Timer.run(() {
      state.refreshQsos();
    });
  }

  MainAxisAlignment currentAlignMain() {
    switch (_currentPage) {
      case 3:
        return MainAxisAlignment.end;
      default:
        return MainAxisAlignment.start;
    }
  }

  CrossAxisAlignment currentAlignCross() {
    switch (_currentPage) {
      default:
        return CrossAxisAlignment.start;
    }
  }

  Widget currentWidget() {
    switch (_currentPage) {
      case 4:
        return _qsoPanel;
      case 3:
        return _qsoTable;
      case 2:
        return _radioStatus;
      case 0:
        return Settings();
      default:
        return Text("Select a tab on the left");
    }
  }

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      body: Row(
        children: <Widget>[
          NavigationRail(
            selectedIndex: _currentPage,
            leading: ClockText(),
            groupAlignment: 1.0,
            onDestinationSelected: (int index) {
              setState(() {
                _currentPage = index;
              });
            },
            labelType: NavigationRailLabelType.all,
            destinations: const <NavigationRailDestination>[
              NavigationRailDestination(
                icon: Icon(Icons.settings),
                selectedIcon: Icon(Icons.settings),
                label: Text('Settings'),
              ),
              NavigationRailDestination(
                icon: Icon(Icons.lan),
                selectedIcon: Icon(Icons.lan),
                label: Text('Network'),
              ),
              NavigationRailDestination(
                icon: Icon(Icons.radio),
                selectedIcon: Icon(Icons.radio),
                label: Text('Radio'),
              ),
              NavigationRailDestination(
                icon: Icon(Icons.receipt_long),
                selectedIcon: Icon(Icons.receipt_long),
                label: Text('Log'),
              ),
              NavigationRailDestination(
                icon: Icon(Icons.book),
                selectedIcon: Icon(Icons.book),
                label: Text('QSO'),
              ),
            ],
          ),
          const VerticalDivider(thickness: 1, width: 1),
          Expanded(
            child: Column(children: <Widget>[
              Row(
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
                  mainAxisAlignment: MainAxisAlignment.start,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: <Widget>[
                    const SizedBox(width: 8, height: 1),
                    Expanded(child: currentWidget()),
                    const SizedBox(width: 8, height: 1),
                  ]),
              const SizedBox(width: 1, height: 8),
            ]),
          ),
        ],
      ),
    );
  }
}
