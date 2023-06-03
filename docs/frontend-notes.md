# Notes:
1. 先用 wagger 確認 api : spec.yaml
2. 進入每一頁前都先執行 isLoginedStatusCheck 確認登入狀態：
    (1) 在 App.vue 中， 生命週期的created()，把 isLoginStatus 傳入的boolean值，給 data 裡的 isLogined
    (2) NavbarComponent 
    ```
    watch: { // 只要有變化就執行
        isLoginedStatusCheck: {
            immediate: true, // watch's 立即執行
            // handler(newVal, oldVal) { // watch's 執行 handler
            //     console.log("isLoginedStatusCheck (NavBar): ", newVal, oldVal)
            // }
        },
    },
    ```
    (3) logout... 

- 在 api/index.js 寫跟api 有關的程式：確認 login status, 打 login api, 打 logout api
- App.vue       - provide: isLoginedStatusCheck, loginAction, logoutAction
- Navbar.vue    - inject:  isLoginedStatusCheck, logoutAction
- LoginView.vue - inject:  isLoginedStatusCheck, loginAction,
                - methods: localLoginAction
