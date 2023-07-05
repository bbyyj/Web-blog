<!--顶部导航栏组件-->
<template>
    <div class="navi-bar">
        <!-- 小站名字-->
        <p>{{ siteName }}</p>
        <!--中间搜索、导航区域-->

        <div class="navi-bar-center">
            <div class="navi-bar-item-group">
                <!--搜索框-->
                <div class="search-warp">
                    <el-input class="search-input" size="mini" placeholder="请输入内容" prefix-icon="el-icon-search"
                        v-model="keyWord" @keyup.enter.native="search">
                    </el-input>
                    <div @mouseleave="showResult = false; keyWord = ''" v-show="showResult" class="search-result">
                        <div class="search-content">
                            <ul>
                                <!-- <li>{{total ? '共' + total + '条记录！' : '没有找到记录！' }}</li> -->
                                <li :key="item.id" v-for="item in searchResult">
                                    <a @click="blogDetail(item.id)">{{ item.title }}</a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
                <!--导航按键-->
                <ul>
                    <li>
                        <NaviItem :activeRoute="activeRoute" index-name="首页" ico-class="icon-home1" rout="/home">
                        </NaviItem>
                    </li>
                    <li>
                        <NaviItem :activeRoute="activeRoute" index-name="软工题库" ico-class="icon-tikuzidian" rout="/ChooseQS">
                        </NaviItem>
                    </li>
                    <li>
                        <NaviItem :activeRoute="activeRoute" index-name="提问箱" ico-class="icon-icon_icon_tiwen_s" rout="/AskBoard">
                        </NaviItem>
                    </li>
                    <li>
                        <NaviItem :activeRoute="activeRoute" index-name="资源库" ico-class="icon-ziyuanku1" rout="/resourceLib">
                        </NaviItem>
                    </li>
                    <li>
                        <NaviItem :activeRoute="activeRoute" index-name="选课通" ico-class="icon-xuanke-2" rout="/HelpChoose">
                        </NaviItem>
                    </li>
                    <li>
                        <NaviItem :activeRoute="activeRoute" index-name="后台管理" ico-class="icon-houtaiguanli"
                            rout="/login"></NaviItem>
                    </li>
                </ul>
                <i @click="fullscreen" class="iconfont icon-quanpingmu fullscreen"></i>
            </div>
        </div>

    </div>
</template>

<script>

import "../assets/icon/iconfont.css"

import NaviItem from "./NaviItem";

export default {
    name: "NavigationBar",
    components: { NaviItem },
    data() {
        return {
            keyWord: "",
            siteName: "中山大学",
            activeRoute: "/home",
            searchResult: [],
            // total: 0,
            showResult: false
        }
    },
    methods: {
        fullscreen() {
            const de = document.documentElement;
            if (de.requestFullscreen) {
                de.requestFullscreen();
            }
            window.scrollTo(0, 1);
        },
        async search() {
            if (!this.keyWord) {
                return
            }
            const { data: res } = await this.$axios.get("/myblog/search", { params: { keyWord: this.keyWord } })
            if (res.status !== 1) {
                this.$message.error("搜索失败，请重试！")
                return
            }
            if (res.data.length > 0) {
                this.searchResult = res.data[0]
                if (res.data.length > 1) {
                    this.total = res.data.length
                }
            }
            this.showResult = true
        },
        blogDetail(id) {
            this.showResult = false
            this.$router.push({
                path: "/blogDetail",
                query: {
                    id: id
                }
            });
        }
    },
    created() {
        this.activeRoute = this.$router.currentRoute.path
    },
    watch: {
        // 监听路由改变，路由改变时给activeRoute赋值，从而让导航栏组件高亮
        $route(to) {
            this.activeRoute = to.path
        }
    }

}
</script>


<style lang="less">
.search-input {
    width: 200px !important;
    border: None;

    input {
        width: 200px;
        border: 1px solid #3d3952 !important;
        border-radius: 20px;
        background: rgba(255, 255, 255, 0.3);
    }

}
</style>

<style lang="less" scoped >
ul {
    margin: 0;
    display: inline-block;
}

li {
    display: inline-block;
    list-style: none;
    margin: 0 15px;
}

.search-warp {
    position: relative;
    display: inline-block;
}

.search-result {
    width: 240px;
    position: absolute;
    background-color: #fff;
    top: 45px;
    left: -20px;
    border-radius: 8px;

    .search-content {
        overflow-y: auto;
        max-height: 300px;
        margin: 8px 0;
    }
}

.search-result::before {
    content: '';
    position: absolute;
    display: inline-block;
    width: 16px;
    height: 16px;
    border-left: 1px solid #fff;
    border-top: 1px solid #fff;
    transform: rotate(45deg);
    top: -8px;
    left: 50px;
    background-color: #fff;
    border-top-left-radius: 3px;
}

.search-result ul {
    padding-left: 16px;
    font-size: 14px;

    li {
        display: block;
        cursor: pointer;
        margin: 2px 0;

        a {
            outline: none;
            text-decoration: none;
            color: #3d3952;
        }
    }

    li:first-child {
        cursor: default;
    }

}

p {
    float: left;
    text-align: center;
    font-size: 25px;
    align-items: center;
    margin: 10px 0 0 50px;
    height: 40px;
    -webkit-background-clip: text;
    font-weight: bolder;
    color: #7378ac;
    user-select: none;
}

.navi-bar {
    z-index: 999;
    width: 100%;
    height: 60px;
    position: fixed;
    background: rgba(252, 252, 252, 0.35);
    font-family: NotoSerifSC-Regular;
    font-weight: normal;
    font-size: 18px;
}

.navi-bar-center {
    margin: 10px 0;
    height: 40px;
    width: 100%;
}

.navi-bar-item-group {
    display: inline-block;
    align-items: center;
    height: 100%;
    float: right;
    margin-right: 70px;
}

.fullscreen {
    color: #7378ac;
    font-size: 27px;
    cursor: pointer;
    margin: 7px 8px 0px 10px;
}
// .iconfont .icon-icon_icon_tiwen_s{
//     font-size: 30px !important;
// }

</style>
