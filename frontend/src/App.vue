<script setup>
import { RouterView } from 'vue-router';
// import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <!-- <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
          <RouterLink to="/">Home</RouterLink>
          <RouterLink to="/about">About</RouterLink>
        </nav>
      </div>
    </header> -->

  <RouterView />
</template>

<style scoped></style>

<script>
// every pages
import { mapActions } from 'pinia';
import { useLoginCheckStore } from './stores/useLoginCheck.js';

export default {
  methods: {
    ...mapActions(useLoginCheckStore, ['fetchLoginApi'])
  },
  created() {
    this.fetchLoginApi()
  }
}

</script>