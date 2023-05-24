<template>
    <div class="panel" id="rp">
        <div class="wrapper">
            <textarea name="" id="" cols="30" rows="10" v-model="postInfo.content" placeholder="输入你的问题..."></textarea>
            <div class="btn-area">
                <button @click="clearAll">清空</button>
                <button @click="commit">发送</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "LeaveMessagePanel",
    props: ["closeBtn", "msgInfo"],
    watch: {
        msgInfo: {
            handler(newVal) {
                this.postInfo.parentId = newVal.parentId
                this.postInfo.topParentId = newVal.topParentId
            },
            deep: true
        }
    },
    data() {
        return {
            postInfo: {
                nickname: "",
                email: "",
                url: "",
                content: "",
                parentId: 0,
                topParentId: 0
            },
            commented: false,
            repeatCommitTimes: 0
        }
    },
    methods: {
        async commit() {
            if (!this.validate()) {
                return
            }

            this.trimSpace()
            const { data: res } = await this.$axios.post("/myblog/leaveMsg", this.postInfo);
            if (res.status === 101) {
                this.$message.success("您的问题已提交,请等待博主的回答！")
                this.postInfo.content = ""
                if (this.closeBtn) {
                    this.close()
                }
            } else {
                this.commented = false
                this.$message.success("网络出现了点小问题，请重试！")
            }
        },
        validate() {
            // 评论后，需要等一分钟才能评论，不能重复提交，重复提交次数>3，就清空所有输入的内容
            if (this.commented) {
                this.$message.error("您的提交太过频繁，请稍后再试！")
                this.repeatCommitTimes++
                if (this.repeatCommitTimes > 3) {
                    this.clearAll()
                }
                return false
            }
            /*
                        // 验证昵称
                        if (this.postInfo.nickname.trim() === "") {
                            this.$message.error("请输入您的昵称！")
                            return false
                        } else if (this.postInfo.nickname.length > 28) {
                            this.$message.error("昵称长度限制为28个字符！")
                            return
                        }
                        // 验证邮箱格式
                        const t = /^[A-Za-zd0-9]+([-_.][A-Za-zd]+)*@([A-Za-zd]+[-.])+[A-Za-zd]{2,5}$/;
                        if (!t.test(this.postInfo.email)) {
                            this.$message.error("邮箱格式错误！")
                            this.postInfo.email = ""
                            return false
                        }
                        // 验证网址格式
                        if (this.postInfo.url !== "") {
                            const reg = /(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?/;
                            if (!reg.test(this.postInfo.url)) {
                                this.$message.error("网址格式错误！")
                                this.postInfo.url = ""
                                return false
                            }
                        }
            */

            // 验证输入内容
            if (this.postInfo.content == "") {
                this.$message.error("请输入提问内容！")
                return false
            } else if (this.postInfo.content.length > 500) {
                this.$message.error("提问内容字符限制最大500字！")
                return false
            }
            this.commented = true
            this.repeatCommitTimes = 0
            setTimeout(() => {
                this.commented = false
            }, 60 * 1000)
            return true
        },
        clearAll() {
            this.postInfo = {
                nickname: "",
                email: "",
                url: "",
                content: ""
            }
        },
        trimSpace() {
            this.postInfo.nickname.trim()
            this.postInfo.email.trim()
            this.postInfo.content.trim()
            this.postInfo.url.trim()
        },
        /*
        close() {
            this.$bus.$emit("hideReplay")
        }*/
    }
}
</script>

<style scoped>
/* 问题框部分 */
.panel {
    border-radius: 20px;
    box-shadow: 0 15px 35px rgb(50 50 93 / 18%), 0 5px 15px rgb(0 0 0 / 18%);
    background-color: #eff0f6;
    width: 1050px;
}
.wrapper {
    padding: 50px 15px 15px;
}

textarea {
    font-size: 40px;
    outline: none;
    resize: none;
    border: none;
    border-bottom: 2px dashed #dedede;
    padding: 10px;
    width: 90%;
}

.btn-area {
    display: flex;
    justify-content: center;
    padding-top: 10px;
    padding-bottom: 10px;
}

button {
    flex: 1;
    background-color: #c1c3da;
    color: #ffffff;
    text-align: center;
    font-size: 25px;
    border: none;
    padding: 1rem 1rem;
    border-radius: 50px;
    margin: 0 150px;
    
}
button:active {
    transform: translateY(2px);
}
</style>