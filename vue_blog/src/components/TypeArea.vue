<template>
    <div class="type-area">
        <svg class="star star1"></svg>
        <svg class="star star2"></svg>
        <div class="title">classification</div>
        <ul>
            <li class="li-item" :key="item.id" v-for="item in types">
                <BlogTypeItem @click.native="getBlogList(item.id)" :typeinfo="item" :activeId="currentTypeId"></BlogTypeItem>
            </li>
        </ul>
    </div>

</template>

<script>
import BlogTypeItem from "../components/BlogTypeItem";

export default {
    components: {BlogTypeItem},
    data() {
        return {
            currentTypeId: 0,
            pages: 1,              // 页面数量
            queryInfo: {
                pageNum: 1,
                pageSize: 5,
                typeId: 0
            },
            types: [{
                    id: 1,
                    name: "我的故事",
                    count: 5
                },
                {
                    id: 2,
                    name: "前端",
                    count: 8
                },
                 {
                    id: 3,
                    name: "后端",
                    count: 5
                }
            ],
        }
    },
    created() {
        this.getTypeList();
    },
    methods: {
        getTypeList: async function() {
            const {data: res} = await this.$axios.get("/myblog/typeList");
            if(res.status === 1) {
                this.types = res.data.length > 0 ? res.data[0] : this.types;
                this.types.sort((prev, next) => {
                    return next.count - prev.count
                })
            }

            await this.getBlogList(this.types[0].id);
        },
    },
}
</script>

<style lang="less" scoped>

.li-item {
    font-family:NotoSerifSC-Regular;
    display: inline-block;
    opacity: 0.75;
}
.type-area {
    display: block;
    width: 290px;
    margin: 0 auto;
    text-align: center;
    
}

.title {
    color: rgb(255, 255, 255);
    opacity: 0.75;
    text-align: center;
    font-family: DancingScript-Bold;
    font-size: 28px;
    margin-bottom: 10px;
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
    left: 210px;
    animation-direction:alternate;
}
.star2 {  
    position: absolute;
    top: 120px;
    left: 50px;
    animation-direction:alternate-reverse;
}
@keyframes twinkle {
    0% {opacity:0.1;}
    100%{opacity:0.8;}
}
</style>
