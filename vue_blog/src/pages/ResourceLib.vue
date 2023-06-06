<template>
    <div class="bg" >
        <div class="main">
            <div class="maintitle" align="center">资源库</div>
            <!-- <div style="text-align: center;">
                    <button @click="getLinks">全部</button>
                    <button @click="showdata('2')">计算机组成原理</button>
                    <button @click="showdata('1')">计算机网络</button>
                    <button @click="showdata('3')">数据结构与算法</button>
                    <button @click="showdata('4')">操作系统</button>
            </div> -->
            <ul class="categories-area">
                <el-dropdown>
                    <el-button type="primary">
                        更多科目<i class="el-icon-arrow-down el-icon--right"></i>
                    </el-button>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item v-for="subject in categories" :key="subject.id"
                            @click.native="showdata(subject.id)">
                            {{ subject.name }}
                        </el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
                
            </ul>


            <transition appear
                        name="animate__animated animate__bounce animate__slow"
                        enter-active-class="animate__fadeInUp">
                <div class="center-area" style="text-align: center;">
                    <div class="category grow" style="text-align: left;"  :key="item.ID" v-for="item in items" >
                        <a @click="downloadFile(item.name, item.url)">
                            <strong class="title">
                            {{ item.name }}
                        </strong>
                        
                        
                        <p class = "desc">
                            {{ item.desc }}
                        </p>
                        <p class = "desc">
                            下载总数：{{ item.downloadnum }}
                            大小：{{ item.filesize }}
                        </p>
                        <p class = "desc">
                            更新时间：{{ dateFormat(item.UpdatedAt) }}
                        </p>
                        </a>
                    </div>
                </div>
            </transition>
            <!-- 分页 -->
            <!-- <Pagination class="pagebar" @jumpPage="jumpPage" :pageInfo="{ pageNum: pagenum, pages: pages }">
                </Pagination> -->

        </div>
    </div>

</template>

<script>

import "animate.css"
import dayjs from "dayjs";

import Pagination from "../components/Pagination";

export default {
    name: "ResourceLib",
    components: { Pagination },
    data() {
        return {
            //接收到的全部数据
            items:[],
           //当前页面
            pagenum: 1,
            //页面大小
            pagesize: 100,
            //页面数量
            pages: 1,
            //当前类别
            categoryid: 1,
            //全部类别及其id
            categories:[],

        }
    },
    methods: { 
        //进入页面默认展示内容   
        async getLinks() {
            let pagenum = this.pagenum;
            let pagesize = this.pagesize;
            const {data: res} = await this.$axios.get("/myblog/t/pageresource", { params: { pagenum: pagenum, pagesize: pagesize } });

            if(res.status = 563){
                this.items = res.data[0];
                console.log(res);
            }
            else{
                this.$message.warning("获取资源失败")
                return
            }
        },

        //获取全部分类
        async getAllCategories() {
            const {data:res} = await this.$axios.get("/admin/categories");
            if (res.status !== 1) {
                this.$message.error("获取列表失败，请重试！")
                return
            }
            this.categories = res.data.length > 0 ? res.data[0] : []
            console.log(res);
        },

        //点击分类展示对应内容
        async showdata(cateid){
            let pagenum = this.pagenum;
            let pagesize = this.pagesize;
            const {data: res1} = await this.$axios.get("/myblog/t/pageresourcebycategoryid", { params: { categoryid:cateid, pagenum: pagenum, pagesize: pagesize } });
            
            if(res1.status = 563){
                this.items = res1.data[0];
                this.categoryid = cateid;
                console.log(res1);
            }
            else{
                this.$message.warning("获取资源失败")
                return
            }

        },

        async downloadFile(fileName, data) {
            if (!data) {
                return;
            }
            const {data: res} = await this.$axios.get("/myblog/t/downloadresource", { params: {  name: fileName } });
            console.log(res);

            let url = window.URL.createObjectURL(new Blob([data]));
            let link = document.createElement('a');
            link.style.display = 'none';
            link.href = url;
            link.setAttribute('download', fileName);
            document.body.appendChild(link);
            link.click();
        },

        dateFormat(d) {
            return dayjs(d).format("YYYY-MM-DD HH:mm:ss")
        },


    },
    created() {
        window.scrollTo(0, 0)
        this.getLinks()
        this.getAllCategories()
    }
}
</script>

<style lang="less" scoped>

button {
    background-color: #a69ec699;
    margin: 0.5%;
    padding: 0.2em 1em;
    margin-top: 2%;
    text-align: center;
    text-decoration: none;
    color: #3d395299;
    border: 2px solid #ffffff00;
    font-size: 20px;
    display: inline-block;
    border-radius: 0.2em;
    transition: all 0.2s ease-in-out;
    position: relative;
    overflow: hidden;
}

button:before {
    content: "";
    background-color: rgba(255,255,255,0.5);
    height: 100%;
    width: 3em;
    display: block;
    position: absolute;
    top: 0;
    left: -4.5em;
    transform: skewX(-45deg) translateX(0);
    transition: none;
  }

button:hover {
    background-color: #7378acb7;
    color: #fff;
    border-bottom: 4px solid darken(#fff, 10%);
  } 


.maintitle{
    font-size: 48px;
    color: #181720;
    margin-bottom: 50px;
    bottom: 0 !important;
    right: 0 !important;
    font-family: "PingFang SC", "Microsoft YaHei", Lato, sans-serif;
    opacity: 0.5;
    padding-top: 6%;
}


.bg {                      
    background: url(../assets/images/SeaClouds.svg) no-repeat;
    height:100%;
    width:100%;
    overflow: hidden;
    background-size:cover;
}

.main {
    padding-top: 80px;
    padding-bottom: 150px;
    min-height: 800px;
}

.center-area {
    width: 66%;
    background-color: rgba(255, 255, 255, .6);
    margin: 0 auto;
    margin-top: 2%;
    border-radius: 5px;
    padding-bottom: 80px;
}

.categories-area {
    margin: -10px 200px 50px 200px;
    overflow: hidden;
}


.category {
    display: inline-block;
    border-radius: 7px;
    background-color: #ffffff87;
    width: 280px;
    height: 150px;
    padding: 20px;
    margin-top: 2%;
    margin-right: 2%;
    transition: all 0.3s linear;
}

.desc{
    font-weight: 500;
    color: #979898;
    margin: 0;
    white-space: nowrap;
    overflow: hidden;
    max-width: 250px;
    text-overflow: ellipsis;
}

.pagebar {
    padding-bottom: 50px;
}

.grow {
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
.grow:hover, .grow:focus, .grow:active {
  -webkit-transform: scale(1.1);
  transform: scale(1.03);
}

.title {
    font-weight: 700;
    color: #7378ac;
}


</style>
