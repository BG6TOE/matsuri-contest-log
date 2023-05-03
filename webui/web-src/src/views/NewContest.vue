<script setup>
import axios from 'axios'
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const contest_def = ref('')
const station_callsign = ref('')
const database_file = ref('')
const custom_fields = ref([])
let contest_metadata = {};
let contest_script = '';

const router = useRouter();

async function parseContestDefinition() {
  let res = await axios({
    method: 'post',
    url: '/call/Gui/ParseContest',
    data: {
      contest_descriptor: contest_def.value,
    }
  })
  let data = res.data.data;
  let fields = [];
  contest_script = contest_def.value;
  contest_metadata = data;
  for (let field in data.FieldInfos) {
    let val = data.FieldInfos[field];
    if (val.FieldType == 'tx_const') {
      fields.push({
        key: field,
        value: "",
        title: val.DisplayName,
        description: val.Description,
      })
    }
  }

  custom_fields.value = fields;
}

async function createContest() {
  let custom_field_values = {};

  for (let field in custom_fields.value) {
    let val = custom_fields.value[field];
    custom_field_values[val.key] = val.value;
  }

  await axios({
    method: 'post',
    url: '/call/Gui/CreateContest',
    data: {
      database_name: database_file.value,
      contest: {
        name: '',
        contest: contest_metadata,
        station: {
          callsign: station_callsign.value.toUpperCase(),
          custom_fields: custom_field_values,
        },
        begin_timestamp: 0,
        end_timestamp: 0,
        contest_script: contest_script,
      },
    }
  })

  await axios({
    method: 'post',
    url: '/call/Gui/LoadContest',
    data: { 'database_name': database_file.value }
  })
  router.replace('/contest/run')
}
</script>

<template>
  <div>
    <div>Contest Definition: <textarea v-model="contest_def"></textarea><button
        @click="parseContestDefinition()">Load</button></div>
    <div>Station Callsign: <input v-model="station_callsign" type="text" /></div>
    <div v-for="(item, _) in custom_fields">
      {{ item.title }}: <input v-model="item.value" /> {{ item.description }}
    </div>
    <div>Database File: <input v-model="database_file" type="text" /></div>
    <button @click="createContest()">Save</button>
  </div>
</template>

<style></style>