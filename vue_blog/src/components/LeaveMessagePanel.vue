<template>
    <div class="panel" id="rp">
        <div class="wrapper">
            <!-- 提问输入框 -->
            <textarea v-model="postInfo.question" placeholder="输入你的问题..." id="" cols="30" rows="10"></textarea>
            <!-- 彩虹模式选择框 -->
            <label class="rainbow">
                <input v-model="postInfo.rainbow" type="checkbox" class="rainbow-input">
                是否为彩虹
            </label>
            <!-- 提交按钮 -->
            <button @click="commit">发送</button>
        </div>
    </div>
</template>

<script>
export default {
    name: "LeaveMessagePanel",
    props: ["closeBtn", "msgInfo"],
    /*watch: {
        msgInfo: {
            handler(newVal) {
                this.postInfo.parentId = newVal.parentId
                this.postInfo.topParentId = newVal.topParentId
            },
            deep: true
        }
    },*/
    data() {
        return {
            postInfo: {
                question: "",  // 提问的内容
                rainbow: false,  // 是否为彩虹
                // 其他数据属性...
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
            // 添加新的问题
            this.addNewQuestion();
        },
        //验证提问
        validate() {
            // 提问后，需要等一分钟才能提问，不能重复提问，重复提问次数>3，就清空所有输入的内容
            if (this.commented) {
                this.$message.error("您的提交太过频繁，请稍后再试！")
                this.repeatCommitTimes++
                if (this.repeatCommitTimes > 3) {
                    this.clearAll()
                }
                return false
            }
            // 验证输入内容
            if (this.postInfo.question == "") {
                this.$message.error("请输入提问内容！")
                return false
            } else if (this.postInfo.question.length > 500) {
                this.$message.error("提问内容字符限制最大500字!")
                return false
            }
            //控制放置短时间多次提问的方式
            this.commented = true
            this.repeatCommitTimes = 0
            setTimeout(() => {
                this.commented = false
            }, 60 * 1000)
            return true
        },
        //添加新问题
        async addNewQuestion() {
            // 构建请求的数据
            const requestData = {
                question: this.postInfo.question,
                rainbow: this.postInfo.rainbow,
            };
            // 发送请求
            const { data: res } = await this.$axios.post("/myblog/addNewQuestion", requestData);
            // 检查响应的状态
            if (res.status !== 101) {
                this.$message.error("提交问题失败！");
                return;
            }
            // 提示成功消息
            this.$message.success("问题提交成功！");
            // 清空问题输入和彩虹标志
            this.postInfo.question = '';
            this.postInfo.rainbow = false;
        },

    }
}
</script>

<style lang="less" scoped>
/* 问题框部分 */
.panel {
    border-radius: 20px;
    box-shadow: 0 15px 35px rgb(50 50 93 / 18%), 0 5px 15px rgb(0 0 0 / 18%);
    background-color: #eff0f6;
    width: 1050px;
}

.wrapper {
    position: relative;
    padding: 50px 15px 15px;
}

textarea {
    font-size: 36px;
    outline: none;
    resize: none;
    border: none;
    border-bottom: 2px dashed #dedede;
    padding: 10px;
    width: 90%;
}

button {
    flex: 1;
    background-color: #c1c3da;
    color: #ffffff;
    text-align: center;
    font-size: 25px;
    border: none;
    padding: 1rem 2rem;
    border-radius: 25px;
    margin: 0 150px;

}

button:active {
    transform: translateY(2px);
}
.rainbow-input{
    -webkit-appearance:none;
    -moz-appearance:none;
    outline: none;
    width:25px;
    height:25px;
    background-color:#ffffff;
    border:solid 1px #dddddd;
    border-radius:50%;
    margin:6px 2px 1px 2px;
    padding:0;
    position:relative;
    display:inline-block;
    vertical-align:top;
    transition:background-color ease 0.1s;
}
.rainbow-input:checked {
    background:#7c80b8;
}
.rainbow-input:checked::after {
    content:'';
    top:8px;
    left:6px;
    position:absolute;
    border:#fff solid 2px;
    border-top:none;
    border-right:none;
    height:6px;
    width:10px;
    transform:rotate(-45deg);
}
.rainbow{
    background-color: #c1c3da;
    color: #ffffff;
    text-align: center;
    font-size: 25px;
    border: none;
    padding: 1rem 2rem;
    border-radius: 25px;
    margin: 0 100px;
    
}

</style>