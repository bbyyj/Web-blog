<template>
    <div class="container">
        <el-select v-model="selectedCategory" placeholder="选择课程类别" @change="fetchCourses" clearable>
            <el-option v-for="item in categories" :key="item" :label="item" :value="item">
            </el-option>
        </el-select>
        <!-- 章节列表栏 -->
        <transition appear name="animate__animated animate__bounce animate__slow" enter-active-class="animate__fadeInUp">
            <div>
                <div>
                    <!-- CourseCard组件的渲染位置 -->
                    <div v-for="(course, index) in courses" :key="index">
                        <CourseCard :courseCode="course.subject_id" :courseName="course.subject_name"
                            :teacher="course.teacher" />
                    </div>
                </div>
                <!--分页-->
                <!-- 分页区域 -->
                <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                    :current-page="queryInfo.pageNum" :page-sizes="[5, 10]" :page-size="queryInfo.pageSize"
                    layout="total, sizes, prev, pager, next, jumper" :total="total">
                </el-pagination>
            </div>
        </transition>
    </div>
</template>

<script>
import axios from 'axios';
import CourseCard from "../components/CourseCard";  // 确保路径正确

export default {
    components: {
        CourseCard
    },
    data() {
        return {
            selectedCategory: '',
            categories: ['公必', '公选', '专必', '专选'],
            courses: [],
            total: 0,
            // 页面数量
            pages: 1,
            queryInfo: {
                pageNum: 1,
                pageSize: 5,
            },
        }
    },
    created() {
        //初始化
        this.fetchAllCourses();
    },
    methods: {
        async fetchCourses() {
            if (!this.selectedCategory) return;  // 没有选择分类则不发送请求
            try {
                const { data: res } = await this.$axios.get("/myblog/electionList", {
                    params: {
                        classification: this.selectedCategory,
                        pageNum: this.queryInfo.pageNum,
                        pageSize: this.queryInfo.pageSize
                    }
                });
                this.courses = res.data[0]; // 根据你的后端返回的数据结构进行修改
                this.total = res.data[1]
            } catch (error) {
                console.error(error);
            }
        },
        async fetchAllCourses() {
            try {
                const { data: res } = await this.$axios.get("myblog/electionListNoClass", {
                    params: {
                        pageNum: this.queryInfo.pageNum,
                        pageSize: this.queryInfo.pageSize
                    }
                });
                this.courses = res.data[0]; // 根据你的后端返回的数据结构进行修改
                this.total = res.data[1]
            } catch (error) {
                console.error(error);
            }
        },
        ///////////////////////分页相关的
        handleSizeChange(newSize) {
            this.queryInfo.pageSize = newSize;
            this.fetchCourses()
        },
        handleCurrentChange(newPage) {
            this.queryInfo.pageNum = newPage;
            this.fetchCourses()
        },
    }
}
</script>
<style scoped>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    width: 100vw;
    box-sizing: border-box;
    padding: 10px;
}
</style>