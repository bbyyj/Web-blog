<template>
    <div class="bg">
        <div class="container">
            <el-select class="top-select" popper-class="select-down" v-model="selectedCategory" placeholder="选择课程类别" @change="fetchCourses" clearable>
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
                    <el-pagination popper-class="select-down" @size-change="handleSizeChange" @current-change="handleCurrentChange"
                        :current-page="queryInfo.pageNum" :page-sizes="[5, 10]" :page-size="queryInfo.pageSize"
                        layout="total, sizes, prev, pager, next, jumper" :total="total">
                    </el-pagination>
                </div>
            </transition>
        </div>
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
<style lang="less" scoped>
.bg{
    background: linear-gradient(-45deg, #f5e2e0, #c2bbce, #837cb8, #3d3952);
    background-size: 300% 300%;
    animation: gradient 8s ease-in-out infinite;
    min-height: 1600px;
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
.container {
    flex-direction: column;
    align-items: center;
    justify-content: center;
    font-family: HYYueXiChuKaiJ-2;
}
.box{
    margin-left: 180px;
    margin-bottom: 50px;
    width: 800px;
}
.top-select{
    margin: 100px 100px 60px 50px;
}
/deep/.el-dropdown-menu {
    background-color: #eff1f6;
    border-radius: 10px;
    padding: 15px 20px 15px 20px;
    font-family: HYYueXiChuKaiJ-2;
}
/deep/ .el-dropdown-menu__item:hover{
    background-color:#e0e5ed;
    color: #7378ac;
    border-radius: 5px;
    padding: 5px 10px 5px 10px;
}

/deep/ .el-pagination__total{
    color: #eff1f6;
    font-size: 20px;
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
    border: 2px solid  #565ca4 !important; 
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
/deep/ .el-pagination__jump{
    color: #eff1f6;
    font-size: 16px;
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