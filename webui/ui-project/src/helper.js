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
        return (freq / 1000).toFixed(3) + " k";
    },
}

export default Helper