<script setup>
import { RouterView } from 'vue-router';
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
import { useLoginStatusStore } from '@/stores/LoginStatus.js';
import { mapActions, } from 'pinia';

export default {
  methods: {
    ...mapActions(useLoginStatusStore, ['checkLoginStatus'])
  },
  created() {
    this.checkLoginStatus()
  }
}
</script>