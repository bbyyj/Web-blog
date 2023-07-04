<template>
    <div class="bg" >
        <div class="main">
            <div class="maintitle" align="center">资源库</div>
            <!-- 分类选择、搜索、上传 -->
            <div class="categories-area">
            <el-row :gutter="15" type="flex" justify="center" align="middle">
                <el-col :span="3.5">
                    <el-dropdown>
                    <el-button class="button1" type="primary" icon="el-icon-notebook-1">
                        更多分类<i class="el-icon-arrow-down el-icon--right"></i>
                    </el-button>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item @click.native="getLinks()">全部</el-dropdown-item>
                        <el-dropdown-item v-for="subject in categories" :key="subject.id"
                            @click.native="showdata(subject.id, '1')">
                            {{ subject.name }}
                        </el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
                </el-col>
                <el-col :span="8">
                    <el-input v-model="info.sname" placeholder="请输入关键词" clearable></el-input>
                </el-col>
                <el-col :span="2.5">
                    <el-button class="button1" type="primary" @click="searchdata('1')" icon="el-icon-search">搜索</el-button>
                </el-col>
                <el-col :span="4">
                    <el-button class="button1" type="primary" @click="adddata" icon="el-icon-paperclip">上传资料</el-button>
                </el-col>
            </el-row>                
            </div>
            <transition appear
                        name="animate__animated animate__bounce animate__slow"
                        enter-active-class="animate__fadeInUp">
                <div class="center-area" style="text-align: center;">

                    <div class="category grow" style="text-align: left;"  :key="item.ID" v-for="item in items" >
                        <!-- <a @click="downloadFile(item.name, item.url)"> -->
                        <a @click="downloadFile(item.name)">

                            <strong class="title">
                            {{ item.name }}
                        </strong>
                        <p class = "desc">
                            {{ item.desc }}
                        </p>
                        <p class = "desc">
                            下载总数：{{ item.downloadnum }}
                            大小：{{ (item.filesize/1048576).toFixed(2) }}MB
                        </p>
                        <p class = "desc">
                            更新时间：{{ dateFormat(item.UpdatedAt) }}
                        </p>
                        </a>
                    </div>
                </div>
            </transition>
            <!-- 分页 -->
            <el-pagination popper-class="select-down"  @size-change="handleSizeChange" @current-change="handleCurrentChange"
                    :current-page="queryInfo.pagenum" :page-sizes="[5, 10]" :page-size="queryInfo.pagesize"
                    layout="total, sizes, prev, pager, next, jumper" :total="total">
                </el-pagination>
            <!-- 编辑、添加显示页面 -->
            <el-dialog title="资源信息" :visible.sync="dialogFormVisible">
                <el-form label-width="80px">
                    <el-form-item label="标题：">
                        <el-input v-model="postInfo.name" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="类别：">
                        <el-select :value="selectedCategory" @visible-change="getAllCategories" @change="changeCategory" clearable placeholder="资源类别">
                            <el-option v-for="item in categories" :key="item.id" :value="item.name">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="描述：">
                        <el-input type="textarea" :rows="3" v-model="postInfo.desc" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="文件：">
                        <el-row :gutter="22">
                            
                            <form @submit.prevent="uploadFile">
                            <input type="file" ref="fileInput" name="f1">
                            <button class="upload" type="submit">上传</button>
                            </form>

                        </el-row>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button class="edt" @click="cancel">取 消</el-button>
                    <el-button class="yes" type="primary" @click="commitLink">确 定</el-button>
                </div>
            </el-dialog>
        </div>
    </div>

</template>

<script>

import "animate.css"
import dayjs from "dayjs";
import axios from "axios";



import Pagination from "../components/Pagination";

export default {
    name: "ResourceLib",
    components: { Pagination },
    data() {
        return {
            //接收到的全部数据
            items:[],
           //当前页面
           queryInfo: {
                pagenum: 1,
                pagesize: 5,
            },
            //页面总数
            pages: 1,
            //资源总条数
            total: 0,
            //当前类别
            categoryid: 1,
            //全部类别及其id
            categories:[],
            //搜索用
            info:{sname: "",},
            //添加资源显示页面
            dialogFormVisible: false,
            //添加资源用
            postInfo: {
                ID: 0,
                name: "",
                desc: "",
                categoryid: 0,
                url: "",
            },
            selectedCategory: "",

        }
    },
    methods: { 
        //进入页面默认展示内容   
        async getLinks() {
            let pagenum = this.queryInfo.pagenum;
            let pagesize = this.queryInfo.pagesize;
            const {data: res} = await this.$axios.get("/myblog/t/pageresource", { params: { pagenum: pagenum, pagesize: pagesize } });

            if(res.status = 563){
                this.items = res.data[0];
                this.total = res.data[1];
                this.categoryid = 0;
                console.log(res);
            }
            else{
                this.$message.warning("获取资源失败")
                return
            }

            //分页相关
            this.pages = Math.ceil(this.total / this.queryInfo.pagesize);
            if (this.pages <= 0) {
                this.pages = 1
            }
        },

        //获取全部分类
        async getAllCategories() {
            const {data:res} = await this.$axios.get("/admin/categories");
            if (res.status !== 1) {
                this.$message.error("获取分类列表失败，请重试！")
                return
            }
            this.categories = res.data.length > 0 ? res.data[0] : []
            console.log(res);
        },

        //点击分类展示对应内容
        async showdata(cateid, flag){
            if(flag == 1){
                this.queryInfo.pagenum = 1;
            }
            
            let pagenum = this.queryInfo.pagenum;
            let pagesize = this.queryInfo.pagesize;
            const {data: res} = await this.$axios.get("/myblog/t/pageresourcebycategoryid", { params: { categoryid:cateid, pagenum: pagenum, pagesize: pagesize } });
            
            if(res.status = 563){
                this.items = res.data[0];
                this.total = res.data[1];
                this.categoryid = cateid;
                console.log(res);
            }
            else{
                this.$message.warning("获取资源失败")
                return
            }

            //分页相关
            this.pages = Math.ceil(this.total / this.queryInfo.pagesize);
            if (this.pages <= 0) {
                this.pages = 1
            }
        },

        //查询
        async searchdata(flag) {
            if(flag == 1){
                this.queryInfo.pagenum = 1;
            }
            let pagenum = this.queryInfo.pagenum;
            let pagesize = this.queryInfo.pagesize;
            const {data: res} = await this.$axios.get("/myblog/t/queryresource", { params: {  name: this.info.sname, pagenum: pagenum, pagesize: pagesize  } });
            console.log(res);

            if(res.status = 563){
                this.items = res.data[0];
                this.total = res.data[1];
                this.categoryid = -1;
                console.log(res);
            }
            else{
                this.$message.warning("获取资源失败")
                return
            }

            //分页相关
            this.pages = Math.ceil(this.total / this.queryInfo.pagesize);
            if (this.pages <= 0) {
                this.pages = 1
            }
        },

        //上传资料
        changeCategory(name) {
            this.selectedCategory = name
            const val = this.categories.find(item => {
                return item.name === name
            })
            this.postInfo.categoryid = val.id
        },

        adddata() {
            this.postInfo.ID = 0
            this.postInfo.url = ""
            this.dialogFormVisible = true
        },

        cancel() {
            this.postInfo = {
                ID: 0,
                name: "",
                desc: "",
                categoryid: 0,
                url: "",
            }
            this.dialogFormVisible = false
            this.selectedCategory = ""
        },

        //提交文件（仅文件）
        uploadFile() {
            const f1 = this.$refs.fileInput.files[0];
            const formData = new FormData();
            formData.append('f1', f1);


            if (!f1.name.endsWith(".zip")) {
                this.$message.error("只能上传 ZIP 文件!");
                return;
            }

            axios.post('/myblog/t/uploadresourcecheck', formData)
                .then(response => {
                // 处理后端返回的数据
                this.postInfo.url = response.data.url
                this.$message.success("上传成功!")
                })
                .catch(error => {
                // 处理错误
                console.log(error);
                });
        },
    

        async commitLink() {
            //若文件名格式错误报错
            if(!this.postInfo.name.endsWith(".zip")) {
                this.$message.error("请在文件名后添加.zip后缀！")
                return
            }
            //若未上传文件报错
            if(this.postInfo.url == "") {
                this.$message.error("请先上传文件！")
                return
            }
            
            
            let res
            res = await this.$axios.post("/myblog/t/addresourcecheck",  { name: this.postInfo.name, desc: this.postInfo.desc, categoryid: this.postInfo.categoryid, url: this.postInfo.url })
            console.log(res);
            
            if (res.data.status !== 101) {
                this.$message.error("操作失败，请重试！")
            } else {
                this.cancel()
                this.$message.success("操作成功，请等待审核！")
                this.postInfo = {
                ID: 0,
                name: "",
                desc: "",
                categoryid: 0,
                url: "",
            }
            }
        },

        uploadSuccess(response) {
            this.postInfo.icon = response;
        },

        //下载
        downloadFile(fileName) {
            

        const url = `/myblog/t/downloadresource?name=${fileName}`;
        const newWindow = window.open('', '_blank');
        newWindow.location.href = axios.defaults.baseURL + url;
        setTimeout(() => {
            newWindow.close();
        }, 500);

           
        },

        //时间格式
        dateFormat(d) {
            return dayjs(d).format("YYYY-MM-DD HH:mm:ss")
        },

        //页面相关
        handleSizeChange(newSize) {
            this.queryInfo.pagesize = newSize;
            if(this.categoryid == 0) {
                this.getLinks()
            } else if(this.categoryid == -1) {
                this.searchdata('0')
            } else {
                this.showdata(this.categoryid,'0');
            }
            
        },

        handleCurrentChange(newPage) {
            this.queryInfo.pagenum = newPage;
            if(this.categoryid == 0) {
                this.getLinks()
            } else if(this.categoryid == -1) {
                this.searchdata('0')
            } else {
                this.showdata(this.categoryid,'0');
            }
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

.button1 {
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

.button1:before {
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

.button1:hover {
    background-color: #9b93b7;
    color: #fff;
    // border-bottom: 4px solid darken(#fff, 10%);
    border-color: #9b93b7;
  }
  
.button1:focus {
    background-color: #9b93b7;
    color: #fff;
    // border-bottom: 4px solid darken(#fff, 10%);
    border-color: #9b93b7;
}


.maintitle{
    font-size: 56px;
    color: #181720;
    margin-bottom: 50px;
    bottom: 0 !important;
    right: 0 !important;
    font-family: AaMeiSong-2;
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
    // border-radius: 5px;
    padding-bottom: 80px;
}

.categories-area {
    margin-left: 5%;
    overflow: hidden;
}


.category {
    display: inline-block;
    // border-radius: 7px;
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
    font-size: 14px;
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


/deep/ .el-pagination__jump{
    color: #eff1f6;
    font-size: 16px;
}
.el-pagination{
    padding-top: 30px;
    padding-left: 500px;
}

.upload {
    display: inline-block;
    line-height: 1;
    white-space: nowrap;
    border: 1px solid #DCDFE6;
    text-align: center;
    box-sizing: border-box;
    outline: 0;
    margin: 0;
    transition: .1s;
    font-weight: 500;
    padding: 12px 20px;
    font-size: 14px;
    border-radius: 4px;
    background-color: #baaaca19;
    color: #917ba7ee;
    border-color: #baaacaee;

}

.edt:hover, .upload:hover {
    background-color: #baaacaee;
    color: #fff;
    border-color: #baaacaee;
}

.yes {
    background-color: #947dabee;
    color: #fff;
    border-color: #87739aee;
}

.yes:hover {
    background-color: #6e5884ee;
    color: #fff;
    border-color: #6e5884ee;
}

::file-selector-button {
    height: 3rem;
    border-radius: .25rem;
    border: 1px solid #baaacaee;
    background-color: #baaaca19;
    color: #917ba7ee;
    cursor: pointer;
  }

  ::file-selector-button:hover {
    background-color: #baaacaee;
    color: #fff;
    border-color: #baaacaee;
  }


</style>
