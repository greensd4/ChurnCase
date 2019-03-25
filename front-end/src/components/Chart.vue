<template>
   <div class="chart">
     <ul>
      <br><br>
      <label> number of complaints/ {{pick2}} </label>
    <apexchart align="center" width="800" type="bar" :options="chartOptions" :series="series"></apexchart>
        <li v-for="(index) in selectedBySeg">
      <button v-on:click="cha('', index)">
      {{index}}
      </button>
    </li>
    <br><br>
    <li class="li1" v-for="index in selectedByAtt">
      <button class="button1" v-on:click="cha(index,'')">
        {{index}}
      </button>
      </li>
  </ul>
   </div>
</template>

<script>
import {APIService} from '../APIService';
const apiService = new APIService();

export default {
    name: 'chart',
    data() {
      return {
        chartOptions: {
          chart: {
            id: 'complaints chart'
          },
          xaxis: {
            categories:[],
            title: {
              text: ''
            }
          },
          plotOptions: {
            bar: {
              distributed: true
            }
          },
          yaxis:[{
              title: {
                text: "#Complaints"
              },
          },
         ],  
        },
        series: [{
          name: 'complaints',
          data: []
        }],
        rows:[],
        d: [],
        selectedBySeg: [1,2,3,4,5,'All'],
        selectedByAtt: ["Years_Customer", "Contract_Value"],
        pick: '',
        pick2: 'Years_Customer',
        cat:[],
        colors:[],

      }
    },
  methods: {
    getCustomers() {
      apiService.getCustomers().then((data) => {
      this.rows = data;
      this.cha('',"All");
      });
    },
    
    cha(indx, segment){
      var i = 0;
      if (indx !== '') {
        this.pick2 = indx;
      }
      if (segment !== '') {
        this.pick = segment;
      }
      this.series[0].data = [];
      this.chartOptions.xaxis.categories = [];
      this.d = [];
      this.cat = [];

      if (this.pick === "All"){
        this.pick = '';
      }
      if (this.pick2 === "Years_Customer") {
        for (i = 0; i < this.filterSearch.length; i++) {
          this.d.push({y: this.filterSearch[i].years_costumer, n: this.filterSearch[i].no_of_complaints});
        }
      } else {
          for (i = 0; i < this.filterSearch.length; i++) {
            this.d.push({y: this.filterSearch[i].contract_value, n: this.filterSearch[i].no_of_complaints});
        }
      }
      this.d.sort((a,b)=>{
          return a.y-b.y;
      })
      for (i = 0; i < this.filterSearch.length; i++) {
        this.cat.push(this.d[i].y);
        this.series[0].data.push(this.d[i].n);
      }
      this.fillcolors();
      this.chartOptions = {
        xaxis: {
          categories : this.cat,
          title : {
            text : this.pick2
          }
        },
        fill: {
         colors: this.colors
        }
      }
    },
    fillcolors() {
      this.colors = [];
      for (var i = 0; i < this.filterSearch.length; i++) {
          if (this.filterSearch[i].gender === "M") {
              this.colors[i] = '#EC457A';
            } else {
              this.colors[i] = '#6DDCD7';
            }
      }
    },
  },
  mounted() {
    this.getCustomers();
  },
  computed: {
    filterSearch:function(){
      return this.rows.filter((row) =>{
        return row.segment.match(this.pick)
      });
    },
  },
};
</script>



<style scoped>
li {
  margin-left: 25%;
}
.li1 {
  margin-left: 35%;
}
button {
    width: 10%; 
}
button {
    float: left;
}
.button1 {
    width: 20%; 
}
.button1 {
    float: left;
}
ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
  background-color: rgba(216, 213, 11, 0.133);
}
</style>