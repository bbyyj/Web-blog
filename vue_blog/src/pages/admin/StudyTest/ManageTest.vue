<template>
    <div>
        <el-card>
            <!-- 列表区域 -->
            <el-table :data="links" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="问题" prop="question" width="300px"></el-table-column>
                <el-table-column label="选项1" prop="selection1"></el-table-column>
                <el-table-column label="选项2" prop="selection2"></el-table-column>
                <el-table-column label="选项3" prop="selection3"></el-table-column>
                <el-table-column label="选项4" prop="selection4"></el-table-column>
                <el-table-column label="正确答案" prop="answer" width="300px"></el-table-column>
                <el-table-column label="操作" width="150">
                    <template slot-scope="scope">
                        <!--编辑内容-->
                        <el-button size="mini" @click="handleEdit(scope.$index)">编辑</el-button>
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
                    <el-form-item label="问题:">
                        <el-input v-model="postInfo.name" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="选项1:">
                        <el-input v-model="postInfo.selection1" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="选项2:">
                        <el-input v-model="postInfo.selection2" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="选项3:">
                        <el-input v-model="postInfo.selection3" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="选项4:">
                        <el-input v-model="postInfo.selection4" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="答案：">
                        <el-input v-model="postInfo.answer" autocomplete="off" clearable></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancel">取 消</el-button>
                    <!--提交编辑后的内容-->
                    <el-button type="primary" @click="commitcontent">确 定</el-button>
                </div>
            </el-dialog>

            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogCategoryVisible = false">取 消</el-button>
            </div>

        </el-card>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: "ManageTacit",
    data() {
        return {
            dialogFormVisible: false,
            dialogCategoryVisible: false,
            innerVisible: false,
            total: 0,
            //编辑的内容
            postInfo: [{
                id: 0,
                question: "",
                selection1: "",
                selection2: "",
                selection3: "",
                selection4: "",
                answer: "",
            }
            ],
            queryInfo: {
                pageNum: 1,
                pageSize: 10
            },
        }
    },
    methods: {
        //获取所有内容
        async getTacitInfo() {
            const { data: res } = await this.$axios.get("/后端接口", { params: this.queryInfo });
            if (res.status !== 1) {
                this.$message.error("获取列表失败，请重试！")
                return
            }
            let postInfo
            if (res.data.length > 2) {
                postInfo = res.data[0]
            } else {
                this.$message.error("获取列表失败，请重试！")
                return
            }
            //总数量
            this.total = 10

            links.forEach((val) => {
                this.links.push({ ...val, category: m.get(val.categoryId) })
            })
        },
        //编辑问题和答案等内容
        async handleEdit(index) {
            this.dialogFormVisible = true
            //获取已有内容
            this.postInfo = { ...this.links[index] }
            const val = this.categories.find(item => {
                return item.id === this.links[index].categoryId
            })
        },
        // 监听pagesize 改变的事件
        handleSizeChange: function (pagesize) {
            this.queryInfo.pageSize = pagesize;
            this.getLinkList();
        },
        // 页码值发送变化
        handleCurrentChange: function (newPage) {
            this.queryInfo.pageNum = newPage;
            this.getLinkList();
        },

        //取消编辑
        cancel() {
            this.postInfo = {
                id: 0,
                question: "",
                selection1: "",
                selection2: "",
                selection3: "",
                selection4: "",
                answer: "",
            }
            this.dialogFormVisible = false
        },

        //提交编辑后的内容
        async commitcontent() {
            let res
            if (this.postInfo.id === 0) {
                res = await this.$axios.post("/admin/增加接口", this.postInfo)
            } else {
                res = await this.$axios.put("/admin/更新接口", this.postInfo)
            }
            if (res.data.status !== 101) {
                this.$message.error("操作失败，请重试！")
            } else {
                this.cancel()
                await this.getTacitInfo()
                this.$message.success("操作成功！")
            }
        },
    },
    created() {
        this.getTacitInfo()
    }
}
</script>

<style scoped></style>