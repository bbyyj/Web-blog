<template>
    <div class="StudyTest">
        <transition appear name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__slideInUp">
                <div class="wrap">
                    <div class="statement">
                    </div>
                    <div>
                        <!-- 问题和答案部分 -->
                        <div class="questionnaire" v-for="(question, qIndex) in questions" :key="qIndex">
                            <h2>{{ question.text }}</h2>
                            <ul>
                                <li v-for="(answer, aIndex) in question.answers" :key="aIndex">
                                    <input type="radio" :id="`q${qIndex}a${aIndex}`" :name="`question${qIndex}`"
                                        :value="answer" v-model="userAnswers[qIndex]">
                                    <label :for="`q${qIndex}a${aIndex}`">{{ answer }}</label>
                                </li>
                            </ul>
                        </div>

                        <!-- 提交按钮 -->
                        <button v-if="!isSubmitted" @click="submitAnswers" :disabled="isScoring">提交</button>
                        <!-- 重新开始按钮 -->
                        <button v-else @click="resetTest" :disabled="isScoring">重新测试</button>

                        <div class="loader" v-bind:class="{ show: isScoring }"></div>

                        <!-- 显示得分 -->
                        <div class="answer" v-if="showScore">
                            <h3>你的得分是：{{ score }}</h3>
                            <h1>{{ messages[score] }}</h1>
                        </div>
                    </div>
                </div>
        </transition>
    </div>
</template>

<script>

import "../assets/icon/iconfont.css"
//声明组件
import TitleArea from "../components/TitleArea";

import "animate.css"

export default {
    components: { TitleArea },
    data() {
        return {
            //所有可选科目
            subjects: [],
            questions: [
                //设置的10个问题
            ],
            //将问题选项的按钮置位空
            userAnswers: Array(10).fill(null),
            //是否显示评分
            showScore: false,
            //评分
            score: 0,
            //设置正在评分
            isScoring: false,
            //是否提交评分过  评分过按钮会变成"重新测试"
            isSubmitted: false,
            //最高分为100
            maxScore: 100,
            //设置不同评分对应的话语
            messages: {
                0: '你似乎还没有开始学习这个主题。',
                10: '你已经初步了解这个主题，但还有许多需要学习的地方。',
                20: '你已经掌握了一些基础知识，但还需要进一步的学习。',
                30: '你对这个主题有一定的理解，但还有很多需要提升的地方。',
                40: '你的知识掌握程度还可以，但还有很大的提升空间。',
                50: '你已经掌握了这个主题的一半知识，继续努力！',
                60: '你的知识掌握程度相当不错，但不要停下来，继续努力！',
                70: '你已经对这个主题有了深入的了解，再继续努力就更完美了！',
                80: '你的知识掌握程度相当高，你已经接近这个主题的专家了！',
                90: '你的知识掌握程度接近完美，你是这个主题的专家了！',
                100: '你的知识掌握太到位了！你已经是这个主题的大师了！'
            }
        }
    },

    created() {
        window.scrollTo(0, 0);
        this.getQuestionInfo()
    },
    methods: {
        //获取对应的题目
        async getQuestionInfo() {
            let title = this.$route.query.title;
            let typename = this.$route.query.typename;
            const { data: res } = await this.$axios.get("/myblog/examList", { params: { typename: typename, title: title } });
            if (res.status === 1) {
                this.questions = res.data[0];
                console.log(res.data[0]);
            }
            else {
                print("fail")
            }
        },

        //点击提交按钮
        submitAnswers() {
            if (this.userAnswers.includes(null)) {
                alert("请回答所有的问题后再进行评分！");
                return;
            }
            if (!confirm("你确定提交你的答案吗")) {
                return;
            }
            //已经提交
            this.isSubmitted = true;
            //显示加载圆圈
            this.isScoring = true;

            //设置时间延迟
            setTimeout(() => {
                //计算评分
                this.score = this.calculateScore();
                //显示评分
                this.showScore = true;
                //关闭加载圆圈
                this.isScoring = false;
            }, 3000);
        },

        //重置按钮
        resetTest() {
            if (!confirm("你重置你的答案吗")) {
                return;
            }
            this.userAnswers = Array(10).fill(null);
            this.showScore = false;
            this.isSubmitted = false;
            this.score = 0;
        },
        //计算评分
        calculateScore() {
            let score = 0;
            for (let i = 0; i < this.questions.length; i++) {
                if (this.userAnswers[i] === this.questions[i].correctAnswer) {
                    score += 10;
                }
            }
            return score;
        },

    }
}
</script>



<style lang="less" scoped>

.StudyTest {
    min-height: 2000px;
    background-attachment: fixed;
    background-repeat: no-repeat;
    background-color: #7378ac;
    background-image:
        radial-gradient(closest-side, #aab7cc, rgba(235, 105, 78, 0)),
        radial-gradient(closest-side, #7285ac, rgba(243, 11, 164, 0)),
        radial-gradient(closest-side, #919fbe, rgba(254, 234, 131, 0)),
        radial-gradient(closest-side, #e0e5ed, rgba(170, 142, 245, 0)),
        radial-gradient(closest-side, #8295b5, rgba(248, 192, 147, 0));
    background-size: 130vmax 130vmax, 80vmax 80vmax, 90vmax 90vmax, 110vmax 110vmax, 90vmax 90vmax;
    background-position: -80vmax -80vmax, 60vmax -30vmax, 10vmax 10vmax, -30vmax -10vmax, 50vmax 50vmax;
    animation: 10s move linear infinite;

}
// 通过修改background的参数形成动画
@keyframes move {

    0%,
    100% {
        background-size: 130vmax 130vmax, 80vmax 80vmax, 90vmax 90vmax, 110vmax 110vmax, 90vmax 90vmax;
        background-position: -80vmax -80vmax, 60vmax -30vmax, 10vmax 10vmax, -30vmax -10vmax, 50vmax 50vmax;
    }

    25% {
        background-size: 100vmax 100vmax, 90vmax 90vmax, 100vmax 100vmax, 90vmax 90vmax, 60vmax 60vmax;
        background-position: -60vmax -90vmax, 50vmax -40vmax, 0vmax -20vmax, -40vmax -20vmax, 40vmax 60vmax;
    }

    50% {
        background-size: 80vmax 80vmax, 110vmax 110vmax, 80vmax 80vmax, 60vmax 60vmax, 80vmax 80vmax;
        background-position: -50vmax -70vmax, 40vmax -30vmax, 10vmax 0vmax, 20vmax 10vmax, 30vmax 70vmax;
    }

    75% {
        background-size: 90vmax 90vmax, 90vmax 90vmax, 100vmax 100vmax, 90vmax 90vmax, 70vmax 70vmax;
        background-position: -50vmax -40vmax, 50vmax -30vmax, 20vmax 0vmax, -10vmax 10vmax, 40vmax 60vmax;
    }
}


/* 加载圆圈的样式 */
@keyframes spin {
    0% {
        transform: rotate(0deg);
    }
    100% {
        transform: rotate(360deg);
    }
}
.loader {
    
    border: 16px solid #f3f3f3;
    border-radius: 50%;
    border-top: 16px solid #7c80b8;
    width: 80px;
    height: 80px;
    position: relative;
    top: 50%;
    left: 50%;
    z-index: 9999;
    margin-top: 50px;
    margin-left: -50px;
    display: none;
    animation: spin 2s linear infinite;
}

.loader.show {
    display: block;
}

ul {
    list-style: none;
    margin: 0;
    padding: 0;
}

.wrap {
    width: 60%;
    border-radius: 10px;
    margin: 0 auto;
    transform: translateY(-20px);
    padding: 100px 0;
    h4 {
        text-align: center;
        margin: 12px 0;
        color: #34495e;
        font-style: italic;
    }

    .questionnaire {
        margin: 30px 0;
        background-color: #eff0f6;
        border-radius: 5px;
        padding: 20px;
        box-shadow: 0px 2px 5px 0px rgba(0, 0, 0, 0.15);
        h2 {
            font-size: 24px;
            margin-bottom: 20px;
        }

        ul {
            padding: 0;

            li {
                display: flex;
                align-items: center;
                margin-bottom: 10px;

                input[type="radio"] {
                    -webkit-appearance:none;
                    -moz-appearance:none;
                    outline: none;
                    width:25px;
                    height:25px;
                    background-color:#ffffff;
                    border:solid 1px #dddddd;
                    border-radius:50%;
                    margin:0px 10px 4px 4px;
                    padding:0;
                    position:relative;
                    display:inline-block;
                    vertical-align:top;
                    transition:background-color ease 0.1s;
                }
                input[type="radio"]:checked{
                    background:#7c80b8;
                }
                input[type="radio"]:checked::after{
                    content:'';
                    top:8px;
                    left:6px;
                    position:absolute;
                    border:#fff solid 2px;
                    border-top:none;
                    border-right:none;
                    height:6px;
                    width:10px;
                    transform:rotate(-45deg);
                }

                label {
                    font-size: 16px;
                }
            }
        }
    }

    button {
        display: block;
        margin: 20px auto;
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        background-color: #eff0f6;
        color: #7c80b8;
        font-size: 18px;
        cursor: pointer;
        transition: background-color 0.3s;
        &:hover {
            background-color: #484b6f;
            color: #eff0f6;
        }
    }

    .answer{
        margin-top: 40px;
        text-align: center;
        h3{
            font-size: 20px;
            color: #eff0f6;
        }
        h1{
            font-size: 20px;
            color: #eff0f6;
        }
    }

}


</style>
