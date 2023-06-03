<!--标签云组件-->
<template>
    <div class="wapper">
        <svg class="star star1"></svg>
        <svg class="star star2"></svg>
        <div class="title">标签</div>
        <div class="content">
            <ul>
                <li :key="item.id" v-for="item in tags" @click="itemClicked(item.id)">
                    <span>{{ item.name }}</span>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
export default {
    name: "TagCloud",
    data() {
      return {
          tags: [
              {id:1, name: "Java"},
              {id:2, name: "C++"},
              {id:3, name: "Golang"},
              {id:4, name: "Nginx"},
              {id:5, name: "Docker"},
              {id:6, name: "Kubernetes"},
              {id:7, name: "Mysql"},
              {id:8, name: "Redis"},
              {id:9, name: "MongoDb"},
              {id:10, name: "Gin"}
          ],
      }
    },
    methods: {
        itemClicked(id) {
            this.$router.push({
                path: "/tags",
                query: {
                    id: id
                }
            })
        },
        // 获取所有博客标签
        async getTagList() {
            const {data: res} = await this.$axios.get("/myblog/tagList");
            if(res.status === 1) {
                this.tags = res.data.length > 0 ? res.data[0] : this.tags;
            } else {
                this.$message.warning("获取标签列表失败！")
            }
        },
    },
    created() {
        this.getTagList()
    }
}
</script>

<style scoped>

.wapper {
    width: 290px;
    background-color: #0925f700;
    margin-top: 0px;
    user-select: none;
}

.title {
    color: rgb(255, 255, 255);
    opacity: 0.75;
    text-align: center;
    font-size: 25px;
    margin-bottom: 10px;
}
.content{
    text-align: center;
}
ul li {
    display: inline-block;
    padding: 2px;
    color:#ffffff;
    opacity: 0.7;
    list-style: none;
    margin: 6px 10px;
    cursor: pointer;
    font-weight: 500;
    transition: all 0.3s;
    
}

ul li:hover {
    transform: scale(1.2);
}

.star {
    position: relative;
    background-color: #ffffff;
    clip-path: path('M29.51068,15.41064C19.78412,13.544,18.456,12.21582,16.58929,2.48926a.60015.60015,0,0,0-1.17871,0C13.54437,12.21582,12.21625,13.544,2.4892,15.41064a.60016.60016,0,0,0,0,1.17872c9.72705,1.86669,11.05517,3.19482,12.92138,12.92138a.60027.60027,0,0,0,1.17871,0c1.8667-9.72656,3.19483-11.05469,12.92139-12.92138a.60016.60016,0,0,0,0-1.17872Z');
    width: 40px;
    height: 40px;
    animation: twinkle 2s ease-out infinite alternate;
}
.star1 {  
    position: absolute;
    top: 220px;
    left: 50px;
    animation-direction:alternate-reverse;
}
.star2 {  
    position: absolute;
    top: -10px;
    left: 210px;
}
@keyframes twinkle {
    0% {opacity:0.1;}
    100%{opacity:0.8;}
 }

</style>

