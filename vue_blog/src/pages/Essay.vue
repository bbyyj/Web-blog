
<template>
    <div class="love_test">
        <TitleArea class="title-area" :info="info"></TitleArea>
        <div class="bottom">
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
                        <div v-if="showScore">
                            <h3>你的得分是：{{ score }}</h3>
                            <h1>{{ messages[score] }}</h1>
                        </div>
                    </div>
                </div>
            </transition>
        </div>
    </div>
</template>

<script>

import "../assets/icon/iconfont.css"
import TitleArea from "../components/TitleArea";

import "animate.css"

export default {
    components: { TitleArea },
    data() {
        return {
            info: {
                title: "默契鉴定",
                desc: "!"
            },
            questions: [
                //设置的10个问题
                {
                    text: "问题1",
                    answers: ["选项1", "选项2", "选项3"],
                    correctAnswer: "选项1"
                },
                {
                    text: "问题2",
                    answers: ["选项1", "选项2", "选项3"],
                    correctAnswer: "选项2"
                },
                // 可以添加更多的问题
            ],
            //将问题选项的按钮置位空
            userAnswers: Array(2).fill(null),
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
                0: '我们真的没有缘分！',
                10: '可能我们只是初识，需要更多的了解',
                20: '默契有点微妙，但是还有希望',
                30: '我们的默契还需提高！',
                40: '虽然有些小默契，但还有很大的提升空间',
                50: '我们的默契程度一般，需要更多交流',
                60: '我们的默契程度还不错，继续保持',
                70: '我们的默契程度很好，让我们更进一步',
                80: '我们的默契程度很高，继续努力',
                90: '我们的默契程度接近完美，让我们继续',
                100: '我们之间的默契棒极了！'
            },
        }
    },
    created() {
        window.scrollTo(0, 0);
        this.getUserInfo()
    },
    methods: {

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
            this.userAnswers = Array(2).fill(null);
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
    z-index: 9999;
    border: 16px solid #f3f3f3;
    border-radius: 50%;
    border-top: 16px solid #3498db;
    width: 120px;
    height: 120px;
    animation: spin 2s linear infinite;
    position: absolute;
    top: 50%;
    left: 50%;
    margin-top: -60px;
    margin-left: -60px;
    display: none;
}

.loader.show {
    display: block;
}

ul {
    list-style: none;
    margin: 0;
    padding: 0;
}

.love_test {
    background-color: #3A3B55;
    background-image:
        radial-gradient(closest-side, #7378ac, rgba(80, 120, 99, 0)),
        radial-gradient(closest-side, #a69ec6, rgba(132, 157, 133, 0)),
        radial-gradient(closest-side, #6c5c74, rgba(94, 129, 137, 0)),
        radial-gradient(closest-side, #fbd5d1, rgba(210, 201, 154, 0)),
        radial-gradient(closest-side, #3d3952, rgba(114, 145, 147, 0));
    background-size:
        130vmax 130vmax,
        80vmax 80vmax,
        90vmax 90vmax,
        110vmax 110vmax,
        90vmax 90vmax;
    background-position:
        -80vmax -80vmax,
        60vmax -30vmax,
        10vmax 10vmax,
        -30vmax -10vmax,
        50vmax 50vmax;
    background-repeat: no-repeat;
    animation: 10s movement linear infinite;

    @keyframes movement {

        0%,
        100% {
            background-size:
                130vmax 130vmax,
                80vmax 80vmax,
                90vmax 90vmax,
                110vmax 110vmax,
                90vmax 90vmax;
            background-position:
                -80vmax -80vmax,
                60vmax -30vmax,
                10vmax 10vmax,
                -30vmax -10vmax,
                50vmax 50vmax;
        }

        25% {
            background-size:
                100vmax 100vmax,
                90vmax 90vmax,
                100vmax 100vmax,
                90vmax 90vmax,
                60vmax 60vmax;
            background-position:
                -60vmax -90vmax,
                50vmax -40vmax,
                0vmax -20vmax,
                -40vmax -20vmax,
                40vmax 60vmax;
        }

        50% {
            background-size:
                80vmax 80vmax,
                110vmax 110vmax,
                80vmax 80vmax,
                60vmax 60vmax,
                80vmax 80vmax;
            background-position:
                -50vmax -70vmax,
                40vmax -30vmax,
                10vmax 0vmax,
                20vmax 10vmax,
                30vmax 70vmax;
        }

        75% {
            background-size:
                90vmax 90vmax,
                90vmax 90vmax,
                100vmax 100vmax,
                90vmax 90vmax,
                70vmax 70vmax;
            background-position:
                -50vmax -40vmax,
                50vmax -30vmax,
                20vmax 0vmax,
                -10vmax 10vmax,
                40vmax 60vmax;
        }
    }
}

.title-area {
    font-size: 450%;
    color: #ffffff;
    margin-bottom: 50px;
    bottom: 0 !important;
    right: 0 !important;
    font-family: 'STXingkai';
    opacity: 0.5;
    padding-top: 6%;
}

.bottom {
    padding-bottom: 80px;
    background-color: #ececec;
}

.wrap {
    width: 60%;
    border-radius: 10px;
    margin: 0 auto;
    transform: translateY(-20px);
    box-shadow: 0 15px 35px rgb(50 50 93 / 10%), 0 5px 15px rgb(0 0 0 / 7%);
    background: #fff url("../assets/images/flowerbg.png") fixed top;
    padding: 50px 0;

    img {
        display: block;
        width: 160px;
        height: 160px;
        border-radius: 50%;
        margin: 0 auto;
    }

    h4 {
        text-align: center;
        margin: 12px 0;
        color: #34495e;
        font-style: italic;
    }

    .questionnaire {
        margin: 20px 0;
        background-color: #f8f8f8;
        border-radius: 5px;
        padding: 20px;
        box-shadow: 0px 2px 5px 0px rgba(0, 0, 0, 0.15);

        h2 {
            font-size: 18px;
            margin-bottom: 10px;
        }

        ul {
            padding: 0;

            li {
                display: flex;
                align-items: center;
                margin-bottom: 10px;

                input[type="radio"] {
                    margin-right: 10px;
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
        background-color: #42b983;
        color: white;
        font-size: 18px;
        cursor: pointer;
        transition: background-color 0.3s;

        &:hover {
            background-color: #2a8a63;
        }
    }

}
</style>
