<template>
    <div>
        
        <el-card>
            <!-- 选择区域 -->
            <el-select v-model="selectedCategory" placeholder="选择课程类别" @change="getCourses" clearable>
                <el-option v-for="item in categories" :key="item" :label="item" :value="item">
                </el-option>
            </el-select>
            <!-- 添加新的课程 -->
            <el-button class="add" type="primary" @click="addCourse">添加新课程</el-button>
            <!-- 列表区域  data中是获取到的数据-->
            <el-table :data="courses" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="课程编号" prop="subject_id" width="90px"></el-table-column>
                <el-table-column label="课程名称" prop="subject_name"></el-table-column>
                <el-table-column label="教师" prop="teacher"></el-table-column>
                <el-table-column label="学分" prop="credit" width="60px"></el-table-column>
                <el-table-column label="学院" prop="college"></el-table-column>
                <el-table-column label="考勤" prop="attendance"></el-table-column>
                <el-table-column label="分数" prop="score"></el-table-column>
                <el-table-column label="考试" prop="evaluation"></el-table-column>
                <!-- 操作按钮 -->
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button class="edt" size="mini" @click="handleEdit(scope.$index)">编辑</el-button>
                        <el-button class="del" size="mini" type="danger" @click="handleDelete(scope.row.subject_id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.pageNum" :page-sizes="[5, 10, 15, 20]" :page-size="queryInfo.pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>

            <!-- 课程信息编辑对话框 -->
            <el-dialog title="编辑课程信息" :visible.sync="dialogEditFormVisible">
                <el-form label-width="80px">
                    <el-form-item label="课程编号：">
                        <el-input v-model="editInfo.subject_id" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="课程类别：">
                        <el-select v-model="editInfo.classification" placeholder="选择课程类别">
                            <el-option v-for="item in categories" :key="item" :label="item" :value="item"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="课程名称：">
                        <el-input v-model="editInfo.subject_name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="教师：">
                        <el-input v-model="editInfo.teacher" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="学分：">
                        <el-select v-model="editInfo.credit" placeholder="选择学分">
                            <el-option v-for="item in credits" :key="item" :label="item" :value="item"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="学院：">
                        <el-select v-model="editInfo.college" placeholder="选择学院">
                            <el-option v-for="item in colleges" :key="item" :label="item" :value="item"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="考勤：">
                        <el-input v-model="editInfo.attendance" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="分数：">
                        <el-input v-model="editInfo.score" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="考试：">
                        <el-input v-model="editInfo.evaluation" autocomplete="off"></el-input>
                    </el-form-item>
                    <!--其他课程信息字段...-->
                </el-form>
                <!--取消编辑 和 确认提交-->
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancelEdit">取 消</el-button>
                    <el-button type="primary" @click="commitEdit">确 定</el-button>
                </div>
            </el-dialog>


            <!-- 添加课程信息对话框 -->
            <el-dialog title="添加新课程" :visible.sync="dialogAddFormVisible">
                <el-form label-width="80px">
                    <!-- 表单内容，你可以根据需要添加更多的表单项 -->
                    <el-form-item label="课程编号：">
                        <el-input v-model="newInfo.subject_id" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="课程名称：">
                        <el-input v-model="newInfo.subject_name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="教师：">
                        <el-input v-model="newInfo.teacher" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="类别：">
                        <el-select v-model="newInfo.classification" placeholder="选择课程类别">
                            <el-option v-for="item in categories" :key="item" :label="item" :value="item"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="学分：">
                        <el-select v-model="newInfo.credit" placeholder="选择学分">
                            <el-option v-for="item in credits" :key="item" :label="item" :value="item"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="学院：">
                        <el-select v-model="newInfo.college" placeholder="选择学院">
                            <el-option v-for="item in colleges" :key="item" :label="item" :value="item"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="考勤：">
                        <el-input v-model="newInfo.attendance" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="分数：">
                        <el-input v-model="newInfo.score" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="考试：">
                        <el-input v-model="newInfo.evaluation" autocomplete="off"></el-input>
                    </el-form-item>
                    <!-- 其他表单项... -->
                </el-form>
                <!--取消添加 和 确认添加-->
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancelAdd">取 消</el-button>
                    <el-button type="primary" @click="commitAdd">确 定</el-button>
                </div>
            </el-dialog>

        </el-card>
    </div>
</template>

<script>

export default {
    name: "manageCourses",
    data() {
        return {
            //类别相关
            selectedCategory: '',
            categories: ['公必', '公选', '专必', '专选'],
            credits: ['1', '2', '3', '4', '5', '6'],
            colleges: ["软件工程学院", "马克思主义学院", "数学学院", "物理与天文学院"],
            courses: [],
            // 控制添加对话框的显示和隐藏
            dialogAddFormVisible: false,
            // 控制编辑对话框的显示和隐藏
            dialogEditFormVisible: false,
            total: 0,
            //编辑的课程信息
            editInfo: {
                subject_id: "",
                subject_name: "",
                teacher: "",
                classification: "",
                credit: "",
                college: "",
                attendance: "",
                score: "",
                evaluation: ""
            },
            newInfo: {
                subject_id: "",
                subject_name: "",
                teacher: "",
                classification: "",
                credit: "",
                college: "",
                attendance: "",
                score: "",
                evaluation: ""
            },
            //查询的信息
            queryInfo: {
                pageNum: 1,
                pageSize: 5
            },
        }
    },
    created() {
        //获取课程列表
        this.getALLCourses();
    },
    methods: {
        //开始的获取全部的这个进行保留
        async getALLCourses() {
            const { data: res } = await this.$axios.get("/admin/electionList", { params: this.queryInfo });
            if (res.status !== 1) {
                this.$message.error("获取课程信息失败，请重试！")
                return
            }
            else {
                //进行赋值
                this.courses = res.data[0],
                    this.total = res.data[1]
            }
        },
        //获取对应类别下的课程
        async getCourses() {
            if (this.selectedCategory === "") {
                this.getALLCourses()
                return
            }
            const { data: res } = await this.$axios.get("/admin/electionByClass", { params: { classification: this.selectedCategory, pageNum: this.queryInfo.pageNum, pageSize: this.queryInfo.pageSize } });
            if (res.status === 1) {
                //进行赋值

                this.courses = res.data[0]
                this.total = res.data[1]
            }
            else {
                this.$message.error("获取课程信息失败，请重试！")
                return
            }
        },

        // 显示添加新课程的对话框
        addCourse() {
            this.newInfo = {};
            this.dialogAddFormVisible = true;
        },
        // 取消添加操作
        cancelAdd() {
            this.newInfo = {};
            this.dialogAddFormVisible = false;
        },
        // 提交添加操作
        async commitAdd() {
            try {
                const { data: res } = await this.$axios.post("/admin/addElection", this.newInfo);
                if (res.status !== 101) {
                    this.$message.error("操作失败，请重试！")
                } else {
                    this.$message.success("操作成功！");
                    this.cancelAdd();
                    await this.getCourses();
                }
            } catch (error) {
                console.error(error);
            }
        },
        //进行编辑
        handleEdit(index) {
            //显示编辑窗口
            this.dialogEditFormVisible = true
            //此处进行绑定到editInfo
            this.editInfo = { ...this.courses[index] }
        },
        //取消编辑操作
        cancelEdit() {
            this.editInfo = {}
            this.dialogEditFormVisible = false
        },

        //删除该条
        async handleDelete(id) {
            this.$messageBox.confirm('确认删除该课程?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                const { data: res } = await this.$axios.delete("/admin/deleteElection", { params: { subject_id: id } });
                if (res.status !== 401) {
                    this.$message.error("删除失败，请重试！")
                } else {
                    this.$message.success("删除成功！")
                    await this.getCourses();
                }
            }, () => {
                this.$message({ type: 'info', message: '已取消删除' });
            });
        },

        //提交编辑操作
        async commitEdit() {
            try {
                const { data: res } = await this.$axios.put("/admin/updateElection", this.editInfo);
                if (res.status !== 101) {
                    this.$message.error("操作失败，请重试！");
                } else {
                    this.$message.success("操作成功！");
                    this.cancelEdit();
                    await this.getCourses();
                }
            } catch (error) {
                console.error(error);
            }
        },
        /////////////////////分页相关
        // 监听pagesize 改变的事件
        handleSizeChange: function (pagesize) {
            this.queryInfo.pageSize = pagesize;
            this.getCourses();
        },
        // 页码值发送变化
        handleCurrentChange: function (newPage) {
            this.queryInfo.pageNum = newPage;
            this.getCourses();
        },

    },

}
</script>

<style scoped>
.add {
    float: right !important;
    margin-right: 50px;
    background-color: #baaaca17;
    color: #baaacaee;
    border-color: #baaacaee;
}

.add:hover, .add:focus {
    background-color: #baaacaee;
    color: #fff;
    border-color: #baaacaee;
}
.show-categories {
    float: right;
    margin-right: 30px;
}

.edt:hover {
    background-color: #baaacaee;
    color: #fff;
    border-color: #baaacaee;
}

.del {
    background-color: #f6727218;
    color: #f67272ac;
    border-color: #f67272ac;
}

.del:hover {
    background-color: #f67272ac;
    color: #fff;
    border-color: #f67272ac;
}

</style>
