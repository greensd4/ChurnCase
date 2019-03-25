<template>
  <div class="ctable" align="center">
    <ul>
    <li v-for="index in selected">
      <button v-on:click="onChange(index)">
      {{index}}
      </button>
    </li>
  </ul>
  <br>
  <input type="text" v-model="search2" placeholder="enter number of complaints.."/>
    <table>
      <thead>
        <tr>
         <th v-for="column in columns">
              {{ column }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="row in filterSearch">
          <td>{{row.id}}</td>
          <td>{{row.gender}}</td>
          <td>{{row.segment}}</td>
          <td>{{row.years_costumer}}</td>
          <td>{{row.no_of_complaints}}</td>
          <td>{{row.contract_value}}</td>
        </tr>
      </tbody>
  </table>
  </div>
</template>



<script>
import {APIService} from '../APIService';
const apiService = new APIService();

export default {
  name: 'ctable',
  data() {
    return {
    selected: [1,2,3,4,5, 'All'],
        
    search: '',
    search2: '',
    
    reverse: false,

    columns: ['ID', 'Gender','Segment','Years_customer','No_of_complaints','Contract_value'],
    rows:[],
    }
  },

  methods: {
    getCustomers() {
      apiService.getCustomers().then((data) => {
      this.rows = data;
      });
    },
    onChange(seg) {
      this.search = seg;
      if (seg === "All") {
        this.search = '';
      }
    },
  },
  mounted() {
    this.getCustomers();
  },
  computed: {
    filterSearch:function(){
      return this.rows.filter((row) =>{
        if (this.search2 == '') {
          return row.segment.match(this.search);  
        }
        return row.segment.match(this.search) && row.no_of_complaints == this.search2;
      });
    },
  }
}
</script>



<style scoped>

li {
  margin-left: 25%;
}

button {
    width: 10%; 
}
button {
    float: left;
}

ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
}
table {
  font-family: 'Open Sans', sans-serif;
  width: 700px;
  border-collapse: collapse;
  border: 3px solid #44475C;
  margin: 10px 10px 10px 0px;
}

table th {
  text-transform: uppercase;
  text-align: left;
  background: #44475C;
  color: #FFF;
  cursor: pointer;
  padding: 8px;
  min-width: 30px;
}
table th:hover {
        background: #717699;
      }
table td {
  text-align: left;
  padding: 8px;
  border-right: 2px solid #7D82A8;
}
table td:last-child {
  border-right: none;
}
table tbody tr:nth-child(2n) td {
  background: rgb(195, 236, 247);
}
</style>