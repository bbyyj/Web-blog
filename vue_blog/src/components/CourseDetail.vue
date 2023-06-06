<template>
    <div class="bg">
        <div class="container">
            <h1>课程详情</h1>
            <div class="course-info">
                <p>课程编号：{{ courseInfo.subject_id }}</p>
                <p>课程名称：{{ courseInfo.subject_name }}</p>
                <p>教师：{{ courseInfo.teacher }}</p>
                <p>课程类型：{{ courseInfo.classification }}</p>
                <p>学分：{{ courseInfo.credit }}</p>
                <p>学院：{{ courseInfo.college }}</p>
                <p>考勤：{{ courseInfo.attendance }}</p>
                <p>分数：{{ courseInfo.score }}</p>
                <p>评价：{{ courseInfo.evaluation }}</p>
            </div><hr class="hr01">
            <div>

                <el-card>
                    <div class="comment-outer">
                        <div class="comment-body">
                            <div style="font-weight: bold">评论</div>
                            <hr class="hr01">
                            <!-- 评论内容 -->
                            <div class="comment" :key="item.id" v-for="item in commentList">
                                <span class="date">
                                    {{ item.createTime | fromatDate("yyyy-MM-dd hh:mm:ss") }}
                                </span>
                                <div class="comment-content">
                                    {{ item.comment }}
                                </div>
                            </div>
                        </div>
                    </div>
                    <!-- 分页区域 -->
                    <el-pagination popper-class="select-down" @size-change="handleSizeChange" @current-change="handleCurrentChange"
                        :current-page="queryInfo.pageNum" :page-sizes="[10, 12, 15, 20]" :page-size="queryInfo.pageSize"
                        layout="total, sizes, prev, pager, next, jumper" :total="total">
                    </el-pagination>
                </el-card>

                <div class="text-bg">
                    <!-- 评论栏 -->
                    <textarea class="replay" v-model="comment.content"></textarea>

                    <!-- 评论 -->
                    <b-container fluid>
                        <b-row>
                            <b-col sm="3">
                                <b-button variant="outline-info" @click="publishComment">
                                    <b-icon icon="pencil-square" font-scale="0.9"></b-icon>
                                    发布评论
                                </b-button>
                            </b-col>
                        </b-row>
                    </b-container>
                </div>
            </div>
        </div>
    </div>
</template>
<script>

export default {
    data() {
        return {
            courseInfo: [],
            total: 0,
            //获取的评论内容
            commentList: [],
            queryInfo: {
                pageNum: 1,
                pageSize: 10,
            },
            comment: {         //评论数据
                subject_id: this.$route.query.subject_id,
                content: "",
            },
        }
    },
    //日期格式相关
    filters: {
        fromatDate: function (val, fmt) {
            var date = new Date(val);
            let ret;
            const opt = {
                "y+": date.getFullYear().toString(),        // 年
                "M+": (date.getMonth() + 1).toString(),     // 月
                "d+": date.getDate().toString(),            // 日
                "h+": date.getHours().toString(),           // 时
                "m+": date.getMinutes().toString(),         // 分
                "s+": date.getSeconds().toString()          // 秒
                // 有其他格式化字符需求可以继续添加，必须转化成字符串
            };
            for (let k in opt) {
                ret = new RegExp("(" + k + ")").exec(fmt);
                if (ret) {
                    fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
                };
            };
            return fmt;
        }
    },
    async created() {
        let subject_id = this.$route.query.subject_id;
        //获取评论内容
        this.getCommentList();
        // 请求对应的课程详情信息
        try {
            const { data: res } = await this.$axios.get("/myblog/electionDetailedList", {
                params: { subject_id: subject_id },
            });
            console.log(res.data[0])
            console.log(subject_id)
            this.courseInfo = res.data[0][0];
        } catch (error) {
            console.error(error);
        }
    },
    methods: {
        getCommentList: async function () {
            //获取评论内容
            let subject_id = this.$route.query.subject_id;
            const { data: res } = await this.$axios.get("/myblog/electionCommentList", { params: { subject_id: subject_id, pageNum: this.queryInfo.pageNum, pageSize: this.queryInfo.pageSize } });
            if (res.status === 1) {
                this.commentList = res.data[0],
                    this.total = res.data[1]
            }
            else {
                window.alert("获取评论失败！")
            }
        },
        publishComment: async function () {
            if (!this.comment.content) {
                this.$message.warning("请输入评论内容")
                return
            }
            const { data: res } = await this.$axios.post("/myblog/addElectionComment", {
                subject_id: this.$route.query.subject_id,
                comment: this.comment.content,
            });

            console.log(this.comment.content)
            if (res.status === 101) {
                this.getCommentList();
                this.comment.content = "";
                window.alert("评论成功")
            }
            else {
                this.getCommentList();
                window.alert("评论失败,请检查网络设置")
            }
        },
        ///////////////////分页相关的
        // 监听pagesize 改变的事件
        handleSizeChange: function (pagesize) {
            this.queryInfo.pageSize = pagesize;
            this.getCommentList()
        },
        // 页码值发送变化
        handleCurrentChange: function (newPage) {
            this.queryInfo.pageNum = newPage;
            this.getCommentList();
        },
    }
}
</script>

<style lang="less" scoped>
.bg{
    background: linear-gradient(-45deg, #f5e2e0, #c2bbce, #837cb8, #3d3952);
    background-size: 300% 300%;
    animation: gradient 20s ease-in-out infinite;
    min-height: 1700px;
}
@keyframes gradient {
	0% {
		background-position: 0% 50%;
	}
	50% {
		background-position: 100% 50%;
	}
	100% {
		background-position: 0% 50%;
	}
}
.hr01{
    margin-top: 30px;
    border: 0;
    height: 2px;
    background-image: linear-gradient(to right, #c2bbce, #fef6f5, #c2bbce);
}
.container {
    padding-top: 80px;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    font-family: HYYueXiChuKaiJ-2;
    h1 {
        color: #f0f0f0;
        font-size: 48px;
        margin-bottom: 1em;
    }
}
// 课程信息
.course-info {
    background-color: #fefefe;
    opacity: 0.6;
    padding: 2em;
    border-radius: 2em;
    font-size: 20px;
    p {
        margin-bottom: 1em;
        color: #3d3952;
    }
}
.btn-outline-info{
    margin: 20px 40px 0px 40px;
    background-color: #c1c3db;
    border: 0;
    color: #565ca4;
    border-radius: 10px;
}
.btn-outline-info:hover, .btn-outline-info:focus, .btn-outline-info:active{
    border: 0 !important;
    background-color: #565ca4 !important;
    color: #fef6f5;
}
.replay {
    width: 90%;
    min-height: 180px;
    background: #c1c3db;
    opacity: 0.8;
    margin: 30px 40px 0px 40px;
    padding: 2em;
    border-radius: .28571429rem;

    border: 0;
    outline: 0;
    color: rgba(0, 0, 0, .87);
    // transition: color .1s ease, border-color .1s ease;
    font-size: 1em;
    line-height: 1.2857;
    resize: vertical;
}
.text-bg{
    margin: 20px 0px 20px 0px;
    padding: 10px;
    background-color: #fef6f5;
    border-radius: 1rem;
}
/*评论相关的样式 */
.el-card.is-always-shadow{
    border-radius: 1rem !important;
}
/deep/ .el-card__body{
    background-color: #fef6f5;
    
}
.comment-outer {
    width: 95%;
    margin-left: 2.5%;
    margin-right: 2.5%;
    border-radius: .28571429rem;
    box-shadow: none;
    border: 1px solid #565ca4;
    padding: 1.5em;

    min-height: 300px;
    background: #ffffff;
    opacity: 0.7;
}

.comment-body {
    position: relative;
    background: #FFF;
    margin: 0.5rem 0;
    padding: 1em;
    border-radius: .28571429rem;
    border: 2px solid #8c90c0;
    min-height: 260px;
    hr{
        background-image: linear-gradient(to right, #c2bbce, #c2bbce, #c2bbce);
        height: 1px;
    }

}

.comment {
    width: 90%;
    padding: 6px 18px;
    margin-left: 10px;
    margin-right: 20px;
    margin-top: 12px;
    display: block;
}


/deep/ .el-pager .number:hover, .el-pager .number:active{
    color: #393f51;
}
/deep/ .el-pager .number{
    color: #565ca4;
    font-size: 16px;
}
/deep/ .el-input__inner{
    color: #565ca4;
    font-size: 16px;
    border: 1px solid #565ca4 !important; 
}

/deep/ .btn-prev{
    border-radius: 4px 0 0 4px ;
}
/deep/ .btn-prev:hover{
    color: #565ca4;
}
/deep/ .btn-next{
    border-radius: 0 4px 4px 0;
}
/deep/ .btn-next:hover{
    color: #565ca4;
}
.el-pagination{
    padding-top: 30px;
    padding-left: 270px;
}
</style>


<style lang="less">
.select-down {
    .el-select-dropdown__item.selected{
        color: #565ca4 !important;
    }
    .el-select-dropdown__item.hover, .el-select-dropdown__item:hover{
        margin: 5px;
    }
	
}

</style>
