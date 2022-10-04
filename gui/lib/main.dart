import 'dart:async';
import 'dart:ui';

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:mcl_gui/gui_state.dart';
import 'package:mcl_gui/proto/proto/mcl.pbenum.dart';
import 'package:mcl_gui/proto/proto/mclgui.pbgrpc.dart';
import 'package:mcl_gui/title.dart';

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

class QSOExchTextField extends TextField {
  QSOExchTextField(String title,
      {super.key, super.onSubmitted, super.focusNode, super.controller})
      : super(
          style: const TextStyle(fontFamily: "monospace"),
          decoration: InputDecoration(
            border: const OutlineInputBorder(),
            labelStyle: const TextStyle(fontFamily: "serif"),
            labelText: title,
          ),
          inputFormatters: <TextInputFormatter>[
            LengthLimitingTextInputFormatter(10),
            UpperCaseTextFormatter(),
          ],
        );
}

class _QSOInputState extends State<QSOInput> {
  final dxCallsignController = TextEditingController();
  var exchangeSentControllers = <TextEditingController>[];
  var exchangeRcvdControllers = <TextEditingController>[];
  var focuser = <FocusNode>[];

  @override
  void initState() {
    super.initState();
    exchangeSentControllers = <TextEditingController>[];
    focuser.add(FocusNode());
    for (int i = 0; i < state.exchangeSentFields.length; i++) {
      exchangeSentControllers.add(TextEditingController());
      focuser.add(FocusNode());
    }
    for (int i = 0; i < state.exchangeRcvdFields.length; i++) {
      exchangeRcvdControllers.add(TextEditingController());
      focuser.add(FocusNode());
    }
  }

  @override
  void dispose() {
    // Clean up the controller when the widget is disposed.
    dxCallsignController.dispose();
    for (var element in exchangeSentControllers) {
      element.dispose();
    }
    for (var element in exchangeRcvdControllers) {
      element.dispose();
    }
    exchangeSentControllers.clear();
    exchangeRcvdControllers.clear();
    super.dispose();
  }

  void onFieldSubmitted(String fieldType, int fieldSeq, String fieldValue) {
    var seq = 0;
    switch (fieldType) {
      case "dx_callsign":
        seq = fieldSeq;
        break;
      case "exch_sent":
        seq = 1 + fieldSeq;
        break;
      case "exch_rcvd":
        seq = 1 + exchangeSentControllers.length + fieldSeq;
        break;
    }
    seq++;
    if (seq == focuser.length) {
      submitQso();
      seq = 0;
    }
    focuser[seq].requestFocus();
  }

  void submitQso() {
    var dxCallsign = dxCallsignController.text;
    var exchangeSent = exchangeSentControllers.map((e) => e.text);
    var exchangeRcvd = exchangeRcvdControllers.map((e) => e.text);

    // TODO: Check response
    unawaited(state.getGuiClient()!.logQSO(QSOMessage(
          dxCallsign: dxCallsign,
          exchangeSent: exchangeSent,
          exchangeRcvd: exchangeRcvd,
          time: $fixnum.Int64(
              DateTime.now().toUtc().millisecondsSinceEpoch ~/ 1000),
          freq: $fixnum.Int64(14000000),
          mode: Mode.phone_lsb,
        )));

    dxCallsignController.text = "";
    for (var element in exchangeSentControllers) {
      element.text = "";
    }
    for (var element in exchangeRcvdControllers) {
      element.text = "";
    }
  }

  @override
  Widget build(BuildContext context) {
    var qsoInputFields = <Widget>[
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
        onSubmitted: (value) => onFieldSubmitted("dx_callsign", 0, value),
        focusNode: focuser[0],
        controller: dxCallsignController,
      )),
    ];
    for (var i = 0; i < state.exchangeSentFields.length; i++) {
      qsoInputFields.add(const SizedBox(width: 4, height: 1));
      qsoInputFields.add(Expanded(
          child: QSOExchTextField(
        state.exchangeSentFields[i],
        onSubmitted: (value) => onFieldSubmitted("exch_sent", i, value),
        focusNode: focuser[1 + i],
        controller: exchangeSentControllers[i],
      )));
    }
    for (var i = 0; i < state.exchangeRcvdFields.length; i++) {
      qsoInputFields.add(const SizedBox(width: 4, height: 1));
      qsoInputFields.add(Expanded(
          child: QSOExchTextField(
        state.exchangeRcvdFields[i],
        onSubmitted: (value) => onFieldSubmitted("exch_rcvd", i, value),
        focusNode: focuser[1 + state.exchangeSentFields.length + i],
        controller: exchangeRcvdControllers[i],
      )));
    }

    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      children: qsoInputFields,
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

class _GuiMainPageState extends State<GuiMainPage> {
  @override
  void initState() {
    super.initState();
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
      body: Column(children: <Widget>[
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
        Spacer(),
        GuiTitle(),
      ]),
    );
  }
}
