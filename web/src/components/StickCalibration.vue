<template>
  <b-card>
    <h5 slot="header" class="mb-0">Stick Calibration</h5>
    <b-row>
      <b-col sm="12">{{ state.stick_calibration_wizard }}</b-col>
      <!--
      <b-col sm="6">
        <b-row v-for="(v, k) in state.rx" :key="k" class="my-2 mx-2">
          <b-col sm="3">{{ channelNames[k] }}</b-col>
          <b-col sm="9">
            <b-progress
              height="20px"
              :value="v"
              min="-1"
              max="1"
              show-progress
            ></b-progress>
          </b-col>
        </b-row>
      </b-col>
      <b-col sm="6">
        <b-row
          v-for="(v, k) in receiver_stick_calibration_limits"
          :key="k"
          class="my-2 mx-2"
        >
          <b-col sm="3">{{ channelNames[k] }}</b-col>
          <b-col sm="3">
            <b-form-input
              :id="`limi-${channelNames[k]}`"
              type="number"
              step="0.1"
              v-model.number="receiver_stick_calibration_limits[k].min"
            ></b-form-input>
          </b-col>
          <b-col sm="3">
            <b-form-input
              :id="`limi-${channelNames[k]}`"
              type="number"
              step="0.1"
              v-model.number="receiver_stick_calibration_limits[k].max"
            ></b-form-input>
          </b-col>
        </b-row>
      </b-col>
      -->
    </b-row>
  </b-card>
</template>

<script>
import { mapState } from "vuex";
import { mapFields } from "@/store/helper.js";

export default {
  name: "StickCalibration",
  components: {},
  computed: {
    ...mapFields("profile", ["receiver.stick_calibration_limits"]),
    ...mapState(["state"]),
  },
  data() {
    return {
      channelNames: ["Roll", "Pitch", "Yaw", "Throttle"],
    };
  },
};
</script>
