<template>
    <div>
        <el-card>
            <!-- 列表区域  data中是获取到的数据-->
            <el-table :data="QandA" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="问题" prop="question" width="300px"></el-table-column>
                <el-table-column label="答案" prop="answer" width="180px"></el-table-column>
                <el-table-column label="操作" width="150">
                    <template slot-scope="scope">
                        <!--回复和删除按钮-->
                        <el-button size="mini" @click="handleResponse(scope.$index)">回复</el-button>
                        <el-button size="mini" type="danger" @click="handleDelete(scope.row.parent_id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.pageNum" :page-sizes="[10, 12, 15, 20]" :page-size="queryInfo.pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>

            <el-dialog title="问答信息" :visible.sync="dialogFormVisible">
                <el-form label-width="80px">
                    <el-form-item label="你的回复：">
                        <el-input v-model="postInfo.answer" autocomplete="off" clearable></el-input>
                    </el-form-item>
                </el-form>
                <!--取消编辑 和 确认提交-->
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancel">取 消</el-button>
                    <el-button type="primary" @click="commitEdit">确 定</el-button>
                </div>
            </el-dialog>
        </el-card>
    </div>
</template>

<script>

export default {
    name: "manageQuestionandAnswer",
    data() {
        return {
            //已经回答过的问题
            QandA: [],
            dialogFormVisible: false,
            total: 0,
            //准备进行回答
            postInfo: {
                parent_id: 0,
                child_id: 0,
                answer: "",
            },
            //查询的信息
            queryInfo: {
                pageNum: 1,
                pageSize: 10
            },
        }
    },
    created() {
        //获取已经回答过的问题
        this.getQandA();
    },
    methods: {
        async getQandA() {
            const { data: res } = await this.$axios.get("/admin/getUnansweredQA", { params: this.queryInfo });
            if (res.status !== 1) {
                this.$message.error("获取已经回答过的问题失败，请重试！")
                return
            } else {
                this.QandA = res.data[0],
                    //问题总数
                    this.total = res.data[1]
            }
        },

        //进行回答
        async handleResponse(index) {
            //显示编辑窗口
            this.dialogFormVisible = true
            //此处进行绑定到postInfo
            this.postInfo = { ...this.QandA[index] }
        },
        //删除该条
        async handleDelete(id) {
            this.$messageBox.confirm('确认删除该条问答?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                //删除对应问答
                const { data: res } = await this.$axios.delete("/admin/deleteQuestion", { params: { parent_id: id } });
                if (res.status !== 401) {
                    this.$message.error("删除失败，请重试！")
                } else {
                    this.$message.success("删除成功！")
                }
                //进行页码等相关转换操作
                if (this.queryInfo.pageNum === Math.ceil(this.total / this.queryInfo.pageSize) && this.QandA.length === 1) {
                    this.queryInfo.pageNum -= 1
                    if (this.queryInfo.pageNum <= 0) {
                        this.queryInfo.pageNum = 1
                        return
                    }
                }
                //刷新列表
                await this.getQandA();
            }, () => {
                //取消删除操作
                this.$message({ type: 'info', message: '已取消删除' });
            });
        },

        // 监听pagesize 改变的事件
        handleSizeChange: function (pagesize) {
            this.queryInfo.pageSize = pagesize;
            this.getQandA();
        },
        // 页码值发送变化
        handleCurrentChange: function (newPage) {
            this.queryInfo.pageNum = newPage;
            this.getQandA();
        },
        //取消编辑操作
        cancel() {
            this.postInfo = {
                parent_id: 0,
                child_id: 0,
                answer: ""
            }
            //消失
            this.dialogFormVisible = false
        },

        //提交编辑操作
        async commitEdit() {
            let res
            //提交编辑操作
            res = await this.$axios.put("/admin/modifyAnswer", this.postInfo)
            if (res.data.status !== 101) {
                this.$message.error("操作失败，请重试！")
            } else {
                this.cancel()
                await this.getQandA()
                this.$message.success("操作成功！")
            }
        },

    },

}
</script>

<style scoped>
.add {
    float: right !important;
    margin-right: 50px;
}

.show-categories {
    float: right;
    margin-right: 30px;
}
</style>