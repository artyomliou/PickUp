import { createPinia } from 'pinia'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'

import FooterComponent from './components/Footer.vue'
import NavBar from './components/NavBar.vue'

import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/css/main.css'

const app = createApp(App)
app.use(createPinia()).use(router)

app.component('NavBar', NavBar).component('FooterComponent', FooterComponent)

app.mount('#app')
