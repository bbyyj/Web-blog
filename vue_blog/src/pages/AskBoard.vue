<template>
    <div>
        <!-- 动态背景设置 -->
        <div class="background-layer"></div>
        <div class="background-layer upper01"></div>
        <div class="background-layer upper02"></div>

        <!-- 提问想内容设置 -->
        <askBoardTitle class="title-area" :info="info"></askBoardTitle>
        <div class="main">
            <transition appear name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__slideInLeft">
                <div class="center-area">
                    <div class="total">
                        <span>回答过的提问({{ total }})</span>
                        <div class="space"></div>
                    </div>
                    <div class="comments">
                        <GroupMessageItem :key="item.myself.id" v-for="item in messages" :groupMsg="item">
                        </GroupMessageItem>
                    </div>
                </div>
            </transition>
            <Pagination @jumpPage="jumpPage" :pageInfo="{ pageNum: queryInfo.pageNum, pages: pages }"></Pagination>
        </div>
    </div>
</template>

<script>

import TitleArea from "../components/TitleArea";
import askBoardTitle from "../components/askBoardTitle";
import LeaveMessagePanel from "../components/LeaveMessagePanel";
import GroupMessageItem from "../components/GroupMessageItem";
import Pagination from "../components/Pagination";

export default {
    name: "MsgBoard",
    components: { TitleArea, LeaveMessagePanel, GroupMessageItem, Pagination, askBoardTitle },
    data() {
        return {
            showReplay: false,
            info: {
                title: "欢迎向我匿名提问!",
                desc: "Ask me anything!"
            },
            queryInfo: {
                pageNum: 1,
                pageSize: 10,
            },
            pages: 1,
            total: 0,
            messages: [
                {
                    myself: {
                        id: 1,
                        question: "为什么要学习数分为什么要学习数分为什么要学习数分为什么要学习数分为什么要学习数分为什么要学习数分为什么要学习数分为什么要学习数分为什么要学习数分为什么要学习数分为什么要学习数分",
                        answer: "",
                        createTime: "2022-03-29T13:28:02+08:00",
                        parentId: 0,
                        topParentId: 0,
                    },
                    children: []
                },
                {
                    myself: {
                        id: 2,
                        question: "为什么要学习数分",
                        answer: "我哪知道。",
                        createTime: "2022-04-29T13:28:02+08:00",
                        parentId: 0,
                        topParentId: 0,
                    },
                    children: []
                },
                {
                    myself: {
                        id: 3,
                        question: "为什么要学习数分",
                        answer: "你也不知道？",
                        createTime: "2022-05-29T13:28:02+08:00",
                        parentId: 0,
                        topParentId: 0,
                    },
                    children: []
                },
            ],
            msgInfo: {
                parentId: 0,
                topParentId: 0
            }
        }
    },
    methods: {
        async getMessageList() {
            // 从缓存中拿数据
            const msg = window.sessionStorage.getItem(`message_page${this.queryInfo.pageNum}`);
            if (msg !== null) {
                this.messages = JSON.parse(msg)
                return
            }

            const { data: res } = await this.$axios.get("/myblog/displayMsg", { params: this.queryInfo })
            if (res.status !== 1) {
                this.$message.error("网络出了点小问题，获取评论失败！")
                return
            }

            this.messages.splice(0, this.messages.length)
            this.total = res.data.length > 2 ? res.data[2] : this.total
            const count = res.data.length > 1 ? res.data[1] : 0
            const data = res.data.length > 0 ? res.data[0] : []

            this.pages = Math.ceil(count / this.queryInfo.pageSize)
            if (this.pages <= 0) {
                this.pages = 1
            }

            // 查找第一个子评论的index，之前的都是父评论
            const index = data.findIndex(val => val.parentId !== 0)
            // 如果index为-1，说明没有子评论
            if (index === -1) {
                data.forEach((val) => {
                    this.messages.push({
                        myself: val,
                        children: []
                    })
                })
            } else {  // 有子评论，过滤子评论
                const parentMsg = data.splice(0, index)
                parentMsg.forEach((val) => {
                    const children = data.filter(v => val.id === v.topParentId)
                    // 查找子评论的子评论，给其添加父评论昵称
                    children.forEach((item) => {
                        if (item.parentId !== item.topParentId) {
                            item.parentName = (children.find(m => m.id === item.parentId)).nickname
                        }
                    })
                    this.messages.push({
                        myself: val,
                        children: children
                    })
                })
            }

            window.sessionStorage.setItem(`message_page${this.queryInfo.pageNum}`, JSON.stringify(this.messages))
        },
        jumpPage(pageNum) {
            this.queryInfo.pageNum = pageNum;
            this.getMessageList();
        }
    },
    created() {
        window.scrollTo(0, 0)
        // 清空缓存的提问
        window.sessionStorage.clear()
        this.getMessageList()
    },
    mounted() {
        this.$bus.$on("replay", (val) => {
            this.msgInfo.parentId = val.id
            this.msgInfo.topParentId = val.topParentId
            this.showReplay = true
        })
        this.$bus.$on("hideReplay", () => {
            this.showReplay = false
        })
    },
    beforeDestroy() {
        this.$bus.$off("replay")
        this.$bus.$off("hideReplay")
    }
}
</script>

<style scoped>

.background-layer {
    animation-name: slide;
    animation-duration: 4s;
    animation-timing-function: ease-in-out;
    animation-direction: alternate;
    animation-iteration-count: infinite;

    background-image: linear-gradient(-60deg, #a69ec6 50%, #fbdedb 50%);
    opacity: .5;
    position: fixed;

    left: -50%;
    right: -50%;
    bottom: 0;
    top: 0;
    z-index: -1;
}
@keyframes slide {
  0% {
    transform: translateX(-25%);
  }
  100% {
    transform: translateX(25%);
  }
}

.upper01 {
    /* 设置设置动画持续时间错开,形成波浪效果 */   
    animation-direction: alternate-reverse;
    animation-duration: 5s;
}

.upper02 {
    animation-duration: 6s;
}

.title-area {
    margin: 100px;
    padding-top: 6%;
}

.main {
    width: 100%;
    padding-top: 60px;
    padding-bottom: 80px;
}

.space {
    height: 30px;
}

.center-area {
    margin: 100px;
    background-color: #fffafa;
    opacity: 0.75;
    border-radius: 20px;
    padding-bottom: 20px;
    padding-top: 1px;
}

/*回答过的问题 */
.total {
    font-size: 32px;
    line-height: 32px;
    font-weight: 1000;
    text-align: center;
    padding-top: 30px;
    padding-bottom: 10px;
    font-family: 华文楷体;
    position: relative;
}

.total span {
    position: absolute;
    left: 20%;
    transform: translateX(-50%);
}

.total a {
    float: right;
    margin-right: 80px;
    text-decoration: none;
}
/* 
.leave-message {
    margin: 0 auto;
    opacity: 0.8;
} */
/* 
@keyframes downcome {
    from {
        bottom: -500px;
    }

    to {
        bottom: 0;
    }
}

.replay-message {
    position: fixed;
    left: 50%;
    transform: translateX(-50%);
    bottom: 0;

} */
/* 
.replay-enter-active {
    animation: downcome 1s linear;
}

.replay-leave-active {
    animation: downcome 1s linear;
    animation-direction: reverse;
} */

.comments {
    width: 950px;
    min-height: 80px;
    border-radius: 20px;
    background-color: #ffffff00;
    margin: 20px auto;
    opacity: 0.8;
}

</style>
