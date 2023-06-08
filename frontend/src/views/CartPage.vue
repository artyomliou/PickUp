<template>
    <h1>購物車</h1>
    <ul>
        <li v-for="(i, idx) in items" :key="i.id">
            {{ i.product.name }} {{ i.amount }}
            <button type="button" class="btn btn-danger" @click="callRemoveItemApi(i.id, idx)">移除</button>
        </li>
    </ul>
    <button class="btn" @click="callCreateOrderApi">送出訂單</button>
</template>

<script>
import { createOrder, fetchCartItemList, removeCartItem } from '../api';
import '../api/model';

export default {
    async beforeMount() {
        const itemList = await fetchCartItemList(this.$route.params.sid)
        console.log(itemList)
        this.items = itemList
    },
    data() {
        return {
            /** @type {CartItem[]} */
            items: []
        }
    },
    methods: {
        async callRemoveItemApi(itemId, idx) {
            const success = await removeCartItem(this.$route.params.sid, itemId)
            if (success) {
                this.items.splice(idx, 1) // Remove item only from UI
            }
        },
        async callCreateOrderApi() {
            const order = await createOrder(this.$route.params.sid)
            this.$router.push(`/order/${order.id}`)
        }
    }
}
</script>