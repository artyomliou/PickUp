import { createRouter, createWebHistory } from 'vue-router'
import { isLoggedIn } from '../api'
import LoginView from '../views/LoginView.vue'
import PasswordView from '../views/PasswordView.vue'
import index from '../views/index.vue'

/**
 * 只允許「沒登入的 user」進入這頁面
 */
async function isGuest() {
    if (await isLoggedIn()) {
        return {
            path: '/'
        }
    }
    return true
}

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
            component: LoginView,
            beforeEnter: [isGuest]
        },
        {
            path: '/password',
            name: 'password',
            component: PasswordView
        }
    ]
})

export default router
