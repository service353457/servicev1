<template>
  <div id="app">
    <button :style="buttonStyle" @click="toggleButton">
      {{ buttonText }}
    </button>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      isLoading: true,
    };
  },
  computed: {
    buttonText() {
      return this.isLoading ? "Загружать" : "Удалять";
    },
    buttonStyle() {
      return {
        backgroundColor: this.isLoading ? "green" : "red",
        color: "white",
        padding: "10px 20px",
        border: "none",
        borderRadius: "5px",
        cursor: "pointer",
      };
    },
  },
  methods: {
    async toggleButton() {
      await this.changeStatus();
    },
    async checkstatus() {
      try {
        const resp = await axios.get("/checkstatus");
        this.isLoading = resp.data.isLoading;
      } catch (error) {
        console.log(error);
      }
    },
    async changeStatus() {
      try {
        this.isLoading = !this.isLoading;
        const data = {
          is_loading: this.isLoading,
        };

        const resp = await axios.post(
          "/changestatus",
          data
        );

        this.isLoading = resp.data.isLoading;
      } catch (error) {
        console.log(error);
      }
    },
  },
  mounted() {
    this.checkstatus();
  },
};
</script>

<style>
#app {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>