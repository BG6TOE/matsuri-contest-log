<template>
  <v-container>
    <table>
      <thead>
        <td>Callsign</td>
        <td>Freq</td>
        <td>Mode</td>
        <td>Time</td>
        <td>Exch Sent</td>
        <td>Exch Rcvd</td>
        <td>RST Sent</td>
        <td>RST Rcvd</td>
      </thead>
      <tbody>
        <tr v-for="qso in qsoList" :key="qso.uid">
          <td>{{ qso.dx_callsign }}</td>
          <td>{{ Helper.FormatFreq(qso.freq_hz) }}</td>
          <td>{{ qso.mode }}</td>
          <td>{{ Helper.FormatTime(qso.time) }}</td>
          <td>{{ qso.exch_sent }}</td>
          <td>{{ qso.exch_rcvd }}</td>
          <td>{{ qso.rst_sent }}</td>
          <td>{{ qso.rst_rcvd }}</td>
        </tr>
      </tbody>
    </table>
  </v-container>
</template>

<script>
import RPC from "../rpc.js";
import Event from "../event.js";
import Helper from "../helper.js";

export default {
  name: "LogList",

  data: () => ({
    qsoList: [],
    Helper,
  }),

  methods: {
    HandleNewQSO(qso) {
      this.qsoList.push(qso);
    },
  },

  created: function() {
    console.log(this);
    Event.$on("NewQSO", this.HandleNewQSO);
    RPC.getQSOList().then((res) => {
      this.qsoList = res.qsos;
    });
  },

  beforeDestroy() {
    Event.$off("NewQSO", this.HandleNewQSO);
  },
};
</script>
