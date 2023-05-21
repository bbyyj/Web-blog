
<template>

    <div>
        <TitleArea class="title-area" :info="info"></TitleArea>        
        <div class="main">
            <transition appear name="animate__animated animate__bounce animate__slow"
                        enter-active-class="animate__fadeInDown" leave-active-class="animate__fadeOutUp">
                <div class="center-area">
                    <!-- 左边时间线 -->
                    <div class="timeline hvr-grow">
                        <div class="article-title">文章列表</div>
                        <div class="article-list ">
                            <TimeAxis :key="index" v-for="(item, index) in blogs" :list="item"></TimeAxis>
                        </div>
                    </div>
                    <!-- 右侧边栏 -->
                    <div class="right-area ">
                        <BlogTypePie class="hvr-grow"></BlogTypePie>
                        <TagCloud class="tag-cloud hvr-grow"></TagCloud>
                    </div>
                </div>
            </transition>
        </div>
    </div>

</template>



<script>

import "../assets/icon/iconfont.css"

import TitleArea from "../components/TitleArea";
import TimeAxis from "../components/TimeAxis";
import BlogTypePie from "../components/BlogTypePie";
import TagCloud from "../components/TagCloud";

export default {
    components: { TitleArea, TimeAxis, BlogTypePie, TagCloud },
    data() {
        return {
            info: {
                title: "文章归档",
            },
            blogsTotal: 3,
            blogs: [
                    [{
                    id: 1,
                    title: "博客系统开发",
                    flag: "原创",
                    createTime: "2021-04-30T13:20:34.000+00:00"
                }],
                [{
                    id: 2,
                    title: "博客系统开发",
                    flag: "原创",
                    createTime: "2021-03-30T13:20:34.000+00:00"
                }],
                [{
                    id: 3,
                    title: "博客系统开发",
                    flag: "原创",
                    createTime: "2020-03-30T13:20:34.000+00:00"
                }]
            ],
        }
    },
    created() {
        window.scrollTo(0, 0)
        this.getBlogList();
    },
    methods: {
        getBlogList: async function() {
            const {data: res} = await this.$axios.get("/myblog/timeLine");
            this.blogs = res.data.length > 0 ? res.data[0] : this.blogs;
            this.blogsTotal = 0;
            for(let i = 0; i < this.blogs.length; i++) {
                this.blogsTotal += this.blogs[i].length;
            }
        },
    },
}
</script>



<style lang="less" scoped>

ul, li {
    margin: 0;
    padding: 0;
}

//标题区域
.title-area {
   background-color: #3d3952;
   
}

//下方整个区域
.main {
    min-height: 800px;
    width: 100%;
    background-image: url("../assets/images/wavetest.svg");
    background-position-y: 45ch;
    background-repeat: no-repeat;
    background-color: #3d3952;
    padding-top: 50px;
}

//下方中心区域
.center-area {
    width: 1200px;
    margin: 0 auto;
}

.timeline {
    width: 840px;
    background-color: #f7d2f9;
    padding: 30px 0;
    border-radius: 6px;
    float: left;
}

.article-title {
    font-size: 36px;
    font-weight: 500;
    font-family: '华文中宋';
    text-align: center;
}


.right-area {
    width: 320px;
    float: left;
    margin-left: 26px;
}

.tag-cloud {
    width: 300px;
    margin: 24px 0 0;
}

.article-list {
    padding: 30px 0 30px 90px;
}

.hvr-grow {
  display: inline-block;
  vertical-align: middle;
  -webkit-transform: perspective(1px) translateZ(0);
  transform: perspective(1px) translateZ(0);
  box-shadow: 0 0 1px rgba(0, 0, 0, 0);
  -webkit-transition-duration: 0.3s;
  transition-duration: 0.3s;
  -webkit-transition-property: transform;
  transition-property: transform;
}
.hvr-grow:hover, .hvr-grow:focus, .hvr-grow:active {
  -webkit-transform: scale(1.05);
  transform: scale(1.05);
}



</style>
