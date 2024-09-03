import { createRouter, createWebHistory } from "vue-router";
import HomePage from "@/views/HomePage.vue";
import ActivitiesList from "@/views/ActivitiesList.vue";
import NoiceDetect from "@/views/NoiceDetect.vue";
import VotePage from "@/views/VotePage.vue";  // 確保這行存在

const routes = [
  {
    path: "/",
    name: "Home",
    component: HomePage,
  },
  {
    path: "/activities",
    name: "activities",
    component: ActivitiesList,
  },
  {
    path: "/detect",
    name: "noice-detect",
    component: NoiceDetect,
  },
  {
    path: "/vote",  // 確保這行存在
    name: "vote",
    component: VotePage,
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
