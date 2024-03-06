import { createApp } from 'vue'
import './assets/css/global.css'
import App from './App.vue'
import router from "./router/router.js"

import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import i18n from './locales/index.js';

// import Mock from '../mock/mock'

const app = createApp(App)
app.use(router).use(ElementPlus).use(i18n).mount('#app')

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}