
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
            <div class="cloud-bg03"></div>
            <!-- <TagCloud class="tag-detail"></TagCloud>
            <TypeArea class="type-detail"></TypeArea>
             -->

            <div class="blog-area">
                <!--用户信息栏-->
                <UserInfoCard :count="blogCount"></UserInfoCard>

                <div id="player"></div>
                <BlogCard :key="item.id" v-for="(item, index) in blogDetails" :item="item" :imgRight="index % 2 === 0"
                    @click.native="blogDetail(item.id)">
                </BlogCard>

                <Pagination @jumpPage="jumpPage" :pageInfo="{ pageNum: queryInfo.pageNum, pages: pages }"></Pagination>

                <div class="wapper-label">
                    <svg class="star star1"></svg>
                    <svg class="star star2"></svg>
                    <div class="title">Label</div>
                    <div class="content">
                        <ul>
                            <li :key="item.id" v-for="item in tags" @click="itemClicked(item.id)">
                                <span>{{ item.name }}</span>
                            </li>
                        </ul>
                    </div>
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

import 'APlayer/dist/APlayer.min.css';
import APlayer from 'APlayer';

import "animate.css"
import '../assets/css/star.css'
import { set } from "vue";
export default {

    name: "Home",
    components: { BackToTop, BlogCard, Pagination, UserInfoCard },
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
                firstPicture: img3,
                createTime: "2021-03-29",
                views: 1024,
                nickname: "Jack Ma",
                typename: "框架底层原理"
            }],
            //获得歌单
            musicList: [],
            //博客的标签_通过函数获取全部标签
            tags: [
                { id: 1, name: "Java" },
                { id: 2, name: "C++" },
                { id: 3, name: "Golang" },
                { id: 4, name: "Nginx" },
            ],
        }
    },
    created() {
        this.getBlogLists();
        this.getmusicList();

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
        this.getTagList();
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
        
        // 进入博客内容页面
        blogDetail(id) {
            this.$router.push({
                path: "/blogDetail",
                query: {
                    id: id
                }
            });
        },

        //页面相关
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
            var words = ['白云山高', '珠江水长', '吾校矗立', '蔚为国光'];
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
        //获取对应标签下的文章
        itemClicked(id) {
            this.getBlogsByTag(id);
        },
        // 获取所有博客标签
        async getTagList() {
            const { data: res } = await this.$axios.get("/myblog/tagList");
            if (res.status === 1) {
                this.tags = res.data.length > 0 ? res.data[0] : this.tags;
            } else {
                this.$message.warning("获取标签列表失败！")
            }
        },
        async getBlogsByTag(tagId) {
            const { data: res } = await this.$axios.get("/myblog/blogsByTag", { params: { pageNum: this.queryInfo.pageNum, pageSize: this.queryInfo.pageSize, tagId: tagId } });
            if (res.status === 1) {
                this.blogDetails = res.data.length > 0 ? res.data[0] : this.blogDetails
                this.blogCount = res.data.length > 1 ? res.data[1] : this.blogCount
                this.pages = Math.ceil(this.blogCount / this.queryInfo.pageSize)
                if (this.pages <= 0) {
                    this.pages = 1
                }
            } else {
                this.$message.error("获取对应标签下的博客失败,请重试")
            }
        },

    }
}
</script>




<style lang="less" scoped>
.animate__animated {
    animation-duration: 3s;
}

#text {
    font-family: NotoSerifSC-Regular;
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

.cloud-bg01 {
    position: absolute;
    top: 1250px;
    right: 0;
    background-image: url('../assets/images/cloud03.png');
    opacity: 0.48;
    background-repeat: no-repeat;
    height: 500px;
    width: 800px;
}

.cloud-bg02 {
    position: absolute;
    top: 1550px;
    left: 0;
    background-image: url('../assets/images/cloud04.png');
    opacity: 0.68;
    background-repeat: no-repeat;
    height: 520px;
    width: 700px;
}

.cloud-bg03 {
    position: absolute;
    top: 850px;
    left: 300px;
    background-image: url('../assets/images/cloud02.png');
    opacity: 0.58;
    background-repeat: no-repeat;
    height: 520px;
    width: 1000px;
}

.wapper-label {
    width: 290px;
    background-color: #0925f700;
    margin-top: 0px;
    user-select: none;

    position: absolute;
    top: 1350px;
    right: 20px;

    .title {
        color: rgb(255, 255, 255);
        opacity: 0.75;
        text-align: center;
        font-family: DancingScript-Bold;
        font-size: 27px;
        margin-bottom: 10px;
    }

    .content {
        text-align: center;
    }

    ul li {
        display: inline-block;
        padding: 2px;
        color: #ffffff;
        opacity: 0.7;
        list-style: none;
        margin: 6px 10px;
        cursor: pointer;
        font-weight: 500;
        transition: all 0.3s;
        font-family: NotoSerifSC-Regular;

    }

    ul li:hover {
        transform: scale(1.2);
    }

    .star {
        position: relative;
        background-color: #ffffff;
        clip-path: path('M29.51068,15.41064C19.78412,13.544,18.456,12.21582,16.58929,2.48926a.60015.60015,0,0,0-1.17871,0C13.54437,12.21582,12.21625,13.544,2.4892,15.41064a.60016.60016,0,0,0,0,1.17872c9.72705,1.86669,11.05517,3.19482,12.92138,12.92138a.60027.60027,0,0,0,1.17871,0c1.8667-9.72656,3.19483-11.05469,12.92139-12.92138a.60016.60016,0,0,0,0-1.17872Z');
        width: 40px;
        height: 40px;
        animation: twinkle 2s ease-out infinite alternate;
    }

    .star1 {
        position: absolute;
        top: 220px;
        left: 50px;
        animation-direction: alternate-reverse;
    }

    .star2 {
        position: absolute;
        top: -10px;
        left: 210px;
    }

    @keyframes twinkle {
        0% {
            opacity: 0.1;
        }

        100% {
            opacity: 0.8;
        }
    }

}
</style>
