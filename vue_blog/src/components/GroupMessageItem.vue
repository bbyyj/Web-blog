<!--一组评论，包含父评论和子评论-->
<template>
    <div class="group-warp">
        <MessageItem :msg="groupMsg.myself"></MessageItem>
        <div class="count" v-if="groupMsg.children.length">
            <span @click="hideChildren">
                {{ groupMsg.children.length }}条回复<span class="arrowUp" :class="{ arrowDown: !showChildren }"></span>
            </span>
        </div>
        <transition name="children">
            <div class="ul-wapper" v-show="showChildren">
                <ul>
                    <li :key="item.id" v-for="item in groupMsg.children">
                        <MessageItem :msg="item"></MessageItem>
                    </li>
                </ul>
            </div>
        </transition>
    </div>
</template>

<script>

import MessageItem from "./MessageItem";

export default {
    name: "GroupMessageItem",
    components: { MessageItem },
    props: ["groupMsg"],
    data() {
        return {
            showChildren: true
        }
    },
    methods: {
        hideChildren() {
            this.showChildren = !this.showChildren
        }
    }
}
</script>

<style scoped>
ul,
li {
    margin: 10px;
    padding: 0px;
}

.group-warp {
    background-color: #c1c3da;
    border: 1px solid #c1c3da;
    box-shadow: 0 15px 35px rgb(50 50 93 / 18%), 0 5px 15px rgba(16, 10, 10, 0.18) !important;

    border-radius: 20px; 
    margin: 40px;
    padding: 20px 30px 3px;
}

.count {
    overflow: hidden;
    padding-bottom: 10px;
    user-select: none;
}

.count>span {
    float: right;
    margin-right: 50px;
    cursor: pointer;
    font-weight: 600;
    font-size: 20px;
    line-height: 1.5;
    position: relative;
}

.arrowUp {
    display: inline-block;
    width: 8px;
    height: 8px;
    /* border-right: 1px solid #000;
    border-bottom: 1px solid #000; */
    /* transform: rotate(-135deg); */
    position: absolute;
    top: 10px;
    right: -20px;
    transition: all 0.3s linear;
}

.arrowDown {
    transform: rotate(45deg);
    top: 5px;
}

.ul-wapper {
    position: relative;
    transition: all 0.5s ease-in;
}

ul::before {
    content: '';
    width: 0;
    height: 80%;
    border-left: 1px solid rgba(0, 0, 0, .3);
    position: absolute;
    left: 20px;
    top: 0;
}

li {
    list-style: none;
    margin-left: 50px !important;
}

@keyframes appe {
    from {
        transform: scale(0);
    }

    to {
        transform: scale(1);
    }
}

.children-enter-active {
    animation: appe 0.5s ease;
}

.children-leave-active {
    animation: appe 0.5s ease;
    animation-direction: reverse;
}
</style>