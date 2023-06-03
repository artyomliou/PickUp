<template>
    <div style="background #aca; text-align:center;"> isLoginedStatusCheck:{{ isLoginedStatusCheck }} |
        isLoginedStatusCheck._value:{{ isLoginedStatusCheck._value }} </div>

    <div>
        <div class="shadow" :class="{ show: sidebarShow }" @click="sidebarToggle"></div>
    </div>

    <nav class="navbar navbar-expand-lg navbar-light main-nav" id="navMain">
        <div class="container-fluid">
            <button @click="sidebarToggle" class="navbar-hamburger" type="button" role="button" data-bs-toggle="collapse"
                data-bs-target="#navbarMainToggle" aria-controls="navbarMainToggle" aria-expanded="false"
                aria-label="Toggle navigation">
                <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="bi bi-list" viewBox="0 0 16 16">
                    <path fill-rule="evenodd"
                        d="M2.5 12a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5zm0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5zm0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5z" />
                </svg>
            </button>
            <a href="/" class="navbar-brand brand-link">
                <img src="@/assets/img/logo.svg" alt="" class="brand-img" />
            </a>

            <div class="btns row justify-content-between">

                <span class="col-auto p-0 btns-choose">
                    <button type="button" class="btn btn-takeout">外帶</button>
                </span>
                <button type="button" class="btn btn-location">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="bi bi-geo-alt-fill"
                        viewBox="0 0 16 16">
                        <path d="M8 16s6-5.686 6-10A6 6 0 0 0 2 6c0 4.314 6 10 6 10zm0-7a3 3 0 1 1 0-6 3 3 0 0 1 0 6z" />
                    </svg>
                    <span class="location-txt">目前地址</span>
                </button>
                <template v-if="!isLoginedStatusCheck_value">
                    <a href="/login" class="btn btn-login">登入</a>
                    <a href="/register" class="btn btn-register">註冊</a>
                </template>
                <template v-else>
                    <button type="button" class="btn btn-cart ms-auto">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="bi bi-cart4" viewBox="0 0 16 16">
                            <path
                                d="M0 2.5A.5.5 0 0 1 .5 2H2a.5.5 0 0 1 .485.379L2.89 4H14.5a.5.5 0 0 1 .485.621l-1.5 6A.5.5 0 0 1 13 11H4a.5.5 0 0 1-.485-.379L1.61 3H.5a.5.5 0 0 1-.5-.5zM3.14 5l.5 2H5V5H3.14zM6 5v2h2V5H6zm3 0v2h2V5H9zm3 0v2h1.36l.5-2H12zm1.11 3H12v2h.61l.5-2zM11 8H9v2h2V8zM8 8H6v2h2V8zM5 8H3.89l.5 2H5V8zm0 5a1 1 0 1 0 0 2 1 1 0 0 0 0-2zm-2 1a2 2 0 1 1 4 0 2 2 0 0 1-4 0zm9-1a1 1 0 1 0 0 2 1 1 0 0 0 0-2zm-2 1a2 2 0 1 1 4 0 2 2 0 0 1-4 0z" />
                        </svg>
                        <span class="cart-txt">購物車</span> •
                        <span class="cart-num">0</span>
                    </button>
                </template>

            </div>

        </div>
    </nav>

    <!-- Sidebar -->
    <!-- <div class="collapse sidebar" id="navbarMainToggle"> -->
    <div class="collapse sidebar" :class="{ show: sidebarShow }">
        <template v-if="!isLoginedStatusCheck_value">
            <!-- 登出中 -->
            <div class="nav-logouting mt-5" id="navLogout">
                <router-link to="/login" class="btn btn-light btn-login">登入</router-link>
                <!-- <ul class="nav-item">
                      <li class="nav-item">
                          <a href="" class="nav-link">XXX</a>
                      </li>
                  </ul> -->
            </div>
        </template>
        <template v-else>
            <!-- 登入中 -->
            <div class="nav-logining" id="navLogin">
                <div class="btn-avatar">
                    <img src="@/assets/img/avatar.jpg" alt="avatar" class="avatar-img" />
                </div>
                <ul class="nav-item nav-logining-list">
                    <li class="ropdown-item nav-item">Mue hung</li>

                    <button @click="logoutAction" type="button" role="button" class="btn btn-light btn-logout">
                        登出
                    </button>
                </ul>
            </div>
        </template>
    </div>
</template>

<script>
// import { inject } from 'vue';
// import { logoutAction } from '../api';
// import { isLoginStatus } from '../js/api/index.js';
export default {
    data() {
        return {
            isLoginStatus: 'isLoginStatus had error?',
            sidebarShow: false, //側邊欄隱藏
        }
    },
    inject: ['isLoginedStatusCheck', 'logoutAction'],
    watch: {
        isLoginedStatusCheck: {
            immediate: true,
            handler(newStatus, oldStatus) {
                if (newStatus == true) {
                    console.log('status true')
                } else {
                    console.log('status false')
                }
                // console.log("isLoginedStatusCheck (NavBar): ", newVal, oldVal)
            }
        },
        logoutAction: {
            isLoginedStatusCheck() { }
        }
    },
    computed: {},
    methods: {
        sidebarToggle() {
            this.sidebarShow = !this.sidebarShow;
        },

    }
}
</script>
