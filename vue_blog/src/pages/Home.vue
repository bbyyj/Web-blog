
<template>
    <div>
        <!--上方背景页面-->
        <div id="bg1" ref="box">
            <span ref="space">Blog</span>
            <!-- <div id="motto" v-if="showMotto">
                <transition-group appear name="animate__animated animate__bounce animate__slow"
                    enter-active-class="animate__bounceIn" leave-active-class="animate__bounceOut">
                    <h1 key="1">{{ firstBGPageInfo.curMotto.ch }}</h1>
                    <p key="2">{{ firstBGPageInfo.curMotto.en }}</p>
                </transition-group>
            </div> -->
            <div class="scroll-down">
                <a class="anchor-down" @click.prevent="anchorDown"></a>
            </div>
        </div>


        <!--博客展示区-->
        <div id="area-blow">
            <div class="blog-area">
                <!--左侧博客信息区-->
                <div class="blog-left">
                    <!--用户信息栏-->
                    <UserInfoCard :count="blogCount"></UserInfoCard>
                    <!--最新推荐栏-->
                    <RecommendList :blogList="newRecommend"></RecommendList>
                    <!--标签云-->
                    <TagCloud></TagCloud>
                    <!--热门推荐-->
                    <RecommendList :blogList="hotBlogs"></RecommendList>

                </div>
                <div id="player"></div>
                <!--右侧博客展示区-->
                <div class="blog-right">
                    <BlogCard :key="item.id" v-for="(item, index) in blogDetails" :item="item" :imgRight="index % 2 === 0"
                        @click.native="blogDetail(item.id)">
                    </BlogCard>
                    <!--分页导航区-->
                    <Pagination @jumpPage="jumpPage" :pageInfo="{ pageNum: queryInfo.pageNum, pages: pages }"></Pagination>

                </div>
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
import RecommendList from "../components/RecommendList";
import TagCloud from "../components/TagCloud";

import 'APlayer/dist/APlayer.min.css';
import APlayer from 'APlayer';

import "animate.css"

export default {

    name: "Home",
    components: { BackToTop, BlogCard, Pagination, UserInfoCard, RecommendList, TagCloud },
    data() {
        return {
            firstBGPageInfo: {
                curBg: "",
                bgs: [],
                curMotto: {
                    ch: "奋斗从未停止, 前进永无止境!",
                    en: "Struggle never stops, and progress never ends!"
                },
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
            musicList: [
                {
                    id: 1,
                    name: '',
                    artist: '',
                    url: '',
                    cover: '',
                    theme: '#ebd0c2'
                },
                {
                    id: 2,
                    name: '',
                    artist: '',
                    url: '',
                    cover: '',
                    theme: '#ebd0c2'
                },
                {
                    id: 3,
                    name: '',
                    artist: '',
                    url: '',
                    cover: '',
                    theme: '#ebd0c2'
                }
            ]
        }
    },
    created() {
        this.getBlogLists();
        this.getMotto()
        this.getNewBlogs()
        this.getHotBlogs()
        this.getmusicList()
    },
    mounted() {
        //音乐播放器
        const ap = new APlayer({
            container: document.getElementById('player'),
            fixed: true,
            audio: this.musicList
            /*
            audio: [
                {
                    name: '借',
                    artist: '毛不易',
                    url: 'http://music.163.com/song/media/outer/url?id=569214250.mp3',
                    cover: 'http://p1.music.126.net/vmCcDvD1H04e9gm97xsCqg==/109951163350929740.jpg?param=130y130',
                    theme: '#ebd0c2'
                },
                {
                    name: '像我这样的人',
                    artist: '毛不易',
                    url: 'http://music.163.com/song/media/outer/url?id=569213220.mp3',
                    cover: 'http://p1.music.126.net/vmCcDvD1H04e9gm97xsCqg==/109951163350929740.jpg?param=130y130',
                    theme: '#46718b'
                },
                {
                    name: '无与伦比的美丽',
                    artist: '苏打绿',
                    url: 'http://music.163.com/song/media/outer/url?id=2023994371.mp3',
                    cover: 'http://p2.music.126.net/zfr5TThFTC3vedNp3L24Zg==/109951168565381063.jpg?param=130y130',
                    theme: '#46718b'
                }
            ]*/

        });
        document.addEventListener('scroll', this.parallax, true)
    },
    deactivated() {
        this.showMotto = false
    },
    activated() {
        this.showMotto = true
        window.scrollTo(0, 0)
    },

    methods: {
        // 获取座右铭
        async getMotto() {
            const { data: res } = await this.$axios.get("/myblog/mottos")
            if (res.status === 1) {
                if (res.data.length > 0) {
                    this.firstBGPageInfo.mottos = res.data[0]
                    const n = Math.round(Math.random() * (res.data[0].length - 1));
                    this.firstBGPageInfo.curMotto = this.firstBGPageInfo.mottos[n]
                }
            }
        },
        //获取音乐列表
        async getmusicList() {
            const { data: res } = await this.$axios.get("/admin/musicList")
            if (res.status === 1) {
                this.musicList = res.data.data;
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
            } else {
                this.$refs.box.style.style.backgroundPosition = ''
            }
        }
    }
}
</script>




<style lang="less" scoped>
.animate__animated {
    animation-duration: 3s !important;
}

//座右铭部分
#motto {
    user-select: none;
    font-weight: 560;
    line-height: 1.25;
    color: white;
    text-align: center;
    position: relative;
    top: 42%;

    h1 {
        padding-bottom: 20px;
    }

    p {
        font-size: 22px;
    }
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
    border-right: 5px solid #d2c99a;
    border-bottom: 5px solid #d2c99a;
    display: block;
    border-bottom-right-radius: 5px;
}

@keyframes bounce-in {
    0% {
        transform: translateY(0);
    }

    20% {
        transform: translateY(0);
    }

    50% {
        transform: translateY(-20px);
        // border-right: #ed1414;
        // border-top: #0f1110;
    }

    80% {
        transform: translateY(0);
    }

    100% {
        transform: translateY(0);
    }
}

//上方背景
#bg1 {
    background-image: url("../assets/images/background03.jpg");
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
    background-color: #263529;
    background-image:
        radial-gradient(closest-side, #507863, rgba(80, 120, 99, 0)),
        radial-gradient(closest-side, #5b8463, rgba(132, 157, 133, 0)),
        radial-gradient(closest-side, #4d7a69, rgba(94, 129, 137, 0)),
        radial-gradient(closest-side, #d8d0a7, rgba(210, 201, 154, 0)),
        radial-gradient(closest-side, #98bdc8, rgba(114, 145, 147, 0));
    background-size:
        130vmax 130vmax,
        80vmax 80vmax,
        90vmax 90vmax,
        110vmax 110vmax,
        90vmax 90vmax;
    background-position:
        -80vmax -80vmax,
        60vmax -30vmax,
        10vmax 10vmax,
        -30vmax -10vmax,
        50vmax 50vmax;
    background-repeat: no-repeat;
    animation: 10s movement linear infinite;

    @keyframes movement {

        0%,
        100% {
            background-size:
                130vmax 130vmax,
                80vmax 80vmax,
                90vmax 90vmax,
                110vmax 110vmax,
                90vmax 90vmax;
            background-position:
                -80vmax -80vmax,
                60vmax -30vmax,
                10vmax 10vmax,
                -30vmax -10vmax,
                50vmax 50vmax;
        }

        25% {
            background-size:
                100vmax 100vmax,
                90vmax 90vmax,
                100vmax 100vmax,
                90vmax 90vmax,
                60vmax 60vmax;
            background-position:
                -60vmax -90vmax,
                50vmax -40vmax,
                0vmax -20vmax,
                -40vmax -20vmax,
                40vmax 60vmax;
        }

        50% {
            background-size:
                80vmax 80vmax,
                110vmax 110vmax,
                80vmax 80vmax,
                60vmax 60vmax,
                80vmax 80vmax;
            background-position:
                -50vmax -70vmax,
                40vmax -30vmax,
                10vmax 0vmax,
                20vmax 10vmax,
                30vmax 70vmax;
        }

        75% {
            background-size:
                90vmax 90vmax,
                90vmax 90vmax,
                100vmax 100vmax,
                90vmax 90vmax,
                70vmax 70vmax;
            background-position:
                -50vmax -40vmax,
                50vmax -30vmax,
                20vmax 0vmax,
                -10vmax 10vmax,
                40vmax 60vmax;
        }
    }
}

// 下面中心区域
.blog-area {
    width: 1165px;
    margin: 0 auto;
    padding-top: 36px;
    padding-bottom: 64px;
    overflow: hidden;
}

.blog-right {
    float: left;
}

.blog-left {
    float: left;
    width: 300px;
    margin-right: 15px;
    padding: 0 15px;
    border-radius: 6px;
    transition: all 0.3s;
}
</style>
