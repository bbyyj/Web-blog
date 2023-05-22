
<template>
    <div class="bg">
        <div class="maintitle" align="center">关于我</div>
        <div>
            <transition appear name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__slideInUp">
                <div class="wrap">

                    <ul class="contact">
                        <li><a :href="userInfo.github" target="_blank" class="iconfont icon-github"></a></li>
                        <li><a :href="userInfo.csdn" target="_blank" class="iconfont icon-csdn"></a></li>
                    </ul>

                    <div class="box-wrap">
                        <span class="side-label">关于我</span>
                        <h5 class="title">关于博主</h5>
                        <p class="about-state">
                            自我描述 //可在后端数据库存储
                        </p>
                        <p class="about-state">
                            自我描述 //可在后端数据库存储
                        </p>
                    </div>

                    <div class="box-wrap">
                        <span class="side-label bgc-pink">关于本站</span>
                        <h5 class="title">本站采用的技术</h5>
                        <h6 class="site-intro"><span>前端：</span>Vue + Element Ui</h6>
                        <h6 class="site-intro"><span>后端：</span>Golang + Gin + Mysql</h6>
                        <h6 class="site-intro"><span>部署方式：</span>提供了Docker-Compose的部署方式,只需少量配置即可一键启动,使用Nginx作为静态资源和反向代理服务器。
                        </h6>
                        <h6 class="site-intro">本站已经开源在Github</h6>
                    </div>

                    <div class="box-wrap">
                        <span class="side-label bgc-orange">联系方式</span>
                        <h6 class="site-intro" style="margin-top: 12px;"><span>Github:</span><a :href="userInfo.github"
                                target="_blank"> {{ userInfo.github }}</a></h6>
                        <h6 class="site-intro"><span>CSDN:</span><a :href="userInfo.csdn" target="_blank"> </a></h6>
                        <h6 class="site-intro"><span>Email:</span>{{ userInfo.email }}</h6>
                    </div>
                </div>
            </transition>
        </div>
    </div>
</template>



<script>

import "../assets/icon/iconfont.css"
import TitleArea from "../components/TitleArea";

import "animate.css"

export default {
    components: { TitleArea },
    data() {
        return {
            userInfo: {
                avatar: "",
                nickname: "",
                github: "",
                csdn: "",
                email: ""
            },
           
        }
    },
}
</script>


<style lang="less" scoped>

.bg{
    background: url("../assets/images/SeaClouds.svg") 0 0 / cover no-repeat;
}

.maintitle{
    font-size: 48px;
    color: #3d3952;
    margin-bottom: 50px;
    bottom: 0 !important;
    right: 0 !important;
    font-family: "PingFang SC", "Microsoft YaHei", Lato, sans-serif;
    opacity: 0.5;
    padding-top: 6%;
}



.wrap {
    width: 60%;
    border-radius: 10px;
    margin: 0 auto;
    transform: translateY(-20px);
    padding: 50px 0;

    .contact {
        text-align: center;
        margin: 20px 0;
        padding-bottom: 6px;
        position: relative;

        li {
            display: inline-block;

            a {
                text-decoration: none;
                font-size: 26px;
                margin: 0 12px;
                color: rgb(63, 63, 64);
            }
        }
    }


    .box-wrap {
        position: relative;
        width: 80%;
        margin: 30px auto 60px;
        border-radius: 8px;
        background-color: rgba(255, 255, 255, .5);
        box-shadow: -12px -17px 38px rgb(0 0 0 / 10%), 7px 18px 25px rgb(0 0 0 / 8%);
        transition: all .3s ease 0s, transform .6s cubic-bezier(.6, .2, .1, 1) 0s, -webkit-transform .6s cubic-bezier(.6, .2, .1, 1) 0s;
        color: #34495e;
        padding-top: 60px;
        padding-bottom: 18px;

        .side-label {
            position: absolute;
            background-color: #b28fce;
            color: #fff;
            font-style: italic;
            font-size: 14px;
            padding-left: 32px;
            padding-right: 24px;
            left: -16px;
            top: 20px;
            height: 32px;
            line-height: 32px;
            border-radius: 0 3px 3px 0;
        }

        .side-label::before {
            content: '';
            position: absolute;
            top: 100%;
            left: 0;
            width: 0;
            height: 0;
            background-color: transparent;
            border-style: solid;
            border-width: 0 16px 16px 0;
            border-color: transparent;
            border-right-color: #b28fce;
            filter: brightness(120%);
        }
    }

    .box-wrap:hover {
        transform: translateY(-5px);
    }

    .bgc-pink {
        background-color: #f596aa !important;
    }

    .bgc-pink::before {
        border-right-color: #f596aa !important;
    }

    .bgc-orange {
        background-color: #fb966e !important;
    }

    .bgc-orange::before {
        border-right-color: #fb966e !important;
    }

    .title {
        padding: 10px 20px;
        font-size: 20px;
        font-weight: 600;

    }

    font-family: 'Architects Daughter',
    cursive;

    .about-state {
        padding: 0 60px 0 36px;
        text-indent: 2rem;
        font-style: italic;
        font-size: 18px;
        font-weight: 520;
        margin-bottom: 8px;
    }

    .site-intro {
        padding-left: 32px;
        padding-right: 60px;
        font-size: 18px;

        span {
            font-weight: 700;
        }

        a {
            text-decoration: none;
            color: #42b983;
        }
    }


}
</style>
