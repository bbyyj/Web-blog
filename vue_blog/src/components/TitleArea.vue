<template>
    <div class="bg">
        <span class="title">{{ info.title }}</span>
    </div>
</template>

<script>
export default {
    name: "TitleArea",
    props: ["info"]
}
</script>

<style scoped>
.bg {
    height: 220px;
    width: 100%;
    position: relative;
    text-align: center;
    z-index: 0;
    user-select: none;
}

.bg::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1;
}

.title {
    font-size: 48px;
    font-weight: 500;
    color: #d0ced9;
    display: block;
    padding-top: 120px;
}


</style>