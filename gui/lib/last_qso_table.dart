import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:mcl_gui/support.dart';
import 'package:mcl_gui/theme.dart';

import 'gui_state.dart';

class LastQsoTable extends StatefulWidget {
  @override
  State<LastQsoTable> createState() => _LastQsoTableState();
}

class _LastQsoTableState extends State<LastQsoTable> {
  List<List<String>> qsos = [];
  List<String> titles = ['', 'Call', 'Date', 'Time', 'Freq'];

  int callbackId = 0;

  @override
  void initState() {
    super.initState();
    callbackId =
        state.registerCallback(CallbackKind.kOnQsoUpdate, _updateQsoList);
    
    state.refreshQsos();
  }

  @override
  void dispose() {
    state.deregisterCallback(CallbackKind.kOnQsoUpdate, callbackId);

    super.dispose();
  }

  void _updateQsoList() {
    setState(() {
      titles = ['', 'Call', 'Date', 'Time', 'Freq'];
      titles.addAll(state.activeContest!.contest.exchangeSent.map((e) => overrideQsoFieldTitle(e, state.activeContest!.contest.fieldInfos)));
      titles.addAll(state.activeContest!.contest.exchangeRcvd.map((e) => overrideQsoFieldTitle(e, state.activeContest!.contest.fieldInfos)));
      var lastQsos = state.activeQsos.sublist(
          max(state.activeQsos.length - 10, 0), state.activeQsos.length);

      qsos.clear();
      for (var i = 0; i < lastQsos.length; i++) {
        var qsoTime =
            DateTime.fromMillisecondsSinceEpoch(lastQsos[i].time.toInt() * 1000)
                .toUtc();
        List<String> qsoStr = [
          'Alt-${(lastQsos.length - i) % 10}',
          lastQsos[i].dxCallsign,
          '${qsoTime.month.toString().padLeft(2, "0")}-${qsoTime.day.toString().padLeft(2, "0")}',
          '${qsoTime.hour.toString().padLeft(2, "0")}:${qsoTime.minute.toString().padLeft(2, "0")}',
          '${lastQsos[i].freq.toInt() ~/ 1000}',
        ];

        for (var exchSentField in state.activeContest!.contest.exchangeSent) {
          var exchSentVal = lastQsos[i].exchangeSent[exchSentField];
          exchSentVal ??= "";
          qsoStr.add(exchSentVal);
        }
        
        for (var exchRcvdField in state.activeContest!.contest.exchangeRcvd) {
          var exchRcvdVal = lastQsos[i].exchangeRcvd[exchRcvdField];
          exchRcvdVal ??= "";
          qsoStr.add(exchRcvdVal);
        }

        qsos.add(qsoStr);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return DataTable(
      columnSpacing: 32,
      dataRowHeight: 20,
      headingRowHeight: 20,
      columns: titles
          .map((e) => DataColumn(label: Expanded(child: Text(e, style: FontConfigs.monospaceFont,))))
          .toList(),
      rows: qsos
          .map((e) => DataRow(cells: e.map((c) => DataCell(Text(c, style: FontConfigs.monospaceFont))).toList()))
          .toList(),
    );
  }
}
