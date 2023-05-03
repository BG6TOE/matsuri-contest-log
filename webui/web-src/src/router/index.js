import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import NewContest from '../views/NewContest.vue'
import RunContest from '../views/RunContest.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/contest/new',
      name: 'new-contest',
      component: NewContest,
    },
    {
      path: '/contest/open',
      name: 'open-contest',
      component: NewContest,
    },
    {
      path: '/contest/run',
      name: 'run-contest',
      component: RunContest,
    }
  ]
})

export default router
