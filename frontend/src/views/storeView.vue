<template>
    <section class="sec-h sec-store">
        <div class="banner" :style="`background-image: url('${storeItem.pic}')`">
        </div>
        <div class="container py-3">
            <div class="row">
                <nav class="col-lg-4">
                    <div class="nav-sticky pt-3">
                        <div class="store-info">
                            <h2 class="store-title">{{ storeItem.name }}</h2>
                            <address>100 台北市中正區鄭州路8號</address>
                            <p class="card-text">{{ storeItem.openedAt }} ~ {{ storeItem.closedAt }}</p>
                            <span class="tag tag-store">{{ storeItem.status }}</span>
                        </div>
                        <div class="store-classes">
                            <ul class="list-group class-list" role="tablist" id="listItems">
                                <template v-for="item in storeItem.categories" :key="item.key">
                                    <a class="list-group-item list-group-item-action list-item-link"
                                        :href="`#list-item-${item.id}`">
                                        <li class="list-item">{{
                                            item.name }}
                                        </li>
                                    </a>
                                </template>
                            </ul>
                        </div>
                    </div>

                </nav>

                <div class="col-lg-8">
                    <template v-for="cat in storeItem.categories" :key="cat.id">
                        <div class="tabContent store-cards" data-bs-spy="scroll" data-bs-target="#listItems"
                            data-bs-offset="0" tabindex="0">
                            <h4 class="series-title" :id="`list-item-${cat.id}`">{{ cat.name }}</h4>
                            <div class="card-group cards">
                                <template v-for="product in cat.products" :key="product.id">
                                    <!-- card -->
                                    <!-- {{ product }} -->
                                    <div class="store-card" role="tabpanel">
                                        <a :href="`${$route.params.storeId}/product/${product.id}`" class="store-link">
                                            <figure class="card-img">
                                                <img src="../../assets/img/store-item2.jpg" alt="" class="img-fluid">
                                            </figure>
                                            <div class="card-info">
                                                <h5 class="item-title">{{ product.name }}</h5>
                                                <p class="item-intro">{{ product.intro }}</p>
                                                <p class="store-prize">$ {{ product.price }}</p>
                                            </div>
                                            <a class="store-btn">
                                                <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor"
                                                    class="bi bi-plus-circle" viewBox="0 0 16 16">
                                                    <path
                                                        d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
                                                    <path
                                                        d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z" />
                                                </svg>
                                            </a>
                                        </a>
                                    </div>
                                </template>
                            </div>
                        </div>
                    </template>
                </div>


            </div>
        </div>
    </section>
    <FooterComponent />
</template>

<script>
// import { Transition } from "vue";
import { storeDataApi } from "../api/index";
import "../api/model";
import "../assets/css/store.css";
export default {
    data() {
        return {
            /** @type {Store} */
            storeItem: {},
        }
    },
    async created() {
        const id = this.$route.params.storeId;
        const data = await storeDataApi(id);
        console.log(data)
        this.storeItem = data
        console.log(this.storeItem)
    },
    methods: {
        // 監聽
    }

}
</script>