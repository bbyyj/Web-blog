
<template>
    <div>
        <!--上方背景页面-->
        <div class='console-container'>
            <span id='text'></span>
            <div class='console-underscore' ref="text-style">&#95;</div>
        </div>

        <div id="bg1" ref="box">
            <span>Blog</span>
            <div class="scroll-down">
                <a class="anchor-down" @click.prevent="anchorDown"></a>
            </div>
        </div>

        <!--博客展示区-->
        <div id="area-blow">
            <div id='stars'></div>
            <div id='stars2'></div>
            <div id='stars3'></div>
            
            <div class="cloud-bg01"></div>
            <div class="cloud-bg02"></div>
            <TagCloud class="tag-detail"></TagCloud>
            <TypeArea class="type-detail"></TypeArea>
            

            <div class="blog-area">
                <!--用户信息栏-->
                <UserInfoCard :count="blogCount"></UserInfoCard>

                <div id="player"></div>
                <BlogCard :key="item.id" v-for="(item, index) in blogDetails" :item="item" :imgRight="index % 2 === 0"
                    @click.native="blogDetail(item.id)">
                </BlogCard>
                <!--分页导航区-->
                <Pagination @jumpPage="jumpPage" :pageInfo="{ pageNum: queryInfo.pageNum, pages: pages }"></Pagination>
            </div>
        </div>

    </div>
</template>


<script>
import img1 from "../assets/1.png"
import img2 from "../assets/2.jpg"
import img3 from "../assets/3.jpg"

import BackToTop from "../components/BackToTop";
import BlogCard from "../components/BlogCard";
import UserInfoCard from "../components/UserInfoCard";
import Pagination from "../components/Pagination";
import TagCloud from "../components/TagCloud";
import TypeArea from "../components/TypeArea";

import 'APlayer/dist/APlayer.min.css';
import APlayer from 'APlayer';

import "animate.css"
import '../assets/css/star.css'
import { set } from "vue";
export default {

    name: "Home",
    components: { BackToTop, BlogCard, Pagination, UserInfoCard, TagCloud, TypeArea },
    data() {
        return {
            firstBGPageInfo: {
                curBg: "",
                bgs: [],
                mottos: [],
            },
            showCat: false,
            showMotto: true,
            pages: 1,              // 页面数量
            blogCount: 66,
            queryInfo: {
                pageNum: 1,
                pageSize: 8
            },
            blogDetails: [{
                id: 1,
                title: "博客系统开发",
                description: "本章主要介绍博客系统的开发工作...",
                firstPicture: img1,
                createTime: "2021-03-29",
                views: 1024,
                nickname: "Jack Ma",
                typename: "框架底层原理"
            },
            {
                id: 2,
                title: "博客系统开发",
                description: "本章主要介绍博客系统的开发工作...",
                firstPicture: img2,
                createTime: "2021-03-29",
                views: 1024,
                nickname: "Jack Ma",
                typename: "框架底层原理"
            },
            {
                id: 3,
                title: "博客系统开发",
                description: "本章主要介绍博客系统的开发工作...",
                firstPicture: img3,
                createTime: "2021-03-29",
                views: 1024,
                nickname: "Jack Ma",
                typename: "框架底层原理"
            }],
            newRecommend: {
                title: "最新推荐",
                icon: "icon-zuixingengxin",
                list: [
                    { id: 1, title: "MarkDownit使用" },
                    { id: 2, title: "Vue日期格式化过滤器" },
                    { id: 3, title: "前后端分离登录验证" },
                    { id: 4, title: "MavonEditor上传图片" },
                    { id: 5, title: "Springboot中PageHelper 分页查询使用方法" },
                    { id: 6, title: "mybatis根据日期查询、查询日期并返回" },
                    { id: 7, title: "maven中静态资源的过滤" }
                ]
            },
            hotBlogs: {
                title: "热门推荐",
                icon: "icon-fire",
                list: [
                    { id: 1, title: "MarkDownit使用" },
                    { id: 2, title: "Vue日期格式化过滤器" },
                    { id: 3, title: "前后端分离登录验证" },
                    { id: 4, title: "MavonEditor上传图片" },
                    { id: 5, title: "Springboot中PageHelper 分页查询使用方法" },
                    { id: 6, title: "mybatis根据日期查询、查询日期并返回" },
                    { id: 7, title: "maven中静态资源的过滤" }
                ]
            },
            //获得歌单
            musicList: []
        }
    },
    created() {
        this.getBlogLists();
        this.getNewBlogs()
        this.getHotBlogs()
        this.getmusicList()

    },
    mounted() {
        //音乐播放器
        setTimeout(() => {
            // 在这里放入需要延迟执行的代码
            // 音乐播放器等相关操作
            const ap = new APlayer({
                container: document.getElementById('player'),
                fixed: true,
                audio: this.musicList
            });
        }, 2000); // 延迟1秒执行


        document.addEventListener('scroll', this.parallax, true);
        this.createConsoleText();
    },
    deactivated() {
        this.showMotto = false
    },
    activated() {
        this.showMotto = true
        window.scrollTo(0, 0)
    },

    methods: {
        // 获取音乐列表
        async getmusicList() {
            const { data: res } = await this.$axios.get("/myblog/getAllMusic")
            if (res.status === 1) {
                console.log(res.data[0])
                this.musicList = res.data[0];
            }
            else {
                window.alert("获取音乐列表失败,请检查网络设置")
            }
        },

        anchorDown() {
            const offsetTop = document.getElementById("area-blow").offsetTop;
            window.scrollTo({ top: offsetTop, behavior: 'smooth' })
        },
        // 获取博客列表
        async getBlogLists() {
            var keyWord = window.sessionStorage.getItem("keyWord");
            if (keyWord) {
                window.sessionStorage.removeItem("keyWord");
                const { data: res } = await this.$axios.get("/myblog/search", { params: { pageNum: this.queryInfo.pageNum, pageSize: this.queryInfo.pageSize, keyWord: keyWord } });
                if (res.status === 1) {
                    this.blogDetails = res.data.length > 0 ? res.data[0] : this.blogDetails
                    this.blogCount = res.data.length > 1 ? res.data[1] : this.blogCount
                    this.pages = Math.ceil(this.blogCount / this.queryInfo.pageSize)
                    if (this.pages <= 0) {
                        this.pages = 1
                    }
                } else {
                    this.$message.error("获取博客失败，请重试")
                }
            } else {
                const { data: res } = await this.$axios.get("/myblog/blogLists", { params: this.queryInfo });
                if (res.status === 1) {
                    this.blogDetails = res.data.length > 0 ? res.data[0] : this.blogDetails
                    this.blogCount = res.data.length > 1 ? res.data[1] : this.blogCount
                    this.pages = Math.ceil(this.blogCount / this.queryInfo.pageSize)
                    if (this.pages <= 0) {
                        this.pages = 1
                    }
                } else {
                    this.$message.error("获取博客失败，请重试")
                }
            }
        },
        // 获取最新推荐博客
        async getNewBlogs() {
            const { data: res } = await this.$axios.get("/myblog/newBlogs", { params: { countLimit: 10 } });
            if (res.status === 1) {
                this.newRecommend.list = res.data.length > 0 ? res.data[0] : this.newRecommend.list;
            } else {
                this.$message.warning("获取最新推荐失败！")
            }
        },
        // 获取热门推荐博客
        async getHotBlogs() {
            const { data: res } = await this.$axios.get("/myblog/hotBlogs", { params: { countLimit: 10 } });
            if (res.status === 1) {
                this.hotBlogs.list = res.data.length > 0 ? res.data[0] : this.hotBlogs.list;
            } else {
                this.$message.warning("获取热门推荐失败！")
            }
        },
        // 进入博客内容页面
        blogDetail(id) {
            this.$router.push({
                path: "/blogDetail",
                query: {
                    id: id
                }
            });
        },
        jumpPage(pageNum) {
            this.queryInfo.pageNum = pageNum;
            this.getBlogLists();
        },
        //滚轮视差
        parallax() {
            const scrollY = window.scrollY
            if (scrollY !== 0) {
                this.$refs.box.style.backgroundPosition = `calc(50% + ${scrollY}px) calc(50% + ${scrollY}px)`
            }
        },

        createConsoleText() {
            var words = ['Hello World.', '不知道写什么', '也不知道选什么颜色', '可能还要换个字体', '先意思一下', '不过紫色还蛮好看'];
            var colors = ['#9990bc', '#8a84b7', '#7b79b1', "#6c6dac", "#5d62a7"];

            var visible = true;
            var con = document.getElementById('console');
            var letterCount = 1;
            var x = 1;
            var waiting = false;
            var target = document.getElementById('text');
            target.setAttribute('style', 'color:' + colors[0]);
            this.consoleText(words, colors, visible, con, letterCount, x, waiting, target);
        },

        consoleText(words, colors, visible, con, letterCount, x, waiting, target) {
            if (colors === undefined) colors = ['#black'];

            target.setAttribute('style', 'color:' + colors[0])
            window.setInterval(function () {
                if (letterCount === 0 && waiting === false) {
                    waiting = true;
                    target.innerHTML = words[0].substring(0, letterCount)
                    window.setTimeout(function () {
                        var usedColor = colors.shift();
                        colors.push(usedColor);
                        var usedWord = words.shift();
                        words.push(usedWord);
                        x = 1;
                        target.setAttribute('style', 'color:' + colors[0])
                        letterCount += x;
                        waiting = false;
                    }, 1000)
                } else if (letterCount === words[0].length + 1 && waiting === false) {
                    waiting = true;
                    window.setTimeout(function () {
                        x = -1;
                        letterCount += x;
                        waiting = false;
                    }, 1000)
                } else if (waiting === false) {
                    target.innerHTML = words[0].substring(0, letterCount)
                    letterCount += x;
                }
            }, 120)
        },
    }
}
</script>




<style lang="less" scoped>
.animate__animated {
    animation-duration: 3s;
}

.console-container {
    font-size: 3em;
    text-align: center;
    height: 560px;
    width: 1000px;
    display: block;
    position: absolute;
    color: #fffcfb;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
}
.console-underscore {
    display: inline-block;
    position: relative;
    top: -0.05em;
    left: 10px;
}
.scroll-down {
    width: 100%;
    height: 60px;
    position: relative;
    top: 75%;
    text-align: center;

    animation: bounce-in 5s 1s infinite;
}

.anchor-down {
    display: block;
    width: 36px;
    height: 30px;
    transform: rotate(45deg);
    position: absolute;
    left: 50%;
    margin-left: -18px;
    cursor: pointer;
}

.anchor-down::after {
    content: '';
    width: 28px;
    height: 28px;
    position: absolute;
    right: 0;
    bottom: 0;
    border-right: 5px solid #fbd5d1;
    border-bottom: 5px solid #fbd5d1;
    display: block;
    border-bottom-right-radius: 5px;
}

@keyframes bounce-in {
    0% {
        transform: translateY(0);
    }
    50% {
        transform: translateY(-20px);
    }
    100% {
        transform: translateY(0);
    }
}

//上方背景
#bg1 {
    background-image: url("../assets/images/forest04.svg");
    background-size: cover;
    background-position: 50% 50%;
    font: 600 25rem '';
    line-height: 95vh;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    position: relative;
    text-align: center;
    overflow: hidden;

}

#bg1::before {
    content: '';
    background-size: cover;
    background-image: inherit;
    background-position: 50% 50%;
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    z-index: -100;
}

//下面区域
#area-blow {
    background-color: #1d1d2b;
}

// 下面中心区域
.blog-area {
    margin: 20px 20px 20px 20px;
    padding-top: 36px;
    padding-bottom: 64px;
    overflow: hidden;
    min-height: 2000px;
}

.cloud-bg01{
    position: absolute;
    top: 1250px;
    right: 0;
    background-image: url('../assets/images/cloud03.png');
    opacity: 0.48;
    background-repeat: no-repeat;
    height: 500px;
    width: 800px;
}
.tag-detail{
    position: absolute;
    top: 1350px;
    right: 20px;
}
.type-detail{
    position: absolute;
    top: 1650px;
    left: 20px;
}
.cloud-bg02{
    position: absolute;
    top: 1550px;
    left: 0;
    background-image: url('../assets/images/cloud04.png');
    opacity: 0.58;
    background-repeat: no-repeat;
    height: 520px;
    width: 700px;
}
</style>
