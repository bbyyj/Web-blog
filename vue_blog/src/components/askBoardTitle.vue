<template>
    <div>
        <div class="bg">
            <span class="title">{{ info.title }}</span>
            <span class="desc">{{ info.desc }}</span>
            <div class="space"> </div>
            <div>
                <button @click="show_ask" class="ask_button">点击此处向我匿名提问</button>
            </div>
            <div class="space" v-if="control"> </div>
            <div class="ask_board" v-if="control">
                <LeaveMessagePanel :closeBtn="false" class="leave-message"></LeaveMessagePanel>
                <transition name="replay">
                    <LeaveMessagePanel :closeBtn="true" :msgInfo="msgInfo" v-show="showReplay" class="replay-message">
                    </LeaveMessagePanel>
                </transition>
            </div>
            <div class="space"> </div>
        </div>
    </div>
</template>

<script>
import LeaveMessagePanel from "../components/LeaveMessagePanel";
export default {
    name: "askBoardTitle",
    props: ["info"],
    components: { LeaveMessagePanel },
    data() {
        return {
            control: false
        }
    },
    methods: {
        show_ask() {
            if (this.control == false)
                this.control = true
            else
                this.control = false
        }
    }
}
</script>

<style scoped>

/* 提问箱背景部分 */
.bg {
    user-select: none;
    text-align: center;
    font-family: GEETYPEQingKongPaoPaoTi-Shan-XiTi;

    font-size: 50px;
    position: relative;
    z-index: 0;
}
.bg::before {
    content: '';
    background-color: #fffafa;
    opacity: 0.75;
    border-radius: 20px;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1;
}

/* 提问箱文字部分 */
.title {
    font-size: 45px;
    font-weight: 500;
    color: #7378ac;
    display: block;
    padding-top: 120px;
}

.desc {
    font-size: 30px;
    color: #7378ac;
    margin-top: 8px;
    display: inline-block;
}

/* 控制必要上下间隔 */
.space {
    height: 40px;
}

/* 提问箱按钮部分 */
button {
    background: #c1c3da;
    color: #ffffff;
    font-size: 25px;
    font-weight: bold;
    border: 0;
    border-radius: 10px;
    padding: 15px 30px;
    outline: none;
    position: relative;
}
button:active {
    transform: translateY(2px);
}
/* 提问箱问题输入框部分 */
.leave-message {
    margin: 0 auto;
    opacity: 0.8;
}

.replay-message {
    position: relative;
    left: 50%;
    transform: translateX(-50%);
}

</style>
