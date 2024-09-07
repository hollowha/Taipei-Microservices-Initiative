import { createRouter, createWebHistory } from "vue-router";
import HomePage from "@/views/HomePage.vue";
import ActivitiesList from "@/views/ActivitiesList.vue";
import NoiceDetect from "@/views/NoiceDetect.vue";
import ReportSubmit from "@/views/ReportSubmit.vue";
import ReportManage from "@/views/ReportManage.vue";
import DataDetect from "@/views/DataDetect.vue";
import AutoFillForm from "@/views/AutoFillForm.vue";
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
    path: "/auto-fill-form",
    name: "auto-fill-form",
    component: AutoFillForm,
  },
  {
    path: "/detect",
    name: "noice-detect",
    component: NoiceDetect,
  },
  {
    path: "/report",
    name: "report",
    component: ReportSubmit,
  },
  {
    path: "/report-manage",
    name: "report-manage",
    component: ReportManage,
  }

  {
    path: "/data-detect",
    name: "data-detect",
    component: DataDetect,
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
