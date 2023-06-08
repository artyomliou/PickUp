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
import { mapActions } from 'pinia';
import { isLoggedInApi } from './api';
import FooterComponent from './components/FooterComponent.vue';
import NavbarComponent from './components/NavbarComponent.vue';
import { useLoginStore } from './store/login';
export default {
  components: { NavbarComponent, FooterComponent },
  data() {
    return {
      loading: false,
      post: null,
      error: null,
    }
  },
  methods: {
    ...mapActions(useLoginStore, ['updateLoginStatus'])
  },
  async beforeMount() {
    this.updateLoginStatus(await isLoggedInApi())
  }
}
</script>