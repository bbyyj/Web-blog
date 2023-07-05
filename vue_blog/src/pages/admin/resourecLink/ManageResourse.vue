<template>
    <div>
        <el-card>
        <el-table :data="links" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="文件名" prop="name" width="220px"></el-table-column>
                <el-table-column label="描述" prop="desc"></el-table-column>
                <el-table-column label="类别" prop="category" width="150px"></el-table-column>
                <el-table-column label="上传时间">
                    <template slot-scope="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{ dateFormat(scope.row.CreatedAt) }}</span>
                </template>
                </el-table-column>
                <el-table-column label="操作"  width="150">
                <template slot-scope="scope">
                    <el-button class="pas" size="mini" :disabled="scope.row.status === 1" @click="handleStatus(scope.row.ID, 1)">通过</el-button>
                    <el-button class="del" size="mini" :disabled="scope.row.status === 2" type="danger" @click="handleStatus(scope.row.ID, 2)">禁止</el-button>
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

import dayjs from "dayjs";

export default {
    name: "ManageResourse",
    data() {
        return {
            //接收到的数据
            links: [],
            //资源总条数
            total: 0,
            //页码相关
            queryInfo: {
                pagenum: 1,
                pagesize: 10
            },
            //所有分类 包括id和名字
            categories: [],
        }
    },
    methods: {
        async getMessageList() {
            const {data:res} = await this.$axios.get("/admin/t/pageresourcecheck", {params: this.queryInfo});
            console.log(res);
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

        async handleStatus(id, status) {
            let res
            console.log(id);
            if (status==1) {
                res = await this.$axios.put("/admin/t/checksucceeded", {id: id});
            } else {
                res = await this.$axios.delete("/admin/t/checkfailed", {params: {id: id}});
            }
            if (res.data.status !== 101) {
                status === 1 ? this.$message.error("文件通过失败，请重试！") : this.$message.error("文件禁止失败，请重试！")
            } else {
                status === 1 ? this.$message.success("文件通过成功！") : this.$message.success("文件禁止成功！")
                await this.getMessageList()
            }
        },

        handleSizeChange: function(pagesize) {     // 监听pagesize 改变的事件
            this.queryInfo.pagesize = pagesize;
            this.getMessageList();
        },

        handleCurrentChange: function(newPage) {  // 页码值发送变化
            this.queryInfo.pagenum = newPage;
            this.getMessageList();
        },
        
        dateFormat(d) {
            return dayjs(d).format("YYYY-MM-DD HH:mm:ss")
        }
    },
    created() {
        this.getMessageList()
    }
}
</script>

<style scoped>

.pas:hover {
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

</style>