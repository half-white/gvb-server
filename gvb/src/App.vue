<template>
    <button v-html="home"></button>
</template>

<script>
export default {
    data() {
        return {
            home:"<a href='http://127.0.0.1:8082/hello'>后端接口</a>"
        }
    },
}

</script>