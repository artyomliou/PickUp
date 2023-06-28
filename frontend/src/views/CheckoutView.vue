<template>
    <NavbarComponent />
    <section class="sec-h sec-cart">{{ $route.params.storeId }}
        <!-- this.$route.params.storeId: {{ this.$route.params.storeId }} -->
        <!--  -->
    </section>
    <div class="" storeId=""></div>
    <FooterComponent />
</template>

<script>
import { cartItemsApi } from "../api";
import "../api/model";
import "../assets/css/product.css";
import '../router/index';
export default {
    props: ['storeId'],
    // data() {
    //     return {
    //         /** @property {number} id */
    //         loclaStoreId: this.$route.params.storeId
    //     }
    // },
    async created() {
        const data = await cartItemsApi(this.storeId);
        console.log(data)
        // const [productId, amount, selectAnswers, customAnswers] 
        // const cart = await cartNewItemApi(storeId);
        // console.log(cart)
    },
    methods: {
        // Check required is filled
    }

}
</script>