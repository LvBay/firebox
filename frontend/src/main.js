import {createApp} from 'vue'
import App from './App.vue'
import router from './router/index'
import naiveUi from "naive-ui"

createApp(App).use(naiveUi).use(router).mount('#app')
