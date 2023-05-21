<template>
    <div class="time-line">
        <div class="year">{{ year(list[0].createTime) }}</div>
        <ul>
            <li :key="item.id" v-for="item in list">
                <span class="date">{{ date(item.createTime) }}</span>
                <span class="title" @click="gotoBlog(item.id)">{{ item.title }}</span>
            </li>
        </ul>
    </div>
</template>

<script>

import dayjs from 'dayjs'

export default {
    name: "TimeAxis",
    props: ["list"],
    methods: {
        year(t) {
            return dayjs(t).format('YYYY')
        },
        date(t) {
            return dayjs(t).format('MM-DD')
        },
        gotoBlog(id) {
            this.$router.push({
                path: "/blogDetail",
                query: {
                    id: id
                }
            });
        }
    }
}
</script>

<style scoped>

ul, li {
   padding: 0;
    margin: 0;
}

.year {
    font-size: 24px;
    font-weight: 700;
    margin: 20px 0 15px;
    line-height: 1.5;
    font-family: "PingFang SC", "Microsoft YaHei", Lato, sans-serif;
}


ul li {
    position: relative;
    list-style: none;
    margin-bottom: 10px;
    line-height: 28px;
    padding-left: 35px;
}

ul li::before {
    content: '';
    width: 7px;
    height: 7px;
    background-color: #565ca4;
    border-radius: 50%;
    display: inline-block;
    position: absolute;
    left: 10px;
    top: 13px;
}

ul li .date {
    margin-right: 20px;
    color: #525bba;
    font-size: 14px;
}

ul li .title {
    position: relative;
    font-size: 16px;
    font-weight: 400;
    font-family: '华文中宋';
    cursor: pointer;
    padding-bottom: 6px;
}

ul li .title::before {
    content: '';
    position: absolute;
    width: 0;
    height: 0;
    bottom: 0;
    transition: 0.4s width linear;
}

ul li .title:hover {
    color: #3d3952;
}

ul li .title:hover::before {
    width: 100%;
    border-bottom: 1px solid #3d3952;
}

</style>
