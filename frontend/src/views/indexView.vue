<template>
    <section class="sec-h sec-home">
        <div class="container">
            <div class="row">
                <div class="col-lg-6 row p-0 cards card-stores">
                    <div v-for="store in storesData" :key="store.id" class="card card-store m-2 p-0" style="width: 16rem;">
                        <a :href="`store/${store.id}`" class="card-link">
                            <figure class="card-img-top">
                                <img :src="store.pic" alt="" class="img-fluid">
                            </figure>
                            <div class="card-body">
                                <h5 class="card-title">{{ store.name }}</h5>
                                <p class="card-text">{{ store.openedAt }} ~ {{ store.closedAt }}</p>
                                <span class="tag tag-store">{{ store.status }}</span>
                            </div>
                        </a>
                    </div>

                </div>
                <div class="col-lg-6">
                    <div id="mapSite">
                        <!-- <iframe
                            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d8849.149936277927!2d121.50771574266926!3d25.044567497194116!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x3442a972a8266295%3A0xece39c57ea7e4e12!2z5Y-w5YyX6LuK56uZ!5e0!3m2!1szh-TW!2stw!4v1685499297618!5m2!1szh-TW!2stw"
                            width="100%" height="450" style="border:0;" allowfullscreen="" loading="lazy"
                            referrerpolicy="no-referrer-when-downgrade"></iframe> -->
                        <img src="../assets/img/temp_map.png" alt="暫時放圖" class="img-fluid">
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
<script>
import { storesListApi } from '../api';
import '../api/model';
export default {
    data() {
        return {
            /** @type {Store[]} */
            storesData: [],
        }
    },
    async mounted() {
        const storeList = await storesListApi()
        console.log(storeList)
        this.storesData = storeList
        console.log(JSON.stringify(this.storesData))
    }
}
</script>