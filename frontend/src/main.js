import { createPinia } from 'pinia'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'

import FooterComponent from './components/Footer.vue'
import NavbarComponent from './components/NavbarComponent.vue'

import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/css/main.css'

const app = createApp(App)
app.use(createPinia()).use(router)

app.component('NavbarComponent', NavbarComponent).component('FooterComponent', FooterComponent)

app.mount('#app')
