
<template>
    <div class="ChooseQS">
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
                <Pagination class="pagebar" @jumpPage="jumpPage" :pageInfo="{ pageNum: queryInfo.pageNum, pages: pages }">
                </Pagination>
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
        this.getChaptersList('计算机组成原理');
    },

    methods: {
        //获取科目列表
        async getSubjectsList(flag) {
            const { data: res } = await this.$axios.get("/myblog/subjectList");
            if (res.status === 1) {
                //将res赋值给subjects
                console.log(res.data[0]);
                this.subjects = res.data[0];
            }
        },
        //获取科目下的章节内容
        async getChaptersList(subject) {
            const { data: res } = await this.$axios.get("/myblog/chapterList", {
                params: { name: subject }
            });
            if (res.status === 1) {
                this.Chapters = res.data[0];
            } else {
                this.$message.error("获取对应章节失败,请重试")
                return
            }
            //分页相关
            const count = res.data.length > 1 ? res.data[1] : 0
            this.pages = Math.ceil(count / this.queryInfo.pageSize);
            if (this.pages <= 0) {
                this.pages = 1
            }
        },
        jumpPage(pageNum) {
            window.scrollTo(0, 0)
            this.queryInfo.pageNum = pageNum;
            this.getChaptersList(this.currentTagId);
        },
        GotoTest(typename, title) {
            this.$router.push({
                path: "/StudyTest",
                query: {
                    typename: typename,
                    title: title
                }
            });
        }
    }
}

</script>



<style lang="less" scoped>
//下拉菜单的样式
.el-dropdown {
    vertical-align: top;
}

.el-dropdown+.el-dropdown {
    margin-left: 15px;
}

.el-icon-arrow-down {
    font-size: 12px;
}

ul,
li {
    margin: 0;
    padding: 0;
}

.ChooseQS {
    background-attachment: fixed;
    min-height: 1000px;
}

.title {
    font-size: 450%;
    color: #ffffff;
    margin-bottom: 50px;
    bottom: 0 !important;
    right: 0 !important;
    font-family: 'STXingkai';
    opacity: 0.5;
    padding-top: 6%;
}

.subjects-area {
    width: 840px;
    margin: 0 auto;
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