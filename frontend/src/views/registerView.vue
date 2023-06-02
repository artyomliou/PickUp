<template>
    <section class="sec-h sec-login">
        <div class="container">
            <div class="row justify-content-center">
                <div class="col-12 col-lg-8 sec-titlebox">
                    <h2 class="sec-title">註冊</h2>
                    <!-- <p class="sec-sub">立即註冊、收信</p> -->
                </div>
                <div class="col-12 col-lg-5">
                    <form class="form">
                        <div class="form-group">
                            <label class="label">
                                <span class="input-head">帳號 (E-mail)</span>
                                <input type="email" class="form-control" v-model="userLogin.email" required />
                            </label>
                        </div>
                        <div class="form-group">
                            <label class="label">
                                <span class="input-head">密碼</span>
                                <input type="password" class="form-control" v-model="userLogin.password" requored />
                            </label>
                        </div>
                        <button @click.prevent="localLoginAction" type="submit" class="btn btn-md mx-auto btn-sumit">
                            送出
                        </button>
                    </form>
                    <div class="text-center">
                        <!-- <button @click="countPlus" class="btn btn-default">{{ click }}忘記密碼？</button> -->
                        <a href="/login" class="btn btn-inline">已註冊？</a>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!--  -->
    <div ref="modalAlert" class="modal fade" tabindex="-1">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <!-- <h5 class="modal-title">{{ alertTxt }}</h5> -->
                    <h5 class="modal-title"></h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
// import { reactive } from 'vue';
// import { loginAction } from '../js/api/index';
// bootstrap modal
// https://getbootstrap.com/docs/5.0/getting-started/webpack/#importing-javascript
export default {
    inject: ['isLoginedStatusId', 'loginAction'],
    // inject: reactive([isLoginedStatusId)],
    data() {
        return {
            userLogin: {
                email: '',
                password: '',
            },
            isLoginStatus: false,
            countSet: 0
        }
    },
    methods: {
        async localLoginAction() {
            this.loginAction(this.userLogin.email, this.userLogin.password);
            // const response = await this.loginAction(this.userLogin.email, this.userLogin.password);
            // if (response.ok) {
            //     // await nextTick();
            //     this.isLoginStatus = await isLoginStatus();
            //     // new Modal().show(this.$refs.modalAlert("你已成功登入"))
            //     await this.$router.push('/');
            //     return response.blob(); // TODO change
            // }
            // new Error('Network response was not okk');
            // Modal(this.$refs.modalAlert("登入錯誤，請重新確認！")).show()
            // new Modal().show(this.$refs.modalAlert("登入錯誤，請重新確認！"))
            // return console.log('Cannot to be ok');
        },
        countPlus() {
            console.count(this.countSet)
        }
    },
    // mounted() {
    //     this.localLoginAction();
    // }
}
</script>
