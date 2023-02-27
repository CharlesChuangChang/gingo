import { createApp } from 'vue'
import App from './App.vue'
import router from "./router";
import store from "./store";
import axios from "axios";
import ElementPlus  from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
const app = createApp(App)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.config.globalProperties.$axios = axios;
app.use(store).use(router).mount("#app")

