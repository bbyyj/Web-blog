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
.bg {
    width: 100%;
    position: relative;
    text-align: center;
    z-index: 0;
    user-select: none;
    font-family: '楷体';
}

.bg::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1;
    background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAQAAAAECAYAAACp8Z5+AAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAKUlEQVQImU3IMREAIAgAwJfNkQCEsH8cijjpMf6vnXlQaIiJFx+omEBfmqIEZLe2jzcAAAAASUVORK5CYII=);
}

.title {
    font-size: 36px;
    font-weight: 500;
    color: #fff;
    display: block;
    padding-top: 160px;
}

.desc {
    display: inline-block;
    margin-top: 8px;
    color: #fff;
    font-size: 18px;
}

.space {
    height: 30px;
}

button {
    border: 0;
    border-radius: 10px;
    background: #2ec4b6;
    text-transform: uppercase;
    color: white;
    font-size: 16px;
    font-weight: bold;
    padding: 15px 30px;
    outline: none;
    position: relative;
    transition: border-radius 3s;
    -webkit-transition: border-radius 3s;
}

button:hover {
    border-bottom-right-radius: 50px;
    border-top-left-radius: 50px;
    border-bottom-left-radius: 10px;
    border-top-right-radius: 10px;
}

.leave-message {
    margin: 0 auto;
    opacity: 0.8;
}

.replay-message {
    position: relative;
    left: 50%;
    transform: translateX(-50%);
}

.replay-enter-active {
    animation: downcome 1s linear;
}

.replay-leave-active {
    animation: downcome 1s linear;
    animation-direction: reverse;
}
</style>
