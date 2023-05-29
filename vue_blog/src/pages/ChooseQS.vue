
<template>
    <div class="ChooseQS">
        <div class="sun">
        </div>

        <div class="title" align="center">答题测试</div>
        <!-- 科目栏 -->
        <transition appear name="animate__animated animate__bounce animate__slow" enter-active-class="animate__fadeInDown"
            leave-active-class="animate__fadeOutUp">
            <ul class="subjects-area">
                <el-dropdown>
                    <el-button type="primary">
                        更多科目<i class="el-icon-arrow-down el-icon--right"></i>
                    </el-button>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item v-for="subject in subjects" :key="subject.id"
                            @click.native="getChaptersList(subject.name)">
                            {{ subject.name }}
                        </el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
                <div class="space"></div>
                <!--标识每个科目-->
                <li :key="item.id" v-for="item in subjects">
                    <!--传入对应科目的名字 获取全部的章节-->
                    <BlogTagItem @click.native="getChaptersList(item.name)" :taginfo="item" :activeId="currentTagId">
                    </BlogTagItem>
                </li>
            </ul>
        </transition>

        <!-- 章节列表栏 -->
        <transition appear name="animate__animated animate__bounce animate__slow" enter-active-class="animate__fadeInUp">
            <div>
                <div class="chapters-area">
                    <!--标识每个科目-->
                    <ChaptersCard class="chapter-item" :key="item.id" v-for="(item, index) in Chapters" :item="item"
                        :imgRight="index % 2 === 0" @click.native="GotoTest(item.typename, item.title)">
                        <!--blogDetail跳转到对应的答题页面-->
                    </ChaptersCard>
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

import ChaptersCard from "../components/ChaptersCard";
import Pagination from "../components/Pagination";
import BlogTagItem from "../components/BlogTagItem";

export default {
    components: { ChaptersCard, Pagination, BlogTagItem },
    data() {
        return {
            currentTagId: 1,
            //选中的科目
            subject: "",
            //对应科目下的章节数量
            total: 0,
            // 页面数量
            pages: 1,
            queryInfo: {
                pageNum: 1,
                pageSize: 5,
                tagId: 0
            },
            //科目选择
            subjects: [],
            //对应科目下的章节等内容
            Chapters: [],
        }
    },
    created() {
        this.getSubjectsList(true);
        this.getChaptersList('计算机网络');
    },
    methods: {
        //获取科目列表
        async getSubjectsList(flag) {
            const { data: res } = await this.$axios.get("/myblog/subjectList");
            if (res.status === 1) {
                //将res赋值给subjects
                //console.log(res.data[0]);
                this.subjects = res.data[0];
            }
        },
        //获取科目下的章节内容
        async getChaptersList(subject) {
            const { data: res } = await this.$axios.get("/myblog/chapterList", {
                params: { name: subject, pageNum: this.queryInfo.pageNum, pageSize: this.queryInfo.pageSize }
            });
            console.log(res.data[0])
            if (res.status === 1) {
                this.subject = subject;
                this.Chapters = res.data[0];
                for (let item of this.subjects) {
                    if (item.name === subject) {
                        this.total = item.count;
                        break;
                    }
                }
            } else {
                this.$message.error("获取对应章节失败,请重试")
                return
            }
            //分页相关
            this.pages = Math.ceil(this.total / this.queryInfo.pageSize);
            if (this.pages <= 0) {
                this.pages = 1
            }
        },
        GotoTest(typename, title) {
            this.$router.push({
                path: "/StudyTest",
                query: {
                    typename: typename,
                    title: title
                }
            });
        },
        handleSizeChange(newSize) {
            this.queryInfo.pageSize = newSize;
            this.getChaptersList(this.subject);
        },
        handleCurrentChange(newPage) {
            this.queryInfo.pageNum = newPage;
            this.getChaptersList(this.subject);
        },
    }
}

</script>

<style lang="less" scoped>
.ChooseQS {
    min-height: 2000px;
    background-attachment: fixed;
    background-repeat: no-repeat;
    background-color: #7378ac;
    background-image:
        radial-gradient(closest-side, #aab7cc, rgba(235, 105, 78, 0)),
        radial-gradient(closest-side, #7285ac, rgba(243, 11, 164, 0)),
        radial-gradient(closest-side, #919fbe, rgba(254, 234, 131, 0)),
        radial-gradient(closest-side, #e0e5ed, rgba(170, 142, 245, 0)),
        radial-gradient(closest-side, #8295b5, rgba(248, 192, 147, 0));
    background-size: 130vmax 130vmax, 80vmax 80vmax, 90vmax 90vmax, 110vmax 110vmax, 90vmax 90vmax;
    background-position: -80vmax -80vmax, 60vmax -30vmax, 10vmax 10vmax, -30vmax -10vmax, 50vmax 50vmax;
    animation: 10s move linear infinite;

}
// 通过修改background的参数形成动画
@keyframes move {

    0%,
    100% {
        background-size: 130vmax 130vmax, 80vmax 80vmax, 90vmax 90vmax, 110vmax 110vmax, 90vmax 90vmax;
        background-position: -80vmax -80vmax, 60vmax -30vmax, 10vmax 10vmax, -30vmax -10vmax, 50vmax 50vmax;
    }

    25% {
        background-size: 100vmax 100vmax, 90vmax 90vmax, 100vmax 100vmax, 90vmax 90vmax, 60vmax 60vmax;
        background-position: -60vmax -90vmax, 50vmax -40vmax, 0vmax -20vmax, -40vmax -20vmax, 40vmax 60vmax;
    }

    50% {
        background-size: 80vmax 80vmax, 110vmax 110vmax, 80vmax 80vmax, 60vmax 60vmax, 80vmax 80vmax;
        background-position: -50vmax -70vmax, 40vmax -30vmax, 10vmax 0vmax, 20vmax 10vmax, 30vmax 70vmax;
    }

    75% {
        background-size: 90vmax 90vmax, 90vmax 90vmax, 100vmax 100vmax, 90vmax 90vmax, 70vmax 70vmax;
        background-position: -50vmax -40vmax, 50vmax -30vmax, 20vmax 0vmax, -10vmax 10vmax, 40vmax 60vmax;
    }
}
.sun {
    position: fixed;
    top: 80px;
    left: 780px;
    background-image: url('../assets/images/sun.svg');
    opacity: 0.3;
    background-repeat: no-repeat;
    height: 400px;
    width: 400px;
}
.title {
    font-size: 80px;
    color: #eff1f6;
    top: 50px;
    bottom: 0 !important;
    right: 0 !important;
    font-family: 'STXingkai';
    opacity: 0.85;
    padding-top: 8%;
}
ul,li {
    display: block;
    list-style: none;
    margin: 0;
    padding: 0;
}
.space{
    height: 50px;
}

.el-dropdown {
    vertical-align: top;
}

.el-icon-arrow-down {
    font-size: 14px;
}

.el-button--primary {
    background-color: #393f51;
    padding: 18px;
    border-radius: 10px;
    border: 0;
}

.el-dropdown-menu {
    background-color: #eff1f6;
    border-radius: 10px;
    padding: 15px 20px 15px 20px;
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
}
/deep/ .el-input__inner:focus, .el-input__inner:hover, .el-input__inner:active{
    border: 2px solid  #565ca4;
}
// 修改不成功
/deep/ .el-select-dropdown__list{
    .el-select-dropdown__item {
		padding: 0 0.2rem 0 0.2rem;
		color:#010101;
		font-size: 16px;
	}
    .el-select-dropdown__item.selected {
		color: #565ca4;
	}
    .el-select-dropdown__item.hover,.el-select-dropdown__item:hover {
		background-color: #ffffff;
	}
}

/deep/ .el-select:hover{
    .el-input__inner:hover{
        border: 2px solid  #565ca4;
    }
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
    padding-left: 500px;
}
.subjects-area {
    margin: -10px 200px 50px 200px;
    overflow: hidden;
}

.chapters-area {
    width: 850px;
    margin: 40px auto 50px;
}

.pagebar {
    padding-bottom: 50px;
}
</style>