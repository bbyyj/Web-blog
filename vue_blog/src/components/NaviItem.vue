<!--导航按钮组件-->
<template>
    <router-link :to="rout" class="nav-item" :class="{active: isActive}">
        <i class="iconfont" :class="icoClass"></i>
        {{ indexName }}
    </router-link>
</template>

<script>

import "../assets/icon/iconfont.css"

export default {
    name: "NaviItem",
    props: ["indexName", "icoClass", "rout", "activeRoute"],
    computed: {
        isActive() {
          return this.rout === this.activeRoute
        }
    }

}
</script>

<style lang="less" scoped>

i {
    color: #7378ac !important;
}

.nav-item {
    line-height: 40px;
    cursor: pointer;
    padding-bottom: 5px;
    text-decoration:none;
    color: #3d3952;
    font-weight: 40px;
}

.nav-item:hover, .active {
    font-weight: bold;
    color: #c1c3da;
    border-bottom: 2px solid #c1c3da;
}


</style>
