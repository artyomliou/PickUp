// import { createPinia } from 'pinia'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'


import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/css/main.css'

// import { isLoginStatus } from './api/index.js'
const app = createApp(App)
// app.provide("isLoginPro", function(){
//     isLoginStatus() // provide key, value
// })
app.use(router)

// app.component('NavbarComponent', NavbarComponent).component('FooterComponent', FooterComponent)

app.mount('#app')
// app.config.unwrapInjectedRef = true
