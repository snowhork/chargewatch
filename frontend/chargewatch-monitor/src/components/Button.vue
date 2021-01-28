<template>
    <button @click="onClick">MyButton</button>
</template>

<script lang="ts">
  import {Component, Vue} from "vue-property-decorator";
  import {GetDevices} from "@/lib/request";

  @Component
  export default class MyButton extends Vue{
    public async onClick(){
        const userID = this.$route.params.userID
        console.log(await GetDevices(userID))
        console.log("end.")

      }
  }
</script>

<style scoped>

</style>