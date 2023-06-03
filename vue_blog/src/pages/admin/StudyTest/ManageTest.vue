<template>
    <div>
        <el-card>
            <!-- 筛选栏 -->
            <div class="filter-container">
                <el-select v-model="selectedSubject" placeholder="选择科目" @change="getChaptersList" clearable>
                    <el-option v-for="item in subjects" :key="item.id" :label="item.name" :value="item.name">
                    </el-option>
                </el-select>
                <el-select v-model="selectedChapter" placeholder="选择章节" @change="getExamQuestion" clearable>
                    <el-option v-for="item in chapters" :key="item.id" :label="item.title" :value="item.title">
                    </el-option>
                </el-select>
                <!-- 新增题目按钮 -->
                <el-button class="add" type="primary" @click="showAddDialog" >新增题目</el-button>
            </div>

            <!-- 列表区域 -->
            <el-table :data="examQuestions" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="问题" prop="text" width="300px"></el-table-column>
                <el-table-column label="选项1" prop="first_answer"></el-table-column>
                <el-table-column label="选项2" prop="second_answer"></el-table-column>
                <el-table-column label="选项3" prop="third_answer"></el-table-column>
                <el-table-column label="选项4" prop="fourth_answer"></el-table-column>
                <el-table-column label="正确答案" prop="correct_answer" width="300px"></el-table-column>
                <el-table-column label="操作" width="150">
                    <template slot-scope="scope">
                        <!-- 编辑和删除按钮 -->
                        <el-button size="mini" @click="handleEdit(scope.$index)">编辑</el-button>
                        <el-button size="mini" @click="handleDelete(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.pageNum" :page-sizes="[10, 12, 15, 20]" :page-size="queryInfo.pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>

            <!-- 新增题目的弹出窗口 -->
            <el-dialog title="问答信息" :visible.sync="dialogFormVisible">
                <el-form label-width="80px">
                    <el-form-item label="问题:">
                        <el-input v-model="postInfo.question" autocomplete="off" clearable></el-input>
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
                    <el-form-item label="正确答案:">
                        <el-input v-model="postInfo.answer" autocomplete="off" clearable></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取 消</el-button>
                    <el-button type="primary" @click="submitPostInfo">确 定</el-button>
                </div>
            </el-dialog>
            <!-- 编辑题目的弹出窗口 -->
            <el-dialog title="编辑题目信息" :visible.sync="dialogEditFormVisible">
                <el-form label-width="80px">
                    <el-form-item label="问题:">
                        <el-input v-model="editInfo.question" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="选项1:">
                        <el-input v-model="editInfo.selection1" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="选项2:">
                        <el-input v-model="editInfo.selection2" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="选项3:">
                        <el-input v-model="editInfo.selection3" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="选项4:">
                        <el-input v-model="editInfo.selection4" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="正确答案:">
                        <el-input v-model="editInfo.answer" autocomplete="off" clearable></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="dialogEditFormVisible = false">取 消</el-button>
                    <el-button type="primary" @click="submitEditInfo">确 定</el-button>
                </div>
            </el-dialog>
        </el-card>
    </div>
</template>
<script>
import axios from "axios";

export default {
    name: "ManageTacit",
    data() {
        return {
            //编辑信息的对话框
            dialogEditFormVisible: false,
            editInfo: {
                id: 0,
                question: "",
                selection1: "",
                selection2: "",
                selection3: "",
                selection4: "",
                answer: "",
            },
            //新增题目的信息对话框
            dialogFormVisible: false,
            postInfo: {
                question: "",
                selection1: "",
                selection2: "",
                selection3: "",
                selection4: "",
                answer: "",
            },

            innerVisible: false,
            total: 0,
            queryInfo: {
                pageNum: 1,
                pageSize: 10,
            },
            subjects: [], //科目列表
            chapters: [], //章节列表
            selectedSubject: '', //选中的科目
            selectedChapter: '', //选中的章节
            examQuestions: [] //存储获取的题目列表
        };
    },
    watch: {
        selectedSubject(newSubject) {
            // 当 selectedSubject 变化时，设置 selectedChapter 为空
            this.selectedChapter = '';
            // 调用 getChaptersList 方法获取新科目的章节列表
            if (newSubject) {
                this.getChaptersList(newSubject);
            }
        }
    },
    methods: {
        showAddDialog() {
            // 检查是否已选择科目和章节
            if (this.selectedSubject.trim() === '') {
                window.alert('请先选择科目和对应章节！');
                return;
            }
            if (this.selectedChapter.trim() === '') {
                window.alert('请先选择章节！');
                return;
            }
            // 清空postInfo
            this.postInfo = {
                question: "",
                selection1: "",
                selection2: "",
                selection3: "",
                selection4: "",
                answer: "",
            };
            // 如果已选择科目和章节，显示新增题目的对话框
            this.dialogFormVisible = true;
        },

        // 监听pagesize 改变的事件
        handleSizeChange: function (pagesize) {
            this.queryInfo.pageSize = pagesize;
            this.getExamQuestion();
        },
        // 页码值发送变化
        handleCurrentChange: function (newPage) {
            this.queryInfo.pageNum = newPage;
            this.getExamQuestion();
        },
        //处理编辑操作
        handleEdit(index) {
            // 在下一个DOM更新周期执行数据赋值操作
            const question = this.examQuestions[index];
            this.editInfo.id = question.id;
            this.editInfo.question = question.text;
            this.editInfo.selection1 = question.first_answer;
            this.editInfo.selection2 = question.second_answer;
            this.editInfo.selection3 = question.third_answer;
            this.editInfo.selection4 = question.fourth_answer;
            this.editInfo.answer = question.correct_answer;

            this.dialogEditFormVisible = true; // 先设置对话框可见状态
        },

        //删除按钮
        async handleDelete(index) {
            // 使用浏览器自带的确认框，提示用户确认删除
            if (window.confirm('此操作将永久删除该题目, 是否继续?')) {
                try {
                    // 发送delete HTTP请求来删除题目/admin/deleteExam
                    const { data: res } = await this.$axios.delete("/admin/deleteExam", { params: { id: index } });
                    console.log(index)
                    if (res.status === 401) {
                        // 显示成功消息
                        window.alert('删除成功!');
                        //进行页码等相关转换操作
                        if (this.queryInfo.pageNum === Math.ceil(this.total / this.queryInfo.pageSize) && this.examQuestions.length === 1) {
                            this.queryInfo.pageNum -= 1
                            if (this.queryInfo.pageNum <= 0) {
                                this.queryInfo.pageNum = 1
                                return
                            }
                        }
                        //刷新列表
                        await this.getExamQuestion();
                    } else {
                        // 显示失败消息
                        window.alert('删除失败!');
                    }
                } catch (error) {
                    // 显示异常消息
                    window.alert('操作失败，发生异常!');
                }
            } else {
                // 用户取消了删除操作
                window.alert('已取消删除');
            }
        },
        //提交新的题目信息
        async submitPostInfo() {
            // 检查所有信息是否已被填充
            for (let key in this.postInfo) {
                if (this.postInfo.hasOwnProperty(key) && key !== 'id') {
                    if (this.postInfo[key].trim() === '') {
                        // 如果有任何一个信息未被填充，显示警告消息并退出函数
                        window.alert('所有信息都必须填充完整！');
                        return;
                    }
                }
            }
            try {
                // 如果所有信息都已被填充，继续提交操作
                const { data: res } = await this.$axios.post("/admin/createExam", {
                    subject: this.selectedSubject,
                    chapter: this.selectedChapter,
                    text: this.postInfo.question,
                    first_answer: this.postInfo.selection1,
                    second_answer: this.postInfo.selection2,
                    third_answer: this.postInfo.selection3,
                    fourth_answer: this.postInfo.selection4,
                    correct_answer: this.postInfo.answer
                });
                console.log(this.selectedChapter)
                if (res.status === 101) {
                    // 显示成功消息
                    window.alert('新增成功!');
                    // 重新获取题目列表
                    this.getExamQuestion();
                    // 关闭编辑对话框
                    this.dialogFormVisible = false;
                } else {
                    // 显示失败消息
                    window.alert('新增失败!');
                }
            } catch (error) {
                // 显示异常消息
                window.alert('操作失败，发生异常!');
            }
        },

        //提交新编辑的章节内容
        async submitEditInfo() {
            // 检查所有信息是否已被填充
            for (let key in this.editInfo) {
                if (this.editInfo.hasOwnProperty(key) && key !== 'id') {
                    if (this.editInfo[key].trim() === '') {
                        // 如果有任何一个信息未被填充，显示警告消息并退出函数
                        window.alert('所有信息都必须填充完整！');
                        return;
                    }
                }
            }
            try {
                // 如果所有信息都已被填充，继续提交操作
                const { data: res } = await this.$axios.put("/admin/updateExam", {
                    id: this.editInfo.id,
                    subject: this.selectedSubject,
                    chapter: this.selectedChapter,
                    text: this.editInfo.question,
                    first_answer: this.editInfo.selection1,
                    second_answer: this.editInfo.selection2,
                    third_answer: this.editInfo.selection3,
                    fourth_answer: this.editInfo.selection4,
                    correct_answer: this.editInfo.answer
                });
                if (res.status === 101) {
                    // 显示成功消息
                    window.alert('编辑成功!');
                    this.dialogEditFormVisible = false;
                    // 重新获取题目列表
                    this.getExamQuestion();
                } else {
                    // 显示失败消息
                    window.alert('编辑失败!');
                }
            } catch (error) {
                // 显示异常消息
                window.alert('操作失败，发生异常!');
            }
        },

        //获取全部学科列表 //已经完成/////////////////////////
        async getSubjectsList() {
            const { data: res } = await axios.get("/myblog/subjectList");
            if (res.status === 1) {
                this.subjects = res.data[0];
            }
        },
        //获取对应学科下全部的章节信息 //已经完成 ////////////////////////////////
        async getChaptersList(subject) {
            const { data: res } = await axios.get("/admin/chapterList", {
                params: { name: subject }
            });
            if (res.status === 1) {
                this.chapters = res.data[0];
            } else {
                this.$message.error("获取对应章节失败,请重试")
                return
            }
        },
        //获取对应科目+章节的全部题目
        async getExamQuestion() {
            const { data: res } = await axios.get("/admin/examList", {
                params: {
                    typename: this.selectedSubject,
                    title: this.selectedChapter,
                    pageNum: this.queryInfo.pageNum,
                    pageSize: this.queryInfo.pageSize
                }
            });

            if (res.status === 1) {
                //赋值题目
                this.total = res.data[1]
                this.examQuestions = res.data[0];
            }
            else {
                this.$message.error("获取题目失败,请重试");
            }
        }
    },
    created() {
        this.getSubjectsList();
    }
}
</script>

<style>

.add {
    float: right !important;
    margin-right: 50px;
    background-color: #baaaca18;
    color: #baaaca;
    border-color: #baaaca;
}

.add:hover {
    background-color: #baaacaee;
    color: #fff;
    border-color: #baaacaee;
}

</style>
