// import { createPinia } from 'pinia'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'


import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/css/main.css'

const app = createApp(App)
app.use(router)

// app.component('NavbarComponent', NavbarComponent).component('FooterComponent', FooterComponent)

app.mount('#app')