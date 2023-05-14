<!--博客卡片组件-->
<template>
    <div class="box">

        <!--图片区域-->
        <div class="image-box" :class="{'image-box-right' : imgRight}">
            <img :src="item.firstPicture">
        </div>

        <!--博客区域-->
        <div class="blog-box" :class="{'blog-box-right' : imgRight}">
            <!--博客标题-->
            <div class="blog-title">
                <a>
                    {{ item.title }}
                </a>
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
    // background-image: linear-gradient(120deg, #84fab0 0%, #8fd3f4 100%);
    background-color: #e4ede8;
    width: 840px;
    padding: 16px 20px;
    overflow: hidden;
    border-radius: 0.25rem;
    margin-bottom: 20px;
    position: relative;
    transition: all .5s;
    cursor: pointer;
}

.box:hover {
    transform:  scale(1.02);
    // box-shadow: 0 14px 28px rgba(255, 255, 255, 0.25), 0 10px 10px rgba(255, 255, 255, 0.22);
}

//图片区域
.box .image-box {
    float: left;
    border-radius: 6px;
    margin: 2px 10px 0;
    cursor: pointer;
    overflow: hidden;
}

//博客信息区域
.box .blog-box {
    float: left;
    width: 450px;
    height: 190px;
    padding-left: 20px;
}

.image-box-right {
    float: right !important;
    margin-right: 10px !important;
}

.blog-box-right {
    float: right !important;
    padding-right: 20px !important;
}

.image-box img {
    width: 300px;
    height: 185px;
    border-radius: 6px;
    transition: all 0.5s;
}

.image-box img:hover {
    transform: scale(1.1);
}

.box::after{
    position:absolute;
    content:'';
    // border:3px;
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
    // border-bottom-color:#183f2d;
    // border-left-color:#183f2d;
    transition:border-color 0s ease-out 0.4s, width 0.2s ease-out 0.4s, height 0.2s ease-out 0.6s;
}
.box:hover::before{
    width:100%;
    height:100%;
    // border-top-color:#183f2d;
    // border-right-color: #183f2d;
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
    color: #183f2d;
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
    border-bottom: 2px solid #7f948a;
}

.box .blog-box .blog-title:hover a {
    color: #7f948a;
}

.box .blog-box .blog-desc {
    height: 110px;
    p {
        padding: 12px 20px;
        font-size: 15px;
        line-height: 1.8;
        color: #183f2d;
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
        color: #4b6a5c;
    }

    i {
        font-weight: bold;
    }

    .icon-yonghu{
        color: #98aaa2;
    }

    .icon-riqi {
        color: #98aaa2;
    }

    .icon-yanjing {
        color: #98aaa2;
    }

    a {
        font-size: 13px;
        text-decoration: none;
        outline: none;
        border: 1px solid #4b6a5c;
        padding: 2px;
        border-radius: 3px;
        color: #4b6a5c;
    }
}


li {
    list-style: none;
    display: inline-block;
    margin-right: 18px;
}

.right-float {
    float: right;
}

.left-float {
    float: left;
}

.text-ellipsis {
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 3;
}

</style>
