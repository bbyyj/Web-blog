<!--一组评论，包含父评论和子评论-->
<template>
    <div>
        <div class="group-warp">
            <MessageItem :msg="groupMsg.myself"></MessageItem>
            <button class="follow-btn" @click="clickLikes(groupMsg.myself)">
                {{ groupMsg.myself.likes }}赞
            </button>
            <button class="follow-btn" v-if="shouldDisplayButton" @click="openAppendDialog">追加子问题</button>
            <div class="count" v-if="groupMsg.children.length">
                <span @click="hideChildren">
                    {{ groupMsg.children.length }}条追问<span class="arrowUp" :class="{ arrowDown: !showChildren }"></span>
                </span>
            </div>
            <transition name="children">
                <div class="ul-wapper" v-show="showChildren">
                    <ul>
                        <li :key="item.id" v-for="(item, index) in groupMsg.children">
                            <MessageItem :msg="item"></MessageItem>
                            <button class="follow-btn" @click="clickLikes(item)">
                                {{ item.likes }}赞
                            </button>
                            <!-- 追加子问题的按钮只会在最后一个问题下面显示 -->
                            <button class="follow-btn" v-if="index === groupMsg.children.length - 1" @click="openAppendDialog">追加子问题</button>
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
                    <input class="rainbow-contain" v-model="newQuestion" type="text">
                </label>
            </div>
            <div>
                <label class="rainbow-label">
                    是否为彩虹：
                    <input v-model="newQuestionRainbow" class="rainbow-input" type="checkbox">
                </label>
            </div>
            <button class="ask-btn" @click="submitAppendDialog(groupMsg.children.length)">提交啦</button>
            <button class="ask-btn" @click="closeAppendDialog">关闭</button>
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
            showChildren: false,
            showAppendDialog: false,
            //新写的问题
            newQuestion: '',
            //新写的问题是否加彩虹
            newQuestionRainbow: false,
        }
    },
    computed: {
        shouldDisplayButton() {
            return this.groupMsg.children.length === 0;
        },

    },
    methods: {
        //显示追加问题的窗口
        openAppendDialog() {
            this.showAppendDialog = true;
        },
        //取消追加子问题
        closeAppendDialog() {
            this.showAppendDialog = false;
            this.newQuestion = '';
            this.newQuestionRainbow = false;
        },

        //提交新的子问题  点击确定按钮
        submitAppendDialog(child_id) {
            console.log(child_id)
            child_id += 1
            console.log()
            console.log(this.newQuestion)
            console.log(this.newQuestionRainbow)
            //提交子问题的操作   参数已经获取到了
            this.appendOldQuestion(child_id, this.newQuestion, this.newQuestionRainbow);
            this.closeAppendDialog();
        },
        //追加子问题的函数
        async appendOldQuestion(child_id, question, rainbow) {
            // const parent_id = this.groupMsg.myself.id;
            const parent_id = this.groupMsg.myself.parent_id;
            console.log(parent_id)
            try {
                const { data: res } = await this.$axios.post("/myblog/appendOldQuestion", {
                    parent_id,
                    child_id,
                    question,
                    rainbow,
                });
                if (res.status !== 101) {
                    this.$message.error("操作失败，请重试！")
                } else {
                    this.$message.success("操作成功！");
                    await this.getMessageList()
                    this.closeAppendDialog();
                }
            } catch (error) {
                console.error(error);
            }
        },
        //点赞相关的函数
        async clickLikes(temple) {
            const child_id = temple.child_id
            // const parent_id = this.groupMsg.myself.id;
            const parent_id = this.groupMsg.myself.parent_id;
            const likes = temple.likes += 1;
            console.log(temple.likes)
            console.log(likes)
            console.log(child_id)
            console.log(parent_id)
            try {
                const { data: res } = await this.$axios.put("/myblog/clickLikes", {
                    likes,
                    parent_id,
                    child_id,
                });
                if (res.status !== 101) {
                    this.$message.error("点赞失败，请重试！")
                } else {
                    this.$message.success("点赞成功！");
                }
            } catch (error) {
                console.error(error);
            }

        },
        //切换  是否显示追问
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
    padding: 40px;
    border-radius: 10px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
    z-index: 1000;
}
.ask-btn{
    margin-top: 10px;
    margin-right: 15px;
    padding: 3px 8px 3px 8px ;
    border: 2px solid #9990bc;
    border-radius: 5px;
}
.rainbow-label{
    margin-top: 18px;
}
.rainbow-input{
    -webkit-appearance:none;
    -moz-appearance:none;
    outline: none;
    width:23px;
    height:23px;
    background-color:#fdeceb;
    border:solid 2px #fcdedb;
    border-radius:50%;
    margin:0px 0px 1px 2px;
    padding:0;
    position:relative;
    display:inline-block;
    vertical-align:top;
    transition:background-color ease 0.1s;
}
.rainbow-input:checked {
    background:#7c80b8;
    border:solid 2px #3d3952;
}
.rainbow-input:checked::after {
    content:'';
    top:6px;
    left:5px;
    position:absolute;
    border:#fff solid 2px;
    border-top:none;
    border-right:none;
    height:6px;
    width:10px;
    transform:rotate(-45deg);
}
.rainbow-contain{
    border: 2px solid #9990bc;
    border-radius: 5px;
}
.rainbow-contain:focus{
    outline: 2px solid #3d3952;
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
.follow-btn{
    margin-bottom: 10px;
    margin-top: 10px;
    margin-right: 15px;
    padding: 8px 15px 8px 15px ;
    border: 2px solid #3d3952;
    background-color: #3d3952;
    color: #fef6f5; 
    border-radius: 5px;
}

.follow-btn:hover{
    background-color: #fef6f5;
    color: #3d3952; 
}
ul,li {
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
    
    font-family: GEETYPEQingKongPaoPaoTi-Shan-XiTi;
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
    font-size: 20px;
    line-height: 1.5;
    position: relative;

    display: block;
    border: 2px solid #3d3952;
    background-color: #fef6f5;
    padding: 8px 15px 8px 15px ;
    border-radius: 5px;
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