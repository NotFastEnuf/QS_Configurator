import Vue from 'vue'
import Vuex from 'vuex'

import { get, post } from "@/api.js";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    status: {
      IsConnected: false,
    },
    profile: {
      Setup: {
        InvertYaw: 1,
      },
      Rate: {
        Mode: 0,
        Silverware: {},
        Betaflight: {},
      },
      Voltage: {}
    }
  },
  mutations: {
    set_status(state, status) {
      if (!status.Port) {
        status.Port = status.AvailablePorts[0]
      }
      state.status = status
    },
    set_profile(state, profile) {
      state.profile = profile
    }
  },
  actions: {
    toggle_connection({ dispatch, state }, port) {
      var path = "/api/connect"
      if (state.status.IsConnected) {
        path = "/api/disconnect";
      }
      return post(path, port)
        .then(() => dispatch("fetch_status"))
    },
    fetch_status({ commit, state, dispatch }) {
      return get("/api/status")
        .then(p => {
          if (p.IsConnected && !state.status.IsConnected) {
            dispatch('fetch_profile')
          }
          commit('set_status', p);
        });
    },
    fetch_profile({ commit }) {
      return get("/api/profile")
        .then(p => commit('set_profile', p));
    },
    apply_profile({ commit }, profile) {
      return post("/api/profile", profile)
        .then(p => commit('set_profile', p));
    }
  }
})
