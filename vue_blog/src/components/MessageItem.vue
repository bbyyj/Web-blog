<template>
    <div>
        <div class="wapper">
            <div class="toprow">
                <div class="info">
                </div>
            </div>
        </div>
        <div class="question">
            {{ msg.question }}
        </div>
        <div class="answer">
            {{ msg.answer }}
        </div>
        <h6 class="date"> {{ dateFormat() }} <a @click.prevent="replay" v-if="msg.answer == ''">回复</a></h6>
    </div>
</template>
<script>

import dayjs from "dayjs";
import axios from 'axios';

export default {
    name: "MessageItem",
    props: ["msg"],
    computed: {
        imgObj() {
            const x = Math.floor((this.msg.avatar - 1) % 7) + 1
            const y = Math.floor((this.msg.avatar - 1) / 7)
            return {
                backgroundPositionX: x * 48 + 'px',
                backgroundPositionY: y * 48 + 'px',
            }
        },
    },
    methods: {
        dateFormat() {
            return dayjs(this.msg.createTime).format('YYYY-MM-DD HH:mm:ss')
        },
        replay() {
            //控制是否显示"回复按钮" 并且在点击时跳转到后台管理页面
            this.$router.push({
                path: "/login",
            });
        }
    },

}
</script>

<style scoped>
.wapper {
    overflow: hidden;
}

.toprow {
    height: 0px;
    overflow: hidden;
    line-height: 48px;
}


.avatar {
    width: 48px;
    height: 48px;
    float: left;
    border-radius: 50%;
    transition: all 0.6s linear;
}

.avatar:hover {
    transform: rotate(360deg);

}

.info {
    float: left;
    margin-left: 12px;
}


h4 {
    font-size: 13px;
    font-weight: 700;
    margin-top: 8px;
    margin-bottom: 5px;
    cursor: pointer;
}

h5 {
    font-size: 12px;
    font-weight: 400;
    margin: 0;
}

a {
    float: right;
    margin-right: 10px;
    outline: none;
    text-decoration: none;
    color: rgb(88 80 236) !important;
    cursor: pointer;
}

a:hover {
    color: rgb(166, 80, 236) !important;
    ;
}

p {
    display: block;
    margin: 5px 54px 10px;
    line-height: 1.8;
}

.space {
    height: 20px;
}

.question {
    /*问题的颜色 */
    background: url("../assets/images/askboard_back.jpg");
    padding: 10px;
    margin: 10px 0;
    border-radius: 5px;
    font-weight: normal;
    font-family: 华文楷体;
}

.answer {
    /*答案的样式 */
    padding: 10px;
    margin: 10px 0;
    font-weight: bolder;
    font-family: 华文楷体;
}
</style>