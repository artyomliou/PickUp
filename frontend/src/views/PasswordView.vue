<script setup>
import FooterComponent from '../components/FooterComponent.vue';
import NavbarComponent from '../components/NavbarComponent.vue';
</script>

<template>
  <NavbarComponent />
  <section class="sec-h sec-password">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-12 col-lg-8 sec-titlebox">
          <h2 class="sec-title">忘記密碼</h2>
          <p class="sec-sub">輸入您的註冊信箱</p>
        </div>
        <div class="col-12 col-lg-5">
          <form action="" class="form">
            <div class="form-group">
              <label class="label">
                <span class="input-head">帳號 (E-mail)</span>
                <input type="text" value="帳號 (E-mail)" class="form-control" />
              </label>
            </div>
            <button type="submit" class="btn btn-md mx-auto btn-sumit">送出</button>
          </form>
          <div class="text-center">
            還是有問題？
            <a href="mailto:service@theproject.com.tw" target="_blank" class="">請來信聯絡
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                class="bi bi-envelope-fill" viewBox="0 0 16 16">
                <path
                  d="M.05 3.555A2 2 0 0 1 2 2h12a2 2 0 0 1 1.95 1.555L8 8.414.05 3.555ZM0 4.697v7.104l5.803-3.558L0 4.697ZM6.761 8.83l-6.57 4.027A2 2 0 0 0 2 14h12a2 2 0 0 0 1.808-1.144l-6.57-4.027L8 9.586l-1.239-.757Zm3.436-.586L16 11.801V4.697l-5.803 3.546Z" />
              </svg>
            </a>
          </div>
        </div>
      </div>
    </div>
  </section>
  <FooterComponent />
</template>
<!-- <script>
  export default {}
</script> -->
