import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import PasswordView from '../views/PasswordView.vue'
import index from '../views/index.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'index',
            component: index
        },
        {
            path: '/login',
            name: 'login',
            component: LoginView
        },
        {
            path: '/password',
            name: 'password',
            component: PasswordView
        }
    ]
})

export default router
