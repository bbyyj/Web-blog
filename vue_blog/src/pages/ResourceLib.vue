<template>
    <div class="bg" >
        <div class="main">
            <div class="maintitle" align="center">资源库</div>
            <div class="butt" style="text-align: center;">
                    <button>全部</button>
                    <button>计算机组成原理</button>
                    <button>计算机网络</button>
            </div>
            <transition appear
                        name="animate__animated animate__bounce animate__slow"
                        enter-active-class="animate__fadeInUp">
                <div class="center-area" style="text-align: center;">
                    <div class="category" style="text-align: left;"  :key="item.id" v-for="item in items" >
                        <strong class="title">
                            {{ item.name }}
                        </strong>
                        <p class = "desc">
                            {{ item.desc }}
                        </p>
                        <p class = "desc">
                            下载总数：{{ item.downloadnum }}
                            大小：{{ item.filesize }}
                        </p>
                        <p class = "desc">
                            上传时间：{{ item.createdat }}
                        </p>
                    </div>
                </div>
            </transition>
        </div>
    </div>

</template>

<script>
import TitleArea from "../components/TitleArea";
import ResourceLabel from "../components/ResourceLabel";

import "animate.css"

export default {
    name: "ResourceLib",
    components: { TitleArea, ResourceLabel },
    data() {
        return {
            items:[]
        }
    },
    methods: {
        async getLinks() {
            const {data: res} = await this.$axios.get("/myblog/t/re/1/10");
            this.items = res.data[0];
            
            if(res.status != 563){
                this.$message.warning("获取资源失败")
                return
            }


            // this.items.shift()
            // let data
            // let items
            // if (res.data.length > 1) {
            //     // data = res.data[0]
            //     items = res.data[1]
            // } else {
            //     return
            // }

            // items.forEach((val) => {
            //     const arr = data.filter((d) => {
            //         return d.name === val.name
            //     })
            //     if (arr.length > 0) {
            //         this.items.push({
            //             // id: val.id,
            //             name: val.name,
            //             // links: arr
            //         })
            //     }
            // })
            console.log(res);

        }
    },
    created() {
        window.scrollTo(0, 0)
        this.getLinks()
    }
}
</script>

<style scoped>

ul, li {
    margin: 0;
    padding: 0;
}

button {
    background-color: #a69ec699;
    margin: 1%;
    padding: 0.2em 1em;
    text-align: center;
    text-decoration: none;
    color: #3d395299;
    border: 2px solid #ffffff00;
    font-size: 20px;
    display: inline-block;
    border-radius: 0.3em;
    transition: all 0.2s ease-in-out;
    position: relative;
    overflow: hidden;
}

button:before {
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

button:hover {
    background-color: #7378acb7;
    color: #fff;
    border-bottom: 4px solid darken(#fff, 10%);
  } 


.maintitle{
    font-size: 48px;
    color: #181720;
    margin-bottom: 50px;
    bottom: 0 !important;
    right: 0 !important;
    font-family: "PingFang SC", "Microsoft YaHei", Lato, sans-serif;
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
    border-radius: 5px;
    padding-bottom: 80px;
}

.category {
    /* padding-left: 30px;
    padding-top: 30px;
    width: 80%;
    background-color: #fcfcfcb1; */
    /* outline: none; */
    display: inline-block;
    border-radius: 7px;
    background-color: #ffffff87;
    width: aoto;
    height: aoto;
    padding: 20px;
    margin-top: 2%;
    margin-right: 2%;
    transition: all 0.3s linear;
}

.desc{
    font-weight: 500;
    color: #979898;
    margin: 0;
    white-space: nowrap;
    overflow: hidden;
    max-width: 250px;
    text-overflow: ellipsis;
}

.title {
    /* height: 40px;
    line-height: 40px;
    font-size: 18px;
    font-weight: 400;
    color: #555;
    position: relative;
    padding-left: 28px; */
    font-weight: 700;
    color: #7378ac;
}

/* .title::before {
    content: '\e611';
    position: absolute;
    left: 5px;
    top: 0px;
    font-family: "iconfont";
    font-size: 20px;
} */

ul {
    padding-left: 20px;
}

ul li {
    float: left;
    list-style: none;
    margin-right: 30px;
    margin-top: 10px;
}


/*清除浮动*/
.clear-fix::after {
    content: '';
    display: block;
    height: 0;
    clear: both;
    visibility: hidden;
}

</style>
