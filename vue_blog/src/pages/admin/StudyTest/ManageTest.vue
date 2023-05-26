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
                    <el-option v-for="item in chapters" :key="item.id" :label="item.name" :value="item.name">
                    </el-option>
                </el-select>
                <!-- 新增题目按钮 -->
                <el-button type="primary" @click="showAddDialog">新增题目</el-button>
            </div>

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
                        <!-- 编辑和删除按钮 -->
                        <el-button size="mini" @click="handleEdit(scope.$index)">编辑</el-button>
                        <el-button size="mini" @click="handleDelete(scope.$index)">删除</el-button>
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
                id: 0,
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

            subjects: [{ name: "机组" }], //科目列表
            chapters: [{ name: "机组" }], //章节列表
            selectedSubject: '', //选中的科目
            selectedChapter: '', //选中的章节
            links: [{
                id: 0,
                question: "test",
                selection1: "1",
                selection2: "2",
                selection3: "3",
                selection4: "4",
                answer: "3",
            }] //存储获取的题目列表
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
                id: 0,
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

        //分页大小改变
        handleSizeChange(val) {
            this.queryInfo.pageSize = val;
            this.getExamQuestion();
        },
        //分页当前页改变
        handleCurrentChange(val) {
            this.queryInfo.pageNum = val;
            this.getExamQuestion();
        },
        handleEdit(index) {
            // 复制被编辑的题目到editInfo
            this.editInfo = { ...this.links[index] };
            // 打开编辑对话框
            this.dialogEditFormVisible = true;
        },
        //删除按钮
        async handleDelete(index) {
            // 使用浏览器自带的确认框，提示用户确认删除
            if (window.confirm('此操作将永久删除该题目, 是否继续?')) {
                try {
                    // 发送delete HTTP请求来删除题目
                    const { data: res } = await axios.delete(`/myblog/tacitInfo/${this.links[index].id}`);
                    if (res.status === 1) {
                        // 显示成功消息
                        window.alert('删除成功!');
                        // 重新获取题目列表
                        this.getExamQuestion();
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
            // 检查是否已选择科目和章节
            if (this.selectedSubject.trim() === '') {
                window.alert('请先选择科目！');
                return;
            }
            if (this.selectedChapter.trim() === '') {
                window.alert('请先选择章节！');
                return;
            }
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
                const { data: res } = this.postInfo.id
                    ? await axios.put("/myblog/tacitInfo", this.postInfo) // 发送PUT HTTP请求来更新题目
                    : await axios.post("/myblog/tacitInfo", this.postInfo); // 发送POST HTTP请求来创建题目
                if (res.status === 1) {
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
                const { data: res } = this.editInfo.id
                    ? await axios.put("/myblog/tacitInfo", this.editInfo) // 发送PUT HTTP请求来更新题目
                    : await axios.post("/myblog/tacitInfo", this.editInfo); // 发送POST HTTP请求来创建题目
                if (res.status === 1) {
                    // 显示成功消息
                    window.alert('编辑成功!');
                    // 重新获取题目列表
                    this.getExamQuestion();
                    // 关闭编辑对话框
                    this.dialogFormVisible = false;
                } else {
                    // 显示失败消息
                    window.alert('编辑失败!');
                }
            } catch (error) {
                // 显示异常消息
                window.alert('操作失败，发生异常!');
            }
        },

        //获取全部学科列表
        async getSubjectsList() {
            const { data: res } = await axios.get("/myblog/subjectList");
            if (res.status === 1) {
                this.subjects = res.data[0];
            }
        },
        //获取对应学科下全部的章节信息
        async getChaptersList(subject) {
            const { data: res } = await axios.get("/myblog/chapterList", {
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
            const { data: res } = await axios.get("/myblog/获取题目的接口", {
                params: {
                    subject: this.selectedSubject,
                    chapter: this.selectedChapter,
                }
            });
            if (res.status === 1) {
                this.links = res.data[0].list;
                this.total = res.data[0].total;
            } else {
                this.$message.error("获取题目失败,请重试");
            }
        }
    },
    created() {
        this.getSubjectsList();
    }
}
</script>
