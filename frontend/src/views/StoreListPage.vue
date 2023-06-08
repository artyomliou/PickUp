<template>
    <ul>
        <li v-for="s in stores" :key="s.id">
            <router-link :to="`/store/${s.id}`">{{ s.name }}</router-link>
            {{ `(${s.openedAt} ~ ${s.closedAt}) ${s.status}` }}
        </li>
    </ul>
</template>

<style></style>

<script>
import { fetchStoreList } from '../api/index';
import '../api/model';

export default {
    data() {
        return {
            /** @type {Store[]} */
            stores: [],
        }
    },
    async mounted() {
        const stores = await fetchStoreList()
        console.log(stores)
        this.stores = stores
    }
}
</script>