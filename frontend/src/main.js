import { createPinia } from 'pinia'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'

import FooterComponent from './components/Footer.vue'
import Navbar from './components/Navbar.vue'

import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/css/main.css'

const app = createApp(App)
app.use(createPinia()).use(router)

app.component('Navbar', Navbar).component('FooterComponent', FooterComponent)

app.mount('#app')

// every pages
import { mapActions } from 'pinia'
import { useLoginCheckStore } from './stores/useLoginCheck.js'

export default {
    data() {
        return {}
    },
    computed: {},
    methods: {
        ...mapActions(useLoginCheckStore, ['fetchLoginApi'])
    },
    created() {
        this.fetchLoginApi()
    }
}
