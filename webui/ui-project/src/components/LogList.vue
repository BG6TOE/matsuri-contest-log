<template>
  <v-container>
    <v-data-table
      dense
      :headers="qsoTableHeaders"
      :items="qsoList"
      :items-per-page="25"
      @dblclick:row="qsoTableDblClick"
    >
      <template v-slot:item.freq_hz="{ item }">
        {{ Helper.FormatFreq(item.freq_hz) }}
      </template>
      <template v-slot:item.time="{ item }">
        {{ Helper.FormatTime(item.time) }}
      </template>
    </v-data-table>

    <v-dialog v-model="dialog" fullscreen>
      <v-card>
        <v-card-title>
          <span class="text-h5"
            >Update Log <code>{{ editedItem.uid }}</code></span
          >
        </v-card-title>

        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12" sm="3" md="4">
                <v-text-field
                  v-model="editedItem.dx_callsign"
                  label="Callsign"
                ></v-text-field>
              </v-col>

              <v-col cols="12" sm="3" md="4">
                <vc-date-picker v-model="editedItem.time" mode="dateTime" is24hr timezone="UTC">
                  <template v-slot="{ inputValue, inputEvents }">
                    <v-text-field
                     :value="Helper.FormatTime(editedItem.time)"
                      v-on="inputEvents"
                    />
                  </template>
                </vc-date-picker>
              </v-col>
              <v-col cols="12" sm="3" md="4">
                <v-text-field
                  v-model="editedItem.freq_hz"
                  label="Freq"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="3" md="4">
                <v-text-field
                  v-model="editedItem.mode"
                  label="Mode"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="3" md="4">
                <v-text-field
                  v-model="editedItem.rst_sent"
                  label="RST Sent"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="3" md="4">
                <v-text-field
                  v-model="editedItem.exch_sent"
                  label="Exch Sent"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="3" md="4">
                <v-text-field
                  v-model="editedItem.rst_rcvd"
                  label="RST Rcvd"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="3" md="4">
                <v-text-field
                  v-model="editedItem.exch_rcvd"
                  label="Exch Rcvd"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="close"> Cancel </v-btn>
          <v-btn color="blue darken-1" text @click="save"> Save </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
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
    qsoTableHeaders: [
      { text: "Callsign", value: "dx_callsign" },
      { text: "Freq", value: "freq_hz" },
      { text: "Mode", value: "mode" },
      { text: "Time", value: "time" },
      { text: "RST Sent", value: "rst_sent" },
      { text: "Exch Sent", value: "exch_sent" },
      { text: "RST Rcvd", value: "rst_rcvd" },
      { text: "Exch Rcvd", value: "exch_rcvd" },
    ],
    editedItem: {},
    dialog: false,
    Helper,
  }),

  methods: {
    HandleNewQSO(qso) {
      this.qsoList.push(qso);
    },
    qsoTableDblClick(event, item) {
      console.log(event);
      console.log(item);
      this.editedItem = JSON.parse(JSON.stringify(item.item));
      this.editedItem.timeVar = this.Helper.FormatTime(this.editedItem.time);
      this.dialog = true;
    },
    save() {
      let updateQSO = JSON.parse(JSON.stringify(this.editedItem));
      updateQSO.time = (new Date(this.editedItem.time)).toISOString();
      console.log(updateQSO.time);
      updateQSO.freq_hz = Number(updateQSO.freq_hz);
      RPC.updateQSO({qso: updateQSO});
      this.qsoList.forEach((e, index) => {
        if (e.uid == updateQSO.uid) {
          this.qsoList.splice(index, 1, updateQSO)
        }
      })
      this.dialog = false;
    },
    close() {
      this.dialog = false;
    },
  },

  created: function () {
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
