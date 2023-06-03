<template>
    <div>
        <el-card>
            <!-- 筛选栏 -->
            <div class="filter-container">
                <el-select v-model="selectedCategory" placeholder="选择类别" @change="getSubject" clearable>
                    <el-option v-for="item in categories" :key="item" :label="item" :value="item">
                    </el-option>
                </el-select>
                <el-select v-model="selectedSubject" placeholder="选择科目" @change="getComments" clearable>
                    <el-option v-for="item in subjects" :key="item.subject_id" :label="item.subject_name"
                        :value="item.subject_name">
                    </el-option>
                </el-select>
                <!-- 新增题目按钮 -->
            </div>
            <!-- 列表区域 -->
            <el-table :data="formattedComments" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="课程编号" prop="subject_id" width="300px"></el-table-column>
                <el-table-column label="课程名称" prop="subject_name" width="300px"></el-table-column>
                <el-table-column label="时间" prop="createTime"></el-table-column>
                <el-table-column label="评论内容" prop="comment"></el-table-column>
                <el-table-column label="操作" width="150">
                    <template slot-scope="scope">
                        <!-- 编辑和删除按钮 -->
                        <el-button class="del" size="mini" @click="handleDelete(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <!-- 分页区域 -->
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.pageNum" :page-sizes="[10, 12, 15, 20]" :page-size="queryInfo.pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </el-card>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: "ManageComments",
    data() {
        return {
            total: 0,
            queryInfo: {
                pageNum: 1,
                pageSize: 10,
            },
            categories: ['公必', '公选', '专必', '专选'], //类型列表
            subjects: [], //科目列表
            selectedCategory: '',   //选中的类别
            selectedSubject: '', //选中的科目
            comments: []  //获取到的评论
        };
    },
    computed: {
        formattedComments() {
            return this.comments.map(comment => {
                return {
                    ...comment,
                    createTime: this.formatDate(comment.createTime, "yyyy-MM-dd hh:mm:ss")
                };
            });
        }
    },
    watch: {
        selectedCategory(newCategory) {
            // 当 selectedCategory 变化时，设置 selectedSubject 为空
            this.selectedSubject = '';
            // 调用 getSubject 方法获取科目列表
            if (newCategory) {
                this.getSubject();
            }
        },

    },
    created() {
        this.getALLcomments()
    },
    methods: {
        //格式化日期
        formatDate(val, fmt) {
            var date = new Date(val);
            let ret;
            const opt = {
                "y+": date.getFullYear().toString(),        // 年
                "M+": (date.getMonth() + 1).toString(),     // 月
                "d+": date.getDate().toString(),            // 日
                "h+": date.getHours().toString(),           // 时
                "m+": date.getMinutes().toString(),         // 分
                "s+": date.getSeconds().toString()          // 秒
            };
            for (let k in opt) {
                ret = new RegExp("(" + k + ")").exec(fmt);
                if (ret) {
                    fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
                };
            };
            return fmt;
        },
        //开始时候获取全部的评论进行渲染
        async getALLcomments() {
            try {
                const { data: res } = await this.$axios.get("/admin/electionCommentList", {
                    params: {
                        pageNum: this.queryInfo.pageNum,
                        pageSize: this.queryInfo.pageSize
                    }
                });
                this.comments = res.data[0]; // 根据你的后端返回的数据结构进行修改
                this.total = res.data[1]
            } catch (error) {
                console.error(error);
            }
        },

        // 监听pagesize 改变的事件
        handleSizeChange: function (pagesize) {
            this.queryInfo.pageSize = pagesize;
            this.getComments()
        },
        // 页码值发送变化
        handleCurrentChange: function (newPage) {
            this.queryInfo.pageNum = newPage;
            this.getComments()
        },
        //删除的操作
        async handleDelete(index) {
            this.$messageBox.confirm('确认删除该条内容?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                //删除对应问答
                const { data: res } = await this.$axios.delete("/admin/deleteElectionComment", { params: { id: index } });
                if (res.status !== 101) {
                    this.$message.error("删除失败，请重试！")
                } else {
                    this.$message.success("删除成功！")
                }
                //进行页码等相关转换操作
                if (this.queryInfo.pageNum === Math.ceil(this.total / this.queryInfo.pageSize) && this.comments.length === 1) {
                    this.queryInfo.pageNum -= 1
                    if (this.queryInfo.pageNum <= 0) {
                        this.queryInfo.pageNum = 1
                        return
                    }
                }
                //刷新列表
                await this.getComments();
            }, () => {
                //取消删除操作
                this.$message({ type: 'info', message: '已取消删除' });
            });
        },
        //获取全部学科列表 //已经完成/////////////////////////
        async getSubject() {
            if (!this.selectedCategory) {
                this.subjects = ""
                return

            }  // 没有选择分类则不发送请求
            const { data: res } = await this.$axios.get("/admin/electionNoPage", {
                params: {
                    classification: this.selectedCategory
                }
            });
            if (res.status === 1) {
                this.subjects = res.data[0];
            }
            else {
                window.alert("获取科目失败，请检查网络设置！")
            }
        },
        //获取对应类型+学科的全部评论
        async getComments() {
            console.log("test")
            const { data: res } = await this.$axios.get("/admin/electionByClassification", {
                params: {
                    classification: this.selectedCategory,
                    pageNum: this.queryInfo.pageNum,
                    pageSize: this.queryInfo.pageSize,
                    subject_name: this.selectedSubject,
                }
            });
            if (res.status === 1) {
                //赋值题目
                this.total = res.data[1]
                this.comments = res.data[0];
            }
            else {
                this.$message.error("获取题目失败,请重试");
            }
        }
    },
}
</script>


<style>

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

</style>