<template>
    <div class="admin-panel">
        <el-card>
        <!--新增科目-->
        <el-form :model="newSubjectForm" @submit="addSubject" label-width="120px">
            <h2>新增科目</h2>
            <el-form-item label="科目名">
                <el-input v-model="newSubjectForm.name"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="confirmAddSubject">提交</el-button>
            </el-form-item>
        </el-form>

        <!--删除科目-->
        <el-form :model="deleteSubjectForm" @submit="deleteSubject" label-width="120px">
            <h2>删除科目</h2>
            <el-form-item label="科目名">
                <el-select v-model="deleteSubjectForm.name" placeholder="请选择">
                    <el-option v-for="item in subjects" :key="item.id" :label="item.name" :value="item.name">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="danger" @click="confirmDeleteSubject">删除</el-button>
            </el-form-item>
        </el-form>

        <!--新增章节-->
        <el-form :model="newChapterForm" @submit="addChapter" label-width="120px">
            <h2>新增章节</h2>
            <el-form-item label="科目名">
                <el-input v-model="newChapterForm.subject"></el-input>
            </el-form-item>
            <el-form-item label="章节名">
                <el-input v-model="newChapterForm.chapter"></el-input>
            </el-form-item>
            <el-form-item label="描述">
                <el-input v-model="newChapterForm.description"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="confirmAddChapter">提交</el-button>
            </el-form-item>
        </el-form>

        <!--删除章节-->
        <el-form :model="deleteChapterForm" @submit="deleteChapter" label-width="120px">
            <h2>删除章节</h2>
            <el-form-item label="科目名">
                <el-select v-model="deleteChapterForm.subject" placeholder="请选择" @change="updateChapterList">
                    <el-option v-for="item in subjects" :key="item.id" :label="item.name" :value="item.name">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="章节名">
                <el-select v-model="deleteChapterForm.chapter" placeholder="请选择">
                    <el-option v-for="item in chapters" :key="item.id" :label="item.title" :value="item.title">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="danger" @click="confirmDeleteChapter">删除</el-button>
            </el-form-item>
        </el-form>
    </el-card>
    </div>
</template>

<script>
export default {
    data() {
        return {
            newSubjectForm: {
                name: ''
            },
            deleteSubjectForm: {
                name: ''
            },
            newChapterForm: {
                subject: '',
                chapter: '',
                description: ''
            },
            deleteChapterForm: {
                subject: '',
                chapter: ''
            },
            //获取到已有的科目和对应的章节 便于实现删除
            subjects: [],  //科目列表
            chapters: []   //章节列表
        }
    },
    created() {
        this.getSubjectsList();
    },
    methods: {
        //获取科目列表
        async getSubjectsList() {
            const { data: res } = await this.$axios.get("/myblog/subjectList");
            if (res.status === 1) {
                this.subjects = res.data[0];
            }
        },
        //获取科目下的章节内容
        async getChaptersList(subject) {
            const { data: res } = await this.$axios.get("/admin/chapterList", {
                params: { name: subject }
            });
            if (res.status === 1) {
                this.chapters = res.data[0];
            } else {
                this.$message.error("获取对应章节失败,请重试")
                return
            }
        },
        updateChapterList(val) {
            this.getChaptersList(val);
        },
        //////////////////////////////////////////////////////////////////////////////
        //一些确认内容
        //确认新增科目
        confirmAddSubject() {
            if (!this.newSubjectForm.name) {
                this.$message({
                    type: 'info',
                    message: '请输入科目名称'
                });
                return;
            }
            if (window.confirm('确认新增此科目吗?')) {
                this.addSubject();
            } else {
                this.$message({
                    type: 'info',
                    message: '已取消新增'
                });
            }
        },
        //确认删除科目
        confirmDeleteSubject() {
            if (!this.deleteSubjectForm.name) {
                this.$message({
                    type: 'info',
                    message: '请选择要删除的科目'
                });
                return;
            }
            if (window.confirm('确认删除此科目吗?')) {
                this.deleteSubject();
            } else {
                this.$message({
                    type: 'info',
                    message: '已取消删除'
                });
            }
        },
        //确认新增章节
        confirmAddChapter() {
            if (!this.newChapterForm.subject || !this.newChapterForm.chapter || !this.newChapterForm.description) {
                this.$message({
                    type: 'info',
                    message: '请输入对应的完整信息'
                });
                return;
            }
            if (window.confirm('确认新增此章节吗?')) {
                this.addChapter();
            } else {
                this.$message({
                    type: 'info',
                    message: '已取消新增'
                });
            }
        },
        //确认删除章节
        confirmDeleteChapter() {
            /* if (!this.deleteChapterForm.subject || !this.deleteChapterForm.chapter) {
                 this.$message({
                     type: 'info',
                     message: '请选择要删除的科目和章节'
                 });
                 return;
             }*/
            if (window.confirm('确认删除此章节吗?')) {
                this.deleteChapter();
            } else {
                this.$message({
                    type: 'info',
                    message: '已取消删除'
                });
            }
        },

        /////////////////////////////////////////////////////////////////////////////////////
        //增加科目
        async addSubject() {
            try {
                const { data: res } = await this.$axios.post("/admin/addSubject", {
                    name: this.newSubjectForm.name, //假设你的新科目名称存在这个变量里
                });
                console.log(this.newSubjectForm.name)
                if (res.status === 101) {
                    this.$message.success("新增科目成功");
                    this.getSubjectsList();  //刷新科目列表
                    this.newSubjectForm.name = ""
                } else {
                    this.$message.error("新增科目失败");
                }
            } catch (err) {
                this.$message.error("新增科目出错");
            }
        },
        //删除科目
        async deleteSubject() {
            try {
                const { data: res } = await this.$axios.delete("/admin/deleteSubject", {
                    params: { name: this.deleteSubjectForm.name }  //假设你要删除的科目名称存在这个变量里
                });
                if (res.status === 401) {
                    this.$message.success("删除科目成功");
                    this.getSubjectsList();  //刷新科目列表
                } else {
                    this.$message.error("删除科目失败");
                }
            } catch (err) {
                this.$message.error("删除科目出错");
            }
        },
        //增加章节
        async addChapter() {
            try {
                const { data: res } = await this.$axios.post("/admin/addChapter", {
                    subjectName: this.newChapterForm.subject,  //假设你要在哪个科目下新增章节的科目名称存在这个变量里
                    chapterName: this.newChapterForm.chapter,   //假设你的新章节名称存在这个变量里
                    description: this.newChapterForm.description
                });
                if (res.status === 101) {
                    this.$message.success("新增章节成功");
                    this.getChaptersList(this.newChapterForm.subject);  //刷新章节列表
                    this.newChapterForm.subject = "",  //假设你要在哪个科目下新增章节的科目名称存在这个变量里
                        this.newChapterForm.chapter = "",   //假设你的新章节名称存在这个变量里
                        this.newChapterForm.description = ""

                } else {
                    this.$message.error("新增章节失败");
                }
            } catch (err) {
                this.$message.error("新增章节出错");
            }
        },
        //删除章节
        async deleteChapter() {
            try {
                const { data: res } = await this.$axios.delete("/admin/deleteChapter", {
                    params: {
                        subjectName: this.deleteChapterForm.subject, //假设你要在哪个科目下删除章节的科目名称存在这个变量里
                        chapterName: this.deleteChapterForm.chapter  //假设你要删除的章节名称存在这个变量里
                    }
                });
                if (res.status === 401) {
                    this.$message.success("删除章节成功");
                    this.getChaptersList(this.deleteChapterForm.subject);  //刷新章节列表
                    console.log(this.deleteChapterForm.subject)
                } else {
                    this.$message.error("删除章节失败");
                }
            } catch (err) {
                this.$message.error("删除章节出错");
            }
        }
    }
}
</script>

<style scoped>
.admin-panel {
    width: auto;
    margin: 0 auto;
}
</style>
