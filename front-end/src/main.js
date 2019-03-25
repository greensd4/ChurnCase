import Vue from 'vue'
import App from './App.vue'
import VueResource from 'vue-resource'
const VueApexCharts = require('vue-apexcharts');

Vue.use(VueApexCharts)
Vue.use(VueResource);

Vue.component('apexchart', VueApexCharts)
Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
