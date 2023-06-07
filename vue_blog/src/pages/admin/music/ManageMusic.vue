<template>
    <div>
        <el-card>
            <el-row>
                <el-button class="add" type="primary" plain @click="handleAdd">添加音乐</el-button>
            </el-row>
            <el-table :data="musicList" border stripe>
                <el-table-column type="index"></el-table-column>
                <el-table-column label="name" prop="name" width="300px"></el-table-column>
                <el-table-column label="artist" prop="artist"></el-table-column>
                <el-table-column label="url" prop="url" width="180px"></el-table-column>
                <el-table-column label="cover" prop="cover"></el-table-column>
                <el-table-column>
                    <template slot-scope="scope">
                        <el-button size="mini" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="queryInfo.pageNum" :page-sizes="[10, 12, 15, 20]" :page-size="queryInfo.pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>

            <el-dialog title="新的音乐信息" :visible.sync="dialogFormVisible">
                <el-form label-width="80px">
                    <el-form-item label="name:">
                        <el-input v-model="postInfo.name" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="artist:">
                        <el-input v-model="postInfo.artist" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="url:">
                        <el-input v-model="postInfo.url" autocomplete="off" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="cover:">
                        <el-input v-model="postInfo.cover" autocomplete="off" clearable></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="cancel">取 消</el-button>
                    <el-button type="primary" @click="commitMusic">确 定</el-button>
                </div>
            </el-dialog>
        </el-card>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: "manageMusic",
    data() {
        return {
            musicList: [],
            dialogFormVisible: false,
            total: 0,
            postInfo: {
                name: "",
                artist: "",
                url: "",
                cover: "",
            },
            queryInfo: {
                pageNum: 1,
                pageSize: 10
            },
        }
    },
    methods: {
        //获取歌单列表
        async getMusicList() {
            const { data: res } = await this.$axios.get("/admin/musicList", {
                params: {
                    pageNum: this.queryInfo.pageNum,
                    pageSize: this.queryInfo.pageSize
                }
            });
            if (res.status === 1) {
                this.musicList = res.data[0];
                this.total = res.data[1];
            } else {
                this.$message.error("获取音乐列表失败，请重试！")
            }
        },
        //显示新增信息的窗口
        async handleAdd() {
            this.dialogFormVisible = true
        },
        //删除音乐
        async handleDelete(id) {
            this.$messageBox.confirm('确认删除该音乐?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                const { data: res } = await this.$axios.delete("/admin/deleteMusic", { params: { id: id } });
                if (res.status !== 401) {
                    this.$message.error("删除失败，请重试！")
                } else {
                    this.$message.success("删除成功！")
                    await this.getMusicList();
                }
            }, () => {
                this.$message({ type: 'info', message: '已取消删除' });
            });
        },
        //////////////////分页相关的
        handleSizeChange: function (pagesize) {
            this.queryInfo.pageSize = pagesize;
            this.getMusicList();
        },
        handleCurrentChange: function (newPage) {
            this.queryInfo.pageNum = newPage;
            this.getMusicList();
        },
        //取消新增操作
        cancel() {
            this.postInfo = {
                name: "",
                artist: "",
                url: "",
                cover: "",
            }
            this.dialogFormVisible = false
        },
        //增加新的歌曲内容
        async commitMusic() {
            const { data: res } = await this.$axios.post("/admin/addMusic", this.postInfo)

            if (res.status !== 101) {
                this.$message.error("操作失败，请重试！")
            } else {
                this.cancel()
                //获取到新的歌单列表
                await this.getMusicList()
                this.$message.success("操作成功！")
            }
        },
    },
    created() {
        this.getMusicList()
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

.add:hover, .add:focus {
    background-color: #baaacaee;
    color: #fff;
    border-color: #baaacaee;
}

</style>
