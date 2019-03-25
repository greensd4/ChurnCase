<template>
  <div id="app">
    <img alt="MiCompanyLogo" src="./assets/logo1.png">
    <h1>
      Welcome to MiCompany Churn Explorative Data Module 
    </h1>
    <Table/>
    <Chart/>
  </div>
</template>

<script>
import Table from './components/Table.vue';
import Chart from './components/Chart.vue';

export default {
  name: 'app',
  components: {
    Table,
    Chart
  },
  }
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  align-content: center;
  color: #2c3e50;
  margin-top: 10px;
  background-color: rgba(194, 191, 17, 0.133);
}
</style>
