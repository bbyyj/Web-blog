
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
            currentTagId: 0,
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
            subjects: [{
                id: 1,
                name: "计算机网络",
                count: 5
            },
            {
                id: 2,
                name: "计算机组成原理",
                count: 8
            },
            {
                id: 3,
                name: "操作系统",
                count: 5
            },
            {
                id: 4,
                name: "数据结构与算法",
                count: 7
            },
            {
                id: 5,
                name: "数据结构与算法",
                count: 7
            },
            {
                id: 6,
                name: "数据结构与算",
                count: 7
            }
            ],
            //对应科目下的章节等内容
            Chapters: [{
                id: 1,
                title: "传输层",
                description: "TCP UDP",
                typename: "计算机网络",
                views: 98
            },
            {
                id: 2,
                title: "文件系统",
                description: "LFS",
                typename: "操作系统",
                views: 10
            },
            {
                id: 3,
                title: "树",
                description: "二叉树",
                typename: "数据结构与算法",
                views: 20
            }],
        }
    },
    created() {
        this.getSubjectsList(true);
    },

    methods: {
        //获取科目列表
        async getSubjectsList(flag) {
            const { data: res } = await this.$axios.get("/myblog/subjectList");
            if (res.status === 1) {
                //将res赋值给subjects
                console.log(res.data[0]);
                this.subjects = res.data[0].data;
                console.log(res.data);
            }
            //调用章节获取的接口
            /*if (flag) {
                await this.getChaptersList(this.$route.query.id);
                //await this.getChaptersList();
            } else {
                await this.getChaptersList(this.subjects[0].id)
            }*/
        },
        //获取科目下的章节内容

        async getChaptersList(subject) {
            //this.currentTagId = id;
            //this.queryInfo.tagId = id;
            const { data: res } = await this.$axios.get("/myblog/章节接口", { params: this.subject });
            if (res.status === 1) {
                this.Chapters = res.data.length > 0 ? res.data[0] : this.Chapters;
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
            this.getBlogList(this.currentTagId);
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