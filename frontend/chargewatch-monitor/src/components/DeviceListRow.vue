<template>
  <div class="container">
    <div class="item">
      {{name}}
    </div>
    <div class="item value">
      {{value}}
    </div>
    <div class="item charging">
      {{chargingDisplay}}
    </div>
    <div class="item timestamp">
      {{timestampDisplay}} updated
    </div>
  </div>
</template>

<script lang="ts">
  import {Component, Prop, Vue} from "vue-property-decorator";
  import {Device} from "@/vals/vals";

  @Component
  export default class DeviceListRow extends Vue {
    @Prop()
    public userID!: string

    @Prop()
    public device!: Device

    public get name(): string {
      return this.device.name
    }

    public get value(): string {
      return `${this.device.charge.value} %`
    }

    public get chargingDisplay(): string {
      return this.device.charge.charging ? "charging" : "not charging"
    }

    public get timestampDisplay(): string {
      const d = new Date(this.device.charge.timestamp/1000000) // timestamp is nano sec
      return d.toISOString()
    }


  }
</script>

<style scoped>
  .container {
    display: flex;
    font-size: 32px;
    align-items: baseline;
  }

  .item {
    margin-left: 2em;
  }

  .value {
    width: 3.5em;
    text-align: right;
  }

  .charging {
    width: 6em;
  }

  .timestamp {
    font-size: 24px;

  }

</style>