<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">{{ state.contest.name }}</div>
      <h1>
        <code>{{ state.rig.mode }} {{ Helper.FormatFreq(state.rig.vfo) }}</code>
      </h1>
      <h1>
        <code>{{ currentInputtingCallsign }}</code>
      </h1>


      <v-spacer></v-spacer>
    </v-app-bar>

    <v-main>
      <LogList />
    </v-main>

    <v-footer padless>
      <v-col cols="12">MCL WebUI Version: {{ version.hash }}</v-col>
    </v-footer>
  </v-app>
</template>

<script>
import LogList from "./components/LogList";
import Event from "./event.js";
import Helper from "./helper.js";

export default {
  name: "App",

  components: {
    LogList,
  },

  data: () => ({
    Helper,
    version: JSON.parse(process.env.VUE_APP_GIT_VERSION),

    state: {
      contest: {
        uid: "00000000",
        name: "DX",
        start_time: "1970-01-01T08:00:00+08:00",
        end_time: "1970-01-01T08:00:00+08:00",
        config: "",
        station_callsign: "XX1XXX",
      },
      operator: { callsign: "YY1YYY" },
      rig: { vfo: 0, mode: "CW" },
      rigConfig: [
        {
          Manufacturer: "",
          model: "",
        },
      ],
    },
    currentInputtingCallsign: "",
  }),

  methods: {
    HandleStateReport(state) {
      console.log(state);
      this.state = JSON.parse(JSON.stringify(state));
    },
    HandleInputtedCallsign(callsign) {
      this.currentInputtingCallsign = callsign
    }
  },

  created: function () {
    Event.$on("StateReport", this.HandleStateReport);
    Event.$on("InputCallsignChanged", this.HandleInputtedCallsign);
  },

  beforeDestroy: function () {
    Event.$off("StateReport", this.HandleStateReport);
    Event.$off("InputCallsignChanged", this.HandleInputtedCallsign);
  },
};
</script>
