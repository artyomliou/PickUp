# Notes:

- 在 api/index.js 寫跟api 有關的程式：確認 login status, 打 login api, 打 logout api
- App.vue       - provide: isLoginedStatusId, loginAction, logoutAction
- Navbar.vue    - inject:  isLoginedStatusId, logoutAction
- LoginView.vue - inject:  isLoginedStatusId, loginAction,
                - methods: localLoginAction
