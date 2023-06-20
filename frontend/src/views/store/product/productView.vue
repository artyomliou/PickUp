<template>
    <NavbarComponent />
    <section class="sec-h sec-product">
        <div class="container py-3">
            <div class="row">
                <!-- <div class="col-12 py-3">
                    
                </div> -->
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
                            <img src="../../../assets/img/store-item2.jpg" alt="" class="img-fluid">
                        </figure>
                    </div>
                </div>
                <div class="col-lg-6 product-infoarea mt-5">
                    <!-- {{ product }} -->
                    <h2 class="product-title">{{ product.name }}</h2>
                    <h5 class="product-price">$ {{ product.price }}</h5>
                    <p class="product-intro">{{ product.intro }}</p>

                    <form action="" class="form form-product">
                        <template v-for="quiz in    product.selectQuestions " :key="quiz.id">
                            <!-- quiz: {{ quiz }}>>><br>
                            <hr> -->
                            <h5 class="head d-flex">
                                <span>{{ quiz.name }}</span>
                                <template v-if="quiz.isRequired">
                                    <span class="tag tag-reqired ms-auto">必填</span>
                                    <span class="tag tag-reqired ms-auto bg-warning">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                                            class="bi bi-exclamation" viewBox="0 0 16 16">
                                            <path
                                                d="M7.002 11a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 4.995a.905.905 0 1 1 1.8 0l-.35 3.507a.553.553 0 0 1-1.1 0L7.1 4.995z" />
                                        </svg>
                                        必填</span>
                                    <span class="tag tag-reqired ms-auto bg-danger text-white">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                                            class="bi bi-check" viewBox="0 0 16 16">
                                            <path
                                                d="M10.97 4.97a.75.75 0 0 1 1.07 1.05l-3.99 4.99a.75.75 0 0 1-1.08.02L4.324 8.384a.75.75 0 1 1 1.06-1.06l2.094 2.093 3.473-4.425a.267.267 0 0 1 .02-.022z" />
                                        </svg> 必填</span>
                                </template>
                            </h5>
                            <template v-for="select in   quiz.options" :key="select.id">
                                <!-- select: {{ select }}>>><br> -->
                                <label :for="`quiz${select.selectQuestionId}`"
                                    class="form-control my-2 d-flex justify-content-between">
                                    <span>{{ select.name }}</span>
                                    <input type="radio" :name="`quiz${select.selectQuestionId}`" class="radio">
                                </label>
                            </template>
                        </template>

                        <template v-for="quizCustom in  product.customQuestions " :key="quizCustom.id">
                            <label for="" class="">
                                <h5 class="head">{{ quizCustom.name }}</h5>
                                <textarea class="form-control form-textarea" name="" :id="quizCustom.id" cols="30" rows="5"
                                    :placeholder="quizCustom.placeholder"></textarea>
                            </label>
                        </template>

                        <h5 class="head">加料</h5>
                        <label for="" class="form-control d-flex my-2">
                            <div class="select-item">
                                <p class="select-name strong selectName">珍珠</p>
                                <p class="select-info">+$10.00</p>
                            </div>
                            <input type="checkbox" name="" id="" class="d-block ms-auto">
                        </label>

                        <label for="" class="py-3 d-flex">
                            <span class="selectName">數量</span>
                            <select name="" id="" class="ms-3">
                                <option value="1">1</option>
                                <option value="2">2</option>
                                <option value="3">3</option>
                            </select>
                        </label>
                        <div class="d-grid py-3">
                            <button type="button" class="btn btn-submit p-3">
                                新增
                                <span class="totalCount">1</span> 產品
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
import { productApi } from "../../../api/index";
import "../../../api/model";
import "../../../assets/css/product.css";
export default {
    data() {
        return {
            /** @type {Store} */
            product: {},
        }
    },
    async created() {
        const storeId = this.$route.params.storeId;
        const prodId = this.$route.params.productId;
        const data = await productApi(storeId, prodId);
        console.log(data)
        this.product = data
        console.log(this.product)

    },
    methods: {
        // Check required is filled
    }

}
</script>