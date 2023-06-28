import { createRouter, createWebHistory } from 'vue-router';
import { useLoginStatusStore } from '../stores/useLoginStatusStore';
import CheckoutView from '../views/CheckoutView.vue';
import LoginView from '../views/LoginView.vue';
import PasswordView from '../views/PasswordView.vue';
import index from '../views/indexView.vue';
import notFound from '../views/notFound.vue';
import registerView from '../views/registerView.vue';
import productView from '../views/store/product/productView.vue';
import storeView from '../views/store/storeView.vue';



// export default defineComponent({
//     computed: {
//     }
// })

/**
 * 只允許「沒登入的 user」進入這頁面
 */
async function isGuest() {
    const store = useLoginStatusStore()
    if (store.isLogin) {
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
            beforeEnter: [isGuest],
            // beforeEnter: async(to,from) => {
            //     const store = useLoginStatusStore();
            //     const fromPath = from.path == '/login'
            //     if(store.isLogin == true || store.isLogin && fromPath) // 登入+從登入頁+去登入頁
            //     to.path.replace == '/'
            // }
        },
        {
            path: '/register',
            name: 'register',
            component: registerView,
            beforeEnter: [isGuest],
            // redirect: () => { 
            //     // const store = useLoginStatusStore()
            //     if (store.isLogin) {
            //         return {
            //             path: '/'
            //         }
            //     } 
            // }
        },
        {
            path: '/password',
            name: 'password',
            component: PasswordView
        },
        {
            path: '/store/:storeId',
            name: 'store/:storeId',
            component: storeView,
            // children: [
            //     {

            //     }
            // ],
            props: true,
        },
        {
            path: '/store/:storeId/product/:productId',
            name: 'product',
            component: productView,
        },
        {
            path: '/store/:storeId/checkout',
            name: 'checkout',
            component: CheckoutView,
            props: true,
            // props: (route) => route.params
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
router.beforeEach( async(to, from) => {
    const store = useLoginStatusStore();
    const fromPath = from.path !== '/login' || '/register'
    const toPath = to.path == '/login' || '/register'
    if(store.isLogin && fromPath || toPath) // 登入+從登入頁+去登入頁
    to.path.replace == '/'
})

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