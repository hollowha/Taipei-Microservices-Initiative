import { createRouter, createWebHistory } from "vue-router";
import HomePage from "@/views/HomePage.vue";
import ActivitiesList from "@/views/ActivitiesList.vue";
// import AutoFillFormtmp from "@/views/AutoFillFormtmp.vue";
import AutoFillFormSelect from "@/views/AutoFillFormSelect.vue";
import AutoFillFormFill from "@/views/AutoFillFormFill.vue";
import AutoFillFormDownload from "@/views/AutoFillFormDownload.vue";
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
  {
    path: '/auto-fill-form',
    name: 'SampleForm',
    component: AutoFillFormSelect
  },
  {
    path: "/auto-fill-form/:formType",
    name: "auto-fill-form",
    component: AutoFillFormFill,
  },
  // set route for`/autofillform/download?formname=${this.formType}&userid=5`
  {
    path: "/autofillform/download",
    name: "auto-fill-form-download",
    component: AutoFillFormDownload,
  },
  
];

// 創建路由器實例
const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
