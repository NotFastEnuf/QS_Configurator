<template>
  <b-card>
    <h5 slot="header" class="mb-0">Stick Calibration</h5>
    <b-row>
      <b-col sm="12">
        {{ wizardStates[state.stick_calibration_wizard] }}
      </b-col>
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
    </b-row>
    <b-row>
      <b-col offset="9" sm="3">
        <b-button variant="primary" class="ml-4 mt-2" v-on:click="cal_sticks()">
          Calibrate
        </b-button>
      </b-col>
    </b-row>
  </b-card>
</template>

<script>
import { mapState, mapActions } from "vuex";
import { mapFields } from "@/store/helper.js";

export default {
  name: "StickCalibration",
  components: {},
  computed: {
    ...mapFields("profile", ["receiver.stick_calibration_limits"]),
    ...mapState(["state"]),
  },
  methods: {
    ...mapActions(["cal_sticks"]),
  },
  data() {
    return {
      channelNames: ["Roll", "Pitch", "Yaw", "Throttle"],
      wizardStates: [
        "INACTIVE",
        "CAPTURE_STICKS",
        "WAIT_FOR_CONFIRM",
        "CALIBRATION_FAILED",
        "CALIBRATION_SUCCESS",
        "TIMEOUT",
        "CALIBRATION_CONFIRMED",
      ],
    };
  },
};
</script>
