import 'dart:async';
import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

void main() {
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
      home: const GuiMainPage(title: 'Flutter Demo Home Page'),
    );
  }
}

class QSOInput extends StatefulWidget {
  const QSOInput({super.key});

  @override
  State<QSOInput> createState() => _QSOInputState();
}

class UpperCaseTextFormatter extends TextInputFormatter {
  @override
  TextEditingValue formatEditUpdate(
      TextEditingValue oldValue, TextEditingValue newValue) {
    return TextEditingValue(
      text: newValue.text.toUpperCase(),
      selection: newValue.selection,
    );
  }
}

class _QSOInputState extends State<QSOInput> {
  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      children: <Widget>[
        Expanded(
            child: TextField(
          style: const TextStyle(fontFamily: "monospace"),
          decoration: const InputDecoration(
            border: OutlineInputBorder(),
            labelText: 'Callsign',
            labelStyle: TextStyle(fontFamily: "serif"),
          ),
          inputFormatters: <TextInputFormatter>[
            LengthLimitingTextInputFormatter(10),
            UpperCaseTextFormatter(),
          ],
        )),
        const SizedBox(width: 4, height: 1),
        Expanded(
            child: TextField(
          style: const TextStyle(fontFamily: "monospace"),
          decoration: const InputDecoration(
            border: OutlineInputBorder(),
            labelStyle: TextStyle(fontFamily: "serif"),
            labelText: 'RST Sent',
          ),
          inputFormatters: <TextInputFormatter>[
            LengthLimitingTextInputFormatter(5),
            UpperCaseTextFormatter(),
          ],
        )),
        const SizedBox(width: 4, height: 1),
        Expanded(
            child: TextField(
          style: const TextStyle(fontFamily: "monospace"),
          decoration: const InputDecoration(
            border: OutlineInputBorder(),
            labelStyle: TextStyle(fontFamily: "serif"),
            labelText: 'RST Rcvd',
          ),
          inputFormatters: <TextInputFormatter>[
            LengthLimitingTextInputFormatter(5),
            UpperCaseTextFormatter(),
          ],
        )),
        const SizedBox(width: 4, height: 1),
        Expanded(
            child: TextField(
          style: const TextStyle(fontFamily: "monospace"),
          decoration: const InputDecoration(
            border: OutlineInputBorder(),
            labelStyle: TextStyle(fontFamily: "serif"),
            labelText: 'Exch Rcvd',
          ),
          inputFormatters: <TextInputFormatter>[
            UpperCaseTextFormatter(),
          ],
        )),
      ],
    );
  }
}

class HotkeyPanel extends StatefulWidget {
  const HotkeyPanel({super.key});

  @override
  State<HotkeyPanel> createState() => _HotkeyPanelState();
}

class HotKeyButton extends TextButton {
  final String text;
  HotKeyButton({super.key, required this.text, required super.onPressed})
      : super(
            child: Text(
              text,
              textAlign: TextAlign.left,
              textWidthBasis: TextWidthBasis.parent,
              softWrap: false,
            ),
            style: ButtonStyle(
                backgroundColor: MaterialStateProperty.all(Colors.black12)));
}

class _HotkeyPanelState extends State<HotkeyPanel> {
  @override
  Widget build(BuildContext context) {
    return Column(mainAxisAlignment: MainAxisAlignment.start, children: <
        Widget>[
      Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Expanded(child: HotKeyButton(text: "F1 [CQ]", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F2 [Exch]", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F3 TU", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F4 BG6TOE", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(
              child: HotKeyButton(text: "F5 [Their Call]", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F6 [Repeat]", onPressed: () {})),
        ],
      ),
      const SizedBox(width: 1, height: 4),
      Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          Expanded(child: HotKeyButton(text: "F7 [Spare]", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F8 AGN?", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F9 NR?", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F10 CL?", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F11 [Spare]", onPressed: () {})),
          const SizedBox(width: 4, height: 1),
          Expanded(child: HotKeyButton(text: "F12 [Wipe]", onPressed: () {})),
        ],
      ),
    ]);
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
      setState(() {
        refreshTime();
      });
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
    final String currentLocal = formatDateTime(_current);

    return Text("Contest Name | $currentUtc | $currentLocal | Current Frequency");
  }
}

class _GuiMainPageState extends State<GuiMainPage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
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
            Expanded(
                child: Column(
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
                  GuiTitle(),
                  const SizedBox(width: 1, height: 8),
                  const QSOInput(),
                  const SizedBox(width: 1, height: 4),
                  const Text(
                    "BY AS/CHINA, Zn 24, ITU 44",
                    textAlign: TextAlign.left,
                    textWidthBasis: TextWidthBasis.parent,
                  ),
                  const SizedBox(width: 1, height: 4),
                  const HotkeyPanel(),
                ])),
            const SizedBox(width: 8, height: 1),
          ]),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
