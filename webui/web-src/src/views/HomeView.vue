<script setup>
import axios from 'axios'
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter();

async function created() {
  try {
    let res = await axios({
      method: 'post',
      url: '/call/Gui/GetActiveContest',
      data: {}
    })
    router.replace('/contest/run')
  } catch {

  }
}

onMounted(created);
</script>

<template>
  <div>
    <RouterLink to="/contest/new">Create</RouterLink> a new contest or <RouterLink to="/contest/open">open</RouterLink> an
    existing contest.
  </div>
</template>

<style></style>