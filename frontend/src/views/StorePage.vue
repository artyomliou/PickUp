<template>
    <div>
        <h1>Store</h1>
        <router-link :to="`/store/${store.id}/cart`">購物車</router-link>
        <ul>
            <li>{{ store.id }}</li>
            <li>{{ store.name }}</li>
            <li>{{ store.openedAt }}</li>
            <li>{{ store.closedAt }}</li>
            <li>{{ store.pic }}</li>
            <li>{{ store.status }}</li>
        </ul>

        <h1>Categories</h1>
        <ul>
            <li v-for="c in store.categories" :key="c.id">
                <a :href="`#${c.name}`">{{ c.name }}</a>
            </li>
        </ul>

        <h1>Products</h1>
        <template v-for="c in store.categories" :key="c.id">
            <h3 :id="c.id">{{ c.name }}</h3>
            <ul>
                <li v-for="p in c.products" :key="p.id">
                    <router-link :to="`/store/${store.id}/product/${p.id}`">
                        {{ p.name }}
                    </router-link>
                </li>
            </ul>
        </template>
    </div>
    <router-view></router-view>
</template>

<style></style>

<script>
import { fetchStore } from '../api';
import '../api/model';

export default {
    async beforeMount() {
        const sid = this.$route.params.sid
        console.log(sid)
        const store = await fetchStore(sid)
        console.log(store)
        this.store = store
    },
    data() {
        return {
            /** @type {Store} */
            store: {},
        }
    },
}
</script>