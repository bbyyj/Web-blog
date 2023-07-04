<template>
    <div>
        <!-- 动态背景设置 -->
        <div class="background-layer"></div>
        <div class="background-layer upper01"></div>
        <div class="background-layer upper02"></div>

        <!-- 提问箱内容设置 -->
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
                        <!--注意对应的渲染关系-->
                        <GroupMessageItem :key="item.id" v-for="item in messages" :groupMsg="item">
                        </GroupMessageItem>
                    </div>
                </div>
            </transition>
            <!-- <Pagination @jumpPage="jumpPage" :pageInfo="{ pageNum: queryInfo.pageNum, pages: pages }"></Pagination> -->
        </div>
    </div>
</template>

<script>


import askBoardTitle from "../components/askBoardTitle";
import LeaveMessagePanel from "../components/LeaveMessagePanel";
import GroupMessageItem from "../components/GroupMessageItem";

export default {
    name: "MsgBoard",
    components: {  LeaveMessagePanel, GroupMessageItem, askBoardTitle },
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
            messages: [],
            msgInfo: {
                parentId: 0,
                topParentId: 0
            }
        }
    },

    methods: {
        //获取对应的
        async getMessageList() {
            try {
                const { data: res } = await this.$axios.get("/myblog/getAnsweredQA");
                if (res.status !== 1) {
                    this.$message.error("网络出了点小问题,获取提问失败！")
                    return
                }
                console.log("test")
                this.total = res.data[1];
                this.messages = res.data[0].map(item => ({
                    myself: item.find(m => m.is_parent === true),
                    children: item.filter(m => m.is_parent === false)
                }));
            } catch (error) {
                console.error(error);
                this.$message.error("网络出了点小问题,获取提问失败！")
            }
        },
        jumpPage(pageNum) {
            this.queryInfo.pageNum = pageNum;
            this.getMessageList();
        }
    },
    created() {
        window.scrollTo(0, 0)
        this.getMessageList()
    },
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
    font-family: GEETYPEQingKongPaoPaoTi-Shan-XiTi;
}

/*回答过的问题 */
.total {
    font-size: 32px;
    line-height: 40px;
    font-weight: 1000;
    text-align: center;
    padding-top: 30px;
    padding-bottom: 10px;
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
