<template>
    <nav class="navbar navbar-expand-lg navbar-light main-nav" id="navMain">
        <div class="container-fluid">
            <a href="/" class="navbar-brand brand-link">
                <img src="@/assets/img/logo.png" alt="" class="brand-img" />
            </a>
            <div class="input-search mr-auto">
                <label class="input-group input-search-label">
                    <input type="search" placeholder="請輸入關鍵字" class="form-control" />
                </label>
                <button type="button" role="button" class="btn btn-light input-search-btn">
                    搜尋
                </button>
            </div>
            <button class="navbar-toggler" type="button" role="button" data-bs-toggle="collapse"
                data-bs-target="#navbarMainToggle" aria-controls="navbarMainToggle" aria-expanded="false"
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse nav-log" id="navbarMainToggle">
                <template v-if="!isLogined">
                    <!-- 登出中 -->
                    <div class="nav-logouting" id="navLogout">
                        <!-- <a
                            :href=""
                            type="button"
                        >
                            登入
                        </a> -->
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
                        <button class="btn-avatar dropdown-toggle" data-bs-toggle="dropdown" role="button"
                            aria-expanded="false">
                            <img src="@/assets/img/avatar.jpg" alt="avatar" class="avatar-img" />
                        </button>
                        <ul class="nav-item dropdown-menu nav-logining-list">
                            <li class="ropdown-item nav-item">Mue hung</li>
                            <li class="ropdown-item nav-item">@muehung92</li>
                            <li class="ropdown-item nav-item">
                                <a href="" class="nav-link">頻道管理
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                                        class="bi bi-arrow-right-square-fill" viewBox="0 0 16 16">
                                        <path
                                            d="M0 14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v12zm4.5-6.5h5.793L8.146 5.354a.5.5 0 1 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708-.708L10.293 8.5H4.5a.5.5 0 0 1 0-1z" />
                                    </svg>
                                </a>
                            </li>

                            <button type="button" role="button" class="btn btn-light btn-logout">
                                登出
                            </button>
                        </ul>
                    </div>
                </template>
            </div>
        </div>
    </nav>
</template>

<script>
import { mapState } from 'pinia';
import { useLoginCheckStore } from '../stores/useLoginCheck.js';

export default {
    computed: {
        ...mapState(useLoginCheckStore, ['isLogined'])
    },
}
</script>
