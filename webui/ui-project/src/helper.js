const Helper = {
    FormatTime: function (time) {
        let newDate = new Date(time)
        let month = newDate.getUTCMonth() + 1
        let day = newDate.getUTCDate()
        let hour = newDate.getUTCHours()
        let minute = newDate.getUTCMinutes()
        return (month < 10 ? "0" + month : month) + "-" +
            (day < 10 ? "0" + day : day) + " " +
            (hour < 10 ? "0" + hour : hour) + ":" +
            (minute < 10 ? "0" + minute : minute);
    },

    FormatFreq: function (freq) {
      // let unit = ["Hz", "kHz", "MHz", "GHz"];
      let ret = "Hz";
      let res = [];
      if (freq < 1) {
          return ""
      }
      while (freq > 1) {
        res = [freq % 1000, ...res];
        freq = Math.round(freq / 1000);
      }
      ret = "" + res[0];
      for (let i = 1; i < res.length; i++) {
        let t = "" + res[i];
        while (t.length < 3) {
          t = "0" + t;
        }
        ret = ret + " " + t;
      }
      return ret + " Hz";
    },
}

export default Helper