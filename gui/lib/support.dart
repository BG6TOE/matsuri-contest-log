import 'package:mcl_gui/proto/proto/common.pbserver.dart';

String overrideQsoFieldTitle(String fieldName, Map<String, ContestFieldInfo> fieldInfo) {
  switch (fieldName) {
    case "rst_sent_":
      return "RST Tx";
    case "rst_rcvd_":
      return "RST Rx";
  }
  var field = fieldInfo[fieldName];
  if (field != null) {
    return field.displayName;
  }
  return fieldName;
}
