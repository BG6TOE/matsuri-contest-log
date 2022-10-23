import 'dart:async';

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'gui_state.dart';
import 'last_qso_table.dart';
import 'proto/proto/mcl.pbenum.dart';
import 'proto/proto/mcl.pb.dart';
import 'proto/proto/mclgui.pb.dart';
import 'theme.dart';

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
          style: FontConfigs.monospaceFont,
          decoration: InputDecoration(
            border: const OutlineInputBorder(),
            labelStyle: FontConfigs.sansFont,
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

  var _qsoFieldSetUpdateCallbackId = 0;

  void _onQsoFieldSetUpdate() {
    exchangeSentControllers = <TextEditingController>[];
    focuser = <FocusNode>[];
    focuser.add(FocusNode());
    for (int i = 0; i < state.displayExchangeSentFields.length; i++) {
      exchangeSentControllers.add(TextEditingController());
      focuser.add(FocusNode());
    }
    for (int i = 0; i < state.displayExchangeRcvdFields.length; i++) {
      exchangeRcvdControllers.add(TextEditingController());
      focuser.add(FocusNode());
    }

    setState(() {});
  }

  @override
  void initState() {
    super.initState();

    _onQsoFieldSetUpdate();

    _qsoFieldSetUpdateCallbackId = state.registerCallback(
        CallbackKind.kOnQsoFieldSetUpdate, _onQsoFieldSetUpdate);
  }

  @override
  void dispose() {
    state.deregisterCallback(
        CallbackKind.kOnQsoFieldSetUpdate, _qsoFieldSetUpdateCallbackId);

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
    var exchangeSent = <String, String>{};
    var exchangeRcvd = <String, String>{};

    for (var i = 0; i < state.displayExchangeSentFields.length; i++) {
      exchangeSent[state.displayExchangeSentFields[i].value]=exchangeSentControllers[i].text;
    }
    for (var i = 0; i < state.displayExchangeRcvdFields.length; i++) {
      exchangeRcvd[state.displayExchangeRcvdFields[i].value] = exchangeRcvdControllers[i].text;
    }

    // TODO: Check response
    unawaited(() async {
      await state.getGuiClient()!.logQSO(QSO(
            dxCallsign: dxCallsign,
            exchangeSent: exchangeSent,
            exchangeRcvd: exchangeRcvd,
            time: $fixnum.Int64(
                DateTime.now().toUtc().millisecondsSinceEpoch ~/ 1000),
            freq: $fixnum.Int64(14000000),
            mode: "LSB",
          ));
      state.refreshQsos();
    }());

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
    for (var i = 0; i < state.displayExchangeSentFields.length; i++) {
      qsoInputFields.add(const SizedBox(width: 4, height: 1));
      qsoInputFields.add(Expanded(
          child: QSOExchTextField(
        state.displayExchangeSentFields[i].title,
        onSubmitted: (value) => onFieldSubmitted("exch_sent", i, value),
        focusNode: focuser[1 + i],
        controller: exchangeSentControllers[i],
      )));
    }
    for (var i = 0; i < state.displayExchangeRcvdFields.length; i++) {
      qsoInputFields.add(const SizedBox(width: 4, height: 1));
      qsoInputFields.add(Expanded(
          child: QSOExchTextField(
        state.displayExchangeRcvdFields[i].title,
        onSubmitted: (value) => onFieldSubmitted("exch_rcvd", i, value),
        focusNode: focuser[1 + state.displayExchangeSentFields.length + i],
        controller: exchangeRcvdControllers[i],
      )));
    }

    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      children: qsoInputFields,
    );
  }
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

class HotkeyPanel extends StatefulWidget {
  const HotkeyPanel({super.key});

  @override
  State<HotkeyPanel> createState() => _HotkeyPanelState();
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

class QsoPanel extends StatefulWidget {
  @override
  State<QsoPanel> createState() => _QsoPanelState();
}

class _QsoPanelState extends State<QsoPanel> {
  @override
  Widget build(BuildContext context) {
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
          SizedBox(height: 220, child: LastQsoTable()),
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
        ]);
  }
}
