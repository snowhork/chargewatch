<template>
  <div class="home">
    <p>
      Hello, {{ userID }}
    </p>

    <ul>
      <li v-for="device in devices" :key="device.id">
        {{device.name}}
        <span style="margin-left: 4px"></span>
        {{device.charge.value}}
        <span style="margin-left: 4px"></span>
        {{device.charge.timestamp}}
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
  import {Component, Prop, Vue} from 'vue-property-decorator';
  import {Device} from "@/vals/vals";

  @Component
  export default class DeviceList extends Vue {
    @Prop()
    public userID!: string

    @Prop()
    public devices!: Array<Device>
  }
</script>

<style scoped>

</style>