<template>
    <div id="app">
        <router-view></router-view>
    </div>
    <Background />
</template>

<script>
import Background from './components/BackGround.vue';

export default {
    name: 'App',
    components: {
        Background,
    }
}
</script>