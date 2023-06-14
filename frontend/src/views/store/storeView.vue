<template>
    <NavbarComponent />
    <section class="sec-h sec-store">
        <div class="banner" :style="`background-image: url('${storeItem.pic}')`">
        </div>
        <div class="container py-3">
            <div class="row">
                <!-- <h5>$route.params.id: {{ $route.params.id }}</h5> -->
                <aside class="col-lg-4">
                    <div class="store-info">
                        <h2 class="store-title">{{ storeItem.name }}</h2>
                        <address>!!100台北市中正區鄭州路8號</address>
                        <p class="card-text">{{ storeItem.openedAt }} ~ {{ storeItem.closedAt }}</p>
                        <span class="tag tag-store">{{ storeItem.status }}</span>
                    </div>
                    <div class="store-classes">
                        <ul class="nav class-list">
                            <tempalte v-for="item in storeItem.categories">
                                <li>{{ item.name }}</li>
                            </tempalte>
                        </ul>
                    </div>
                </aside>

                <template v-for="cat in storeItem.categories" :key="cat.id">
                    <div class="col-lg-8 store-cards">
                        <h4 class="series-title">{{ cat.name }}</h4>
                        <div class="card-group justify-content-center cards">

                            <!-- card -->
                            <div class="store-card">
                                <a href="/item" class="store-link">
                                    <figure class="card-img">
                                        <img src="../../assets/img/store-item2.jpg" alt="" class="img-fluid">
                                    </figure>
                                    <div class="card-info">
                                        <h4 class="item-title">紫色特調飲</h4>
                                        <p class="store-prize">$ 70.00</p>
                                    </div>
                                    <button class="btn store-btn">
                                        <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor"
                                            class="bi bi-plus-circle" viewBox="0 0 16 16">
                                            <path
                                                d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z" />
                                            <path
                                                d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z" />
                                        </svg>
                                    </button>
                                </a>
                            </div>

                        </div>
                    </div>
                </template>

            </div>
        </div>
    </section>
    <FooterComponent />
</template>

<script>
import { storeDataApi } from "../../api/index";
import "../../api/model";
import "../../assets/css/store.css";
export default {
    data() {
        return {
            /** @type {Store} */
            storeItem: {},
        }
    },
    // computed: {
    //     destinationId() {
    //         return parseInt(this.$route.params.id)
    //     }
    // },
    async created() {
        const id = this.$route.params.id;
        const data = await storeDataApi(id);
        console.log(data)
        this.storeItem = data
        console.log(this.storeItem)

        // this.storeItem = data[0].find(id) => this.storeItem.id == id;
        // async destination() {
        //     console.log('------mounted----')
        //     const store = await storeDataApi()
        //     // const data = Object.values(store)
        //     console.log('destination store')
        //     // console.log(...data[0])
        //     // return data
        //     console.log(store)
        //     return store
        // }
    },

}
</script>