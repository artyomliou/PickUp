<template>
    <h1>訂單 {{ order.id }}</h1>
    <span>(即將在 {{ refreshCountdown }} 秒後更新狀態)</span>
    <ul>
        <li>售價：{{ order.price }}</li>
        <li>狀態：{{ order.status }}</li>
    </ul>
</template>

<script>
import { fetchOrder, fetchOrderStatus } from '../api';

export default {
    async mounted() {
        const order = await fetchOrder(this.$route.params.oid)
        console.log(order)
        this.order = order

        setInterval(async () => {
            if (this.refreshCountdown > 0) {
                this.refreshCountdown -= 1
            } else {
                this.refreshCountdown = 5
                this.order.status = await fetchOrderStatus(this.order.id)
            }
        }, 1000)
    },
    data() {
        return {
            /** @type {Order} */
            order: [],
            refreshCountdown: 5,
        }
    },
}
</script>