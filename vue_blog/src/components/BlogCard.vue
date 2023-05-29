<!--博客卡片组件-->
<template>
    <div class="box">

        <!--图片区域-->
        <div class="image-box">
            <img :src="item.firstPicture">
        </div>

        <!--博客区域-->
        <div class="blog-box">
            <!--博客标题-->
            <div class="blog-title">
                <a>{{ item.title }}</a>
            </div>

            <!--博客信息-->
            <div class="blog-info">
                <ul>
                    <li>
                        <i class="iconfont icon-yonghu"></i>
                        <span>{{item.nickname}}</span>
                    </li>
                    <li>
                        <i class="iconfont icon-riqi"></i>
                        <span>{{ (item.createTime || '').split('T')[0]}}</span>
                    </li>
                    <li>
                        <i class="iconfont icon-yanjing"></i>
                        <span>{{item.views}}</span>
                    </li>
                    <li>
                        <a href="#">
                            {{item.typename}}
                        </a>
                    </li>
                </ul>
            </div>
            <div class="blog-desc">
                <p class="text-ellipsis">{{item.description}}</p>
            </div>
        </div>
    </div>
</template>

<script>

import "../assets/icon/iconfont.css"

export default {
    name: "BlogCards",
    props: ["item", "imgRight"],
    data() {
        return {
        }
    }

}
</script>


<style lang="less" scoped>

* {
    padding: 0;
    margin: 0;
}

//博客展示item
.box {
    background-color: #d9dbeb;
    opacity: 0.8;
    height: 300px;
    overflow: hidden;
    border-radius: 20px;
    padding: 16px 20px;
    margin: 20px 200px 50px 200px;
    position: relative;
    transition: all .5s;
    cursor: pointer;
}

.box:hover {
    transform:  scale(1.02);
}

//图片区域
.image-box {
    float: left;
    border-radius: 6px;
    margin: 40px 30px;
    cursor: pointer;
    overflow: hidden;
}

//博客信息区域
.blog-box {
    width: 650px;
    height: 250px;
    padding-left: 20px;
}

.image-box img {
    width: 180px;
    height: 180px;
    border-radius: 6px;
    transition: all 0.5s;
}

.image-box img:hover {
    transform: scale(1.1);
}

.box::after{
    position:absolute;
    content:'';
    width:0;
    height:0;
    z-index:-1;
    left: -6px;
    top: 0;
    transition: border-color  0s  ease-in 0.8s , width 0.2s ease-in 0.6s, height 0.2s ease-in 0.4s;
    -webkit-transition: border-color  0s  ease-in 0.8s , width 0.2s ease-in 0.6s, height 0.2s ease-in 0.4s;
}
.box::before{
    position:absolute;
    content:'';
    border:3px dashed transparent;
    box-sizing: border-box;
    width:0;
    height:0;
    z-index:-1;
    bottom: 0;
    right: 0;
    transition: border-color 0s ease-in 0.4s, width 0.2s ease-in 0.2s, height 0.2s ease-in ;
}

.box:hover::before{
    width:100%;
    height:100%;
    transition: width 0.2s ease-out, height 0.2s ease-out 0.2s;
    -webkit-transition:width 0.2s ease-out, height 0.2s ease-out 0.2s;
}


.box .blog-box .blog-title {
    cursor: pointer;
    height: 46px;
    line-height: 46px;
    text-align: center;
    font-size: 1.28rem;
}

.box .blog-box .blog-title a {
    font-weight: 500;
    text-decoration: none;
    color: #3d3952;
    padding-bottom: 5px;
    position: relative;
}

.box .blog-box .blog-title a::before {
    content: '';
    position: absolute;
    bottom: 0;
    right: 50%;
    width: 0;
    transition: 0.5s width ease-out;
}

.box .blog-box .blog-title a::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 50%;
    width: 0;
    transition: 0.5s width ease-out;
}

.box .blog-box .blog-title:hover a::before,
.box .blog-box .blog-title:hover a::after {
    width: 50%;
    border-bottom: 2px solid #7c80b8;
}

.box .blog-box .blog-title:hover a {
    color: #7c80b8;
}

.box .blog-box .blog-desc {
    height: 110px;
    p {
        padding: 12px 24px;
        font-size: 16px;
        line-height: 1.8;
        color: #3d3952;
        max-height: 100px;
    }
}

.box .blog-box .blog-info {
    height: 25px;
    text-align: center;
    margin-top: 5px;

    span {
        margin-left: 10px;
        font-size: 14px;
        font-weight: 700;
        color: #696eae;
    }

    i {
        font-weight: bold;
    }

    .icon-yonghu, .icon-riqi, .icon-yanjing {
        color: #3d3952;
    }

    a {
        font-size: 13px;
        text-decoration: none;
        outline: none;
        border: 1px solid #696eae;
        padding: 4px;
        border-radius: 3px;
        color: #696eae;
    }
}


li {
    list-style: none;
    display: inline-block;
    margin-right: 18px;
}

.text-ellipsis {
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 3;
}

</style>
