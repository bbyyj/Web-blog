<template>
    <div class="box" @click="navigateToCourseDetail">
        <div class="course-box">
            <!--课程编号-->
            <div class="course-title">
                <a>
                    {{ courseCode }}
                </a>
            </div>
            <!--课程名称-->
            <div class="course-info">
                {{ courseName }}
            </div>
            <!--教师名称-->
            <div class="teacher-info">
                {{ teacher }}
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "CourseCard",
    props: ["courseCode", "courseName", "teacher"],
    methods: {
        navigateToCourseDetail() {
            //这里你可以导航到课程的详细信息页面
            console.log(this.courseCode)
            this.$router.push({ path: '/CourseDetail', query: { subject_id: this.courseCode } });
        }
    }
}
</script>

<style scoped>
/* 这里我仅保留了一部分的CSS样式，并进行了一些修改 */
.box {
    background-color: #d9dbeb;
    width: 840px;
    padding: 16px 20px;
    overflow: hidden;
    border-radius: 0.25rem;
    margin-bottom: 20px;
    position: relative;
    transition: all .5s;
    cursor: pointer;
}

.box:hover {
    transform: scale(1.02);
}

.box .course-box {
    width: 450px;
    height: 190px;
    padding-left: 20px;
}

.box .course-box .course-title {
    cursor: pointer;
    height: 46px;
    line-height: 46px;
    text-align: center;
    font-size: 2rem;
}

.box .course-box .course-title a {
    font-weight: 500;
    text-decoration: none;
    color: #3d3952;
    padding-bottom: 5px;
    position: relative;
}

.box .course-box .course-title:hover a {
    color: #7c80b8;
}

.box .course-box .course-info,
.box .course-box .teacher-info {
    height: 25px;
    text-align: center;
    margin-top: 5px;
    font-size: 20px;
    font-weight: 700;
    color: #696eae;
}
</style>
