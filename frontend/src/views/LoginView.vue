<template>
    <Navbar />
    <section class="sec-h sec-login">
        <div class="container">
            <div class="row justify-content-center">
                <div class="col-12 col-lg-8 sec-titlebox">
                    <h2 class="sec-title">登入</h2>
                    <p class="sec-sub">歡迎回來！</p>
                </div>
                <div class="col-12 col-lg-5">
                    <form @submit.prevent="login" class="form">
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
                        <button @click="loginSubmit()" type="submit" class="btn btn-md mx-auto btn-sumit">
                            送出
                        </button>
                    </form>
                    <div class="text-center">
                        <a href="" class="btn btn-inline">忘記密碼？</a>
                        <a href="" class="btn btn-inline">還沒註冊？</a>
                    </div>
                </div>
            </div>
        </div>
    </section>
    <FooterComponent />

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
// bootstrap modal
// https://getbootstrap.com/docs/5.0/getting-started/webpack/#importing-javascript
import Modal from 'bootstrap/js/dist/modal'
import { mapActions, mapState } from 'pinia'
import { login } from '../api'
import { useLoginCheckStore } from '../stores/useLoginCheck.js'

export default {
    data() {
        return {
            userLogin: {
                email: '',
                password: ''
            },
            // alertTxt: ''
        }
    },
    computed: {
        ...mapState(useLoginCheckStore, ['isLogined'])
    },
    methods: {
        ...mapActions(useLoginCheckStore, ['checkLoginStatus']),
        async loginSubmit() {
            try {
                const response = await login(this.userLogin.email, this.userLogin.password);
                if (response.ok) {
                    this.showModal("你已成功登入");
                    this.$router.push('/')
                } else {
                    this.showModal("登入錯誤，請重新確認！")
                }
            } catch (error) {
                console.error(error)
            }
        },
        showModal(msg) {
            this.$refs.modalAlert.querySelector('h5').innerText = msg;

            const modalAlert = new Modal(this.$refs.modalAlert)
            modalAlert.show()
        }
    },
}
</script>
