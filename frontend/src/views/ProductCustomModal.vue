<template>
    <div class="modal" tabindex="-1" id="productmodal">
        <div class="modal-dialog modal-xl">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">{{ product.name }}</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <form id="productmodal-form" @submit.prevent>
                        <div class="d-flex">
                            <div class="col">
                                <p>{{ product.intro }}</p>
                                <p>售價：{{ product.price }}</p>
                            </div>
                            <div class="col">
                                <template v-for="q in product.selectQuestions" :key="q.id">
                                    <hr>
                                    <h3>{{ q.name }}</h3>
                                    <span class="text-warning" v-if="q.isRequired">必填</span>
                                    <div class="form-check" v-for="o in q.options" :key="o.id">
                                        <input class="form-check-input" v-model="selectAnswers[q.id]" :name="`sq-${q.id}`"
                                            :required="q.isRequired ? 'required' : ''"
                                            :type="q.atMost == 1 ? 'radio' : 'checkbox'" :value="o.id">
                                        <label class="form-check-label" for="flexCheckDefault">
                                            {{ o.name }}
                                        </label>
                                    </div>
                                </template>

                                <template v-for="q in product.customQuestions" :key="q.id">
                                    <hr>
                                    <h3>{{ q.name }}</h3>
                                    <textarea class="form-control" rows="3" :placeholder="q.placeholder"
                                        v-model="customAnswers[q.id]"></textarea>
                                </template>

                                <hr>
                                <h3>數量</h3>
                                <select class="form-select" aria-label="Chhose mount" v-model="amount">
                                    <option value="1">1</option>
                                    <option value="2">2</option>
                                    <option value="3">3</option>
                                    <option value="4">4</option>
                                    <option value="5">5</option>
                                    <option value="6">6</option>
                                </select>
                            </div>
                        </div>
                    </form>
                </div>
                <div class="modal-footer">
                    <!-- <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button> -->
                    <button type="submit" form="productmodal-form" @click="handleInputValidation($event) && callSaveApi()"
                        class="btn btn-primary">加入購物車</button>
                </div>
            </div>
        </div>
    </div>
</template>

<style></style>

<script>
import { Modal } from 'bootstrap';
import { createCartItem, fetchProduct } from '../api';
import '../api/model';

let modal = null

export default {
    data() {
        return {
            /** @type {Product} */
            product: {},

            amount: 1,
            /** @type {Object.<number, number | number[]>} */
            selectAnswers: {},
            /** @type {Object.<number, number | number[]>} */
            customAnswers: {}
        }
    },
    async mounted() {
        // 抓資料
        console.log(`sid=${this.$route.params.sid} pid=${this.$route.params.pid}`)
        const product = await fetchProduct(this.$route.params.sid, this.$route.params.pid)
        console.log(product)
        this.product = product

        // 初始化表單
        this.product.selectQuestions.forEach(q => {
            this.selectAnswers[q.id] = []
        })
        this.product.customQuestions.forEach(q => {
            this.customAnswers[q.id] = ''
        })

        // 顯示燈箱
        const node = document.querySelector('#productmodal')
        modal = new Modal(node)
        modal.show()

        // 當關閉燈箱（點叉叉、點燈箱周圍）時更新 router
        node.addEventListener('hidden.bs.modal', () => {
            this.$router.back()
        })
    },
    methods: {
        handleInputValidation(ev) {
            /** @type {HTMLFormElement} */
            const form = ev.target.form
            return form.checkValidity()
        },
        async callSaveApi() {
            // Transform from object of records to array of objects
            const selectAnswers = []
            for (const [qid, options] of Object.entries(this.selectAnswers)) {
                /** @type {SelectAnswer} */
                const answer = {
                    selectQuestionId: Number.parseInt(qid, 10),
                    options: Array.isArray(options) ? options : [options]
                }
                selectAnswers.push(answer)
            }

            // Transform from object of records to array of objects
            const customAnswers = []
            for (const [qid, text] of Object.entries(this.customAnswers)) {
                /** @type {CustomAnswer} */
                const answer = {
                    customQuestionId: Number.parseInt(qid, 10),
                    text
                }
                customAnswers.push(answer)
            }

            const ok = await createCartItem(this.product.storeId, this.product.id, this.amount, selectAnswers, customAnswers)
            if (ok) {
                if (modal) {
                    modal.hide()
                }
                this.$router.push(`/store/${this.$route.params.sid}/cart`)
            } else {
                alert("發生錯誤，請稍後再試")
            }
        }
    },
}
</script>