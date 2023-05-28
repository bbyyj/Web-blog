<!--一组评论，包含父评论和子评论-->
<template>
    <div>
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
                        <li :key="item.id" v-for="(item, index) in groupMsg.children">
                            <MessageItem :msg="item"></MessageItem>
                            <!-- 追加子问题的按钮只会在最后一个问题下面显示 -->
                            <button v-if="index === groupMsg.children.length - 1" @click="openAppendDialog">追加子问题</button>
                        </li>
                    </ul>
                </div>
            </transition>
        </div>
        <!-- 追加子问题的对话框 -->
        <div v-if="showAppendDialog" class="append-dialog">
            <div>
                <label>
                    问题内容：
                    <input v-model="newQuestion" type="text">
                </label>
            </div>
            <div>
                <label>
                    是否为彩虹：
                    <input v-model="newQuestionRainbow" type="checkbox">
                </label>
            </div>
            <button @click="submitAppendDialog(groupMsg.children.length)">提交</button>
            <button @click="closeAppendDialog">关闭</button>
        </div>
        <!-- 阴影效果 -->
        <div v-if="showAppendDialog" class="overlay" @click="closeAppendDialog"></div>
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
            showChildren: true,
            showAppendDialog: false,
            newQuestion: '',
            newQuestionRainbow: false,
        }
    },
    methods: {
        openAppendDialog() {
            this.showAppendDialog = true;
        },
        closeAppendDialog() {
            this.showAppendDialog = false;
            this.newQuestion = '';
            this.newQuestionRainbow = false;
        },
        submitAppendDialog(child_id) {
            this.appendOldQuestion(child_id, this.newQuestion, this.newQuestionRainbow);
            this.closeAppendDialog();
        },
        appendOldQuestion(child_id, question, rainbow) {
            this.openAppendDialog();
            const parent_id = this.groupMsg.myself.id;
            axios.post('http://127.0.0.1:8080/api/myblog/appendOldQuestion', {
                parent_id,
                child_id,
                question,
                rainbow,
            }).then(response => {
                if (response.data.status === 101) {
                    this.groupMsg.children.push({
                        id: child_id,
                        content: question,
                        rainbow: rainbow,
                        likes: 0,
                        reply: '',
                        replyTime: new Date().toLocaleString(),
                        askedTime: new Date().toLocaleString(),
                        answered: false,
                        isParent: false,
                    });
                } else {
                    console.error('追加子问题失败: ', response.data.message);
                }
            }).catch(error => {
                console.error('追加子问题时发生错误: ', error);
            });
        },
        clickLikes(child_id) {
            const parent_id = this.groupMsg.myself.id;
            const likes = this.groupMsg.children.find(child => child.id === child_id).likes;
            axios.put('http://127.0.0.1:8080/api/myblog/clickLikes', {
                parent_id,
                child_id,
                likes,
            }).then(response => {
                if (response.data.status === 101) {
                    this.groupMsg.children.find(child => child.id === child_id).likes += 1;
                } else {
                    console.error('点赞问题失败: ', response.data.message);
                }
            }).catch(error => {
                console.error('点赞问题时发生错误: ', error);
            });
        },
        hideChildren() {
            this.showChildren = !this.showChildren
        }
    }
}
</script>

<style scoped>
.append-dialog {
    position: fixed;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    background-color: #fff;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
    z-index: 1000;
}

.overlay {
    /*可以实现点击对话框的周围也退出对话框*/
    position: fixed;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 999;
}

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