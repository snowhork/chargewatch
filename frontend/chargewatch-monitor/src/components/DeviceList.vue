<template>
  <div class="home">
    <p>
      Hello, {{ userID }}
    </p>
    <tr v-for="device in devices" :key="device.id">
      <DeviceListRow :device="device"></DeviceListRow>
    </tr>
  </div>
</template>

<script lang="ts">
  import {Component, Prop, Vue} from 'vue-property-decorator';
  import {Device} from "@/vals/vals";
  import DeviceListRow from "@/components/DeviceListRow.vue";
  @Component({
    components: {DeviceListRow}
  })
  export default class DeviceList extends Vue {
    @Prop()
    public userID!: string

    @Prop()
    public devices!: Array<Device>
  }
</script>

<style scoped>

</style>