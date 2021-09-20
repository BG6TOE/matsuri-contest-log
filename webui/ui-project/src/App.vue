<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">MCL</div>

      <v-spacer></v-spacer>
      <code>{{ version }}</code>
    </v-app-bar>

    <v-main>
      <LogList />
    </v-main>

    <v-footer padless>
      <v-col cols="12"> MCL WebUI Version: {{ version.raw }}</v-col>
    </v-footer>
  </v-app>
</template>

<script>
import LogList from "./components/LogList";

export default {
  name: "App",

  components: {
    LogList,
  },

  data: () => ({
    version: process.env.VUE_APP_GIT_VERSION,
  }),
};
</script>
