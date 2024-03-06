<script>
export default {
    name: "Background",
}
</script>

<template>
    <div class="background fixed">
        <div class="backdrop fixed"></div>
        <div class="stripe stripe-blue fixed s-1"></div>
        <div class="stripe stripe-blue fixed s-2"></div>
        <div class="stripe stripe-cyan fixed s-3"></div>
        <div class="stripe stripe-cyan fixed s-4"></div>
    </div>
</template>

<style scoped>
.backdrop {
    left: 0;
    top: 30%;
    height: 200%;
    width: 200%;
    background-color: #fff;
    transform: rotate(170deg);
    z-index: 1;
}

.stripe {
    height: 40px;
    z-index: 2;
    opacity: 0.9;
    transform: skew(170deg) rotate(170deg);
    transition: all 1s ease-out;
}

.stripe-blue {
    width: 500px;
    background-color: #409EFF;
}

.stripe-cyan {
    width: 200px;
    background-color: #80E9FF;
}

.s-1 {
    left: -20px;
    top: 56%;
}

.s-2 {
    left: 75%;
    top: 33%;
}

.s-3 {
    left: 10%;
    top: 58%;
}

.s-4 {
    left: 80%;
    top: 31%;
}

.fixed {
    position: fixed;
}


</style>