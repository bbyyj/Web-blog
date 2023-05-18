<!--返回顶部组件-->
<template>
    <transition appear
                name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__bounceInDown"
                leave-active-class="animate__bounceOutUp"
    >
        <div class="back-to-top" @click="toTop">
        </div>
    </transition>

</template>

<script>

import "animate.css"

export default {
    name: "BackToTop",
    methods: {
        toTop() {
            window.scrollTo({top: 0, behavior: 'smooth'})
        }
    }
}
</script>

<style scoped>

.back-to-top {
    cursor: pointer;
    position: fixed;
    right: 80px;
    top: -200px;
    z-index: 999;
    width: 70px;
    height: 900px;
    background: url('data:image/BackToTop.png')
}

.animate__animated{
    animation-duration: 1.5s !important;
}

</style>
