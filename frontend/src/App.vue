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
import { mapWritableState } from 'pinia';
import { isLoginStatus } from './api';
import FooterComponent from './components/FooterComponent.vue';
import NavbarComponent from './components/NavbarComponent.vue';
import { useLoginStatusStore } from './stores/useLoginStatusStore';
export default {
  components: { NavbarComponent, FooterComponent },
  data() {
    return {
      loading: false,
      post: null,
      error: null,
    }
  },
  computed: {
    ...mapWritableState(useLoginStatusStore, ['isLogin']),
  },
  async watch() {

  },
  async mounted() {
    // 打api後拿到 isLogin 值，存入 store
    const store = useLoginStatusStore();
    let status = await isLoginStatus()
    console.log('status---')
    console.log(status)
    store.isLogin = status;
    if (store.isLogin) {
      console.log('store.isLogin = true')
      return store.isLogin
    } else {
      store.isLogin = false;
      console.log('store.isLogin = false;')
      return store.isLogin
    }
  },
}
</script>