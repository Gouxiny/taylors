<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
          <el-row>
            <el-col :span="5">
              <el-form-item label="市值">
                <el-input-number placeholder="最小" v-model="searchInfo.marketCapitalMin" :controls="false"></el-input-number>
              </el-form-item>
            </el-col>
            <el-col :span="5">
              <el-form-item label="">
                <el-input-number placeholder="最大"  v-model="searchInfo.marketCapitalMax" :controls="false"></el-input-number>
              </el-form-item>
            </el-col>
          </el-row>
        <el-row>
          <el-col :span="5">
            <el-form-item label="涨幅">
              <el-input-number placeholder="最小"  v-model="searchInfo.percentMin" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.percentMax" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="5">
            <el-form-item label="量比">
              <el-input-number placeholder="最小"  v-model="searchInfo.volume_ratio_min" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.volume_ratio_max" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe :default-sort = "{prop: ['market_capital','percent','volume_ratio','high','limit_down','chg','low','volume','amount','open','last_close'], order: 'descending'}">
      <el-table-column label="名称" min-width="150" prop="symbol"></el-table-column>
      <el-table-column label="市值" min-width="130" prop="market_capital" sortable ></el-table-column>

      <el-table-column label="涨幅" min-width="80" prop="percent" sortable></el-table-column>
      <el-table-column label="量比" min-width="80" prop="volume_ratio" sortable></el-table-column>

<!--      <el-table-column label="涨停" min-width="80" prop="high" sortable></el-table-column>
      <el-table-column label="跌停" min-width="80" prop="limit_down" sortable></el-table-column>-->

      <el-table-column label="最高" min-width="80" prop="chg" sortable></el-table-column>
      <el-table-column label="最低" min-width="80" prop="low" sortable></el-table-column>

      <el-table-column label="成交量" min-width="120" prop="volume" sortable></el-table-column>
      <el-table-column label="成交额" min-width="120" prop="amount" sortable></el-table-column>

<!--      <el-table-column label="今开" min-width="80" prop="open" sortable></el-table-column>
      <el-table-column label="昨收" min-width="80" prop="last_close" sortable></el-table-column>-->
    </el-table>
  </div>
</template>


<script>
  import {
    getTopList
  } from '@/api/stockTop'
  import infoList from '@/components/mixins/infoList'
  export default {
    name: 'Top',
    mixins: [infoList],
    data() {
      return {
        listApi: getTopList,
        searchInfo: {
          marketCapitalMin: undefined,
          marketCapitalMax: undefined,
          percentMin: undefined,
          percentMax: undefined,
          volume_ratio_min: undefined,
          volume_ratio_max: undefined
        },
      }
    },
    methods: {
      //搜索
      onSubmit() {
        this.getTableData()
        setInterval(()=>{
          this.getTableData()
        },10000)
      }
    },
    created(){
      this.getTableData()
    }
  }
</script>
<style scoped lang="scss">
  .button-box {
    padding: 10px 20px;
  }
  .el-button {
    float: right;
  }

  .el-tag--mini {
    margin-left: 5px;
  }
  .warning {
    color: #dc143c;
  }
</style>