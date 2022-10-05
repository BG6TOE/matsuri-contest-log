import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

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

    _updateQsoList();
  }

  @override
  void dispose() {
    state.deregisterCallback(CallbackKind.kOnQsoUpdate, callbackId);

    super.dispose();
  }

  void _updateQsoList() {
    setState(() {
      titles = ['Call', 'Date', 'Time', 'Freq'];
      titles.addAll(state.exchangeSentFields);
      titles.addAll(state.exchangeRcvdFields);
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

        for (var j = 0;
            j < lastQsos[i].exchSent.length &&
                j < state.exchangeSentFields.length;
            j++) {
          qsoStr.add(lastQsos[i].exchSent[j]);
        }
        for (var j = 0;
            j < lastQsos[i].exchRcvd.length &&
                j < state.exchangeRcvdFields.length;
            j++) {
          qsoStr.add(lastQsos[i].exchRcvd[j]);
        }

        qsos.add(qsoStr);
      }

      print(qsos);
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
