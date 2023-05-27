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
    overflow: hidden;
    height: 0px;
    line-height: 48px;
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
p {
    display: block;
    margin: 5px 54px 10px;
    line-height: 1.8;
}

a {
    display: block;
    background-color: #30246b;
    color: rgb(255, 255, 255) !important;
    border-radius: 10px;
    float: right;
    padding: 10px;
    margin-right: 10px;
    outline: none;
    text-decoration: none;   
    cursor: pointer;
}

a:hover {
    color: #30246b !important;
    background-color: #ffffff;
    opacity: 0.75;
}

.space {
    height: 20px;
}
/* 问题部分 */
.question {
    padding: 10px;
    margin: 0px 0;
    border-radius: 5px;
    font-size: 24px;
    font-weight: normal;
    font-family: 华文楷体;
}
/* 答案部分 */
.answer {
    padding: 10px;
    margin: 10px 0;
    font-size: 24px;
    font-weight: bolder;
    font-family: 华文楷体;
}
</style>