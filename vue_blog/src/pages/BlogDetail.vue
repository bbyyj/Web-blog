
<template>

    <div class="blog_bg">
        <!-- 标题 -->
        <div class="title">
            <h1>{{ blog.title }}</h1>
        </div>

        <!-- 博客信息 -->
        <div class="blogInfo">
            <span class="flag">{{ blog.flag }}</span>
            <span>
                <b-icon icon="person" font-scale="0.9"></b-icon>
                {{ blog.nickname }}
            </span>
            <span>
                <b-icon icon="calendar2-check" font-scale="1"></b-icon>
                {{ (blog.createTime || '').split('T')[0] }}
            </span>
            <span>
                <b-icon icon="journals" font-scale="0.8"></b-icon>
                {{ blog.typename }}
            </span>
            <span>
                <b-icon icon="eye" font-scale="1"></b-icon>
                {{ blog.views }}
            </span>
        </div>

        <!-- 正文 -->
        <div class="content-outer">
            <div class="tags">
                <b-icon icon="tags" font-scale="1.2"></b-icon>
                <span :key="item.id" v-for="item in tags"> {{ item.name }} </span>
            </div>
            <div class="content-inner" v-html="blog.content">

            </div>

            <!-- 赞赏 -->
            <div align="center" v-if="blog.appreciation">
                <b-button id="popover-button-variant" pill variant="outline-info">赞赏</b-button>
                <b-popover target="popover-button-variant"  placement="bottom" variant="info" triggers="focus">
                    <img :src="require('../assets/alipay1.jpg')" width="120" height="140" >
                    <img :src="require('../assets/wechat1.png')" width="120" height="140" >
                </b-popover>
            </div>
        </div>

        <!-- 版权声明 -->
        <div class="rights">
            <ul class="list">
                <li>作者：{{blog.nickname}}</li>
                <li>更新时间：{{ blog.updateTime | fromatDate("yyyy-MM-dd hh:mm:ss") }} </li>
                <li>版权声明：自由转载-非商用-非衍生-保持署名</li>
                <li>转载声明：如果是转载栈主转载的文章，请附上原文链接</li>
            </ul>
        </div>

        <!-- 评论 -->
        <div class="comment-outer">
             <div class="comment-body">
                <div style="font-weight: bold">评论</div>
                <hr>

                <!-- 评论内容 -->
                <div class="comment" :key="item.id" v-for="item in commentList">

                    <b-avatar class="comment-avatar" :src="item.avatar"></b-avatar>
                    <!-- 昵称和日期 -->
                    <!-- <div class="comment-author"> -->
                        <span class="author">
                            {{ item.nickname }}
                        </span>
                        <span class="date">
                            {{ item.createTime | fromatDate("yyyy-MM-dd hh:mm:ss") }}
                        </span>
                    <!-- </div> -->
                    <!-- 回复内容 -->
                    <div class="comment-content">
                        {{ item.content }}
                    </div>
                </div>

            </div>

            <!-- 回复栏 -->
            <textarea class="replay" v-model="comment.content">
            </textarea>

            <b-container fluid>
                <b-row >
                    <b-col sm="3" style="padding-left: 0;">
                        <b-form-input  type="text" placeholder="昵称" v-model="comment.nickname">
                        </b-form-input>
                    </b-col>
                    <b-col sm="3">
                        <b-form-input  type="email" placeholder="邮箱" v-model="comment.email">
                        </b-form-input>
                    </b-col>
                    <b-col sm="3">

                        <b-button variant="outline-info" @click="publishComment">
                            <b-icon icon="pencil-square" font-scale="0.9"></b-icon>
                            发布
                        </b-button>
                    </b-col>
                </b-row>
            </b-container>
        </div>

        <div style="height: 200px;"></div>
    </div>

</template>



<script>
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

<style>

/*改变博客中的图片的最大宽度*/
.blog_image img {
    max-width: 1000px !important;
    cursor: pointer;
}

</style>

<style lang="less" scoped>

.title {
    color: #FFFFFF;
    font-size: 40px !important;
    font-family: STSong;
    text-align: center;
    margin-top: 20px;
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
        border-color: #3d3952;
        border: 1px solid;
        box-shadow: none;
        padding: 2px;
        border-radius: .28571429rem;
        font-size: 14px;
    }

}


.content-outer {
    width: 76%;
    min-height: 800px;
    margin-top: 60px;
    margin-left: 12%;
    margin-right: 12%;

    top: 0;
    bottom: 0;
    border-radius: .28571429rem;
    box-shadow: none;
    border: 1px solid #D4D4D5;
    padding: 1.5em;


    background: #FFF;
    opacity: 0.9;

}

.content-inner {
    padding: 4em 40px 2em;
    box-sizing: border-box;
    font: 1em/1.5 Tahoma,Helvetica,Arial,'华文中宋',sans-serif !important;
    // font-weight: normal;
    color: #3d3952e7;
    height: auto;

}

.tags {
    float: right;
    padding-top: 3px;
    padding-bottom: 3px;
    margin-right: 10px;

    span {
        margin-left: 12px;
        background-color: #FFF;
        color: #565ca4;
        border-color: #3d3952;
        border: 1px solid;
        box-shadow: none;
        padding: 2px;
        border-radius: .28571429rem;
        font-size: 14px;
    }
}

.rights {
    width: 76%;
    margin-left: 12%;
    margin-right: 12%;
    border-radius: .28571429rem;
    box-shadow: 0 0 0 1px #a3c293 inset;
    border: 1px solid #a3c293;
    padding: 1.5em;


    background: #FCFFF5;
    color: #565ca4;
    opacity: 0.9;

    .list {
        text-align: left;
        padding: 0;
        opacity: .85;
        list-style-position: inside;
        margin-bottom: 0.8em;

    }
}

.comment-outer {
    width: 76%;
    margin-left: 12%;
    margin-right: 12%;
    border-radius: .28571429rem;
    box-shadow: none;
    border: 1px solid #D4D4D5;
    padding: 1.5em;

    min-height: 300px;
    background: #FFF;
    opacity: 0.9;


    //margin-bottom: 60px;
}

.comment-body {
    position: relative;
    background: #FFF;
    box-shadow: 0 1px 2px 0 rgb(34 36 38 / 15%);
    margin: 0.5rem 0;
    padding: 1em;
    border-radius: .28571429rem;
    border: 1px solid rgba(34,36,38,.15);
    border-top: 2px solid #565ca4;
    min-height: 260px;

}

.comment {
    width: 90%;
    padding: 6px 18px;
    margin-left: 10px;
    margin-right: 20px;
    margin-top: 12px;
    display: block;

    .comment-avatar {
        float: left;
        margin-top: 6px;
        margin-right: 12px;
    }

    .author {
        font-size: 1.1em;
        color: rgba(0,0,0,.87);
        font-weight: 700;
        margin-left: 8px;
    }

    .date {
        display: inline-block;
        color: rgba(0,0,0,.4);
        font-size: .875em;
        margin-left: 8px;
    }


    .comment-content {
        margin-top: 4px;
        font-size: 15px;
        word-wrap: break-word;
        color: rgba(0,0,0,.87);
        line-height: 1.3;
        padding-left: 60px;
    }
}

.replay {
    width: 100%;
    min-height: 180px;
    background: #FFF;
    margin: 0.5rem 0;
    padding: 1em;
    border-radius: .28571429rem;

    padding: .78571429em 1em;
    border: 1px solid rgba(34,36,38,.15);
    outline: 0;
    color: rgba(0,0,0,.87);
    box-shadow: 0 0 0 0 transparent inset;
    transition: color .1s ease,border-color .1s ease;
    font-size: 1em;
    line-height: 1.2857;
    resize: vertical;
}

.blog_bg {
    min-height: 1000px;
    padding-top: 8%;
    background-color: #b5b0cc;
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='260' height='260' viewBox='0 0 260 260'%3E%3Cg fill-rule='evenodd'%3E%3Cg fill='%236c5c74' fill-opacity='0.5'%3E%3Cpath d='M24.37 16c.2.65.39 1.32.54 2H21.17l1.17 2.34.45.9-.24.11V28a5 5 0 0 1-2.23 8.94l-.02.06a8 8 0 0 1-7.75 6h-20a8 8 0 0 1-7.74-6l-.02-.06A5 5 0 0 1-17.45 28v-6.76l-.79-1.58-.44-.9.9-.44.63-.32H-20a23.01 23.01 0 0 1 44.37-2zm-36.82 2a1 1 0 0 0-.44.1l-3.1 1.56.89 1.79 1.31-.66a3 3 0 0 1 2.69 0l2.2 1.1a1 1 0 0 0 .9 0l2.21-1.1a3 3 0 0 1 2.69 0l2.2 1.1a1 1 0 0 0 .9 0l2.21-1.1a3 3 0 0 1 2.69 0l2.2 1.1a1 1 0 0 0 .86.02l2.88-1.27a3 3 0 0 1 2.43 0l2.88 1.27a1 1 0 0 0 .85-.02l3.1-1.55-.89-1.79-1.42.71a3 3 0 0 1-2.56.06l-2.77-1.23a1 1 0 0 0-.4-.09h-.01a1 1 0 0 0-.4.09l-2.78 1.23a3 3 0 0 1-2.56-.06l-2.3-1.15a1 1 0 0 0-.45-.11h-.01a1 1 0 0 0-.44.1L.9 19.22a3 3 0 0 1-2.69 0l-2.2-1.1a1 1 0 0 0-.45-.11h-.01a1 1 0 0 0-.44.1l-2.21 1.11a3 3 0 0 1-2.69 0l-2.2-1.1a1 1 0 0 0-.45-.11h-.01zm0-2h-4.9a21.01 21.01 0 0 1 39.61 0h-2.09l-.06-.13-.26.13h-32.31zm30.35 7.68l1.36-.68h1.3v2h-36v-1.15l.34-.17 1.36-.68h2.59l1.36.68a3 3 0 0 0 2.69 0l1.36-.68h2.59l1.36.68a3 3 0 0 0 2.69 0L2.26 23h2.59l1.36.68a3 3 0 0 0 2.56.06l1.67-.74h3.23l1.67.74a3 3 0 0 0 2.56-.06zM-13.82 27l16.37 4.91L18.93 27h-32.75zm-.63 2h.34l16.66 5 16.67-5h.33a3 3 0 1 1 0 6h-34a3 3 0 1 1 0-6zm1.35 8a6 6 0 0 0 5.65 4h20a6 6 0 0 0 5.66-4H-13.1z'/%3E%3Cpath id='path6_fill-copy' d='M284.37 16c.2.65.39 1.32.54 2H281.17l1.17 2.34.45.9-.24.11V28a5 5 0 0 1-2.23 8.94l-.02.06a8 8 0 0 1-7.75 6h-20a8 8 0 0 1-7.74-6l-.02-.06a5 5 0 0 1-2.24-8.94v-6.76l-.79-1.58-.44-.9.9-.44.63-.32H240a23.01 23.01 0 0 1 44.37-2zm-36.82 2a1 1 0 0 0-.44.1l-3.1 1.56.89 1.79 1.31-.66a3 3 0 0 1 2.69 0l2.2 1.1a1 1 0 0 0 .9 0l2.21-1.1a3 3 0 0 1 2.69 0l2.2 1.1a1 1 0 0 0 .9 0l2.21-1.1a3 3 0 0 1 2.69 0l2.2 1.1a1 1 0 0 0 .86.02l2.88-1.27a3 3 0 0 1 2.43 0l2.88 1.27a1 1 0 0 0 .85-.02l3.1-1.55-.89-1.79-1.42.71a3 3 0 0 1-2.56.06l-2.77-1.23a1 1 0 0 0-.4-.09h-.01a1 1 0 0 0-.4.09l-2.78 1.23a3 3 0 0 1-2.56-.06l-2.3-1.15a1 1 0 0 0-.45-.11h-.01a1 1 0 0 0-.44.1l-2.21 1.11a3 3 0 0 1-2.69 0l-2.2-1.1a1 1 0 0 0-.45-.11h-.01a1 1 0 0 0-.44.1l-2.21 1.11a3 3 0 0 1-2.69 0l-2.2-1.1a1 1 0 0 0-.45-.11h-.01zm0-2h-4.9a21.01 21.01 0 0 1 39.61 0h-2.09l-.06-.13-.26.13h-32.31zm30.35 7.68l1.36-.68h1.3v2h-36v-1.15l.34-.17 1.36-.68h2.59l1.36.68a3 3 0 0 0 2.69 0l1.36-.68h2.59l1.36.68a3 3 0 0 0 2.69 0l1.36-.68h2.59l1.36.68a3 3 0 0 0 2.56.06l1.67-.74h3.23l1.67.74a3 3 0 0 0 2.56-.06zM246.18 27l16.37 4.91L278.93 27h-32.75zm-.63 2h.34l16.66 5 16.67-5h.33a3 3 0 1 1 0 6h-34a3 3 0 1 1 0-6zm1.35 8a6 6 0 0 0 5.65 4h20a6 6 0 0 0 5.66-4H246.9z'/%3E%3Cpath d='M159.5 21.02A9 9 0 0 0 151 15h-42a9 9 0 0 0-8.5 6.02 6 6 0 0 0 .02 11.96A8.99 8.99 0 0 0 109 45h42a9 9 0 0 0 8.48-12.02 6 6 0 0 0 .02-11.96zM151 17h-42a7 7 0 0 0-6.33 4h54.66a7 7 0 0 0-6.33-4zm-9.34 26a8.98 8.98 0 0 0 3.34-7h-2a7 7 0 0 1-7 7h-4.34a8.98 8.98 0 0 0 3.34-7h-2a7 7 0 0 1-7 7h-4.34a8.98 8.98 0 0 0 3.34-7h-2a7 7 0 0 1-7 7h-7a7 7 0 1 1 0-14h42a7 7 0 1 1 0 14h-9.34zM109 27a9 9 0 0 0-7.48 4H101a4 4 0 1 1 0-8h58a4 4 0 0 1 0 8h-.52a9 9 0 0 0-7.48-4h-42z'/%3E%3Cpath d='M39 115a8 8 0 1 0 0-16 8 8 0 0 0 0 16zm6-8a6 6 0 1 1-12 0 6 6 0 0 1 12 0zm-3-29v-2h8v-6H40a4 4 0 0 0-4 4v10H22l-1.33 4-.67 2h2.19L26 130h26l3.81-40H58l-.67-2L56 84H42v-6zm-4-4v10h2V74h8v-2h-8a2 2 0 0 0-2 2zm2 12h14.56l.67 2H22.77l.67-2H40zm13.8 4H24.2l3.62 38h22.36l3.62-38z'/%3E%3Cpath d='M129 92h-6v4h-6v4h-6v14h-3l.24 2 3.76 32h36l3.76-32 .24-2h-3v-14h-6v-4h-6v-4h-8zm18 22v-12h-4v4h3v8h1zm-3 0v-6h-4v6h4zm-6 6v-16h-4v19.17c1.6-.7 2.97-1.8 4-3.17zm-6 3.8V100h-4v23.8a10.04 10.04 0 0 0 4 0zm-6-.63V104h-4v16a10.04 10.04 0 0 0 4 3.17zm-6-9.17v-6h-4v6h4zm-6 0v-8h3v-4h-4v12h1zm27-12v-4h-4v4h3v4h1v-4zm-6 0v-8h-4v4h3v4h1zm-6-4v-4h-4v8h1v-4h3zm-6 4v-4h-4v8h1v-4h3zm7 24a12 12 0 0 0 11.83-10h7.92l-3.53 30h-32.44l-3.53-30h7.92A12 12 0 0 0 130 126z'/%3E%3Cpath d='M212 86v2h-4v-2h4zm4 0h-2v2h2v-2zm-20 0v.1a5 5 0 0 0-.56 9.65l.06.25 1.12 4.48a2 2 0 0 0 1.94 1.52h.01l7.02 24.55a2 2 0 0 0 1.92 1.45h4.98a2 2 0 0 0 1.92-1.45l7.02-24.55a2 2 0 0 0 1.95-1.52L224.5 96l.06-.25a5 5 0 0 0-.56-9.65V86a14 14 0 0 0-28 0zm4 0h6v2h-9a3 3 0 1 0 0 6H223a3 3 0 1 0 0-6H220v-2h2a12 12 0 1 0-24 0h2zm-1.44 14l-1-4h24.88l-1 4h-22.88zm8.95 26l-6.86-24h18.7l-6.86 24h-4.98zM150 242a22 22 0 1 0 0-44 22 22 0 0 0 0 44zm24-22a24 24 0 1 1-48 0 24 24 0 0 1 48 0zm-28.38 17.73l2.04-.87a6 6 0 0 1 4.68 0l2.04.87a2 2 0 0 0 2.5-.82l1.14-1.9a6 6 0 0 1 3.79-2.75l2.15-.5a2 2 0 0 0 1.54-2.12l-.19-2.2a6 6 0 0 1 1.45-4.46l1.45-1.67a2 2 0 0 0 0-2.62l-1.45-1.67a6 6 0 0 1-1.45-4.46l.2-2.2a2 2 0 0 0-1.55-2.13l-2.15-.5a6 6 0 0 1-3.8-2.75l-1.13-1.9a2 2 0 0 0-2.5-.8l-2.04.86a6 6 0 0 1-4.68 0l-2.04-.87a2 2 0 0 0-2.5.82l-1.14 1.9a6 6 0 0 1-3.79 2.75l-2.15.5a2 2 0 0 0-1.54 2.12l.19 2.2a6 6 0 0 1-1.45 4.46l-1.45 1.67a2 2 0 0 0 0 2.62l1.45 1.67a6 6 0 0 1 1.45 4.46l-.2 2.2a2 2 0 0 0 1.55 2.13l2.15.5a6 6 0 0 1 3.8 2.75l1.13 1.9a2 2 0 0 0 2.5.8zm2.82.97a4 4 0 0 1 3.12 0l2.04.87a4 4 0 0 0 4.99-1.62l1.14-1.9a4 4 0 0 1 2.53-1.84l2.15-.5a4 4 0 0 0 3.09-4.24l-.2-2.2a4 4 0 0 1 .97-2.98l1.45-1.67a4 4 0 0 0 0-5.24l-1.45-1.67a4 4 0 0 1-.97-2.97l.2-2.2a4 4 0 0 0-3.09-4.25l-2.15-.5a4 4 0 0 1-2.53-1.84l-1.14-1.9a4 4 0 0 0-5-1.62l-2.03.87a4 4 0 0 1-3.12 0l-2.04-.87a4 4 0 0 0-4.99 1.62l-1.14 1.9a4 4 0 0 1-2.53 1.84l-2.15.5a4 4 0 0 0-3.09 4.24l.2 2.2a4 4 0 0 1-.97 2.98l-1.45 1.67a4 4 0 0 0 0 5.24l1.45 1.67a4 4 0 0 1 .97 2.97l-.2 2.2a4 4 0 0 0 3.09 4.25l2.15.5a4 4 0 0 1 2.53 1.84l1.14 1.9a4 4 0 0 0 5 1.62l2.03-.87zM152 207a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm6 2a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm-11 1a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm-6 0a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm3-5a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm-8 8a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm3 6a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm0 6a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm4 7a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm5-2a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm5 4a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm4-6a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm6-4a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm-4-3a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm4-3a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm-5-4a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm-24 6a1 1 0 1 1 2 0 1 1 0 0 1-2 0zm16 5a5 5 0 1 0 0-10 5 5 0 0 0 0 10zm7-5a7 7 0 1 1-14 0 7 7 0 0 1 14 0zm86-29a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm19 9a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm-14 5a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm-25 1a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm5 4a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm9 0a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm15 1a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm12-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm-11-14a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm-19 0a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm6 5a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm-25 15c0-.47.01-.94.03-1.4a5 5 0 0 1-1.7-8 3.99 3.99 0 0 1 1.88-5.18 5 5 0 0 1 3.4-6.22 3 3 0 0 1 1.46-1.05 5 5 0 0 1 7.76-3.27A30.86 30.86 0 0 1 246 184c6.79 0 13.06 2.18 18.17 5.88a5 5 0 0 1 7.76 3.27 3 3 0 0 1 1.47 1.05 5 5 0 0 1 3.4 6.22 4 4 0 0 1 1.87 5.18 4.98 4.98 0 0 1-1.7 8c.02.46.03.93.03 1.4v1h-62v-1zm.83-7.17a30.9 30.9 0 0 0-.62 3.57 3 3 0 0 1-.61-4.2c.37.28.78.49 1.23.63zm1.49-4.61c-.36.87-.68 1.76-.96 2.68a2 2 0 0 1-.21-3.71c.33.4.73.75 1.17 1.03zm2.32-4.54c-.54.86-1.03 1.76-1.49 2.68a3 3 0 0 1-.07-4.67 3 3 0 0 0 1.56 1.99zm1.14-1.7c.35-.5.72-.98 1.1-1.46a1 1 0 1 0-1.1 1.45zm5.34-5.77c-1.03.86-2 1.79-2.9 2.77a3 3 0 0 0-1.11-.77 3 3 0 0 1 4-2zm42.66 2.77c-.9-.98-1.87-1.9-2.9-2.77a3 3 0 0 1 4.01 2 3 3 0 0 0-1.1.77zm1.34 1.54c.38.48.75.96 1.1 1.45a1 1 0 1 0-1.1-1.45zm3.73 5.84c-.46-.92-.95-1.82-1.5-2.68a3 3 0 0 0 1.57-1.99 3 3 0 0 1-.07 4.67zm1.8 4.53c-.29-.9-.6-1.8-.97-2.67.44-.28.84-.63 1.17-1.03a2 2 0 0 1-.2 3.7zm1.14 5.51c-.14-1.21-.35-2.4-.62-3.57.45-.14.86-.35 1.23-.63a2.99 2.99 0 0 1-.6 4.2zM275 214a29 29 0 0 0-57.97 0h57.96zM72.33 198.12c-.21-.32-.34-.7-.34-1.12v-12h-2v12a4.01 4.01 0 0 0 7.09 2.54c.57-.69.91-1.57.91-2.54v-12h-2v12a1.99 1.99 0 0 1-2 2 2 2 0 0 1-1.66-.88zM75 176c.38 0 .74-.04 1.1-.12a4 4 0 0 0 6.19 2.4A13.94 13.94 0 0 1 84 185v24a6 6 0 0 1-6 6h-3v9a5 5 0 1 1-10 0v-9h-3a6 6 0 0 1-6-6v-24a14 14 0 0 1 14-14 5 5 0 0 0 5 5zm-17 15v12a1.99 1.99 0 0 0 1.22 1.84 2 2 0 0 0 2.44-.72c.21-.32.34-.7.34-1.12v-12h2v12a3.98 3.98 0 0 1-5.35 3.77 3.98 3.98 0 0 1-.65-.3V209a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4v-24c.01-1.53-.23-2.88-.72-4.17-.43.1-.87.16-1.28.17a6 6 0 0 1-5.2-3 7 7 0 0 1-6.47-4.88A12 12 0 0 0 58 185v6zm9 24v9a3 3 0 1 0 6 0v-9h-6z'/%3E%3Cpath d='M-17 191a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm19 9a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H3a1 1 0 0 1-1-1zm-14 5a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm-25 1a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm5 4a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm9 0a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm15 1a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm12-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2H4zm-11-14a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm-19 0a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2h-2zm6 5a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1zm-25 15c0-.47.01-.94.03-1.4a5 5 0 0 1-1.7-8 3.99 3.99 0 0 1 1.88-5.18 5 5 0 0 1 3.4-6.22 3 3 0 0 1 1.46-1.05 5 5 0 0 1 7.76-3.27A30.86 30.86 0 0 1-14 184c6.79 0 13.06 2.18 18.17 5.88a5 5 0 0 1 7.76 3.27 3 3 0 0 1 1.47 1.05 5 5 0 0 1 3.4 6.22 4 4 0 0 1 1.87 5.18 4.98 4.98 0 0 1-1.7 8c.02.46.03.93.03 1.4v1h-62v-1zm.83-7.17a30.9 30.9 0 0 0-.62 3.57 3 3 0 0 1-.61-4.2c.37.28.78.49 1.23.63zm1.49-4.61c-.36.87-.68 1.76-.96 2.68a2 2 0 0 1-.21-3.71c.33.4.73.75 1.17 1.03zm2.32-4.54c-.54.86-1.03 1.76-1.49 2.68a3 3 0 0 1-.07-4.67 3 3 0 0 0 1.56 1.99zm1.14-1.7c.35-.5.72-.98 1.1-1.46a1 1 0 1 0-1.1 1.45zm5.34-5.77c-1.03.86-2 1.79-2.9 2.77a3 3 0 0 0-1.11-.77 3 3 0 0 1 4-2zm42.66 2.77c-.9-.98-1.87-1.9-2.9-2.77a3 3 0 0 1 4.01 2 3 3 0 0 0-1.1.77zm1.34 1.54c.38.48.75.96 1.1 1.45a1 1 0 1 0-1.1-1.45zm3.73 5.84c-.46-.92-.95-1.82-1.5-2.68a3 3 0 0 0 1.57-1.99 3 3 0 0 1-.07 4.67zm1.8 4.53c-.29-.9-.6-1.8-.97-2.67.44-.28.84-.63 1.17-1.03a2 2 0 0 1-.2 3.7zm1.14 5.51c-.14-1.21-.35-2.4-.62-3.57.45-.14.86-.35 1.23-.63a2.99 2.99 0 0 1-.6 4.2zM15 214a29 29 0 0 0-57.97 0h57.96z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
}


//// 多行文本溢出显示省略号
//.text-ellipsis {
//    //弹性伸缩盒子模型
//    display: -webkit-box;
//    //限制在一个块元素显示的文本行数
//    -webkit-line-clamp: 3;
//    //设置或检索伸缩盒子对象的子元素的排列方式
//    -webkit-box-orient: vertical;
//}

</style>




