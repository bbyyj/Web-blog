
<template>
    <el-container class="home_container">
        <!-- 头部区域 -->
        <el-header>
            <div>
                <span>博客管理后台</span>
            </div>
            <el-button type="info" @click="logout">退出</el-button>
        </el-header>

        <el-container>
            <!-- 侧边区域 -->
            <el-aside :width="isCollapse ? '65px' : '200px'">
                <div class="toggle-button" @click="toggleCollapse">|||</div>
                <el-menu text-color="#eaf0ef" active-text-color="#eaf0ef" unique-opened
                    :collapse="isCollapse" :collapse-transition="false" router :default-active="activePath">

                    <el-submenu :index="menu.id" v-for="(menu, i1) in menus" :key="menu.id">
                        <!-- 一级菜单的模板区域 -->
                        <template slot="title">
                            <!-- 图标 -->
                            <i :class="menu.icon"></i>
                            <!-- 文本 -->
                            <span>{{ menu.name }}</span>
                        </template>

                        <!-- 二级菜单 -->
                        <el-menu-item :key="subMenu.id" v-for="(subMenu, i2) in menu.submenus" :index="subMenu.index"
                            @click="saveNavState(i1, i2)">
                            <template slot="title">
                                <!-- 图标 -->
                                <i :class="subMenu.icon"></i>
                                <!-- 文本 -->
                                <span>{{ subMenu.name }}</span>
                            </template>
                        </el-menu-item>
                    </el-submenu>
                </el-menu>
            </el-aside>

            <!-- 主体区域 -->
            <el-main>
                <el-breadcrumb separator-class="el-icon-arrow-right">
                    <el-breadcrumb-item>{{ breadcrumbItem1 }}</el-breadcrumb-item>
                    <el-breadcrumb-item>{{ breadcrumbItem2 }}</el-breadcrumb-item>
                </el-breadcrumb>
                <!-- 路由占位符 -->
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>



<script>
export default {
    data() {
        return {
            isCollapse: false,
            activeMenu: -1,
            activeSubMenu: -1,
            menus: [
                {
                    id: "1",
                    icon: "el-icon-edit",
                    name: "博客管理",
                    submenus: [
                        { id: "11", index: "/listBlogs", icon: "el-icon-menu", name: "查看博客" },
                        { id: "12", index: "/addBlog", icon: "el-icon-plus", name: "新增博客" },
                        { id: "13", index: "/listMottos", icon: "el-icon-menu", name: "查看格言" }
                    ]
                },
                {
                    id: "2",
                    icon: "el-icon-notebook-2",
                    name: "分类管理",
                    submenus: [
                        { id: "21", index: "/listTypes", icon: "el-icon-menu", name: "查看分类" },
                    ]
                },
                {
                    id: "3",
                    icon: "el-icon-paperclip",
                    name: "标签管理",
                    submenus: [
                        { id: "31", index: "/listTags", icon: "el-icon-menu", name: "查看标签" },
                    ]
                },
                // {
                //     id :"4",
                //     icon: "el-icon-picture-outline",
                //     name: "图片管理",
                //     submenus: [
                //         {id:"41", index: "/listImages", icon:"el-icon-menu", name:"查看图片"},
                //     ]
                // },
                {
                    id: "5",
                    icon: "el-icon-notebook-1",
                    name: "音乐管理",
                    submenus: [
                        { id: "51", index: "/manageMusic", icon: "el-icon-menu", name: "管理音乐" },
                    ]
                },
                {
                    id: "6",
                    icon: "el-icon-chat-line-square",
                    name: "留言管理",
                    submenus: [
                        { id: "61", index: "/manageMsg", icon: "el-icon-menu", name: "管理留言" },
                    ]
                },
                {
                    id: "7",
                    icon: "el-icon-star-off",
                    name: "资源库管理",
                    submenus: [
                        { id: "71", index: "/manageLink", icon: "el-icon-menu", name: "管理文件" },
                    ]
                },
                {
                    id: "8",
                    icon: "el-icon-star-off",
                    name: "题库管理",
                    submenus: [
                        { id: "81", index: "/manageTest", icon: "el-icon-menu", name: "管理题库" },
                        { id: "82", index: "/manageSubjectsAndChapters", icon: "el-icon-menu", name: "管理科目和章节" },
                    ]
                },
                {
                    id: "9",
                    icon: "el-icon-star-off",
                    name: "提问箱管理",
                    submenus: [
                        { id: "91", index: "/ManagehaveAsked", icon: "el-icon-menu", name: "已回答问题" },
                        { id: "92", index: "/ManagenoAsk", icon: "el-icon-menu", name: "未回答问题" },
                    ]
                }
            ]
        }
    },
    methods: {
        logout() {

            // 清空token
            window.sessionStorage.clear();
            // 跳转登录页
            this.$router.push("/login");
        },
        toggleCollapse() {   //单机按钮实现菜单的折叠与展开
            this.isCollapse = !this.isCollapse;
        },
        saveNavState(i1, i2) {     //保存链接的激活状态
            this.activeMenu = i1
            this.activeSubMenu = i2
        }
    },
    computed: {
        activePath() {
            return this.activeMenu === -1 || this.activeSubMenu === -1 ?
                null : this.menus[this.activeMenu].submenus[this.activeSubMenu].index
        },
        breadcrumbItem1() {
            return this.activeMenu === -1 ? null : this.menus[this.activeMenu].name
        },
        breadcrumbItem2() {
            return this.activeMenu === -1 || this.activeSubMenu === -1 ?
                null : this.menus[this.activeMenu].submenus[this.activeSubMenu].name
        }
    }
}
</script>



<style lang="less" scoped>
.home_container {
    height: 100%;
}
.el-header {
    background-color: #1c436d;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: #F9F5EB;
    padding-top: 10px;
    padding-bottom: 10px;
    >div {
        display: flex;
        align-items: center;

        span {
            margin-left: 100px;
            font-size: 30px;
            font-weight: bold;
            
        }
    }
}
.el-button--info{
    background-color: #EA5455;
    border: none;
    color: #F9F5EB;
    font-size: 16px;
}
.el-aside {
    background-color: #1c436d;

    ele-menu {
        border-right: 1px solid #1c436d;
    }
}

.el-main {
    background-color: #F9F5EB;
}

.el-menu {
    border-right: 1px solid #1c436d;
    background-color: #1c436d;
    // color: #1d3d39;
    opacity: 1;
    font-weight: 500;
    font-size: 50px;
}

/deep/ .el-submenu .el-submenu__title:hover, .el-menu-item:hover, .el-submenu .el-submenu__title:focus, .el-menu-item:focus{
    outline: 0 !important;
    background-color: #7189a4 !important;
}
.el-menu-item {
    background-color: #1c436d !important;
    border: none;
}
.toggle-button {
    background-color: #1c436d;
    font-size: 10px;
    line-height: 30px;
    color: #F9F5EB;
    text-align: center;
    letter-spacing: 0.2em;
    cursor: pointer;
}
</style>
