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
                    </div>
                </div>
                <div class=" col-lg-6 product-infoarea mt-5">
                    <h2 class="product-title">{{ product.name }}</h2>
                    <h5 class="product-price">$ {{ product.price }}</h5>
                    <p class="product-intro">{{ product.intro }}</p>

                    <form @submit.prevent="submitProdunctForm" method="post" class="form form-product">
                        <template v-for="selectQuestion in product.selectQuestions " :key="selectQuestion.id">
                            <!-- <div v-for="s in selectQuestion" :key="s.id">
                                <div v-for=" ss in s" :key="ss.id">ss name: {{ ss.name }}</div>
                            </div> -->
                            {{ requiredAnswer }}
                            <template v-if="!!selectQuestion.isRequired">
                                <h5 class="head d-flex">
                                    <span>{{ selectQuestion.name }}</span>
                                </h5>
                                <span class="tag tag-reqired ms-auto">必填</span>

                                <template v-for="option in selectQuestion.options" :key="option.id">
                                    <label class="form-control my-2 d-flex justify-content-between">
                                        <span>{{ option.name }}</span>
                                        <!-- selectQuestionId -->
                                        <input type="radio" :name="selectQuestion.id" :value="option.id"
                                            v-model="requiredAnswer[selectQuestion.id]" class="radio"> <!-- like map -->
                                        <!-- {{ option.selectQuestionId }} -->
                                    </label>
                                </template>
                                !!!!!!!!
                            </template>

                            <!-- 非必填 -->
                            <template v-else>
                                <h4 class="text-danger">非必填</h4>
                                <h5 class="head d-flex">
                                    <span>{{ selectQuestion.name }}</span>
                                    {{ unrequiredAnswer }}
                                </h5>
                                <template v-for="option in selectQuestion.options" :key="option.id">
                                    <label class="form-control my-2 d-flex justify-content-between">
                                        <p class="select-name strong selectName">{{ option.name }}</p>
                                        <p class="select-info">{{ option.info }}</p>
                                        <p class="price">$ {{ option.price }}</p>
                                        <input v-model="unrequiredAnswer" :value="option.id" type="checkbox" class="radio">
                                    </label>
                                </template>
                            </template>
                        </template>

                        <template v-for="selectQuestionCustom in  product.customQuestions " :key="selectQuestionCustom.id">
                            <label for="" class="">
                                <h5 class="head">{{ selectQuestionCustom.name }}</h5>
                                <textarea v-model="product.notes" class="form-control form-textarea" name=""
                                    :id="selectQuestionCustom.id" cols="30" rows="5"
                                    :placeholder="selectQuestionCustom.placeholder"></textarea>
                            </label>
                        </template>

                        <label for="" class="py-3 d-flex">
                            <span class="selectName">數量</span>
                            <select v-model="productPiece" id="" class="ms-3">
                                <option value="1" selected>1</option>
                                <option value="2">2</option>
                                <option value="3">3</option>
                            </select>
                        </label>

                        <div class="d-grid py-3">
                            <button type="submit" class="btn btn-submit p-3">
                                新增
                                <span class="totalCount">{{ productPiece }}</span> 產品
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
            requiredAnswer: {},
            unrequiredAnswer: [],
            productPiece: 1,
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