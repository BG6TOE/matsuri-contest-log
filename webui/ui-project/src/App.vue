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

    <v-snackbar v-model="snackbar" timeout="2000">
      {{ snackbarText }}
    </v-snackbar>

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
    snackbarText: "",
    snackbar: false,
    websocketConnected: false,
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
      this.currentInputtingCallsign = callsign;
    },
    localEventHandler(event) {
      switch (event.name) {
        case "WS.OnOpen":
          this.websocketConnected = true;
          this.snackbarText = "Connected to MCL program.";
          this.snackbar = true;
          break;
        case "WS.OnClose":
          if (this.websocketConnected != false) {
            this.snackbarText = "Disconnected from MCL program.";
            this.snackbar = true;
          }
          this.websocketConnected = false;
          break;
        default:
      }
    },
  },

  created: function () {
    Event.$on("StateReport", this.HandleStateReport);
    Event.$on("InputCallsignChanged", this.HandleInputtedCallsign);
    Event.$on("LocalEvent", this.localEventHandler);
  },

  beforeDestroy: function () {
    Event.$off("StateReport", this.HandleStateReport);
    Event.$off("InputCallsignChanged", this.HandleInputtedCallsign);
    Event.$off("LocalEvent", this.localEventHandler);
  },
};
</script>
