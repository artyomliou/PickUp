<template>
    <NavbarComponent />
    <section class="sec-h sec-product">
        <div class="container py-3">
            <div class="row">
                <div class="col-12">
                    <a href="store/{{this.$route.params.storeId}}"></a>
                </div>
                <div class="col-lg-6">
                    <figure>
                        <img src="../../../assets/img/store-item2.jpg" alt="" class="img-fluid">
                    </figure>
                </div>
                <div class="col-lg-6">
                    <!-- {{ product }} -->
                    <h2 class="product-title">{{ product.name }}</h2>
                    <h5 class="product-price">$ {{ product.price }}</h5>
                    <p class="product-intro">{{ product.intro }}</p>

                    <form action="" class="from from-product">
                        <template v-for="quiz in    product.selectQuestions " :key="quiz.id">
                            <!-- quiz: {{ quiz }}>>><br>
                            <hr> -->
                            <h5 class="">
                                {{ quiz.name }}
                                <template v-if="quiz.isRequired">
                                    <span class="tag tag-reqired">required</span>
                                </template>
                            </h5>
                            <template v-for="select in   quiz.options" :key="select.id">
                                <!-- select: {{ select }}>>><br> -->
                                <label :for="`quiz${select.selectQuestionId}`" class="form-control m-2">
                                    {{ select.name }}
                                    <input type="radio" :name="`quiz${select.selectQuestionId}`" class="radio">
                                </label>
                            </template>
                        </template>

                        <template v-for="quizCustom in  product.customQuestions " :key="quizCustom.id">
                            <label for="" class="from-control">
                                {{ quizCustom.name }}
                                <textarea class="from-control form-textarea" name="" id="quizCustom.id" cols="30" rows="5"
                                    :placeholder="quizCustom.placeholder"></textarea>
                            </label>
                        </template>

                        <select name="" id="" class="">
                            <option value="1">1</option>
                            <option value="1">2</option>
                            <option value="1">3</option>
                        </select>
                        <div class="d-grid p-3">
                            <button type="button" class="btn btn-submit">新增</button>
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