import { createPinia } from 'pinia'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'


import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/css/main.css'

const pinia = createPinia()
const app = createApp(App)
app.use(router).use(pinia)

// app.component('NavbarComponent', NavbarComponent).component('FooterComponent', FooterComponent)

app.mount('#app')