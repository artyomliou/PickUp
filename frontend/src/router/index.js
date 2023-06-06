import { createRouter, createWebHistory } from 'vue-router';
import { useLoginStatusStore } from '../stores/useLoginStatusStore';
import LoginView from '../views/LoginView.vue';
import PasswordView from '../views/PasswordView.vue';
import index from '../views/indexView.vue';
import notFound from '../views/notFound.vue';
import registerView from '../views/registerView.vue';


// export default defineComponent({
//     computed: {
//     }
// })

/**
 * 只允許「沒登入的 user」進入這頁面
 */
async function isGuest() {
    const store = useLoginStatusStore()
    if (await store.isLoginStatus()) {
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
            alias: ['/home, /index'],
            component: index
        },
        {
            path: '/login',
            name: 'login',
            component: LoginView,
            beforeEnter: [isGuest]
        },
        {
            path: '/register',
            name: 'register',
            component: registerView,
            beforeEnter: [isGuest]
        },
        {
            path: '/password',
            name: 'password',
            component: PasswordView
        },
        { 
            // form Kuro https://book.vue.tw/CH4/4-2-route-settings.html
            path: '/:pathMatch(.*)*',
            name: 'notFound',
            component: notFound
        },
    ]
});


// 每一個 router 都會經過
// router.beforeEach((to) => {
//     const store = useLoginCheckStore()
//     if(to.meta.requiresAuth && !store.isLogined)
//     return '/login'
// })

// router.beforeEach( async(to) => {
//     //如果已登入(isLogined)就轉到首頁
//     const store = useLoginCheckStore()
//     if(
//         store.isLogined && //檢查登入
//         to.name !== 'login'
//     ) {
//         return { name: 'index'}
//     }
// });
// // note: meta https://router.vuejs.org/zh/guide/advanced/meta.html


export default router