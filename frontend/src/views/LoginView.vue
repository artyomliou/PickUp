<template>
    <section class="sec-h sec-login">
        <div class="container">
            <div class="row justify-content-center">
                <div class="col-12 col-lg-8 sec-titlebox">
                    <h2 class="sec-title">登入</h2>
                    <p class="sec-sub">歡迎回來！</p>
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
                        <button @click.prevent="login" type="submit" class="btn btn-md mx-auto btn-sumit">
                            送出
                        </button>
                    </form>
                    <div class="text-center">
                        <button @click="countPlus" class="btn btn-default">{{ click }}忘記密碼？</button>
                        <a href="" class="btn btn-inline">還沒註冊？</a>
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
import { mapActions } from 'pinia'
import { loginApi } from '../api'
import { useLoginStore } from '../store/login'
export default {
    data() {
        return {
            userLogin: {
                email: '',
                password: '',
            },
            countSet: 0
        }
    },
    methods: {
        ...mapActions(useLoginStore, ['updateLoginStatus']),
        async login() {
            const ok = await loginApi(this.userLogin.email, this.userLogin.password)
            if (ok) {
                this.updateLoginStatus(true)
                this.$router.push('/')
            }
        },
        countPlus() {
            console.count(this.countSet)
        }
    },
}
</script>
