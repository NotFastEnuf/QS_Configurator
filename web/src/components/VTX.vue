<template>
  <b-card>
    <h5 slot="header" class="mb-0">
      VTX
      <b-button size="sm" class="ml-4" v-on:click="fetch_vtx_settings()"
        >refresh</b-button
      >
    </h5>
    <div v-if="vtx.settings.detected">
      <b-row>
        <b-col sm="4" class="my-2">
          <label for="vtx-band">Protocol</label>
        </b-col>
        <b-col sm="8" class="my-2">{{
          protocolNames[vtx.settings.detected]
        }}</b-col>
      </b-row>
      <b-row>
        <b-col sm="4" class="my-2">
          <label for="vtx-band">Frequency</label>
        </b-col>
        <b-col sm="8" class="my-2">{{
          frequencyTable[vtx.settings.band][vtx.settings.channel]
        }}</b-col>
      </b-row>
      <b-row>
        <b-col sm="4" class="my-2">
          <label for="vtx-band">Band</label>
        </b-col>
        <b-col sm="8" class="my-2">
          <b-form-select
            id="vtx-band"
            v-model.number="vtx.settings.band"
            :options="vtxBandOptions"
          ></b-form-select>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="4" class="my-2">
          <label for="vtx-channel">Channel</label>
        </b-col>
        <b-col sm="8" class="my-2">
          <b-form-select
            id="vtx-channel"
            v-model.number="vtx.settings.channel"
            :options="vtxChannelOptions"
          ></b-form-select>
        </b-col>
      </b-row>
      <b-row v-if="vtx.settings.pit_mode != 2">
        <b-col sm="4" class="my-2">
          <label for="vtx-pit-mode">Pit Mode</label>
        </b-col>
        <b-col sm="8" class="my-2">
          <b-form-select
            id="vtx-pit-mode"
            v-model.number="vtx.settings.pit_mode"
            :options="vtxPitModeOptions"
          ></b-form-select>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="4" class="my-2">
          <label for="vtx-power-level">Power Level</label>
        </b-col>
        <b-col sm="8" class="my-2">
          <b-form-select
            id="vtx-power-level"
            v-model.number="vtx.settings.power_level"
            :options="vtxPowerLevelOptions"
          ></b-form-select>
        </b-col>
      </b-row>
      <b-row>
        <b-col offset="9" sm="3">
          <b-button
            class="ml-4 mt-2"
            v-on:click="apply_vtx_settings(vtx.settings)"
            >Apply</b-button
          >
        </b-col>
      </b-row>
    </div>
    <div v-else>Not detected</div>
  </b-card>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  name: "vtx",
  data() {
    return {
      protocolNames: ["INVALID", "TRAMP", "SMARTAUDIO"],
      vtxBandOptions: [
        { value: 0, text: "VTX_BAND_A" },
        { value: 1, text: "VTX_BAND_B" },
        { value: 2, text: "VTX_BAND_E" },
        { value: 3, text: "VTX_BAND_F" },
        { value: 4, text: "VTX_BAND_R" },
      ],
      vtxChannelOptions: [
        { value: 0, text: "VTX_CHANNEL_1" },
        { value: 1, text: "VTX_CHANNEL_2" },
        { value: 2, text: "VTX_CHANNEL_3" },
        { value: 3, text: "VTX_CHANNEL_4" },
        { value: 4, text: "VTX_CHANNEL_5" },
        { value: 5, text: "VTX_CHANNEL_6" },
        { value: 6, text: "VTX_CHANNEL_7" },
        { value: 7, text: "VTX_CHANNEL_8" },
      ],
      vtxPitModeOptions: [
        { value: 0, text: "Off" },
        { value: 1, text: "On" },
        // { value: 2, text: "NO SUPPORT" }
      ],
      vtxPowerLevelOptions: [
        { value: 0, text: "VTX_POWER_LEVEL_1" },
        { value: 1, text: "VTX_POWER_LEVEL_2" },
        { value: 2, text: "VTX_POWER_LEVEL_3" },
        { value: 3, text: "VTX_POWER_LEVEL_4" },
      ],
      frequencyTable: [
        [5865, 5845, 5825, 5805, 5785, 5765, 5745, 5725],
        [5733, 5752, 5771, 5790, 5809, 5828, 5847, 5866],
        [5705, 5685, 5665, 5645, 5885, 5905, 5925, 5945],
        [5740, 5760, 5780, 5800, 5820, 5840, 5860, 5880],
        [5658, 5695, 5732, 5769, 5806, 5843, 5880, 5917],
      ],
    };
  },
  computed: {
    ...mapState(["vtx"]),
  },
  methods: {
    ...mapActions(["apply_vtx_settings", "fetch_vtx_settings"]),
  },
  created() {
    this.fetch_vtx_settings();
  },
};
</script>
