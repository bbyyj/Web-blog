<template>
    <div>
        <el-card>
            <el-row :gutter="15">
                <el-col :span="9">
                    <el-input v-model="info.sname" placeholder="请输入关键词" clearable></el-input>
                </el-col>
                <el-col :span="3">
                    <el-button class="add" type="primary" @click="searchdata" icon="el-icon-search">搜索</el-button>
                </el-col>
                <el-button class="add" type="primary" plain @click="handleAdd">添加资源</el-button>
                <el-button class="show-categories" type="primary" plain @click="showCategories">查看资源分类</el-button>
            </el-row>
            <!-- 列表区域 -->
            <el-table :data="links" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="文件名" prop="name" width="220px"></el-table-column>
                <el-table-column label="描述" prop="desc"></el-table-column>
                <el-table-column label="类别" prop="category" width="150px"></el-table-column>
                <el-table-column label="更新时间">
                    <template slot-scope="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{ dateFormat(scope.row.UpdatedAt) }}</span>
                </template>
                </el-table-column>
                <el-table-column label="操作"  width="150">

                    <template scope="scope">
                        <el-button class="edt" size="mini" @click="handleEdit(scope.$index)">编辑</el-button>
                        <el-button class="del" size="mini" type="danger" @click="handleDelete(scope.row.ID)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                           :current-page="queryInfo.pagenum" :page-sizes="[10, 12, 15, 20]" :page-size="queryInfo.pagesize"
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
                            <input   type="file" ref="fileInput" name="f1">
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
            <!-- 分类信息展示 -->
            <el-dialog title="资源分类信息" :visible.sync="dialogCategoryVisible">
                <el-row>
                    <el-button class="add" type="primary" plain @click="innerVisible=true;categoryPost.id=0">添加资源分类</el-button>
                </el-row>
                <el-table :data="categories" border stripe>
                    <el-table-column type="index"></el-table-column>
                    <el-table-column label="名称" prop="name"></el-table-column>
                    <el-table-column label="操作"  width="150">
                        <template slot-scope="scope">
                            <el-button size="mini" @click="editCategory(scope.$index)">编辑</el-button>
                            <el-button size="mini" type="danger" @click="deleteCategory(scope.row.id)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
                <el-dialog
                    width="40%"
                    title="添加资源分类"
                    :visible.sync="innerVisible"
                    append-to-body>
                    <el-input placeholder="分类名称" v-model="categoryPost.name"></el-input>
                    <div slot="footer" class="dialog-footer">
                        <el-button class="edt" @click="addCategoryCancel">取 消</el-button>
                        <el-button class="yes" type="primary" @click="commitCategory">确 定</el-button>
                    </div>
                </el-dialog>
            </el-dialog>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogCategoryVisible=false">取 消</el-button>
            </div>
        </el-card>
    </div>
</template>

<script>
import axios from "axios";
import dayjs from "dayjs";

export default {
    name: "manageResLink",
    data() {
        return {
            //接收到的该页数据
            links: [],
            //编辑、添加资源信息显示对话
            dialogFormVisible: false,
            //展示分类列表对话
            dialogCategoryVisible: false,
            innerVisible: false,
            //资源总条数
            total: 0,
            selectedCategory: "",
            postInfo: {
                ID: 0,
                name: "",
                desc: "",
                categoryid: 0,
                url: "",
            },
            //添加、更新资源分类递交用
            categoryPost: {
                id: 0,
                name: ""
            },
            //页码相关
            queryInfo: {
                pagenum: 1,
                pagesize: 10
            },
            //所有分类 包括id和名字
            categories: [],
            //搜索用
            info:{sname: "",},
            //
        }
    },
    methods: {
        async getLinkList() {
            const {data:res} = await this.$axios.get("/admin/t/pageresource", {params: this.queryInfo});
            if(res.status !== 563) {
                this.$message.error("获取列表失败，请重试！")
                return
            }

            let links, categories, count
            if (res.data.length > 2 ) {
                links = res.data[0]
                categories = res.data[1]
                //总数
                count = res.data[2]
            } else {
                this.$message.error("获取列表失败，请重试！")
                return
            }

            this.total = count
            this.categories = categories
            const m = new Map()
            this.links.splice(0, this.links.length)
            categories.forEach((val => {
                m.set(val.id, val.name)
            }))
            links.forEach((val) => {
                this.links.push({...val, category: m.get(val.categoryid)})
            })
        },
        async getAllCategories() {
            const {data:res} = await this.$axios.get("/admin/categories");
            if (res.status !== 1) {
                this.$message.error("获取列表失败，请重试！")
                return
            }
            this.categories = res.data.length > 0 ? res.data[0] : []
        },
        changeCategory(name) {
            this.selectedCategory = name
            const val = this.categories.find(item => {
                return item.name === name
            })
            this.postInfo.categoryid = val.id
        },

         //查询
         async searchdata() {
            let pagenum = this.queryInfo.pagenum;
            let pagesize = this.queryInfo.pagesize;
            const {data: res} = await this.$axios.get("/myblog/t/queryresource", { params: {  name: this.info.sname, pagenum: pagenum, pagesize: pagesize  } });

            if(res.status = 563){
                this.links = res.data[0];
                this.total = res.data[1];
            }
            else{
                this.$message.warning("获取资源失败")
                return
            }
        },

        //时间格式
        dateFormat(d) {
            return dayjs(d).format("YYYY-MM-DD HH:mm:ss")
        },

        //添加资源button对应
        async handleAdd() {
            this.postInfo.ID = 0
            this.postInfo.url = ""
            this.dialogFormVisible = true
        },

        //表格内编辑button对应
        async handleEdit(index) {
            this.dialogFormVisible = true
            this.postInfo = {...this.links[index]}
            const val = this.categories.find(item => {
                return item.id === this.links[index].categoryid
            })
            this.selectedCategory = val.name
            this.postInfo.url = ""
            
        },

        //表格内删除button对应
        async handleDelete(id) {
            this.$messageBox.confirm('确认删除该资源?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                //删除资源
                const {data:res} = await this.$axios.delete("/admin/t/deleteresource", {params: {id: id}});
                if (res.status !== 101) {
                    this.$message.error("删除失败，请重试！")
                } else {
                    this.$message.success("删除成功！")

                }
                if (this.queryInfo.pagenum === Math.ceil(this.total / this.queryInfo.pagesize) && this.links.length === 1) {
                    this.queryInfo.pagenum -= 1
                    if(this.queryInfo.pagenum <= 0) {
                        this.queryInfo.pagenum = 1
                        return
                    }
                }

                //刷新列表
                await this.getLinkList();
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });
        },
        handleSizeChange: function(pagesize) {     // 监听pagesize 改变的事件
            this.queryInfo.pagesize = pagesize;
            this.getLinkList();
        },
        handleCurrentChange: function(newPage) {  // 页码值发送变化
            this.queryInfo.pagenum = newPage;
            this.getLinkList();
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

            axios.post('/admin/t/uploadresource', formData)
                .then(response => {
                this.postInfo.url = response.data.url
                this.$message.success("上传成功！")
                })
                .catch(error => {
                console.log(error);
                });
        },

        

        async commitLink() {
            //若文件名格式错误报错
            if(!this.postInfo.name.endsWith(".zip")) {
                this.$message.error("请在文件名后添加.zip后缀！")
                return
            }
            
            let res
            if(this.postInfo.ID === 0) {
                //若未上传文件报错
                if(this.postInfo.url == "" ) {
                    this.$message.error("请先上传文件！")
                    return
                }
                res = await this.$axios.post("/admin/t/addresource",  { name: this.postInfo.name, desc: this.postInfo.desc, categoryid: this.postInfo.categoryid, url: this.postInfo.url })
            } else {
                res = await this.$axios.put("/admin/t/updateresource", { ID: this.postInfo.ID, name: this.postInfo.name, desc: this.postInfo.desc, categoryid: this.postInfo.categoryid, url: this.postInfo.url })
            }
            if (res.data.status !== 101) {
                this.$message.error("操作失败，请重试！")
            } else {
                this.cancel()
                await this.getLinkList()
                this.$message.success("操作成功！")
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
            this.postInfo.url = response;
        },
        //显示资源分类信息
        showCategories() {
            this.dialogCategoryVisible = true
        },
        addCategoryCancel() {
            this.categoryPost.name = ''
            this.categoryPost.id = 0
            this.innerVisible = false
        },


        async commitCategory() {
            let res
            if(this.categoryPost.id === 0) {
                res = await this.$axios.post("/admin/addCategory", this.categoryPost);
            } else {
                res = await this.$axios.put("/admin/updateCategory", this.categoryPost)
            }

            if (res.data.status === 101) {
                this.$message.success("操作成功")
                this.addCategoryCancel()
                await this.getAllCategories()
            } else {
                this.$message.error("添加失败，请重试！")
            }
        },
        editCategory(index) {
            this.categoryPost.name = this.categories[index].name
            this.categoryPost.id = this.categories[index].id
            this.innerVisible = true
        },
        async deleteCategory(id) {
            this.$messageBox.confirm('确认删除该资源?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                const {data:res} = await this.$axios.delete("/admin/deleteCategory", {params: {id: id}});
                if (res.status === 401) {
                    this.$message.success("删除成功")
                } else {
                    this.$message.error("添加失败，请重试！")
                }
                //刷新列表
                await this.getAllCategories();
            }, () => {
                this.$message({type: 'info', message: '已取消删除'});
            });

        }
    },
    created() {
        this.getLinkList()
    }
}
</script>

<style scoped>

.add {
    float: right !important;
    margin-right: 50px;
    background-color: #baaaca1a;
    color: #baaacaee;
    border-color: #baaacaee;
}

.add:hover, .add:focus, .show-categories:hover, .show-categories:focus {
    background-color: #baaacaee;
    color: #fff;
    border-color: #baaacaee;
}

.show-categories {
    float: right;
    margin-right: 30px;
    background-color: #baaaca1a;
    color: #baaacaee;
    border-color: #baaacaee;
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

.del {
    background-color: #f6727218;
    color: #f67272ac;
    border-color: #f67272ac;
}

.del:hover {
    background-color: #f67272ac;
    color: #fff;
    border-color: #f67272ac;
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