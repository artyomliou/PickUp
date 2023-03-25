<script setup>
import Navbar from '../components/Navbar.vue'
import FooterComponent from '../components/Footer.vue'
</script>

<template>
  <Navbar />
  <section class="sec-h sec-login">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-12 col-lg-8 sec-titlebox">
          <h2 class="sec-title">登入</h2>
          <p class="sec-sub">歡迎回來！</p>
        </div>
        <div class="col-12 col-lg-5">
          <form action="" class="form">
            <div class="form-group">
              <label class="label">
                <span class="input-head">帳號 (E-mail)</span>
                <input type="text" value="帳號 (E-mail)" class="form-control" />
              </label>
            </div>
            <div class="form-group">
              <label class="label">
                <span class="input-head">密碼</span>
                <input type="password" value="密碼" class="form-control" />
              </label>
            </div>
            <button type="submit" class="btn btn-md mx-auto btn-sumit">送出</button>
          </form>
          <div class="text-center">
            <a href="" class="btn btn-inline">忘記密碼？</a>
            <a href="" class="btn btn-inline">還沒註冊？</a>
          </div>
        </div>
      </div>
    </div>
  </section>
  <FooterComponent />
</template>
<script>
export default {}
</script>
