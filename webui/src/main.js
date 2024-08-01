import { createApp } from 'vue'
import router from './router'
import axios from 'axios'
import App from './App.vue'

const app = createApp(App)
app.use(router)
app.config.globalProperties.$http = axios 
app.config.globalProperties.$apiurl = "http://localhost:3000"

app.mount('#app')