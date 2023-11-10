<template>
    <section class="sec-h sec-product">
        <div class="container py-3">
            <div class="row">
                <div class="col-12 py-3">
                    <!-- {{ product }} -->
                </div>
                <div class="col-lg-6 product-picarea">
                    <div class="nav-sticky">
                        <a href="../" class="btn-goback d-block mb-2 pt-3">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="bi bi-arrow-left-short"
                                viewBox="0 0 16 16">
                                <path fill-rule="evenodd"
                                    d="M12 8a.5.5 0 0 1-.5.5H5.707l2.147 2.146a.5.5 0 0 1-.708.708l-3-3a.5.5 0 0 1 0-.708l3-3a.5.5 0 1 1 .708.708L5.707 7.5H11.5a.5.5 0 0 1 .5.5z" />
                            </svg>
                            返回 60 嵐
                        </a>
                        <figure>
                            <img src="../assets/img/product/gold-olong.jpg" alt="" class="img-fluid">
                        </figure>
                        <div v-for="item in product.selectQuestions" :key="item.id">
                            <!-- <div class="card my-3">{{ item }}</div> -->
                            <div class="card my-3">{{ item.options }}</div>
                        </div>
                    </div>
                </div>
                <div class=" col-lg-6 product-infoarea mt-5">
                    <h2 class="product-title">{{ product.name }}</h2>
                    <h5 class="product-price">$ {{ product.price }}</h5>
                    <p class="product-intro">{{ product.intro }}</p>

                    <form @submit.prevent="submitProdunctForm" method="post" class="form form-product">
                        <template v-for="quiz in product.selectQuestions " :key="quiz.id">
                            <!-- quiz: {{ quiz }}>>><br> -->
                            <h5 class="head d-flex">
                                <span>{{ quiz.name }}</span>
                            </h5>
                            <template v-if="quiz.isRequired">
                                <span class="tag tag-reqired ms-auto">必填</span>

                                <!-- 必填沒填時 顯現 -->
                                <!-- <span class="tag tag-reqired ms-auto bg-danger text-white">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                                        class="bi bi-check" viewBox="0 0 16 16">
                                        <path
                                            d="M10.97 4.97a.75.75 0 0 1 1.07 1.05l-3.99 4.99a.75.75 0 0 1-1.08.02L4.324 8.384a.75.75 0 1 1 1.06-1.06l2.094 2.093 3.473-4.425a.267.267 0 0 1 .02-.022z" />
                                    </svg> 必填
                                </span> -->
                                <!-- 必填沒填時 顯現 -->
                                {{ customAnswer }}
                                <template v-for="option in   quiz.options" :key="option.id">
                                    <label class="form-control my-2 d-flex justify-content-between">
                                        <span>{{ option.name }}</span>
                                        <!-- selectQuestionId -->
                                        <input type="radio" :name="quiz.id" :value="option.id"
                                            v-model="customAnswer[quiz.id]" class="radio"> <!-- like map -->
                                        {{ option.selectQuestionId }}
                                    </label>
                                </template>

                            </template>

                            <!-- 非必填 -->
                            <template v-else>
                                <template v-for="frSelect in   quiz.options" :key="frSelect.id">
                                    <label class="form-control my-2 d-flex justify-content-between">
                                        <p class="select-name strong selectName">{{ frSelect.name }}</p>
                                        <p class="select-info">{{ frSelect.info }}</p>
                                        <p class="price">$ {{ frSelect.price }}</p>
                                        <input v-model="product.selectQuestionId" :value="`otherAdded${frSelect.id}`"
                                            type="checkbox" class="radio">
                                    </label>
                                </template>
                            </template>
                        </template>

                        <template v-for="quizCustom in  product.customQuestions " :key="quizCustom.id">
                            <label for="" class="">
                                <h5 class="head">{{ quizCustom.name }}</h5>
                                <textarea v-model="product.notes" class="form-control form-textarea" name=""
                                    :id="quizCustom.id" cols="30" rows="5" :placeholder="quizCustom.placeholder"></textarea>
                            </label>
                        </template>

                        <label for="" class="py-3 d-flex">
                            <span class="selectName">數量</span>
                            <select v-model="product.cups" id="" class="ms-3">
                                <option value="1" selected>1</option>
                                <option value="2">2</option>
                                <option value="3">3</option>
                            </select>
                        </label>

                        <div class="d-grid py-3">
                            <button type="submit" class="btn btn-submit p-3">
                                新增
                                <span class="totalCount">{{ product.cups }}</span> 產品
                                • 至購物車
                            </button>
                        </div>
                    </form>

                </div>
            </div>
        </div>
    </section>
    <FooterComponent />
</template>

<script>
import { productApi } from "../api/index";
import "../api/model";
import "../assets/css/product.css";
export default {
    data() {
        return {
            /** @type {Store} */
            product: {},
            customAnswer: {},

        }
    },
    async created() {
        const storeId = this.$route.params.storeId;
        const prodId = this.$route.params.productId;
        const data = await productApi(storeId, prodId);
        this.product = data;
        // console.log(data);
        // this.product = data;
        // renderData(data);
        // console.log(this.product.cups)
    },
    methods: {
        submitProdunctForm(e) {
            e.preventDefault();
            console.log('this.product');
            console.log(this.product);
            // 1. Check value validate

            // 2. api post 回傳

        }
    },
    // computed: {
    //     renderData(data) {
    //         this.product.ice = data.ice,
    //             this.product.suger = data.suger,
    //             this.product.add = data.add,
    //             this.product.notes = data.notes,
    //             this.product.cups = data.cups,
    //     }
    // }

}
</script>