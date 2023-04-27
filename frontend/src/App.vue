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
import { isLoginStatus, loginAction, logoutAction } from './api/index';
import FooterComponent from './components/FooterComponent.vue';
import NavbarComponent from './components/NavbarComponent.vue';
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
      isLoginedStatusId: computed(() => this.isLogined),
      loginAction: async (email, password) => {
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