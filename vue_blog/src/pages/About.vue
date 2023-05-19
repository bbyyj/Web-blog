
<template>
    <div>
        <TitleArea class="title-area" :info="info"></TitleArea>
        <div class="bottom">
            <transition appear name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__slideInUp">
                <div class="wrap">



                    <img :src="userInfo.avatar">
                    <h4>{{ userInfo.nickname }}</h4>

                    <ul class="contact">
                        <li><a :href="userInfo.github" target="_blank" class="iconfont icon-github"></a></li>
                        <li><a :href="userInfo.csdn" target="_blank" class="iconfont icon-csdn"></a></li>
                    </ul>
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

                    <div class="box-wrap">
                        <span class="side-label">关于我</span>
                        <h5 class="title">关于博主</h5>
                        <p class="about-state">
                            自我描述 //可在后端数据库存储
                        </p>
                        <p class="about-state">
                            自我描述 //可在后端数据库存储
                        </p>
                    </div>

                    <div class="box-wrap">
                        <span class="side-label bgc-pink">关于本站</span>
                        <h5 class="title">本站采用的技术</h5>
                        <h6 class="site-intro"><span>前端：</span>Vue + Element Ui</h6>
                        <h6 class="site-intro"><span>后端：</span>Golang + Gin + Mysql</h6>
                        <h6 class="site-intro"><span>部署方式：</span>提供了Docker-Compose的部署方式,只需少量配置即可一键启动,使用Nginx作为静态资源和反向代理服务器。
                        </h6>
                        <h6 class="site-intro">本站已经开源在Github</h6>
                    </div>

                    <div class="box-wrap">
                        <span class="side-label bgc-orange">联系方式</span>
                        <h6 class="site-intro" style="margin-top: 12px;"><span>Github:</span><a :href="userInfo.github"
                                target="_blank"> {{ userInfo.github }}</a></h6>
                        <h6 class="site-intro"><span>CSDN:</span><a :href="userInfo.csdn" target="_blank"> </a></h6>
                        <h6 class="site-intro"><span>Email:</span>{{ userInfo.email }}</h6>
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
                title: "关于我",
                desc: "写点什么吧!"
            },
            userInfo: {
                avatar: "",
                nickname: "",
                github: "",
                csdn: "",
                email: ""
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
            //是否提交评分过  评分过按钮会变成“重新测试”
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
        async getUserInfo() {
            //获取用户信息
            const { data: res } = await this.$axios.get("/myblog/userInfo")
            if (res.status !== 1) {
                this.$message.warning("获取用户信息失败！")
            }
            if (res.data.length > 0) {
                const data = res.data[0]
                this.userInfo.avatar = data.avatar
                this.userInfo.nickname = data.nickname
                this.userInfo.github = data.github
                this.userInfo.csdn = data.csdn
                this.userInfo.email = data.email
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

.title-area {
    background: url("../assets/images/back12.jpg") 0 0 / cover no-repeat;
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

    .contact {
        text-align: center;
        margin: 20px 0;
        padding-bottom: 6px;
        position: relative;

        li {
            display: inline-block;

            a {
                text-decoration: none;
                font-size: 26px;
                margin: 0 12px;
                color: rgb(63, 63, 64);
            }
        }
    }


    .box-wrap {
        position: relative;
        width: 80%;
        margin: 30px auto 60px;
        border-radius: 8px;
        background-color: rgba(255, 255, 255, .5);
        box-shadow: -12px -17px 38px rgb(0 0 0 / 10%), 7px 18px 25px rgb(0 0 0 / 8%);
        transition: all .3s ease 0s, transform .6s cubic-bezier(.6, .2, .1, 1) 0s, -webkit-transform .6s cubic-bezier(.6, .2, .1, 1) 0s;
        color: #34495e;
        padding-top: 60px;
        padding-bottom: 18px;

        .side-label {
            position: absolute;
            background-color: #b28fce;
            color: #fff;
            font-style: italic;
            font-size: 14px;
            padding-left: 32px;
            padding-right: 24px;
            left: -16px;
            top: 20px;
            height: 32px;
            line-height: 32px;
            border-radius: 0 3px 3px 0;
        }

        .side-label::before {
            content: '';
            position: absolute;
            top: 100%;
            left: 0;
            width: 0;
            height: 0;
            background-color: transparent;
            border-style: solid;
            border-width: 0 16px 16px 0;
            border-color: transparent;
            border-right-color: #b28fce;
            filter: brightness(120%);
        }
    }

    .box-wrap:hover {
        transform: translateY(-5px);
    }

    .bgc-pink {
        background-color: #f596aa !important;
    }

    .bgc-pink::before {
        border-right-color: #f596aa !important;
    }

    .bgc-orange {
        background-color: #fb966e !important;
    }

    .bgc-orange::before {
        border-right-color: #fb966e !important;
    }

    .title {
        padding: 10px 20px;
        font-size: 20px;
        font-weight: 600;

    }

    font-family: 'Architects Daughter',
    cursive;

    .about-state {
        padding: 0 60px 0 36px;
        text-indent: 2rem;
        font-style: italic;
        font-size: 18px;
        font-weight: 520;
        margin-bottom: 8px;
    }

    .site-intro {
        padding-left: 32px;
        padding-right: 60px;
        font-size: 18px;

        span {
            font-weight: 700;
        }

        a {
            text-decoration: none;
            color: #42b983;
        }
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
