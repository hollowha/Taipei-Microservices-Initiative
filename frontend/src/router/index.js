import { createRouter, createWebHistory } from "vue-router";
import HomePage from "@/views/HomePage.vue";
import ActivitiesList from "@/views/ActivitiesList.vue";

// 定義路由配置
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
];

// 創建路由器實例
const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
