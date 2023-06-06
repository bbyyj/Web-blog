
<template>

    <div class="blog_bg">
        <div id='stars'></div>
        <div id='stars2'></div>
        <div id='stars3'></div>

        <!-- 标题 -->
        <div class="title">
            <h1>{{ blog.title }}</h1>
        </div>

        <!-- 博客信息 -->
        <div class="blogInfo">
            <span>
                <i class="iconfont icon-yonghu-yuan"></i>
                {{ blog.nickname }}
            </span>
            <span>
                <i class="iconfont icon-rili"></i>
                {{ (blog.createTime || '').split('T')[0] }}
            </span>
            <span>
                <i class="iconfont icon-chakan"></i>
                {{ blog.views }}
            </span>
            <br><div class="space"></div>

            <i class="iconfont icon-xitongguanli"></i>

            <span class="tags">    
                <span class="tag1">{{ blog.flag }}</span>
                <span class="tag2">{{ blog.typename }}</span> 
                <span class="tag3" :key="item.id" v-for="item in tags"> {{ item.name }} </span>
            </span>
            
        </div>

        <!-- 正文 -->
        <div class="content-outer">
           <div class="content-inner" v-html="blog.content">
            </div>
        </div>

    </div>

</template>



<script

>import '../assets/css/star.css'

export default {
    data() {
        return {
            blog: {
                id: 0,
                title: "上邪~~",
                content: `上邪,<br> 我欲与君相知，<br>长命无绝衰。<br>山无陵，<br>江水为竭。<br>冬雷震震，<br>夏雨雪。<br>天地合，<br>乃敢与君绝。`,
                flag: "原创",
                updateTime: "2021-03-30",
                createTime: "2021-03-30",
                views: 1024,
                appreciation: false,
                typename: "框架底层原理",
                nickname: "Jack Ma"
            },
            tags: [
                {id: 1, name: "Java"},
                {id: 2, name: "Spring"},
                {id: 3, name: "框架"}
            ],
            comment: {         //评论数据
                nickname: "",
                email: "",
                content: "",
                avatar: "http://localhost:8080/images/firstPic/avatar1.png",
                blogId: 0
            },
            commentList: [
                {
                    id: 1,
                    nickname: "小兔叽",
                    email: "rabbit@blog.com",
                    content: "very good!",
                    avatar: "http://localhost:8080/firstPic/avatar1.png",
                    createTime: "2021-03-31 22:35",
                    blogId: 0,
                },
                {
                    id: 2,
                    nickname: "小脑斧",
                    email: "tiger@blog.com",
                    content: "excellent!",
                    avatar: "http://localhost:8080/firstPic/avatar1.png",
                    createTime: "2021-03-31 22:35",
                    blogId: 0,
                }
            ]
        }
    },
    created() {
        window.scrollTo(0, 0);
        this.getBlogDetails();
    },
    filters: {
        fromatDate: function(val, fmt) {
            var date = new Date(val);
            let ret;
            const opt = {
                "y+": date.getFullYear().toString(),        // 年
                "M+": (date.getMonth() + 1).toString(),     // 月
                "d+": date.getDate().toString(),            // 日
                "h+": date.getHours().toString(),           // 时
                "m+": date.getMinutes().toString(),         // 分
                "s+": date.getSeconds().toString()          // 秒
                // 有其他格式化字符需求可以继续添加，必须转化成字符串
            };
            for (let k in opt) {
                ret = new RegExp("(" + k + ")").exec(fmt);
                if (ret) {
                    fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
                };
            };
            return fmt;
        }
    },
    methods: {
        getBlogDetails: async function() {
            let blogId = this.$route.query.id;
            blogId = parseInt(blogId);
            const {data: res} = await this.$axios.get("/myblog/detailedBlog", {params: {id: blogId}});
            if(res.status === 1) {
                this.blog = res.data.length > 0 ? res.data[0] : this.blog;
                this.tags = res.data.length > 1 ? res.data[1] : this.tags;
                var hljs = require('highlight.js');
                var md = require('markdown-it')({
                                    html: true,
                                    linkify: true,
                                    typographer: true,
                                    highlight: function (str, lang) {
                                        if (lang && hljs.getLanguage(lang)) {
                                            try {
                                                return '<pre class="hljs"><code>' +
                                                    hljs.highlight(lang, str, true).value +
                                                    '</code></pre>';
                                            } catch (__) {}
                                        }

                                        return '<pre class="hljs"><code>' + md.utils.escapeHtml(str) + '</code></pre>';
                                    }

                });

                // 让图片居中
                md.renderer.rules.image = (token, idx, options, env, self) => {
                    return "<div align='center' class='blog_image'>" + md.renderer.renderToken(token, idx, options, env, self) + "</div>"
                }
                this.blog.content = md.render(this.blog.content);

                this.getCommentList();
            }
        },
        publishComment: async function () {
            this.comment.blogId = this.blog.id;
            if(!this.comment.content) {
                this.$message.warning("请输入评论内容")
                return
            }
            if(!this.comment.nickname) {
                this.$message.warning("请输入您的称呼")
                return
            }
            if(!this.comment.email) {
                this.$message.warning("请输入邮箱")
                return
            }
            if(!this.checkEmail()) {
                this.$message.warning("邮箱不合法")
                return
            }

            await this.$axios.post("/myblog/publishComment", this.comment);
            this.comment.nickname = "";
            this.comment.email = "";
            this.comment.content = "";

            this.getCommentList();
        },
        getCommentList: async function() {
            const {data: res} = await this.$axios.get("/myblog/commentList", {params: {id: this.blog.id}});
            if(res.status === 1) {
                this.commentList = res.data.length > 0 ? res.data[0] : this.commentList;
            }
        },
        checkEmail: function() {
            var reg = new RegExp("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"); //正则表达式
            if(this.comment.email === ""){ //输入不能为空
                alert("输入不能为空!");
                return false;
            }else if(!reg.test(this.comment.email)){ //正则验证不通过，格式不对
                return false;
            }else{
                return true;
            }
        }

    }
}
</script>

<style lang="less" scoped>
.blog_bg {
    min-height: 2000px;
    background-color: #1d1d2b;
    font-family: NotoSerifSC-Regular;
}
.title {
    margin-top: 80px;
    color: #FFFFFF;
    text-align: center;
    h1{
        font-size: 50px;
    }
}
.space{
    height: 20px;
}
.blogInfo {
    color: #FFFFFF;
    text-align: center;
    margin-top: 30px;
    font-size: 18px;
    padding-top: 5px;
    padding-bottom: 5px;

    span {
        margin-left: 12px;
    }
    .flag {
        background-color: #FFF;
        color: #3d3952;
        border: 1px solid;
        box-shadow: none;
        padding: 4px 8px 4px 8px;
        border-radius: 5px;
        font-size: 14px;
    }

}

.content-outer {
    width: 72%;
    min-height: 1000px;
    margin-top: 80px;
    margin-left: 14%;
    margin-right: 14%;

    border-radius: 20px;
    box-shadow: none;
    padding: 1.5em;

    background: #FFF;
    opacity: 0.8;

}

.content-inner {
    padding: 4em 40px 2em;
    box-sizing: border-box;

    color: #1d1d2b;
    height: auto;

}

.icon-xiangqingchakan{
    color: #3d3952;
}
.tags {
    padding-top: 3px;
    padding-bottom: 3px;
    margin-left: 10px;

    span {
        margin-left: 15px;
        background-color: #FFF;
        color: #1d1d2b;
        box-shadow: none;
        padding: 6px 10px 6px 10px;
        border-radius: .28571429rem;
        font-size: 14px;
    }
    .tag1{
        background-color: #ecedf5;
    }
    .tag2{
        background-color: #c7c9e1;
    }
    .tag3{
        background-color: #a1a4cc;
    }
}

.iconfont{
    font-size: 20px !important;
}
</style>




