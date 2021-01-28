<template>
    <div class="home">
        <img alt="Vue logo" src="../assets/logo.png">
        <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/>
        <MyButton></MyButton>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import MyButton from "@/components/Button.vue";

@Component({
  components: {
    MyButton,
  },
})
export default class Home extends Vue {}
</script>
