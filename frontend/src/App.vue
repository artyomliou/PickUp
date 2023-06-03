<script setup>
import { RouterView } from 'vue-router';
</script>


<template>
  <NavbarComponent />
  <div class="post">
    <div v-if="loading" class="loading">Loading...</div>

    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else-if="post" class="content">
      <h2>{{ post.title }}</h2>
    </div>
  </div>
  <RouterView />
  <FooterComponent />
</template>

<script>
import { computed } from 'vue';
import FooterComponent from './components/FooterComponent.vue';
import NavbarComponent from './components/NavbarComponent.vue';
import { isLoginStatus, loginAction, logoutAction } from './js/api/index';
export default {
  components: { NavbarComponent, FooterComponent },
  data() {
    return {
      loading: false,
      post: null,
      error: null,
      isLoginedStatus: "isLoginStatus had error?",
      isLogined: false,
    }
  },
  async created() {
    this.isLogined = await isLoginStatus()
  },
  provide() {
    return {
      isLoginedStatusCheck: computed(() => this.isLogined), //確認是否登入 boolean
      loginAction: async (email, password) => { // 輸入帳密給登入 function
        if (await loginAction(email, password)) {
          this.isLogined = true
          this.$router.push('/')
        }
      },
      logoutAction: async () => {
        if (await logoutAction()) {
          this.isLogined = false
        }
      }
    }
  },
}
</script>