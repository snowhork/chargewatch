<template>
  <div>
    UserPage
    <p>{{userID}}</p>
    <DeviceList :userID="userID" :devices="devices">
    </DeviceList>
  </div>
</template>

<script lang="ts">
  import {Component, Prop, Vue} from 'vue-property-decorator';
  import DeviceList from "../components/DeviceList.vue";
  import {NavigationGuardNext, Route} from "vue-router";
  import {GetDevices} from "@/lib/request";
  import {Device} from "@/vals/vals";

  Component.registerHooks([
    'beforeRouteUpdate'
  ])

  @Component({
      components: {DeviceList}
  })

  export default class User extends Vue {
    private userID = ""
    private devices: Array<Device> = []

    async created() {
      await this.fetchData(this.$route);
    }

    async beforeRouteUpdate(
      to: Route,
      from: Route,
      next: NavigationGuardNext) {
      await this.fetchData(to);
      next();
    }

    async fetchData(route: Route) {
      this.userID = route.params.userID

      const resp = await GetDevices(this.userID)
      this.devices = resp.devices
    }
  }
</script>

<style scoped>

</style>