import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:mcl_gui/support.dart';

import 'gui_state.dart';

class LastQsoTable extends StatefulWidget {
  @override
  State<LastQsoTable> createState() => _LastQsoTableState();
}

class _LastQsoTableState extends State<LastQsoTable> {
  List<List<String>> qsos = [];
  List<String> titles = ['Call', 'Date', 'Time', 'Freq'];

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
      titles = ['Call', 'Date', 'Time', 'Freq'];
      titles.addAll(state.activeContest!.contest.exchSent.map((e) => '${overrideQsoFieldTitle(e)} Sent'));
      titles.addAll(state.activeContest!.contest.exchRcvd.map((e) => '${overrideQsoFieldTitle(e)} Rcvd'));
      var lastQsos = state.activeQsos.sublist(
          max(state.activeQsos.length - 10, 0), state.activeQsos.length);

      qsos.clear();
      for (var i = 0; i < lastQsos.length; i++) {
        var qsoTime =
            DateTime.fromMillisecondsSinceEpoch(lastQsos[i].time.toInt() * 1000)
                .toUtc();
        List<String> qsoStr = [
          lastQsos[i].dxCallsign,
          '${qsoTime.month.toString().padLeft(2, "0")}-${qsoTime.day.toString().padLeft(2, "0")}',
          '${qsoTime.hour.toString().padLeft(2, "0")}:${qsoTime.minute.toString().padLeft(2, "0")}',
          '${lastQsos[i].freq.toInt() ~/ 1000}',
        ];

        for (var exchSentField in state.activeContest!.contest.exchSent) {
          var exchSentVal = lastQsos[i].exchSent[exchSentField];
          exchSentVal ??= "";
          qsoStr.add(exchSentVal!);
        }
        
        for (var exchRcvdField in state.activeContest!.contest.exchRcvd) {
          var exchRcvdVal = lastQsos[i].exchRcvd[exchRcvdField];
          exchRcvdVal ??= "";
          qsoStr.add(exchRcvdVal!);
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
          .map((e) => DataColumn(label: Expanded(child: Text(e))))
          .toList(),
      rows: qsos
          .map((e) => DataRow(cells: e.map((e) => DataCell(Text(e))).toList()))
          .toList(),
    );
  }
}
