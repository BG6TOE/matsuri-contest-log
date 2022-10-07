String overrideQsoFieldTitle(String fieldName) {
  switch (fieldName) {
    case "_rst":
      return "RST";
  }
  return fieldName;
}
